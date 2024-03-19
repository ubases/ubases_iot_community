package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/controls"
	entitys2 "cloud_platform/iot_app_api_service/controls/common/entitys"
	_const "cloud_platform/iot_app_api_service/controls/user/const"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/rpc"
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
	"sort"
	"strings"
	"sync"
	"time"

	"go-micro.dev/v4/metadata"

	"github.com/gin-gonic/gin"

	"go-micro.dev/v4/logger"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppHomeService struct {
	Ctx      context.Context
	AppKey   string
	TenantId string
	Lang     string
}

func (s AppHomeService) SetContext(ctx context.Context) AppHomeService {
	s.Ctx = ctx
	return s
}
func (s AppHomeService) SetApp(appKey, tenantId string) AppHomeService {
	s.AppKey = appKey
	s.TenantId = tenantId
	return s
}

func (s AppHomeService) AddHome(req entitys.UcHomeEntitys, userId int64) error {
	saveObj := entitys.UcHomeReq_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Sid = 0 //todo 待处理
	saveObj.UserId = userId
	saveObj.CreatedBy = userId
	_, err := rpc.UcHomeService.AddHome(s.Ctx, saveObj)
	if err != nil {
		return err
	}
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(userId))); err != nil {
		return err
	}
	return nil
}

func (s AppHomeService) Details(c *gin.Context, homeId int64, userId int64) (*entitys.UcHomeDetailEntitys, int, error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	homeInfo := &entitys.UcHomeDetailEntitys{}
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(userId)), homeInfo)
	if err == nil {
		//切换语言之后，获取房间翻译数据
		defaultRooms := GetDefaultRooms(lang, tenantId, appKey)
		//设置家庭设备的升级
		s.setHomeDeviceList(homeInfo, defaultRooms)
		//设置共享设备
		//s.setSharedDeviceList(userId, homeInfo)
		//设置群组信息
		// s.setGroupDeviceList(homeId, homeInfo)
		homeInfo.Name = HomeLanguage(lang, homeInfo.Name)
		//设置产品的面板更新信息
		s.setProductPanel(homeInfo)
		for i, room := range homeInfo.RoomList {
			homeInfo.RoomList[i].Sort = room.Sort
			roomName := room.Name
			if room.TemplateId != 0 {
				if dfVal, ok := defaultRooms[iotutil.ToString(room.TemplateId)]; ok {
					roomName = dfVal
				}
			}
			homeInfo.RoomList[i].Name = roomName
		}
		return homeInfo, 0, nil
	}
	//查询完整的家庭详细信息
	res, err := rpc.UcHomeService.HomeDetail(s.Ctx, &protosService.UcHomeDetailRequest{
		HomeId: homeId,
		UserId: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("调用UcHomeService.HomeDetail异常，%s", err.Error())
		return nil, 0, err
	}
	if res.Code != 200 {
		iotlogger.LogHelper.Errorf("调用UcHomeService.HomeDetail异常，- %s", res.Message)
		return nil, 0, errors.New(res.Message)
	}
	if res.Data == nil || res.Data.Data == nil {
		return nil, ioterrs.ERROR_HOME_ID.Code, errors.New(ioterrs.ERROR_HOME_ID.Msg)
	}
	var homeDetail = res.Data
	homeInfo = entitys.UcHome_pb2db(homeDetail.Data)
	// 查询国家名称，省份名称，城市名称
	if len(homeInfo.CountryId) != 0 {
		id, err := iotutil.ToInt32Err(homeInfo.CountryId)
		if err != nil {
			return homeInfo, 0, err
		}
		areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
			Id: int64(id),
		})
		if err != nil {
			return homeInfo, 0, err
		}
		if res.Code != 200 {
			return homeInfo, 0, errors.New(areaData.Message)
		}
		homeInfo.CountryName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
	}
	if len(homeInfo.ProvinceId) != 0 {
		id, err := iotutil.ToInt32Err(homeInfo.ProvinceId)
		if err != nil {
			return homeInfo, 0, err
		}
		areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
			Id: int64(id),
		})
		if err != nil {
			return homeInfo, 0, err
		}
		if res.Code != 200 {
			return homeInfo, 0, errors.New(areaData.Message)
		}
		homeInfo.ProvinceName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
	}
	if len(homeInfo.CityId) != 0 {
		id, err := iotutil.ToInt32Err(homeInfo.CityId)
		if err != nil {
			return homeInfo, 0, err
		}
		areaData, err := rpc.ClientAreaService.FindById(s.Ctx, &protosService.SysAreaFilter{
			Id: int64(id),
		})
		if err != nil {
			return homeInfo, 0, err
		}
		if res.Code != 200 {
			return homeInfo, 0, errors.New(areaData.Message)
		}
		homeInfo.CityName = getNameByLang(lang, areaData.Data[0].ChineseName, areaData.Data[0].EnglishName)
	}
	//如果省份为空，则将选择省份的名称赋值给Province
	if homeInfo.Province == "" {
		homeInfo.Province = homeInfo.ProvinceName
	}
	//房间名称翻译
	homeInfo.Name = HomeLanguage(lang, homeInfo.Name)
	//获取家庭所在区域
	mqttInfo, _ := GetRegionMqttById(iotutil.ToString(homeDetail.Data.Sid))

	for _, v := range homeDetail.UserList {
		if v.Uid == iotutil.ToString(userId) {
			homeInfo.CurrentUserRole = v.Role
		}
		homeInfo.UserList = append(homeInfo.UserList, entitys.HomeUser{
			UserId:   v.Uid,
			Role:     v.Role,
			Photo:    v.Photo,
			NickName: v.NickName,
		})
	}
	if homeInfo.CurrentUserRole == 0 {
		err = errors.New(ioterrs.ERROR_HOME_ID.Msg)
		return nil, ioterrs.ERROR_HOME_ID.Code, err
	}
	//tenantId, lang, appKey string
	defaultRooms := GetDefaultRooms(lang, tenantId, appKey)
	roomDeviceCount := map[string]int32{}
	var deviceMap sync.Map
	for _, v := range homeDetail.DeviceList {
		devInfo := v.Data
		devStatus := s.GetDeviceStatus(devInfo.Did)
		roomName := devInfo.RoomName
		if devInfo.RoomTemplateId != 0 {
			if dfVal, ok := defaultRooms[iotutil.ToString(devInfo.RoomTemplateId)]; ok {
				roomName = dfVal
			}
		}
		devInfo.RoomName = roomName
		homeInfo.DeviceList = append(homeInfo.DeviceList, *entitys.HomeDevice_2e(homeInfo.Name, devInfo, devStatus, mqttInfo))

		deviceCount, _ := roomDeviceCount[devInfo.RoomId]
		deviceCount = deviceCount + 1
		roomDeviceCount[devInfo.RoomId] = deviceCount
		deviceMap.Store(devInfo.Did, true)
	}
	//家庭房间信息
	homeInfo.SetRoom(defaultRooms, homeDetail.RoomList, roomDeviceCount)
	//设置共享设备
	s.setSharedDeviceList(userId, homeInfo, deviceMap)
	//设置群组信息
	s.setGroupDeviceList(homeId, homeInfo)
	//设置产品的面板更新信息
	s.setProductPanel(homeInfo)

	//根据sort进行排序
	sort.Slice(homeInfo.DeviceList, func(i, j int) bool {
		return homeInfo.DeviceList[i].Time > homeInfo.DeviceList[j].Time
	})
	//设置缓存
	if err := cached.RedisStore.Set(persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(userId)), homeInfo, 600*time.Second); err != nil {
		return homeInfo, 0, err
	}
	return homeInfo, 0, err
}

