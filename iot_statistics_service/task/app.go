package task

import (
	"cloud_platform/iot_common/iotutil"
	appOrm "cloud_platform/iot_model/db_app/orm"
	appBuildOrm "cloud_platform/iot_model/db_app_oem/orm"
	openSystemOrm "cloud_platform/iot_model/db_open_system/orm"
	statisticsModel "cloud_platform/iot_model/db_statistics/model"
	statisticsOrm "cloud_platform/iot_model/db_statistics/orm"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"errors"
	"time"

	"github.com/xxl-job/xxl-job-executor-go"
)

//APP统计表数据来源:
//-- app列表
//SELECT id,app_key,tenant_id,created_by FROM t_oem_app
//
//-- 开发者名称
//SELECT id,user_name FROM t_open_user
//
//-- app注册用户数
//SELECT app_key,SUM(register_sum) FROM t_app_user_month GROUP BY app_key
//
//-- app最新已构建版本
//SELECT app_id ,MAX(`version`) FROM  t_oem_app_build_record WHERE `status` = 3  GROUP BY app_id
//
//-- app版本
//SELECT app_id ,`version` FROM  t_oem_app_build_record  GROUP BY app_id,`version`
//
//-- app用户反馈
//SELECT app_key,COUNT(*) FROM t_uc_user_feedback GROUP BY app_key

func AppListStatistics(cxt context.Context, param *xxl.RunReq) (msg string) {
	applist, err := GetAppList()
	if err != nil {
		return err.Error()
	}
	appLastVersionmap, err := GetAppLastVersionList()
	if err != nil {
		return err.Error()
	}
	appVersionSummap, err := GetAppVersionSumList()
	if err != nil {
		return err.Error()
	}
	appFeedbackSummap, err := GetAppUserFeedback()
	if err != nil {
		return err.Error()
	}
	devAccounmap, err := GetDeveloperAccountList()
	if err != nil {
		return err.Error()
	}
	registerUserSummap, err := GetAppRegisterUserStatistics()
	if err != nil {
		return err.Error()
	}
	activeAppusermap, err := GetAppLast7DayActiveUserStatistics()
	if err != nil {
		return err.Error()
	}
	now := time.Now()
	list := make([]*statisticsModel.TPmAppData, 0, len(applist))
	for _, v := range applist {
		obj := statisticsModel.TPmAppData{AppId: v.Id, AppKey: v.AppKey, AppName: v.Name, RegisterUserSum: 0, ActiveUserSum: 0, VersionSum: 0, FeedbackSum: 0, UpdatedAt: now}
		if ver, ok := appLastVersionmap[v.Id]; ok {
			obj.LastVersion = ver
		}
		if verSum, ok := appVersionSummap[v.Id]; ok {
			obj.VersionSum = verSum
		}
		if feedbackSum, ok := appFeedbackSummap[v.AppKey]; ok {
			obj.FeedbackSum = feedbackSum
		}
		if account, ok := devAccounmap[v.CreatedBy]; ok {
			obj.DevAccount = account
		}
		if regSum, ok := registerUserSummap[v.AppKey]; ok {
			obj.RegisterUserSum = regSum
		}
		if activeUser, ok := activeAppusermap[v.AppKey]; ok {
			obj.ActiveUserSum = activeUser
		}
		list = append(list, &obj)
	}
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return "iot_statistics数据库未初始化"
	}

	q := statisticsOrm.Use(statDB)
	err = q.Transaction(func(tx *statisticsOrm.Query) error {
		//先删除再插入
		_, err = tx.TPmAppData.WithContext(context.Background()).Where(tx.TPmAppData.AppId.Gte(0)).Delete()
		if err != nil {
			return err
		}
		return tx.TPmAppData.WithContext(context.Background()).CreateInBatches(list, len(list))
	})
	if err != nil {
		return err.Error()
	}

	return "AppListStatistics success"
}

func GetAppList() ([]AppInfo, error) {
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//-- app列表
	//SELECT id,app_key,tenant_id,created_by FROM t_oem_app
	var appInfoList []AppInfo
	tOemApp := appBuildOrm.Use(appBuild).TOemApp
	err := tOemApp.WithContext(context.Background()).Select(tOemApp.Id, tOemApp.AppKey, tOemApp.Name, tOemApp.TenantId,
		tOemApp.CreatedBy).Scan(&appInfoList)
	if err != nil {
		return nil, err
	}
	return appInfoList, nil
}

func GetAppLastVersionList() (map[int64]string, error) {
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//-- app最新已构建版本
	//SELECT app_id ,`version` FROM  t_oem_app_build_record WHERE `status` = 3
	var appVersionList []AppLastVersion
	tOemApp := appBuildOrm.Use(appBuild).TOemAppBuildRecord
	//err := tOemApp.WithContext(context.Background()).Select(tOemApp.AppId, tOemApp.Version.Max().As("version")).Group(tOemApp.AppId).Scan(&appVersionList)
	err := tOemApp.WithContext(context.Background()).Select(tOemApp.AppId, tOemApp.Version).Scan(&appVersionList)
	if err != nil {
		return nil, err
	}
	mapAppVersionList := make(map[int64]string, len(appVersionList))
	for _, v := range appVersionList {
		if old, ok := mapAppVersionList[v.AppId]; ok {
			//新版本大于旧版本
			if ret, err1 := iotutil.VerCompare(v.Version, old); err1 == nil {
				if ret == 1 {
					mapAppVersionList[v.AppId] = v.Version
				}
			}
		} else {
			mapAppVersionList[v.AppId] = v.Version
		}
	}
	return mapAppVersionList, nil
}

