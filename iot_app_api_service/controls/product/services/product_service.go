package services

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/product/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	goerrors "go-micro.dev/v4/errors"
	"sort"

	"go-micro.dev/v4/metadata"
)

type ProductService struct {
	Ctx context.Context
}

func (s ProductService) SetContext(ctx context.Context) ProductService {
	s.Ctx = ctx
	return s
}

// GetProductByApp get Product list  data
func (s ProductService) GetProductByApp(filter entitys.AppQueryProductForm, userId, language string) (rets []*entitys.AppProductDto, total int, err error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)

	userRet, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: iotutil.ToInt64(userId),
	})
	if err != nil {
		return nil, 0, err
	}
	if len(userRet.Data) == 0 {
		return nil, 0, errors.New("当前用户参数异常")
	}

	ret, err := rpc.ProductService.AppLists(s.Ctx, &protosService.AppOpmProductListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.AppOpmProduct{
			WifiFlag:        filter.WifiFlag,
			ProductTypeId:   filter.ProductTypeId,
			BaseProductId:   filter.BaseProductId,
			ProductTypeName: filter.ProductTypeName,
			WifiFlags:       filter.WifiFlags},
	})

	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Message)
	}
	homeId := iotutil.ToInt64(userRet.Data[0].DefaultHomeId)
	nickName := userRet.Data[0].NickName
	account := userRet.Data[0].UserName
	homeName := userRet.Data[0].DefaultHomeName
	userIdInt := iotutil.ToInt64(userId)

	homeRes, err := rpc.UcHomeService.FindById(context.Background(), &protosService.UcHomeFilter{Id: homeId})
	if err != nil {
		return nil, 0, err
	}
	if homeRes != nil && homeRes.Code != 200 {
		return nil, 0, errors.New(homeRes.Message)
	}

	//指定缓存key获取
	//langMap := make(map[string]string)
	//if lang != "" {
	//	sourceRowIds := []string{}
	//	for _, data := range ret.Data {
	//		sourceRowIds = append(sourceRowIds, fmt.Sprintf("%s_%d_name", lang, data.Id))
	//	}
	//	slice, err := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_PM_PRODUCT_TYPE, sourceRowIds...).Result()
	//	if err == nil {
	//		langMap = iotutil.ArrayUnionInterfaces(sourceRowIds, slice)
	//	}
	//}

	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_NAME)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	resultList := make([]*entitys.AppProductDto, 0)
	for _, pro := range ret.Data {
		info := entitys.AppProductDto{}
		info.Id = iotutil.ToString(pro.Id)
		info.ProductTypeId = iotutil.ToString(pro.ProductTypeId)
		info.Name = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_name", lang, pro.ProductKey)], pro.Name)
		deviceToken, _ := GenerateNetworkToken(iotstruct.DeviceNetworkTokenCacheModel{
			UserId:       userIdInt,
			ProductId:    pro.Id,
			DeviceNature: pro.DeviceNatureKey,
			HomeId:       homeId,
			UserName:     nickName,
			Account:      account,
			HomeName:     homeName,
			ProductName:  info.Name,
			ProductKey:   pro.ProductKey,
			Devices:      make([]string, 0),
			DevicesMap:   make(map[string]iotstruct.DeviceResult),
			Lat:          homeRes.Data[0].Lat,
			Lng:          homeRes.Data[0].Lng,
			Country:      homeRes.Data[0].Country,
			Province:     homeRes.Data[0].Province,
			City:         homeRes.Data[0].City,
			District:     homeRes.Data[0].District,
			AppKey:       appKey,
			TenantId:     pro.TenantId,
		})
		//if language == "en" {
		//	info.Name = pro.NameEn
		//}
		info.Model = pro.ProductKey
		info.ImageUrl = controls.ConvertProImg(pro.ImageUrl)
		info.WifiFlag = pro.WifiFlag
		info.NetworkType = iotutil.ToInt32(pro.NetworkType)
		info.Token = deviceToken
		//TODO 不需要额外增加 DistributionType 直接使用 NetworkType
		networkTypeInt := iotutil.ToInt32(pro.NetworkType)
		switch networkTypeInt {
		case 1, 2:
			info.DistributionType = 1
		case 3, 4:
			info.DistributionType = 3
		default:
			info.DistributionType = 4
		}
		//info.DistributionType = iotutil.ToInt32(pro.NetworkType)
		resultList = append(resultList, &info)
	}
	return resultList, total, err
}

