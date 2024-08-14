package services

import (
	"bytes"
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

type ProductService struct {
}

func boolToint32(b bool) int32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func int32Tobool(n int32) bool {
	if n > 0 {
		return true
	} else {
		return false
	}
}

func (s ProductService) GetProductListCached() ([]*protosService.TPmProductRequest, error) {
	if c := iotredis.GetClient().Get(context.Background(), iotconst.PRODUCT_TYPE_ALL_DATA); c.Err() == nil {
		var data []*protosService.TPmProductRequest
		err := json.Unmarshal([]byte(c.Val()), &data)
		if err != nil {
			goto reload
		}
		if data == nil || len(data) == 0 {
			goto reload
		}
		return data, nil
	}
reload:
	pros, err := rpc.ClientProductService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{})
	if err == nil {
		c := iotredis.GetClient().Set(context.Background(), iotconst.PRODUCT_TYPE_ALL_DATA, iotutil.ToString(pros.List), 0)
		if c.Err() != nil {
			return pros.List, c.Err()
		}
	}
	return pros.List, nil
}

// GetProductTypeMap 获取产品类型类型数据（非产品分类数据）
func (s ProductService) GetProductTypeMap() (map[int64]*protosService.TPmProductRequest, error) {
	res := make(map[int64]*protosService.TPmProductRequest)
	proList, err := s.GetProductListCached()
	if err != nil {
		return res, err
	}
	for _, p := range proList {
		res[p.Id] = p
	}
	return res, nil
}

// CreateProduct create one record
func (s ProductService) CreateProduct(req *entitys.CreateProductForm) (ret int64, err error) {
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = protosService.TPmProductRequest{}
	)

	mapstructure.WeakDecode(req, &data)
	//参数填充
	data.ProductTypeId = iotutil.ToInt64(req.ProductTypeId)
	data.ProductTypeIdPath = req.ProductTypeIdPath
	data.ProductKey = iotutil.GetProductKeyRandomString()
	data.Identifier = iotutil.Uuid()
	data.Model = iotutil.Uuid()
	data.Status = req.IsPublish
	data.Desc = req.Remark
	ctx := context.Background()
	//物模型Id转换
	var (
		thingModelPropertyIds = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		thingModelServiceIds  = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		thingModelEventIds    = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		pIndex                int
		sIndex                int
		eIndex                int
	)
	var thingModels = req.ThingModels
	if req.ThingModels != nil && len(req.ThingModels) > 0 {
		for _, model := range thingModels {
			if model.FuncType == "属性" {
				thingModelPropertyIds = append(thingModelPropertyIds, &protosService.ThingModelInfo{
					ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier:    model.Identifier,
					TriggerCond:   boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				pIndex++
				continue
			}
			if model.FuncType == "服务" {
				thingModelServiceIds = append(thingModelServiceIds, &protosService.ThingModelInfo{ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier: model.Identifier, TriggerCond: boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				sIndex++
				continue
			}
			if model.FuncType == "事件" {
				thingModelEventIds = append(thingModelEventIds, &protosService.ThingModelInfo{ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier: model.Identifier, TriggerCond: boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				eIndex++
				continue
			}
		}
		data.ThingModelPropertyIds = thingModelPropertyIds[0:pIndex]
		data.ThingModelServiceIds = thingModelServiceIds[0:sIndex]
		data.ThingModelEventIds = thingModelEventIds[0:eIndex]
	}

	res, err := rpc.ClientProductService.CreateTPmProduct(ctx, &data)
	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Msg)
	}
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return 0, err
	}

	//设置上传图片对应业务是否成功
	if data.ImageUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmProduct, iotutil.ToString(res.Data.Id), data.ImageUrl)
	}
	return res.Data.Id, err
}

