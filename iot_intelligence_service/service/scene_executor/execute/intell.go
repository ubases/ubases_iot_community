package execute

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	rules2 "cloud_platform/iot_intelligence_service/service/scene_executor/rules"
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_model/db_device/orm"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
	"time"
)

// SceneIntelligenceExecute 开关智能场景
func SceneIntelligenceExecute(resultId int64, runtime int64,
	taskItem *protosService.SceneIntelligenceTask) (*model.TSceneIntelligenceResultTask, error) {
	var (
		ctx                 = context.Background()
		intelligenceId      = taskItem.IntelligenceId
		otherIntelligenceId = taskItem.ObjectId
	)
	//结果
	theTaskResult := &model.TSceneIntelligenceResultTask{
		Id:             iotutil.GetNextSeqInt64(),
		IntelligenceId: intelligenceId,
		TaskId:         taskItem.Id,
		IsSuccess:      0,
		ResultId:       resultId,
		ResultMsg:      iotconst.RUN_CONTINUTE,
		TaskImg:        taskItem.TaskImg,
		TaskDesc:       taskItem.TaskDesc,
		TaskType:       taskItem.TaskType,
		FuncDesc:       taskItem.FuncDesc,
		FuncValue:      taskItem.FuncValue,
		ObjectDesc:     taskItem.ObjectDesc,
		ObjectId:       taskItem.ObjectId,
		FuncKey:        taskItem.FuncKey,
		StartTime:      time.Unix(runtime, 0),
	}
	db := orm.Use(iotmodel.GetDB())
	resultTaskDo := db.TSceneIntelligenceResultTask.WithContext(ctx)
	err := resultTaskDo.Save(theTaskResult)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		theTaskResult.IsSuccess = 2
		theTaskResult.ResultMsg = err.Error()
		err = resultTaskDo.Save(theTaskResult)
		return theTaskResult, err
	}
	err = RemoveRules(iotutil.ToString(intelligenceId))
	if err != nil {
		iotlogger.LogHelper.Error("开关智能场景移除规则异常", err)
		//TODO 考虑移除RULE报错的场景，决定是否需要直接返回错误
	}

	status, err := iotutil.ToInt32Err(taskItem.FuncValue)
	if err != nil {
		theTaskResult.IsSuccess = 2
		theTaskResult.ResultMsg = fmt.Sprintf("智能场景的开关功能值异常，应为Int类型，实际值为%v", taskItem.FuncValue)
	} else {
		//修改另一个智能场景的状态
		_, err = db.TSceneIntelligence.WithContext(ctx).Select(db.TSceneIntelligence.Status).Updates(model.TSceneIntelligence{
			Id:     iotutil.ToInt64(otherIntelligenceId),
			Status: status,
		})
		if err != nil {
			theTaskResult.IsSuccess = 0
			theTaskResult.ResultMsg = err.Error()
		} else {
			theTaskResult.IsSuccess = 1
			theTaskResult.ResultMsg = iotconst.RUN_SUCESS
		}
		go func(id string, state int32) {
			rules2.CreateRuleChan <- rules2.CreateRuleChanData{
				Id:     id,
				Desc:   "",
				Status: state,
			}
		}(otherIntelligenceId, status)
	}
	//结果
	err = resultTaskDo.Save(theTaskResult)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		theTaskResult.IsSuccess = 2
		theTaskResult.ResultMsg = err.Error()
		err = resultTaskDo.Save(theTaskResult)
		return theTaskResult, err
	}
	return theTaskResult, nil
}

// SceneIntelligenceClose 关闭智能场景
func SceneIntelligenceClose(intellId int64) (err error) {
	db := orm.Use(iotmodel.GetDB())
	//修改另一个智能场景的状态
	_, err = db.TSceneIntelligence.WithContext(context.Background()).Select(db.TSceneIntelligence.Status).Updates(model.TSceneIntelligence{
		Id:     intellId,
		Status: 2,
	})
	RemoveRules(iotutil.ToString(intellId))
	return err
}

// RemoveRules 移除规则
func RemoveRules(intelligenceId string) error {
	if valscene.RuleBuilder == nil {
		return nil
	}
	var err error
	for i := 0; i < 3; i++ {
		//移除规则
		err = valscene.RuleBuilder.RemoveRules([]string{intelligenceId})
		if err == nil {
			break
		}
	}
	return err
}
