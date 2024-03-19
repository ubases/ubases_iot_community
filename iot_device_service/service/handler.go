package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_device_service/config"
)

func PushNatsData(productkey, deviceKey, gid, topic, subject string, ts int64, payload interface{}) {
	//推送到消息队列
	natsObj := iotstruct.MqttToNatsData{ProductKey: productkey, DeviceId: deviceKey, Gid: gid, Time: ts, Topic: topic, Payload: payload}
	data, err := json.MarshalToString(natsObj)
	if err != nil {
		iotlogger.LogHelper.Errorf("PushNatsData: MarshalToString:%v,error:%s", natsObj, err.Error())
		return
	}
	GetJsPublisherMgrJob().PushData(&NatsPubDataJob{Subject: subject, Data: data})
}

func RunJob() {
	// 初始化nats消息队列，新增发布者
	if err := GetJsPublisherMgrJob().AddPublisher(iotconst.NATS_SUBJECT_DEVICE_JOB, config.Global.Nats.Addrs); err != nil {
		iotlogger.LogHelper.Error("nats.AddPublisher failed:%s", err)
		return
	}

	go GetJsPublisherMgrJob().Run()

	// cron定时任务初始化
	NewCron()
	defer GetCron().Stop()
	InitCron()
}
