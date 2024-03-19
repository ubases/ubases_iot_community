package scene_executor

import (
	cron2 "cloud_platform/iot_intelligence_service/service/scene_executor/cron"
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	goerrors "go-micro.dev/v4/errors"

	"github.com/robfig/cron/v3"
)

// ObserverTimerItems 定时器管理
type ObserverTimerItems struct {
	observerHasMap sync.Map
}

func (s *ObserverTimerItems) initSub() {
	s.observerHasMap = sync.Map{}
}

func (s *ObserverTimerItems) register(o observerTimer) (bool, error) {
	cronStr := s.generateCron(o)
	if cronStr == "" {
		return false, errors.New("执行表达式生成异常")
	}
	iotlogger.LogHelper.WithTag("Intelligence", "Condition.Timer.register").
		Infof("定时器表达式为，cron:%v", cronStr)
	//创建定时器 "0 */15 * * * *"
	entryId, err := cron2.CronCtx.AddFunc(cronStr, func() {
		s.notify(o.getRuleId())
	})
	if err != nil {
		return false, err
	}
	o.setEntryId(entryId)
	s.observerHasMap.Store(o.getRuleId(), o)
	return true, nil
}
func (s *ObserverTimerItems) deregister(o observerTimer) (bool, error) {
	s.removeFormSlice(o)
	return true, nil
}
func (s *ObserverTimerItems) removeFormSlice(o observerTimer) {
	mapKey, _ := s.observerHasMap.Load(o.getEntryId())
	ob := mapKey.(observerTimer)
	cron2.CronCtx.Remove(ob.getEntryId())
	s.observerHasMap.Delete(o.getEntryId())
}
func (s *ObserverTimerItems) generateCron(o observerTimer) string {
	return convertSpec(o.getTimer())
}
func (s *ObserverTimerItems) notifyAll() {
	//天气发送变化通知所有观察者
	s.observerHasMap.Range(func(key, value interface{}) bool {
		o := value.(observerTimer)
		if o != nil {
			o.run()
		}
		return true
	})
}
func (s *ObserverTimerItems) notify(id string) {
	mapKey, _ := s.observerHasMap.Load(id)
	ob := mapKey.(observerTimer)
	ob.run()
}

// 定时器变化
type observerTimer interface {
	run() bool
	getEntryId() cron.EntryID
	setEntryId(entryId cron.EntryID)
	getRuleId() string
	getTimer() (string, string, string)
	getTimezone() string
	getRegionServerId() int64
}

// TimerObserver 定时器观察者
type TimerObserver struct {
	id             string
	entryId        cron.EntryID
	weekVal        string //周值（1-7） 1,2,3,4,5,6,7， 注意：cron中需要转换为 7 需要转换为 0
	timeVal        string //时间值（几点:几分） 01:00
	timezone       string //时区
	regionServerId int64  //区域
}

func (w TimerObserver) run() bool {
	if valscene.Gengine == nil {
		return false
	}
	if valscene.TimerRuleBuilder == nil {
		return false
	}
	//当前任务为执行一次，还是重复执行
	if w.weekVal == "" {
		cron2.CronCtx.Remove(w.entryId)
	}
	err := valscene.Gengine.ExecuteSelectedRules(valscene.TimerRuleBuilder, []string{w.getRuleId()})
	if err != nil {
		return false
	}
	return true
}

// 获取规则Id，这个Id为场景Id
func (w TimerObserver) getRuleId() string {
	return w.id
}

// 设置cron执行Id
func (w TimerObserver) setEntryId(entryId cron.EntryID) {
	w.entryId = entryId
}

// 获取cron执行Id
func (w TimerObserver) getEntryId() cron.EntryID {
	return w.entryId
}

// 获取执行时间
func (w TimerObserver) getTimer() (string, string, string) {
	return w.timeVal, w.weekVal, w.timezone
}

// 获取区域服务器Id
func (w TimerObserver) getRegionServerId() int64 {
	return w.regionServerId
}

// 获取时区
func (w TimerObserver) getTimezone() string {
	return w.timezone
}

// 转换定时任务表达式格式（按规律执行）
// 0 0 9 * * 1,2,3,4,5,6,7
func convertSpec(timeV string, weeks string, timezone string) string {
	var (
		userLoc     = getTimeLocation(timezone)  //当前用户所在时区
		localLoc, _ = time.LoadLocation("Local") //本地时区
		localTime   = time.Now().In(localLoc)    //服务器时区的当前时间
	)

	if timeV == "" {
		return ""
	}
	values := strings.Split(timeV, ":")
	hourStr := values[0]
	minuteStr := values[1]
	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		iotlogger.LogHelper.Error("1-时间转换错误：", hourStr, err)
		return ""
	}
	minute, err := strconv.Atoi(minuteStr)
	if err != nil {
		iotlogger.LogHelper.Error("2-时间转换错误：", minuteStr, err)
		return ""
	}

	//将客户需要执行的时间转换为服务器的时间
	runTime, err := userTimeToLocalTime(userLoc, localLoc, int32(localTime.Year()), int32(localTime.Month()), int32(localTime.Day()), int32(hour), int32(minute))
	if err != nil {
		iotlogger.LogHelper.Errorf("3-时间转换错误：timeV:%s, err: %v", timeV, err.Error())
		return ""
	}
	fmt.Println("runTime：", iotutil.TimeFullFormat(runTime))

	//如果添加时间已经过去，设置为第二天
	if runTime.Before(localTime) {
		nextDay, _ := time.ParseDuration("24h")
		runTime = runTime.Add(nextDay)
	}
	if weeks == "" {
		minute, hour := runTime.Minute(), runTime.Hour()
		return convertSpecOnlyOne(minute, hour)
	}
	theWeeks := strings.Replace(weeks, "7", "0", 1)
	//当前秒 11 12 * * 1,2
	//spec := fmt.Sprintf("%d %d %d * * %s", localTime.Second(), minute, hour, theWeeks)
	spec := fmt.Sprintf("0 %d %d * * %s", runTime.Minute(), runTime.Hour(), theWeeks)
	return spec
}

// 转换定时任务表达式格式（只执行一次）
func convertSpecOnlyOne(minute int, hour int) string {
	curTime := time.Now()
	//如: 2020年3月11日13点27分15秒,?指的是不考虑星期几
	//    表达式：15 27 13 11 3 ? 2020
	//spec := fmt.Sprintf("%d %d %d %d %d ?",
	//	curTime.Second(), minute,
	//	hour, curTime.Day(), curTime.Month())
	spec := fmt.Sprintf("0 %d %d %d %d ?", minute, hour, curTime.Day(), curTime.Month())
	return spec
}

func userTimeToLocalTime(userLoc, localLoc *time.Location, year, month, day, hour, minute int32) (time.Time, error) {
	formatLay := "2006-01-02 15:04:05"
	timeStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", year, month, day, hour, minute)
	inRunTime, err := time.ParseInLocation(formatLay, timeStr, userLoc)
	if err != nil {
		return inRunTime, goerrors.New("", err.Error(), ioterrs.ErrUserTimeToLocalTime)
	}
	return inRunTime.In(localLoc), nil //运行时间
}

func getTimeLocation(timezone string) *time.Location {
	locStr := timezone
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	return loc
}