// UpdateProduct edit Product one record
func (s ProductService) UpdateProduct(req *entitys.UpProductForm) (err error) {
	ctx := context.Background()
	if err = req.Valid(); err != nil {
		return
	}
	var (
		data = protosService.TPmProductRequest{}
	)
	mapstructure.WeakDecode(req, &data)
	data.Id = iotutil.ToInt64(req.Id)
	if !iotutil.IsEmpty(req.ProductTypeId) {
		data.ProductTypeId = iotutil.ToInt64(req.ProductTypeId)
	}
	data.ProductTypeId, _ = iotutil.ToInt64AndErr(req.ProductTypeId)
	data.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	data.Desc = req.Remark
	//物模型Id转换
	var (
		thingModelPropertyIds = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		thingModelServiceIds  = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		thingModelEventIds    = make([]*protosService.ThingModelInfo, 0, len(req.ThingModels))
		pIndex                int
		sIndex                int
		eIndex                int
	)
	var thingModels = req.ThingModels
	if req.ThingModels != nil && len(req.ThingModels) > 0 {
		for _, model := range thingModels {
			if model.FuncType == "属性" {
				thingModelPropertyIds = append(thingModelPropertyIds, &protosService.ThingModelInfo{ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier: model.Identifier, TriggerCond: boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				pIndex++
				continue
			}
			if model.FuncType == "服务" {
				thingModelServiceIds = append(thingModelServiceIds, &protosService.ThingModelInfo{ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier: model.Identifier, TriggerCond: boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				sIndex++
				continue
			}
			if model.FuncType == "事件" {
				thingModelEventIds = append(thingModelEventIds, &protosService.ThingModelInfo{ThingModelIds: iotutil.ToInt64(model.Id),
					Identifier: model.Identifier, TriggerCond: boolToint32(model.TriggerCond), ExecCond: boolToint32(model.ExecCond)})
				eIndex++
				continue
			}
		}
		data.ThingModelPropertyIds = thingModelPropertyIds[0:pIndex]
		data.ThingModelServiceIds = thingModelServiceIds[0:sIndex]
		data.ThingModelEventIds = thingModelEventIds[0:eIndex]
	}
	res, err := rpc.ClientProductService.UpdateTPmProduct(ctx, &data)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Msg)
	}
	//services.SetDefaultTranslate(ctx, "t_pm_product", data.Id, "name", data.Name, data.NameEn)
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.PRODUCT_TYPE_DATA, req.Id)); err != nil {
		return err
	}
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return err
	}

	//设置上传图片对应业务是否成功
	if data.ImageUrl != "" {
		commonGlobal.SetAttachmentStatus(model.TableNameTPmProduct, iotutil.ToString(res.Data.Id), data.ImageUrl)
	}
	return nil
}

// GetProductBatch get Product list  data
func (s ProductService) GetProductList(filter *entitys.QueryProductForm) (rets []*entitys.TPmProductVo, total int64, err error) {
	var (
		queryObj = &protosService.TPmProductFilter{
			Name:        filter.Name,
			NetworkType: filter.NetworkType,
		}
	)
	if filter.Status == nil {
		queryObj.Status = -1
	} else {
		queryObj.Status = *filter.Status
	}
	if filter.AttributeType == nil {
		queryObj.AttributeType = -1
	} else {
		queryObj.AttributeType = *filter.AttributeType
	}

	if !iotutil.IsEmpty(filter.ProductTypeId) {
		queryObj.ProductTypeId = iotutil.ToInt64(filter.ProductTypeId)
	}

	ret, err := rpc.ClientProductService.ListTPmProduct(context.Background(), &protosService.TPmProductFilterPage{
		Page:     int64(filter.Page),
		Limit:    int64(filter.Limit),
		QueryObj: queryObj,
	})

	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Msg)
	}
	for _, data := range ret.List {
		rets = append(rets, &entitys.TPmProductVo{
			Id:              iotutil.ToString(data.Id),
			Identifier:      data.Identifier,
			ProductKey:      data.ProductKey,
			Name:            data.Name,
			ImageURL:        data.ImageUrl,
			Model:           data.Model,
			ProductTypeName: data.ProductTypeName,
			ProductTypeId:   iotutil.ToString(data.ProductTypeId),
			WifiFlag:        data.WifiFlag,
			NetworkType:     data.NetworkType,
			AttributeType:   data.AttributeType,
			Status:          data.Status,
		})
	}
	return rets, ret.Total, nil
}

