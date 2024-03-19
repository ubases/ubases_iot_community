package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	services2 "cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go-micro.dev/v4/metadata"
)

type AppDevService struct {
	Ctx      context.Context
	AppKey   string
	TenantId string
}

func (s AppDevService) SetContext(ctx context.Context) AppDevService {
	s.Ctx = ctx
	return s
}
func (s AppDevService) SetApp(appKey, tenantId string) AppDevService {
	s.AppKey = appKey
	s.TenantId = tenantId
	return s
}

func (s AppDevService) setProduct(rep *protosService.CurrentIotDeviceInfo) {
	productInfo, err := rpc.ProductService.Find(context.Background(), &protosService.OpmProductFilter{
		Id: iotutil.ToInt64(rep.ProductId),
	})
	if err == nil && len(productInfo.Data) > 0 {
		proInfo := productInfo.Data[0]
		rep.Pic = proInfo.ImageUrl
		rep.NetworkMode = iotutil.ToString(proInfo.NetworkType)
		rep.Model = proInfo.ProductKey
		rep.ProductName = proInfo.Name
		rep.ProductNameEn = proInfo.NameEn
		rep.IsShowImg = proInfo.IsShowImg
		rep.PanelProImg = proInfo.PanelProImg
		rep.StyleLinkage = proInfo.StyleLinkage
	}

}

func (s AppDevService) setRooms(rep *protosService.CurrentIotDeviceInfo, lang, tenantId, appKey string) []*protosService.HomeRoomInfo {
	defaultRooms := services2.GetDefaultRooms(lang, tenantId, appKey)

	roomList, err := rpc.UcHomeRoomService.Lists(context.Background(), &protosService.UcHomeRoomListRequest{
		Query: &protosService.UcHomeRoom{
			HomeId: iotutil.ToInt64(rep.HomeId),
		},
	})
	homeRoomInfoList := make([]*protosService.HomeRoomInfo, 0)
	if err == nil && len(roomList.Data) > 0 {
		for _, v := range roomList.Data {
			if v.Id == iotutil.ToInt64(rep.RoomId) {
				rep.RoomName = v.RoomName
			}
			if v.RoomTemplateId != 0 {
				if dfVal, ok := defaultRooms[iotutil.ToString(v.RoomTemplateId)]; ok {
					v.RoomName = dfVal
				}
			}
			homeRoomInfoList = append(homeRoomInfoList, &protosService.HomeRoomInfo{
				RoomId: v.Id,
				Name:   v.RoomName,
				Icon:   v.IconUrl,
				Sort:   v.Sort,
			})
		}
	}
	return homeRoomInfoList
}

// DeviceInfo 设备信息查询
func (s AppDevService) DeviceInfo(devId string, userId int64, devSecret string, language string) (*entitys.CurrentIotDeviceInfo, error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	rep, err := rpc.IotDeviceInfoService.CurrentDeviceInfo(context.Background(), &protosService.CurrentDeviceInfoFilter{
		DevId:     devId,
		UserId:    userId,
		DevSecret: devSecret,
		Lang:      lang,
	})
	if err != nil {
		return nil, err
	}
	if rep.Data == nil {
		return nil, errors.New("未获取到设备信息")
	}

	rep.Data.HomeRoomList = s.setRooms(rep.Data, lang, tenantId, appKey)
	s.setProduct(rep.Data)

	OpmThingModelResult, err := rpc.ClientOpmThingModelService.ProductThingModel(context.Background(), &protosService.OpmThingModelByProductRequest{
		ProductId: rep.Data.ProductId,
		Custom:    -1,
	})

	var deviceInfoProto *protosService.CurrentIotDeviceInfo
	deviceInfoProto = rep.Data
	deviceInfo := entitys.CurrentDeviceInfo_pb2db(deviceInfoProto)
	if language == "en" {
		deviceInfo.ProductName = deviceInfoProto.ProductNameEn
	} else {
		deviceInfo.ProductName = deviceInfoProto.ProductName
	}
	props := map[string]entitys.TslInfo{}

	//物模型
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	productKey := deviceInfo.Model
	for _, v := range OpmThingModelResult.Data.Properties {
		name := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%s_%s_name", lang, productKey, v.Identifier)], v.Name)
		//客户自定义功能描述获取
		if rep.Data.FuncDescMap != nil {
			if v, ok := rep.Data.FuncDescMap[fmt.Sprintf("%s_%s", productKey, v.Dpid)]; ok && v != "" {
				name = v
			}
		}
		dataSpecsList := s.convertJsonByLang(lang, productKey, v.Dpid, v.Identifier, v.DataSpecsList, langMap, rep.Data.FuncDescMap)
		props[v.Identifier] = entitys.TslInfo{
			DpId:          v.Dpid,
			DataType:      v.DataType,
			Name:          name,
			RwFlag:        v.RwFlag,
			DataSpecs:     v.DataSpecs,
			DataSpecsList: dataSpecsList,
			Required:      v.Required,
			Identifier:    v.Identifier,
			DefaultVal:    v.DefaultVal,
		}
	}
	deviceInfo.Props = props
	deviceInfo.DeviceType = 1

	//共享设备
	deviceShareReceiveData, err := rpc.IotDeviceShareReceiveService.Find(context.Background(), &protosService.IotDeviceShareReceiveFilter{
		UserId:   userId,
		DeviceId: devId,
		IsAgree:  2, //已同意
	})
	if err == nil && len(deviceShareReceiveData.Data) > 0 {
		deviceInfo.DeviceType = 2 //共享设备
		deviceInfo.BelongUserName = deviceShareReceiveData.Data[0].BelongUserName
		deviceInfo.ReceiveShareId = iotutil.ToString(deviceShareReceiveData.Data[0].Id)
	}

	return deviceInfo, nil
}

