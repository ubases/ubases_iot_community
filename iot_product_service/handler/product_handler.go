package handler

import (
	"bytes"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_model/db_product/orm"
	"cloud_platform/iot_product_service/convert"
	"cloud_platform/iot_product_service/service"
	_ "cloud_platform/iot_product_service/service"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gen/field"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// The Register tPmProduct handler.
func RegisterTPmProductHandler(service micro.Service) error {
	err := protosService.RegisterTPmProductHandler(service.Server(), new(TPmProductHandler))
	if err != nil {
		iotlogger.LogHelper.Errorf("RegisterTPmProductHandler发生错误:%s", err.Error())
	}
	return err
}

type TPmProductUpdate struct {
	Id                int64  `json:"id"`                //主键（雪花算法19位）
	ProductTypeId     int64  `json:"productTypeId"`     //产品类型ID
	ProductTypeIdPath string `json:"productTypeIdPath"` //产品类型ID
	ProductKey        string `json:"productKey"`        //产品唯一标识
	Name              string `json:"name"`              //产品名称
	NameEn            string `json:"nameEn"`            //产品名称（英文）
	Identifier        string `json:"identifier"`        //属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。
	Model             string `json:"model"`             //产品型号
	ImageUrl          string `json:"imageUrl"`          //产品图片
	WifiFlag          string `json:"wifiFlag"`          //WIFI标识
	NetworkType       int32  `json:"networkType"`       //通信协议（WIFI, BLE, WIFI&#43;BLE）
	AttributeType     int32  `json:"attributeType"`     //设备性质（0:普通设备，1：网关设备）
	Status            int32  `json:"status"`            //状态（0：未发布，1：已发布，2：停用）
	IsVirtualTest     int32  `json:"isVirtualTest"`     //是否支持虚拟测试（0：否，1：是）
	PowerConsumeType  int32  `json:"powerConsumeType,omitempty"`
	Desc              string `json:"desc"`      //描述
	CreatedBy         int64  `json:"createdBy"` //创建人
	UpdatedBy         int64  `json:"updatedBy"` //修改人
}

type TPmProductHandler struct{}

func (s TPmProductHandler) Exists(ctx context.Context, filter *protosService.TPmProductFilter, response *protosService.TPmProductResponse) error {
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	do = do.Where(t.Name.Eq(filter.Name))
	//编辑的时候验证名称是否重复
	if filter.Id != 0 {
		do = do.Where(t.Id.Neq(filter.Id))
	}
	count, err := do.Count()
	if err != nil {
		return err
	}

	if err != nil {
		response.Msg = err.Error()
		return nil
	}
	if count > 0 {
		response.Msg = "产品品类名称已存在"
		return nil
	}
	response.Code = 200
	response.Msg = ""
	return nil
}

func (s TPmProductHandler) UpdateStatus(ctx context.Context, request *protosService.TPmProductStatusRequest, response *protosService.TPmProductResponse) error {
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())

	//检查数据
	if request.Status == 1 {
		err := s.checkData(request.Id)
		if err != nil {
			response.Code = 0
			response.Msg = err.Error()
			return err
		}
	}

	//先查找产品发布状态
	//状态（0：未发布，1：已发布，2：停用）
	//发布：未发布和停用状态下才可以发布，点击，弹窗二次确认是否发布产品，发布后，开放平台才可以创建这个产品类型的产品
	//停用：已发布状态下才可以停用，点击，弹窗二次确认是否停用，停用后，开放平台中将无法新创建这款产品，已创建的产品，不受影响
	var status int32
	err := do.Select(t.Status).Where(t.Id.Eq(request.Id)).Scan(&status)
	if err != nil {
		response.Code = 0
		response.Msg = err.Error()
		return err
	}
	//发布：未发布和停用状态下才可以发布，点击，弹窗二次确认是否发布产品，发布后，开放平台才可以创建这个产品类型的产品
	if request.Status == 1 && status == 1 { //已发布状态，不允许再次发布
		err = errors.New("已发布状态，不允许再次发布")
		response.Code = 0
		response.Msg = err.Error()
		return err
	} else if request.Status == status { //其它情况
		response.Code = 200
		response.Msg = ""
		return nil
	}

	//停用：已发布状态下才可以停用，点击，弹窗二次确认是否停用，停用后，开放平台中将无法新创建这款产品，已创建的产品，不受影响
	if request.Status == 2 && status != 1 { //未发布状态，不允许停用
		err = errors.New("未发布状态，不允许停用")
		response.Code = 0
		response.Msg = err.Error()
		return err
	}

	_, err = do.Where(t.Id.Eq(request.Id)).Update(t.Status, request.Status)
	if err != nil {
		response.Code = 0
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "UpdateTIotDeviceTriad").Error(err)
		return err
	}
	response.Code = 200
	response.Msg = ""
	return nil
}

// GetTPmProduct query struct by struct
func (TPmProductHandler) GetTPmProduct(ctx context.Context, request *protosService.TPmProductFilter, response *protosService.TPmProductResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetTPmProduct request")
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	// 判断参数进行查询
	if request.Id != 0 {
		do = do.Where(t.Id.Eq(request.Id))
	}

	info, err := do.First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetTPmProduct").Error(err)
		return err
	}
	response.Code = 200
	//此方法效率较低，需要尽快找到替换方式
	data := protosService.TPmProductRequest{}
	//mapstructure.WeakDecode(req, &data)
	mapstructure.WeakDecode(info, &data)
	response.Data = &data
	response.Data.ProductTypeIdPath = strings.Split(info.ProductTypeIdPath, ",")
	return nil
}

// ListTPmProduct query list by paging
func (TPmProductHandler) ListTPmProduct(ctx context.Context, request *protosService.TPmProductFilterPage, response *protosService.TPmProductResponseList) error {
	iotlogger.LogHelper.Info("Received Handler.ListTPmProduct request")
	t := orm.Use(iotmodel.GetDB())
	p := t.TPmProduct
	pt := t.TPmProductType

	do := p.WithContext(context.Background()).Select(p.ALL, pt.Name.As("ProductTypeName"))
	// 判断参数进行查询
	if request.QueryObj != nil {
		if request.QueryObj.Id > 0 {
			do = do.Where(p.Id.Eq(request.QueryObj.Id))
		}
		if request.QueryObj.ProductTypeId > 0 {
			do = do.Where(p.ProductTypeId.Eq(int64(request.QueryObj.ProductTypeId)))
		}
		if !iotutil.IsEmpty(request.QueryObj.Name) {
			do = do.Where(p.Name.Like("%" + request.QueryObj.Name + "%"))
		}
		if request.QueryObj.NetworkType != 0 {
			do = do.Where(p.NetworkType.Eq(request.QueryObj.NetworkType))
		}
		if request.QueryObj.AttributeType > 0 {
			do = do.Where(p.AttributeType.Eq(request.QueryObj.AttributeType))
		}
		if request.QueryObj.Status >= 0 {
			do = do.Where(p.Status.Eq(request.QueryObj.Status))
		}
	}
	//if request.SearchKey != "" {
	//    do = do.Or(t.Name.Like(request.QueryObj.Remark), t.Remark.Like(request.QueryObj.Remark))
	//}
	//关联产品品类表
	do = do.Join(pt, p.ProductTypeId.EqCol(pt.Id))

	//排序（update_at倒序）
	orderCol, ok := p.GetFieldByName("updated_at")
	if ok {
		do.Order(orderCol.Desc())
	}

	var (
		list  []*model.TPmProduct
		count int64
		err   error
	)
	hasPage := request.Page != 0 || request.Limit != 0
	if hasPage {
		offset := request.Limit * (request.Page - 1)
		list, count, err = do.FindByPage(int(offset), int(request.Limit))
	} else {
		list, err = do.Find()
	}
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "ListTPmProduct").Error(err)
		return err
	}
	response.Code = 200
	response.Total = count
	response.List = make([]*protosService.TPmProductRequest, len(list))
	for i, _ := range response.List {
		response.List[i] = new(protosService.TPmProductRequest)
		mapstructure.WeakDecode(list[i], &response.List[i])
		response.List[i].ProductTypeIdPath = strings.Split(list[i].ProductTypeIdPath, ",")
	}
	return nil
}

