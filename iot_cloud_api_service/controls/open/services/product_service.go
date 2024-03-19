package services

import (
	"bufio"
	"bytes"
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls/common"
	services2 "cloud_platform/iot_cloud_api_service/controls/global"
	entitys3 "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/controls/lang/services"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	entitys2 "cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/mitchellh/mapstructure"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OpmProductService struct {
	Ctx context.Context
}

var tempPath = iotconst.GetWorkTempDir() + string(filepath.Separator)

func (s OpmProductService) SetContext(ctx context.Context) OpmProductService {
	s.Ctx = ctx
	return s
}

// 产品详细
func (s OpmProductService) GetOpmProductDetail(id string) (*entitys.OpmProductEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	if len(req.Data) == 0 {
		return nil, errors.New("not found")
	}
	var data = req.Data[0]
	return entitys.OpmProduct_pb2e(data), err
}

// 产品详细
func (s OpmProductService) GetOpmProductAllDetail(id string) (*entitys.OpmProductAllEntitys, error) {
	if id == "" {
		return nil, errors.New("产品编号不能为空")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmProductService.FindByAllDetails(s.Ctx, &protosService.OpmProductPrimarykey{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	resObj := new(entitys.OpmProductAllEntitys)
	resObj.Pd2Entity(req)
	return resObj, err
}

// 产品详细
func (s OpmProductService) GetCompleteDevelopDetail(id string) (*entitys.CompleteDevelopDetail, error) {
	if id == "" {
		return nil, errors.New("产品编号不能为空")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmProductService.FindByAllDetails(s.Ctx, &protosService.OpmProductPrimarykey{Id: rid})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	resObj := new(entitys.CompleteDevelopDetail)
	resObj.Product = entitys.CompleteDevelopDetailProduct{
		ProductTypeName: req.Product.ProductTypeName,
		Name:            req.Product.Name,
		NameEn:          req.Product.NameEn,
		Identifier:      req.Product.Identifier,
		Model:           req.Product.Model,
		NetworkType:     req.Product.NetworkType,
	}

	//1-功能定义，2-硬件开发，3-设备面板 ，4 -基础配置
	resObj.List = []*entitys.CompleteDevelopDetailItems{}
	if req.Module != nil {
		resObj.List = append(resObj.List, &entitys.CompleteDevelopDetailItems{
			Code:  2,
			Key:   "使用模组",
			Valid: true,
			Title: req.Module.ModuleName,
		})
	}

	propertiesArr := make([]string, 0)
	for _, property := range req.ThingModes.Properties {
		propertiesArr = append(propertiesArr, property.Name)
	}
	resObj.List = append(resObj.List, &entitys.CompleteDevelopDetailItems{
		Code:  1,
		Key:   "功能点",
		Valid: len(propertiesArr) > 0,
		Title: strings.Join(propertiesArr, "，"),
	})
	//验证固件选择数据
	firmwareName := ""
	firmwareValid := false
	if req.Module != nil {
		firmwareName = req.Module.FirmwareName
		firmwareValid = true
	}
	resObj.List = append(resObj.List, &entitys.CompleteDevelopDetailItems{
		Code:  2,
		Key:   "使用固件",
		Valid: firmwareValid,
		Title: firmwareName,
	})

	panelName, url := "", ""
	if req.ControlPanel != nil {
		panelName, url = req.ControlPanel.Name, req.ControlPanel.PreviewUrl
	}
	resObj.List = append(resObj.List, &entitys.CompleteDevelopDetailItems{
		Code:  3,
		Key:   "使用面板",
		Valid: req.Product.ControlPanelId != 0,
		Title: panelName,
		Url:   url,
	})
	return resObj, err
}

// QueryOpmProductList 产品列表
func (s OpmProductService) QueryOpmProductList(filter entitys.OpmProductQuery) ([]*entitys.OpmProductEntitys, int64, error) {
	if err := filter.QueryCheck(); err != nil {
		return nil, 0, err
	}
	rep, err := rpc.ClientOpmProductService.Lists(s.Ctx, &protosService.OpmProductListRequest{
		Page:       int64(filter.Page),
		PageSize:   int64(filter.Limit),
		SearchKey:  filter.SearchKey,
		IsPlatform: filter.IsPlatform,
		Query:      entitys.OpmProductFilter_e2pb(filter.Query),
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmProductEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmProduct_pb2e(item))
	}
	return resultList, rep.Total, err
}

// AddOpmProduct 新增产品
func (s OpmProductService) AddOpmProduct(req entitys.OpmProductEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	saveObj := entitys.OpmProduct_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.ProductKey = iotutil.GetProductKeyRandomString()
	saveObj.Status = 2
	res, err := rpc.ClientOpmProductService.Create(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//go s.SetProductNameTranslate(saveObj.Name, saveObj.NameEn)
	// 同步创建产品帮助中心配置项
	langs := []entitys.Lang{
		{
			Lang:     "en",
			LangName: "英文",
		},
		{
			Lang:     "zh",
			LangName: "简体中文",
		},
	}
	langsBytes, err := json.Marshal(langs)
	if err != nil {
		return "", err
	}
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	reqConf := &protosService.ProductHelpConf{
		Id:            saveObj.Id,
		TenantId:      tenantId,
		ProductKey:    saveObj.ProductKey,
		ProductName:   saveObj.Name,
		ProductTypeId: saveObj.ProductTypeId,
		Langs:         string(langsBytes),
		RemainLang:    "zh",
		Status:        2,
		CreatedAt:     timestamppb.New(time.Now()),
		UpdatedAt:     timestamppb.New(time.Now()),
	}
	resp, err := rpc.ClientProductHelpConfService.Create(s.Ctx, reqConf)
	if err != nil {
		return "", err
	}
	if resp.Code != ioterrs.Success {
		return "", errors.New(resp.Message)
	}
	return iotutil.ToString(saveObj.Id), err
}

// SetProductNameTranslate 设置翻译
func (s OpmProductService) SetProductNameTranslate(name string, nameEn string) {
	defer iotutil.PanicHandler()
	if name == "" {
		return
	}
	langSvr := services.LangTranslateService{Ctx: s.Ctx}
	langReq := entitys3.LangTranslateEntitys{
		SourceTable: "",
		SourceRowId: "",
		TranslateList: []entitys3.BatchSaveTranslateItem{
			{
				Lang:       "zh",
				FieldName:  "name",
				FieldValue: name,
			},
		},
	}
	if nameEn == "" {
		langReq.TranslateList = append(langReq.TranslateList, entitys3.BatchSaveTranslateItem{
			Lang:       "en",
			FieldName:  "name",
			FieldValue: nameEn,
		})
	}
	langSvr.BatchInsert(langReq)
}

// UpdateOpmProduct 修改产品
func (s OpmProductService) UpdateOpmProduct(req entitys.OpmProductEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	res, err := rpc.ClientOpmProductService.Update(s.Ctx, entitys.OpmProduct_e2pb(&req))
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//go s.SetProductNameTranslate(req.Name, req.NameEn)
	reqConf := &protosService.ProductHelpConfUpdateFieldsRequest{
		Fields: []string{"product_name", "updated_at"},
		Data: &protosService.ProductHelpConf{
			Id:          iotutil.ToInt64(req.Id),
			ProductName: req.Name,
			UpdatedAt:   timestamppb.New(time.Now()),
		},
	}
	resp, err := rpc.ClientProductHelpConfService.UpdateFields(s.Ctx, reqConf)
	if err != nil {
		return "", err
	}
	if resp.Code != ioterrs.Success {
		return "", errors.New(resp.Message)
	}
	return iotutil.ToString(req.Id), err
}

func (s OpmProductService) UpdateOpmProductPanelInfo(req entitys.OpmProductEntitys) (string, error) {
	resPro, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{Id: req.Id})
	if err != nil {
		return "", err
	}
	if resPro.Code != 200 {
		return "", errors.New(resPro.Message)
	}
	if len(resPro.Data) == 0 {
		return "", errors.New("not found")
	}
	var data = resPro.Data[0]
	productKey := data.ProductKey

	//如果传入了图片，则视为需要显示图片
	var isShowImg int32 = 0
	if req.PanelProImg != "" {
		isShowImg = 1
	}

	res, err := rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{
		Id:           req.Id,
		PanelProImg:  req.PanelProImg,
		IsShowImg:    isShowImg, //iotutil.IfInt32(req.IsShowImg, 1, 2),
		StyleLinkage: req.StyleLinkage,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//设备缓存
	if cerr := iotredis.GetClient().HMSet(context.Background(), fmt.Sprintf(iotconst.PANEL_RULE_SETTINGS_DATA, productKey, req.PanelId), map[string]interface{}{
		"panelProImg":  req.PanelProImg,
		"isShowImg":    iotutil.IfInt32(req.IsShowImg, 1, 2),
		"styleLinkage": req.StyleLinkage,
	}).Err(); cerr != nil {
		iotlogger.LogHelper.Errorf("面板和产品绑定设置缓存失败，Err：%v", cerr.Error())
	}
	return iotutil.ToString(req.Id), err
}

// RefreshUpdateTime 刷新产品更新时间
func (s OpmProductService) RefreshUpdateTime(req entitys.OpmProductEntitys) {
	rpc.ClientOpmProductService.UpdateFields(s.Ctx, &protosService.OpmProductUpdateFieldsRequest{
		Fields: []string{"updated_at"},
		Data: &protosService.OpmProduct{
			Id:        iotutil.ToInt64(req.Id),
			UpdatedAt: timestamppb.New(time.Now()),
		},
	})
}

// DeleteOpmProduct 删除产品
func (s OpmProductService) DeleteOpmProduct(req entitys.OpmProductFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	// 先查询产品
	respProduct, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if respProduct.Code != 200 {
		return errors.New(respProduct.Message)
	}
	// 删除产品帮助中心文档
	_, err = rpc.ClientProductHelpDocService.Delete(s.Ctx, &protosService.ProductHelpDoc{
		TenantId:   tenantId,
		ProductKey: respProduct.Data[0].ProductKey,
	})
	if err != nil {
		return err
	}
	// 删除产品帮助中心配置
	_, err = rpc.ClientProductHelpConfService.DeleteById(s.Ctx, &protosService.ProductHelpConf{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	// 删除产品
	rep, err := rpc.ClientOpmProductService.Delete(s.Ctx, &protosService.OpmProduct{
		Id: iotutil.ToInt64(req.Id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// SetStatusOpmProduct 禁用/启用产品
func (s OpmProductService) SetStatusOpmProduct(req entitys.OpmProductFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.Status == 0 {
		return errors.New("status not found")
	}

	//开发完成重新生成数据库日志表接口
	if req.Status == 1 {
		//生成数据库表
		err := s.createLogTable(req.Id)
		if err != nil {
			return err
		}
	}

	rep, err := rpc.ClientOpmProductService.UpdateFields(s.Ctx, &protosService.OpmProductUpdateFieldsRequest{
		Fields: []string{"status"},
		Data: &protosService.OpmProduct{
			Id:     iotutil.ToInt64(req.Id),
			Status: req.Status,
		},
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

func (s OpmProductService) createLogTable(productId int64) error {
	//产品详细信息查询
	productRes, err := rpc.ClientOpmProductService.FindByAllDetails(s.Ctx, &protosService.OpmProductPrimarykey{Id: productId})
	if err != nil {
		return err
	}
	if productRes.Code != 200 {
		return errors.New(productRes.Message)
	}

	thingModels := make(map[string]string)
	thingModelInfo := make(map[string]string)
	thingModelInfo["imageUrl"] = productRes.Product.ImageUrl
	thingModelInfo["productKey"] = productRes.Product.ProductKey
	//thingModelInfo["networkType"] = productRes.Product.NetworkType
	thingModelInfo["name"] = productRes.Product.Name
	thingModelInfo["nameEn"] = productRes.Product.NameEn
	thingModelInfo["wifiFlag"] = productRes.Product.WifiFlag
	thingModelInfo["firmwareId"] = iotutil.ToString(productRes.Module.FirmwareId)
	for _, property := range productRes.ThingModes.Properties {
		thingModels[property.Identifier] = property.DataType
		//缓存物模型的内容
		thingModelInfo[iotconst.FIELD_PREFIX_TLS+iotutil.ToString(property.Dpid)] = iotutil.ToString(map[string]interface{}{
			"identifier":    property.Identifier,
			"dataType":      property.DataType,
			"name":          property.Name,
			"rwFlag":        property.RwFlag,
			"dataSpecs":     property.DataSpecs,
			"dataSpecsList": property.DataSpecsList,
			"custom":        property.Custom,
			"dpid":          property.Dpid,
		})
	}
	//缓存物模型数据
	rdCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productRes.Product.ProductKey, thingModelInfo)
	if rdCmd.Err() != nil {
		return rdCmd.Err()
	}

	res, err := rpc.ClientIotDeviceLogServer.CreateProductLogTable(s.Ctx, &protosService.CreateProductLogTableResponse{
		ProductId:   productRes.Product.ProductId,
		ProductKey:  productRes.Product.ProductKey,
		ThingModels: thingModels,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// QueryProductNetworkGuide 产品配网引导数据查询
func (s OpmProductService) QueryProductNetworkGuide(id string) ([]*entitys.OpmNetworkGuideEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}
	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmNetworkGuideService.FindByProductId(s.Ctx, &protosService.OpmNetworkGuideFilter{
		Id:        rid,
		ProductId: rid,
		//Type: ?
	})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	var list []*entitys.OpmNetworkGuideEntitys
	for _, data := range req.Data {
		steps := []*entitys.OpmNetworkGuideStepEntitys{}
		for _, step := range data.Steps {
			steps = append(steps, entitys.OpmNetworkGuideStep_pb2e(step))
		}
		list = append(list, &entitys.OpmNetworkGuideEntitys{
			Id:        data.Id,
			ProductId: data.ProductId,
			Type:      data.Type,
			Steps:     steps,
		})
	}
	return list, err
}

// QueryProductNetworkGuide 产品配网引导数据查询(默认配网方式）
func (s OpmProductService) QueryProductDefaultNetworkGuide(id string, networkGuideType int32) (*entitys.OpmNetworkGuideEntitys, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}

	rid := iotutil.ToInt64(id)
	req, err := rpc.ClientOpmNetworkGuideService.FindDefaultByProductId(s.Ctx, &protosService.OpmNetworkGuideFilter{
		ProductId: rid,
		Type:      networkGuideType,
	})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	if len(req.Data) == 0 {
		return nil, errors.New("未获取到默认配网方式")
	}

	var list []*entitys.OpmNetworkGuideEntitys
	for _, data := range req.Data {
		steps := []*entitys.OpmNetworkGuideStepEntitys{}
		for _, step := range data.Steps {
			steps = append(steps, entitys.OpmNetworkGuideStep_pb2e(step))
		}
		list = append(list, &entitys.OpmNetworkGuideEntitys{
			Id:        data.Id,
			ProductId: data.ProductId,
			Type:      data.Type,
			Steps:     steps,
		})
	}

	//考虑之后可能有多个默认返回的情况
	return list[0], nil
}

// SaveProductNetworkGuide 新增产品
func (s OpmProductService) SaveProductNetworkGuide(req entitys.OpmNetworkGuideEntitys) (string, error) {
	steps := []*protosService.OpmNetworkGuideStep{}
	for _, step := range req.Steps {
		steps = append(steps, entitys.OpmNetworkGuideStep_e2pb(step))
	}
	saveObj := &protosService.OpmNetworkGuide{
		Id:        req.Id,
		ProductId: req.ProductId,
		Type:      req.Type,
		Steps:     steps,
	}

	saveObj.Id = iotutil.GetNextSeqInt64()
	res, err := rpc.ClientOpmNetworkGuideService.CreateAndUpdate(s.Ctx, saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//刷新产品更新时间
	if req.ProductId != 0 {
		rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.ProductId})
	}
	return iotutil.ToString(saveObj.Id), err
}

// SetProductNetworkGuideType 设置配网引导类型
func (s OpmProductService) SetProductNetworkGuideType(req entitys.ChangeNetworkGuideRequest) error {
	saveObj := &protosService.SetNetworkGuideTypeRequest{
		ProductId: req.ProductId,
		Type:      req.Type,
	}
	res, err := rpc.ClientOpmNetworkGuideService.SetNetworkGuideTypes(s.Ctx, saveObj)
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

// QueryProductThingModel 开放平台-获取产品基础物模型数据
func (s OpmProductService) QueryProductThingModel(productId string, isCustom int32) (*entitys.OpmThingModelList, error) {
	if productId == "" {
		return nil, errors.New("产品编号不存在")
	}
	result := new(entitys.OpmThingModelList)
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	lang, _ := metadata.Get(s.Ctx, "lang")
	//if tenantId == "" {
	//	return nil, errors.New("租户不存在")
	//}
	if lang == "" {
		lang = "zh"
	}
	//if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, productId), result); err == nil {
	//	return result, nil
	//}
	productIdInt := iotutil.ToInt64(productId)
	//兼容云管平台查询
	if tenantId == "" {
		proRes, err := rpc.ClientOpmProductService.FindById(context.Background(), &protosService.OpmProductFilter{Id: productIdInt})
		if err != nil {
			return nil, err
		}
		tenantId = proRes.Data[0].TenantId
	}
	res, err := rpc.ClientOpmProductVoiceService.GetVoiceProductFunc(s.Ctx, &protosService.OpmVoiceProductListReq{
		ProductId: productIdInt,
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}

	req, err := rpc.ClientOpmThingModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productIdInt,
		Custom:    isCustom,
	})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}
	result.WebMQTT = config.Global.WebMQTT.Addr
	result.Model = entitys.OpmThingModel_pb2e(req.Data.Model)

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	for _, property := range req.Data.Properties {
		properties := entitys.OpmThingModelProperties_pb2e(property)
		properties.Name = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, property.ProductKey, property.Identifier)], property.Name)
		properties.DataSpecsList = ConvertJsonByLang(lang, property.ProductKey, property.Identifier, property.DataSpecsList, langMap)
		if _, ok := res.FuncMap[properties.Identifier]; ok {
			properties.NotAllowEdit = true
		}
		result.List = append(result.List, properties)
	}
	//result := new(entitys.OpmThingModelAllList)
	//result.Model = entitys.OpmThingModel_pb2e(req.Data.Model)
	//for _, property := range req.Data.Properties {
	//	result.Properties = append(result.Properties, entitys.OpmThingModelProperties_pb2e(property))
	//}
	//for _, event := range req.Data.Events {
	//	result.Events = append(result.Events, entitys.OpmThingModelEvents_pb2e(event))
	//}
	//for _, services := range req.Data.Services {
	//	result.Services = append(result.Services, entitys.OpmThingModelServices_pb2e(services))
	//}
	//err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, productId), result, 600*time.Second)
	//if err != nil {
	//	return result, err
	//}
	return result, err
}

// QueryProductThingModel 开放平台-获取产品基础物模型数据
func (s OpmProductService) QueryStandardThingModel(productId string) (*entitys.OpmThingModelList, error) {
	if productId == "" {
		return nil, errors.New("产品编号不存在")
	}
	productIdStr := iotutil.ToInt64(productId)
	req, err := rpc.ClientOpmThingModelService.StandardThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productIdStr,
	})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	result := new(entitys.OpmThingModelList)
	//result.Model = entitys.OpmThingModel_pb2e(req.Data.Model)
	for _, property := range req.Data.Properties {
		result.List = append(result.List, entitys.OpmThingModelProperties_pb2e(property))
	}
	return result, err
}

// 从标准物模型添加
func (s OpmProductService) AddOpmThingModel(req entitys.OpmThingModelPropertiesEntitys) (string, error) {
	if err := req.AddCheck(); err != nil {
		return "", err
	}
	req.Required = 2 //默认不必填
	saveObj := protosService.OpmThingModel{
		ProductId:  iotutil.ToInt64(req.ProductId),
		Version:    "v1.0.0",
		Standard:   0,
		Properties: entitys.OpmThingModelProperties_e2pb(&req),
	}
	res, err := rpc.ClientOpmThingModelService.Create(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新产品修改时间，用于我的产品列表最近更新时间字段
	_, err = rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: iotutil.ToInt64(req.ProductId), TslUpdatedAt: timestamppb.Now()})
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

func (s OpmProductService) AddOpmThingModelByStandard(req entitys.AddStandardThingModelRequest) (string, error) {
	if req.ProductId == "" {
		return "", errors.New("产品编号不存在")
	}

	funcs := []*protosService.OpmStandardFuncs{}
	for _, f := range req.FuncList {
		funcs = append(funcs, &protosService.OpmStandardFuncs{
			ModelId:    f.ModelId,
			FuncId:     f.Id,
			FuncType:   "properties",
			Identifier: f.Identifier,
		})
	}

	saveObj := protosService.OpmThingModel{
		ProductId:     iotutil.ToInt64(req.ProductId),
		Version:       "v1.0.0",
		Standard:      1,
		StandardFuncs: funcs,
	}
	res, err := rpc.ClientOpmThingModelService.Create(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新产品修改时间，用于我的产品列表最近更新时间字段
	_, err = rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: iotutil.ToInt64(req.ProductId), TslUpdatedAt: timestamppb.Now()})
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// 修改物模型
func (s OpmProductService) UpdateOpmThingModel(req entitys.OpmThingModelPropertiesEntitys) (string, error) {
	if err := req.UpdateCheck(); err != nil {
		return "", err
	}
	updateObj := protosService.OpmThingModel{
		ProductId:  iotutil.ToInt64(req.ProductId),
		Version:    "v1.0.0",
		Standard:   req.Standard,
		Properties: entitys.OpmThingModelProperties_e2pb(&req),
	}
	res, err := rpc.ClientOpmThingModelService.Update(s.Ctx, &updateObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: iotutil.ToInt64(req.ProductId), TslUpdatedAt: timestamppb.Now()})
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// 检查语音配置
func (s *OpmProductService) getTslAndVoiceRelation(productKey, identifier string) (map[string]bool, error) {
	rep, err := rpc.ClientOpmVoiceProductMapService.Lists(s.Ctx, &protosService.OpmVoiceProductMapListRequest{
		Query: &protosService.OpmVoiceProductMap{
			ProductKey: productKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	res := make(map[string]bool)
	for _, d := range rep.Data {
		if d.AttrCode == identifier {
			res[d.VoiceNo] = true
		}
	}
	return res, nil
}

func (s OpmProductService) getTslByDpid(productId string, id int64) (*protosService.OpmThingModelProperties, error) {
	tslList, err := rpc.ClientOpmThingModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: iotutil.ToInt64(productId),
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if tslList.Data == nil {
		return nil, nil
	}
	var resTsl *protosService.OpmThingModelProperties
	for _, t := range tslList.Data.Properties {
		if t.Id == id {
			resTsl = t
			break
		}
	}
	return resTsl, nil
}

// 删除物模型
func (s OpmProductService) DeleteOpmThingModel(req entitys.OpmThingModelPropertiesEntitys) (*entitys.DeleteThingsModelResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id not found")
	}
	//查询查询信息
	proInfo, err := s.GetOpmProductDetail(req.ProductId)
	if err != nil {
		return nil, err
	}
	//已开放完成的产品无法删除物模型
	if proInfo.Status == 1 {
		return nil, errors.New("已开发完成的产品无法删除物模型")
	}

	//查询功能详情
	tslInfo, err := s.getTslByDpid(req.ProductId, req.Id)
	if err != nil {
		return nil, err
	}
	tslVoiceRel := &entitys.DeleteThingsModelResponse{
		ShowDeleteMsg: false,
		Voice:         make(map[string]bool),
	}
	//获取功能与语音的关联关系
	if tslInfo != nil && tslInfo.Identifier != "" {
		//删除之前检查使用情况
		tslVoiceRel.Voice, err = s.getTslAndVoiceRelation(proInfo.ProductKey, tslInfo.Identifier)
		if err != nil {
			return nil, err
		}
		for _, b := range tslVoiceRel.Voice {
			if b {
				tslVoiceRel.ShowDeleteMsg = true
				break
			}
		}
	}
	//验证不通过，不需要执行删除
	if tslVoiceRel.ShowDeleteMsg == false {
		rep, err := rpc.ClientOpmThingModelService.DeleteThingModel(s.Ctx, &protosService.OpmThingModelDeleteRequest{
			FuncId:   req.Id,
			FuncType: req.FuncType,
		})
		if err != nil {
			return nil, err
		}
		if rep.Code != 200 {
			return nil, errors.New(rep.Message)
		}
		if req.ProductId != "" {
			//刷新产品更新时间
			rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: iotutil.ToInt64(req.ProductId), TslUpdatedAt: timestamppb.Now()})
			cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
		}
	}
	return tslVoiceRel, nil
}

// 开放平台-根据产品类型ID获取模组SDK列表
func (s OpmProductService) GetOpenModuleListByProductId(productId int64) ([]*entitys2.PmModuleEntitys, error) {
	////查询产品类型ID关联模组列表
	//productModuleRelationObj, err := rpc.ClientProductModuleRelationService.Lists(context.Background(), &protosService.PmProductModuleRelationListRequest{
	//	Page:     1,
	//	PageSize: 100,
	//	Query:    &protosService.PmProductModuleRelation{ProductId: productId},
	//})
	//if err != nil {
	//	logger.Errorf("GetOpenModuleListByProductId error : %s", err.Error())
	//	return nil, err
	//}
	//if len(productModuleRelationObj.Data) <= 0 {
	//	return nil, errors.New("不存在关联模组")
	//}
	//
	//var ids = make([]int64, len(productModuleRelationObj.Data))
	//for index, relation := range productModuleRelationObj.Data {
	//	ids[index] = relation.ModuleId
	//}
	if productId == 0 {
		return nil, errors.New("产品编号异常")
	}
	ret, err := rpc.ClientOpmProductService.ModuleLists(context.Background(), &protosService.ModuleIdsRequest{
		//ModuleIds: ids,
		ProductId: productId,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	result := make([]*entitys2.PmModuleEntitys, 0)
	for _, v := range ret.Data {
		result = append(result, entitys2.PmModuleVo_pb2e(v))
	}
	return result, err
}

// 开放平台-根据产品类型ID获取控制面板列表
func (s OpmProductService) GetOpenControlPanelsListByProductId(baseProductId int64, productId int64) (resp []*entitys2.PmControlPanelsEntitys, err error) {
	//查询产品类型ID关联模组列表
	//productPanelRelationObj, err := rpc.ClientProductPanelRelationService.Lists(context.Background(), &protosService.PmProductPanelRelationListRequest{
	//	Page:     1,
	//	PageSize: 100,
	//	Query:    &protosService.PmProductPanelRelation{ProductId: baseProductId},
	//})
	//if err != nil {
	//	logger.Errorf("GetOpenControlPanelsListByProductId error : %s", err.Error())
	//	return nil, err
	//}
	//if len(productPanelRelationObj.Data) <= 0 {
	//	return nil, errors.New("不存在关联模组")
	//}
	//
	//var id = make([]int64, len(productPanelRelationObj.Data))
	//for index, relation := range productPanelRelationObj.Data {
	//	id[index] = relation.ControlPanelId
	//}
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	if baseProductId == 0 {
		return nil, errors.New("产品编号异常")
	}
	ret, err := rpc.ClientOpmProductService.ControlPanelsLists(s.Ctx, &protosService.ControlPanelIdsRequest{
		ProductId: baseProductId,
	})
	if err != nil {
		return nil, err
	}
	if ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	result := make([]*entitys2.PmControlPanelsEntitys, 0)
	for _, v := range ret.Data {
		item := entitys2.PmControlPanelsVo_pb2e(v)
		item.AppPanelType = 1
		result = append(result, item)
	}

	//productId
	rep, err := rpc.ClientOpmPanelService.Lists(s.Ctx, &protosService.OpmPanelListRequest{
		OrderDesc: "desc",
		OrderKey:  "updated_at",
		Query: &protosService.OpmPanel{
			ProductId: productId,
			PanelType: 1, //自定义面板
			TenantId:  tenantId,
			Status:    3, //已发布固件
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(ret.Message)
	}
	for _, item := range rep.Data {
		if item.PanelType == 2 {
			continue
		}
		panelInfo := entitys.OpmPanels_pb2panels(item)
		panelInfo.AppPanelType = 2
		result = append(result, panelInfo)
	}
	return result, err
}

// 开放平台-保存产品与模组关系
func (s OpmProductService) SaveOpenProductAndModuleRelation(req entitys.OpmProductModuleRelationEntitys) (err error) {
	//删除开放平台该产品下的模组关联
	//_, err = rpc.ClientOpmProductModuleRelationService.Delete(context.Background(), &protosService.OpmProductModuleRelation{
	//	ProductId: productModuleRelations[0].ProductId,
	//})
	relations := []*protosService.OpmProductModuleRelation{}
	relations = append(relations, &protosService.OpmProductModuleRelation{
		ProductId:  req.ProductId,
		ModuleId:   req.ModuleId,
		FirmwareId: req.FirmwareId,
		IsCustom:   req.IsCustom,
	})
	//创建开放平台该产品下的模组关联
	ret, err := rpc.ClientOpmProductModuleRelationService.BatchCreate(s.Ctx, &protosService.OpmProductModuleRelationList{
		ProductModuleRelations: relations,
	})
	if err != nil {
		logger.Errorf("SaveOpenProductAndModuleRelation error : %s", err.Error())
		return err
	}
	if ret.Code != 200 {
		return errors.New(ret.Message)
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.ProductId})
	//rpc.ClientOpmProductFirmwareRelationService.BatchCreate(context.Background(), &protosService.OpmProductModuleRelationList{
	//	ProductModuleRelations
	//})
	return err
}

// RemoveProductFirmwareRelation 移除自定义固件
func (s OpmProductService) RemoveProductFirmwareRelation(req entitys.OpmProductModuleRelationEntitys) (err error) {
	//创建开放平台该产品下的模组关联
	_, err = rpc.ClientOpmProductModuleRelationService.Delete(s.Ctx, &protosService.OpmProductModuleRelation{
		Id: req.Id,
		//FirmwareId: req.FirmwareId,
		//ProductId:  req.ProductId,
		//IsCustom:   1,
	})
	if err != nil {
		logger.Errorf("SaveOpenProductAndModuleRelation error : %s", err.Error())
		return err
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.ProductId})
	return err
}

// 开放平台-保存产品与自定义固件关系
func (s OpmProductService) SaveOpenProductAndCustomFirmwareRelation(req entitys.OpmProductFirmwareRelationEntitys) (err error) {
	_, err = rpc.ClientOpmProductFirmwareRelationService.Create(s.Ctx, &protosService.OpmProductFirmwareRelation{
		ProductId:  req.ProductId,
		FirmwareId: req.FirmwareId,
	})
	if err != nil {
		logger.Errorf("SaveOpenProductAndCustomFirmwareRelation error : %s", err.Error())
		return err
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.ProductId})
	return err
}

// 开放平台-保存产品与控制面板关系
func (s OpmProductService) SaveOpenProductAndControlPanelRelation(req entitys.OpmProductPanelRelationEntitys) (panel *entitys.ControlPanel, err error) {
	////删除开放平台该产品下的控制面板关联
	//_, err = rpc.ClientOpmProductModuleRelationService.Delete(context.Background(), &protosService.OpmProductModuleRelation{
	//	ProductId: productPanelRelations[0].ProductId,
	//})
	//查询产品信息
	proInfoRes, err := rpc.ClientOpmProductService.FindById(s.Ctx, &protosService.OpmProductFilter{Id: req.ProductId})
	if err != nil {
		logger.Errorf("SaveOpenProductAndControlPanelRelation ClientOpmProductService.FindById error : %s", err.Error())
		return nil, err
	}
	if proInfoRes.Code != 200 {
		logger.Errorf("SaveOpenProductAndControlPanelRelation ClientOpmProductService.FindById error : %s", proInfoRes.Message)
		return nil, errors.New(proInfoRes.Message)
	}
	relations := make([]*protosService.OpmProductPanelRelation, 0)
	relations = append(relations, &protosService.OpmProductPanelRelation{
		ProductId:      req.ProductId,
		ControlPanelId: req.ControlPanelId,
		AppPanelType:   iotutil.GetInt32AndDef(req.AppPanelType, 1),
	})
	//创建开放平台该产品下的控制面板关联
	resq, err := rpc.ClientOpmProductPanelRelationService.BatchCreate(s.Ctx, &protosService.OpmProductPanelRelationList{
		ProductPanelRelations: relations,
	})
	if err != nil {
		logger.Errorf("SaveOpenProductAndControlPanelRelation error : %s", err.Error())
		return nil, err
	}
	if resq.Code != 200 {
		return nil, errors.New(resq.Message)
	}
	if len(resq.Data) == 0 {
		return nil, errors.New("未获取到面板信息")
	}
	//刷新产品更新时间 (BatchCreate已经有更新修改时间）
	//rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.ProductId})
	panel = &entitys.ControlPanel{
		Id:           resq.Data[0].ControlPanel.Id,
		FileURL:      resq.Data[0].ControlPanel.Url,
		PreviewURL:   resq.Data[0].ControlPanel.PreviewUrl,
		Name:         resq.Data[0].ControlPanel.Name,
		PanelProImg:  proInfoRes.Data[0].PanelProImg,
		IsShowImg:    iotutil.IntToBoolean(proInfoRes.Data[0].IsShowImg),
		StyleLinkage: proInfoRes.Data[0].StyleLinkage,
	}
	go ClearProductPanelLang(req.ProductId, "")
	return panel, nil
}

//==================================

// CreateOpenThingModel 开放平台-创建物模型（属性/服务/事件）  （备份）
func (s OpmProductService) CreateOpenThingModel(thingModel entitys.TOpmThingModel) (err error) {
	switch thingModel.FuncType {
	case "属性":
		//_, err = rpc.ClientOpmThingModelPropertiesService.Create(context.Background(), &protosService.OpmThingModelProperties{
		//	Id:            iotutil.GetNextSeqInt64(),
		//	ModelId:       thingModel.ModelId,
		//	ProductKey:    thingModel.ProductKey,
		//	CreateTs:      time.Now().Format("2006-01-02 15:04:05"),
		//	Identifier:    thingModel.Identifier,
		//	DataType:      thingModel.DataType,
		//	Name:          thingModel.FuncName,
		//	RwFlag:        thingModel.RwFlag,
		//	DataSpecs:     thingModel.Specs,
		//	DataSpecsList: thingModel.SpecsList,
		//	Required:      thingModel.Required,
		//	Custom:        thingModel.Custom,
		//})
	case "服务":
		//_, err = rpc.ClientOpmThingModelServicesService.Create(context.Background(), &protosService.OpmThingModelServices{
		//	Id:           iotutil.GetNextSeqInt64(),
		//	ModelId:      thingModel.ModelId,
		//	ProductKey:   thingModel.ProductKey,
		//	CreateTs:     time.Now().Format("2006-01-02 15:04:05"),
		//	Identifier:   thingModel.Identifier,
		//	ServiceName:  thingModel.FuncName,
		//	InputParams:  thingModel.InputParams,
		//	OutputParams: thingModel.OutputParams,
		//	Required:     thingModel.Required,
		//	CallType:     thingModel.CallType,
		//	Custom:       thingModel.Custom,
		//})
	case "事件":
		//_, err = rpc.ClientOpmThingModelEventsService.Create(context.Background(), &protosService.OpmThingModelEvents{
		//	Id:         iotutil.GetNextSeqInt64(),
		//	ModelId:    thingModel.ModelId,
		//	ProductKey: thingModel.ProductKey,
		//	CreateTs:   time.Now().Format("2006-01-02 15:04:05"),
		//	Identifier: thingModel.Identifier,
		//	EventName:  thingModel.FuncName,
		//	EventType:  thingModel.EventType,
		//	Outputdata: thingModel.Outputdata,
		//	Required:   thingModel.Required,
		//	Custom:     thingModel.Custom,
		//})
	}

	if err != nil {
		return err
	}
	return err
}

// DelOpenThingModel 开放平台-删除物模型（属性/服务/事件）  （备份）
func (s OpmProductService) DelOpenThingModel(thingModel entitys.TOpmThingModel) (err error) {
	//switch thingModel.FuncType {
	//case "属性":
	//	_, err = rpc.ClientOpmThingModelPropertiesService.DeleteById(context.Background(), &protosService.OpmThingModelProperties{Id: iotutil.ToInt64(thingModel.Id)})
	//case "服务":
	//	_, err = rpc.ClientOpmThingModelServicesService.DeleteById(context.Background(), &protosService.OpmThingModelServices{Id: iotutil.ToInt64(thingModel.Id)})
	//case "事件":
	//	_, err = rpc.ClientOpmThingModelEventsService.DeleteById(context.Background(), &protosService.OpmThingModelEvents{Id: iotutil.ToInt64(thingModel.Id)})
	//}

	if err != nil {
		return err
	}
	return err
}

// 开放平台-创建产品基础信息 （备份）
func (s OpmProductService) CreateOpenProduct(req *entitys.OpmProductEntitys) (ret int32, err error) {
	if err = req.AddCheck(); err != nil {
		return
	}
	var (
		data = protosService.OpmProduct{}
	)
	mapstructure.WeakDecode(req, &data)
	//参数填充
	data.ProductTypeId = iotutil.ToInt64(req.ProductTypeId)
	//data.ProductId = req.ProductId
	data.ProductKey = iotutil.GetRandomString(16)
	data.Identifier = iotutil.Uuid()
	data.CreatedAt = timestamppb.Now()
	data.UpdatedAt = timestamppb.Now()

	res, err := rpc.ClientOpmProductService.Create(context.Background(), &data)
	if err != nil {
		return 0, err
	}
	if res.Code != 200 {
		return 0, errors.New(res.Message)
	}
	return res.Code, err
}

// 开放平台-根据产品唯一标识获取标准物模型 （备份）
func (s OpmProductService) GetOpenStandardThingModelDetail(productKey string) (resp []*entitys.TOpmThingModelVo, err error) {
	//查询物模型基础信息
	thingModelObj, err := rpc.ClientThingModelService.GetTPmThingModel(context.Background(), &protosService.TPmThingModelFilter{
		ProductKey: productKey,
		Version:    "V1.0.0",
		Standard:   1,
	})
	var (
		thingModel *protosService.TPmThingModelRequest
	)
	thingModel = thingModelObj.GetData()
	if thingModelObj.GetData() == nil {
		err = fmt.Errorf("所属品类不存在标准物模型，productKey：%s，Version：%s", productKey, "V1.0.0")
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
	var thingModelVoResp = make([]*entitys.TOpmThingModelVo, cap)
	if err = s.mergeOpenPropertyAttributeDesc(propertyRes, serviceRes, eventRes, thingModelVoResp); err != nil {
		return nil, err
	}

	return thingModelVoResp, err
}

// 合并物模型属性/方法/事件
func (s OpmProductService) mergePropertyAttributeDesc(propertyRes *protosService.TPmThingModelPropertiesResponseList, serviceRes *protosService.TPmThingModelServicesResponseList, eventRes *protosService.TPmThingModelEventsResponseList, thingModelVoResp []*entitys.TOpmThingModelVo) error {
	var (
		attr  string
		index int
		err   error
	)
	//合并物模型-属性/服务/事件
	if propertyRes != nil && len(propertyRes.List) > 0 {
		for i, property := range propertyRes.List {
			entity := &entitys.TOpmThingModelVo{
				Id:         iotutil.ToString(property.Id),
				FuncType:   "属性",
				FuncName:   property.Name,
				Required:   property.Required,
				Identifier: property.Identifier,
				RwFlag:     property.RwFlag,
				DataType:   property.DataType,
			}
			attr, err = s.transformPropertyAttributeDesc(property.GetDataType(), property.GetDataSpecs(), property.GetDataSpecsList())
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
			entity := &entitys.TOpmThingModelVo{
				Id:         iotutil.ToString(service.Id),
				FuncType:   "服务",
				FuncName:   service.ServiceName,
				Required:   service.Required,
				Identifier: service.Identifier,
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
			entity := &entitys.TOpmThingModelVo{
				Id:         iotutil.ToString(event.Id),
				FuncType:   "事件",
				FuncName:   event.EventName,
				Required:   event.Required,
				Identifier: event.Identifier,
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

// 开放平台-合并物模型属性/方法/事件
func (s OpmProductService) mergeOpenPropertyAttributeDesc(propertyRes *protosService.TPmThingModelPropertiesResponseList, serviceRes *protosService.TPmThingModelServicesResponseList, eventRes *protosService.TPmThingModelEventsResponseList, thingModelVoResp []*entitys.TOpmThingModelVo) error {
	var (
		attr  string
		index int
		err   error
	)
	//合并物模型-属性/服务/事件
	if propertyRes != nil && len(propertyRes.List) > 0 {
		for i, property := range propertyRes.List {
			entity := &entitys.TOpmThingModelVo{
				Id:         iotutil.ToString(property.Id),
				FuncType:   "属性",
				FuncName:   property.Name,
				Required:   property.Required,
				Identifier: property.Identifier,
				RwFlag:     property.RwFlag,
				DataType:   property.DataType,
				Space:      property.DataSpecs,
			}
			attr, err = s.transformPropertyAttributeDesc(property.GetDataType(), property.GetDataSpecs(), property.GetDataSpecsList())
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
			entity := &entitys.TOpmThingModelVo{
				Id:           iotutil.ToString(service.Id),
				FuncType:     "服务",
				FuncName:     service.ServiceName,
				Required:     service.Required,
				Identifier:   service.Identifier,
				InputParams:  service.InputParams,
				OutputParams: service.OutputParams,
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
			entity := &entitys.TOpmThingModelVo{
				Id:         iotutil.ToString(event.Id),
				FuncType:   "事件",
				FuncName:   event.EventName,
				Required:   event.Required,
				Identifier: event.Identifier,
				Outputdata: event.Outputdata,
				EventType:  event.EventType,
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

// 转换属性描述
func (OpmProductService) transformPropertyAttributeDesc(dataType string, space string, spaceLists string) (string, error) {
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
	case "ENUM", "BOOL":
		var (
			buff bytes.Buffer
		)
		if dataType == "ENUM" {
			buff.WriteString("枚举值：")
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

// UploadTestReport 上传测试报告
func (s OpmProductService) UploadTestReport(req *entitys.OpmProductTestReportEntitys) error {
	if err := req.AddCheck(); err != nil {
		return err
	}
	res, err := rpc.ClientOpmProductTestReportService.Create(s.Ctx, entitys.OpmProductTestReport_e2pb(req))
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return err
}

// GetTestReport 获取测试报告信息
func (s OpmProductService) GetTestReport(productId int64) (*entitys.OpmProductTestReportEntitys, error) {
	if productId == 0 {
		return nil, errors.New("ProductId不能为空")
	}
	res, err := rpc.ClientOpmProductTestReportService.Lists(s.Ctx, &protosService.OpmProductTestReportListRequest{
		Query: &protosService.OpmProductTestReport{ProductId: productId},
	})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if len(res.Data) == 0 {
		return nil, nil
	}
	return entitys.OpmProductTestReport_pb2e(res.Data[0]), nil
}

// GenerateMcuSdkCode 生成MCU SDK代码
func (s OpmProductService) GenerateMcuSdkCode(ctx context.Context, productId string) (string, error) {
	productInfo, err := s.GetOpmProductDetail(productId)
	if err != nil {
		return "", err
	}
	resp, err := s.QueryProductThingModel(productId, -1)
	if err != nil {
		return "", err
	}
	models := entitys.ThingsModels{}
	for i := range resp.List {
		var dataSpecsList []map[string]interface{}
		if resp.List[i].DataSpecsList != "" {
			if err := json.Unmarshal([]byte(resp.List[i].DataSpecsList), &dataSpecsList); err != nil {
				return "", err
			}
		}
		var values []interface{}
		for i := range dataSpecsList {
			value, ok := dataSpecsList[i]["value"]
			if ok {
				values = append(values, value)
			}
		}
		model := entitys.ThingsModel{
			Dpid:         resp.List[i].Dpid,
			Name:         resp.List[i].Name,
			Identifier:   strings.ToUpper(resp.List[i].Identifier),
			CaseCamel:    iotutil.Case2CamelAndUcfirst(resp.List[i].Identifier),
			IdenLowCase:  strings.ToLower(resp.List[i].Identifier),
			DataType:     "DATA_TYPE_" + iotconst.DataTypeMap[strings.ToUpper(resp.List[i].DataType)],
			Type:         iotutil.Case2CamelAndUcfirst(strings.ToLower(iotconst.DataTypeMap[resp.List[i].DataType])),
			VarType:      iotconst.VarValueType[strings.ToUpper(resp.List[i].DataType)],
			DefaultValue: iotconst.VarDefaultValue[strings.ToUpper(resp.List[i].DataType)],
			RwFlagDesc:   resp.List[i].RwFlagDesc,
			IsControl:    strings.Contains(resp.List[i].RwFlag, "WRITE"),
			Values:       values,
		}
		models = append(models, model)
	}
	sort.Sort(models)
	pe := entitys.ProductEntry{
		ProductKey: productInfo.ProductKey,
		WifiFlag:   productInfo.WifiFlag,
		Models:     models,
	}
	return generateCodeAndZip(pe)
}

func generateCodeAndZip(pe entitys.ProductEntry) (string, error) {
	// 执行shell脚本，拉取模板代码并拷贝目录
	if err := execMcuSdkScript(config.Global.Service.McuSdkDir, pe.ProductKey); err != nil {
		return "", err
	}
	// 读取mcu sdk模板，并执行生成mcu sdk代码
	incTplPath := config.Global.Service.McuSdkDir + "/iot-mcu-sdk-template/template/inc/config.tpl"
	genIncCodePath := config.Global.Service.McuSdkDir + "/" + pe.ProductKey + "/mcu_sdk/inc/"
	if err := generateCode(incTplPath, genIncCodePath, "config.h", pe); err != nil {
		return "", err
	}
	srcTplPath := config.Global.Service.McuSdkDir + "/iot-mcu-sdk-template/template/src/protocol.tpl"
	genSrcCodePath := config.Global.Service.McuSdkDir + "/" + pe.ProductKey + "/mcu_sdk/src/"
	if err := generateCode(srcTplPath, genSrcCodePath, "protocol.c", pe); err != nil {
		return "", err
	}
	srcDir := config.Global.Service.McuSdkDir + "/" + pe.ProductKey + "/mcu_sdk"
	zipFile := config.Global.Service.McuSdkDir + "/" + pe.ProductKey + "/MCU_SDK.zip"
	iotutil.Zip(srcDir, zipFile)
	return zipFile, nil
}

// 根据模板，生成MCU SDK代码
func generateCode(srcPath, desPath, filename string, data interface{}) error {
	tpl := template.Must(template.New(filename).ParseGlob(srcPath))
	if err := iotutil.IsNotExistMkDir(desPath); err != nil {
		return err
	}
	outFile, err := iotutil.Open(desPath+filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer outFile.Close()
	if err := tpl.ExecuteTemplate(outFile, filename, data); err != nil {
		return err
	}
	return nil
}

func execMcuSdkScript(mcuSdkDir, productKey string) error {
	ctx := context.Background()
	c := exec.CommandContext(ctx, config.Global.Service.McuSdkScript, mcuSdkDir, productKey)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go read(ctx, &wg, stderr)
	go read(ctx, &wg, stdout)
	err = c.Start()
	wg.Wait()
	return err
}

func read(ctx context.Context, wg *sync.WaitGroup, std io.ReadCloser) {
	reader := bufio.NewReader(std)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			iotlogger.LogHelper.Helper.Info("exec mcu sdk shell script: ", readString)
		}
	}
}

func (s OpmProductService) ChangeVersionSubmit(req entitys.OpmProductFilter) error {
	if req.Id == 0 {
		return errors.New("id not found")
	}
	if req.FirmwareVersionId == 0 {
		return errors.New("firmwareVersionId not found")
	}
	rep, err := rpc.ClientOpmProductModuleRelationService.ChangeOpmProductModuleRelation(s.Ctx, &protosService.OpmProductModuleRelationChangeVersion{
		ProductId:         req.Id,
		FirmwareId:        req.FirmwareId,
		FirmwareVersionId: req.FirmwareVersionId,
		FirmwareVersion:   req.FirmwareVersion,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: req.Id})
	return nil
}

// QueryProductMap 查询产品信息，Map数据结构，Key=productId
func (s OpmProductService) QueryProductMap(isPlatform bool, productName string) (map[int64]*protosService.OpmProduct, error) {
	proRes, err := rpc.ClientOpmProductService.Lists(s.Ctx, &protosService.OpmProductListRequest{
		IsPlatform: isPlatform,
		SearchKey:  productName,
	})
	if err != nil {
		return nil, err
	}
	if proRes.Code != 200 {
		return nil, errors.New(proRes.Message)
	}
	proMap := make(map[int64]*protosService.OpmProduct, 0)
	for _, pro := range proRes.Data {
		proMap[pro.Id] = pro
	}
	return proMap, nil
}

// GetTaskOrWhereByProduct
func (s OpmProductService) GetTaskOrWhereByProduct(productId int64, condType string) (rets *entitys.ProductThingsModel, err error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	ret, err := rpc.ClientOpmThingModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productId,
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	retPro, err := rpc.ClientOpmProductService.FindById(context.Background(), &protosService.OpmProductFilter{Id: productId})
	if err != nil {
		return nil, err
	}
	if retPro != nil && retPro.Code != 200 {
		return nil, errors.New(retPro.Message)
	}

	var (
		thingsModel        = new(entitys.ProductThingsModel)
		productKey  string = retPro.Data[0].ProductKey
	)

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	dataTransferType, _ := new(services2.DictTempData).GetDictByCode(iotconst.Dict_type_data_transfer_type)
	//抽离公共方法，绑定属性
	var setProperties = func(property *protosService.OpmThingModelProperties) {
		name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, property.Identifier)], property.Name)
		nameEn := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", "en", productKey, property.Identifier)], property.Name)
		rwFlagDesc := dataTransferType.ValueStr(property.RwFlag)
		//是否启用作为条件
		thingsModel.Properties = append(thingsModel.Properties, &entitys.ThingModelProperties{
			Id:            property.Id,
			Name:          name,
			NameEn:        nameEn,
			RwFlag:        property.RwFlag,
			RwFlagDesc:    rwFlagDesc,
			Identifier:    property.Identifier,
			DpId:          property.Dpid,
			DataType:      property.DataType,
			DataSpecs:     property.DataSpecs,
			DefaultVal:    property.DefaultVal,
			DataSpecsList: ConvertJsonByLang(lang, productKey, property.Identifier, property.DataSpecsList, langMap),
		})
	}
	for _, property := range ret.Data.Properties {
		switch condType {
		case "triggerCond":
			//是否启用作为条件
			if property.TriggerCond == 1 {
				setProperties(property)
			}
		case "execCond":
			//是否启用作为执行动作
			if property.ExecCond == 1 {
				setProperties(property)
			}
		case "all":
			setProperties(property)
		}
	}
	return thingsModel, nil
}

func ConvertJsonByLang(lang string, productKey string, identifier string, jsonStr string, langMap map[string]string) string {
	if jsonStr == "" {
		return jsonStr
	}
	var resObj []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &resObj)
	if err != nil {
		panic(err)
	}
	if resObj == nil {
		return jsonStr
	}
	for i, item := range resObj {
		val := iotutil.ToString(item["value"])
		desc := item["desc"]
		name := item["name"]
		if desc == "" || desc == nil {
			desc = name
		}
		dataType := iotutil.ToString(item["dataType"])
		//数值转换（BOOL类型特殊处理）
		if dataType == "BOOL" {
			if val == "1" || val == "true" {
				item["value"] = "true"
			} else if val == "0" {
				item["value"] = "false"
			}
		}
		if langMap != nil {
			langKey := fmt.Sprintf("zh_%s_%s_%v_name", productKey, identifier, item["value"])
			//中文
			desc = iotutil.MapGetStringVal(langMap[langKey], desc)
			resObj[i]["desc"] = desc
			//英文
			langKeyEn := fmt.Sprintf("en_%s_%s_%v_name", productKey, identifier, item["value"])
			nameEn := iotutil.MapGetStringVal(langMap[langKeyEn], name)
			resObj[i]["desc_en"] = nameEn
		} else {
			resObj[i]["desc"] = desc
			resObj[i]["desc_en"] = name
			//resObj[i]["desc_en"] = name
		}
		resObj[i]["value"] = item["value"]
	}
	return iotutil.ToString(resObj)
}

// Export 导出
func (s *OpmProductService) Export(productId string) (string, string, error) {
	//excel 样式文件
	headerStyle := common.ExcelHeaderStyle()
	contentStyle := common.ExcelContentStyle()

	//导出生成excel附件
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("data")

	sheet.SetColWidth(0, 0, 10)
	sheet.SetColWidth(1, 1, 14)
	sheet.SetColWidth(2, 2, 16)
	sheet.SetColWidth(3, 3, 10)
	sheet.SetColWidth(4, 4, 10)
	sheet.SetColWidth(5, 5, 30)

	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "dpid"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "identifier"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "name"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "rwFlag"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	cell.Value = "dataType"
	cell = headerRow.AddCell()
	cell.SetStyle(headerStyle)
	//cell.Value = "properties"
	//cell = headerRow.AddCell()
	cell.Value = "mark"

	res, err := s.QueryProductThingModel(productId, -1)
	if err != nil {
		return "", "", err
	}
	for _, row := range res.List {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = iotutil.ToString(row.Dpid)
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Identifier
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Name
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.RwFlagDesc
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.DataTypeDesc
		//cell = headerRow.AddCell()
		//cell.SetStyle(contentStyle)
		//cell.Value = row.DataSpecs
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.DataSpecsDesc
	}

	tempPathFile := tempPath + iotutil.Uuid() + ".xlsx"
	err = file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		return "", "", err
	}
	//发送完文件后删除对应文件
	//defer func() {
	//	os.Remove(tempPathFile)
	//}()
	fileName := "deviceInfo-" + time.Now().Format("20060102150400") + ".xlsx"
	return fileName, tempPathFile, nil
}

func (s OpmProductService) ResetStandardFunc(productId int64) (string, error) {
	if productId == 0 {
		return "", errors.New("产品编号不存在")
	}
	res, err := rpc.ClientOpmProductService.ResetOpmProductThingsModel(s.Ctx, &protosService.OpmProduct{
		Id: productId,
	})
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	//更新产品修改时间，用于我的产品列表最近更新时间字段
	rpc.ClientOpmProductService.Update(s.Ctx, &protosService.OpmProduct{Id: productId, TslUpdatedAt: timestamppb.Now()})
	return "", err
}

// SetThingsModelSceneFunc 设置场景功能
func (s OpmProductService) SetThingsModelSceneFunc(req entitys.AddStandardThingModelRequest) (string, error) {
	funcList := make([]*protosService.OpmStandardFuncs, 0)
	for _, f := range req.FuncList {
		funcList = append(funcList, &protosService.OpmStandardFuncs{
			FuncId:          f.Id,
			FuncType:        "properties",
			ExecCond:        f.ExecCond,
			ExecCondArgs:    f.ExecCondArgs,
			TriggerCond:     f.TriggerCond,
			TriggerCondArgs: f.TriggerCondArgs,
		})
	}
	saveObj := protosService.OpmThingModel{
		StandardFuncs: funcList,
	}
	res, err := rpc.ClientOpmThingModelService.SetThingsModelSceneFunc(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// SetAppointmentFunc 设置预约功能
func (s OpmProductService) SetAppointmentFunc(req entitys.AddStandardThingModelRequest) (string, error) {
	funcList := make([]*protosService.OpmStandardFuncs, 0)
	for _, f := range req.FuncList {
		funcList = append(funcList, &protosService.OpmStandardFuncs{
			FuncId:           f.Id,
			FuncType:         "properties",
			AppointmentArgs:  f.AppointmentArgs,
			AllowAppointment: f.AllowAppointment,
		})
	}
	saveObj := protosService.OpmThingModel{
		StandardFuncs: funcList,
	}
	res, err := rpc.ClientOpmThingModelService.SetAppointmentFunc(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// SetFuncLevel 设置功能层级
func (s OpmProductService) SetFuncLevel(req entitys.SetFuncLevelRequest) (string, error) {
	funcList := make([]*protosService.OpmStandardFuncs, 0)
	funcList = append(funcList, &protosService.OpmStandardFuncs{
		FuncId:   req.Id,
		FuncType: "properties",
		IsTop:    req.IsTop,
		ParentId: req.ParentId,
	})
	saveObj := protosService.OpmThingModel{
		StandardFuncs: funcList,
	}
	res, err := rpc.ClientOpmThingModelService.SetFuncLevel(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// SetFuncSort 设置功能排序
func (s OpmProductService) SetFuncSort(req entitys.SetFuncSortRequest) (string, error) {
	properties := &protosService.OpmThingModelProperties{
		Id:       req.Id,
		FuncType: "properties",
		Sort:     req.Sort,
	}
	saveObj := protosService.OpmThingModel{
		Properties: properties,
	}
	res, err := rpc.ClientOpmThingModelService.SetFuncSort(s.Ctx, &saveObj)
	if err != nil {
		return "", err
	}
	if res.Code != 200 {
		return "", errors.New(res.Message)
	}
	if err != nil {
		return "", err
	}
	err = cached.RedisStore.Delete(persist.GetRedisKey(iotconst.OPEN_PRODUCT_FUNC_LIST_DATA, req.ProductId))
	if err != nil {
		return "", err
	}
	return "", err
}

// QueryAppointmentFuncList 开放平台-获取产品基础物模型数据
func (s OpmProductService) QueryAppointmentFuncList(productId string) (*entitys.OpmThingModelList, error) {
	if productId == "" {
		return nil, errors.New("产品编号不存在")
	}
	productIdStr := iotutil.ToInt64(productId)
	req, err := rpc.ClientOpmThingModelService.ProductThingModel(s.Ctx, &protosService.OpmThingModelByProductRequest{
		ProductId: productIdStr,
		Custom:    -1,
	})
	if err != nil {
		return nil, err
	}
	if req.Code != 200 {
		return nil, errors.New(req.Message)
	}

	result := new(entitys.OpmThingModelList)
	for _, property := range req.Data.Properties {
		if property.AllowAppointment == 1 {
			result.List = append(result.List, entitys.OpmThingModelProperties_pb2e(property))
		}
	}
	sort.Slice(result.List, func(i, j int) bool {
		return result.List[i].Sort < result.List[j].Sort
	})
	return result, err
}
