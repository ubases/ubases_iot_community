package push

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"cloud_platform/iot_message_service/service/push/wechatclient"
	"fmt"
)

type WechatClient struct {
	Cfg         *config.JpushCfg
	WechatCli   *wechatclient.WechatClient
	LangCt      map[string]pushModel.MessageRequestModel
	InputTarget pushModel.MessageTarget
	Message     pushModel.MessageRequestModel
}

func (g *WechatClient) ClearAlias(userId, appKey string) error {
	//TODO implement me
	return nil
}

func (g *WechatClient) checkConfig(appKey string) error {
	var err error
	if g.Cfg == nil {
		g.Cfg, err = getJsPushConfig(appKey)
		if err != nil {
			iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
			return err
		}
	}
	return nil
}

func (g *WechatClient) checkLang(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	if g.LangCt == nil {
		g.LangCt = map[string]pushModel.MessageRequestModel{}
		if inputTarget.IsPublic {
			//公共消息无法进行翻译，直接推送
		} else {
			//处理翻译（中文、英文）
			for _, lang := range langs {
				var m pushModel.MessageRequestModel
				iotutil.StructToStruct(message, &m)
				setTranslate(lang, &m)
				g.LangCt[lang] = m
			}
		}
	}
	return nil
}

func (g *WechatClient) PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	iotlogger.LogHelper.Infof("WechatClient.PushMessage wechat开始推送")
	g.InputTarget = inputTarget
	g.Message = message
	defer iotutil.PanicHandler()
	g.checkConfig(message.AppKey)
	g.checkLang(inputTarget, message)
	requets := make([]wechatclient.NoticeRequest, 0)
	for _, token := range inputTarget.PushTokens {
		if token.Lang == "" {
			iotlogger.LogHelper.Errorf("WechatClient.PushMessage 未设置语言")
			continue
		}
		//初始化
		appPushPlatform, _ := iotutil.ToIntErr(token.AppPushPlatform)
		switch appPushPlatform {
		case iotconst.PlatformWechat:
			data := wechatclient.NewSdkRequest()
			data.SetGMap(map[string]interface{}{
				"title":     g.LangCt[token.Lang].Title,
				"content":   g.LangCt[token.Lang].Content,
				"url":       g.LangCt[token.Lang].Url,
				"messageId": g.LangCt[token.Lang].MessageId,
				"childType": g.LangCt[token.Lang].ChildType,
				"type":      g.LangCt[token.Lang].Type,
				"model":     g.LangCt[token.Lang].Model,
			})
			//g.checkPlatformTokens(requets, token.Lang, iotconst.PlatformWechat)
			requets = append(requets, wechatclient.NoticeRequest{
				OpenId:     token.AppPushId,
				TemplateId: config.Global.ThirdPartyLogin.MiniProgram.TemplateId,
				PageCode:   config.Global.ThirdPartyLogin.MiniProgram.Page,
				Data:       data,
			})
		}
	}

	var msgInfo map[string]interface{}
	sendMessage := pushModel.MessageSendModel{}.SetModel(message)
	iotutil.StructToStruct(sendMessage, &msgInfo)
	if len(requets) > 0 {
		err := g.WechatCli.SendPush(requets)
		if err != nil {
			iotlogger.LogHelper.Errorf("g.GorushCli.SendPush error: %v", err.Error())
			return err
		}
	}
	iotlogger.LogHelper.Infof("WechatClient.PushMessage Gorush推送完成", len(requets))
	return nil
}

func (g *WechatClient) checkPlatformTokens(platformTokens map[int]map[string][]string, lang string, appPushPlatform int) {
	if platformTokens[appPushPlatform] == nil {
		platformTokens[appPushPlatform] = make(map[string][]string)
	}
	if platformTokens[appPushPlatform][lang] == nil {
		platformTokens[appPushPlatform][lang] = make([]string, 0)
	}
}