// CreateTPmProduct create
func (s TPmProductHandler) CreateTPmProduct(ctx context.Context, request *protosService.TPmProductRequest, response *protosService.TPmProductResponse) error {
	iotlogger.LogHelper.Info("Received Handler.CreateTPmProduct request")

	//参数判断 request
	if iotutil.IsEmpty(request.Id) {
		return errors.New("id 主键（雪花算法19位） is null")
	}
	if iotutil.IsEmpty(request.ProductTypeId) {
		return errors.New("product_type_id 产品分类Id is null")
	}
	if iotutil.IsEmpty(request.ProductTypeIdPath) {
		return errors.New("ProductTypeIdPath 产品分类Id路径 is null")
	}
	if iotutil.IsEmpty(request.ProductKey) {
		return errors.New("product_key 产品唯一标识 is null")
	}
	if iotutil.IsEmpty(request.Name) {
		return errors.New("name 产品名称 is null")
	}
	if iotutil.IsEmpty(request.NameEn) {
		return errors.New("name_en 产品名称（英文） is null")
	}
	if iotutil.IsEmpty(request.Identifier) {
		return errors.New("identifier 属性的标识符。可包含大小写英文字母、数字、下划线（_），长度不超过50个字符。 is null")
	}
	if iotutil.IsEmpty(request.Model) {
		return errors.New("model 产品型号 is null")
	}
	if iotutil.IsEmpty(request.WifiFlag) {
		return errors.New("wifi_flag WIFI标识 is null")
	}
	if iotutil.IsEmpty(request.NetworkType) {
		return errors.New("network_type 通信协议（WIFI, BLE, WIFI&#43;BLE） is null")
	}
	if iotutil.IsEmpty(request.AttributeType) {
		return errors.New("attribute_type 设备性质（0:普通设备，1：网关设备） is null")
	}
	if iotutil.IsEmpty(request.Status) {
		return errors.New("status 状态（0：未发布，1：已发布，2：停用） is null")
	}
	if iotutil.IsEmpty(request.IsVirtualTest) {
		return errors.New("is_virtual_test 是否支持虚拟测试（0：否，1：是） is null")
	}

	var (
		saveObj = model.TPmProduct{}
		err     error
	)
	//分类名称去重
	isExists, err := s.existsByName(request.Name, request.NameEn, 0)
	if err != nil {
		response.Msg = err.Error()
		return nil
	}
	if isExists {
		response.Msg = "产品品类名称已存在"
		return nil
	}

	defer func() {
		if err != nil {
			response.Msg = err.Error()
			iotlogger.LogHelper.WithTag("func", "CreateTPmProduct").Error(err)
		}
	}()

	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		mapstructure.WeakDecode(request, &saveObj)
		do := tx.TPmProduct.WithContext(context.Background())
		saveObj.Id = iotutil.GetNextSeqInt64()
		request.Id = saveObj.Id
		saveObj.ProductTypeIdPath = strings.Join(request.ProductTypeIdPath, ",")
		err = do.Create(&saveObj)
		if err != nil {
			return err
		}
		//创建物模型
		var (
			_this      = new(TPmProductHandler)
			handle     = new(TPmThingModelHandler)
			now        = time.Now().Format("2006-01-02 15:04:05")
			thingModel = &protosService.TPmThingModelRequest{
				ProductKey:    request.ProductKey,
				ProductTypeId: request.ProductTypeId,
				Standard:      0,
				Version:       "V1.0.0",
				Description:   fmt.Sprintf("【%s】的物模型", request.GetName()),
				CreatedTime:   now,
				UpdatedTime:   now,
			}
		)
		thingsModelRes, err := handle.SaveTPmThingModel(tx, ctx, thingModel)
		if err != nil {
			return err
		}
		//var modelId = iotutil.ToString(thingModelResponse.GetData().Id)

		//更新关联-物模型属性/服务/事件
		err = _this.relationThingModellForCreate(tx, ctx, request, thingsModelRes.Id)
		if err != nil {
			return err
		}

		var (
			networkGuideIds []int64
		)
		//关联固件
		err = _this.relationProductFirmware(tx, request)
		if err != nil {
			logger.Errorf("CreatePmFirmware relationProductFirmware error : %s", err.Error())
			return err
		}
		//关联模组
		err = _this.relationProductModule(tx, request)
		if err != nil {
			logger.Errorf("CreatePmFirmware relationProductModule error : %s", err.Error())
			return err
		}
		//关联控制面板
		err = _this.relationProductPanel(tx, request)
		if err != nil {
			logger.Errorf("CreatePmFirmware relationProductPanel error : %s", err.Error())
			return err
		}
		//关联配网引导
		networkGuideIds, err = _this.relationNetworkGuide(tx, request)
		if err != nil {
			logger.Errorf("CreatePmFirmware relationNetworkGuide error : %s", err.Error())
			return err
		}
		//关联配网引导步骤
		if len(networkGuideIds) > 0 {
			err = _this.relationNetworkGuideStep(tx, request, networkGuideIds)
			if err != nil {
				logger.Errorf("CreatePmFirmware relationNetworkGuideStep error : %s", err.Error())
				return err
			}
		}
		return nil
	})

	if err != nil {
		logger.Errorf("CreatePmFirmware error : %s", err.Error())
		return err
	}

	response.Code = 200
	response.Data = &protosService.TPmProductRequest{Id: saveObj.Id}

	//通知设置多语言
	//service.GetJsPublisherMgr().PushData(&service.NatsPubData{
	//	Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
	//	Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, saveObj.Id, "name", saveObj.Name, saveObj.NameEn),
	//})
	iotnatsjs.GetJsClientPub().PushData(&iotnatsjs.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, saveObj.Id, "name", saveObj.Name, saveObj.NameEn),
	})
	return nil
}