// 获取房间，并读取翻译
func GetDefaultRooms(lang, tenantId, appKey string) map[string]string {
	res := map[string]string{}
	key := fmt.Sprintf("%s_%s_%s_%s", tenantId, lang, iotconst.LANG_OEM_APP_ROOMS, appKey)
	var roomList []entitys2.TConfigDictData
	rooms := iotredis.GetClient().Get(context.Background(), key)
	if rooms.Err() == nil && rooms.Val() != "" {
		err := json.Unmarshal([]byte(rooms.Val()), &res)
		if err == nil {
			for _, r := range roomList {
				res[iotutil.ToString(r.DictCode)] = r.DictLabel
			}
			return res
		}
	}
	//获取默认房间信息
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		return res
	}
	if oemAppResult.Code != 200 {
		return res
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppUiConfigService.Find(context.Background(), &protosService.OemAppUiConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	roomList = []entitys2.TConfigDictData{}
	if err == nil && oemAppResult.Code == 200 && len(oemAppFunctionConfig.Data) > 0 && oemAppFunctionConfig.Data[0].Room != "" {
		roomConfigs := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Room)
		for _, roomConfig := range roomConfigs {
			roomList = append(roomList, entitys2.TConfigDictData{
				DictCode:  iotutil.ToInt64(roomConfig["roomId"]),
				DictLabel: iotutil.ToString(roomConfig["roomName"]),
				DictValue: iotutil.ToString(roomConfig["roomSort"]),
				DictType:  "",
				Listimg:   iotutil.ToString(roomConfig["roomImage"]),
			})
		}
	}
	langMap := make(map[string]string)
	if lang != "" {
		langKey := fmt.Sprintf("%s_%s%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX, iotconst.LANG_OEM_APP_ROOMS, appKey)
		langMap, err = iotredis.GetClient().HGetAll(context.Background(), langKey).Result()
		if err != nil {
			langMap = make(map[string]string)
		}
	}
	for i, dict := range roomList {
		dictLabel := iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%v_name", lang, dict.DictCode)], dict.DictLabel)
		roomList[i].DictLabel = dictLabel
		res[iotutil.ToString(dict.DictCode)] = dictLabel
	}
	iotredis.GetClient().Set(context.Background(), key, iotutil.ToString(roomList), 0)
	return res
}

// 设置家庭设备列表（赋值升级状态）
func (s AppHomeService) setHomeDeviceList(homeInfo *entitys.UcHomeDetailEntitys, defaultRooms map[string]string) {
	//区域兼容处理regionId := controls.GetRegion(c)
	mqttInfo, _ := GetRegionMqttByHomeId(iotutil.ToInt64(homeInfo.Id))
	for i, v := range homeInfo.DeviceList {
		if v.DevType != 3 {
			devStatus := s.GetDeviceStatus(v.Did)
			//设备在线状态
			homeInfo.DeviceList[i].OnlineState = devStatus.OnlineStatus
			//设备开关状态
			homeInfo.DeviceList[i].DevSwitch = devStatus.PowerState
			//是否需要升级
			homeInfo.DeviceList[i].HasForceUpgrade = devStatus.UpgradeMode == 2 && devStatus.OtaUpgradeStatus == 1
			homeInfo.DeviceList[i].Version = devStatus.Version
			//兼容，保留原来MqttServer的逻辑，如果区域有mqtt地址，则采用区域服务地址
			if mqttInfo != nil && mqttInfo.WebsocketServer != "" {
				homeInfo.DeviceList[i].MqttServer = mqttInfo.WebsocketServer
			}
			//是否升级中
			var upgradeStatus int32 = 0
			if devStatus.OtaUpgradeStatus == 1 {
				if devStatus.UpgradeState == "Downloading" || devStatus.UpgradeState == "Installing" {
					upgradeStatus = 1
				}
			}
			homeInfo.DeviceList[i].OtaUpgradeStatus = upgradeStatus
			//图片OSS样式转换
			homeInfo.DeviceList[i].ProductPic = controls.ConvertProImg(homeInfo.DeviceList[i].ProductPic)
			//homeInfo.DeviceList[i].OtaUpgradeStatus = devStatus.OtaUpgradeStatus
			//是否强制升级
			//homeInfo.DeviceList[i].UpgradeMode = devStatus.UpgradeMode
			//房间名称翻译
			roomName := v.RoomName
			if v.TemplateId != 0 {
				if dfVal, ok := defaultRooms[iotutil.ToString(v.TemplateId)]; ok {
					roomName = dfVal
				}
			}
			homeInfo.DeviceList[i].RoomName = roomName
		}
	}
	return
}

// 设置共享设备
func (s AppHomeService) setSharedDeviceList(userId int64, homeInfo *entitys.UcHomeDetailEntitys, deviceMap sync.Map) {
	DeviceSharedResult, err := rpc.IotDeviceSharedService.Lists(context.Background(), &protosService.IotDeviceSharedListRequest{
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.IotDeviceShared{
			UserId: userId,
		},
	})
	if err == nil && DeviceSharedResult.Code == 200 {
		mqtts := map[int64]*protosService.SysRegionServer{}
		for _, v := range DeviceSharedResult.Data {
			//排除你拥有的设备，确保列表中不会出现重复的设备
			if _, ok := deviceMap.Load(v.DeviceId); ok {
				continue
			}
			//获取区域服务器地址
			var mqttInfo *protosService.SysRegionServer
			if m, ok := mqtts[v.HomeId]; ok {
				mqttInfo = m
			} else {
				mqttInfo, _ = GetRegionMqttByHomeId(v.HomeId)
				mqtts[v.HomeId] = mqttInfo
			}
			//获取设备的缓存状态
			devStatus := s.GetDeviceStatus(v.DeviceId)
			//转换升级状态
			var upgradeStatus int32 = 0
			if devStatus.OtaUpgradeStatus == 1 {
				if devStatus.UpgradeState == "Downloading" || devStatus.UpgradeState == "Installing" {
					upgradeStatus = 1
				}
			}
			//mqtt地址兼容处理
			mqttServer := v.MqttServer
			if mqttInfo != nil && mqttInfo.WebsocketServer != "" {
				mqttServer = mqttInfo.WebsocketServer
			}
			homeInfo.DeviceList = append(homeInfo.DeviceList, entitys.HomeDevice{
				Did:        v.DeviceId,
				ProductId:  v.ProductId,
				DeviceName: v.CustomName,
				ProductPic: controls.ConvertProImg(v.ProductPic),
				RoomName:   "",
				RoomId:     "",
				HomeName:   homeInfo.Name,
				ProductKey: v.ProductKey,
				MqttServer: mqttServer,
				//设备在线状态
				OnlineState: devStatus.OnlineStatus,
				//设备开关状态
				DevSwitch:       devStatus.PowerState,
				HasForceUpgrade: devStatus.UpgradeMode == 2 && devStatus.OtaUpgradeStatus == 1,
				Version:         devStatus.Version,
				//是否需要升级
				OtaUpgradeStatus: upgradeStatus,
				//是否强制升级
				//UpgradeMode: devStatus.UpgradeMode,
				DevType: 2, //共享设备
				Time:    v.SharedTime.AsTime().Unix(),
			})
		}
	}
}

// 设置产品的面板更新信息
func (s AppHomeService) setProductPanel(homeInfo *entitys.UcHomeDetailEntitys) {
	productIds := make([]int64, 0)
	for _, v := range homeInfo.DeviceList {
		if v.DevType != 3 {
			productIds = append(productIds, v.ProductId)
		}
	}
	if len(productIds) == 0 {
		return
	}
	panels, err := s.getProductsList(productIds)
	if err != nil {
		iotlogger.LogHelper.Infof("Home Details setProductPanel异常，%s", err.Error())
		return
	}
	homeInfo.PanelList = make([]entitys.ProductPanel, 0)
	for _, product := range panels {
		homeInfo.PanelList = append(homeInfo.PanelList, entitys.ProductPanel{
			Url:        product.PanelUrl,
			Key:        product.PanelKey,
			ProductKey: product.ProductKey,
		})
	}
}

// 设置产品的面板更新信息 - 通过Id获取产品信息
func (s AppHomeService) getProductsList(productIds []int64) ([]*protosService.OpmProduct, error) {
	if len(productIds) == 0 {
		return []*protosService.OpmProduct{}, nil
	}
	productsRes, err := rpc.ProductService.PanelListsByProductIds(context.Background(), &protosService.ListsByProductIdsRequest{
		ProductIds: productIds,
	})
	if err != nil {
		return nil, err
	}
	return productsRes.Data, nil
}

// 设置群组信息
func (s AppHomeService) setGroupDeviceList(homeId int64, homeInfo *entitys.UcHomeDetailEntitys) {
	deviceGroupResult, err := rpc.IotDeviceGroupService.Lists(context.Background(), &protosService.IotDeviceGroupListRequest{
		OrderKey:  "created_at",
		OrderDesc: "desc",
		Query: &protosService.IotDeviceGroup{
			HomeId: homeId,
		},
	})
	if err == nil && deviceGroupResult.Code == 200 {
		for _, v := range deviceGroupResult.Data {
			var state, devSwitch int32
			deviceGroupResponse, err := rpc.IotDeviceGroupService.FindById(context.Background(), &protosService.IotDeviceGroupFilter{
				Id: v.Id,
			})
			if err != nil {
				continue
			}
			if deviceGroupResponse.Code != 200 {
				continue
			}

			deviceGroupData := deviceGroupResponse.Data[0]
			homeId = deviceGroupData.HomeId

			deviceGroupListResponse, err := rpc.IotDeviceGroupListService.Lists(context.Background(), &protosService.IotDeviceGroupListListRequest{
				Query: &protosService.IotDeviceGroupList{
					GroupId: v.Id,
					HomeId:  homeId,
					//UserId:  userId,
				},
			})
			if err != nil {
				continue
			}
			if deviceGroupListResponse.Code != 200 || len(deviceGroupListResponse.Data) == 0 {
				continue
			}
			deviceGroupList := deviceGroupListResponse.Data[0]

			deviceHomeResp, err := rpc.IotDeviceHomeService.Find(context.Background(), &protosService.IotDeviceHomeFilter{
				DeviceId: deviceGroupList.DevId,
			})
			if err != nil {
				iotlogger.LogHelper.Infof("调用IotDeviceHomeService.Find异常，%s", err.Error())
				continue
			}
			if deviceHomeResp.Code != 200 {
				iotlogger.LogHelper.Infof("调用IotDeviceHomeService.Find异常，- %s", deviceHomeResp.Message)
				continue
			}

			productInfo, err := rpc.ProductService.FindById(context.Background(), &protosService.OpmProductFilter{
				Id: deviceHomeResp.Data[0].ProductId,
			})
			if err != nil {
				iotlogger.LogHelper.Infof("调用ProductService.FindById异常，%s", err.Error())
				continue
			}
			if productInfo.Code != 200 {
				iotlogger.LogHelper.Infof("调用ProductService.FindById异常，- %s", productInfo.Message)
				continue
			}

			strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_GROUP_DATA+iotutil.ToString(v.Id), iotconst.FIELD_PREFIX_DPID+iotutil.ToString(1))
			if strCmd.Val() == "true" || strCmd.Val() == "1" {
				devSwitch = 1
			} else {
				devSwitch = 0
			}
			homeInfo.DeviceList = append(homeInfo.DeviceList, entitys.HomeDevice{
				GroupId:     iotutil.ToString(v.Id),
				DeviceName:  v.Name,
				ProductPic:  productInfo.Data[0].ImageUrl,
				RoomName:    v.RoomName,
				RoomId:      iotutil.ToString(v.RoomId),
				HomeName:    "",
				ProductKey:  "",
				SecretKey:   "",
				MqttServer:  v.MqttServer,
				OnlineState: state,
				DevSwitch:   devSwitch,
				DevType:     3, //群组设备
				DevCount:    iotutil.ToInt32(len(deviceGroupListResponse.Data)),
			})
		}
	}
}

