package services

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	goerrors "go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AppTimerService struct {
	Ctx context.Context
}

func (s AppTimerService) SetContext(ctx context.Context) AppTimerService {
	s.Ctx = ctx
	return s
}

func getTimeLocation(timezone string) *time.Location {
	locStr := timezone
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	return loc
}

// AddTimer 添加云端定时
func (s AppTimerService) AddTimer(req entitys.IotDeviceTimerEntitys) (int64, error) {
	saveObj := entitys.IotDeviceTimer_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	saveObj.Enabled = 1
	var (
		userLoc     = getTimeLocation(req.Timezone) //当前用户所在时区
		localLoc, _ = time.LoadLocation("Local")    //本地时区
		localTime   = time.Now().In(localLoc)       //服务器时区的当前时间
	)

	//是否存在时间交集
	if err := hasTimeBeMixed(req.DeviceId, req.Type, saveObj, localTime); err != nil {
		return 0, err
	}
	//将客户需要执行的时间转换为服务器的时间
	runTime, err := userTimeToLocalTime(userLoc, localLoc, int32(localTime.Year()), int32(localTime.Month()), int32(localTime.Day()), req.Hour, req.Minute)
	if err != nil {
		return 0, err
	}
	var runEndTime time.Time
	if req.Type == 1 {
		runEndTime, err = userTimeToLocalTime(userLoc, localLoc, int32(localTime.Year()), int32(localTime.Month()), int32(localTime.Day()), req.EndHour, req.EndMinute)
		if err != nil {
			return 0, err
		}
		iotlogger.LogHelper.Infof("开始运行：%v，结束运行: %v，服务器当前时间：%v", iotutil.TimeFullFormat(runTime), iotutil.TimeFullFormat(runEndTime), iotutil.TimeFullFormat(localTime))
		//runEndTime = time.Date(localTime.Year(), localTime.Month(), localTime.Day(), iotutil.ToInt(saveObj.EndHour), iotutil.ToInt(saveObj.EndMinute), 0, 0, time.Local)
	} else {
		iotlogger.LogHelper.Infof("运行时间: %v, 服务器当前时间：%v", iotutil.TimeFullFormat(runTime), iotutil.TimeFullFormat(localTime))
	}
	//如果添加时间已经过去，设置为第二天
	if runTime.Before(localTime) {
		nextDay, _ := time.ParseDuration("24h")
		runTime = runTime.Add(nextDay)
		runEndTime = runEndTime.Add(nextDay)
	}
	saveObj.FirstTime = timestamppb.New(runTime)
	if saveObj.Weeks == "" {
		saveObj.DaysMode = 1
	} else {
		saveObj.DaysMode = GetDaysModeByWeek(strings.Split(saveObj.Weeks, ","))
	}
	//是否循环
	if saveObj.DaysMode == 1 { //不重复
		saveObj.Cron = ConvertSpecOnlyOne(runTime)
		if req.Type == 1 {
			saveObj.EndCron = ConvertSpecOnlyOne(runEndTime)
		}
	} else {
		//saveObj.Cron = ConvertSpec(int(saveObj.DaysMode), iotutil.ToInt(saveObj.Minute), iotutil.ToInt(saveObj.Hour), strings.Split(saveObj.Weeks, ","))
		//if req.Type == 1 {
		//	saveObj.EndCron = ConvertSpec(int(saveObj.DaysMode), iotutil.ToInt(saveObj.EndMinute), iotutil.ToInt(saveObj.EndHour), strings.Split(saveObj.Weeks, ","))
		//}
		//取运行时间的时分
		minute, hour := runTime.Minute(), runTime.Hour() //iotutil.ToInt(saveObj.Minute), iotutil.ToInt(saveObj.Hour)
		saveObj.Cron = ConvertSpec(int(saveObj.DaysMode), minute, hour, strings.Split(saveObj.Weeks, ","))
		if req.Type == 1 {
			//取运行时间的时分
			endMinute, endHour := runEndTime.Minute(), runEndTime.Hour() //iotutil.ToInt(saveObj.Minute), iotutil.ToInt(saveObj.Hour)
			saveObj.EndCron = ConvertSpec(int(saveObj.DaysMode), endMinute, endHour, strings.Split(saveObj.Weeks, ","))
		}
	}
	//global.GVA_LOG.Info(fmt.Sprintf("spec %s", saveObj.Cron))

	_, err = rpc.IotDeviceTimerService.Create(context.Background(), saveObj)
	if err != nil {
		return 0, err
	}
	return saveObj.Id, nil
}

