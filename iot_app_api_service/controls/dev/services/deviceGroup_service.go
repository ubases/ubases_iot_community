package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	services2 "cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"go-micro.dev/v4/logger"
)

type AppDeviceGroupService struct {
	Ctx context.Context
}

func (s AppDeviceGroupService) SetContext(ctx context.Context) AppDeviceGroupService {
	s.Ctx = ctx
	return s
}

// DevListByProductKey 通过产品Key获取设备列表
func (s AppDeviceGroupService) DevListByProductKey(productKey, homeId, groupIdStr string) ([]entitys.DeviceGroupListDto, error) {
	groupId, _ := iotutil.ToInt64AndErr(groupIdStr)

	//如果groupId不为0，则获取已分配的设备Id
	devIds, err := s.getGroupDevices(groupId, homeId)
	if err != nil {
		return nil, err
	}

	//获取产品信息
	productData, err := rpc.ProductService.Find(s.Ctx, &protosService.OpmProductFilter{
		ProductKey: productKey,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(err)
		return nil, err
	}
	if productData.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(productData.Message)
		return nil, errors.New("record not found")
	}

	//获取同类型的产品
	products, productKeys, err := s.getSameTypeProduct(productKey)
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(err)
		return nil, err
	}
	//获取家庭信息
	ret, err := s.getHomeInfo(homeId)
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	resDevList := make([]entitys.DeviceGroupListDto, 0) //需要返回的设备列表
	for _, d := range ret.DeviceList {
		devInfo := d.Data
		//检查产品Key是否在获取的产品列表中
		//检查设备Id是否在已分配的设备Id列表中
		if iotutil.IsContainsByList(productKeys, devInfo.ProductKey) == false || iotutil.IsContainsByList(devIds, devInfo.Did) == true {
			continue
		}
		resDev := entitys.DeviceGroupListDto{}
		resDev.DevName = devInfo.DeviceName
		for _, productInfo := range products {
			if devInfo.ProductKey == productInfo.ProductKey {
				resDev.Pic = productInfo.ImageUrl
				break
			}
		}
		resDev.DevId = devInfo.Did
		resDev.ProductKey = devInfo.ProductKey
		resDevList = append(resDevList, resDev)
	}
	return resDevList, nil
}

// 获取分组已存在设备Id
func (s AppDeviceGroupService) getGroupDevices(groupId int64, homeId string) ([]string, error) {
	devIds := make([]string, 0)
	if groupId != 0 {
		dgRes, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
			Query: &protosService.IotDeviceGroupList{
				GroupId: groupId,
				HomeId:  iotutil.ToInt64(homeId),
			},
		})
		if err != nil {
			return nil, err
		}
		if dgRes.Code == 200 {
			for _, deviceGroupListInfo := range dgRes.Data {
				devIds = append(devIds, deviceGroupListInfo.DevId)
			}
		}
	}
	return devIds, nil
}

// 获取家庭信息
func (s AppDeviceGroupService) getHomeInfo(homeId string) (*protosService.UcHomeDetail, error) {
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: iotutil.ToInt64(homeId),
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return nil, err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
		return nil, errors.New("record not found")
	}
	return ret.Data, nil
}

// 获取相同类型的产品信息
func (s AppDeviceGroupService) getSameTypeProduct(productKey string) ([]*protosService.OpmProduct, []string, error) {
	//获取产品信息
	productData, err := rpc.ProductService.Find(s.Ctx, &protosService.OpmProductFilter{
		ProductKey: productKey,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(err)
		return nil, nil, err
	}
	if productData.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(productData.Message)
		return nil, nil, errors.New("record not found")
	}

	productRes, err := rpc.ProductService.Lists(s.Ctx, &protosService.OpmProductListRequest{
		Query: &protosService.OpmProduct{ProductTypeId: productData.Data[0].ProductId}, //这里没问题，是用base_product_id这个来比较
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(err)
		return nil, nil, err
	}
	if productData.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "DevListByProductKey").Error(productData.Message)
		return nil, nil, errors.New("record not found")
	}
	productKeys := make([]string, 0)
	for _, productInfo := range productRes.Data {
		productKeys = append(productKeys, productInfo.ProductKey)
	}
	return productRes.Data, productKeys, nil
}