// UpdateTPmProduct update
func (s TPmProductHandler) UpdateTPmProduct(ctx context.Context, request *protosService.TPmProductRequest, response *protosService.TPmProductResponse) error {
	iotlogger.LogHelper.Info("Received Handler.UpdateTPmProduct request")

	var err error
	//更新并发布，检查所有参数是否完整
	if request.Status == 2 {
		productDetails := &protosService.TPmProductResponseObject{}
		err = s.GetByIdTPmProduct(ctx, &protosService.TPmProductFilterById{
			Id: request.Id,
		}, productDetails)
		if err != nil {
			return err
		}
		//检查产品数据是否完整
		//模组、固件配置是否完整
		if productDetails.Data.ModuleIds == nil || len(productDetails.Data.ModuleIds) == 0 {
			return errors.New("请选择产品芯片模组")
		}
		if productDetails.Data.FirmwareIds == nil || len(productDetails.Data.FirmwareIds) == 0 {
			return errors.New("请选择产品固件")
		}
		if productDetails.Data.NetworkGuides == nil || len(productDetails.Data.NetworkGuides) == 0 {
			return errors.New("请选择产品配网方式")
		}
		//面板是否完整
		if productDetails.Data.ControlPanelIds == nil || len(productDetails.Data.ControlPanelIds) == 0 {
			return errors.New("请选择产品面板")
		}
		//物模型是否完整
		if (productDetails.Data.ThingModelPropertyIds == nil || len(productDetails.Data.ThingModelPropertyIds) == 0) &&
			(productDetails.Data.ThingModelServiceIds == nil || len(productDetails.Data.ThingModelServiceIds) == 0) &&
			(productDetails.Data.ThingModelEventIds == nil || len(productDetails.Data.ThingModelEventIds) == 0) {
			return errors.New("请选择产品物模型")
		}
	}
	//
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		err = s.updateBaseInfo(tx, request)
		//service.GetJsPublisherMgr().PushData(&service.NatsPubData{
		//	Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
		//	Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, request.Id, "name", request.Name, request.NameEn),
		//})
		iotnatsjs.GetJsClientPub().PushData(&iotnatsjs.NatsPubData{
			Subject: iotconst.NATS_SUBJECT_LANGUAGE_UPDATE,
			Data:    iotstruct.TranslatePush{}.SetContent(iotconst.LANG_PRODUCT_NAME, request.Id, "name", request.Name, request.NameEn),
		})
		//更新关联-物模型属性/服务/事件 (场景条件）
		err = s.relationThingModelForUpdate(tx, ctx, request)
		if err != nil {
			return err
		}
		//关联控制面板
		err = s.relationProductPanel(tx, request)
		if err != nil {
			return err
		}
		//关联模组
		err = s.relationProductModule(tx, request)
		if err != nil {
			return err
		}
		//关联固件
		err = s.relationProductFirmware(tx, request)
		if err != nil {
			return err
		}
		var networkGuideIds []int64
		//关联配网引导
		networkGuideIds, err = s.relationNetworkGuide(tx, request)
		if err != nil {
			return err
		}
		//关联配网引导步骤
		if len(networkGuideIds) > 0 {
			err = s.relationNetworkGuideStep(tx, request, networkGuideIds)
			if err != nil {
				return err
			}
		}
		if request.Status == 2 {
			s.updateProductStatus(tx, request)
		}
		return nil
	})
	if err != nil {
		logger.Errorf("CreatePmFirmware error : %s", err.Error())
		return err
	}
	response.Code = 200
	response.Data = &protosService.TPmProductRequest{Id: request.Id}
	return nil
}

func (s TPmProductHandler) updateBaseInfo(tx *orm.Query, request *protosService.TPmProductRequest) error {
	t := tx.TPmProduct
	do := t.WithContext(context.Background())
	//分类名称去重
	if request.Name != "" {
		isExists, err := s.existsByName(request.Name, request.NameEn, request.Id)
		if err != nil {
			return err
		}
		if isExists {
			return errors.New("产品品类名称已存在")
		}
	}

	// 赋值参数赋值
	var updateObj = model.TPmProduct{}
	//此方法效率较低，需要尽快找到替换方式
	mapstructure.WeakDecode(request, &updateObj)
	updateObj.ProductTypeIdPath = strings.Join(request.ProductTypeIdPath, ",")

	//修改所有数据，防止0的排除。
	var updateField []field.Expr
	updateField = append(updateField, t.ProductTypeId)
	updateField = append(updateField, t.ProductKey)
	updateField = append(updateField, t.Name)
	updateField = append(updateField, t.NameEn)
	updateField = append(updateField, t.Identifier)
	updateField = append(updateField, t.Model)
	updateField = append(updateField, t.ImageUrl)
	updateField = append(updateField, t.WifiFlag)
	updateField = append(updateField, t.NetworkType)
	updateField = append(updateField, t.AttributeType)
	updateField = append(updateField, t.PowerConsumeType)
	updateField = append(updateField, t.Status)
	updateField = append(updateField, t.IsVirtualTest)
	updateField = append(updateField, t.Desc)
	updateField = append(updateField, t.UpdatedBy)
	updateField = append(updateField, t.ProductTypeName)
	updateField = append(updateField, t.ProductTypeIdPath)
	_, err := do.Select(updateField...).Where(t.Id.Eq(request.Id)).Updates(updateObj)
	if err != nil {
		return err
	}
	return nil
}

func (s TPmProductHandler) updateProductStatus(tx *orm.Query, request *protosService.TPmProductRequest) error {
	t := tx.TPmProduct
	do := t.WithContext(context.Background())
	_, err := do.Where(t.Id.Eq(request.Id)).Update(t.Status, request.Status)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTPmProduct delete
func (TPmProductHandler) DeleteTPmProduct(ctx context.Context, request *protosService.TPmProductRequest, response *protosService.TPmProductResponse) error {
	iotlogger.LogHelper.Info("Received Handler.DeleteTPmProduct request")
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())

	_, err := do.Where(t.Id.Eq(int64(request.Id))).Delete()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "DeleteTPmProduct").Error(err)
		return err
	}
	response.Code = 200
	return nil
}

