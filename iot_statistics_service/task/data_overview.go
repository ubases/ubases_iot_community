package task

import (
	"cloud_platform/iot_common/iotutil"
	appOrm "cloud_platform/iot_model/db_app/orm"
	deviceOrm "cloud_platform/iot_model/db_device/orm"
	openSystemOrm "cloud_platform/iot_model/db_open_system/orm"
	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/xxl-job/xxl-job-executor-go"
)

func HourDataOveriewStatistics(cxt context.Context, param *xxl.RunReq) (msg string) {
	curr := iotutil.New(time.Now()).BeginningOfHour()
	preHour := curr.Add(-1 * time.Hour)
	err := HourDataOveriew(preHour)
	if err != nil {
		return err.Error()
	}
	err = DeveloperRegisterSumStatistics()
	if err != nil {
		return err.Error()
	}
	return "AppListStatistics success"
}

func MonthDataOveriewStatistics(cxt context.Context, param *xxl.RunReq) (msg string) {
	start := iotutil.New(time.Now()).BeginningOfMonth()
	err := MonthDataOveriew(start)
	if err != nil {
		return err.Error()
	}
	return "AppListStatistics success"
}

func HistoryMonthDataOveriewActive(cxt context.Context, param *xxl.RunReq) (msg string) {
	timeList := strings.Split(param.ExecutorParams, ";")
	if len(timeList) != 2 {
		return "补历史数据需要开始和结束时间"
	}
	start, err := time.ParseInLocation("2006-01", timeList[0], time.Local)
	if err != nil {
		return "开始时间无法解析，请确保格式类似2006-01"
	}

	end, err := time.ParseInLocation("2006-01", timeList[1], time.Local)
	if err != nil {
		return "结束时间无法解析，请确保格式类似2006-01"
	}
	preHour := start
	for preHour.Before(end) || preHour == end {
		err = MonthDataOveriew(preHour)
		if err != nil {
			return err.Error()
		}
		preHour = preHour.AddDate(0, 1, 0)
	}
	return "HistoryActive Success"
}

func HistoryHourDataOveriewActive(cxt context.Context, param *xxl.RunReq) (msg string) {
	timeList := strings.Split(param.ExecutorParams, ";")
	if len(timeList) != 2 {
		return "补历史数据需要开始和结束时间"
	}
	start, err := time.ParseInLocation("2006-01-02T15", timeList[0], time.Local)
	if err != nil {
		return "开始时间无法解析，请确保格式类似2006-01-02T15"
	}

	end, err := time.ParseInLocation("2006-01-02T15", timeList[1], time.Local)
	if err != nil {
		return "结束时间无法解析，请确保格式类似2006-01-02T15"
	}
	preHour := start
	for preHour.Before(end) || preHour == end {
		err = HourDataOveriew(preHour)
		if err != nil {
			return err.Error()
		}
		preHour = preHour.Add(time.Hour)
	}
	return "HistoryActive Success"
}

func HourDataOveriew(preHour time.Time) error {
	tenantIDs, err := GetAllTenantID()
	if err != nil {
		return err
	}
	mapActive, err := GetDeviceActiveStatistics(preHour, 0)
	if err != nil {
		return err
	}
	total, err := GetDeveloperRegister(preHour, 0)
	if err != nil {
		return err
	}
	mapAppUserRegister, err := GetAppUserRegister(preHour, 0)
	if err != nil {
		return err
	}
	now := time.Now()
	all := statisticsModel.TDataOverviewHour{
		DataTime:             preHour,
		TenantId:             "",
		DeviceFaultSum:       0,
		DeveloperRegisterSum: total,
		UpdatedAt:            now,
	}
	datalist := make([]*statisticsModel.TDataOverviewHour, 0, len(tenantIDs)+1)
	for _, v := range tenantIDs {
		obj := statisticsModel.TDataOverviewHour{
			DataTime:             preHour,
			TenantId:             v,
			DeviceActiveSum:      mapActive[v],
			DeviceFaultSum:       0,
			DeveloperRegisterSum: 0,
			UserRegisterSum:      mapAppUserRegister[v],
			UpdatedAt:            now,
		}
		all.UserRegisterSum += obj.UserRegisterSum
		all.DeviceActiveSum += obj.DeviceActiveSum
		datalist = append(datalist, &obj)
	}
	datalist = append(datalist, &all)

	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	q := statisticsOrm.Use(statDB)
	err = q.Transaction(func(tx *statisticsOrm.Query) error {
		//先删除再插入
		_, err = tx.TDataOverviewHour.WithContext(context.Background()).Where(tx.TDataOverviewHour.DataTime.Eq(preHour)).Delete()
		if err != nil {
			return err
		}
		return tx.TDataOverviewHour.WithContext(context.Background()).CreateInBatches(datalist, len(datalist))
	})
	return err
}