func (s AppHomeService) UpdateHome(userId, homeId string, req entitys.UpdateHomeEntitys) error {
	updateObj := entitys.UpdateHome_e2pb(&req)
	updateObj.UpdatedAt = timestamppb.Now()
	updateObj.CreatedBy = iotutil.ToInt64(userId)
	updateObj.Id = iotutil.ToInt64(homeId)
	_, err := rpc.UcHomeService.Update(context.Background(), updateObj)
	if err != nil {
		return err
	}
	// 删除家庭详情缓存
	s.clearHomeDetailsCached(updateObj.Id, nil)
	return nil
}

// 清理家庭缓存
func (s AppHomeService) clearHomeDetailsCached(homeId int64, keys []string) error {
	defer iotutil.PanicHandler("清理家庭详情缓存", homeId)
	if homeId == 0 {
		return errors.New("家庭Id不能为空")
	}
	// 删除家庭详情缓存
	if keys == nil {
		keys = make([]string, 0)
	}
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: homeId,
		},
	})
	if err != nil {
		return err
	}
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(resp.Data[i].UserId)),
			fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, homeId, iotutil.ToString(resp.Data[i].UserId)))
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

func (s AppHomeService) SendInvitationCode(homeId string, userId int64, appKey, tenantId string) (string, string) {
	//todo 还需要判断当前家庭是否存在

	code := ""
	//防止邀请码重复
	var i int = 5
	for {
		if i <= 0 {
			break
		}
		tempCode := strings.ToUpper(iotutil.GetRandomStringCombination(6))
		resp := iotredis.GetClient().Get(context.Background(), iotconst.APP_INVITE_CODE+tempCode)
		if resp.Val() == "" {
			code = tempCode
			break
		}
		continue
	}
	if code == "" {
		return "", "邀请码生成失败"
	}
	joinHome := map[string]interface{}{
		"userId":   userId,
		"homeId":   homeId,
		"appKey":   appKey,
		"tenantId": tenantId,
	}
	//设置过期,有效期3天，60*60*24*3
	res := iotredis.GetClient().Set(context.Background(), iotconst.APP_INVITE_CODE+code, iotutil.ToString(joinHome), 259200*time.Second) //有效期3天
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("缓存邀请码失败:%s", res.Err().Error())
		return "", "缓存邀请码失败"
	}
	return code, ""
}