type Cond struct {
	TriggerCond int32
	ExecCond    int32
}

func MergeCond(p []*protosService.ThingModelInfo, s []*protosService.ThingModelInfo, e []*protosService.ThingModelInfo) map[string]Cond {
	m := make(map[string]Cond)
	for _, v := range p {
		m[strconv.Itoa(int(v.ThingModelIds))] = Cond{TriggerCond: v.TriggerCond, ExecCond: v.ExecCond}
	}
	for _, v := range s {
		m[strconv.Itoa(int(v.ThingModelIds))] = Cond{TriggerCond: v.TriggerCond, ExecCond: v.ExecCond}
	}
	for _, v := range e {
		m[strconv.Itoa(int(v.ThingModelIds))] = Cond{TriggerCond: v.TriggerCond, ExecCond: v.ExecCond}
	}
	return m
}

// GetProduct get Product one record
func (s ProductService) GetProduct(id string) (res *entitys.TPmProductVo, err error) {
	var (
		ret *protosService.TPmProductResponseObject
	)
	//resv := &entitys.TPmProductVo{}
	//if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.PRODUCT_TYPE_DATA, id), resv); err == nil {
	//	return resv, nil
	//}
	ret, err = rpc.ClientProductService.GetByIdTPmProduct(context.Background(), &protosService.TPmProductFilterById{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Msg)
	}
	if err = mapstructure.WeakDecode(ret.GetData(), &res); err != nil {
		return
	}
	//配网引导转换
	if len(ret.GetData().NetworkGuides) > 0 {
		res.NetworkGuides = make([]*entitys.PmNetworkGuideEntitys, len(ret.GetData().NetworkGuides))
		for i, guide := range ret.GetData().NetworkGuides {
			mapstructure.WeakDecode(guide, &res.NetworkGuides[i])
		}
	}

	res.Remark = ret.GetData().Desc

	//查询标准物模型列表
	var (
		thingModelStandardVos []*entitys.TPmThingModelVo
		index                 int
	)
	mCond := MergeCond(ret.Data.ThingModelPropertyIds, ret.Data.ThingModelServiceIds, ret.Data.ThingModelEventIds)
	modelid, err := strconv.Atoi(ret.GetData().ModelId)
	if err != nil {
		return nil, err
	}
	thingModelStandardVos, err = s.GetThingModelDetail(int64(modelid))
	if err != nil {
		return res, nil
	}
	thingModelVos := make([]*entitys.TPmThingModelVo, len(thingModelStandardVos))
	for _, vo := range thingModelStandardVos {
		if strings.Index(ret.GetData().Identifiers, vo.Identifier) != -1 {
			if con, ok := mCond[vo.Id]; ok {
				vo.TriggerCond = intTobool(con.TriggerCond)
				vo.ExecCond = intTobool(con.ExecCond)
			}
			thingModelVos[index] = vo
			index++
		}
	}
	res.ThingModels = thingModelVos[0:index]
	//if err := cached.RedisStore.Set(persist.GetRedisKey(iotconst.PRODUCT_TYPE_DATA, id), res, 0); err != nil {
	//	return nil, err
	//}
	return res, err
}

func intTobool(n int32) bool {
	if n > 0 {
		return true
	}
	return false
}

