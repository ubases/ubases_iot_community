package execute

import (
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_common/iotlogger"
)

func SetResult(result *model.TSceneIntelligenceResult) error {
	db := iotmodel.GetDB()
	err := db.Save(result).Error
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return err
	}
	return nil
}
