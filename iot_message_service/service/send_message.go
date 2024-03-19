package service

import (
	"bytes"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/cached"
	"cloud_platform/iot_message_service/convert"
	"cloud_platform/iot_message_service/service/push"
	"cloud_platform/iot_message_service/service/push/pushModel"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_message/model"
	"cloud_platform/iot_model/db_message/orm"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"text/template"
	"time"

	"go-micro.dev/v4/logger"
)

type SendMessageSvc struct {
	Ctx context.Context
}

// 消息类型转换
func (s SendMessageSvc) convertMessageType(messageType int32) string {
	//需要从缓存获取。
	msgTypeStr := ""
	switch messageType {
	case 1:
		msgTypeStr = "home"
	case 2:
		msgTypeStr = "notice"
	case 3:
		msgTypeStr = "device"
	case 4:
		msgTypeStr = "alrame"
	}
	return msgTypeStr
}

/*
原消息格式
msg := map[string]interface{}{
	"userids": userid,
	"homeid":homeid,
	"ispublic":0,
	"type":msgtype,
	"objectid": objectId,
	"title": title,
	"content": content,
	"childtype": objectType,
	"params": params,
}

流程：
通过tplCode, 参数、
通过tplCode可以得到messageType、agentType、pushType、tplContent  childType应该也可以得到、

请求参数定义
1、tplCode
2、参数的map[string]interface
3、time
4、SourceTable
5、SourceRowId
6、homeId
7、userId
8、tags、alias、regIds
   alias别名可以为： homeid、userid
   tags标签可以为：  productKey_version（固件升级通知）
                  ChinaArea
                  ....
9、isPublic
10、url
*/

// 从文本内容加载
func (s *SendMessageSvc) paramIntoContent(templateContent string, params interface{}) (string, error) {
	tmp, err := template.New("TemplateContent").Parse(templateContent)
	if err != nil {
		return templateContent, err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, params); err != nil {
		return templateContent, err
	}
	return buf.String(), nil
}

// SendMessage 发送消息
func (s SendMessageSvc) SendMessage(messageId int64, message *protosService.SendMessageRequest) error {
	//读取模板
	tplSvc := MpMessageTemplateSvc{Ctx: s.Ctx}
	templateInfo, err := tplSvc.GetMessageTemplateByCode(message.TplCode)
	if err != nil {
		return err
	}
	message.Subject = templateInfo.TplName

	//如果是发送给设备，需要指定这边id
	devId := ""
	if message.PushTo == "device" {
		devId = message.SourceRowId
	}
	//content, _ := s.paramIntoContent(templateInfo.TplContent, message.Params)
	//转换是否为公共消息
	var isPublicInt int32 = 0
	if message.IsPublic {
		isPublicInt = 1
	}

	messageType := s.convertMessageType(templateInfo.MessageType)
	pushMessage := pushModel.MessageRequestModel{
		Type:        messageType,
		ChildType:   iotutil.ToString(message.ChildType),
		ObjectType:  message.SourceTable,
		ObjectId:    message.SourceRowId,
		Model:       message.ProductKey,
		Title:       message.Subject,
		TplCode:     message.TplCode,
		Content:     templateInfo.TplContent,
		Devids:      devId,
		Url:         message.Url,
		IsRead:      false,
		UnSetExpire: false,
		IsPublic:    isPublicInt,
		TagType:     0, //iotutil.ToInt32(message.MsgTag), //0：alias 、1：tag、 2：regid
		HomeId:      iotutil.ToString(message.HomeId),
		UserId:      iotutil.ToString(message.UserId),
		Result:      "",
		CreateTime:  time.Now().Unix(),
		ExpireTime:  time.Now().Add(1 * time.Hour).Unix(),
		CreatedAt:   time.Now().Unix(),
		AppKey:      message.AppKey,
		TenantId:    message.TenantId,
		Params:      message.Params,
	}

	q := orm.Use(iotmodel.GetDB())
	//存储推送记录
	saveUserIn := convert.SetMpMessageUserIn(messageId, message.UserId[0], message, &pushMessage)
	saveUserIn.MessageType = templateInfo.MessageType
	saveUserIn.Params = iotutil.ToString(message.Params)
	t := q.TMpMessageUserIn
	err = t.WithContext(context.Background()).Create(saveUserIn)
	if err != nil {
		iotlogger.LogHelper.Errorf("CreateMpMessageUserIn error : %s", err.Error())
		return err
	}

	if len(message.UserId) == 0 {
		iotlogger.LogHelper.Errorf("CreateMpMessageUserIn error : no userid")
		return errors.New("no userid")
	}
	//TODO 通过用户Id，获取需要推送的手机Id(AppPushId)
	t1 := q.TAppPushTokenUser
	t2 := q.TAppPushToken
	var list []pushModel.PushTokenItem
	err = t1.WithContext(context.Background()).LeftJoin(t2, t1.AppPushId.EqCol(t2.AppPushId)).
		Select(t2.Lang, t1.UserId, t1.AppPushId, t1.AppKey, t1.TenantId, t1.RegionId, t2.AppToken, t2.AppPushPlatform, t2.AppPacketName).
		Where(t1.UserId.In(message.UserId...)).Scan(&list)
	if err != nil {
		iotlogger.LogHelper.Errorf("CreateMpMessageUserIn TAppPushTokenUser error : %s", err.Error())
		return err
	}

	userIdStr := []string{}
	for _, user := range message.UserId {
		userIdStr = append(userIdStr, iotutil.ToString(user))
	}
	pushMessage.MessageId = iotutil.ToString(saveUserIn.Id)
	//执行推送
	push.PushMgr.PushPush(push.PushInfo{
		pushModel.MessageTarget{
			IsPublic:   message.IsPublic,
			Alias:      userIdStr,
			PushTokens: list,
		}, pushMessage,
	})
	//pushMsg := push.PushMessage{}
	//pushMsg.SendMessage(pushModel.MessageTarget{
	//	IsPublic: message.IsPublic,
	//	Alias:    userIdStr,
	//}, pushMessage)

	//消息红点处理
	for _, user := range message.UserId {
		s.SetRedDot(user, templateInfo.MessageType, saveUserIn.Id)
	}
	return err
}

