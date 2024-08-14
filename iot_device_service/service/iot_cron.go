package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_device_service/config"
	"cloud_platform/iot_device_service/service/job"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron/v3"
	"sync"
)

var (
	cronObj *CronObj
	json    = jsoniter.ConfigCompatibleWithStandardLibrary
)

type CronObj struct {
	c      *cron.Cron
	m      map[int64]*DeviceJob
	endM   map[int64]*DeviceJob
	mux    *sync.RWMutex
	endMux *sync.RWMutex
}

func NewCron() {
	cronObj = &CronObj{
		c:      cron.New(cron.WithSeconds()),
		m:      make(map[int64]*DeviceJob),
		endM:   make(map[int64]*DeviceJob),
		mux:    new(sync.RWMutex),
		endMux: new(sync.RWMutex),
	}
}

func GetCron() *CronObj {
	return cronObj
}

type DeviceJob struct {
	Id         int64
	ProductKey string
	DeviceId   string
	TaskId     int64
	TaskType   int32
	TimerType  int32 // 1: 单个时间点的定时，2：时间段的定时
	IsEndTimer int32 // 1: 开始定时，2：结束定时
	Cron       string
	Data       string
	EntryID    cron.EntryID
}

func (dj DeviceJob) Run() {
	iotlogger.LogHelper.Infof("send job %s data %s to nats", dj.DeviceId, dj.Data)
	var dataMap map[string]interface{}
	if err := json.Unmarshal([]byte(dj.Data), &dataMap); err != nil {
		iotlogger.LogHelper.Errorf("解析载荷信息失败,内容[%s],错误:%s", dj.Data, err.Error())
	} else {
		job.PubControl(dj.ProductKey, dj.DeviceId, dataMap)
	}
	if dj.TaskType == iotconst.Task_Countdown_Job {
		// 调用iot_device_service接口, 关闭倒计时任务
		req := &protosService.IotDeviceCountdownJobReq{
			Id: dj.Id,
		}
		countdownCtx := IotDeviceCountdownSvcEx{Ctx: context.Background()}
		err := countdownCtx.StopIotDeviceCountdownJob(req)
		if err != nil {
			iotlogger.LogHelper.Errorf("倒计时任务已执行完，关闭倒计时任务失败：%v", err)
		}
	} else if dj.TaskType == iotconst.Task_Timer_Job_Once {
		// 调用iot_device_service接口, 关闭定时任务(只执行一次)
		if dj.TimerType == 1 {
			req := &protosService.IotDeviceTimerJobReq{
				Id: dj.Id,
			}
			timerCtx := IotDeviceTimerSvcEx{Ctx: context.Background()}
			err := timerCtx.StopIotDeviceTimerJob(req)
			if err != nil {
				iotlogger.LogHelper.Errorf("定时任务(Once)已执行完，关闭倒计时任务失败：%v", err)
			}
		} else if dj.TimerType == 2 {
			if dj.IsEndTimer == 2 {
				req := &protosService.IotDeviceTimerJobReq{
					Id: dj.Id,
				}
				timerCtx := IotDeviceTimerSvcEx{Ctx: context.Background()}
				err := timerCtx.StopIotDeviceTimerJob(req)
				if err != nil {
					iotlogger.LogHelper.Errorf("定时任务(Once)已执行完，关闭倒计时任务失败：%v", err)
				}
			}
		}
	}
}

// 创建定时任务
func (co *CronObj) CreateJob(job *protosService.IotJob) error {
	dj := &DeviceJob{
		Id:         job.Id,
		ProductKey: job.ProductKey,
		DeviceId:   job.DeviceId,
		TaskId:     job.TaskId,
		TaskType:   job.TaskType,
		TimerType:  1,
		Cron:       job.Cron,
		Data:       job.Data,
	}
	EntryID, err := co.c.AddJob(dj.Cron, dj)
	if err != nil {
		return err
	}
	dj.EntryID = EntryID
	if job.EndCron != "" && job.EndData != "" {
		// 如果是时间段定时任务，则重置开始定时的类型
		dj.TimerType = 2
		dj.IsEndTimer = 1
		djEnd := &DeviceJob{
			Id:         job.Id,
			ProductKey: job.ProductKey,
			DeviceId:   job.DeviceId,
			TaskId:     job.TaskId,
			TaskType:   job.TaskType,
			TimerType:  2,
			IsEndTimer: 2,
			Cron:       job.EndCron,
			Data:       job.EndData,
		}
		EndEntryID, err := co.c.AddJob(djEnd.Cron, djEnd)
		if err != nil {
			return err
		}
		djEnd.EntryID = EndEntryID
		co.endMux.Lock()
		defer co.endMux.Unlock()
		co.endM[djEnd.Id] = djEnd
	}
	co.mux.Lock()
	defer co.mux.Unlock()
	co.m[dj.Id] = dj
	return nil
}

// 删除定时任务
func (co *CronObj) DeleteJob(job *protosService.IotJob) error {
	co.mux.RLock()
	defer co.mux.RUnlock()
	dj, ok := co.m[job.Id]
	if !ok {
		return nil
	}
	co.c.Remove(dj.EntryID)
	delete(co.m, job.Id)
	co.endMux.RLock()
	defer co.endMux.RUnlock()
	djEnd, ok := co.endM[job.Id]
	if !ok {
		return nil
	}
	co.c.Remove(djEnd.EntryID)
	delete(co.endM, job.Id)
	return nil
}

func (co *CronObj) Entries() []cron.Entry {
	return co.c.Entries()
}

func (co *CronObj) Start() {
	co.c.Start()
}

func (co *CronObj) Stop() {
	co.c.Stop()
}

func InitCron() error {
	svc := new(IotJobSvc)
	req := &protosService.IotJobListRequest{
		Query: &protosService.IotJob{
			Enabled:        1,
			RegionServerId: config.Global.Service.ServerId, //config中读取服务器Id
		},
	}
	jobs, _, err := svc.GetListIotJob(req)
	if err != nil {
		return err
	}

	for i := range jobs {
		if err := cronObj.CreateJob(jobs[i]); err != nil {
			iotlogger.LogHelper.Errorf("create task:[%s] job error: [%s]", jobs[i].TaskId, err)
			continue
		}
	}
	cronObj.c.Start()
	return nil
}
