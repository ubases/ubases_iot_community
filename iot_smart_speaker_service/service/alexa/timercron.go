package alexa

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"fmt"
	"sync"

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

var RefreshTokenUserIds *sync.Map = &sync.Map{}

func ScheduleTask() {
	CronCtx = NewCron()
	//定时刷新alexa token
	_, err := CronCtx.AddFunc("0 */50 * * * *", func() {
		defer iotutil.PanicHandler()
		RefreshTokenUserIds.Range(func(key, value interface{}) bool {
			GetAlexaRefreshToken(key.(string))
			return true
		})
	})
	if err != nil {
		fmt.Println(err)
	}
	CronCtx.Start()
}

func InitTimerCron() {
	getNeedRefreshUserId()
	ScheduleTask()
}

// 初始获取需要刷新Token的用户
func getNeedRefreshUserId() {
	userCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.AlexaVoiceAllUserIdKey)
	if userCmd.Err() != nil {
		return
	}
	for k, _ := range userCmd.Val() {
		RefreshTokenUserIds.Store(k, 0)
	}
}