func (s AppDevService) convertJsonByLang(lang string, productKey string, dpid int32, identifier string, jsonStr string, langMap map[string]string, funcSetMap map[string]string) string {
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
				val = "true"
				item["value"] = true
			} else {
				val = "false"
				item["value"] = false
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
	}
	return iotutil.ToString(resObj)
}

func (s AppDevService) CheckDevice(devId, secretKey, ssid string, homeId int64) (bool, error) {
	// 防止已存在的设备重现配网，直接成功问题
	rep, err := rpc.IotDeviceHomeService.Find(context.Background(), &protosService.IotDeviceHomeFilter{
		HomeId:   homeId,
		DeviceId: devId,
		Secrtkey: secretKey,
	})
	if err != nil {
		return false, err
	}
	//如果存在当前设备的绑定记录则为成功
	isok := rep.Total > 0

	//todo 需要保存ssid
	//if isok && origin == 1 {
	//	//origin 表示来源，1：app或者固件
	//	SaveSSID(bson.ObjectIdHex(homeid), devid, ssid, database)
	//}
	return isok, nil
}

func (s AppDevService) RemoveDev(req entitys.RemoveDevFilter, userId int64) error {
	homeId, err := iotutil.ToInt64AndErr(req.HomeId)
	if err != nil {
		return err
	}
	//推送消息需要用到家庭详情接口
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(err)
		return err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(ret.Message)
		return errors.New(ret.Message)
	}
	data := ret.Data

	var currentUserRole int32
	for _, v := range data.UserList {
		if v.Uid == iotutil.ToString(userId) {
			currentUserRole = v.Role
			break
		}
	}
	if currentUserRole == 3 {
		err := errors.New("家庭成员不能删除家庭设备")
		return err
	}

	keys := services2.ReadHomeRoomListsCachedKey(homeId)
	for i := range data.UserList {
		keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(data.UserList[i].Uid)))
		keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, req.HomeId, iotutil.ToString(data.UserList[i].Uid)))
	}
	//移除设备，需要清空共享者的缓存
	if len(req.DevIdList) > 0 {
		sharedUsers, err := rpc.IotDeviceSharedService.Lists(context.Background(), &protosService.IotDeviceSharedListRequest{
			Query: &protosService.IotDeviceShared{DeviceIds: req.DevIdList},
		})
		if err == nil && sharedUsers.Code == 200 {
			for _, d := range sharedUsers.Data {
				u, err := rpc.TUcUserService.FindById(context.Background(), &protosService.UcUserFilter{Id: d.UserId})
				if err == nil && len(u.Data) > 0 {
					keys = append(keys,
						persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(d.UserId)),
						persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, u.Data[0].DefaultHomeId, iotutil.ToString(d.UserId)),
					)
				}
			}
		}
	}

	_, err = rpc.IotDeviceHomeService.RemoveDev(context.Background(), &protosService.RemoveDevRequest{
		DevId:     req.DevId,
		HomeId:    req.HomeId,
		RoomId:    req.RoomId,
		UserId:    iotutil.ToString(userId),
		DevIdList: req.DevIdList,
	})
	if err != nil {
		return err
	}

	ctx := context.Background()
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)

		iotlogger.LogHelper.Info("clear keys ==>", keys)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}

	du := iotstruct.DeviceRedisUpdate{
		UserId:     iotutil.ToString(userId),
		UpdateType: iotstruct.UPDATE_TYPE_REMOVE_DEVICE,
		DeviceIds:  req.DevIdList,
	}
	duBytes, err := json.Marshal(du)
	if err != nil {
		return err
	}
	err = cached.RedisStore.GetClient().Publish(context.Background(), strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, iotutil.ToString(userId)}, "."), string(duBytes)).Err()
	if err != nil {
		return err
	}

	//TODO 这里需要迁移到微服务中（目前已经迁移清除设备定时器、设备分享 （迁移到设备服务中处理））
	go s.clearDevice(req, userId, data)
	return nil
}

