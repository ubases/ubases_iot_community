package task

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	deviceOrm "cloud_platform/iot_model/db_device/orm"
	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"errors"
	"strings"
	"time"

	"cloud_platform/iot_statistics_service/config"
	"context"

	"github.com/xxl-job/xxl-job-executor-go"
)

type MonFaultData struct {
	ProductKey string
	Total      int64
}

type FaultTypeData struct {
	ProductKey string
	FaultType  string
	Total      int64
}

func MonthFault(ctx context.Context, param *xxl.RunReq) (msg string) {
	t := time.Now()
	if t.Hour() == 0 && iotutil.New(t).Day() == 1 {
		//每月1日0点，统计上个月的
		t = t.Add(-1 * time.Hour)
	}
	start := iotutil.New(t).BeginningOfMonth()
	end := iotutil.New(t).EndOfDay()
	if err := MonFaultStatistics(ctx, start, end); err != nil {
		iotlogger.LogHelper.Helper.Error("get device month fault error: ", err)
		return err.Error()
	}
	if err := FaultTypeStatistics(ctx, start, end); err != nil {
		iotlogger.LogHelper.Helper.Error("get device fault type error: ", err)
		return err.Error()
	}
	return "MonthFault Success"
}

func MonFaultStatistics(ctx context.Context, start, end time.Time) error {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	var datas []MonFaultData
	t := deviceOrm.Use(deviceDB).TIotDeviceFault
	if err := t.WithContext(ctx).Select(t.ProductKey, t.FaultName.Count().IfNull(0).As("total")).
		Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end)).Group(t.ProductKey).Scan(&datas); err != nil {
		return err
	}
	if len(datas) > 0 {
		objList := make([]*statisticsModel.TProductFaultMonth, 0, len(datas))
		for _, v := range datas {
			obj := statisticsModel.TProductFaultMonth{
				Id:         iotutil.GetNextSeqInt64(),
				ProductKey: v.ProductKey,
				Month:      start,
				Total:      v.Total,
				UpdatedAt:  time.Now(),
			}
			objList = append(objList, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TProductFaultMonth.WithContext(context.Background()).Where(tx.TProductFaultMonth.Month.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TProductFaultMonth.WithContext(context.Background()).CreateInBatches(objList, len(objList))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func FaultTypeStatistics(ctx context.Context, start, end time.Time) error {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	var datas []FaultTypeData
	t := deviceOrm.Use(deviceDB).TIotDeviceFault
	if err := t.WithContext(ctx).Select(t.ProductKey, t.FaultName.As("fault_type"), t.FaultName.Count().IfNull(0).As("total")).
		Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end)).Group(t.ProductKey, t.FaultName, t.FaultCode).Scan(&datas); err != nil {
		return err
	}
	if len(datas) > 0 {
		objList := make([]*statisticsModel.TProductFaultType, 0, len(datas))
		for _, v := range datas {
			obj := statisticsModel.TProductFaultType{
				Id:         iotutil.GetNextSeqInt64(),
				ProductKey: v.ProductKey,
				Month:      start,
				FaultType:  v.FaultType,
				Total:      v.Total,
				UpdatedAt:  time.Now(),
			}
			objList = append(objList, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TProductFaultType.WithContext(context.Background()).Where(tx.TProductFaultType.Month.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TProductFaultType.WithContext(context.Background()).CreateInBatches(objList, len(objList))
		})
		if err != nil {
			return err
		}

	}
	return nil
}

func HistoryFault(cxt context.Context, param *xxl.RunReq) (msg string) {
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
	if err := HistoryHourFault(iotutil.GetTodaySartTime(start), iotutil.GetTodaySartTime(end).Add(24*time.Hour)); err != nil {
		return err.Error()
	}
	iotlogger.LogHelper.Helper.Info(start, end)
	return "HistoryFault Success"
}

// 补处理历史
func HistoryHourFault(start, end time.Time) error {
	for start.Before(end) {
		iotlogger.LogHelper.Helper.Info(start, end)
		err := MonFaultStatistics(context.Background(), start, end)
		if err != nil {
			return err
		}
		err = FaultTypeStatistics(context.Background(), start, end)
		if err != nil {
			return err
		}
		start = start.Add(24 * time.Hour)
	}
	return nil
}

// 获取设备故障数
func GetDeviceFaultCount(start time.Time, flag int) (int64, error) {
	if start.Minute() != 0 || start.Second() != 0 {
		return 0, errors.New("start时间必须是整点时间")
	}
	end := getEndTime(start, flag)
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return 0, errors.New("iot_device数据库未初始化")
	}
	var total int64
	t := deviceOrm.Use(deviceDB).TIotDeviceFault
	err := t.WithContext(context.Background()).Select(t.CreatedAt.Count().As("total")).
		Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end)).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// 每小时设备故障
func GetDeviceFaultStatistics(start time.Time, flag int) (map[string]int64, error) {
	if start.Minute() != 0 || start.Second() != 0 {
		return nil, errors.New("start时间必须是整点时间")
	}
	end := getEndTime(start, flag)
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return nil, errors.New("iot_device数据库未初始化")
	}
	var datas []TenantIdData
	t := deviceOrm.Use(deviceDB).TIotDeviceFault
	err := t.WithContext(context.Background()).Select(t.TenantId, t.TenantId.Count().As("total")).
		Where(t.CreatedAt.Gte(start), t.CreatedAt.Lt(end), t.TenantId.IsNotNull()).
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
