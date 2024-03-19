package service

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_model/db_statistics/model"
	"cloud_platform/iot_model/db_statistics/orm"
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"errors"
	"time"
)

type DeviceActiveSvc struct {
	Ctx context.Context
}

func (s *DeviceActiveSvc) Lists(req *proto.DeviceActiveListFilter) (*proto.DeviceActiveResponse, error) {
	//获取总计
	var total int64
	var db, ok = config.DBMap["iot_statistics"]
	if !ok {
		return nil, errors.New("数据库未初始化")
	}
	var err error
	t := orm.Use(db).TDeviceDataSum
	do := t.WithContext(context.Background()).Select(t.ActiveSum.IfNull(0).As("total")).
		Where(t.TenantId.Eq(req.TenantId), t.ProductKey.Eq(req.ProductKey))
	err = do.Scan(&total)
	if err != nil {
		return nil, err
	}

	var ret proto.DeviceActiveResponse
	ret.Data = &proto.OpenActiveEntitys{
		DeviceMonActive: &proto.Data{Total: total},
		DeviceDayActive: &proto.Data{Total: 0},
		//DeviceTodayActive: 0,
		//Device7DayActive:  0,
		DeviceActiveAll: int32(total),
	}
	nowDay := iotutil.New(time.Now()).BeginningOfDay()
	beginDay := GetLast30Day()
	nowMonth := iotutil.New(time.Now()).BeginningOfMonth()
	beginMonth := GetLast12Month()

	//查询月数据
	tMonth := orm.Use(db).TDeviceActiveMonth
	doMonth := tMonth.WithContext(context.Background()).Where(tMonth.TenantId.Eq(req.TenantId), tMonth.ProductKey.Eq(req.ProductKey),
		tMonth.DataTime.Gte(beginMonth), tMonth.DataTime.Lte(nowMonth))
	doMonth = doMonth.Order(tMonth.DataTime)
	var list []*model.TDeviceActiveMonth
	list, err = doMonth.Find()
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		data := &proto.TimeData{Time: v.DataTime.Format("2006-01"), Total: v.ActiveSum}
		ret.Data.DeviceMonActive.Data = append(ret.Data.DeviceMonActive.Data, data)
	}
	ret.Data.DeviceMonActive.Data = FillTimeData(ret.Data.DeviceMonActive.Data, 1, beginMonth, nowMonth)

	//查询日数据
	total = 0
	tDay := orm.Use(db).TDeviceActiveDay
	doDay := tDay.WithContext(context.Background()).Where(tDay.TenantId.Eq(req.TenantId), tDay.ProductKey.Eq(req.ProductKey),
		tDay.DataTime.Gte(beginDay), tDay.DataTime.Lte(nowDay))
	doDay = doDay.Order(tDay.DataTime)
	var list2 []*model.TDeviceActiveDay
	list2, err = doDay.Find()
	if err != nil {
		return nil, err
	}
	for _, v := range list2 {
		data := &proto.TimeData{Time: v.DataTime.Format("2006-01-02"), Total: v.ActiveSum}
		ret.Data.DeviceDayActive.Data = append(ret.Data.DeviceDayActive.Data, data)
		total += v.ActiveSum
	}
	ret.Data.DeviceDayActive.Data = FillTimeData(ret.Data.DeviceDayActive.Data, 0, beginDay, nowDay)
	ret.Data.DeviceDayActive.Total = total
	ret.Data.DeviceTodayActive = int32(total)

	//查询近7日统计
	total = 0
	beginDay = GetLast7Day()
	err = tDay.WithContext(context.Background()).Select(tDay.ActiveSum.Sum().IfNull(0).As("total")).Where(tDay.TenantId.Eq(req.TenantId), tDay.ProductKey.Eq(req.ProductKey),
		tDay.DataTime.Gte(beginDay), tDay.DataTime.Lte(nowDay)).Scan(&total)
	if err != nil {
		return nil, err
	}
	ret.Data.Device7DayActive = int32(total)
	return &ret, nil
}

// 获取最近12个月
func GetLast12Month() time.Time {
	t := time.Now()
	t1 := t.AddDate(0, -11, 0)
	return iotutil.New(t1).BeginningOfMonth()
}

// 获取最近30天
func GetLast30Day() time.Time {
	t := time.Now()
	t1 := t.AddDate(0, 0, -30)
	return iotutil.New(t1).BeginningOfDay()
}

// 获取最近7天
func GetLast7Day() time.Time {
	t := time.Now()
	t1 := t.AddDate(0, 0, -7)
	return iotutil.New(t1).BeginningOfDay()
}

// 注意:srcList必须是按照时间顺序升序排列
func FillTimeData(srcList []*proto.TimeData, flag int, begin, end time.Time) []*proto.TimeData {
	src := srcList
	curr := begin
	var next time.Time
	strTime := ""
	i := 0
	for curr.Before(end) || curr == end {
		if flag == 1 { //月
			strTime = curr.Format("2006-01")
			next = curr.AddDate(0, 1, 0)
		} else {
			strTime = curr.Format("2006-01-02")
			next = curr.AddDate(0, 0, 1)
		}
		if len(src) > i {
			if src[i].Time != strTime {
				obj := proto.TimeData{Time: strTime, Total: 0}
				src = append(src[:i], append([]*proto.TimeData{&obj}, src[i:]...)...)
			}
		} else {
			obj := proto.TimeData{Time: strTime, Total: 0}
			src = append(src, &obj)
		}
		curr = next
		i++
	}
	return src
}