// 加入家庭
func (s AppHomeService) JoinHome(smsCode, userId, appKey, tenantId string) (code int32, err error) {
	resp := iotredis.GetClient().Get(context.Background(), iotconst.APP_INVITE_CODE+smsCode)
	if resp.Val() == "" {
		iotlogger.LogHelper.Errorf("邀请码无效，请联系邀请者重新生成")
		return int32(ioterrs.ERROR_INVALIDINVITATIONCODE.Code), errors.New(ioterrs.ERROR_INVALIDINVITATIONCODE.Msg)
	}

	joinInfo, err := iotutil.JsonToMapErr(resp.Val())
	if err != nil {
		iotlogger.LogHelper.Errorf("邀请码无效，请联系邀请者重新生成")
		return int32(ioterrs.ERROR_INVALIDINVITATIONCODE.Code), errors.New(ioterrs.ERROR_INVALIDINVITATIONCODE.Msg)
	}
	homeIdInt := iotutil.ToInt64(joinInfo["homeId"])
	fromUserId := iotutil.ToInt64(joinInfo["userId"])
	homeAppKey := iotutil.ToString(joinInfo["appKey"])
	homeTenantId := iotutil.ToString(joinInfo["tenantId"])
	toUserId := iotutil.ToInt64(userId)
	if appKey != homeAppKey || tenantId != homeTenantId {
		iotlogger.LogHelper.Errorf("用户和家庭属于不同开发者")
		return int32(ioterrs.ERROR_INVALID_INVITATION_CODE.Code), errors.New(ioterrs.ERROR_INVALID_INVITATION_CODE.Msg)
	}

	joinHomeResp, err := rpc.UcHomeService.JoinHome(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeIdInt,
		UserId: toUserId,
	})
	if joinHomeResp.Code != 200 {
		iotlogger.LogHelper.Error("加入家庭失败，原因:%s", joinHomeResp.Message)
		return joinHomeResp.Code, errors.New(joinHomeResp.Message)
	}

	//删除redis中验证码
	iotredis.GetClient().Del(context.Background(), iotconst.APP_INVITE_CODE+smsCode)

	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeIdInt,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "JoinHome").Error(err)
		return joinHomeResp.Code, err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "JoinHome").Error(ret.Message)
		return joinHomeResp.Code, errors.New("get home detail error")
	}

	keys := []string{}
	for i := range ret.Data.UserList {
		keys = append(keys,
			persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(ret.Data.UserList[i].Uid)),
			persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeIdInt), iotutil.ToString(ret.Data.UserList[i].Uid)),
		)
	}
	ctx := context.Background()
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return joinHomeResp.Code, err
		}
	}
	go SendJoinHomeMessage(SetAppInfoByContext(s.Ctx), ret, homeIdInt, fromUserId, toUserId)
	return 0, nil
}