// GetProduct get Product one record
func (s ProductService) GetDefaultNetworkGuides() (res []*entitys.PmNetworkGuideEntitys, err error) {
	var (
		defaultApNetworkGuide  = new(entitys.PmNetworkGuideEntitys)
		defaultEzNetworkGuide  = new(entitys.PmNetworkGuideEntitys)
		defaultBleNetworkGuide = new(entitys.PmNetworkGuideEntitys)
		steps                  = make([]*entitys.PmNetworkGuideStepEntitys, 3)
	)
	//配网类型[0:AP配网,1:EZ配网,2:蓝牙配网]
	defaultApNetworkGuide.Type = 1
	defaultEzNetworkGuide.Type = 2
	defaultBleNetworkGuide.Type = 3
	//配网步骤
	for i := 0; i < 3; i++ {
		step := new(entitys.PmNetworkGuideStepEntitys)
		if i == 0 {
			step.Instruction = "配网步骤一：插上电源"
			step.InstructionEn = "Step 1 of distribution network: plug in the power supply"
			step.ImageUrl = "http://dummyimage.com/400x400"
			step.VideoUrl = "http://bfgn.ph/djomk"
			step.Sort = 1
		}
		if i == 1 {
			step.Instruction = "配网步骤二：按AP配网开关进入配网模式"
			step.InstructionEn = "Step 2 of distribution network: press the AP distribution network switch to enter the distribution network mode"
			step.ImageUrl = "http://dummyimage.com/400x400"
			step.VideoUrl = "http://lwwvm.ph/rjjvi"
			step.Sort = 2
		}
		if i == 2 {
			step.Instruction = "配网步骤三：配网成功"
			step.InstructionEn = "Distribution network step 3: distribution network success"
			step.ImageUrl = "http://dummyimage.com/400x400"
			step.VideoUrl = "http://bfgn.ph/djomk"
			step.Sort = 3
		}
		steps[i] = step
	}
	defaultApNetworkGuide.Steps = steps
	defaultEzNetworkGuide.Steps = steps
	defaultBleNetworkGuide.Steps = steps

	//组装返回数据
	res = make([]*entitys.PmNetworkGuideEntitys, 3)
	res[0] = defaultApNetworkGuide
	res[1] = defaultEzNetworkGuide
	res[2] = defaultBleNetworkGuide
	return res, err
}

// DelProduct delete Product one record
func (s ProductService) DelProduct(id int64) (err error) {
	var (
		data = protosService.TPmProductRequest{}
	)
	data.Id = id
	ret, err := rpc.ClientProductService.DeleteTPmProduct(context.Background(), &data)
	if err != nil {
		return err
	}
	if ret != nil && ret.Code != 200 {
		return errors.New(ret.Msg)
	}
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.PRODUCT_TYPE_DATA, iotutil.ToString(id))); err != nil {
		return err
	}
	if err := cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA); err != nil {
		return err
	}
	return err
}

func (s ProductService) GetThingModelDetail(modelid int64) (resp []*entitys.TPmThingModelVo, err error) {

	var (
		thingModelPropertiesFilter = &protosService.TPmThingModelPropertiesFilter{
			ModelId: modelid,
		}
		thingModelServicesFilter = &protosService.TPmThingModelServicesFilter{
			ModelId: modelid,
		}
		thingModelEventsFilter = &protosService.TPmThingModelEventsFilter{
			ModelId: modelid,
		}
		propertyRes = new(protosService.TPmThingModelPropertiesResponseList)
		serviceRes  = new(protosService.TPmThingModelServicesResponseList)
		eventRes    = new(protosService.TPmThingModelEventsResponseList)
	)

	//查询物模型属性
	var cap = 0
	propertyRes, err = rpc.ClientThingModelPropertiesService.ListTPmThingModelProperties(context.Background(), &protosService.TPmThingModelPropertiesFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelPropertiesFilter,
	})
	if propertyRes != nil && len(propertyRes.List) > 0 {
		cap += len(propertyRes.List)
	}

	//查询物模型服务
	serviceRes, err = rpc.ClientThingModelServicesService.ListTPmThingModelServices(context.Background(), &protosService.TPmThingModelServicesFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelServicesFilter,
	})
	if serviceRes != nil && len(serviceRes.List) > 0 {
		cap += len(serviceRes.List)
	}

	//查询物模型事件
	eventRes, err = rpc.ClientThingModelEventsService.ListTPmThingModelEvents(context.Background(), &protosService.TPmThingModelEventsFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelEventsFilter,
	})
	if eventRes != nil && len(eventRes.List) > 0 {
		cap += len(eventRes.List)
	}

	//合并&转换物模型属性/方法/事件
	var thingModelVoResp = make([]*entitys.TPmThingModelVo, cap)
	if err = s.mergePropertyAttributeDesc(propertyRes, serviceRes, eventRes, thingModelVoResp); err != nil {
		return nil, err
	}

	return thingModelVoResp, err
}

