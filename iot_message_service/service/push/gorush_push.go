package push

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/platforms"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"fmt"
)

type GorushClient struct {
	Cfg         *config.JpushCfg
	GorushCli   *gorushclient.GorushClient
	LangCt      map[string]pushModel.MessageRequestModel
	InputTarget pushModel.MessageTarget
	Message     pushModel.MessageRequestModel
}

func (g *GorushClient) ClearAlias(userId, appKey string) error {
	return nil
}

func (g *GorushClient) checkConfig(appKey string) error {
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
func (g *GorushClient) checkLang(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
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

func (g *GorushClient) PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	iotlogger.LogHelper.WithTag("mode", "gorush").Infof("GorushClient.PushMessage Gorush开始推送")
	g.InputTarget = inputTarget
	g.Message = message
	defer iotutil.PanicHandler()
	g.checkConfig(message.AppKey)
	g.checkLang(inputTarget, message)
	//token分组
	platformTokens := platforms.TokenGroup(inputTarget.PushTokens)

	//日志输出
	iotlogger.LogHelper.WithTag("mode", "gorush").Debugf(fmt.Sprintf("start manufacturer push, platform total: %v", len(inputTarget.PushTokens)))
	for p, tokens := range platformTokens {
		if len(tokens) > 0 {
			iotlogger.LogHelper.WithTag("mode", "gorush").Debugf(fmt.Sprintf("%v platform: push token count: %v", getPushPlatformName(p), len(tokens)))
		}
	}
	req := gorushclient.PushNotificationRequest{[]gorushclient.PushNotification{}}

	var msgInfo map[string]interface{}
	sendMessage := pushModel.MessageSendModel{}.SetModel(message)
	iotutil.StructToStruct(sendMessage, &msgInfo)

	for platform, items := range platformTokens {
		p := platforms.PlatformSelect(platform, g.Cfg, g.LangCt)
		if len(items) > 0 {
			p.GetCredentialInfo(msgInfo)
			for lang, tokens := range items {
				n := p.SetPushNotification(msgInfo, tokens, lang, nil)
				req.Notifications = append(req.Notifications, *n)
			}
		}
	}
	if len(req.Notifications) > 0 {
		err := g.GorushCli.SendPush(req)
		if err != nil {
			iotlogger.LogHelper.WithTag("mode", "gorush").Errorf("g.GorushCli.SendPush error: %v", err.Error())
			return err
		}
	}
	iotlogger.LogHelper.WithTag("mode", "gorush").Infof("GorushClient.PushMessage Gorush推送完成，通知数量：%v", len(req.Notifications))
	return nil
}

func getIntentUrl(pkgName string, platform int, data map[string]interface{}) string {
	return fmt.Sprintf("intent:#Intent;action=android.intent.action.MAIN;component=%v/%v.MainActivity;S.PushExtra=%v;end",
		pkgName, pkgName, iotutil.ToString(data),
	)
}

func checkPlatformTokens(platformTokens map[int]map[string][]string, lang string, appPushPlatform int) {
	if platformTokens[appPushPlatform] == nil {
		platformTokens[appPushPlatform] = make(map[string][]string)
	}
	if platformTokens[appPushPlatform][lang] == nil {
		platformTokens[appPushPlatform][lang] = make([]string, 0)
	}
}