// GetByIdTPmProduct query struct by id
func (TPmProductHandler) GetByIdTPmProduct(ctx context.Context, request *protosService.TPmProductFilterById, response *protosService.TPmProductResponseObject) error {
	iotlogger.LogHelper.Info("Received Handler.GetByIdTPmProduct request")
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	//创建查询条件
	info, err := do.Where(t.Id.Eq(request.Id)).First()
	if err != nil {
		response.Msg = err.Error()
		iotlogger.LogHelper.WithTag("func", "GetByIdTPmProduct").Error(err)
		return err
	}
	response.Data = new(protosService.TPmProductRequest)
	mapstructure.WeakDecode(info, &response.Data)
	response.Data.ProductTypeIdPath = strings.Split(info.ProductTypeIdPath, ",")

	var (
		thingModelHandle           = new(TPmThingModelHandler)
		propertyHandle             = new(TPmThingModelPropertiesHandler)
		serviceHandle              = new(TPmThingModelServicesHandler)
		eventHandle                = new(TPmThingModelEventsHandler)
		productModuleRelationSvc   = new(service.PmProductModuleRelationSvc)
		productFirmwareRelationSvc = new(service.PmProductFirmwareRelationSvc)
		productPanelRelationSvc    = new(service.PmProductPanelRelationSvc)
		networkGuideSvc            = new(service.PmNetworkGuideSvc)
		networkGuideStepSvc        = new(service.PmNetworkGuideStepSvc)
	)

	//查询物模型
	var (
		thingModelFilter = &protosService.TPmThingModelFilter{
			ProductKey: info.ProductKey,
		}
		thingModelResponse = new(protosService.TPmThingModelResponseObject)
	)
	err = thingModelHandle.GetTPmThingModel(ctx, thingModelFilter, thingModelResponse)
	if err != nil {
		return err
	}
	response.Data.ModelId = iotutil.ToString(thingModelResponse.Data.Id)

	//查询物模型-属性
	var (
		flag                       = false
		thingModelPropertiesFilter = &protosService.TPmThingModelPropertiesFilterPage{
			QueryObj: &protosService.TPmThingModelPropertiesFilter{
				ModelId: thingModelResponse.GetData().GetId(),
			},
		}
		thingModelPropertiesResponseList = new(protosService.TPmThingModelPropertiesResponseList)
	)
	propertyHandle.ListTPmThingModelProperties(ctx, thingModelPropertiesFilter, thingModelPropertiesResponseList)
	if err != nil {
		return err
	}
	response.Data.ThingModelPropertyIds = make([]*protosService.ThingModelInfo, 0, len(thingModelPropertiesResponseList.GetList()))
	identifiers := new(bytes.Buffer)
	for _, propertiesRequest := range thingModelPropertiesResponseList.GetList() {
		response.Data.ThingModelPropertyIds = append(response.Data.ThingModelPropertyIds, &protosService.ThingModelInfo{
			ThingModelIds: propertiesRequest.Id,
			TriggerCond:   propertiesRequest.TriggerCond,
			ExecCond:      propertiesRequest.ExecCond,
			Valid:         propertiesRequest.Valid,
		})
		identifiers.WriteString(",")
		identifiers.WriteString(propertiesRequest.Identifier)
		flag = true
	}

	//查询物模型-服务
	var (
		thingModelServicesFilter = &protosService.TPmThingModelServicesFilterPage{
			QueryObj: &protosService.TPmThingModelServicesFilter{
				ModelId: thingModelResponse.GetData().GetId(),
			},
		}
		thingModelServicesResponseList = new(protosService.TPmThingModelServicesResponseList)
	)
	serviceHandle.ListTPmThingModelServices(ctx, thingModelServicesFilter, thingModelServicesResponseList)
	if err != nil {
		return err
	}
	response.Data.ThingModelServiceIds = make([]*protosService.ThingModelInfo, 0, len(thingModelServicesResponseList.GetList()))
	for _, servicesRequest := range thingModelServicesResponseList.GetList() {
		response.Data.ThingModelServiceIds = append(response.Data.ThingModelServiceIds, &protosService.ThingModelInfo{
			ThingModelIds: servicesRequest.Id,
			TriggerCond:   servicesRequest.TriggerCond,
			ExecCond:      servicesRequest.ExecCond,
			Valid:         servicesRequest.Valid,
		})

		identifiers.WriteString(",")
		identifiers.WriteString(servicesRequest.Identifier)
		flag = true
	}

	//查询物模型-事件
	var (
		thingModelEventsFilter = &protosService.TPmThingModelEventsFilterPage{
			QueryObj: &protosService.TPmThingModelEventsFilter{
				ModelId: thingModelResponse.GetData().GetId(),
			},
		}
		thingModelEventsResponseList = new(protosService.TPmThingModelEventsResponseList)
	)
	eventHandle.ListTPmThingModelEvents(ctx, thingModelEventsFilter, thingModelEventsResponseList)
	if err != nil {
		return err
	}
	response.Data.ThingModelEventIds = make([]*protosService.ThingModelInfo, 0, len(thingModelEventsResponseList.GetList()))
	for _, eventsRequest := range thingModelEventsResponseList.GetList() {
		response.Data.ThingModelEventIds = append(response.Data.ThingModelEventIds, &protosService.ThingModelInfo{
			ThingModelIds: eventsRequest.Id,
			TriggerCond:   eventsRequest.TriggerCond,
			ExecCond:      eventsRequest.ExecCond,
			Valid:         eventsRequest.Valid,
		})

		identifiers.WriteString(",")
		identifiers.WriteString(eventsRequest.Identifier)
		flag = true
	}
	if flag {
		//物模型属性
		response.Data.Identifiers = identifiers.String()[1:]
	}

	//查询模组
	var (
		productModuleRelationFilter = &protosService.PmProductModuleRelationListRequest{
			Query: &protosService.PmProductModuleRelation{
				ProductId: request.GetId(),
			},
		}
	)
	productModuleRelations, _, err := productModuleRelationSvc.GetListPmProductModuleRelation(productModuleRelationFilter)
	if err != nil {
		return err
	}
	if len(productModuleRelations) > 0 {
		response.Data.ModuleIds = make([]int64, len(productModuleRelations))
		for i, relation := range productModuleRelations {
			response.Data.ModuleIds[i] = relation.ModuleId
		}
	}

	//查询固件
	var (
		productFirmwareRelationFilter = &protosService.PmProductFirmwareRelationListRequest{
			Query: &protosService.PmProductFirmwareRelation{
				ProductId: request.GetId(),
			},
		}
	)
	productFirmwareRelations, _, err := productFirmwareRelationSvc.GetListPmProductFirmwareRelation(productFirmwareRelationFilter)
	if err != nil {
		return err
	}
	if len(productFirmwareRelations) > 0 {
		response.Data.FirmwareIds = make([]int64, len(productFirmwareRelations))
		for i, relation := range productFirmwareRelations {
			response.Data.FirmwareIds[i] = relation.FirmwareId
		}
	}

	//查询面板
	var (
		productPanelRelationFilter = &protosService.PmProductPanelRelationListRequest{
			Query: &protosService.PmProductPanelRelation{
				ProductId: request.GetId(),
			},
		}
	)
	productPanelRelations, _, err := productPanelRelationSvc.GetListPmProductPanelRelationAndPanelStatus(productPanelRelationFilter)
	if err != nil {
		return err
	}
	if len(productPanelRelations) > 0 {
		response.Data.ControlPanelIds = make([]int64, len(productPanelRelations))
		for i, relation := range productPanelRelations {
			response.Data.ControlPanelIds[i] = relation.ControlPanelId
		}
	}

	//查询配网引导
	var (
		networkGuideFilter = &protosService.PmNetworkGuideListRequest{
			Query: &protosService.PmNetworkGuide{
				ProductId: request.GetId(),
			},
		}
	)
	networkGuides, _, err := networkGuideSvc.GetListPmNetworkGuide(networkGuideFilter)
	if err != nil {
		return err
	}
	if len(networkGuides) > 0 {
		response.Data.NetworkGuides = make([]*protosService.PmNetworkGuideObj, len(networkGuides))
		for i, guide := range networkGuides {
			mapstructure.WeakDecode(guide, &response.Data.NetworkGuides[i])
		}
	}

	//查询配网引导步骤
	if len(networkGuides) > 0 {
		for i, guide := range networkGuides {
			networkGuideSteps, _, err := networkGuideStepSvc.GetListPmNetworkGuideStep(&protosService.PmNetworkGuideStepListRequest{
				Query: &protosService.PmNetworkGuideStep{
					NetworkGuideId: guide.Id,
				},
			})
			if err != nil {
				return err
			}
			if len(networkGuideSteps) > 0 && response.Data.NetworkGuides[i] != nil {
				response.Data.NetworkGuides[i].Steps = make([]*protosService.PmNetworkGuideStepObj, len(networkGuideSteps))
				for j, step := range networkGuideSteps {
					response.Data.NetworkGuides[i].Steps[j] = new(protosService.PmNetworkGuideStepObj)
					mapstructure.WeakDecode(step, response.Data.NetworkGuides[i].Steps[j])
				}
			}
		}
	}
	response.Code = 200
	return nil
}