// GetThingModelList get ThingModel list data
func (s ProductService) GetStandardThingModelDetail(productTypeId string) (resp []*entitys.TPmThingModelVo, err error) {
	var (
		_this = new(ProductService)
	)
	//查询物模型基础信息
	thingModelObj, err := rpc.ClientThingModelService.GetTPmThingModel(context.Background(), &protosService.TPmThingModelFilter{
		ProductTypeId: iotutil.ToInt64(productTypeId),
		Version:       "V1.0.0",
		Standard:      1,
	})
	var (
		thingModel *protosService.TPmThingModelRequest
	)
	thingModel = thingModelObj.GetData()
	if thingModelObj.GetData() == nil {
		err = fmt.Errorf("所属品类不存在标准物模型，ProductTypeId：%s，Version：%s", productTypeId, "V1.0.0")
		return nil, err
	}

	var (
		thingModelPropertiesFilter = &protosService.TPmThingModelPropertiesFilter{
			ModelId: thingModel.Id,
		}
		thingModelServicesFilter = &protosService.TPmThingModelServicesFilter{
			ModelId: thingModel.Id,
		}
		thingModelEventsFilter = &protosService.TPmThingModelEventsFilter{
			ModelId: thingModel.Id,
		}
		propertyRes = new(protosService.TPmThingModelPropertiesResponseList)
		serviceRes  = new(protosService.TPmThingModelServicesResponseList)
		eventRes    = new(protosService.TPmThingModelEventsResponseList)
	)

	//查询物模型属性
	var cap = 0
	propertyRes, err = rpc.ClientThingModelPropertiesService.ListTPmThingModelProperties(context.Background(), &protosService.TPmThingModelPropertiesFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelPropertiesFilter,
	})
	if propertyRes != nil && len(propertyRes.List) > 0 {
		cap += len(propertyRes.List)
	}

	//查询物模型服务
	serviceRes, err = rpc.ClientThingModelServicesService.ListTPmThingModelServices(context.Background(), &protosService.TPmThingModelServicesFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelServicesFilter,
	})
	if serviceRes != nil && len(serviceRes.List) > 0 {
		cap += len(serviceRes.List)
	}

	//查询物模型事件
	eventRes, err = rpc.ClientThingModelEventsService.ListTPmThingModelEvents(context.Background(), &protosService.TPmThingModelEventsFilterPage{
		Page:     0,
		Limit:    100,
		QueryObj: thingModelEventsFilter,
	})
	if eventRes != nil && len(eventRes.List) > 0 {
		cap += len(eventRes.List)
	}

	//合并&转换物模型属性/方法/事件
	var thingModelVoResp = make([]*entitys.TPmThingModelVo, cap)
	if err = _this.mergePropertyAttributeDesc(propertyRes, serviceRes, eventRes, thingModelVoResp); err != nil {
		return nil, err
	}

	return thingModelVoResp, err
}

