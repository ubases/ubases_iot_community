package service

import (
	"cloud_platform/iot_common/iotutil"
	model "cloud_platform/iot_model"
	models "cloud_platform/iot_model/ch_log/model"
	"time"
)

func CreateAppLogRecords(alr []models.AppLogRecords) error {
	if err := model.GetDB().Model(&models.AppLogRecords{}).Create(alr).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAppLogRecords(al *models.AppLogRecords) error {
	var alr []models.AppLogRecords
	if err := model.GetDB().Debug().Model(&models.AppLogRecords{}).Delete(&alr, "account = ? and tenant_id = ? and app_key = ?", al.Account, al.TenantId, al.AppKey).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAppLogRecordsByTime(delTime time.Time) error {
	var alr []models.AppLogRecords
	if err := model.GetDB().Debug().Model(&models.AppLogRecords{}).Delete(&alr, "created_at < ?", delTime).Error; err != nil {
		return err
	}
	return nil
}

func GetAppLogRecordsList(account, appKey, tenantId, startTime, endTime, eventName, logType string, pageNum, pageSize, regionServerId int64) (alr []models.AppLogRecords, total int64, err error) {
	condition := ""
	if regionServerId != 0 {
		condition = "region_server_id = " + iotutil.ToString(regionServerId)
	}
	if tenantId != "" {
		if condition == "" {
			condition = "tenant_id = '" + tenantId + "'"
		} else {
			condition += " and tenant_id = '" + tenantId + "'"
		}
	}
	if appKey != "" {
		if condition == "" {
			condition = "app_key = '" + appKey + "'"
		} else {
			condition += " and app_key = '" + appKey + "'"
		}
	}
	if account != "" {
		if condition == "" {
			condition = "account = '" + account + "'"
		} else {
			condition += " and account = '" + account + "'"
		}
	}
	if logType != "" {
		if condition == "" {
			condition = "log_type = '" + logType + "'"
		} else {
			condition += " and log_type = '" + logType + "'"
		}
	}
	if eventName != "" {
		if condition == "" {
			condition = "event_name = '" + eventName + "'"
		} else {
			condition += " and event_name = '" + eventName + "'"
		}
	}
	if startTime != "" {
		if condition == "" {
			condition = "created_at >= '" + startTime + "'"
		} else {
			condition += " and created_at >= '" + startTime + "'"
		}
	}
	if endTime != "" {
		if condition == "" {
			condition = "created_at < '" + endTime + "'"
		} else {
			condition += " and created_at < '" + endTime + "'"
		}
	}
	if err := model.GetDB().Model(&models.AppLogRecords{}).Where(condition).Order("created_at desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&alr).Error; err != nil {
		return alr, total, err
	}
	if err := model.GetDB().Model(&models.AppLogRecords{}).Where(condition).Count(&total).Error; err != nil {
		return alr, total, err
	}
	return alr, total, nil
}