// UpsertGroup 保存设备群组
func (s AppDeviceGroupService) UpsertGroup(req entitys.UpsertGroup, userId int64) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	ugInfo := entitys.UpsertGroup_db2pb(req)
	ugInfo.UserId = userId
	ugInfo.HomeId = iotutil.ToInt64(req.HomeId)
	ugRes, err := rpc.IotDeviceGroupService.UpsertGroup(context.Background(), ugInfo)
	if err != nil {
		return result, err
	}
	if ugRes.Code != 200 {
		return result, errors.New(ugRes.Message)
	}
	result["groupId"] = iotutil.ToString(ugRes.Data[0].Id)
	// 删除家庭详情缓存
	keys := make([]string, 0)
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: ugInfo.HomeId,
		},
	})
	if err != nil {
		return result, err
	}
	if resp.Code != 200 {
		return result, errors.New(resp.Message)
	}
	// 清理家庭信息缓存
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, req.HomeId, iotutil.ToString(resp.Data[i].UserId)))
	}
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// DevGroupDevList 设备群组设备列表
func (s AppDeviceGroupService) DevGroupDevList(groupId int64) ([]entitys.DeviceGroupListDto, error) {
	result := make([]entitys.DeviceGroupListDto, 0)
	//获取群组的家庭Id，建议改为前端传入
	dgData, err := s.getDeviceGroupInfo(groupId)
	if err != nil {
		return result, err
	}
	homeId := dgData.HomeId

	dgListRes, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
		Query: &protosService.IotDeviceGroupList{
			GroupId: groupId,
			HomeId:  homeId,
		},
	})
	if err != nil {
		return result, err
	}
	if dgListRes.Code != 200 {
		return result, nil
	}

	//获取分组中的设备Id
	devIds := make([]string, 0)
	for _, d := range dgListRes.Data {
		devIds = append(devIds, d.DevId)
	}
	deviceListRes, err := rpc.IotDeviceInfoService.DeviceInfoListByDevIds(context.Background(), &protosService.DeviceInfoListByDevIdsFilter{DevIds: devIds})
	deviceMap := make(map[string]*protosService.IotDeviceInfo)
	productIds := make([]int64, 0)
	for _, deviceInfo := range deviceListRes.Data {
		deviceMap[deviceInfo.Did] = deviceInfo
		productIds = append(productIds, deviceInfo.ProductId)
	}
	productRes, err := rpc.ProductService.ListsByProductIds(context.Background(), &protosService.ListsByProductIdsRequest{
		ProductIds: productIds,
	})
	productMap := make(map[string]string)
	for _, p := range productRes.Data {
		productMap[p.ProductKey] = p.ImageUrl
	}
	return entitys.DevGroupInfo_pb2db(dgListRes.Data, deviceMap, productMap), nil
}

// 获取设备群组信息
func (s AppDeviceGroupService) getDeviceGroupInfo(groupId int64) (*protosService.IotDeviceGroup, error) {
	dgRes, err := rpc.IotDeviceGroupService.FindById(context.Background(), &protosService.IotDeviceGroupFilter{
		Id: groupId,
	})
	if err != nil {
		return nil, err
	}
	if dgRes.Code != 200 {
		return nil, errors.New(dgRes.Message)
	}
	return dgRes.Data[0], nil
}

func (s AppDeviceGroupService) getDeviceGroupList(groupId int64) {

}

// RemoveGroup 移除设备分组
func (s AppDeviceGroupService) RemoveGroup(groupId, userId int64) error {
	//查询是为了获取群组的名称
	dgData, err := s.getDeviceGroupInfo(groupId)
	if err != nil {
		return err
	}

	//删除群组和群组设备列表
	_, err = rpc.IotDeviceGroupService.Delete(context.Background(), &protosService.IotDeviceGroup{
		Id: groupId,
		//UserId:    userId,
	})
	if err != nil {
		return err
	}

	respUser, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{
		Id: userId,
	})
	if err != nil {
		return err
	}

	// TODO 删除家庭详情缓存，需要通用家庭缓存清理
	keys := []string{}
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: iotutil.ToInt64(respUser.Data[0].DefaultHomeId),
		},
	})
	if err != nil {
		return err
	}

	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, respUser.Data[0].DefaultHomeId, iotutil.ToString(resp.Data[i].UserId)))
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}

	go services2.SendDisbandGroupMessage(services2.SetAppInfoByContext(s.Ctx), userId, iotutil.ToInt64(respUser.Data[0].DefaultHomeId), dgData.Name)
	return nil
}