// QueryProductNetworkGuide 查询产品的配网引导
func (s ProductService) QueryProductNetworkGuide(id string) (*entitys.OpmNetworkGuideEntitys, error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	//networkType 从产品获取
	rid := iotutil.ToInt64(id)
	productResp, err := rpc.ProductService.FindById(context.Background(), &protosService.OpmProductFilter{
		Id: rid,
	})
	if err != nil {
		return nil, err
	}
	if productResp.Code != 200 {
		return nil, errors.New(productResp.Message)
	}
	productInfo := productResp.Data[0]
	networkType := productInfo.NetworkType

	networkTypeInt := iotutil.ToInt32(networkType)
	distributionType := 0
	switch networkTypeInt {
	case 1, 2:
		distributionType = 1
	case 3, 4:
		distributionType = 3
	default:
		distributionType = 1
	}

	opmNetworkGuideResp, err := rpc.ClientOpmNetworkGuideService.FindByProductId(context.Background(), &protosService.OpmNetworkGuideFilter{
		Id: rid,
	})
	if err != nil {
		return nil, err
	}
	if opmNetworkGuideResp.Code != 200 {
		return nil, errors.New(opmNetworkGuideResp.Message)
	}
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_NETWORK_GUIDE)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	dataInfo := entitys.OpmNetworkGuideEntitys{}
	dataInfo.NetworkMode = iotutil.ToInt32(productInfo.NetworkType)
	dataInfo.WifiFlag = productInfo.WifiFlag
	dataInfo.ProductKey = productInfo.ProductKey
	isSelected := false
	for _, data := range opmNetworkGuideResp.Data {
		if data.Type == int32(distributionType) {
			setNetworkGuide(lang, data, langMap, &dataInfo)
			isSelected = true
			break
		}
	}
	//如果没有选择则默认显示第一个
	if isSelected == false && len(opmNetworkGuideResp.Data) > 0 {
		setNetworkGuide(lang, opmNetworkGuideResp.Data[0], langMap, &dataInfo)
	}
	return &dataInfo, err
}

func setNetworkGuide(lang string, data *protosService.OpmNetworkGuide, langMap map[string]string, res *entitys.OpmNetworkGuideEntitys) {
	steps := []*entitys.OpmNetworkGuideStepEntitys{}
	for _, step := range data.Steps {
		item := entitys.OpmNetworkGuideStep_pb2e(step, lang)
		langKey := fmt.Sprintf("%s_%d_instruction", lang, step.Id)
		item.Instruction = iotutil.MapGetStringVal(langMap[langKey], item.Instruction)
		steps = append(steps, item)
	}
	//根据sort进行排序
	sort.Slice(steps, func(i, j int) bool {
		return steps[i].Sort < steps[j].Sort
	})
	res.Id = data.Id
	res.ProductId = data.ProductId
	res.Type = data.Type
	res.Steps = steps
}

// GetProductList get Product list  data
func (s ProductService) GetOpmProductList(filter entitys.OpmProductQuery) (rets []*entitys.OpmProductEntitys, total int64, err error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
	)

	ret, err := rpc.ProductService.Lists(s.Ctx, &protosService.OpmProductListRequest{
		Page:     int64(filter.Page),
		PageSize: int64(filter.Limit),
		Query: &protosService.OpmProduct{
			TenantId: tenantId,
		},
	})

	if err != nil {
		return nil, 0, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, 0, errors.New(ret.Message)
	}

	resultList := make([]*entitys.OpmProductEntitys, 0)
	for i, _ := range ret.Data {
		item := entitys.OpmProduct_pb2e(ret.Data[i])
		if lang == "en" {
			item.Name = item.NameEn
		}
		resultList = append(resultList, item)
	}
	return resultList, ret.Total, err
}