func (s ProductService) GetProductThingModelDetail(prodcutKey string) (resp []*entitys.TPmThingModelVo, err error) {
	var (
		_this = new(ProductService)
	)
	//查询物模型基础信息
	var (
		thingModelPropertiesFilter = &protosService.TPmThingModelPropertiesFilter{
			ProductKey: prodcutKey,
		}
		thingModelServicesFilter = &protosService.TPmThingModelServicesFilter{
			ProductKey: prodcutKey,
		}
		thingModelEventsFilter = &protosService.TPmThingModelEventsFilter{
			ProductKey: prodcutKey,
		}
		propertyRes = new(protosService.TPmThingModelPropertiesResponseList)
		serviceRes  = new(protosService.TPmThingModelServicesResponseList)
		eventRes    = new(protosService.TPmThingModelEventsResponseList)
	)

	//查询物模型属性
	var cap = 0
	propertyRes, err = rpc.ClientThingModelPropertiesService.ListTPmThingModelProperties(context.Background(), &protosService.TPmThingModelPropertiesFilterPage{
		Page:     0,
		Limit:    200,
		QueryObj: thingModelPropertiesFilter,
	})
	if propertyRes != nil && len(propertyRes.List) > 0 {
		cap += len(propertyRes.List)
	}

	//查询物模型服务
	serviceRes, err = rpc.ClientThingModelServicesService.ListTPmThingModelServices(context.Background(), &protosService.TPmThingModelServicesFilterPage{
		Page:     0,
		Limit:    200,
		QueryObj: thingModelServicesFilter,
	})
	if serviceRes != nil && len(serviceRes.List) > 0 {
		cap += len(serviceRes.List)
	}

	//查询物模型事件
	eventRes, err = rpc.ClientThingModelEventsService.ListTPmThingModelEvents(context.Background(), &protosService.TPmThingModelEventsFilterPage{
		Page:     0,
		Limit:    200,
		QueryObj: thingModelEventsFilter,
	})
	if eventRes != nil && len(eventRes.List) > 0 {
		cap += len(eventRes.List)
	}

	//合并&转换物模型属性/方法/事件
	var thingModelVoResp = make([]*entitys.TPmThingModelVo, cap)
	if err = _this.mergePropertyAttributeDesc(propertyRes, serviceRes, eventRes, thingModelVoResp); err != nil {
		return nil, err
	}

	return thingModelVoResp, err
}

// 获取产品类型的故障物模型数据
func (s ProductService) GetProductFaultThingModel(prodcutKey string) (resp []*entitys.TPmThingModelVo, err error) {
	var (
		_this                      = new(ProductService)
		thingModelPropertiesFilter = &protosService.TPmThingModelPropertiesFilter{
			ProductKey: prodcutKey,
			DataType:   iotconst.Dict_type_data_type,
		}
		propertyRes = new(protosService.TPmThingModelPropertiesResponseList)
	)
	//查询物模型属性
	propertyRes, err = rpc.ClientThingModelPropertiesService.ListTPmThingModelProperties(context.Background(), &protosService.TPmThingModelPropertiesFilterPage{
		QueryObj: thingModelPropertiesFilter,
	})
	//合并&转换物模型属性/方法/事件
	var thingModelVoResp = make([]*entitys.TPmThingModelVo, 0)
	if err = _this.mergePropertyAttributeDesc(propertyRes, nil, nil, thingModelVoResp); err != nil {
		return nil, err
	}
	return thingModelVoResp, err
}