// Execute 执行控制
func (s AppDeviceGroupService) Execute(req entitys.GroupExecute) error {
	devIds := make([]string, 0)
	//获取已分配的设备Id
	dgRes, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
		Query: &protosService.IotDeviceGroupList{
			GroupId: iotutil.ToInt64(req.GroupId),
		},
	})
	if err != nil {
		return err
	}
	if dgRes.Code != 200 {
		return errors.New(dgRes.Message)
	}
	for _, d := range dgRes.Data {
		devIds = append(devIds, d.DevId)
	}

	deviceRes, err := rpc.IotDeviceInfoService.DeviceInfoListByDevIds(context.Background(), &protosService.DeviceInfoListByDevIdsFilter{DevIds: devIds})

	if err != nil {
		logger.Errorf("execute error : %s", err.Error())
		return err
	}
	if deviceRes.Code != 200 {
		logger.Errorf(deviceRes.Message)
		return errors.New(deviceRes.Message)
	}

	deviceMap := make(map[string]string)
	for _, d := range deviceRes.Data {
		deviceMap[d.Did] = d.ProductKey
	}

	appHomeSvc := services2.AppHomeService{}
	for _, deviceId := range devIds {
		//设备正在进行升级，无法推送群组控制
		devState := appHomeSvc.GetDeviceStatus(deviceId)
		if devState.HasForceUpgrade {
			iotlogger.LogHelper.Infof("device need upgrade : devId: %s, push error", deviceId)
			continue
		}
		//包装控制指令推送控制
		productKey := deviceMap[deviceId]
		//pushData := iotutil.MapStringToInterface(map[string]string{
		//	req.Dpid: iotutil.ToString(req.Value),
		//})
		_, _, err = PubControl(productKey, deviceId, map[string]interface{}{
			req.Dpid: req.Value,
		})
		if err != nil {
			logger.Errorf("rpcClient.PubControl.Publish error : %s", err.Error())
		}
	}

	//更新群组属性信息
	devInfo := make(map[string]string)
	devInfo[iotconst.FIELD_PREFIX_DPID+req.Dpid] = iotutil.ToString(req.Value)
	//缓存物模型数据
	rdCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_GROUP_DATA+req.GroupId, devInfo)
	if rdCmd.Err() != nil {
		return rdCmd.Err()
	}

	return nil
}

// DevGroupInfo 设备群组信息
func (s AppDeviceGroupService) DevGroupInfo(groupId int64) (entitys.DeviceGroup, error) {
	result := entitys.DeviceGroup{}
	dgData, err := s.getDeviceGroupInfo(groupId)
	if err != nil {
		return result, err
	}
	var homeId int64
	if dgData != nil {
		homeId = dgData.HomeId
	}

	deviceGroupListResponse, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
		Query: &protosService.IotDeviceGroupList{
			GroupId: groupId,
			HomeId:  homeId,
		},
	})
	if err != nil {
		return result, err
	}
	if deviceGroupListResponse.Code != 200 {
		//iotlogger.LogHelper.WithTag("method", "DevGroupInfo").Error(deviceGroupListResponse.Message)
		return result, nil
	}
	deviceGroupList := deviceGroupListResponse.Data[0]

	deviceHomeResp, err := rpc.IotDeviceHomeService.Find(context.Background(), &protosService.IotDeviceHomeFilter{
		DeviceId: deviceGroupList.DevId,
	})
	if err != nil {
		iotlogger.LogHelper.Infof("调用IotDeviceHomeService.Find异常，%s", err.Error())
		return result, err
	}
	if deviceHomeResp.Code != 200 {
		iotlogger.LogHelper.Infof("调用IotDeviceHomeService.Find异常，- %s", deviceHomeResp.Message)
		return result, nil
	}

	productInfo, err := rpc.ProductService.FindById(context.Background(), &protosService.OpmProductFilter{
		Id: deviceHomeResp.Data[0].ProductId,
	})
	if err != nil {
		iotlogger.LogHelper.Infof("调用ProductService.FindById异常，%s", err.Error())
		return result, err
	}
	if productInfo.Code != 200 {
		iotlogger.LogHelper.Infof("调用ProductService.FindById异常，- %s", productInfo.Message)
		return result, nil
	}

	deviceGroupInfo := entitys.DeviceGroup{
		Id:            iotutil.ToString(dgData.Id),
		Name:          dgData.Name,
		RoomId:        iotutil.ToString(dgData.RoomId),
		RoomName:      dgData.RoomName,
		HomeId:        iotutil.ToString(dgData.HomeId),
		Time:          dgData.CreatedAt.AsTime().Unix(),
		UserId:        iotutil.ToString(dgData.UserId),
		DevCount:      len(deviceGroupListResponse.Data),
		Pic:           productInfo.Data[0].ImageUrl,
		ProductKey:    productInfo.Data[0].ProductKey,
		ProductTypeId: "",
	}

	return deviceGroupInfo, nil
}