// GetFunctionRules 获取功能规则
func (s ProductService) GetFunctionRules(productKey string, dataOrigin int32, dpidsMap map[int32]*protosService.OpmThingModelProperties) ([]entitys.ThingModelRuleItemResponse, error) {
	ret, err := rpc.ClientOpmThingModelRuleService.Lists(s.Ctx, &protosService.OpmThingModelRuleListRequest{
		Query: &protosService.OpmThingModelRule{
			ProductKey: productKey,
			DataOrigin: dataOrigin,
			Status:     1, //获取已启用的数据
		},
	})
	if err != nil {
		return nil, err
	}
	if ret != nil && ret.Code != 200 {
		return nil, errors.New(ret.Message)
	}

	resMap := make([]entitys.ThingModelRuleItemResponse, 0)
	for _, d := range ret.Data {
		var ifSpecs []entitys.ThingModelRuleItem
		err := iotutil.JsonToStruct(d.IfSpecs, &ifSpecs)
		if err != nil {
			continue
		}
		var thenSpecs []entitys.ThingModelRuleItem
		err = iotutil.JsonToStruct(d.ThatSpecs, &thenSpecs)
		if err != nil {
			continue
		}
		ifSpecRes := []entitys.ThingModelRuleItem{}
		for _, d := range ifSpecs {
			tsIdentifier := d.Identifier
			if dpidsMap != nil {
				if v, ok := dpidsMap[int32(d.DpId)]; ok {
					tsIdentifier = v.Identifier
				} else {
					continue
				}
			}
			ifSpecRes = append(ifSpecRes, entitys.ThingModelRuleItem{
				DpId:       d.DpId,
				Identifier: tsIdentifier,
				Operate:    d.Operate,
				Value:      d.Value,
			})
		}
		ifSpec := ifSpecs[0]
		thenSpecRes := []entitys.ThingModelRuleItem{}
		for _, d := range thenSpecs {
			tsIdentifier := d.Identifier
			if dpidsMap != nil {
				if v, ok := dpidsMap[int32(d.DpId)]; ok {
					tsIdentifier = v.Identifier
				} else {
					continue
				}
			}
			thenSpecRes = append(thenSpecRes, entitys.ThingModelRuleItem{
				DpId:       d.DpId,
				Identifier: tsIdentifier,
				Operate:    d.Operate,
				Value:      d.Value,
			})
		}
		if len(thenSpecRes) == 0 {
			continue
		}
		resMap = append(resMap, entitys.ThingModelRuleItemResponse{
			DpId:          ifSpec.DpId,
			Identifier:    ifSpec.Identifier,
			Operate:       ifSpec.Operate,
			Value:         ifSpec.Value,
			ConditionType: ret.Data[0].ConditionType,
			IfSpecs:       ifSpecRes,
			Specs:         thenSpecRes,
			Sort:          ifSpec.DpId,
		})
	}
	sort.Slice(resMap, func(i, j int) bool {
		return resMap[i].Sort > resMap[j].Sort
	})
	if productKey != "" {
		resMap, _ = checkMakeUpFuncs(productKey, resMap)
	}
	return resMap, err
}

func getProductTsl(productKey string) (map[string]string, error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey)
	return strCmd.Val(), strCmd.Err()
}

// checkMakeUpFuncs 检查补全Functions
func checkMakeUpFuncs(productKey string, rules []entitys.ThingModelRuleItemResponse) ([]entitys.ThingModelRuleItemResponse, error) {
	//通过产品Key获取功能缓存
	tslMap, err := getProductTsl(productKey)
	if err != nil {
		return rules, goerrors.New("", err.Error(), ioterrs.ErrDevJsonMarshal)
	}
	for i, f := range rules {
		if v, ok := tslMap[fmt.Sprintf("tls_%v", f.DpId)]; ok {
			funcsMap, err := iotutil.JsonToMapErr(v)
			if err == nil {
				if funcsMap["identifier"] != nil {
					f.Identifier = funcsMap["identifier"].(string)
				}
			}
		}
		for i2, spec := range f.IfSpecs {
			if v, ok := tslMap[fmt.Sprintf("tls_%v", spec.DpId)]; ok {
				funcsMap, err := iotutil.JsonToMapErr(v)
				if err == nil {
					if funcsMap["identifier"] != nil {
						f.IfSpecs[i2].Identifier = funcsMap["identifier"].(string)
					}
				}
			}
		}
		for i2, spec := range f.Specs {
			if v, ok := tslMap[fmt.Sprintf("tls_%v", spec.DpId)]; ok {
				funcsMap, err := iotutil.JsonToMapErr(v)
				if err == nil {
					if funcsMap["identifier"] != nil {
						f.Specs[i2].Identifier = funcsMap["identifier"].(string)
					}
				}
			}
		}
		rules[i] = f
	}
	return rules, nil
}

// QueryOpmDocumentsList 列表
func (s ProductService) QueryOpmDocumentsList(productKey, docCodes string) ([]*entitys.OpmDocumentsEntitys, error) {
	rep, err := rpc.ClientOpmDocumentsService.Lists(s.Ctx, &protosService.OpmDocumentsListRequest{
		Query: &protosService.OpmDocuments{OriginKey: productKey, DocCodes: docCodes},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	var resultList = []*entitys.OpmDocumentsEntitys{}
	for _, item := range rep.Data {
		resultList = append(resultList, entitys.OpmDocuments_pb2e(item))
	}
	return resultList, err
}