func (s AppHomeService) SetRole(userId, homeId, thirdUserId int64, roleType int32) string {
	_, err := rpc.UcHomeService.SetRole(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId:      homeId,
		UserId:      userId,
		RoleType:    roleType,
		ThirdUserId: thirdUserId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("成员角色设置失败，原因:%s", err.Error())
		return err.Error()
	}
	// 删除家庭详情缓存
	keys := []string{
		persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(userId)),
		persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(thirdUserId)),
		persist.GetRedisKey(iotconst.APP_HOME_ROOM_LIST_DATA, iotutil.ToString(homeId)),
	}
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: homeId,
		},
	})
	if err != nil {
		return err.Error()
	}
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(resp.Data[i].UserId)))
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err.Error()
		}
	}
	return ""
}

func (s AppHomeService) RemoveMembers(userId, homeId, thirdUserId int64, ip string) string {
	//推送消息需要用到家庭详情接口
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(err)
		return err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(ret.Message)
		return ret.Message
	}

	_, err = rpc.UcHomeService.RemoveMembers(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId:      homeId,
		UserId:      userId,
		ThirdUserId: thirdUserId,
		Ip:          ip,
	})
	if err != nil {
		iotlogger.LogHelper.Error("移除成员失败，原因:%s", err.Error())
		return err.Error()
	}

	keys := []string{}
	for i := range ret.Data.UserList {
		keys = append(keys,
			persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(ret.Data.UserList[i].Uid)),
			persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(ret.Data.UserList[i].Uid)),
		)
	}
	ctx := context.Background()
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err.Error()
		}
	}

	go SendRemoveMembersMessage(SetAppInfoByContext(s.Ctx), ret, homeId, thirdUserId)
	return ""
}

func (s AppHomeService) TransferOwnership(userId, homeId, thirdUserId int64) string {
	_, err := rpc.UcHomeService.TransferOwnership(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId:      homeId,
		UserId:      userId,
		ThirdUserId: thirdUserId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("转移家庭所有权失败，原因:%s", err.Error())
		return err.Error()
	}
	// 删除家庭详情缓存
	keys := []string{}
	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			HomeId: homeId,
		},
	})
	if err != nil {
		return err.Error()
	}
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(resp.Data[i].UserId)))
	}

	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err.Error()
		}
	}
	return ""
}

func (s AppHomeService) Quit(userId, homeId int64) string {
	//推送消息需要用到家庭详情接口
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(err)
		return err.Error()
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendQuitHomeMessage").Error(ret.Message)
		return ret.Message
	}

	_, err = rpc.UcHomeService.Quit(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
		UserId: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("离开家庭失败，原因:%s", err.Error())
		return err.Error()
	}

	keys := []string{}
	for i := range ret.Data.UserList {
		keys = append(keys,
			persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(ret.Data.UserList[i].Uid)),
			persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(ret.Data.UserList[i].Uid)),
		)
	}
	ctx := context.Background()
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err.Error()
		}
	}

	go SendQuitHomeMessage(SetAppInfoByContext(s.Ctx), ret, homeId, userId)
	return ""
}

