package service

import (
	"cloud_platform/iot_common/iotnatsjs"
	"cloud_platform/iot_common/iotstruct"

	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
)

func pushOtaPublishLog(log iotstruct.OtaPublishLog) error {
	defer iotutil.PanicHandler(log)
	data, err := json.Marshal(log)
	if err != nil {
		return err
	}
	pd := &iotnatsjs.NatsPubData{
		Subject: iotconst.NATS_SUBJECT_PRODUCT_PUBLISH,
		Data:    string(data),
	}
	iotnatsjs.GetJsClientPub().PushData(pd)
	iotlogger.LogHelper.Helper.Debugf("pushOtaPublishLog subject: %s data: %s", pd.Subject, pd.Data)
	return nil
}