// SendLaserMessage 只发送激光消息，不保存消息记录和设置消息红点
func (s SendMessageSvc) SendLaserMessage(message *protosService.SendMessageRequest) error {
	//pushMsg := push.PushMessage{}
	//读取模板
	tplSvc := MpMessageTemplateSvc{Ctx: s.Ctx}
	templateInfo, err := tplSvc.GetMessageTemplateByCode(message.TplCode)
	if err != nil {
		return err
	}
	message.Subject = templateInfo.TplName

	//如果是发送给设备，需要指定这边id
	devId := ""
	if message.PushTo == "device" {
		devId = message.SourceRowId
	}
	content, _ := s.paramIntoContent(templateInfo.TplContent, message.Params)
	//转换是否为公共消息
	var isPublicInt int32 = 0
	if message.IsPublic {
		isPublicInt = 1
	}

	messageType := s.convertMessageType(templateInfo.MessageType)

	pushMessage := pushModel.MessageRequestModel{
		Type:        messageType,
		ChildType:   iotutil.ToString(message.ChildType),
		ObjectType:  message.SourceTable,
		ObjectId:    message.SourceRowId,
		Model:       message.ProductKey,
		Title:       message.Subject,
		TplCode:     message.TplCode,
		Content:     content,
		Devids:      devId,
		Url:         message.Url,
		IsRead:      false,
		UnSetExpire: false,
		IsPublic:    isPublicInt,
		TagType:     0,
		HomeId:      iotutil.ToString(message.HomeId),
		UserId:      iotutil.ToString(message.UserId),
		Result:      "",
		CreateTime:  time.Now().Unix(),
		ExpireTime:  time.Now().Add(1 * time.Hour).Unix(),
		CreatedAt:   time.Now().Unix(),
		AppKey:      message.AppKey,
		TenantId:    message.TenantId,
	}

	q := orm.Use(iotmodel.GetDB())
	//TODO 通过用户Id，获取需要推送的手机Id(AppPushId)
	t1 := q.TAppPushTokenUser
	t2 := q.TAppPushToken
	var list []pushModel.PushTokenItem
	err = t1.WithContext(context.Background()).LeftJoin(t2, t1.AppPushId.EqCol(t2.AppPushId)).
		Select(t2.Lang, t1.UserId, t1.AppPushId, t1.AppKey, t1.TenantId, t1.RegionId, t2.AppToken, t2.AppPushPlatform, t2.AppPacketName).
		Where(t1.UserId.In(message.UserId...)).Scan(&list)
	if err != nil {
		iotlogger.LogHelper.Errorf("CreateMpMessageUserIn TAppPushTokenUser error : %s", err.Error())
		return err
	}

	userIdStr := make([]string, 0)
	for _, user := range message.UserId {
		userIdStr = append(userIdStr, iotutil.ToString(user))
	}

	//执行推送
	//pushMsg.SendMessage(pushModel.MessageTarget{
	//	IsPublic:   message.IsPublic,
	//	Alias:      userIdStr,
	//	PushTokens: list,
	//}, pushMessage)

	//执行推送
	push.PushMgr.PushPush(push.PushInfo{
		pushModel.MessageTarget{
			IsPublic:   message.IsPublic,
			Alias:      userIdStr,
			PushTokens: list,
		}, pushMessage,
	})

	return nil
}