// 用户时间转本地时区时间
func userTimeToLocalTime(userLoc, localLoc *time.Location, year, month, day, hour, minute int32) (time.Time, error) {
	formatLay := "2006-01-02 15:04:05"
	timeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", year, month, day, hour, minute)
	inRunTime, err := time.ParseInLocation(formatLay, timeStr, userLoc)
	if err != nil {
		return inRunTime, goerrors.New("", err.Error(), ioterrs.ErrUserTimeToLocalTime)
	}
	return inRunTime.In(localLoc), nil //运行时间
	//runTime := time.Date(currTime.Year(), currTime.Month(), currTime.Day(), iotutil.ToInt(saveObj.Hour), iotutil.ToInt(saveObj.Minute), 0, 0, time.Local)
}

// UpdateTimer 修改云端定时
func (s AppTimerService) UpdateTimer(req entitys.IotDeviceTimerUpdate) error {
	saveObj := entitys.IotDeviceTimer_e2pbUpdate(&req)
	// currTime := time.Now()
	var (
		userLoc     = getTimeLocation(req.Timezone) //当前用户所在时区
		localLoc, _ = time.LoadLocation("Local")    //本地时区
		localTime   = time.Now().In(localLoc)       //服务器时区的当前时间
	)
	// 针对时间段定时，添加得时候需要判断是否存在时间段交集
	if err := hasTimeBeMixed(req.DeviceId, req.Type, saveObj, localTime); err != nil {
		return err
	}

	//将客户需要执行的时间转换为服务器的时间
	runTime, err := userTimeToLocalTime(userLoc, localLoc, int32(localTime.Year()), int32(localTime.Month()), int32(localTime.Day()), req.Hour, req.Minute)
	if err != nil {
		return err
	}
	var runEndTime time.Time
	if req.Type == 1 {
		runEndTime, err = userTimeToLocalTime(userLoc, localLoc, int32(localTime.Year()), int32(localTime.Month()), int32(localTime.Day()), req.EndHour, req.EndMinute)
		if err != nil {
			return err
		}
		iotlogger.LogHelper.Infof("开始运行：%v，结束运行: %v，服务器当前时间：%v", iotutil.TimeFullFormat(runTime), iotutil.TimeFullFormat(runEndTime), iotutil.TimeFullFormat(localTime))
	} else {
		iotlogger.LogHelper.Infof("运行时间: %v, 服务器当前时间：%v", iotutil.TimeFullFormat(runTime), iotutil.TimeFullFormat(localTime))
	}
	// 如果添加时间已经过去，设置为第二天
	if runTime.Before(localTime) {
		nextDay, _ := time.ParseDuration("24h")
		runTime = runTime.Add(nextDay)
		runEndTime = runEndTime.Add(nextDay)
	}
	saveObj.FirstTime = timestamppb.New(runTime)
	if saveObj.Weeks == "" {
		saveObj.DaysMode = 1
	} else {
		saveObj.DaysMode = GetDaysModeByWeek(strings.Split(saveObj.Weeks, ","))
	}

	//是否循环
	if saveObj.DaysMode == 1 { //不重复
		saveObj.Cron = ConvertSpecOnlyOne(runTime)
		if req.Type == 1 {
			saveObj.EndCron = ConvertSpecOnlyOne(runEndTime)
		}
	} else {
		//取运行时间的时分
		minute, hour := runTime.Minute(), runTime.Hour() //iotutil.ToInt(saveObj.Minute), iotutil.ToInt(saveObj.Hour)
		saveObj.Cron = ConvertSpec(int(saveObj.DaysMode), minute, hour, strings.Split(saveObj.Weeks, ","))
		if req.Type == 1 {
			//取运行时间的时分
			endMinute, endHour := runEndTime.Minute(), runEndTime.Hour() //iotutil.ToInt(saveObj.Minute), iotutil.ToInt(saveObj.Hour)
			saveObj.EndCron = ConvertSpec(int(saveObj.DaysMode), endMinute, endHour, strings.Split(saveObj.Weeks, ","))
		}
	}
	_, err = rpc.IotDeviceTimerService.Update(context.Background(), saveObj)
	if err != nil {
		return err
	}
	return nil
}