func MonthDataOveriew(start time.Time) error {
	tenantIDs, err := GetAllTenantID()
	if err != nil {
		return err
	}
	mapActive, err := GetDeviceActiveStatistics(start, 2)
	if err != nil {
		return err
	}
	total, err := GetDeveloperRegister(start, 2)
	if err != nil {
		return err
	}
	mapAppUserRegister, err := GetAppUserRegister(start, 2)
	if err != nil {
		return err
	}
	now := time.Now()
	all := statisticsModel.TDataOverviewMonth{
		DataTime:             start,
		TenantId:             "",
		DeviceFaultSum:       0,
		DeveloperRegisterSum: total,
		UpdatedAt:            now,
	}
	datalist := make([]*statisticsModel.TDataOverviewMonth, 0, len(tenantIDs)+1)
	for _, v := range tenantIDs {
		obj := statisticsModel.TDataOverviewMonth{
			DataTime:             start,
			TenantId:             v,
			DeviceActiveSum:      mapActive[v],
			DeviceFaultSum:       0,
			DeveloperRegisterSum: 0,
			UserRegisterSum:      mapAppUserRegister[v],
			UpdatedAt:            now,
		}
		all.UserRegisterSum += obj.UserRegisterSum
		all.DeviceActiveSum += obj.DeviceActiveSum
		datalist = append(datalist, &obj)
	}
	datalist = append(datalist, &all)

	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	q := statisticsOrm.Use(statDB)
	err = q.Transaction(func(tx *statisticsOrm.Query) error {
		//先删除再插入
		_, err = tx.TDataOverviewMonth.WithContext(context.Background()).Where(tx.TDataOverviewMonth.DataTime.Eq(start)).Delete()
		if err != nil {
			return err
		}
		return tx.TDataOverviewMonth.WithContext(context.Background()).CreateInBatches(datalist, len(datalist))
	})
	return err
}

func GetAllTenantID() ([]string, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//查找购买模组数量
	var tenantId []string
	t := openSystemOrm.Use(openSystem).TOpenCompany
	err := t.WithContext(context.Background()).Pluck(t.TenantId, &tenantId)
	if err != nil {
		return nil, err
	}
	return tenantId, nil
}

// 获取租户激活数据
func GetDeviceActiveStatistics(start time.Time, flag int) (map[string]int64, error) {
	if start.Minute() != 0 || start.Second() != 0 {
		return nil, errors.New("start时间必须是整点时间")
	}
	end := getEndTime(start, flag)
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return nil, errors.New("iot_device数据库未初始化")
	}
	var datas []TenantIdData
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	err := t.WithContext(context.Background()).Select(t.TenantId, t.TenantId.Count().As("total")).
		Where(t.LastActivatedTime.Gte(start), t.LastActivatedTime.Lt(end), t.UseType.Eq(0), t.TenantId.IsNotNull()).
		Group(t.TenantId).Scan(&datas)
	if err != nil {
		return nil, err
	}
	mapTenantIdData := make(map[string]int64, len(datas))
	for _, v := range datas {
		mapTenantIdData[v.TenantId] = v.Total
	}
	return mapTenantIdData, nil
}

// 获取开发者注册数
func GetDeveloperRegister(start time.Time, flag int) (int64, error) {
	if start.Minute() != 0 || start.Second() != 0 {
		return 0, errors.New("start时间必须是整点时间")
	}
	end := getEndTime(start, flag)
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return 0, errors.New("iot_open_system数据库未初始化")
	}
	var total int64
	t := openSystemOrm.Use(openSystem).TOpenUser
	err := t.WithContext(context.Background()).Select(t.CreatedAt.Count().As("total")).Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end)).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// 获取开发者累计数
func DeveloperRegisterSumStatistics() error {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return errors.New("iot_open_system数据库未初始化")
	}
	var total int64
	t := openSystemOrm.Use(openSystem).TOpenUser
	err := t.WithContext(context.Background()).Select(t.CreatedAt.Count().IfNull(0).As("total")).Scan(&total)
	if err != nil {
		return err
	}

	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	q := statisticsOrm.Use(statDB)
	err = q.Transaction(func(tx *statisticsOrm.Query) error {
		//先删除再插入
		_, err = tx.TDeveloperSum.WithContext(context.Background()).Where(tx.TDeveloperSum.UpdatedAt.IsNotNull()).Delete()
		if err != nil {
			return err
		}
		return tx.TDeveloperSum.WithContext(context.Background()).Create(&statisticsModel.TDeveloperSum{DeveloperSum: total, UpdatedAt: time.Now()})
	})

	return err
}

// 获取租户app注册用户数
func GetAppUserRegister(start time.Time, flag int) (map[string]int64, error) {
	if start.Minute() != 0 || start.Second() != 0 {
		return nil, errors.New("start时间必须是整点时间")
	}
	end := getEndTime(start, flag)
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return nil, errors.New("iot_app数据库未初始化")
	}
	var datas []TenantIdData
	t := appOrm.Use(appBuild).TUcUser
	err := t.WithContext(context.Background()).Select(t.TenantId, t.TenantId.Count().As("total")).
		Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end),
			t.TenantId.IsNotNull()).Group(t.TenantId).Scan(&datas)
	if err != nil {
		return nil, err
	}
	mapTenantIdData := make(map[string]int64, len(datas))
	for _, v := range datas {
		mapTenantIdData[v.TenantId] = v.Total
	}
	return mapTenantIdData, nil
}

func getEndTime(curr time.Time, flag int) time.Time {
	switch flag {
	case 0: //小时级别
		return curr.Add(time.Hour)
	case 1: //天级别
		return curr.AddDate(0, 0, 1)
	case 2: //月级别
		return curr.AddDate(0, 1, 0)
	}
	return time.Time{}
}

type TenantIdData struct {
	TenantId string
	Total    int64
}
