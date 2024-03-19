package execute

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_proto/protos/protosService"
	"strconv"
	"strings"
	"time"
)

func DelayedExecute(resultId int64, runtime int64, taskItem *protosService.SceneIntelligenceTask) (*model.TSceneIntelligenceResultTask, error) {
	delayed := taskItem.FuncValue //Delayed 00:01
	taskInfo := &model.TSceneIntelligenceResultTask{
		Id:             iotutil.GetNextSeqInt64(),
		IntelligenceId: taskItem.IntelligenceId,
		TaskId:         taskItem.Id,
		IsSuccess:      0,
		ResultId:       resultId,
		ResultMsg:      iotconst.RUN_CONTINUTE,
		TaskImg:        taskItem.TaskImg,
		TaskDesc:       taskItem.TaskDesc,
		TaskType:       taskItem.TaskType,
		FuncDesc:       taskItem.FuncDesc,
		StartTime:      time.Unix(runtime, 0),
	}
	db := iotmodel.GetDB()
	err := db.Save(taskInfo).Error
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return taskInfo, err
	}
	delayeds := strings.Split(delayed, ":")
	delayedHandle([]string{delayeds[0], delayeds[1]})
	taskInfo.IsSuccess = 1
	taskInfo.ResultMsg = iotconst.RUN_SUCESS

	err = db.Save(taskInfo).Error
	if err != nil {
		iotlogger.LogHelper.Error(err)
		taskInfo.IsSuccess = 2
		taskInfo.ResultMsg = err.Error()
		return taskInfo, nil
	}
	return taskInfo, nil
}

// 延时功能
// [小时 ,分钟 ]
// ['00','02']
func delayedHandle(delayed []string) {
	//延时处理
	minute, _ := strconv.Atoi(delayed[0])
	second, _ := strconv.Atoi(delayed[1])
	totalSecond := minute*60 + second
	if totalSecond > 0 {
		time.Sleep(time.Duration(totalSecond) * time.Second)
	}
}