// 合并物模型属性/方法/事件
func (ProductService) mergePropertyAttributeDesc(propertyRes *protosService.TPmThingModelPropertiesResponseList, serviceRes *protosService.TPmThingModelServicesResponseList, eventRes *protosService.TPmThingModelEventsResponseList, thingModelVoResp []*entitys.TPmThingModelVo) error {
	var (
		_this = new(ProductService)
		attr  string
		index int
		err   error
	)
	//合并物模型-属性/服务/事件
	if propertyRes != nil && len(propertyRes.List) > 0 {
		for i, property := range propertyRes.List {
			entity := &entitys.TPmThingModelVo{
				Id:          iotutil.ToString(property.Id),
				FuncType:    "属性",
				FuncName:    property.Name,
				Required:    property.Required,
				Identifier:  property.Identifier,
				RwFlag:      property.RwFlag,
				DataType:    property.DataType,
				Desc:        property.Desc,
				TriggerCond: int32Tobool(property.TriggerCond),
				ExecCond:    int32Tobool(property.ExecCond),
				Valid:       int32Tobool(property.Valid),
				Dpid:        property.Dpid,
			}
			attr, err = _this.transformPropertyAttributeDesc(property.GetDataType(), property.GetDataSpecs(), property.GetDataSpecsList())
			if err != nil {
				return err
			}
			entity.Attribute = attr
			thingModelVoResp[i] = entity
			index++
		}
	}
	if serviceRes != nil && len(serviceRes.List) > 0 {
		var initIndex = index
		for i, service := range serviceRes.List {
			entity := &entitys.TPmThingModelVo{
				Id:          iotutil.ToString(service.Id),
				FuncType:    "服务",
				FuncName:    service.ServiceName,
				Required:    service.Required,
				Identifier:  service.Identifier,
				Desc:        service.Desc,
				TriggerCond: int32Tobool(service.TriggerCond),
				ExecCond:    int32Tobool(service.ExecCond),
				Valid:       int32Tobool(service.Valid),
				Dpid:        service.Dpid,
			}
			switch service.CallType {
			case 1:
				entity.Attribute = "调用方式： 异步调用"
			case 0:
				entity.Attribute = "调用方式： 同步调用"
			}
			thingModelVoResp[initIndex+i] = entity
			index++
		}
	}
	if eventRes != nil && len(eventRes.List) > 0 {
		for i, event := range eventRes.List {
			entity := &entitys.TPmThingModelVo{
				Id:          iotutil.ToString(event.Id),
				FuncType:    "事件",
				FuncName:    event.EventName,
				Required:    event.Required,
				Identifier:  event.Identifier,
				Desc:        event.Desc,
				TriggerCond: int32Tobool(event.TriggerCond),
				ExecCond:    int32Tobool(event.ExecCond),
				Valid:       int32Tobool(event.Valid),
				Dpid:        event.Dpid,
			}
			switch event.EventType {
			case "INFO_EVENT_TYPE":
				entity.Attribute = "事件类型：信息"
			case "ALERT_EVENT_TYPE":
				entity.Attribute = "事件类型：告警"
			case "ERROR_EVENT_TYPE":
				entity.Attribute = "事件类型：故障"
			}
			thingModelVoResp[index+i] = entity
		}
	}
	return nil
}

//
////开放平台-合并物模型属性/方法/事件
//func (ProductService) mergeOpenPropertyAttributeDesc(propertyRes *protosService.TPmThingModelPropertiesResponseList, serviceRes *protosService.TPmThingModelServicesResponseList, eventRes *protosService.TPmThingModelEventsResponseList, thingModelVoResp []*entitys.TOpmThingModelVo) error {
//	var (
//		_this = new(ProductService)
//		attr  string
//		index int
//		err   error
//	)
//	//合并物模型-属性/服务/事件
//	if propertyRes != nil && len(propertyRes.List) > 0 {
//		for i, property := range propertyRes.List {
//			entity := &entitys.TOpmThingModelVo{
//				Id:         iotutil.ToString(property.Id),
//				FuncType:   "属性",
//				FuncName:   property.Name,
//				Required:   property.Required,
//				Identifier: property.Identifier,
//				RwFlag:     property.RwFlag,
//				DataType:   property.DataType,
//				Space:      property.DataSpecs,
//			}
//			attr, err = _this.transformPropertyAttributeDesc(property.GetDataType(), property.GetDataSpecs(), property.GetDataSpecsList())
//			if err != nil {
//				return err
//			}
//			entity.Attribute = attr
//			thingModelVoResp[i] = entity
//			index++
//		}
//	}
//	if serviceRes != nil && len(serviceRes.List) > 0 {
//		var initIndex = index
//		for i, service := range serviceRes.List {
//			entity := &entitys.TOpmThingModelVo{
//				Id:           iotutil.ToString(service.Id),
//				FuncType:     "服务",
//				FuncName:     service.ServiceName,
//				Required:     service.Required,
//				Identifier:   service.Identifier,
//				InputParams:  service.InputParams,
//				OutputParams: service.OutputParams,
//			}
//			switch service.CallType {
//			case 1:
//				entity.Attribute = "调用方式： 异步调用"
//			case 0:
//				entity.Attribute = "调用方式： 同步调用"
//			}
//			thingModelVoResp[initIndex+i] = entity
//			index++
//		}
//	}
//	if eventRes != nil && len(eventRes.List) > 0 {
//		for i, event := range eventRes.List {
//			entity := &entitys.TOpmThingModelVo{
//				Id:         iotutil.ToString(event.Id),
//				FuncType:   "事件",
//				FuncName:   event.EventName,
//				Required:   event.Required,
//				Identifier: event.Identifier,
//				Outputdata: event.Outputdata,
//				EventType:  event.EventType,
//			}
//			switch event.EventType {
//			case "INFO_EVENT_TYPE":
//				entity.Attribute = "事件类型：信息"
//			case "ALERT_EVENT_TYPE":
//				entity.Attribute = "事件类型：告警"
//			case "ERROR_EVENT_TYPE":
//				entity.Attribute = "事件类型：故障"
//			}
//			thingModelVoResp[index+i] = entity
//		}
//	}
//	return nil
//}

