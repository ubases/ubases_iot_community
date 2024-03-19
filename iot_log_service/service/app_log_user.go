package service

import (
	"cloud_platform/iot_common/iotutil"
	model "cloud_platform/iot_model"
	models "cloud_platform/iot_model/ch_log/model"
	"time"
)

func CreateAppLogUser(al *models.AppLogUser) error {
	if err := model.GetDB().Debug().Model(&models.AppLogUser{}).Create(al).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAppLogUser(al *models.AppLogUser) error {
	data := map[string]interface{}{
		"app_name":   al.AppName,
		"region":     al.Region,
		"login_time": time.Now(),
	}
	if err := model.GetDB().Debug().Model(&models.AppLogUser{}).Where("account = ? and tenant_id = ? and app_key = ?", al.Account, al.TenantId, al.AppKey).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAppLogUser(al *models.AppLogUser) error {
	var alr []models.AppLogUser
	if err := model.GetDB().Debug().Model(&models.AppLogUser{}).Delete(&alr, "account = ? and tenant_id = ? and app_key = ?", al.Account, al.TenantId, al.AppKey).Error; err != nil {
		return err
	}
	return nil
}

func GetAppLogUser(al *models.AppLogUser) (*models.AppLogUser, error) {
	if err := model.GetDB().Debug().Model(&models.AppLogUser{}).Find(al, "region_server_id = ? and tenant_id = ? and app_key = ? and account = ?", al.RegionServerId, al.TenantId, al.AppKey, al.Account).Error; err != nil {
		return al, err
	}
	return al, nil
}

func GetAppLogUserList(account, appName string, pageNum, pageSize, regionServerId int64) (als []models.AppLogUser, total int64, err error) {
	condition := ""
	if regionServerId != 0 {
		condition = "region_server_id = " + iotutil.ToString(regionServerId)
	}
	if appName != "" {
		if condition == "" {
			condition = "app_name like '%" + appName + "%'"
		} else {
			condition += " and app_name like '%" + appName + "%'"
		}
	}
	if account != "" {
		if condition == "" {
			condition = "account like '%" + account + "%'"
		} else {
			condition += " and account like '%" + account + "%'"
		}
	}
	if err := model.GetDB().Model(&models.AppLogUser{}).Where(condition).Order("login_time desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&als).Error; err != nil {
		return als, total, err
	}
	if err := model.GetDB().Model(&models.AppLogUser{}).Where(condition).Count(&total).Error; err != nil {
		return als, total, err
	}
	return als, total, nil
}
