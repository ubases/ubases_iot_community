package execute

import (
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
)

func WriteLog(logObj *model.TSceneIntelligenceLog) error {
	logObj.Id = iotutil.GetNextSeqInt64()
	db := iotmodel.GetDB()
	err := db.Save(logObj).Error
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return err
	}
	return nil
}