func (s AppHomeService) RoomList(homeId int64) ([]entitys.UcHomeRoomList, string) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	homeRoomList := make([]entitys.UcHomeRoomList, 0)
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.APP_HOME_ROOM_LIST_DATA, lang, iotutil.ToString(homeId)), &homeRoomList)
	if err == nil {
		return homeRoomList, ""
	}
	UcHomeRoomResponse, err := rpc.UcHomeRoomService.Lists(context.Background(), &protosService.UcHomeRoomListRequest{
		Query: &protosService.UcHomeRoom{
			HomeId: homeId,
		},
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取房间列表失败，原因:%s", err.Error())
		return homeRoomList, err.Error()
	}
	roomList := UcHomeRoomResponse.Data
	devCountList, err := rpc.IotDeviceHomeService.DevCount(context.Background(), &protosService.IotDeviceHomeDevCount{HomeId: homeId})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间设备数量失败，原因:%s", err.Error())
		return homeRoomList, err.Error()
	}
	defaultRooms := GetDefaultRooms(lang, tenantId, appKey)
	for _, room := range roomList {
		roomObj := entitys.UcHomeRoomList{}
		roomObj.RoomId = iotutil.ToString(room.Id)
		roomObj.Icon = room.IconUrl
		roomObj.Sort = room.Sort
		roomName := room.RoomName
		if room.RoomTemplateId != 0 {
			if dfVal, ok := defaultRooms[iotutil.ToString(room.RoomTemplateId)]; ok {
				roomName = dfVal
			}
		}
		roomObj.Name = roomName
		for _, devCount := range devCountList.Keys {
			if devCount.RoomId == room.Id {
				roomObj.DeviceCount = devCount.DevCount
			}
		}
		homeRoomList = append(homeRoomList, roomObj)
	}

	//根据sort进行排序
	sort.Slice(homeRoomList, func(i, j int) bool {
		return homeRoomList[i].Sort < homeRoomList[j].Sort
	})
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.APP_HOME_ROOM_LIST_DATA, lang, iotutil.ToString(homeId)), &homeRoomList, 600*time.Second)
	if err != nil {
		return homeRoomList, err.Error()
	}
	return homeRoomList, ""
}

func (s AppHomeService) DeviceList(homeId int64) ([]entitys.HomeDevice, string) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	deviceHomeList, err := rpc.IotDeviceHomeService.UserDevList(s.Ctx, &protosService.IotDeviceHomeHomeId{HomeId: homeId})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭设备列表失败，原因:%s", err.Error())
		return nil, err.Error()
	}
	if deviceHomeList.Code != 200 {
		return nil, deviceHomeList.Message
	}
	roomMap, err := s.GetRoomMap(homeId)
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间失败，原因:%s", err.Error())
		return nil, err.Error()
	}
	deviceList := []entitys.HomeDevice{}
	//从缓存获取产品信息
	productCached := controls.ProductCachedData{}
	//默认房间翻译
	defaultRooms := GetDefaultRooms(lang, tenantId, appKey)
	for _, v := range deviceHomeList.DevList {
		devInfo := entitys.IotDeviceInfo_pb2db(v)
		productInfo, err := productCached.GetProduct(iotutil.ToString(devInfo.ProductKey))
		if err == nil && productInfo != nil {
			devInfo.ProductKey = productInfo.ProductKey
			devInfo.ProductPic = productInfo.ImageUrl
		}
		roomId, _ := iotutil.ToInt64AndErr(devInfo.RoomId)
		roomInfo, ok := roomMap[roomId]
		var roomName string = v.RoomName
		var roomTemplateId int64 = 0
		if ok {
			roomName = roomInfo.RoomName
			roomTemplateId = roomInfo.RoomTemplateId
		}
		if roomTemplateId != 0 {
			if dfVal, ok := defaultRooms[iotutil.ToString(roomTemplateId)]; ok {
				roomName = dfVal
			} else {
				iotlogger.LogHelper.Infof("通过房间模板Id获取翻译失败，房间模板%v, defaultRooms: %v", devInfo.RoomTemplateId, iotutil.ToString(defaultRooms))
			}
		}
		deviceList = append(deviceList, entitys.HomeDevice{
			Did:         devInfo.Did,
			ProductId:   devInfo.ProductId,
			DeviceName:  devInfo.DeviceName,
			ProductPic:  controls.ConvertProImg(devInfo.ProductPic),
			RoomName:    roomName,
			RoomId:      devInfo.RoomId,
			ProductKey:  devInfo.ProductKey,
			SecretKey:   devInfo.SecretKey,
			MqttServer:  devInfo.MqttServer,
			OnlineState: iotutil.ToInt32(devInfo.OnlineStatus),
			DevSwitch:   0,
			DevType:     1,
			Time:        roomId,
		})
	}
	//根据sort进行排序
	sort.Slice(deviceList, func(i, j int) bool {
		return deviceList[i].Time > deviceList[j].Time
	})
	return deviceList, ""
}

