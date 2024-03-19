package cron

import (
	"github.com/robfig/cron/v3"
)

var (
	CronCtx *cron.Cron
)

func NewCron() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func ScheduleTask() {
	CronCtx = NewCron()
	//RunTimerTask("0 0 * * * *")
	CronCtx.Start()
}

func InitTimerCron() {
	ScheduleTask()
}