// DevGroupTsl 获取群组设备的功能
func (s AppDeviceGroupService) DevGroupTsl(groupId int64, language, tenantId string) ([]entitys.TslInfo, error) {
	dgData, err := s.getDeviceGroupInfo(groupId)
	if err != nil {
		return nil, err
	}
	homeId := dgData.HomeId

	dgListRes, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
		Query: &protosService.IotDeviceGroupList{
			GroupId: groupId,
			HomeId:  homeId,
		},
	})
	if err != nil {
		return nil, err
	}
	if dgListRes.Code != 200 {
		return nil, nil
	}

	devIds := make([]string, 0)
	for _, d := range dgListRes.Data {
		devIds = append(devIds, d.DevId)
	}
	//获取设备对应的产品Id
	productIds := make([]int64, 0)
	//兼容逻辑，如果存在ProductKey则直接使用，不需要再次查询ProductId
	deviceRes, err := rpc.IotDeviceInfoService.DeviceInfoListByDevIds(context.Background(), &protosService.DeviceInfoListByDevIdsFilter{DevIds: devIds})
	if err != nil {
		return nil, err
	}
	if len(deviceRes.Data) == 0 {
		return nil, errors.New("异常数据")
	}
	for _, deviceInfo := range deviceRes.Data {
		productIds = append(productIds, deviceInfo.ProductId)
	}
	productIds = iotutil.RemoveRepeatInt64Element(productIds)

	thingsModelRes, err := rpc.ProductService.MergeProductThingsModel(context.Background(), &protosService.ListsByProductIdsRequest{ProductIds: productIds})
	if err != nil {
		return nil, err
	}
	if thingsModelRes.Code != 200 {
		return nil, errors.New(thingsModelRes.Message)
	}

	//读取物模型的缓存
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()

	//获取功能属性设置值
	funcDescMap := deviceRes.Data[0].FuncDescMap

	//读取群组控制的值
	strCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_GROUP_DATA+iotutil.ToString(groupId))

	tslList := make([]entitys.TslInfo, 0)
	for _, v := range thingsModelRes.Data.Properties {
		//过滤掉只上报的
		if v.RwFlag == "READ" {
			continue
		}
		var dpIdValue string
		if val, ok := strCmd.Val()[iotconst.FIELD_PREFIX_DPID+iotutil.ToString(v.Dpid)]; ok {
			dpIdValue = iotutil.ToString(val)
		}
		name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", language, v.ProductKey, v.Identifier)], v.Name)
		if funcDescMap != nil {
			if v, ok := funcDescMap[fmt.Sprintf("%s_%s", v.ProductKey, v.Dpid)]; ok && v != "" {
				name = v
			}
		}
		//变更默认值的返回
		var defaultVal interface{}
		switch v.DataType {
		case "ENUM", "INT":
			defaultVal, _ = iotutil.ToInt32Err(v.DefaultVal)
		case "BOOL":
			if v.DataType == "1" || v.DataType == "true" {
				defaultVal = "true"
			} else {
				defaultVal = "false"
			}
		case "DOUBLE", "FLOAT":
			defaultVal, _ = iotutil.ToFloat64Err(v.DefaultVal)
		default:
			defaultVal = iotutil.ToString(v.DefaultVal)
		}
		tslList = append(tslList, entitys.TslInfo{
			DpId:          v.Dpid,
			DataType:      v.DataType,
			Name:          name,
			RwFlag:        "",
			DataSpecs:     v.DataSpecs,
			DataSpecsList: ConvertJsonByLang(language, v.ProductKey, v.Dpid, v.Identifier, v.DataSpecsList, langMap, funcDescMap),
			Required:      0,
			Identifier:    v.Identifier,
			Value:         dpIdValue,
			DefaultVal:    defaultVal,
		})
	}
	//根据sort进行排序
	sort.Slice(tslList, func(i, j int) bool {
		return tslList[i].DpId > tslList[j].DpId
	})
	return tslList, nil
}