// 获取房间列表，名称为值
func (s AppHomeService) GetRoomList(homeIds ...int64) (roomMap map[int64]string, err error) {
	query := &protosService.UcHomeRoom{}
	if len(homeIds) == 1 {
		query.HomeId = homeIds[0]
	} else {
		query.HomeIds = homeIds
	}
	roomRes, err := rpc.UcHomeRoomService.Lists(s.Ctx, &protosService.UcHomeRoomListRequest{
		Query: query,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间失败，原因:%s", err.Error())
		return nil, err
	}
	if roomRes.Code != 200 {
		return nil, errors.New(roomRes.Message)
	}
	roomMap = make(map[int64]string)
	for _, datum := range roomRes.Data {
		roomMap[datum.Id] = datum.RoomName
	}
	return
}

// 获取房间列表，房间信息为值
func (s AppHomeService) GetRoomMap(homeId int64) (roomMap map[int64]*protosService.UcHomeRoom, err error) {
	roomRes, err := rpc.UcHomeRoomService.Lists(s.Ctx, &protosService.UcHomeRoomListRequest{
		Query: &protosService.UcHomeRoom{HomeId: homeId},
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间失败，原因:%s", err.Error())
		return nil, err
	}
	if roomRes.Code != 200 {
		return nil, errors.New(roomRes.Message)
	}
	roomMap = make(map[int64]*protosService.UcHomeRoom)
	for _, r := range roomRes.Data {
		roomMap[r.Id] = r
	}
	return
}

func (s AppHomeService) UserDeviceList(userId int64) ([]entitys.HomeDevice, string) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
		appKey, _   = metadata.Get(s.Ctx, "appKey")
	)
	homeListInfo, err := rpc.TUcUserService.HomeList(context.Background(), &protosService.UcUser{
		Id: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取用户家庭列表失败，原因:%s", err.Error())
		return nil, err.Error()
	}
	homeMap := make(map[int64]string)
	homeIds := make([]int64, 0)
	for _, v := range homeListInfo.HomeUsers {
		homeMap[v.Id] = v.Name
		homeIds = append(homeIds, iotutil.ToInt64(v.Id))
	}

	deviceHomeList, err := rpc.IotDeviceHomeService.UserDevList(context.Background(), &protosService.IotDeviceHomeHomeId{HomeIds: homeIds})
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭设备列表失败，原因:%s", err.Error())
		return nil, err.Error()
	}

	roomMap, err := s.GetRoomList(homeIds...)
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭房间失败，原因:%s", err.Error())
		return nil, err.Error()
	}

	deviceList := []entitys.HomeDevice{}
	productCached := controls.ProductCachedData{}
	//默认房间翻译
	defaultRooms := GetDefaultRooms(lang, tenantId, appKey)
	for _, devInfo := range deviceHomeList.DevList {
		homeName := ""
		if val, ok := homeMap[iotutil.ToInt64(devInfo.HomeId)]; ok {
			homeName = val
		}
		productInfo, err := productCached.GetProduct(iotutil.ToString(devInfo.ProductKey))
		if err == nil && productInfo != nil {
			devInfo.ProductKey = productInfo.ProductKey
			devInfo.ProductPic = productInfo.ImageUrl
		}
		roomName := iotutil.GetMapVal(roomMap, devInfo.RoomId)
		if devInfo.RoomTemplateId != 0 {
			if dfVal, ok := defaultRooms[iotutil.ToString(devInfo.RoomTemplateId)]; ok {
				roomName = dfVal
			}
		}
		deviceList = append(deviceList, entitys.HomeDevice{
			Did:         devInfo.Did,
			ProductId:   devInfo.ProductId,
			DeviceName:  devInfo.DeviceName,
			ProductPic:  controls.ConvertProImg(devInfo.ProductPic),
			RoomName:    roomName,
			HomeName:    homeName,
			RoomId:      devInfo.RoomId,
			ProductKey:  devInfo.ProductKey,
			SecretKey:   devInfo.SecretKey,
			MqttServer:  devInfo.MqttServer,
			OnlineState: iotutil.ToInt32(devInfo.OnlineStatus),
			DevSwitch:   0,
			DevType:     1,
		})
	}
	return deviceList, ""
}

func (s AppHomeService) Delete(homeId int64, userId int64, ip string) (code int32, err error) {
	//删除之前查询，不然消息中心查询不到原来的家庭信息
	ret, err := rpc.UcHomeService.HomeDetail(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(err)
		return -1, err
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendRemoveHomeMessage").Error(ret.Message)
		return -1, err
	}

	//geo, err := iotutil.Geoip(ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode) //根据ip获取位置信息
	geo, err := controls.Geoip(ip)
	if err != nil {
		logger.Errorf("get address by ip[%s], error:%s", ip, err.Error())
	}
	resp, err := rpc.UcHomeService.Delete(s.Ctx, &protosService.UcHome{
		Id:       homeId,
		Lat:      geo.Lat,
		Lng:      geo.Lng,
		Country:  geo.Country,
		Province: geo.Province,
		City:     geo.City,
		District: geo.District,
	})
	if resp.Code != 200 {
		iotlogger.LogHelper.Error("删除家庭失败，原因:%s", resp.Message)
		return resp.Code, errors.New(resp.Message)
	}
	keys := []string{}
	for i := range ret.Data.UserList {
		keys = append(keys,
			persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, iotutil.ToString(ret.Data.UserList[i].Uid)),
			persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(ret.Data.UserList[i].Uid)),
		)
	}
	ctx := context.Background()
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return -1, err
		}
	}
	go SendRemoveHomeMessage(SetAppInfoByContext(s.Ctx), ret, userId, homeId)
	return 0, nil
}

