package execute

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"
	"fmt"
	"time"
)

type NoticeParams struct {
	HomeId           int64
	UserId           []int64
	IntelligenceName string
	AppKey           string // APP Key
	TenantId         string // 开发者租户编号
}

func NoticeExecute(params NoticeParams, resultId int64, runtime int64, taskItem *protosService.SceneIntelligenceTask) (*model.TSceneIntelligenceResultTask, error) {
	var (
		err              error
		isSuccess        int32  = 0
		resultMsg        string = iotconst.RUN_CONTINUTE
		subject          string = fmt.Sprintf("一键执行消息提醒")
		homeId           int64  = params.HomeId
		userIds                 = []int64{}
		intelligenceName string = params.IntelligenceName
	)

	iotlogger.LogHelper.Info("任务参数：", params, "通知参数：", taskItem)

	//TODO 推送给家庭成员，需要通过家庭ID获取家庭成员人员列表
	users, err := GetUsersByHomeId(homeId)
	iotlogger.LogHelper.Info("发送消息给1：", userIds)
	if err != nil {
		iotlogger.LogHelper.Error("获取家庭用户异常", err)
		return nil, err
	}
	iotlogger.LogHelper.Info("==========>555：")
	if users.Code != 200 {
		iotlogger.LogHelper.Error("未获取到家庭用户")
		return nil, errors.New("未获取到家庭用户")
	}
	iotlogger.LogHelper.Info("==========>666：")
	for _, u := range users.Data {
		userIds = append(userIds, u.UserId)
	}
	iotlogger.LogHelper.Info("发送消息给：", userIds)
	//发送消息
	err = SendMessage(&protosService.SendMessageRequest{
		TplCode:     "SceneIntelligenceNotice",
		Params:      map[string]string{"title": intelligenceName},
		TimeUnix:    time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		SourceTable: "t_scene_intelligence",
		SourceRowId: iotutil.ToString(taskItem.IntelligenceId),
		HomeId:      homeId,
		UserId:      userIds,
		IsPublic:    false,
		PushTo:      "device",
		ChildType:   22,
		Subject:     subject,
		Lang:        "", //不指定语言则，则全语言推送
		AppKey:      params.AppKey,
		TenantId:    params.TenantId,
	})
	if err != nil {
		resultMsg = iotconst.RUN_ERORR
		isSuccess = 2
	} else {
		resultMsg = iotconst.RUN_SUCESS
		isSuccess = 1
	}
	//结果
	saveObj := &model.TSceneIntelligenceResultTask{
		Id:             iotutil.GetNextSeqInt64(),
		IntelligenceId: taskItem.IntelligenceId,
		TaskId:         taskItem.Id,
		ResultId:       resultId,
		IsSuccess:      isSuccess,
		ResultMsg:      resultMsg,
		TaskImg:        taskItem.TaskImg,
		TaskDesc:       taskItem.TaskDesc,
		TaskType:       taskItem.TaskType,
		FuncDesc:       taskItem.FuncDesc,
		StartTime:      time.Unix(runtime, 0),
	}
	db := iotmodel.GetDB()
	err = db.Save(saveObj).Error
	if err != nil {
		return saveObj, err
	}
	return saveObj, nil
}