func ConvertJsonByLang(lang string, productKey string, dpid int32, identifier string, jsonStr string, langMap map[string]string, funcSetMap map[string]string) string {
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
		val := item["value"]
		valStr := iotutil.ToString(val)
		desc := item["desc"]
		name := item["name"]
		if desc == "" || desc == nil {
			desc = name
		}
		dataType := iotutil.ToString(item["dataType"])
		//数值转换（BOOL类型特殊处理）
		if dataType == "BOOL" {
			if valStr == "1" || valStr == "true" {
				val = "true"
			} else {
				val = "false"
			}
		}
		//如果有设置自定名称，则已自定义名称显示，否则使用翻译显示；
		if funcSetMap != nil {
			if v, ok := funcSetMap[fmt.Sprintf("%s_%v_%v", productKey, dpid, val)]; ok && v != "" {
				name = v
				resObj[i]["name"] = name
				resObj[i]["desc"] = name
				continue
			}
		}
		langKey := fmt.Sprintf("%s_%s_%s_%v_name", lang, productKey, identifier, val)
		desc = iotutil.MapGetStringVal(langMap[langKey], desc)
		resObj[i]["name"] = name
		resObj[i]["desc"] = desc
		resObj[i]["value"] = val
	}
	return iotutil.ToString(resObj)
}

// 数据列表转换为翻译数据
func (s AppDeviceGroupService) convertDataTranslate(productKey string, res []*protosService.OpmThingModelProperties, langTypes []*entitys.DictKeyVal, language string) ([]*protosService.OpmThingModelProperties, error) {
	//将物模型数据转换为语言key，并查询出翻译内容
	langList, err := s.getTranslateData(productKey, res)
	if err != nil {
		return nil, err
	}

	result := []*protosService.OpmThingModelProperties{}
	for _, data := range res {
		langKey := fmt.Sprintf("%s_%s", productKey, data.Identifier)
		var langs []*entitys.LangTranslateEntitys
		if val, ok := langList[langKey]; ok {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "identifier", data.Name, "", langTypes, val)
		} else {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "identifier", data.Name, "", langTypes, []*entitys.LangTranslateEntitys{})
		}
		for _, lang := range langs {
			if lang.Lang == language {
				data.Name = lang.FieldValue
				break
			}
		}
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err = json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					funcLangKey := fmt.Sprintf("%s_%s_%s", productKey, data.Identifier, val)
					var langSpecs []*entitys.LangTranslateEntitys
					if val, ok := langList[funcLangKey]; ok {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, funcLangKey, "identifier_value", "", "", langTypes, val)
					} else {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, funcLangKey, "identifier_value", "", "", langTypes, []*entitys.LangTranslateEntitys{})
					}
					for _, lang := range langSpecs {
						if lang.Lang == language {
							spec["name"] = lang.FieldValue
						}
						break
					}
				}
				data.DataSpecsList = iotutil.ObjToString(mapSpecs)
			}
		}
		result = append(result, data)
	}
	return result, nil
}

func (s AppDeviceGroupService) SetLangs(sourceTable, sourceRowId, fieldName, fieldValue, fieldValueEn string, langTypes []*entitys.DictKeyVal, data []*entitys.LangTranslateEntitys) []*entitys.LangTranslateEntitys {
	defaultData := []*entitys.LangTranslateEntitys{}
	for _, langType := range langTypes {
		var currData *entitys.LangTranslateEntitys
		for _, d := range data {
			if langType.Code == d.Lang {
				currData = d
			}
		}
		defaultItem := &entitys.LangTranslateEntitys{
			Lang:       langType.Code,
			FieldName:  fieldName,
			FieldValue: "",
		}
		if langType.Code == "zh" {
			defaultItem.FieldValue = fieldValue
		} else if langType.Code == "en" {
			defaultItem.FieldValue = fieldValueEn
		}
		if currData != nil {
			defaultItem = currData
		}
		defaultData = append(defaultData, defaultItem)
	}
	return defaultData
}