// 是否存在时间交集
func hasTimeBeMixed(deviceId string, cronType int, saveObj *protosService.IotDeviceTimer, currTime time.Time) error {
	// 针对时间段定时，添加得时候需要判断是否存在时间段交集
	if cronType == 1 {
		res, err := rpc.IotDeviceTimerService.Lists(context.Background(), &protosService.IotDeviceTimerListRequest{
			Query: &protosService.IotDeviceTimer{
				DeviceId: deviceId,
				Enabled:  1,
			},
		})
		if err != nil {
			return err
		}
		curWeeks := saveObj.Weeks
		for i := range res.Data {
			if res.Data[i].EndCron == "" || res.Data[i].Id == saveObj.Id {
				continue
			}
			dataWeeks := res.Data[i].Weeks
			if len(dataWeeks) == 0 {
				// 一次性定时需要做单独处理
				curHour := currTime.Local().Hour()
				curMinute := currTime.Local().Minute()
				time1 := formatTime(iotutil.ToString(curHour)) + ":" + formatTime(iotutil.ToString(curMinute))
				time2 := formatTime(res.Data[i].Hour) + ":" + formatTime(res.Data[i].Minute)
				if time2 <= time1 {
					dataWeeks = iotutil.ToString(int(currTime.Local().Weekday()) + 1)
				} else {
					dataWeeks = iotutil.ToString(ConvertWeek(int(currTime.Local().Weekday())))
				}
			}
			if len(curWeeks) == 0 {
				// 一次性定时需要做单独处理
				curHour := currTime.Local().Hour()
				curMinute := currTime.Local().Minute()
				time1 := formatTime(iotutil.ToString(curHour)) + ":" + formatTime(iotutil.ToString(curMinute))
				time2 := formatTime(saveObj.Hour) + ":" + formatTime(saveObj.Minute)
				if time2 <= time1 {
					curWeeks = iotutil.ToString(int(currTime.Local().Weekday()) + 1)
				} else {
					curWeeks = iotutil.ToString(ConvertWeek(int(currTime.Local().Weekday())))
				}
			}
			startTime1 := formatTime(saveObj.Hour) + ":" + formatTime(saveObj.Minute)
			endTime1 := formatTime(saveObj.EndHour) + ":" + formatTime(saveObj.EndMinute)
			startTime2 := formatTime(res.Data[i].Hour) + ":" + formatTime(res.Data[i].Minute)
			endTime2 := formatTime(res.Data[i].EndHour) + ":" + formatTime(res.Data[i].EndMinute)
			if interTime(startTime1, endTime1, startTime2, endTime2) && interWeek(strings.Split(curWeeks, ","), strings.Split(dataWeeks, ",")) {
				return goerrors.New("", "同一时间段不允许有两个预约", ioterrs.ErrInterAppointment)
			}
		}
	}
	return nil
}

func GetDaysModeByWeek(weeks []string) (daysmode int32) {
	weekCount := len(weeks)
	daysmode = 1
	switch weekCount {
	case 0:
		daysmode = 1 //不重复
	default:
		daysmode = 0 //重复
	}
	return
}

// ConvertSpecOnlyOne 转换定时任务表达式格式（只执行一次）
func ConvertSpecOnlyOne(runTime time.Time) string {
	//如: 2020年3月11日13点27分15秒,?指的是不考虑星期几
	//    表达式：15 27 13 11 3 ? 2020
	spec := fmt.Sprintf("%d %d %d %d %d ?",
		runTime.Second(), runTime.Minute(),
		runTime.Hour(), runTime.Day(), runTime.Month())
	return spec
}

// ConvertSpec 转换定时任务表达式格式（按规律执行）
func ConvertSpec(daysmode int, minute int, hour int, weeks []string) string {
	var theWeeks string
	if len(weeks) == 7 {
		theWeeks = "*"
	} else {
		newWeeks := []string{}
		for _, v := range weeks {
			if v == "7" {
				v = "0"
			}
			newWeeks = append(newWeeks, v)
		}
		theWeeks = strings.Join(newWeeks, ",")
	}
	//0 11 12 * * 1,2
	spec := fmt.Sprintf("0 %d %d * * %s", minute, hour, theWeeks)
	return spec
}

