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
	"fmt"
	"time"
)

type ReportDeviceSvc struct {
	Data  *iotstruct.MqttToNatsData
	DevId string
}

// ReportDevice 上报控制
func (s *ReportDeviceSvc) ReportDevice() {
	defer iotutil.PanicHandler()
	var (
		err error
	)
	if err := s.check(); err != nil {
		return
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Payload.(map[string]interface{})
	var payloadControl map[string]interface{}
	if _, ok := payloadMap["control"]; !ok {
		return
	}
	switch payloadMap["control"].(type) {
	case map[string]interface{}:
		payloadControl = payloadMap["control"].(map[string]interface{})
	default:
		iotlogger.LogHelper.Error("control格式异常")
		return
	}
	///**
	//{"header":{"ns":"iot.device.report","name":"control","mid":"b2694a93-15f8-4504-8d6f-73f7f17cdb0d","gid":"","ts":1668060656,"ver":"1.0.0"},
	//"payload":{"control":{"1":false}}}
	//{
	//  "header": {
	//    "ns": "iot.device.control",
	//    "name": "control",
	//    "mid": "1bd5d003-31b9-476f-ad03-71d471922820",
	//    "ts": 1632618621717,
	//    "ver": "2.0.6.0002",
	//    "gid": ""
	//  },
	//  "payload": {
	//    "control": {
	//      "1": true,
	//      "2": 1,
	//      "3": 4,
	//      "4": 12,
	//      "5": 53432,
	//    }
	//  }
	//}
	//*/
	device, err := s.getDeviceCached()
	if err != nil {
		iotlogger.LogHelper.Error("设备信息获取失败，设备Id=" + s.DevId)
		return
	}
	svc := IotDeviceLogProductSvc{Ctx: context.Background()}
	err = svc.ReportControl(s.Data, s.DevId, device["productId"], payloadControl)
	if err != nil {
		iotlogger.LogHelper.Error("清理缓存逻辑", s.DevId, err.Error())
		return
	}
	return
}

// ReportAckDevice 上报结果（过时）
func (s *ReportDeviceSvc) ReportAckDevice() {
	defer iotutil.PanicHandler()
	var (
		err error
	)
	if err := s.check(); err != nil {
		return
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Payload.(map[string]interface{})
	//**
	//{
	//  "header": {
	//    "ns": "iot.device.control",
	//    "name": "control",
	//    "mid": "1bd5d003-31b9-476f-ad03-71d471922820",
	//    "ts": 1632618621717,
	//    "ver": "2.0.6.0002",
	//    "gid": ""
	//  },
	//  "payload": {
	//    "code": 0,
	//    "control": {
	//      "1": true,
	//      "2": 1,
	//      "3": 4,
	//      "4": 12,
	//      "5": 53432,
	//    }
	//  }
	//}
	//*/
	device, err := s.getDeviceCached()
	if err != nil {
		iotlogger.LogHelper.Error("设备信息获取失败，设备Id=" + s.DevId)
		return
	}
	productKey := device["productKey"]
	if productKey == "" {
		iotlogger.LogHelper.Error("产品Key获取失败，设备Id=" + s.DevId)
		return
	}
	svc := IotDeviceLogProductSvc{Ctx: context.Background()}
	count, err := svc.FindReportLog(s.Data.MID, productKey)
	if err != nil {
		iotlogger.LogHelper.Error("日志获取失败，消息Id=" + s.Data.MID)
		return
	}
	if count == 0 {
		var payloadControl map[string]interface{}
		if _, ok := payloadMap["device"]; !ok {
			return
		}
		switch payloadMap["device"].(type) {
		case map[string]interface{}:
			payloadControl = payloadMap["device"].(map[string]interface{})
		default:
			iotlogger.LogHelper.Error("control格式异常")
			return
		}
		err = svc.ReportControl(s.Data, s.DevId, device["productId"], payloadControl)
		if err != nil {
			iotlogger.LogHelper.Error("日志存储失败", s.DevId, err.Error())
			return
		}
	} else {
		err = svc.ReportControlAck(s.Data, s.DevId, productKey, payloadMap)
		if err != nil {
			iotlogger.LogHelper.Error("日志ack修改失败", s.DevId, err.Error())
		}
	}
	return
}

// ReportMsgDevice 上报结果判断，清理app首页缓存 （TODO 待优化）
func (s *ReportDeviceSvc) ReportMsgDevice() {
	// 清除app家庭详情缓存
	ctx := context.Background()
	svc := IotDeviceHomeSvc{Ctx: ctx}
	reqDevHome := &protosService.IotDeviceHomeFilter{
		DeviceId: s.Data.DeviceId,
	}
	respDevHome, err := svc.FindIotDeviceHome(reqDevHome)
	if err != nil || respDevHome == nil {
		iotlogger.LogHelper.Error("查询设备家庭信息错误: ", s.Data.DeviceId, err.Error())
		return
	}
	// 删除家庭详情缓存
	pushClearHomeCached(respDevHome.HomeId)
	return
}

func (s *ReportDeviceSvc) check() error {
	if s.Data == nil {
		return errors.New("必须初始化data")
	}
	if s.Data.Payload == nil {
		return errors.New("参数异常，Payload")
	}
	return nil
}

// findDeviceTriad 设置设备信息
func (s *ReportDeviceSvc) getDeviceTriad() (*proto.IotDeviceTriad, error) {
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

func (s *ReportDeviceSvc) getDeviceCached() (map[string]string, error) {
	if s.DevId == "" {
		return nil, errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId)
	if newDeviceStatusCmd.Err() != nil {
		return nil, newDeviceStatusCmd.Err()
	}
	return newDeviceStatusCmd.Val(), nil
}

func (s *ReportDeviceSvc) sendAppMessage(appKey, tenantId, devId string, homeId int64, userId int64, isSuccess bool, params map[string]string) {
	defer iotutil.PanicHandler(appKey, tenantId, devId, homeId, userId, params, rpcClient.ClientAppMessage)
	tplCode := iotconst.APP_MESSAGE_DEVICE_UPGRADE_SUCCESS
	if !isSuccess {
		tplCode = iotconst.APP_MESSAGE_DEVICE_UPGRADE_FAIL
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
		fmt.Println(err)
		return
	}
	if ret.Code == 200 {
		fmt.Println(ret.Message)
		return
	}
}
