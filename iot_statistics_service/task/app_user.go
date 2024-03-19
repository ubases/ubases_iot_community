package task

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	appOrm "cloud_platform/iot_model/db_app/orm"
	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/xxl-job/xxl-job-executor-go"
)

func MonthAppUser(cxt context.Context, param *xxl.RunReq) (msg string) {
	curr := iotutil.New(time.Now()).BeginningOfHour()
	preHour := curr.Add(-1 * time.Hour)
	start := iotutil.New(curr).BeginningOfMonth()
	end := iotutil.New(curr).EndOfMonth()
	err := AppRegisterUserMonthStatistics(start, end)
	if err != nil {
		return err.Error()
	}
	//每月1日0点，重新统计上个月的数据
	if preHour.Hour() == 23 && start.Day() == iotutil.New(start).EndOfMonth().Day() {
		start = iotutil.New(preHour).BeginningOfMonth()
		end = iotutil.New(preHour).EndOfMonth()
		err = AppRegisterUserMonthStatistics(start, end)
		if err != nil {
			return err.Error()
		}
	}
	err = AppRegisterUserStatistics()
	if err != nil {
		return err.Error()
	}
	return "MonthAppUser"
}

func DayAppActiveUser(cxt context.Context, param *xxl.RunReq) (msg string) {
	curr := iotutil.New(time.Now()).BeginningOfHour()
	preHour := curr.Add(-1 * time.Hour)
	start := iotutil.New(curr).BeginningOfDay()
	//统计当日
	err := AppActiveDayUserStatistics(start)
	if err != nil {
		return err.Error()
	}
	err = AppActive30DayUserStatistics(start)
	if err != nil {
		return err.Error()
	}
	//如果是0点多，再次统计昨天整天的
	if preHour.Hour() == 23 {
		err = AppActiveDayUserStatistics(preHour)
		if err != nil {
			return err.Error()
		}
		err = AppActive30DayUserStatistics(preHour)
		if err != nil {
			return err.Error()
		}
	}
	return "DayAppActiveUser"
}

func MonthHistoryAppUser(cxt context.Context, param *xxl.RunReq) (msg string) {
	timeList := strings.Split(param.ExecutorParams, ";")
	if len(timeList) != 2 {
		return "补历史数据需要开始和结束时间"
	}
	start, err := time.ParseInLocation("2006-01", timeList[0], time.Local)
	if err != nil {
		return "开始时间无法解析，请确保格式类似2006-01-02T15"
	}

	end, err := time.ParseInLocation("2006-01", timeList[1], time.Local)
	if err != nil {
		return "结束时间无法解析，请确保格式类似2006-01-02T15"
	}
	for start.Before(end) {
		startMonth := iotutil.New(start).BeginningOfMonth()
		endMonth := iotutil.New(start).EndOfMonth()
		err = AppRegisterUserMonthStatistics(startMonth, endMonth)
		if err != nil {
			iotlogger.LogHelper.Errorf("AppRegisterUserMonthStatistics start=%v,end=%v error:%v", startMonth, endMonth, err)
		}
		start = startMonth.AddDate(0, 1, 0)
	}
	//重新统计总数
	err = AppRegisterUserStatistics()
	if err != nil {
		return err.Error()
	}
	return "HistoryActive Success"
}