func (s AppDevService) clearDevice(req entitys.RemoveDevFilter, userId int64, data *protosService.UcHomeDetail) error {
	defer iotutil.PanicHandler(req, userId)
	if req.DevId != "" {
		req.DevIdList = append(req.DevIdList, req.DevId)
	}

	if req.DevIdList != nil && len(req.DevIdList) > 0 && data != nil {
		deviceGroupListResponse, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
			Query: &protosService.IotDeviceGroupList{
				HomeId: iotutil.ToInt64(req.HomeId),
			},
		})
		if err == nil {
			//用户下无群组不需要调用群组清理和消息推送
			if len(deviceGroupListResponse.Data) > 0 {
				for _, devId := range req.DevIdList {
					//// 清除设备定时器、设备分享 （迁移到设备服务中处理）
					groupMap := map[int64][]*protosService.IotDeviceGroupList{}
					for _, r := range deviceGroupListResponse.Data {
						if _, ok := groupMap[r.GroupId]; ok {
							groupMap[r.GroupId] = append(groupMap[r.GroupId], r)
						} else {
							groupMap[r.GroupId] = []*protosService.IotDeviceGroupList{r}
						}
					}
					groupMapNew := map[int64][]*protosService.IotDeviceGroupList{}
					groupIds := []*protosService.IotDeviceGroupPrimarykey{}
					groupListIds := []*protosService.IotDeviceGroupListPrimarykey{}
					for groupId, groupList := range groupMap {
						hasDevId := false
						for _, item := range groupList {
							if item.DevId == devId {
								hasDevId = true
							}
						}
						if len(groupList) <= 2 && hasDevId {
							groupIds = append(groupIds, &protosService.IotDeviceGroupPrimarykey{Id: groupId})
							groupMapNew[groupId] = groupList
							for _, l := range groupList {
								groupListIds = append(groupListIds, &protosService.IotDeviceGroupListPrimarykey{Id: l.Id})
							}
						}
					}

					if len(groupMapNew) > 0 {
						_, err = rpc.IotDeviceGroupService.DeleteByIds(context.Background(), &protosService.IotDeviceGroupBatchDeleteRequest{
							Keys: groupIds,
						})
						if err != nil {
							return err
						}
						_, err = rpc.IotDeviceGroupListService.DeleteByIds(context.Background(), &protosService.IotDeviceGroupListBatchDeleteRequest{
							Keys: groupListIds,
						})
						if err != nil {
							return err
						}
						for _, groupList := range groupMapNew {
							services2.SendAutoDisbandGroupMessage(services2.SetAppInfoByContext(s.Ctx), data, userId, iotutil.ToInt64(req.HomeId), groupList[1].GroupName)
						}
					}
				}
			}
		} else {
			iotlogger.LogHelper.WithTag("method", "clearDevice").Error(err.Error())
		}
		services2.SendRemoveDeviceMessage(services2.SetAppInfoByContext(s.Ctx), data, userId, iotutil.ToInt64(req.HomeId), req.DevIdList...)
	}
	return nil
}

func (s AppDevService) UpdateDev(req entitys.UpdateDevFilter, userId int64) error {
	roomId, _ := iotutil.ToInt64AndErr(req.RoomId)
	_, err := rpc.IotDeviceHomeService.UpdateDeviceInfo(context.Background(), &protosService.IotDeviceHome{
		HomeId:     iotutil.ToInt64(req.HomeId),
		DeviceId:   req.DevId,
		RoomId:     roomId,
		CustomName: req.DevName,
	})
	if err != nil {
		return err
	}
	//修改设备名称
	if req.DevName != "" {
		iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+req.DevId, map[string]interface{}{
			"deviceName": req.DevName,
		})
	}
	// 删除家庭详情缓存
	keys := []string{}
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: iotutil.ToInt64(req.HomeId),
		},
	})
	if err != nil {
		return err
	}
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, req.HomeId, iotutil.ToString(resp.Data[i].UserId)))
	}
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}
	du := iotstruct.DeviceRedisUpdate{
		UserId:    iotutil.ToString(userId),
		HomeId:    req.HomeId,
		DeviceIds: []string{req.DevId},
	}
	duBytes, err := json.Marshal(du)
	if err != nil {
		return err
	}
	err = cached.RedisStore.GetClient().Publish(context.Background(), strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, iotutil.ToString(userId)}, "."), string(duBytes)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s AppDevService) AddDev(req entitys.AddDevFilter) error {
	_, err := rpc.IotDeviceHomeService.UpdateDeviceInfo(context.Background(), &protosService.IotDeviceHome{
		HomeId:   iotutil.ToInt64(req.HomeId),
		DeviceId: req.DevId,
		RoomId:   iotutil.ToInt64(req.RoomId),
		Sort:     req.Sort,
	})
	if err != nil {
		return err
	}
	return nil
}
