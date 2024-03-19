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

type OnlineDeviceSvc struct {
	Data  *iotstruct.MqttToNatsData
	DevId string
}

// OnlineDevice 在线设备处理
func (s *OnlineDeviceSvc) OnlineDevice() {
	defer iotutil.PanicHandler()
	if err := s.check(); err != nil {
		return
	}
	s.DevId = s.Data.DeviceId
	payloadMap := s.Data.Payload.(map[string]interface{})
	onlineStatus := payloadMap[iotconst.FIELD_ONLINE]
	var (
		isPushOfflineMsg bool   = false //离线之前是否推送消息
		isPushOnlineMsg  bool   = false //上线之前是否推送过
		productKey       string = s.Data.ProductKey
		appKey           string
		tenantId         string
		userId           int64
		deviceIsRemove   bool = false
	)
	//从缓存获取设备信息
	device, err := s.getDeviceCached()
	if err != nil || device == nil || device["productKey"] == "" {
		//如果缓存读取失败从，查询数据获取设备信息（移除设备之后，缓存数据会被清理，所以获取缓存会报错）
		deviceInfo, err := s.getDeviceTriad()
		if err != nil {
			iotlogger.LogHelper.WithTag("online", "OnlineDevice").Error("设备信息获取失败，设备Id=" + s.DevId)
			return
		}
		deviceIsRemove = deviceInfo.Status == 0
		//productKey = deviceInfo.ProductKey
		appKey = deviceInfo.AppKey
		tenantId = deviceInfo.TenantId
		userId = deviceInfo.UserId
	} else {
		isPushOfflineMsg = device[iotconst.FIELD_IS_PUSH_OFFLINE_MSG] == "false"
		isPushOnlineMsg = device[iotconst.FIELD_IS_PUSH_ONLINE_MSG] == "false"
		//productKey = device["productKey"]
		appKey = device["appKey"]
		tenantId = device["tenantId"]
		userId, _ = iotutil.ToInt64AndErr(device["userId"])
	}
	//修改设备信息
	s.updateDeviceInfo(iotutil.ToString(onlineStatus))

	//存储上下线信息
	svc := IotDeviceLogProductSvc{Ctx: context.Background()}
	err = svc.ReportOnline(s.Data, s.DevId, iotutil.ToString(onlineStatus))
	if err != nil {
		iotlogger.LogHelper.Error("存储上报上下线信息失败", s.DevId, err.Error())
	}

	//设备未激活，但是收到设备离线消息，认为设备已删除，推送移除设备
	if deviceIsRemove {
		if productKey == "" {
			iotlogger.LogHelper.WithTag("online", "OnlineDevice").Error("参数异常:", s.DevId)
			return
		}
		publishRemoveDevice([]*model.TIotDeviceHome{
			{
				ProductKey: productKey,
				DeviceId:   s.DevId,
			},
		})
	} else {
		if err = s.checkCached(productKey, appKey, tenantId); err != nil {
			iotlogger.LogHelper.WithTag("online", "OnlineDevice").Error("参数异常:", err.Error())
			return
		}
		//设备离线是否推送离线消息
		if onlineStatus == "offline" {
			//如果设备离线，并且离线之前未推送过，则执行推送（避免反复推送）
			s.setDeviceCached(map[string]string{
				iotconst.FIELD_IS_PUSH_OFFLINE_MSG: "true",
				iotconst.FIELD_IS_PUSH_ONLINE_MSG:  "false",
				iotconst.FIELD_UPGRADE_RUNNING:     "false", //设备上报在线状态，需要清理ota升级状态
			})
			if isPushOfflineMsg {
				s.sendAppMessage(appKey, tenantId, s.DevId, 0, userId, iotconst.APP_MESSAGE_DEVICE_OFFLINE, "设备离线", map[string]string{})
			}
		} else {
			//如果是在线，并且未推送过在线消息，则执行推送（避免反复推送）
			s.setDeviceCached(map[string]string{
				iotconst.FIELD_IS_PUSH_OFFLINE_MSG: "false",
				iotconst.FIELD_IS_PUSH_ONLINE_MSG:  "true",
				iotconst.FIELD_UPGRADE_RUNNING:     "false", //设备上报在线状态，需要清理ota升级状态
			})
			if isPushOnlineMsg {
				s.sendAppMessage(appKey, tenantId, s.DevId, 0, userId, iotconst.APP_MESSAGE_DEVICE_ONLINE, "设备上线", map[string]string{})
			}
		}
		//增加清理缓存逻辑
		s.OnlineMsgDevice()
		if err != nil {
			iotlogger.LogHelper.Error("清理缓存逻辑", s.DevId, err.Error())
			return
		}
	}
	return
}

// OnlineMsgDevice 上报结果判断，清理app首页缓存 （TODO 待优化）
func (s *OnlineDeviceSvc) OnlineMsgDevice() {
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
	pushClearHomeCached(respDevHome.HomeId)
	return
}

func (s *OnlineDeviceSvc) check() error {
	if s.Data == nil {
		return errors.New("必须初始化data")
	}
	if s.Data.Payload == nil {
		return errors.New("参数异常，Payload")
	}
	return nil
}

func (s *OnlineDeviceSvc) checkCached(productKey, appKey, tenantId string) error {
	if productKey == "" {
		return errors.New("产品Key获取失败")
	}
	if appKey == "" {
		return errors.New("APP Key获取失败")
	}
	if tenantId == "" {
		return errors.New("租户Id获取失败")
	}
	return nil
}

// findDeviceTriad 设置设备信息
func (s *OnlineDeviceSvc) getDeviceTriad() (*proto.IotDeviceTriad, error) {
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

func (s *OnlineDeviceSvc) getDeviceCached() (map[string]string, error) {
	if s.DevId == "" {
		return nil, errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId)
	if newDeviceStatusCmd.Err() != nil {
		return nil, newDeviceStatusCmd.Err()
	}
	return newDeviceStatusCmd.Val(), nil
}

func (s *OnlineDeviceSvc) setDeviceCached(data map[string]string) error {
	if s.DevId == "" {
		return errors.New("did不能为空")
	}
	newDeviceStatusCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+s.DevId, data)
	if newDeviceStatusCmd.Err() != nil {
		return newDeviceStatusCmd.Err()
	}
	return nil
}

// findDeviceTriad 设置设备信息
func (s *OnlineDeviceSvc) updateDeviceInfo(onlineStatus string) error {
	if s.DevId == "" {
		return errors.New("did不能为空")
	}
	deviceInfoSvc := IotDeviceInfoSvc{Ctx: context.Background()}
	var onlineStatusInt32 int32 = 2
	if onlineStatus == "online" {
		onlineStatusInt32 = 1
	}
	err := deviceInfoSvc.SetOnlineStatus(s.DevId, onlineStatusInt32)
	if err != nil {
		return err
	}
	return nil
}

func (s *OnlineDeviceSvc) sendAppMessage(appKey, tenantId, devId string, homeId int64, userId int64, tplCode, subject string, params map[string]string) {
	defer iotutil.PanicHandler(appKey, tenantId, devId, homeId, userId, params, rpcClient.ClientAppMessage)
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
		Subject:     subject,
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
