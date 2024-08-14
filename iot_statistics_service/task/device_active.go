package task

import (
	"cloud_platform/iot_common/iotutil"
	deviceOrm "cloud_platform/iot_model/db_device/orm"
	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"errors"
	"fmt"
	"strings"

	//deviceModel "cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"time"

	"github.com/xxl-job/xxl-job-executor-go"
)

type ActiveData struct {
	TenantId   string
	ProductKey string
	Total      int64
}

// 小时统计激活数,每次统计前一个小时的数据，整点10~50分执行
// 如果当前是0点，除了统计前一个小时的外，需要统计前一日的数据
// 如果当前是1日0点，除了统计前一日的数据外，需要统计前一个月的数据
// 注意规则: 日统计依赖小时统计，月统计依赖日统计。日和月统计不要基于原始表统计。
func HourActive(cxt context.Context, param *xxl.RunReq) (msg string) {
	curr := iotutil.New(time.Now()).BeginningOfHour()
	preHour := curr.Add(-1 * time.Hour)
	err := hourActiveStatistics(cxt, preHour)
	if err != nil {
		return err.Error()
	}
	fmt.Println(curr.Day())
	//如果当前统计23点的数据，则统计完后，执行日统计
	if preHour.Hour() == 23 {
		err = DayActiveStatistics(cxt, iotutil.New(preHour).BeginningOfDay())
		if err != nil {
			return err.Error()
		}
		//统计本月
		err = MonthActiveStatistics(cxt, iotutil.New(preHour).BeginningOfMonth())
		if err != nil {
			return err.Error()
		}
	}

	err = DeviceDataSum()
	if err != nil {
		return err.Error()
	}
	return "HourActive Success"
}

func HistoryActive(cxt context.Context, param *xxl.RunReq) (msg string) {
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
	HistoryHourActive(start, end)
	return "HistoryActive Success"
}

