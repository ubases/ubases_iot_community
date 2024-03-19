package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_log_service/config"
	"time"

	"github.com/robfig/cron/v3"
)

var cronJob *CronJob

type CronJob struct {
	C *cron.Cron
	// TTL time.Duration
}

// var c *cron.Cron

func InitCron() {
	cronJob = &CronJob{
		C: cron.New(),
		// TTL: ttl,
	}
}

func GetCronJob() *CronJob {
	return cronJob
}

func (cj *CronJob) StartCronJob() error {
	cl := ClearLog{
		Days: config.Global.Database.Days,
	}
	// _, err = cj.C.AddJob("@daily", cl)
	// if err != nil {
	// 	return err
	// }
	_, err := cj.C.AddJob("*/1 * * * *", cl)
	if err != nil {
		return err
	}
	cj.C.Start()
	return nil
}

func (cj *CronJob) StopCronJob() {
	cj.C.Stop()
}

type ClearLog struct {
	Days int
}

func (cl ClearLog) Run() {
	delTime := iotutil.GetTodaySartTime(time.Now()).AddDate(0, 0, -cl.Days)
	iotlogger.LogHelper.Helper.Debugf("delTime: %s", delTime)
	if err := DeleteAppLogRecordsByTime(delTime); err != nil {
		iotlogger.LogHelper.Helper.Errorf("clear app log records time < %s error: %v", delTime, err)
		return
	}
}