func AppRegisterUserMonthStatistics(start, end time.Time) error {
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return errors.New("iot_app数据库未初始化")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	//SELECT tenant_id,app_key,COUNT(*) FROM t_uc_user
	//WHERE created_at >= '2022-07-28 00:00:00' AND created_at < '2022-07-29 00:00:00'
	//GROUP BY tenant_id,app_key
	var appList []AppTotal
	tucUser := appOrm.Use(appBuild).TUcUser
	do := tucUser.WithContext(context.Background())
	do = do.Select(tucUser.AppKey, tucUser.TenantId, tucUser.AppKey.Count().IfNull(0).As("total"))
	do = do.Where(tucUser.TenantId.IsNotNull(), tucUser.TenantId.Neq(""), tucUser.AppKey.IsNotNull(), tucUser.AppKey.Neq(""),
		/*tucUser.CancelTime.Eq(0),*/ tucUser.CreatedAt.Gte(start), tucUser.CreatedAt.Lt(end)).Group(tucUser.AppKey, tucUser.TenantId)
	if err := do.Scan(&appList); err != nil {
		return err
	}
	if len(appList) > 0 {
		var err error
		now := time.Now()
		list := make([]*statisticsModel.TAppUserMonth, 0, len(appList))
		for _, v := range appList {
			obj := statisticsModel.TAppUserMonth{DataTime: start, TenantId: v.TenantId, AppKey: v.AppKey, RegisterSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TAppUserMonth.WithContext(context.Background()).Where(tx.TAppUserMonth.DataTime.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TAppUserMonth.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func AppRegisterUserStatistics() error {
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return errors.New("iot_app数据库未初始化")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	//SELECT tenant_id,app_key,COUNT(*) FROM t_uc_user
	//GROUP BY tenant_id,app_key
	var appList []AppTotal
	tucUser := appOrm.Use(appBuild).TUcUser
	do := tucUser.WithContext(context.Background())
	do = do.Select(tucUser.AppKey, tucUser.TenantId, tucUser.AppKey.Count().IfNull(0).As("total"))
	do = do.Where(tucUser.TenantId.IsNotNull(), tucUser.TenantId.Neq(""), tucUser.AppKey.IsNotNull(),
		tucUser.AppKey.Neq("")).Group(tucUser.AppKey, tucUser.TenantId)
	if err := do.Scan(&appList); err != nil {
		return err
	}
	if len(appList) > 0 {
		var err error
		now := time.Now()
		list := make([]*statisticsModel.TAppUserSum, 0, len(appList))
		for _, v := range appList {
			obj := statisticsModel.TAppUserSum{TenantId: v.TenantId, AppKey: v.AppKey, RegisterSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TAppUserSum.WithContext(context.Background()).Where(tx.TAppUserSum.UpdatedAt.IsNotNull()).Delete()
			if err != nil {
				return err
			}
			return tx.TAppUserSum.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 统计某日活跃用户数
func AppActiveDayUserStatistics(start time.Time) error {
	if start.Minute() != 0 || start.Second() != 0 {
		return errors.New("AppActiveDayUserStatistics参数错误")
	}
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return errors.New("iot_app数据库未初始化")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	//SELECT tenant_id,app_key,COUNT(DISTINCT user_id)
	//FROM t_uc_user_operate WHERE opt_time >= '2022-09-01 00:00:00' AND opt_time < '2022-09-02 00:00:00'
	//GROUP BY tenant_id,app_key
	end := start.AddDate(0, 0, 1)
	var appList []AppTotal
	tucOpt := appOrm.Use(appBuild).TUcUserOperate
	do := tucOpt.WithContext(context.Background())
	do = do.Select(tucOpt.AppKey, tucOpt.TenantId, tucOpt.UserId.Distinct().Count().IfNull(0).As("total"))
	do = do.Where(tucOpt.TenantId.IsNotNull(), tucOpt.TenantId.Neq(""), tucOpt.AppKey.IsNotNull(),
		tucOpt.AppKey.Neq(""), tucOpt.OptTime.Gte(start), tucOpt.OptTime.Lt(end)).Group(tucOpt.AppKey, tucOpt.TenantId)
	if err := do.Scan(&appList); err != nil {
		return err
	}
	if len(appList) > 0 {
		var err error
		now := time.Now()
		list := make([]*statisticsModel.TAppUserActiveDay, 0, len(appList))
		for _, v := range appList {
			obj := statisticsModel.TAppUserActiveDay{DataTime: start, TenantId: v.TenantId, AppKey: v.AppKey, ActiveSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TAppUserActiveDay.WithContext(context.Background()).Where(tx.TAppUserActiveDay.DataTime.Eq(start)).Delete()
			if err != nil {
				return err
			}
			return tx.TAppUserActiveDay.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 统计最近30天活跃用户数
func AppActive30DayUserStatistics(end time.Time) error {
	if end.Minute() != 0 || end.Second() != 0 {
		return errors.New("AppActive30DayUserStatistics参数错误")
	}
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return errors.New("iot_app数据库未初始化")
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return errors.New("iot_statistics数据库未初始化")
	}
	//SELECT tenant_id,app_key,COUNT(DISTINCT user_id)
	//FROM t_uc_user_operate WHERE opt_time >= '2022-09-01 00:00:00' AND opt_time < '2022-09-02 00:00:00'
	//GROUP BY tenant_id,app_key
	start := end.AddDate(0, 0, -29)
	endnow := end.AddDate(0, 0, 1)
	var appList []AppTotal
	tucOpt := appOrm.Use(appBuild).TUcUserOperate
	do := tucOpt.WithContext(context.Background())
	do = do.Select(tucOpt.AppKey, tucOpt.TenantId, tucOpt.UserId.Distinct().Count().IfNull(0).As("total"))
	do = do.Where(tucOpt.TenantId.IsNotNull(), tucOpt.TenantId.Neq(""), tucOpt.AppKey.IsNotNull(),
		tucOpt.AppKey.Neq(""), tucOpt.OptTime.Gte(start), tucOpt.OptTime.Lt(endnow)).Group(tucOpt.AppKey, tucOpt.TenantId)
	if err := do.Scan(&appList); err != nil {
		return err
	}
	if len(appList) > 0 {
		var err error
		now := time.Now()
		list := make([]*statisticsModel.TAppUserActive30day, 0, len(appList))
		for _, v := range appList {
			obj := statisticsModel.TAppUserActive30day{DataTime: end, TenantId: v.TenantId, AppKey: v.AppKey, ActiveSum: v.Total, UpdatedAt: now}
			list = append(list, &obj)
		}
		q := statisticsOrm.Use(statDB)
		err = q.Transaction(func(tx *statisticsOrm.Query) error {
			//先删除再插入
			_, err = tx.TAppUserActive30day.WithContext(context.Background()).Where(tx.TAppUserActive30day.DataTime.Eq(end)).Delete()
			if err != nil {
				return err
			}
			return tx.TAppUserActive30day.WithContext(context.Background()).CreateInBatches(list, len(list))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type AppTotal struct {
	AppKey   string
	TenantId string
	Total    int64
}

func AppRegisterUserTodayStatistics(appKey string) (int32, error) {
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return 0, errors.New("iot_app数据库未初始化")
	}
	//SELECT COUNT(*) FROM t_uc_user WHERE app_key = ''
	var total int64
	t := appOrm.Use(appBuild).TUcUser
	err := t.WithContext(context.Background()).Select(t.Id.Count().As("total")).
		Where(t.TenantId.IsNotNull(), t.TenantId.Neq(""), t.AppKey.Eq(appKey),
			t.CreatedAt.Between(iotutil.BeginningOfDay(), iotutil.EndOfDay())).Scan(&total)
	if err != nil {
		return 0, err
	}
	return int32(total), nil
}
