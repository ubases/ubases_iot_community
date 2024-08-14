package service

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_statistics/model"
	"cloud_platform/iot_model/db_statistics/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_statistics_service/config"
	"cloud_platform/iot_statistics_service/task"
	"context"
	"errors"
	"time"
)

type AppUserStatisticsSvc struct {
	Ctx context.Context
}

func (s *AppUserStatisticsSvc) GetAppUserStatistics(req *proto.AppUserStatisticsFilter) (*proto.AppUserStatisticsResponse, error) {
	db, ok := config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("数据库未初始化")
	}
	//获取注册用户总数
	var total int64 //累计
	t := orm.Use(db).TAppUserSum
	err := t.WithContext(context.Background()).Select(t.RegisterSum.IfNull(0).As("total")).Where(t.AppKey.Eq(req.GetAppKey())).Scan(&total)
	if err != nil {
		return nil, err
	}
	nowDay := iotutil.New(time.Now()).BeginningOfDay()
	beginDay := GetLast30Day()
	nowMonth := iotutil.New(time.Now()).BeginningOfMonth()
	beginMonth := GetLast12Month()
	aue := &proto.AppUserEntitys{
		AppUser:    &proto.Data{Total: total},
		ActiveUser: &proto.Data{Total: 0},
		//AppUserTodayActive: 0,
		//AppUserToday:       0,
		AppUserAll: int32(total),
	}

	aue.AppUserToday, _ = task.AppRegisterUserTodayStatistics(req.GetAppKey())

	//查询月数据
	tMonth := orm.Use(db).TAppUserMonth
	doMonth := tMonth.WithContext(context.Background())
	doMonth = doMonth.Where(tMonth.AppKey.Eq(req.AppKey))
	doMonth = doMonth.Where(tMonth.DataTime.Gte(beginMonth))
	doMonth = doMonth.Where(tMonth.DataTime.Lte(nowMonth))
	doMonth = doMonth.Order(tMonth.DataTime)
	var list []*model.TAppUserMonth
	list, err = doMonth.Find()
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		data := &proto.TimeData{Time: v.DataTime.Format("2006-01"), Total: v.RegisterSum}
		aue.AppUser.Data = append(aue.AppUser.Data, data)
	}
	aue.AppUser.Data = FillTimeData(aue.AppUser.Data, 1, beginMonth, nowMonth)

	//查询近30天的用户活跃数据，每天一条
	tappUserActiveDay := orm.Use(db).TAppUserActiveDay
	activelist, err := tappUserActiveDay.WithContext(context.Background()).Where(
		tappUserActiveDay.DataTime.Gte(beginDay),
		tappUserActiveDay.DataTime.Lte(nowDay),
		tappUserActiveDay.AppKey.Eq(req.GetAppKey())).Order(tappUserActiveDay.DataTime).Find()
	if err != nil {
		return nil, err
	}
	for _, v := range activelist {
		data := &proto.TimeData{Time: v.DataTime.Format("2006-01-02"), Total: v.ActiveSum}
		aue.ActiveUser.Data = append(aue.ActiveUser.Data, data)
		//今天激活
		if nowDay.Equal(iotutil.New(v.DataTime).BeginningOfDay()) {
			aue.AppUserTodayActive = int32(v.ActiveSum)
		}
	}

	//近30日活跃用户总数，总共一条
	tappUserActive30Day := orm.Use(db).TAppUserActive30day
	activeUser30Day, err := tappUserActive30Day.WithContext(context.Background()).Where(
		tappUserActive30Day.DataTime.Eq(nowDay),
		tappUserActive30Day.AppKey.Eq(req.GetAppKey())).First()
	if err == nil {
		aue.ActiveUser.Total = activeUser30Day.ActiveSum
	}
	aue.ActiveUser.Data = FillTimeData(aue.ActiveUser.Data, 0, beginDay, nowDay)
	ret := &proto.AppUserStatisticsResponse{
		Code:    200,
		Message: "success",
		Data:    aue,
	}
	return ret, nil
}
