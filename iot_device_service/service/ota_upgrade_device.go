package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_device_service/rpc/rpcClient"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_proto/protos/protosService"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"
)

type OtaUpgradeDeviceSvc struct {
	Data  *iotstruct.MqttToNatsData
	DevId string
}

// OtaUpgradeDevice 在线设备处理
func (s *OtaUpgradeDeviceSvc) OtaUpgradeDevice() {
	defer iotutil.PanicHandler()
	var (
		err      error
		state    string
		progress *int32
		code     *int32
		version  string
		pubId    int64
	)
	if err := s.check(); err != nil {
		return
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Payload.(map[string]interface{})

	if val, has := payloadMap["otaState"]; has {
		state = iotutil.ToString(val)
	}
	if val, has := payloadMap["progress"]; has {
		theVal := iotutil.ToInt32(val)
		progress = &theVal
	}
	if val, has := payloadMap["code"]; has {
		theVal := iotutil.ToInt32(val)
		code = &theVal
	}
	if val, has := payloadMap["otaVer"]; has {
		version = iotutil.ToString(val)
	}
	if val, has := payloadMap["pubId"]; has {
		pubId, _ = iotutil.ToInt64AndErr(val)
	}
	if pubId == 0 {
		iotlogger.LogHelper.Error("OTA结果上报异常，pubId=0, payload=" + iotutil.ToString(payloadMap))
		return
	}
	iotlogger.LogHelper.Info(state, code, progress)
	//otaState: 该字段表示OTA的状态，主要有以下三种状态
	//	Downloading: 下载中，当OTA处于此状态时，设备将会上报下载进度。而上报进度时间间隔以10 - 30%的间隔上报。
	//	Installing: 安装中，当OTA处于此状态时，设备将会处于离线状态。而该状态是设备即将进入离线状态前发送给云端的。
	//code：
	//	0： 表示OTA成功
	//	1：表示下载失败
	//	2：表示安装失败
	//	3：协议数据错误
	//	4：OTA数据包错误
	//progress：OTA上报的进度
	//TODO getDeviceTriad 改为redis缓存
	//device, err := s.getDeviceTriad()
	//if err != nil {
	//	iotlogger.LogHelper.Error("设备信息获取失败，设备Id=" + s.DevId)
	//	return
	//}
	var (
		productKey string
		userId     int64
		homeId     int64
		tenantId   string
		appKey     string
		area       string
		fwVer      string
		//forceUpgradeVer string //强制升级版本
		//hasOtaUpgrade   bool   //是否存在升级
		//upgardeMode     int32  //是否强制升级
	)

	//{\"code\":0,\"progress\":90,\"otaState\":\"Downloading\",\"otaVer\":\"1.0.83\",\"pubId\":\"8679005822161879040\"}}

	device, err := s.getDeviceCached()
	if err != nil {
		iotlogger.LogHelper.Error("设备信息获取失败，设备Id=" + s.DevId)
		return
	} else {
		userId, _ = iotutil.ToInt64AndErr(device["userId"])
		homeId, _ = iotutil.ToInt64AndErr(device["homeId"])
		appKey = device["appKey"]
		tenantId = device["tenantId"]
		productKey = device["productKey"]
		area = device["country"]
		fwVer = device["fwVer"]
		//forceUpgradeVer = device[iotconst.FIELD_UPGRADE_FORCE_VER]
		//hasOtaUpgrade = device[iotconst.FIELD_UPGRADE_HAS] == "true"
		//upgardeMode, _ = iotutil.ToInt32Err(device[iotconst.FIELD_UPGRADE_MODE])
	}
	svc := IotOtaUpgradeRecordSvc{Ctx: context.Background()}
	err = svc.ReportUpgradeResult(s.DevId, productKey, version, tenantId, pubId, area, fwVer, state, code, progress)
	if err != nil {
		iotlogger.LogHelper.Error("修改进度失败", s.DevId, state, code, progress)
		return
	}
	isSuccess, err := s.checkStatusIsUpgradeSuccess(state, code)
	if err != nil {
		//推送固件升级结果消息
		s.sendAppMessage(appKey, tenantId, s.DevId, homeId, userId, false, map[string]string{})
		////推送升级结果通知
		//s.batchPublishOtaResultNotice(s.Data.ProductKey, s.DevId, version, 2)
	} else {
		if isSuccess {
			var needSetCached = true
			//判断版本当前升级版本是否大于ota升级版本
			compareRes, _ := iotutil.VerCompare(version, fwVer)
			if compareRes == 1 {
				needSetCached = false
			}
			if needSetCached {
				//升级成功需要更新升级信息
				s.setDeviceCached(map[string]string{
					iotconst.FIELD_UPGRADE_HAS:      "false",
					iotconst.FIELD_UPGRADE_MODE:     "0",
					iotconst.FIELD_UPGRADE_RUNNING:  "false",
					iotconst.FIELD_UPGRADE_STATE:    "",
					iotconst.FIELD_UPGRADE_PROGRESS: "0",
					//iotconst.FIELD_IS_FW_VER:        version,
				})
			} else {
				//升级成功需要更新升级信息
				s.setDeviceCached(map[string]string{
					//iotconst.FIELD_IS_FW_VER:        version,
					iotconst.FIELD_UPGRADE_HAS:      "false",
					iotconst.FIELD_UPGRADE_STATE:    "",
					iotconst.FIELD_UPGRADE_PROGRESS: "0",
					iotconst.FIELD_UPGRADE_RUNNING:  "false",
				})
			}
			//推送固件升级结果消息
			s.sendAppMessage(appKey, tenantId, s.DevId, homeId, userId, true, map[string]string{})
			//推送升级成功
			//s.batchPublishOtaResultNotice(s.Data.ProductKey, s.DevId, version, 1)
		} else {
			if s.isUpgradeRunning(state, code) {
				s.setDeviceCached(map[string]string{
					iotconst.FIELD_UPGRADE_HAS:      "true",
					iotconst.FIELD_UPGRADE_PROGRESS: iotutil.ToString(progress),
					iotconst.FIELD_UPGRADE_STATE:    state,
					iotconst.FIELD_UPGRADE_RUNNING:  "true",
				})
				////推送升级成功
				//s.batchPublishOtaResultNotice(s.Data.ProductKey, s.DevId, version, 3)
			}
		}
	}
	return
}

//func (s *OtaUpgradeDeviceSvc) batchPublishOtaResultNotice(productKey, deviceId, version string, code int32) error {
//	iotlogger.LogHelper.Debugf("开始推送OTA升级结果通知, productKey:%v, deviceId:%v, version:%v, code:%v", productKey, deviceId, version, code)
//	topics := make([]string, 0)
//	topics = append(topics, iotprotocol.GetTopic(iotprotocol.TP_E_NOTICE, productKey, deviceId))
//	//推送升级通知
//	var obj iotprotocol.PackNotice
//	dataMap := make(map[string]interface{})
//	dataMap["version"] = version
//	switch code {
//	case 1:
//		//升级成功
//		dataMap["otaUpgradeStatus"] = 0
//		dataMap["hasForceUpgrade"] = false
//	case 2:
//		//升级失败
//		dataMap["otaUpgradeStatus"] = 0
//		dataMap["hasForceUpgrade"] = false
//	case 3:
//		//升级中
//		dataMap["otaUpgradeStatus"] = 1
//	}
//	buf, _ := obj.Encode(iotprotocol.NOTICE_HEAD_UPGRADE_NOTICE_NAME, dataMap)
//	_, pubErr := rpcClient.ClientMqttService.BatchPublish(context.Background(), &proto.BatchPublishMessage{
//		TopicFullNameList: topics,
//		MessageContent:    string(buf),
//		Qos:               proto.Qos_ExactlyOnce,
//		Retained:          false,
//	})
//	if pubErr != nil {
//		iotlogger.LogHelper.Errorf("CreateIotOtaUpgradeRecord BatchPublish error : %s", pubErr.Error())
//		return pubErr
//	}
//	return nil
//}

func (s *OtaUpgradeDeviceSvc) isUpgradeRunning(otaState string, otaCode *int32) (isSuccess bool) {
	//转换ota状态
	if otaState == "Downloading" {
		return true //下载中
	} else if otaState == "Installing" {
		return true //安装中
	} else {
		return false
	}
}

func (s *OtaUpgradeDeviceSvc) checkStatusIsUpgradeSuccess(otaState string, otaCode *int32) (isSuccess bool, err error) {
	//转换ota状态
	if otaState == "Downloading" {
		//下载中
	} else if otaState == "Installing" {
		//安装中
	} else {
		if *otaCode == 0 {
			isSuccess = true
		} else {
			err = errors.New("升级失败")
		}
	}
	return
}

func (s *OtaUpgradeDeviceSvc) check() error {
	if s.Data == nil {
		return errors.New("必须初始化data")
	}
	if s.Data.Payload == nil {
		return errors.New("参数异常，Payload")
	}
	return nil
}

// findDeviceTriad 设置设备信息
func (s *OtaUpgradeDeviceSvc) getDeviceTriad() (*proto.IotDeviceTriad, error) {
	if s.DevId == "" {
		return nil, errors.New("did不能为空")
	}
	deviceTriadSvc := IotDeviceTriadSvc{Ctx: context.Background()}
	triad, err := deviceTriadSvc.FindIotDeviceTriad(&proto.IotDeviceTriadFilter{Did: s.DevId})
	if err != nil {
		return nil, err
	}
	return triad, nil
}

func (s *OtaUpgradeDeviceSvc) getDeviceCached() (map[string]string, error) {
	if s.DevId == "" {
		return nil, errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId)
	if newDeviceStatusCmd.Err() != nil {
		return nil, newDeviceStatusCmd.Err()
	}
	return newDeviceStatusCmd.Val(), nil
}

// 设备设备缓存
func (s *OtaUpgradeDeviceSvc) setDeviceCached(data map[string]string) error {
	if s.DevId == "" {
		return errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId, data)
	if newDeviceStatusCmd.Err() != nil {
		return newDeviceStatusCmd.Err()
	}
	return nil
}

func (s *OtaUpgradeDeviceSvc) sendAppMessage(appKey, tenantId, devId string, homeId int64, userId int64, isSuccess bool, params map[string]string) {
	defer iotutil.PanicHandler(appKey, tenantId, devId, homeId, userId, params, rpcClient.ClientAppMessage)
	tplCode := iotconst.APP_MESSAGE_DEVICE_UPGRADE_SUCCESS
	if !isSuccess {
		tplCode = iotconst.APP_MESSAGE_DEVICE_UPGRADE_FAIL
	}
	if userId == 0 {
		iotlogger.LogHelper.Error("消息推送失败,userId不能为空")
		return
	}
	//发送消息  测试消息推送
	ret, err := rpcClient.ClientAppMessage.SendMessage(context.Background(), &protosService.SendMessageRequest{
		TplCode:     tplCode,
		Params:      params,
		TimeUnix:    time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		SourceTable: model.TableNameTIotDeviceInfo,
		SourceRowId: devId,
		HomeId:      homeId,
		UserId:      []int64{userId},
		IsPublic:    false,
		PushTo:      "device",
		ChildType:   9,
		Subject:     "设备升级",
		Lang:        "", //不指定语言则，则全语言推送
		AppKey:      appKey,
		TenantId:    tenantId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("消息推送失败", err.Error())
		return
	}
	if ret.Code == 200 {
		iotlogger.LogHelper.Error("消息推送失败", ret.Message)
		return
	}
}