// RemoveTimer 移除云端定时
func (s AppTimerService) RemoveTimer(id int64) error {
	_, err := rpc.IotDeviceTimerService.DeleteById(context.Background(), &protosService.IotDeviceTimer{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}

// TimerInfo 云端定时详情
func (s AppTimerService) TimerInfo(id int64) (*entitys.IotDeviceTimerVo, error) {
	timerInfo, err := rpc.IotDeviceTimerService.FindById(context.Background(), &protosService.IotDeviceTimerFilter{
		Id: id,
	})
	if err != nil {
		return &entitys.IotDeviceTimerVo{}, err
	}

	if len(timerInfo.Data) == 0 {
		return &entitys.IotDeviceTimerVo{}, nil
	}
	//翻译处理
	//langMap := controls.GetProductTslCached()

	return entitys.IotDeviceTimer_vo(timerInfo.Data[0]), nil
}

// TimerList 云端定时列表
func (s AppTimerService) TimerList(devId string, filter entitys.TimerListQuery) ([]*entitys.IotDeviceTimerVo, int64, error) {
	var (
		lang, _     = metadata.Get(s.Ctx, "lang")
		tenantId, _ = metadata.Get(s.Ctx, "tenantId")
	)
	var resultList = []*entitys.IotDeviceTimerVo{}
	timerList, err := rpc.IotDeviceTimerService.Lists(context.Background(), &protosService.IotDeviceTimerListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.IotDeviceTimer{
			DeviceId: devId,
		},
		OrderKey:  "created_at",
		OrderDesc: "desc",
	})
	if err != nil {
		return resultList, 0, err
	}
	if len(timerList.Data) == 0 {
		return resultList, 0, nil
	}
	langMap := controls.GetProductTslCached(tenantId)

	for _, item := range timerList.Data {
		timerInfo := entitys.IotDeviceTimer_vo(item).SetDesc("")
		//获取定时的产品信息
		deviceInfo := controls.GetDeviceInfoCached(item.DeviceId)
		if deviceInfo != nil {
			productKey := deviceInfo["productKey"]
			msgList := make([]string, 0)
			//编辑功能列表，拼凑定时器描述
			for _, v := range timerInfo.Functions {
				//功能名称
				langKey := fmt.Sprintf("%s_%s_%s_name", lang, productKey, v.FuncIdentifier)
				funcDesc := iotutil.MapGetStringVal(langMap[langKey], v.FuncDesc)
				funcDesc = iotutil.MapGetStringVal(funcDesc, v.FuncIdentifier)
				//功能值翻译
				langKey = fmt.Sprintf("%v_%v_%v_%v_name", lang, productKey, v.FuncIdentifier, v.FuncValue)
				funcValueDesc := v.FuncValueDesc
				if !iotutil.IsNumeric(v.FuncValueDesc) {
					funcValueDesc = iotutil.MapGetStringVal(langMap[langKey], v.FuncValue)
				}
				msgList = append(msgList, fmt.Sprintf("%v: %v", funcDesc, funcValueDesc))
			}
			timerInfo.FuncDesc = strings.Join(msgList, "、")
		}
		resultList = append(resultList, timerInfo)
	}
	return resultList, timerList.Total, nil
}

// EnabledTimer 启用云端定时
func (s AppTimerService) EnabledTimer(id int64, tType int) error {
	var (
		localLoc, _ = time.LoadLocation("Local") //本地时区
		localTime   = time.Now().In(localLoc)    //服务器时区的当前时间
	)
	timerInfo, err := rpc.IotDeviceTimerService.FindById(context.Background(), &protosService.IotDeviceTimerFilter{
		Id: id,
	})
	if err != nil {
		return err
	}
	if len(timerInfo.Data) == 0 {
		return nil
	}
	// 针对时间段定时，添加得时候需要判断是否存在时间段交集
	saveObj := timerInfo.Data[0]
	if err := hasTimeBeMixed(saveObj.DeviceId, tType, saveObj, localTime); err != nil {
		return err
	}
	res, err := rpc.IotDeviceTimerService.StartIotDeviceTimerJob(context.Background(), &protosService.IotDeviceTimerJobReq{
		Id: id,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

// DisabledTimer 禁用云端定时
func (s AppTimerService) DisabledTimer(id int64) error {
	timerInfo, err := rpc.IotDeviceTimerService.FindById(context.Background(), &protosService.IotDeviceTimerFilter{
		Id: id,
	})
	if err != nil {
		return err
	}
	if len(timerInfo.Data) == 0 {
		return nil
	}
	res, err := rpc.IotDeviceTimerService.StopIotDeviceTimerJob(context.Background(), &protosService.IotDeviceTimerJobReq{
		Id: id,
	})
	if err != nil {
		return err
	}
	if res.Code != 200 {
		return errors.New(res.Message)
	}
	return nil
}

func formatTime(t string) string {
	if len(t) == 1 {
		return "0" + t
	}
	return t
}

func interTime(startTime1, endTime1, startTime2, endTime2 string) bool {
	if startTime1 >= startTime2 && startTime1 <= endTime2 {
		return true
	} else if endTime1 >= startTime2 && endTime1 <= endTime2 {
		return true
	} else if startTime2 >= startTime1 && startTime2 <= endTime1 {
		return true
	} else if endTime2 > startTime1 && endTime2 <= endTime1 {
		return true
	}
	return false
}

func interWeek(week1, week2 []string) bool {
	checked := map[string]struct{}{}
	for i := range week1 {
		checked[week1[i]] = struct{}{}
	}
	for i := range week2 {
		if _, exist := checked[week2[i]]; exist {
			return true
		}
	}
	return false
}

func ConvertWeek(weekday int) int {
	if weekday == 0 {
		return 7
	}
	return weekday
}

func AddWeek(weekday int) int {
	return weekday + 1
}