// CreateTPmProduct create
func (TPmProductHandler) UploadControlPanel(ctx context.Context, request *protosService.PmControlPanelObj, response *protosService.TPmProductResponse) error {
	//var (
	//	saveObj = model.TPmControlPanel{}
	//)
	//t := orm.Use(iotmodel.GetDB()).TPmControlPanels
	//do := t.WithContext(context.Background())
	//mapstructure.WeakDecode(request, &saveObj)
	//err := do.Create(&saveObj)
	//if err != nil {
	//	logger.Errorf("UploadControlPanel error : %s", err.Error())
	//	return err
	//}
	//
	//response.Code = 200
	//response.Data = &protosService.TPmProductRequest{Id: saveObj.Id}
	return nil
}

func (TPmProductHandler) ControlPanelsLists(ctx context.Context, req *protosService.ControlPanelIdsRequest, opts ...client.CallOption) (*protosService.PmControlPanelsVoResponse, error) {
	var err error
	t := orm.Use(iotmodel.GetDB()).TPmControlPanels
	do := t.WithContext(context.Background())
	do = do.Where(t.Id.In(req.ControlPanelIds...))

	orderCol, ok := t.GetFieldByName("updated_at")
	if !ok {
		orderCol = t.CreatedAt
	}
	do = do.Order(orderCol)
	list, err := do.Find()

	if err != nil {
		logger.Errorf("ControlPanelsLists error : %s", err.Error())
		return nil, err
	}
	if len(list) == 0 {
		return nil, err
	}
	result := make([]*protosService.PmControlPanels, len(list))
	for i, v := range list {
		result[i] = convert.PmControlPanels_db2pb(v)
	}
	var (
		data = make([]*protosService.PmControlPanelsVo, len(result))
	)
	mapstructure.WeakDecode(result, &data)
	return &protosService.PmControlPanelsVoResponse{Data: data}, err
}

func (TPmProductHandler) ModuleLists(ctx context.Context, req *protosService.ModuleIdsRequest, opts ...client.CallOption) (*protosService.PmModuleVoResponse, error) {
	var err error
	t := orm.Use(iotmodel.GetDB()).TPmModule
	do := t.WithContext(context.Background())
	do = do.Where(t.Id.In(req.ModuleIds...))

	orderCol, ok := t.GetFieldByName("updated_at")
	if !ok {
		orderCol = t.CreatedAt
	}
	do = do.Order(orderCol)
	list, err := do.Find()

	if err != nil {
		logger.Errorf("ModuleLists error : %s", err.Error())
		return nil, err
	}
	if len(list) == 0 {
		return nil, err
	}
	result := make([]*protosService.PmModule, len(list))
	for i, v := range list {
		result[i] = convert.PmModule_db2pb(v)
	}
	var (
		data = make([]*protosService.PmModuleVo, len(result))
	)
	mapstructure.WeakDecode(result, &data)
	return &protosService.PmModuleVoResponse{Data: data}, err
}

//func GetModelItemIds(modellist []*protosService.ThingModelInfo) []int64 {
//	ids := make([]int64, len(modellist))
//	for i, v := range modellist {
//		ids[i] = v.ThingModelIds
//	}
//	return ids
//}

type Cond struct {
	ExecCond    int32
	TriggerCond int32
}

