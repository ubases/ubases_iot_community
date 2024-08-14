package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_device_service/cached"
	"cloud_platform/iot_device_service/rpc/rpcClient"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_model/db_device/orm"
	"cloud_platform/iot_proto/protos/protosService"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gen/field"

	"go-micro.dev/v4/logger"
)

type ActiveDeviceSvc struct {
	Data            iotstruct.MqttToNatsData
	DevId           string
	localIp         string
	ssid            string
	rssi            string
	secretKey       string
	macAddress      string
	deviceVersion   string
	mcuVer          string
	hwVer           string
	extends         []map[string]interface{}
	zigbeeVer       string
	btVer           string
	sn              string
	userName        string
	password        string
	token           string
	useType         int32
	firstActiveTime *timestamppb.Timestamp
	cachedInfo      *iotstruct.DeviceNetworkTokenCacheModel
}

// ActiveDevice 激活设备处理
func (s *ActiveDeviceSvc) ActiveDevice() error {
	defer iotutil.PanicHandler()
	var err error
	//如果存在token，则配网初次提交
	if err = s.check(); err != nil {
		return err
	}
	payLoadMap := s.Data.Payload.(map[string]interface{})
	s.DevId = s.Data.DeviceId
	tokenVal, ok := payLoadMap["token"]
	if !ok || tokenVal == "" {
		s.setVersion(s.DevId, payLoadMap)
		return nil
	}
	s.token = iotutil.ToString(tokenVal)
	if tokenVal == "" {
		return nil
	}
	s.secretKey = iotutil.ToString(payLoadMap["secrtKey"]) //TODO secretKey
	if val, ok := payLoadMap["netif"]; ok {
		netif := val.(map[string]interface{})
		s.localIp = iotutil.ToString(netif["localIp"])
	}
	s.macAddress = iotutil.ToString(payLoadMap["mac"])
	s.deviceVersion = iotutil.ToString(payLoadMap["fwVer"])
	s.mcuVer = iotutil.ToString(payLoadMap["mcuVer"])
	s.hwVer = iotutil.ToString(payLoadMap["hwVer"])
	s.zigbeeVer = iotutil.ToString(payLoadMap["zigbeeVer"])
	s.btVer = iotutil.ToString(payLoadMap["btVer"])
	if val, ok := payLoadMap["extends"]; ok {
		if val != nil && iotutil.ToString(val) != "" && iotutil.ToString(val) != "null" {
			var extends []map[string]interface{}
			if err := json.Unmarshal([]byte(iotutil.ToString(val)), &extends); err == nil {
				s.extends = extends
			}
		}
	}
	//获取网络参数
	if val, ok := payLoadMap["ap"]; ok {
		ap := val.(map[string]interface{})
		ssid := iotutil.ToString(ap["ssid"])
		if ssid != "" {
			s.ssid = ssid
		}
		rssi := iotutil.ToString(ap["rssi"])
		if rssi != "" {
			s.rssi = rssi
		}
	}

	//通过token去获取配网的用户信息
	s.cachedInfo, err = s.checkCached()
	if err != nil {
		s.networkTokenSetDeviceResult(s.token, s.DevId, 2, err.Error())
		iotlogger.LogHelper.Error(err.Error())
		return err
	}
	//记录消息
	//s.Data.Topic
	defer pushAppLog(s.cachedInfo.Account, "操作日志", "设备激活",
		s.localIp, s.deviceVersion, s.cachedInfo.AppKey, s.cachedInfo.TenantId, err)

	if err = s.findDeviceTriad(); err != nil {
		s.networkTokenSetDeviceResult(s.token, s.DevId, 2, "三元组查询失败"+err.Error())
		iotlogger.LogHelper.Error("三元组不存在, token="+s.token, err.Error())
		return err
	}

	//TODO 验证是否为同一个开发者的设备
	//if triad.TenantId != cachedInfo.TenantId {
	//	s.networkTokenSetDeviceResult(token, devId, 2, "当前APP无法配网该产品")
	//	iotlogger.LogHelper.Error("当前APP无法配网该产品, tenantId=" + cachedInfo.TenantId)
	//	return err
	//}
	//tenantId := triad.TenantId
	q := orm.Use(iotmodel.GetDB())
	err = q.Transaction(func(tx *orm.Query) error {
		//激活数据存储
		//将三元组信息激活
		deviceTriadSvc := IotDeviceTriadSvc{Ctx: context.Background()}
		t := tx.TIotDeviceTriad
		var updateField []field.Expr
		updateField = append(updateField, t.UserId)
		updateField = append(updateField, t.UserAccount)
		updateField = append(updateField, t.AppName)
		updateField = append(updateField, t.AppKey)
		updateField = append(updateField, t.Status)
		updateField = append(updateField, t.ProductKey)
		updateField = append(updateField, t.ProductId)
		updateField = append(updateField, t.DeviceNatureKey)
		//updateField = append(updateField, t.DeviceSecret)
		if s.firstActiveTime == nil || !s.firstActiveTime.IsValid() || s.firstActiveTime.Nanos == 0 {
			updateField = append(updateField, t.FirstActiveTime)
		}
		err := deviceTriadSvc.SetDeviceUser(tx, s.DevId, updateField, &model.TIotDeviceTriad{
			UserId:          s.cachedInfo.UserId,
			AppName:         s.cachedInfo.AppName,
			UserAccount:     s.cachedInfo.Account,
			AppKey:          s.cachedInfo.AppKey,
			ProductId:       s.cachedInfo.ProductId,
			ProductKey:      s.cachedInfo.ProductKey,
			DeviceNatureKey: s.cachedInfo.DeviceNature,
			FirstActiveTime: time.Now(),
			Status:          1,
		})
		if err != nil {
			iotlogger.LogHelper.Error("设备家庭绑定失败, err=" + err.Error())
			return err
		}
		//创建设备信息
		if err := s.saveDeviceInfo(tx); err != nil {
			iotlogger.LogHelper.Error("设备家庭绑定失败, err=" + err.Error())
			return err
		}
		//缓存设备信息
		if err := s.cachedDeviceInfo(); err != nil {
			logger.Errorf("FindByIdIotDeviceInfo redis error : %s", err.Error())
			return err
		}
		//激活设备（设备信息激活，设备家庭关系帮规定）
		if s.saveDeviceHome(tx) != nil {
			iotlogger.LogHelper.Error("设备家庭绑定失败, err=" + err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		//配网结果存储
		s.networkTokenSetDeviceResult(s.token, s.DevId, 2, err.Error())
		iotlogger.LogHelper.Error("设备家庭绑定失败, err=" + err.Error())
	} else {
		//配网结果存储
		s.networkTokenSetDeviceResult(s.token, s.DevId, 1, "ok")
		s.sendDeviceActiveAppMessage(s.cachedInfo.AppKey, s.cachedInfo.TenantId, s.DevId, s.cachedInfo.HomeId, s.cachedInfo.UserId, map[string]string{
			"userName":   s.cachedInfo.UserName,
			"homeName":   s.cachedInfo.HomeName,
			"deviceName": s.cachedInfo.ProductName,
		})
	}
	// 推送清理家庭缓存
	pushClearHomeCached(s.cachedInfo.HomeId)
	//推送设备数据更新
	s.pushDeviceInfoUpdate()
	return nil
}

// 推送清理家庭缓存
func pushClearHomeCached(homeId int64) error {
	if homeId == 0 {
		return nil
	}
	du := iotstruct.DeviceRedisUpdate{
		HomeId: iotutil.ToString(homeId),
	}
	duBytes, err := json.Marshal(du)
	if err != nil {
		return err
	}
	//err = cached.RedisStore.GetClient().Publish(context.Background(), strings.Join([]string{iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX, iotutil.ToString(homeId)}, "."), string(duBytes)).Err()
	//if err != nil {
	//	return err
	//}

	data := iotnatsjs.NatsPubData{
		Subject: strings.Join([]string{iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX, iotutil.ToString(homeId)}, "."),
		Data:    string(duBytes),
	}
	iotnatsjs.GetJsClientPub().PushData(&data)
	return nil
}

// 推送通知数据更新（语音控制服务）
func (s *ActiveDeviceSvc) pushDeviceInfoUpdate() error {
	if s.cachedInfo.UserId == 0 {
		return nil
	}
	du := iotstruct.DeviceRedisUpdate{
		UserId: iotutil.ToString(s.cachedInfo.UserId),
	}
	duBytes, err := json.Marshal(du)
	if err != nil {
		return err
	}
	err = cached.RedisStore.GetClient().Publish(context.Background(), strings.Join([]string{iotconst.HKEY_UPDATE_DATA_PUB_PREFIX, iotutil.ToString(s.cachedInfo.UserId)}, "."), string(duBytes)).Err()
	if err != nil {
		return err
	}
	return nil
}

// 设备设备缓存
func (s *ActiveDeviceSvc) setDeviceCached(data map[string]string) error {
	if s.DevId == "" {
		return errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId, data)
	if newDeviceStatusCmd.Err() != nil {
		return newDeviceStatusCmd.Err()
	}
	return nil
}

func (s *ActiveDeviceSvc) check() error {
	if s.Data.Payload == nil {
		return errors.New("参数异常")
	}
	return nil
}

func (s *ActiveDeviceSvc) checkCached() (cacheInfo *iotstruct.DeviceNetworkTokenCacheModel, err error) {
	//通过token去获取配网的用户信息
	cachedInfo, err := s.getTokenCacheInfo(s.token)
	if err != nil {
		s.networkTokenSetDeviceResult(s.token, s.DevId, 2, "配网Token无效，"+s.token)
		iotlogger.LogHelper.Error("token获取用户信息失败, token=" + s.token)
		return nil, err
	}
	if cachedInfo.HomeId == 0 {
		s.networkTokenSetDeviceResult(s.token, s.DevId, 2, "配网Token无效，"+s.token)
		iotlogger.LogHelper.Error("token获取用户信息失败, homeId=0")
		return nil, err
	}
	return cachedInfo, nil
}

// SetVersion 设置Redis设备版本信息
func (s *ActiveDeviceSvc) setVersion(devId string, payLoadMap map[string]interface{}) {
	var (
		localIp  = ""
		setRedis = make(map[string]interface{})
	)
	//升级成功需要更新升级信息
	setRedis[iotconst.FIELD_UPGRADE_STATE] = ""
	setRedis[iotconst.FIELD_UPGRADE_PROGRESS] = 0

	if val, ok := payLoadMap["netif"]; ok {
		netIf := val.(map[string]interface{})
		localIp = iotutil.ToString(netIf["localIp"])
		if localIp != "" {
			setRedis["localIp"] = localIp
		}
	}
	if val, ok := payLoadMap["ap"]; ok {
		ap := val.(map[string]interface{})
		ssid := iotutil.ToString(ap["ssid"])
		if ssid != "" {
			setRedis["ssid"] = ssid
		}
		rssi := iotutil.ToString(ap["rssi"])
		if rssi != "" {
			setRedis["rssi"] = rssi
		}
	}
	if val, ok := payLoadMap["fwVer"]; ok {
		fwVer := iotutil.ToString(val)
		if fwVer != "" {
			setRedis["fwVer"] = fwVer
		}
	}
	if val, ok := payLoadMap["mcuVer"]; ok {
		mcuVer := iotutil.ToString(val)
		if mcuVer != "" {
			setRedis["mcuVer"] = mcuVer
		}
	}
	if val, ok := payLoadMap["hwVer"]; ok {
		hwVer := iotutil.ToString(val)
		if hwVer != "" {
			setRedis["hwVer"] = hwVer
		}
	}
	if val, ok := payLoadMap["mac"]; ok {
		mac := iotutil.ToString(val)
		if mac != "" {
			setRedis["mac"] = mac
		}
	}
	if val, ok := payLoadMap["memFree"]; ok {
		memFree := iotutil.ToString(val)
		if memFree != "" {
			setRedis["memFree"] = memFree
		}
	}
	if val, ok := payLoadMap["zigbeeVer"]; ok {
		zigbeeVer := iotutil.ToString(val)
		if zigbeeVer != "" {
			setRedis["zigbeeVer"] = zigbeeVer
		}
	}
	if val, ok := payLoadMap["btVer"]; ok {
		btVer := iotutil.ToString(val)
		if btVer != "" {
			setRedis["btVer"] = btVer
		}
	}
	if val, ok := payLoadMap["extends"]; ok {
		setRedis["extends"] = iotutil.ToString(val)
	}

	if len(setRedis) > 0 {
		newDeviceStatusCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+devId, setRedis)
		if newDeviceStatusCmd.Err() != nil {
			logger.Errorf("FindByIdIotDeviceInfo SetVersion error : %s", newDeviceStatusCmd.Err().Error())
		}
	}
}

// saveDeviceInfo 设置设备信息
func (s *ActiveDeviceSvc) saveDeviceInfo(tx *orm.Query) error {
	deviceInfoSvc := IotDeviceInfoSvc{Ctx: context.Background()}
	err := deviceInfoSvc.TranCreate(tx, &model.TIotDeviceInfo{
		Id:                iotutil.GetNextSeqInt64(),
		Did:               s.DevId,
		ProductId:         s.cachedInfo.ProductId,
		ProductKey:        s.cachedInfo.ProductKey,
		OnlineStatus:      1,
		DeviceName:        s.cachedInfo.ProductName,
		Sn:                s.sn,
		DeviceModel:       s.cachedInfo.ProductKey,
		UserName:          s.userName,
		Passward:          s.password,
		Salt:              s.secretKey,          //triad.Salt
		DeviceSecretHttp:  iotutil.GetSecret(6), //t
		DeviceSecretMqtt:  iotutil.GetSecret(6),
		IpAddress:         s.localIp,
		Lat:               s.cachedInfo.Lat,
		Lng:               s.cachedInfo.Lng,
		Country:           s.cachedInfo.Country,
		Province:          s.cachedInfo.Province,
		City:              s.cachedInfo.City,
		District:          s.cachedInfo.District,
		ActivatedTime:     time.Now(),
		MacAddress:        s.macAddress,
		DeviceVersion:     s.deviceVersion,
		ActiveStatus:      "1",
		LastActivatedTime: time.Now(),
		ActiveUserId:      s.cachedInfo.UserId,
		ActiveUserName:    s.cachedInfo.UserName,
		TenantId:          s.cachedInfo.TenantId,
		AppKey:            s.cachedInfo.AppKey,
		UseType:           s.useType,
		RegionServerId:    s.cachedInfo.RegionServerId,
	})
	if err != nil {
		return err
	}
	return nil
}

// saveDeviceInfo 设置设备信息
func (s *ActiveDeviceSvc) findDeviceTriad() error {
	deviceTriadSvc := IotDeviceTriadSvc{Ctx: context.Background()}
	triad, err := deviceTriadSvc.FindIotDeviceTriad(&proto.IotDeviceTriadFilter{Did: s.DevId})
	if err != nil {
		return err
	}
	s.userName = triad.UserName
	s.password = triad.Passward
	s.sn = triad.SerialNumber
	s.useType = triad.UseType
	s.firstActiveTime = triad.FirstActiveTime
	return nil
}

// cachedDeviceInfo 缓存设备信息
func (s *ActiveDeviceSvc) cachedDeviceInfo() error {
	//先清理在创建
	iotredis.GetClient().Del(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId)
	//缓存设备信息
	var deviceInitStatus map[string]interface{} = map[string]interface{}{
		"productKey":   s.Data.ProductKey,
		"serialNumber": s.sn,
		"productId":    s.cachedInfo.ProductId,
		"did":          s.DevId,
		"productName":  s.cachedInfo.ProductName,
		"userId":       s.cachedInfo.UserId,
		"homeId":       s.cachedInfo.HomeId,
		"appKey":       s.cachedInfo.AppKey,
		"tenantId":     s.cachedInfo.TenantId,
		"country":      s.cachedInfo.Country,
		"province":     s.cachedInfo.Province,
		"city":         s.cachedInfo.City,
		"district":     s.cachedInfo.District,
		"lat":          s.cachedInfo.Lat,
		"lng":          s.cachedInfo.Lng,
		"localIp":      s.localIp,
		"mac":          s.macAddress,
		"ssid":         s.ssid,
		"rssi":         s.rssi,
		"fwVer":        s.deviceVersion,
		"extends":      iotutil.ToString(s.extends),
		"zigbeeVer":    s.zigbeeVer,
		"btVer":        s.btVer,
	}
	newDeviceStatusCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId, deviceInitStatus)
	if newDeviceStatusCmd.Err() != nil {
		return newDeviceStatusCmd.Err()
	}
	return nil
}

// saveDeviceHome 设置设备家庭信息
func (s *ActiveDeviceSvc) saveDeviceHome(tx *orm.Query) error {
	deviceHomeSvc := IotDeviceHomeSvc{Ctx: context.Background()}
	err := deviceHomeSvc.TranCreate(tx, &model.TIotDeviceHome{
		Id:         iotutil.GetNextSeqInt64(),
		HomeId:     s.cachedInfo.HomeId,
		DeviceId:   s.DevId,
		RoomId:     0,
		ProductId:  s.cachedInfo.ProductId,
		ProductKey: s.cachedInfo.ProductKey,
		CustomName: s.cachedInfo.ProductName,
		BindTime:   time.Now(),
		Secrtkey:   s.secretKey,
		AddMethod:  0,
		CreatedBy:  s.cachedInfo.UserId,
		UpdatedBy:  s.cachedInfo.UserId,
	})
	if err != nil {
		return err
	}
	return nil
}

// 从token中获取用户名称（redis，token作为key）
func (s *ActiveDeviceSvc) getTokenCacheInfo(token string) (cacheInfo *iotstruct.DeviceNetworkTokenCacheModel, err error) {
	var cacheData iotstruct.DeviceNetworkTokenCacheModel
	valueCmd := iotredis.GetClient().Get(context.Background(), token)
	val := valueCmd.Val()
	if val == "" {
		return nil, errors.New("无任何数据")
	}
	err = iotutil.JsonToStruct(val, &cacheData)
	if err != nil {
		err = errors.New("redis cache convert error " + val)
		return nil, err
	}
	return &cacheData, nil
}

// 设置设备配网结果，如果成功一个设备，会将设备的devid写入到redis的devices中
// 考虑到全局redis过期时间不能被修改，所以采用SETRANGE
// result = 1 成功 = 0 等待 = 2 失败
func (s *ActiveDeviceSvc) networkTokenSetDeviceResult(token string, devId string, result int, msg string) (err error) {
	var cacheData iotstruct.DeviceNetworkTokenCacheModel
	valueCmd := iotredis.GetClient().Get(context.Background(), token)
	val := valueCmd.Val()
	if val == "" {
		return errors.New("无任何数据")
	}
	err = iotutil.JsonToStruct(val, &cacheData)
	if err != nil {
		err = errors.New("redis cache convert error " + val)
		return
	}
	cacheData.Devices = append(cacheData.Devices, devId)
	if cacheData.DevicesMap == nil {
		cacheData.DevicesMap = map[string]iotstruct.DeviceResult{}
	}
	cacheData.DevicesMap[devId] = iotstruct.DeviceResult{
		Code: result,
		Msg:  msg,
	}
	setCmd := iotredis.GetClient().SetRange(context.Background(), token, 0, iotutil.ToString(cacheData))
	_, err = setCmd.Result()
	if err != nil {
		err = errors.New("缓存配网结果失败 " + err.Error())
	}
	return
}

func (s *ActiveDeviceSvc) sendDeviceActiveAppMessage(appKey, tenantId, devId string, homeId int64, userId int64, params map[string]string) {
	defer iotutil.PanicHandler(appKey, tenantId, devId, homeId, userId, params, rpcClient.ClientAppMessage)
	//发送消息  测试消息推送
	ret, err := rpcClient.ClientAppMessage.SendMessage(context.Background(), &protosService.SendMessageRequest{
		TplCode:     iotconst.APP_MESSAGE_DEVICE_ACTIVE,
		Params:      params,
		TimeUnix:    time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		SourceTable: model.TableNameTIotDeviceInfo,
		SourceRowId: devId,
		HomeId:      homeId,
		UserId:      []int64{userId},
		IsPublic:    false,
		PushTo:      "device",
		ChildType:   9,
		Subject:     "添加设备",
		Lang:        "", //不指定语言则，则全语言推送
		AppKey:      appKey,
		TenantId:    tenantId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if ret.Code == 200 {
		fmt.Println(ret.Message)
		return
	}
}
