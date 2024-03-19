package controls

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_model/db_app/model"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"time"
)

type MessageAppInfo struct {
	AppKey   string
	TenantId string
}

func SetAppInfo(tenantId, appKey string) (res MessageAppInfo) {
	res = MessageAppInfo{}
	defer iotutil.PanicHandler()
	res = MessageAppInfo{
		TenantId: tenantId,
		AppKey:   appKey,
	}
	return res
}

// SendAppMessage 推送消息
// homeId 推送给家庭
// userIds 推送给用户
// tplCode 消息模板，在云管平台消息模板中进行配置
// AddDevice
// subject 推送消息的主题，建议放到常量中
// params 模板对应的动态参数
func SendAppMessage(appInfo MessageAppInfo, pushTo string, homeId int64, userIds []int64,
	tplCode string, subject string, ChildType int32, sourceTable, sourceRowId string, params map[string]string) {
	defer iotutil.PanicHandler(appInfo, pushTo, homeId, userIds, tplCode, subject, ChildType, params)
	pushMsg := &protosService.SendMessageRequest{
		TplCode:     tplCode,
		Params:      params,
		TimeUnix:    time.Now().Add(time.Duration(1) * time.Hour).Unix(), //消息一小时有效
		HomeId:      homeId,
		UserId:      userIds,
		IsPublic:    false,
		PushTo:      pushTo,
		ChildType:   ChildType,
		Subject:     subject,
		AppKey:      appInfo.AppKey,
		TenantId:    appInfo.TenantId,
		SourceRowId: sourceRowId,
		SourceTable: sourceTable,
		Lang:        "", //不指定语言则，则全语言推送
	}
	//发送消息  测试消息推送
	//TODO 修改为消息队列推送
	ret, err := rpc.ClientMessageService.SendMessage(context.Background(), pushMsg)
	if err != nil {
		iotlogger.LogHelper.WithTag("method", "SendAppMessage").Error(err)
		return
	}
	if ret.Code != 200 {
		iotlogger.LogHelper.WithTag("method", "SendAppMessage").Error(ret.Message)
		return
	}
}

// SendFeedbackReplyMessage 用户消息（user)	修改密码-强制退出	14
func SendFeedbackReplyMessage(appInfo MessageAppInfo, userId int64, feedbackId string) {
	defer iotutil.PanicHandler()
	if appInfo.AppKey == "" && appInfo.TenantId == "" {
		iotlogger.LogHelper.Infof("SendFeedbackReplyMessage error tenantId:%s, appKey: %s", appInfo.TenantId, appInfo.AppKey)
		return
	}
	iotlogger.LogHelper.Infof("SendFeedbackReplyMessage start")
	var (
		pushTo    string = "notice"
		childType int32  = 7
		subject   string = "反馈回复通知"
		tplCode   string = iotconst.APP_MESSAGE_FEEDBACK_REPLY
	)
	SendAppMessage(appInfo, pushTo, 0, []int64{userId}, tplCode, subject, childType, model.TableNameTUcUserFeedback, feedbackId, map[string]string{})
}