// 关联物模型
/*
func (TPmProductHandler) relationThingModel(ctx context.Context, request *protosService.TPmProductRequest, modelId string) error {
	var (
		err error
		now = time.Now().Format("2006-01-02 15:04:05")
	)
	//物模型-属性
	if len(request.ThingModelPropertyIds) > 0 {
		var (
			propertyHandle = new(TPmThingModelPropertiesHandler)
			deleteResp     = new(protosService.TPmThingModelPropertiesResponse)
			deleteReq      = &protosService.TPmThingModelPropertiesRequest{
				ModelId:    iotutil.ToInt64(modelId),
				ProductKey: request.ProductKey,
			}
		)
		mapCond := make(map[int64]Cond)
		for _, v := range request.ThingModelPropertyIds {
			mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
		}
		//删除物模型-属性关系
		err = propertyHandle.DeleteTPmThingModelProperties(ctx, deleteReq, deleteResp)
		if err != nil {
			return err
		}

		var (
			thingModelPropertiesRequests   []*protosService.TPmThingModelPropertiesRequest
			thingModelPropertyResponseList = new(protosService.TPmThingModelPropertiesResponseList)
			propertyeQueryObj              = &protosService.TPmThingModelPropertiesFilter{
				Ids: GetModelItemIds(request.ThingModelPropertyIds),
			}
		)
		//创建物模型-属性关系
		err = propertyHandle.ListTPmThingModelProperties(ctx, &protosService.TPmThingModelPropertiesFilterPage{QueryObj: propertyeQueryObj}, thingModelPropertyResponseList)
		if err != nil {
			return err
		}
		thingModelPropertiesRequests = thingModelPropertyResponseList.GetList()
		for _, value := range thingModelPropertiesRequests {
			if con, ok := mapCond[value.Id]; ok {
				value.ExecCond = con.ExecCond
				value.TriggerCond = con.TriggerCond
			}
			value.Id = iotutil.GetNextSeqInt64()
			value.ModelId = iotutil.ToInt64(modelId)
			value.ProductKey = request.ProductKey
			value.Custom = 1
			value.CreatedTime = now
			value.UpdatedTime = now
		}
		var (
			thingModelPropertyResponse = new(protosService.TPmThingModelPropertiesResponse)
		)
		err = propertyHandle.CreateTPmThingModelPropertiesBatch(ctx, thingModelPropertiesRequests, thingModelPropertyResponse)
		if err != nil {
			return err
		}
	}

	//物模型-服务
	if len(request.ThingModelServiceIds) > 0 {
		var (
			serviceHandle = new(TPmThingModelServicesHandler)
			deleteResp    = new(protosService.TPmThingModelServicesResponse)
			deleteReq     = &protosService.TPmThingModelServicesRequest{
				ModelId:    iotutil.ToInt64(modelId),
				ProductKey: request.ProductKey,
			}
		)
		mapCond := make(map[int64]Cond)
		for _, v := range request.ThingModelServiceIds {
			mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
		}
		//删除物模型-服务关系
		err = serviceHandle.DeleteTPmThingModelServices(ctx, deleteReq, deleteResp)
		if err != nil {
			return err
		}

		var (
			thingModelServicesRequests    []*protosService.TPmThingModelServicesRequest
			thingModelServiceResponseList = new(protosService.TPmThingModelServicesResponseList)
			serviceQueryObj               = &protosService.TPmThingModelServicesFilter{
				Ids: GetModelItemIds(request.ThingModelServiceIds),
			}
		)
		//创建物模型-服务关系
		err = serviceHandle.ListTPmThingModelServices(ctx, &protosService.TPmThingModelServicesFilterPage{QueryObj: serviceQueryObj}, thingModelServiceResponseList)
		if err != nil {
			return err
		}
		thingModelServicesRequests = thingModelServiceResponseList.GetList()
		for _, value := range thingModelServicesRequests {
			if con, ok := mapCond[value.Id]; ok {
				value.ExecCond = con.ExecCond
				value.TriggerCond = con.TriggerCond
			}
			value.Id = iotutil.GetNextSeqInt64()
			value.ModelId = iotutil.ToInt64(modelId)
			value.ProductKey = request.ProductKey
			value.Custom = 1
			value.CreatedTime = now
			value.UpdatedTime = now
		}
		var (
			thingModelServiceResponse = new(protosService.TPmThingModelServicesResponse)
		)
		err = serviceHandle.CreateTPmThingModelServicesBatch(ctx, thingModelServicesRequests, thingModelServiceResponse)
		if err != nil {
			return err
		}
	}

	//物模型-事件
	if len(request.ThingModelEventIds) > 0 {
		var (
			eventHandle = new(TPmThingModelEventsHandler)
			deleteResp  = new(protosService.TPmThingModelEventsResponse)
			deleteReq   = &protosService.TPmThingModelEventsRequest{
				ModelId:    iotutil.ToInt64(modelId),
				ProductKey: request.ProductKey,
			}
		)
		mapCond := make(map[int64]Cond)
		for _, v := range request.ThingModelEventIds {
			mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
		}

		//删除物模型-事件关系
		err = eventHandle.DeleteTPmThingModelEvents(ctx, deleteReq, deleteResp)
		if err != nil {
			return err
		}

		var (
			thingModelEventsRequests    []*protosService.TPmThingModelEventsRequest
			thingModelEventResponseList = new(protosService.TPmThingModelEventsResponseList)
			eventQueryObj               = &protosService.TPmThingModelEventsFilter{
				Ids: GetModelItemIds(request.ThingModelEventIds),
			}
		)
		//创建物模型-事件关系
		err = eventHandle.ListTPmThingModelEvents(ctx, &protosService.TPmThingModelEventsFilterPage{QueryObj: eventQueryObj}, thingModelEventResponseList)
		if err != nil {
			return err
		}
		thingModelEventsRequests = thingModelEventResponseList.GetList()
		for _, value := range thingModelEventsRequests {
			if con, ok := mapCond[value.Id]; ok {
				value.ExecCond = con.ExecCond
				value.TriggerCond = con.TriggerCond
			}
			value.Id = iotutil.GetNextSeqInt64()
			value.ModelId = iotutil.ToInt64(modelId)
			value.ProductKey = request.ProductKey
			value.Custom = 1
			value.CreatedTime = now
			value.UpdatedTime = now
			value.CreateTs = now
		}
		var (
			thingModelEventResponse = new(protosService.TPmThingModelEventsResponse)
		)
		err = eventHandle.CreateTPmThingModelEventsBatch(ctx, thingModelEventsRequests, thingModelEventResponse)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
*/

// 关联物模型
func (TPmProductHandler) relationThingModellForCreate(tx *orm.Query, ctx context.Context, request *protosService.TPmProductRequest, modelId int64) error {
	mapCond := make(map[int64]Cond)
	for _, v := range request.ThingModelPropertyIds {
		mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
	}
	for _, v := range request.ThingModelServiceIds {
		mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
	}
	for _, v := range request.ThingModelEventIds {
		mapCond[v.ThingModelIds] = Cond{ExecCond: v.ExecCond, TriggerCond: v.TriggerCond}
	}
	err := CopyTypeToProductForModelServices(tx, request.ProductTypeId, modelId, request.ProductKey, mapCond)
	if err != nil {
		return err
	}
	err = CopyTypeToProductForModelProperties(tx, request.ProductTypeId, modelId, request.ProductKey, mapCond)
	if err != nil {
		return err
	}
	err = CopyTypeToProductForModelEvents(tx, request.ProductTypeId, modelId, request.ProductKey, mapCond)
	if err != nil {
		return err
	}
	return nil
}

