package service

import (
	"time"

	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
)

type AppLog struct {
	Id        int64             `json:"id"`
	Account   string            `json:"account"`
	AppKey    string            `json:"appKey"`
	TenantId  string            `json:"tenantId"`
	LogType   string            `json:"logType"`
	EventName string            `json:"eventName"`
	Details   map[string]string `json:"details"`
	CreatedAt time.Time         `json:"createdAt"`
}

func pushAppLog(account, logType, eventName, ip, sys, appKey, tenantId string, err error) error {
	defer iotutil.PanicHandler(account, logType, eventName, ip, sys, appKey, tenantId)
	errMsg := "ok"
	if err != nil {
		errMsg = err.Error()
	}
	appLog := AppLog{
		Id:        iotutil.GetNextSeqInt64(),
		Account:   account,
		AppKey:    appKey,
		TenantId:  tenantId,
		LogType:   logType,
		EventName: eventName,
		Details: map[string]string{
			"ip":     ip,
			"system": sys,
			"msg":    errMsg,
		},
		CreatedAt: time.Now(),
	}
	data, err := json.Marshal(appLog)
	if err != nil {
		return err
	}
	pd := &NatsPubData{
		Subject: iotconst.NATS_SUBJECT_RECORDS,
		Data:    string(data),
	}
	GetJsPublisherMgr().PushData(pd)
	iotlogger.LogHelper.Helper.Debugf("subject: %s data: %s", pd.Subject, pd.Data)
	return nil
}