// 补处理历史
func HistoryHourActive(start, end time.Time) {
	for start.Before(end) {
		err := hourActiveStatistics(context.Background(), start)
		if err != nil {
			fmt.Println(err)
			return
		}
		//如果当前统计23点的数据，则统计完后，执行日统计
		if start.Hour() == 23 {
			err = DayActiveStatistics(context.Background(), iotutil.New(start).BeginningOfDay())
			if err != nil {
				fmt.Println(err)
				return
			}
			//补数据不需要每个小时统计月数据，这里特殊处理，只月末统计
			if start.Day() == iotutil.New(start).EndOfMonth().Day() {
				err = MonthActiveStatistics(context.Background(), iotutil.New(start).BeginningOfMonth())
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
		start = start.Add(time.Hour)
	}
	_ = DeviceDataSum()
}

func hourActiveStatistics(cxt context.Context, start time.Time) error {
	if start.Minute() != 0 || start.Second() != 0 {
		return errors.New("start时间必须是整点时间")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	//例如，统计2022-07-01的0点的数据
	//SELECT tenant_id,product_key,COUNT(*) AS total FROM t_iot_device_info
	//WHERE last_activated_time >= '2022-07-01 00:00:00' AND last_activated_time < '2022-07-01 01:00:00' AND tenant_id IS NOT NULL
	//GROUP BY tenant_id,product_key
	var datas []ActiveData
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	if err := t.WithContext(cxt).Select(t.TenantId, t.ProductKey, t.TenantId.Count().As("total")).
		Where(t.LastActivatedTime.Gte(start), t.LastActivatedTime.Lt(start.Add(time.Hour)) /*t.UseType.Eq(0), */, t.TenantId.IsNotNull()).
		Group(t.TenantId, t.ProductKey).Scan(&datas); err != nil {
		return err
	}
	if len(datas) > 0 {
		now := time.Now()
		list := make([]*statisticsModel.TDeviceActiveHour, 0, len(datas))
		for _, v := range datas {
			obj := statisticsModel.TDeviceActiveHour{DataTime: start, TenantId: v.TenantId, ProductKey: v.ProductKey,
				ActiveSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TDeviceActiveHour.WithContext(context.Background()).Where(tx.TDeviceActiveHour.DataTime.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TDeviceActiveHour.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func DayActiveStatistics(cxt context.Context, start time.Time) error {
	if start.Hour() != 0 || start.Minute() != 0 || start.Second() != 0 {
		return errors.New("start时间必须是0点整")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	var datas []ActiveData
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	if err := t.WithContext(cxt).Select(t.TenantId, t.ProductKey, t.TenantId.Count().As("total")).
		Where(t.LastActivatedTime.Gte(start), t.LastActivatedTime.Lt(start.AddDate(0, 0, 1)) /*t.UseType.Eq(0),*/, t.TenantId.IsNotNull()).
		Group(t.TenantId, t.ProductKey).Scan(&datas); err != nil {
		return err
	}
	if len(datas) > 0 {
		now := time.Now()
		list := make([]*statisticsModel.TDeviceActiveDay, 0, len(datas))
		for _, v := range datas {
			obj := statisticsModel.TDeviceActiveDay{DataTime: start, TenantId: v.TenantId, ProductKey: v.ProductKey,
				ActiveSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TDeviceActiveDay.WithContext(context.Background()).Where(tx.TDeviceActiveDay.DataTime.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TDeviceActiveDay.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func MonthActiveStatistics(cxt context.Context, start time.Time) error {
	if start.Hour() != 0 || start.Minute() != 0 || start.Second() != 0 {
		return errors.New("start时间必须是月初一0点整")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}

	//end := iotutil.New(start).EndOfMonth()
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	var datas []ActiveData
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	if err := t.WithContext(cxt).Select(t.TenantId, t.ProductKey, t.TenantId.Count().As("total")).
		Where(t.LastActivatedTime.Gte(start), t.LastActivatedTime.Lt(start.AddDate(0, 1, 0)) /*t.UseType.Eq(0),*/, t.TenantId.IsNotNull()).
		Group(t.TenantId, t.ProductKey).Scan(&datas); err != nil {
		return err
	}
	if len(datas) > 0 {
		now := time.Now()
		list := make([]*statisticsModel.TDeviceActiveMonth, 0, len(datas))
		for _, v := range datas {
			obj := statisticsModel.TDeviceActiveMonth{DataTime: start, TenantId: v.TenantId, ProductKey: v.ProductKey,
				ActiveSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TDeviceActiveMonth.WithContext(context.Background()).Where(tx.TDeviceActiveMonth.DataTime.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TDeviceActiveMonth.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func DeviceDataSum() error {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	deviceDB, ok := config.DBMap["iot_device"]
	if !ok {
		return errors.New("iot_device数据库未初始化")
	}
	var datas []ActiveData
	t := deviceOrm.Use(deviceDB).TIotDeviceInfo
	if err := t.WithContext(context.Background()).Select(t.TenantId, t.ProductKey, t.Id.Count().As("total")).
		Where( /*t.UseType.Eq(0),*/ t.TenantId.IsNotNull()).
		Group(t.TenantId, t.ProductKey).Scan(&datas); err != nil {
		return err
	}

	var faultData []ActiveData
	tf := deviceOrm.Use(deviceDB).TIotDeviceFault
	if err := tf.WithContext(context.Background()).Select(tf.TenantId, tf.ProductKey, tf.Id.Count().As("total")).
		Where(tf.TenantId.IsNotNull()).
		Group(tf.TenantId, tf.ProductKey).Scan(&faultData); err != nil {
		return err
	}
	mapFaultData := make(map[string]int64)
	for _, v := range faultData {
		mapFaultData[v.TenantId+v.ProductKey] = v.Total
	}
	//fixme 统计设备故障数据
	if len(datas) > 0 {
		now := time.Now()
		list := make([]*statisticsModel.TDeviceDataSum, 0, len(datas))
		for _, v := range datas {
			var faultTotal int64
			if count, ok := mapFaultData[v.TenantId+v.ProductKey]; ok {
				faultTotal = count
			}
			obj := statisticsModel.TDeviceDataSum{TenantId: v.TenantId, ProductKey: v.ProductKey, ActiveSum: v.Total, FaultSum: faultTotal, UpdatedAt: now}
			list = append(list, &obj)
		}
		var err error
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TDeviceDataSum.WithContext(context.Background()).Where(tx.TDeviceDataSum.UpdatedAt.IsNotNull()).Delete()
			if err != nil {
				return err
			}
			return tx.TDeviceDataSum.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}
