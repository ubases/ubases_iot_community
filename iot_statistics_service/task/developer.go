package task

import (
	appBuildOrm "cloud_platform/iot_model/db_app_oem/orm"
	deviceOrm "cloud_platform/iot_model/db_device/orm"
	openSystemOrm "cloud_platform/iot_model/db_open_system/orm"
	"errors"
	"time"

	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"cloud_platform/iot_statistics_service/config"

	"context"

	"github.com/xxl-job/xxl-job-executor-go"
)

type TenantTotal struct {
	TenantId string
	Total    int64
}

//-- 购买模组数量
//SELECT tenant_id,SUM(auth_quantity) FROM t_open_auth_quantity WHERE  tenant_id IS NOT NULL GROUP BY tenant_id

//-- 已激活数量
//SELECT tenant_id,COUNT(*) FROM t_iot_device_info WHERE active_status = 1 AND tenant_id IS NOT NULL GROUP BY tenant_id

//-- 已开发app
//SELECT tenant_id,COUNT(*) FROM  t_oem_app WHERE  tenant_id IS NOT NULL GROUP BY tenant_id

func DeveloperStatistics(cxt context.Context, param *xxl.RunReq) (msg string) {
	tenantIDs, err := getAllDeveloper()
	if err != nil {
		return err.Error()
	}
	mapauthList, err := AuthQuantityStatistics()
	if err != nil {
		return err.Error()
	}
	mapactiveDeviceList, err := DeviceStatistics()
	if err != nil {
		return err.Error()
	}
	mapappList, err := AppStatistics()
	if err != nil {
		return err.Error()
	}
	//删除历史统计
	err = ClearDeveloperStatistics()
	if err != nil {
		return err.Error()
	}

	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return "iot_statistics数据库未初始化"
	}

	now := time.Now()
	var list []*statisticsModel.TPmDevelopData
	for _, k := range tenantIDs {
		deviceSum, _ := mapauthList[k]
		actives, _ := mapactiveDeviceList[k]
		apps, _ := mapappList[k]
		obj := statisticsModel.TPmDevelopData{
			TenantId:        k,
			DeviceSum:       deviceSum,
			DeviceActiveSum: actives,
			AppSum:          apps,
			UpdatedAt:       now,
		}
		list = append(list, &obj)
	}
	if len(list) > 0 {
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TPmDevelopData.WithContext(context.Background()).Where(tx.TPmDevelopData.UpdatedAt.IsNotNull()).Delete()
			if err != nil {
				return err
			}
			return tx.TPmDevelopData.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
	}
	return "MonthDeveloper"
}

func ClearDeveloperStatistics() error {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	tdeveloper := statisticsOrm.Use(statDB).TPmDevelopData
	_, err := tdeveloper.WithContext(context.Background()).Where(tdeveloper.UpdatedAt.IsNotNull()).Delete()
	return err
}

func AppStatistics() (map[string]int64, error) {
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//查找开发app数量
	var appList []TenantTotal
	tOemApp := appBuildOrm.Use(appBuild).TOemApp
	err := tOemApp.WithContext(context.Background()).Select(tOemApp.TenantId, tOemApp.TenantId.Count().As("total")).
		Where(tOemApp.TenantId.IsNotNull()).Group(tOemApp.TenantId).Scan(&appList)
	if err != nil {
		return nil, err
	}
	mapappList := make(map[string]int64, len(appList))
	for _, v := range appList {
		mapappList[v.TenantId] = v.Total
	}
	return mapappList, nil
}

func DeviceStatistics() (map[string]int64, error) {
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return nil, errors.New("iot_device数据库未初始化")
	}
	//查找激活设备数量
	var activeDeviceList []TenantTotal
	tDeviceInfo := deviceOrm.Use(deviceDB).TIotDeviceInfo
	err := tDeviceInfo.WithContext(context.Background()).Select(tDeviceInfo.TenantId, tDeviceInfo.TenantId.Count().As("total")).
		Where(tDeviceInfo.TenantId.IsNotNull() /*tDeviceInfo.UseType.Eq(0)*/).Group(tDeviceInfo.TenantId).Scan(&activeDeviceList)
	if err != nil {
		return nil, err
	}
	mapactiveDeviceList := make(map[string]int64, len(activeDeviceList))
	for _, v := range activeDeviceList {
		mapactiveDeviceList[v.TenantId] = v.Total
	}
	return mapactiveDeviceList, nil
}

func AuthQuantityStatistics() (map[string]int64, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//查找购买模组数量
	var authList []TenantTotal
	t := openSystemOrm.Use(openSystem).TOpenAuthQuantity
	err := t.WithContext(context.Background()).Select(t.TenantId, t.AuthQuantity.Sum().As("total")).
		Where(t.TenantId.IsNotNull()).Group(t.TenantId).Scan(&authList)
	if err != nil {
		return nil, err
	}
	mapauthList := make(map[string]int64, len(authList))
	for _, v := range authList {
		mapauthList[v.TenantId] = v.Total
	}
	return mapauthList, nil
}

func getAllDeveloper() ([]string, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//查找购买模组数量
	var tenantIds []string
	t := openSystemOrm.Use(openSystem).TOpenCompany
	err := t.WithContext(context.Background()).Select(t.TenantId).Where(t.TenantId.IsNotNull()).Group(t.TenantId).Scan(&tenantIds)
	if err != nil {
		return nil, err
	}
	return tenantIds, nil
}