// ResetProductThingModels 刷新产品物模型
func (s TPmProductHandler) ResetProductThingModels(ctx context.Context, request *protosService.TPmProductRequest, response *protosService.Response) error {
	response.Code = -1
	q := orm.Use(iotmodel.GetDB())
	product, err := q.TPmProduct.WithContext(context.Background()).Where(q.TPmProduct.Id.Eq(request.Id)).First()
	if err != nil {
		response.Message = err.Error()
		return nil
	}
	productTypeId := product.ProductTypeId
	productKey := product.ProductKey
	models, err := q.TPmThingModel.WithContext(context.Background()).Where(q.TPmThingModel.ProductKey.Eq(productKey)).Find()
	if err != nil {
		response.Message = err.Error()
		return nil
	}
	var modelId int64
	if len(models) == 0 {
		//创建thingsModel
		//创建物模型
		handleSvc := new(TPmThingModelHandler)
		thingModel := &protosService.TPmThingModelRequest{
			ProductKey:    productKey,
			ProductTypeId: productTypeId,
			Standard:      0,
			Version:       "V1.0.0",
			Description:   fmt.Sprintf("【%s】的物模型", request.GetName()),
			CreatedTime:   time.Now().Format("2006-01-02 15:04:05"),
		}
		thingsModelRes, err := handleSvc.SaveTPmThingModel(q, ctx, thingModel)
		if err != nil {
			return err
		}
		modelId = thingsModelRes.Id
	} else {
		modelId = models[0].Id
	}
	err = q.Transaction(func(tx *orm.Query) error {
		//先删除旧的物模型功能点
		ts := tx.TPmThingModelServices
		_, err := ts.WithContext(ctx).Unscoped().Where(ts.ProductKey.Eq(productKey)).Delete()
		if err != nil {
			return err
		}
		te := tx.TPmThingModelEvents
		_, err = te.WithContext(ctx).Unscoped().Where(te.ProductKey.Eq(productKey)).Delete()
		if err != nil {
			return err
		}
		tp := tx.TPmThingModelProperties
		_, err = tp.WithContext(ctx).Unscoped().Where(tp.ProductKey.Eq(productKey)).Delete()
		if err != nil {
			return err
		}
		err = CopyTypeToProductForModelServices(tx, productTypeId, modelId, productKey, nil)
		if err != nil {
			return err
		}
		err = CopyTypeToProductForModelProperties(tx, productTypeId, modelId, productKey, nil)
		if err != nil {
			return err
		}
		err = CopyTypeToProductForModelEvents(tx, productTypeId, modelId, productKey, nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Errorf("ResetProductThingModels error : %s", err.Error())
		return err
	}
	response.Code = SUCCESS
	response.Message = "success"
	return nil
}

// 关联物模型
func (TPmProductHandler) ResetProductThingModell(ctx context.Context, ProductTypeId, modelId int64, ProductKey string) error {
	q := orm.Use(iotmodel.GetDB())
	err := q.Transaction(func(tx *orm.Query) error {
		//先删除旧的物模型功能点
		ts := tx.TPmThingModelServices
		_, err := ts.WithContext(ctx).Unscoped().Where(ts.ModelId.Eq(modelId), ts.ProductKey.Eq(ProductKey)).Delete()
		if err != nil {
			return err
		}
		te := tx.TPmThingModelEvents
		_, err = te.WithContext(ctx).Unscoped().Where(te.ModelId.Eq(modelId), te.ProductKey.Eq(ProductKey)).Delete()
		if err != nil {
			return err
		}
		tp := tx.TPmThingModelProperties
		_, err = tp.WithContext(ctx).Unscoped().Where(tp.ModelId.Eq(modelId), tp.ProductKey.Eq(ProductKey)).Delete()
		if err != nil {
			return err
		}

		err = CopyTypeToProductForModelServices(tx, ProductTypeId, modelId, ProductKey, nil)
		if err != nil {
			return err
		}
		err = CopyTypeToProductForModelProperties(tx, ProductTypeId, modelId, ProductKey, nil)
		if err != nil {
			return err
		}
		err = CopyTypeToProductForModelEvents(tx, ProductTypeId, modelId, ProductKey, nil)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// 关联物模型
func (TPmProductHandler) relationThingModelForUpdate(tx *orm.Query, ctx context.Context, request *protosService.TPmProductRequest) error {
	//检查参数
	if request.ModelId == "" || request.ModelId == "0" {
		return errors.New("请选择模组芯片")
	}
	if request.ProductKey == "" {
		return errors.New("产品Key不能为空")
	}
	//先置原物模型无效
	modelid, err := strconv.Atoi(request.ModelId)
	if err != nil {
		return err
	}
	err = UpdateInvalidForModelServices(tx, int64(modelid), request.ProductKey)
	if err != nil {
		return err
	}
	err = UpdateInvalidForModelProperties(tx, int64(modelid), request.ProductKey)
	if err != nil {
		return err
	}
	err = UpdateInvalidForModelEvents(tx, int64(modelid), request.ProductKey)
	if err != nil {
		return err
	}

	//物模型-属性
	if len(request.ThingModelPropertyIds) > 0 {
		var propertyHandle = new(TPmThingModelPropertiesHandler)
		for _, v := range request.ThingModelPropertyIds {
			req := &protosService.TPmThingModelPropertiesRequest{
				Id:          v.ThingModelIds,
				TriggerCond: v.TriggerCond,
				ExecCond:    v.ExecCond,
				Valid:       1,
			}
			_ = propertyHandle.UpdateInfo(tx, ctx, req, &protosService.TPmThingModelPropertiesResponse{})
		}
	}
	//物模型-服务
	if len(request.ThingModelServiceIds) > 0 {
		var serviceHandle = new(TPmThingModelServicesHandler)
		for _, v := range request.ThingModelServiceIds {
			req := &protosService.TPmThingModelServicesRequest{
				Id:          v.ThingModelIds,
				TriggerCond: v.TriggerCond,
				ExecCond:    v.ExecCond,
				Valid:       1,
			}
			_ = serviceHandle.UpdateInfo(tx, ctx, req, &protosService.TPmThingModelServicesResponse{})
		}
	}
	//物模型-事件
	if len(request.ThingModelEventIds) > 0 {
		var eventHandle = new(TPmThingModelEventsHandler)
		for _, v := range request.ThingModelEventIds {
			req := &protosService.TPmThingModelEventsRequest{
				Id:          v.ThingModelIds,
				TriggerCond: v.TriggerCond,
				ExecCond:    v.ExecCond,
				Valid:       1,
			}
			_ = eventHandle.UpdateInfo(tx, ctx, req, &protosService.TPmThingModelEventsResponse{})
		}
	}
	return nil
}

// 更新固件关联关系
func (TPmProductHandler) relationProductFirmware(tx *orm.Query, request *protosService.TPmProductRequest) error {
	if len(request.FirmwareIds) == 0 {
		return nil
	}
	var (
		err                           error
		productFirmwareRelationDelete = &protosService.PmProductFirmwareRelation{
			ProductId: request.Id,
		}
		productFirmwareRelationSvc = new(service.PmProductFirmwareRelationSvc)
	)
	_, err = productFirmwareRelationSvc.DeletePmProductFirmwareRelation(tx, productFirmwareRelationDelete)
	if err != nil {
		return err
	}
	var productFirmwareRelations = make([]*protosService.PmProductFirmwareRelation, len(request.FirmwareIds))
	for i, id := range request.FirmwareIds {
		productFirmwareRelations[i] = new(protosService.PmProductFirmwareRelation)
		productFirmwareRelations[i].Id = iotutil.GetNextSeqInt64()
		productFirmwareRelations[i].ProductId = request.Id
		productFirmwareRelations[i].FirmwareId = id
	}
	_, err = productFirmwareRelationSvc.BatchCreatePmProductFirmwareRelation(tx, productFirmwareRelations)
	if err != nil {
		return err
	}
	return nil
}

// 更新模组关联关系
func (TPmProductHandler) relationProductModule(tx *orm.Query, request *protosService.TPmProductRequest) error {
	if len(request.ModuleIds) == 0 {
		return nil
	}
	var (
		err                         error
		productModuleRelationDelete = &protosService.PmProductModuleRelation{
			ProductId: request.Id,
		}
		productModuleRelations   = make([]*protosService.PmProductModuleRelation, len(request.ModuleIds))
		productModuleRelationSvc = new(service.PmProductModuleRelationSvc)
	)
	_, err = productModuleRelationSvc.DeletePmProductModuleRelation(tx, productModuleRelationDelete)
	if err != nil {
		return err
	}
	for i, id := range request.ModuleIds {
		productModuleRelations[i] = new(protosService.PmProductModuleRelation)
		productModuleRelations[i].Id = iotutil.GetNextSeqInt64()
		productModuleRelations[i].ProductId = request.Id
		productModuleRelations[i].ModuleId = id
	}
	_, err = productModuleRelationSvc.BatchCreatePmProductModuleRelation(tx, productModuleRelations)
	if err != nil {
		return err
	}
	return nil
}

// 更新控制面板关联关系
func (TPmProductHandler) relationProductPanel(tx *orm.Query, request *protosService.TPmProductRequest) error {
	if len(request.ControlPanelIds) == 0 {
		return nil
	}
	var (
		err                        error
		productPanelRelationDelete = &protosService.PmProductPanelRelation{
			ProductId: request.Id,
		}
		productPanelRelations   = make([]*protosService.PmProductPanelRelation, len(request.ControlPanelIds))
		productPanelRelationSvc = new(service.PmProductPanelRelationSvc)
	)
	_, err = productPanelRelationSvc.DeletePmProductPanelRelation(tx, productPanelRelationDelete)
	if err != nil {
		return err
	}
	for i, id := range request.ControlPanelIds {
		productPanelRelations[i] = new(protosService.PmProductPanelRelation)
		productPanelRelations[i].Id = iotutil.GetNextSeqInt64()
		productPanelRelations[i].ProductId = request.Id
		productPanelRelations[i].ControlPanelId = id
	}
	_, err = productPanelRelationSvc.BatchCreatePmProductPanelRelation(tx, productPanelRelations)
	if err != nil {
		return err
	}
	return nil
}

// 更新配网引导
func (TPmProductHandler) relationNetworkGuide(tx *orm.Query, request *protosService.TPmProductRequest) ([]int64, error) {
	if request.NetworkGuides == nil {
		return nil, nil
	}
	var (
		err                error
		networkGuideDelete = &protosService.PmNetworkGuide{
			ProductId: request.Id,
		}
		networkGuideSvc = new(service.PmNetworkGuideSvc)
	)
	_, err = networkGuideSvc.DeletePmNetworkGuide(tx, networkGuideDelete)
	if err != nil {
		return nil, err
	}

	var (
		ret *protosService.PmNetworkGuide
		ids = make([]int64, len(request.NetworkGuides))
	)
	for i, guide := range request.NetworkGuides {
		ret, err = networkGuideSvc.CreatePmNetworkGuide(tx, &protosService.PmNetworkGuide{
			Id:            iotutil.GetNextSeqInt64(),
			ProductId:     request.Id,
			ProductTypeId: request.ProductTypeId,
			Type:          guide.Type,
			CreatedAt:     timestamppb.New(time.Now()),
			UpdatedAt:     timestamppb.New(time.Now()),
		})
		ids[i] = ret.Id
	}
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// 更新配网引导步骤
func (TPmProductHandler) relationNetworkGuideStep(tx *orm.Query, request *protosService.TPmProductRequest, networkGuideIds []int64) error {
	var (
		err                    error
		networkGuideStepDelete = &protosService.PmNetworkGuideStep{
			NetworkGuideId: request.Id,
		}
		networkGuideSteps   = make([]*protosService.PmNetworkGuideStep, len(networkGuideIds)*10)
		networkGuideStepSvc = new(service.PmNetworkGuideStepSvc)
	)
	_, err = networkGuideStepSvc.DeletePmNetworkGuideStep(tx, networkGuideStepDelete)
	if err != nil {
		return err
	}
	var (
		index int
	)
	for i, networkGuideId := range networkGuideIds {
		for _, step := range request.NetworkGuides[i].Steps {
			networkGuideSteps[index] = new(protosService.PmNetworkGuideStep)
			networkGuideSteps[index].Id = iotutil.GetNextSeqInt64()
			networkGuideSteps[index].NetworkGuideId = networkGuideId
			networkGuideSteps[index].Instruction = step.Instruction
			networkGuideSteps[index].InstructionEn = step.InstructionEn
			networkGuideSteps[index].ImageUrl = step.ImageUrl
			networkGuideSteps[index].VideoUrl = step.VideoUrl
			networkGuideSteps[index].Sort = step.Sort
			networkGuideSteps[index].CreatedAt = timestamppb.New(time.Now())
			networkGuideSteps[index].UpdatedAt = timestamppb.New(time.Now())
			index++
		}
	}

	if networkGuideIds == nil {
		return nil
	}
	networkGuideSteps = networkGuideSteps[0:index]

	_, err = networkGuideStepSvc.BatchCreatePmNetworkGuideStep(tx, networkGuideSteps)
	if err != nil {
		return err
	}
	return nil
}

// 新增和修改的时候判断分类名称是否重复
func (s *TPmProductHandler) existsByName(name string, nameEn string, id int64) (bool, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	do = do.Where(do.Where(t.Name.Eq(name)).Or(do.Where(t.NameEn.Eq(nameEn))))

	//编辑的时候验证名称是否重复
	if id != 0 {
		do = do.Where(t.Id.Neq(id))
	}
	count, err := do.Count()
	if err != nil {
		return true, err
	}

	return count > 0, err
}

func (s TPmProductHandler) checkData(productId int64) error {
	//更新并发布，检查所有参数是否完整
	//if request.Status == 2 {
	productDetails := &protosService.TPmProductResponseObject{}
	err := s.GetByIdTPmProduct(context.Background(), &protosService.TPmProductFilterById{
		Id: productId,
	}, productDetails)
	if err != nil {
		return err
	}
	//检查产品数据是否完整
	//模组、固件配置是否完整
	if productDetails.Data.ModuleIds == nil || len(productDetails.Data.ModuleIds) == 0 {
		return errors.New("请选择产品芯片模组")
	}
	if productDetails.Data.FirmwareIds == nil || len(productDetails.Data.FirmwareIds) == 0 {
		return errors.New("请选择产品固件")
	}
	if productDetails.Data.NetworkGuides == nil || len(productDetails.Data.NetworkGuides) == 0 {
		return errors.New("请选择产品配网方式")
	}
	//面板是否完整
	if productDetails.Data.ControlPanelIds == nil || len(productDetails.Data.ControlPanelIds) == 0 {
		return errors.New("请选择产品面板")
	}
	//物模型是否完整
	//!s.hasThingsModels(productDetails.Data.ThingModelServiceIds) &&
	//!s.hasThingsModels(productDetails.Data.ThingModelEventIds)
	if !s.hasThingsModels(productDetails.Data.ThingModelPropertyIds) {
		return errors.New("请选择产品物模型")
	}
	return err
}

func (s TPmProductHandler) hasThingsModels(list []*protosService.ThingModelInfo) bool {
	if list == nil || len(list) == 0 {
		return false
	}
	newList := make([]*protosService.ThingModelInfo, 0)
	for _, info := range list {
		if info.Valid == 1 {
			newList = append(newList, info)
		}
	}
	return len(newList) > 0
}

func (TPmProductHandler) GetBaseProductInfo(baseProductId int64) (*model.TPmProduct, error) {
	t := orm.Use(iotmodel.GetDB()).TPmProduct
	do := t.WithContext(context.Background())
	do = do.Where(t.Id.Eq(baseProductId))
	info, err := do.First()
	if err != nil {
		return info, err
	}
	return info, nil
}