func (s AppHomeService) AddDev(userId, homeId, thirdUserId int64) string {
	_, err := rpc.UcHomeService.TransferOwnership(context.Background(), &protosService.UcHomeDetailRequest{
		HomeId:      homeId,
		UserId:      userId,
		ThirdUserId: thirdUserId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("转移家庭所有权失败，原因:%s", err.Error())
		return err.Error()
	}
	return ""
}

func (s AppHomeService) SetDevSort(req entitys.SetDevSort) error {
	paramList := []*protosService.DevSortParam{}
	for _, v := range req.ParamList {
		paramList = append(paramList, &protosService.DevSortParam{
			Sort:  iotutil.ToInt32(v.Sort),
			DevId: v.DevId,
		})
	}
	_, err := rpc.IotDeviceHomeService.SetDevSort(context.Background(), &protosService.SetDevSortRequest{
		HomeId: iotutil.ToInt64(req.HomeId),
		DevId:  "",
		Data:   paramList,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s AppHomeService) GetDeviceStatus(did string) (res *entitys.DeviceCachedData) {
	res = &entitys.DeviceCachedData{}
	deviceStatus, redisErr := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+did).Result()
	iotlogger.LogHelper.Info("deviceStatus:", deviceStatus)
	if redisErr == nil {
		//在线状态
		if val, ok := deviceStatus[iotconst.FIELD_ONLINE]; ok {
			if val == "online" {
				iotlogger.LogHelper.Info("deviceStatus:online")
				res.OnlineStatus = 1
			} else {
				res.OnlineStatus = 0
			}
		}
		//开关状态
		if val, ok := deviceStatus["1"]; ok {
			if val == "true" {
				iotlogger.LogHelper.Info("switch:true")
				res.PowerState = 1
			} else {
				res.PowerState = 0
			}
		}
		//是否有升级
		if val, ok := deviceStatus[iotconst.FIELD_UPGRADE_HAS]; ok {
			if val == "true" {
				iotlogger.LogHelper.Info("hasOtaUpgrade:true")
				res.OtaUpgradeStatus = 1
			} else {
				res.OtaUpgradeStatus = 0
			}
		}

		//升级方式 1: APP提醒升级, 2: APP强制升级, 3: APP检测升级
		if val, ok := deviceStatus[iotconst.FIELD_UPGRADE_MODE]; ok {
			res.UpgradeMode, _ = iotutil.ToInt32Err(val)
		}
		//固件版本号
		if val, ok := deviceStatus[iotconst.FIELD_IS_FW_VER]; ok {
			res.Version = val
		}
		//UpgradeState
		if val, ok := deviceStatus[iotconst.FIELD_UPGRADE_STATE]; ok {
			res.UpgradeState = val
		}

		//有强制升级
		res.HasForceUpgrade = res.UpgradeMode == 2 && res.OtaUpgradeStatus == 1
		//是否正在升级
		if val, ok := deviceStatus[iotconst.FIELD_UPGRADE_RUNNING]; ok {
			var needClear int = 0 //1表示可以清理 0表示无需清理
			res.IsUpgradeRunning = val == "true"
			//如果正在升级需要判断，则需要检查是否已超时
			if res.IsUpgradeRunning {
				//获取超时时间
				var overtime int32 = 300
				if val, ok := deviceStatus[iotconst.FIELD_UPGRADE_TIMEOUT]; ok {
					tOvertime, err := iotutil.ToInt32Err(val)
					if err == nil && tOvertime != 0 {
						overtime = tOvertime
					}
				}
				//是否有升级
				if tVal, ok := deviceStatus[iotconst.FIELD_UPGRADE_TIME]; ok {
					if tVal != "" {
						upgradeTimeInt, err := iotutil.ToInt64AndErr(tVal)
						if err != nil {
							needClear = 1
						} else {
							currUtc := time.Now().UTC().Add(-time.Duration(overtime) * time.Second).Unix()
							//如果当前时间减去300秒，还是大于升级时间，则为升级超时
							//另外如果设备推送上线消息，也将清理正在升级状态
							if currUtc > upgradeTimeInt {
								needClear = 1
							}
						}
					}
				}
				//是否需要清理
				if needClear != 0 {
					res.IsUpgradeRunning = false
					res.UpgradeState = ""
					//清除升级状态
					iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+did, map[string]string{
						iotconst.FIELD_UPGRADE_RUNNING: "false",
						iotconst.FIELD_UPGRADE_STATE:   "",
					}).Result()
				}
			} else {
				res.UpgradeState = "" //防止IsUpgradeRunning=true，但是upgradeState存在的情况
			}
		}
	}
	return
}

func (s AppHomeService) translateHomeName(homeName, lang string) string {
	if homeName == _const.DefaultHomeName {
		return homeName
	}
	//读取翻译内容
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_DFAULT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
	}
	if len(langMap) == 0 {
		return homeName
	}
	return iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_default_home_name", lang)], _const.DefaultHomeName)
}

func GetRegionMqtt(c *gin.Context) (*protosService.SysRegionServer, error) {
	regionId := controls.GetRegion(c)
	return GetRegionMqttById(regionId)
}

// TODO 需要考虑缓存处理
func GetRegionMqttById(sid string) (*protosService.SysRegionServer, error) {
	regionId := sid
	//区域兼容处理regionId := controls.GetRegion(c)
	if regionId != "" {
		//读取区域数据
		regionIdInt, err := iotutil.ToInt64AndErr(regionId)
		if err == nil {
			rep, err := rpc.SysRegionServerService.FindById(context.Background(), &protosService.SysRegionServerFilter{Id: regionIdInt})
			if err == nil {
				return rep.Data[0], nil
			}
		} else {
			return nil, errors.New("区域参数错误，region: " + regionId)
		}
	}
	return nil, errors.New("异常")
}

func GetRegionMqttByHomeId(homeId int64) (*protosService.SysRegionServer, error) {
	home, err := rpc.UcHomeService.FindById(context.Background(), &protosService.UcHomeFilter{Id: homeId})
	if err != nil {
		return nil, err
	}
	if home.Code != 200 {
		return nil, errors.New(home.Message)
	}
	return GetRegionMqttById(iotutil.ToString(home.Data[0].Sid))
}