// 设置消息红点
func (s SendMessageSvc) SetRedDot(userId int64, messageType int32, messageId int64) {
	defer iotutil.PanicHandler(userId, messageType, messageId)
	q := orm.Use(iotmodel.GetDB())
	redDot, _ := q.TMpMessageRedDot.WithContext(context.Background()).
		Where(q.TMpMessageRedDot.UserId.Eq(userId)).Find()
	if redDot == nil || len(redDot) == 0 {
		redDot = []*model.TMpMessageRedDot{
			&model.TMpMessageRedDot{
				Id:     iotutil.GetNextSeqInt64(),
				UserId: userId,
			}}
	}
	//消息类型
	switch messageType {
	case 1: //"home":
		redDot[0].HomeMsg = 1
		redDot[0].HomeMsgNum++
	case 2: //"notice":
		redDot[0].SystemMsg = 1
		redDot[0].SystemMsgId = messageId
		redDot[0].SystemMsgNum++
	case 3, 4: //"device":"alarm":
		redDot[0].DeviceMsg = 1
		redDot[0].DeviceMsgNum++
	}
	err := q.TMpMessageRedDot.WithContext(context.Background()).Save(redDot[0])
	if err != nil {
		logger.Errorf("set message redhot error : %s", err.Error())
	}

	// 清除缓存的用户红点信息，下次从数据库重新读取，再加入缓存
	if err := cached.RedisStore.Delete(persist.GetRedisKey(iotconst.RED_DOT_DATA, userId)); err != nil {
		logger.Errorf("cache red dot message to redis error : %s", err.Error())
		return
	}
}

// SendMessageExt 发送消息
func (s SendMessageSvc) SendMessageExt(message protosService.MpMessageUserIn) {
	pushMsg := push.PushMessage{}
	messageType := s.convertMessageType(message.MessageType)
	pushMsg.SendMessage(pushModel.MessageTarget{
		IsPublic: message.IsPublic == 1,
		Alias:    []string{iotutil.ToString(message.UserId)},
	}, pushModel.MessageRequestModel{
		Type:       messageType,
		ChildType:  iotutil.ToString(message.ChildType),
		ObjectType: message.SourceTable,
		ObjectId:   iotutil.ToString(message.SourceRowId),
		Model:      message.ProductKey,
		//DevImg:      "",
		//HomeName:    message.HomeId,
		Title:   message.PushTitle,
		Content: message.PushContent,
		Devids:  message.Did,
		//Extands:     ,
		Url: message.Url,
		//Params:      message,
		IsRead:      false,
		UnSetExpire: false,
		IsPublic:    message.IsPublic,
		TagType:     iotutil.ToInt32(message.MsgTag),
		HomeId:      iotutil.ToString(message.HomeId),
		UserId:      iotutil.ToString(message.UserId),
		//UserIds:     nil,
		Result:     "",
		CreateTime: time.Now().Unix(),
		ExpireTime: time.Now().Add(1 * time.Hour).Unix(),
		CreatedAt:  time.Now().Unix(),
	})
}
