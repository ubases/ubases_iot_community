package services

import (
	"bytes"
	"cloud_platform/iot_app_api_service/controls/message/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"fmt"
	"html/template"

	"go-micro.dev/v4/metadata"

	"context"
	"errors"
	"sort"
)

type MessageService struct {
	Ctx context.Context
}

func (s *MessageService) SetContext(ctx context.Context) *MessageService {
	s.Ctx = ctx
	return s
}

// GetMessageRedDot 获取消息统计数据（红点数据）
func (s *MessageService) GetMessageRedDot(lang string, userId int64) (*entitys.MpMessageRedDotEntitys, error) {
	if userId == 0 {
		return nil, errors.New("参数异常")
	}
	resp := &entitys.MpMessageRedDotEntitys{}
	//if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.RED_DOT_DATA, userId), resp); err == nil {
	//	return resp, nil
	//}
	rep, err := rpc.MessageRedDotService.Find(s.Ctx, &protosService.MpMessageRedDotFilter{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}

	homeListInfo, err := rpc.TUcUserService.HomeList(context.Background(), &protosService.UcUser{
		Id: userId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("获取用户家庭列表失败，原因:%s", err.Error())
		return nil, err
	}
	homeMap := make(map[int64]string)
	homeIds := make([]int64, 0)
	for _, v := range homeListInfo.HomeUsers {
		homeMap[v.Id] = services.HomeLanguage(lang, v.Name)
		homeIds = append(homeIds, iotutil.ToInt64(v.Id))
	}

	deviceList := []string{}
	deviceHomeList, deviceErr := rpc.IotDeviceHomeService.UserDevList(context.Background(), &protosService.IotDeviceHomeHomeId{HomeIds: homeIds})
	if deviceErr != nil {
		iotlogger.LogHelper.Error("获取家庭设备列表失败，原因:%s", deviceErr.Error())
		//return nil, err
	} else {
		for _, v := range deviceHomeList.DevList {
			deviceList = append(deviceList, v.Did)
		}
	}
	if len(rep.Data) != 0 {
		resp = entitys.MpMessageRedDot_pb2e(rep.Data[0], deviceList)
	}
	//if err := cached.RedisStore.Set(persist.GetRedisKey(iotconst.RED_DOT_DATA, userId), resp, 600*time.Second); err != nil {
	//	return nil, err
	//}
	return resp, err
}

// GetMessageRedDotOld 获取消息统计数据（红点数据） 兼容历史数据
func (s *MessageService) GetMessageRedDotOld(userId int64) ([]*entitys.MpMessageRedDotOld, error) {
	if userId == 0 {
		return nil, errors.New("参数异常")
	}
	rep, err := rpc.MessageRedDotService.Find(s.Ctx, &protosService.MpMessageRedDotFilter{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}
	redDotList := make([]*entitys.MpMessageRedDotOld, 0)

	alarm := entitys.MpMessageRedDotOld{}
	home := entitys.MpMessageRedDotOld{}
	notice := entitys.MpMessageRedDotOld{}
	device := entitys.MpMessageRedDotOld{}
	if len(rep.Data) == 0 {
		redDotList = append(redDotList, alarm.Set("alarm", "", "", false))
		redDotList = append(redDotList, home.Set("home", "", "", false))
		redDotList = append(redDotList, notice.Set("notice", "", "", false))
		redDotList = append(redDotList, device.Set("device", "", "", false))
	} else {
		data := rep.Data[0]
		redDotList = append(redDotList, alarm.Set("alarm", "", "", false))
		redDotList = append(redDotList, home.Set("home", "", "", data.HomeMsg > 0))
		redDotList = append(redDotList, notice.Set("notice", "", "", data.SystemMsg > 0))
		redDotList = append(redDotList, device.Set("device", "", "", data.DeviceMsg > 0))
	}
	return redDotList, err
}

// ClearMessage 清空消息
func (s *MessageService) ClearMessage(userId int64, typeStr, deviceId string) error {
	if userId == 0 {
		return errors.New("参数异常")
	}
	messageType := s.getMessageType(typeStr)
	rep, err := rpc.MessageUserInService.Delete(s.Ctx, &protosService.MpMessageUserIn{
		MessageType: messageType,
		UserId:      userId,
		Did:         deviceId,
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

func (s *MessageService) DeleteByIdMessage(id string) error {
	if id == "" {
		return errors.New("参数异常")
	}
	rep, err := rpc.MessageUserInService.DeleteById(s.Ctx, &protosService.MpMessageUserIn{
		Id: iotutil.ToInt64(id),
	})
	if err != nil {
		return err
	}
	if rep.Code != 200 {
		return errors.New(rep.Message)
	}
	return nil
}

// QueryMessageGroupList 分组消息列表
func (s *MessageService) QueryMessageGroupList(userId int64, typeStr string, devId, messageId string) ([]*entitys.MpMessageUserInGroupItem, error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	if userId == 0 {
		return nil, errors.New("参数异常")
	}
	if typeStr == "" {
		return nil, errors.New("参数异常, type")
	}
	messageType := s.getMessageType(typeStr)
	var messageIdInt int64 = 0
	if messageId != "" {
		messageIdInt = iotutil.ToInt64(messageId)
	}
	rep, err := rpc.MessageUserInService.GroupLists(s.Ctx, &protosService.MpMessageUserInListRequest{
		Query: &protosService.MpMessageUserIn{
			MessageType: messageType,
			UserId:      userId,
			Did:         devId,
			MessageId:   messageIdInt,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, errors.New(rep.Message)
	}

	//读取翻译缓存
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_MESSAGE_TEMPLATE).Result()
	if err != nil {
		langMap = make(map[string]string)
	}

	var resultList = []*entitys.MpMessageUserInGroupItem{}
	for _, item := range rep.Data {
		var items []*entitys.MpMessageUserInEntitys
		for _, d := range item.Data {
			row := entitys.MpMessageUserIn_pb2e(d)
			//翻译转换
			fKey := fmt.Sprintf("%s_%s_tplContent", lang, d.TplCode)
			params := iotutil.JsonToMap(d.Params)
			if langMap[fKey] != "" {
				row.PushContent, _ = s.paramIntoContent(iotutil.ToString(langMap[fKey]), params)
			} else {
				row.PushContent, _ = s.paramIntoContent(d.PushContent, params)
			}
			subjectKey := fmt.Sprintf("%s_%s_tplSubject", lang, d.TplCode)
			if langMap[subjectKey] != "" {
				row.PushTitle = iotutil.ToString(langMap[subjectKey])
			}
			items = append(items, row)
		}
		resultList = append(resultList, &entitys.MpMessageUserInGroupItem{
			Date: item.Date,
			Data: items,
		})
	}
	//根据sort进行排序
	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].Date > resultList[j].Date
	})

	return resultList, err
}

func (s *MessageService) paramIntoContent(templateContent string, params interface{}) (string, error) {
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

// QueryMessageList 查询消息列表
func (s *MessageService) QueryMessageList(userId int64, typeStr string, filter entitys.MpMessageUserOutQuery) ([]*entitys.MpMessageUserInEntitys, int64, error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	if userId == 0 {
		return nil, 0, errors.New("参数异常")
	}
	if typeStr == "" {
		return nil, 0, errors.New("参数异常, type")
	}
	messageType := s.getMessageType(typeStr)
	rep, err := rpc.MessageUserInService.Lists(s.Ctx, &protosService.MpMessageUserInListRequest{
		Page:     filter.Page,
		PageSize: filter.Limit,
		Query: &protosService.MpMessageUserIn{
			MessageType: messageType,
			UserId:      userId,
			Did:         filter.Query.Did,
			Id:          filter.Query.Id,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	if rep.Code != 200 {
		return nil, 0, errors.New(rep.Message)
	}
	var resultList = []*entitys.MpMessageUserInEntitys{}

	//读取翻译缓存
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_MESSAGE_TEMPLATE).Result()
	if err != nil {
		langMap = make(map[string]string)
	}

	for _, item := range rep.Data {
		row := entitys.MpMessageUserIn_pb2e(item)
		//翻译转换
		fKey := fmt.Sprintf("%s_%s_tplContent", lang, item.TplCode)
		if langMap[fKey] != "" {
			params := iotutil.JsonToMap(item.Params)
			row.PushContent, _ = s.paramIntoContent(iotutil.ToString(langMap[fKey]), params)
		}
		subjectKey := fmt.Sprintf("%s_%s_tplSubject", lang, item.TplCode)
		if langMap[subjectKey] != "" {
			row.PushTitle = iotutil.ToString(langMap[subjectKey])
		}
		resultList = append(resultList, row)
	}
	return resultList, rep.Total, err
}

func (s MessageService) getMessageType(typeStr string) int32 {
	var messageType int32 = 0
	switch typeStr {
	case "home":
		messageType = 1
	case "notice":
		messageType = 2
	case "device":
		messageType = 3
	case "alarm":
		messageType = 4
	}
	return messageType
}