func GetAppVersionSumList() (map[int64]int64, error) {
	appBuild, ok := config.DBMap["iot_app_build"]
	if !ok {
		return nil, errors.New("iot_app_build数据库未初始化")
	}
	//-- app版本数
	//SELECT app_id ,`version` FROM  t_oem_app_build_record
	//GROUP BY app_id,`version`
	var appVersionList []AppVersion
	t := appBuildOrm.Use(appBuild).TOemAppBuildRecord
	err := t.WithContext(context.Background()).Select(t.AppId, t.Version).Group(t.AppId, t.Version).Scan(&appVersionList)
	if err != nil {
		return nil, err
	}
	mapAppVersionSumList := make(map[int64]int64)
	for _, v := range appVersionList {
		mapAppVersionSumList[v.AppId] = mapAppVersionSumList[v.AppId] + 1
	}
	return mapAppVersionSumList, nil
}

func GetAppUserFeedback() (map[string]int64, error) {
	appBuild, ok := config.DBMap["iot_app"]
	if !ok {
		return nil, errors.New("iot_app数据库未初始化")
	}
	//-- app用户反馈
	//SELECT app_key,COUNT(*) FROM t_uc_user_feedback GROUP BY app_key
	var appFeedbackList []AppUserTotal
	t := appOrm.Use(appBuild).TUcUserFeedback
	do := t.WithContext(context.Background())
	do = do.Select(t.AppKey, t.AppKey.Count().As("total"))
	do = do.Group(t.AppKey)
	if err := do.Scan(&appFeedbackList); err != nil {
		return nil, err
	}
	mapappFeedbackList := make(map[string]int64, len(appFeedbackList))
	for _, v := range appFeedbackList {
		mapappFeedbackList[v.AppKey] = v.Total
	}
	return mapappFeedbackList, nil
}

func GetDeveloperAccountList() (map[int64]string, error) {
	openSystem, ok := config.DBMap["iot_open_system"]
	if !ok {
		return nil, errors.New("iot_open_system数据库未初始化")
	}
	//-- 开发者名称
	//SELECT id,user_name FROM t_open_user
	var devInfoList []DeveloperInfo
	t := openSystemOrm.Use(openSystem).TOpenUser
	err := t.WithContext(context.Background()).Select(t.Id, t.UserName).Scan(&devInfoList)
	if err != nil {
		return nil, err
	}
	mapAccountList := make(map[int64]string, len(devInfoList))
	for _, v := range devInfoList {
		mapAccountList[v.Id] = v.UserName
	}
	return mapAccountList, nil
}

func GetAppRegisterUserStatistics() (map[string]int64, error) {
	statDB, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("iot_statistics数据库未初始化")
	}
	//-- app注册用户数
	//SELECT app_key,SUM(register_sum) FROM t_app_user_month GROUP BY app_key
	var userList []AppUserTotal
	t := statisticsOrm.Use(statDB).TAppUserMonth
	do := t.WithContext(context.Background())
	do = do.Select(t.AppKey, t.RegisterSum.Sum().As("total"))
	do = do.Group(t.AppKey)
	if err := do.Scan(&userList); err != nil {
		return nil, err
	}
	mapUser := make(map[string]int64, len(userList))
	for _, v := range userList {
		mapUser[v.AppKey] = v.Total
	}
	return mapUser, nil
}

func GetAppLast7DayActiveUserStatistics() (map[string]int64, error) {
	appDB, ok := config.DBMap["iot_app"]
	if !ok {
		return nil, errors.New("iot_app数据库未初始化")
	}
	//-- app注册用户数
	//SELECT app_key,COUNT(DISTINCT user_id) FROM t_uc_user_operate
	//WHERE DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= DATE("opt_time") GROUP BY app_key
	begin := GetStartTime(2)
	var userList []AppUserTotal
	t := appOrm.Use(appDB).TUcUserOperate
	err := t.WithContext(context.Background()).Select(t.AppKey, t.UserId.Distinct().Count().As("total")).Where(t.AppKey.Neq(""), t.OptTime.Gte(begin)).Group(t.AppKey).Scan(&userList)
	if err != nil {
		return nil, err
	}
	mapUser := make(map[string]int64, len(userList))
	for _, v := range userList {
		mapUser[v.AppKey] = v.Total
	}
	return mapUser, nil
}

type AppInfo struct {
	Id        int64
	AppKey    string
	Name      string
	TenantId  string
	CreatedBy int64 //开发者ID
}

type DeveloperInfo struct {
	Id       int64
	UserName string
}

type AppUserTotal struct {
	AppKey string
	Total  int64
}

type AppLastVersion struct {
	AppId   int64
	Version string
}

type AppVersion struct {
	AppId   int64
	Version string
}

func GetStartTime(flag int) time.Time {
	t0 := iotutil.New(time.Now()).BeginningOfDay()
	t := t0
	switch flag {
	case 1: //今日
	case 2: //近7日
		t = t0.Add(-6 * 24 * time.Hour)
	case 3: //近30日
		t = t0.Add(-29 * 24 * time.Hour)
	case 4: //近60日
		t = t0.Add(-59 * 24 * time.Hour)
	default:
	}
	return t
}
