package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_message_service/config"
	model "cloud_platform/iot_model"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	config.InitTest("../")
	dbcnf := model.DBConfig{
		DBType:     config.Global.Database.Driver,
		ConnectStr: config.Global.Database.Connstr,
	}
	err := model.InitDB(dbcnf)
	if err != nil {
		iotlogger.LogHelper.Error(err)
		return
	}
	svr := &SendMessageSvc{Ctx: context.Background()}
	svr.SendMessage(7940641025993637888, &protosService.SendMessageRequest{
		TplCode:     "SceneIntelligenceNotice",
		Params:      nil,
		TimeUnix:    time.Now().Unix(),
		SourceTable: "Intelligence",
		SourceRowId: "7940641025993637888",
		HomeId:      0,
		UserId:      []int64{7800843423417532416},
		ProductKey:  "",
		IsPublic:    false,
		Url:         "",
		PushTo:      "device",
		ChildType:   1,
		Subject:     "推送消息",
		Lang:        "en",
	})
}