// 转换属性描述
func (ProductService) transformPropertyAttributeDesc(dataType string, space string, spaceLists string) (string, error) {
	var (
		ret string
		err error
	)
	//按照类型赋值属性
	switch dataType {
	case "INT", "FLOAT", "DOUBLE":
		var (
			data = new(entitys.ValueDataSpecs)
		)
		if err = iotutil.JsonToStruct(space, data); err != nil {
			return "", err
		}
		var rets []string
		if data.Max != nil && data.Min != nil {
			rets = append(rets, fmt.Sprintf("数值范围: %v-%v", data.Min, data.Max))
		}
		if data.Step != nil {
			rets = append(rets, fmt.Sprintf("间距: %v", data.Step))
		}
		if data.Unit != nil {
			rets = append(rets, fmt.Sprintf("单位: %v", data.Unit))
		}
		ret = strings.Join(rets, ", ")
	case "TEXT", "DATE", "JSON":
		var (
			data = new(entitys.StringDataSpecs)
		)
		err = iotutil.JsonToStruct(space, data)

		var rets []string
		if data.Length != nil {
			rets = append(rets, fmt.Sprintf("长度: %v", data.Length))
		}
		if data.DefaultValue != nil {
			rets = append(rets, fmt.Sprintf("默认值: %v", data.DefaultValue))
		}
		ret = strings.Join(rets, ", ")
	case "ENUM", "BOOL", "FAULT":
		var (
			buff bytes.Buffer
		)
		if dataType == "ENUM" {
			buff.WriteString("枚举值：")
		} else if dataType == "FAULT" {
			buff.WriteString("故障值：")
		} else {
			buff.WriteString("布尔值：")
		}

		var slice []map[string]interface{}
		err := json.Unmarshal([]byte(spaceLists), &slice)
		if err != nil {
			return "", err
		}
		for _, v := range slice {
			tmpJson, err := json.Marshal(v)
			if err != nil {
				return ret, err
			}
			enum := entitys.EnumDataSpaces{}
			json.Unmarshal(tmpJson, &enum)
			buff.WriteString("\n")
			buff.WriteString(iotutil.ToString(enum.Value))
			buff.WriteString(" - ")
			buff.WriteString(enum.Name)
		}
		ret = buff.String()
	}
	return ret, err
}

// CheckExists 检查是否存在
func (s ProductService) CheckExists(id string, name string, nameEn string) (bl bool, err error) {
	var idInt int64 = 0
	if id != "" {
		idInt, err = iotutil.ToInt64AndErr(id)
		if err != nil {
			return false, errors.New("非法Id")
		}
	}
	ret, err := rpc.ClientProductService.Exists(context.Background(), &protosService.TPmProductFilter{
		Id:     idInt,
		Name:   name,
		NameEn: nameEn,
	})
	if err != nil {
		return false, err
	}
	if ret != nil && ret.Code != 200 {
		return false, errors.New(ret.Msg)
	}
	return true, nil
}

// AddPmModule 新增模组芯片
func (s ProductService) AddPmModule(req entitys.PmModuleEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.PmModule_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Status = 2
	res, err := rpc.ClientModuleService.Create(context.Background(), saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}