// 将物模型数据转换为语言key，并查询出翻译内容
func (s AppDeviceGroupService) getTranslateData(productKey string, res []*protosService.OpmThingModelProperties) (map[string][]*entitys.LangTranslateEntitys, error) {
	sourceIds := []string{}
	for _, data := range res {
		sourceIds = append(sourceIds, fmt.Sprintf("%s_%s", productKey, data.Identifier))
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err := json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					dataType := iotutil.ToString(spec["dataType"])
					if dataType == "BOOL" {
						if val == "1" || val == "true" {
							val = "true"
						} else {
							val = "false"
						}
					}
					langKey := fmt.Sprintf("%s_%s_%s", productKey, data.Identifier, val)
					sourceIds = append(sourceIds, langKey)
				}
			}
		}
	}
	rep, err := rpc.ClientLangTranslateService.Lists(context.Background(), &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable:     iotconst.LANG_PRODUCT_THINGS_MODEL,
			SourceRowIdList: sourceIds,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	resList := map[string][]*entitys.LangTranslateEntitys{}
	for _, item := range rep.Data {
		if _, ok := resList[item.SourceRowId]; !ok {
			resList[item.SourceRowId] = []*entitys.LangTranslateEntitys{}
		}
		resList[item.SourceRowId] = append(resList[item.SourceRowId], &entitys.LangTranslateEntitys{
			Lang:       item.Lang,
			FieldName:  item.FieldName,
			FieldType:  item.FieldType,
			FieldValue: item.FieldValue,
		})
	}
	if err != nil {
		return nil, err
	}
	return resList, nil
}

func (s AppDeviceGroupService) ExecuteSwitch(groupId string, userId int64) error {
	devIdList := []string{}
	deviceGroupListResponse, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
		Query: &protosService.IotDeviceGroupList{
			GroupId: iotutil.ToInt64(groupId),
			//UserId:  userId,
		},
	})
	if err != nil {
		return err
	}
	if deviceGroupListResponse.Code != 200 {
		return errors.New(deviceGroupListResponse.Message)
	}

	mapHomeIds := map[int64]struct{}{}

	for _, deviceGroupListInfo := range deviceGroupListResponse.Data {
		devIdList = append(devIdList, deviceGroupListInfo.DevId)
		mapHomeIds[deviceGroupListInfo.HomeId] = struct{}{}
	}

	//todo  从缓存取DeviceInfo信息
	deviceListResponse, err := rpc.IotDeviceInfoService.DeviceInfoListByDevIds(context.Background(), &protosService.DeviceInfoListByDevIdsFilter{DevIds: devIdList})

	if err != nil {
		logger.Errorf("execute error : %s", err.Error())
		return err
	}
	if deviceListResponse.Code != 200 {
		logger.Errorf(deviceListResponse.Message)
		return errors.New(deviceListResponse.Message)
	}

	deviceKeyValue := make(map[string]string)
	for _, deviceInfo := range deviceListResponse.Data {
		deviceKeyValue[deviceInfo.Did] = deviceInfo.ProductKey
	}

	executeResult := true
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_GROUP_DATA+groupId, iotconst.FIELD_PREFIX_DPID+iotutil.ToString(1))
	if strCmd.Val() == "" || strCmd.Val() == "false" || strCmd.Val() == "0" {
		executeResult = true
	} else if strCmd.Val() == "true" || strCmd.Val() == "1" {
		executeResult = false
	}

	appHomeSvc := services2.AppHomeService{}
	for _, deviceId := range devIdList {
		devState := appHomeSvc.GetDeviceStatus(deviceId)
		if devState.HasForceUpgrade {
			logger.Infof("device need upgrade : devId: %s, push error", deviceId)
			continue
		}
		productKey := deviceKeyValue[deviceId]
		_, _, err = PubControl(productKey, deviceId, map[string]interface{}{
			"1": executeResult,
		})
		if err != nil {
			logger.Errorf("rpcClient.PubControl.Publish error : %s", err.Error())
		}
		//todo 这里统计下成功多少次、失败多少次，返回一个错误码给前端去处理
	}

	dpidInfo := make(map[string]string)
	dpidInfo[iotconst.FIELD_PREFIX_DPID+"1"] = iotutil.ToString(executeResult)
	//缓存物模型数据
	rdCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_GROUP_DATA+groupId, dpidInfo)
	if rdCmd.Err() != nil {
		return rdCmd.Err()
	}

	keys := []string{}
	ctx := context.Background()
	for k := range mapHomeIds {
		resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
			Query: &protosService.UcHomeUser{
				HomeId: k,
			},
		})
		if err != nil {
			iotlogger.LogHelper.Helper.Error("get home user list error: ", err)
			continue
		}
		for i := range resp.Data {
			keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(k), iotutil.ToString(resp.Data[i].UserId)))
		}
	}
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}
