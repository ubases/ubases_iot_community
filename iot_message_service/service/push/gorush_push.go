package push

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
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
	//TODO implement me
	panic("implement me")
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
	iotlogger.LogHelper.Infof("GorushClient.PushMessage Gorush开始推送")
	g.InputTarget = inputTarget
	g.Message = message
	defer iotutil.PanicHandler()
	g.checkConfig(message.AppKey)
	g.checkLang(inputTarget, message)
	platformTokens := make(map[int]map[string][]string)
	for _, token := range inputTarget.PushTokens {
		if token.Lang == "" {
			//token.Lang = "zh"
			iotlogger.LogHelper.Errorf("GorushClient.PushMessage 未设置语言")
			continue
		}
		//初始化
		appPushPlatform, _ := iotutil.ToIntErr(token.AppPushPlatform)
		switch appPushPlatform {
		case iotconst.PlatformHuawei:
			checkPlatformTokens(platformTokens, token.Lang, gorushclient.PlatformHuawei)
			platformTokens[gorushclient.PlatformHuawei][token.Lang] = append(platformTokens[gorushclient.PlatformHuawei][token.Lang], token.AppToken)
		case iotconst.PlatformIos:
			checkPlatformTokens(platformTokens, token.Lang, gorushclient.PlatformIos)
			platformTokens[gorushclient.PlatformIos][token.Lang] = append(platformTokens[gorushclient.PlatformIos][token.Lang], token.AppToken)
		case iotconst.PlatformAndroid:
			checkPlatformTokens(platformTokens, token.Lang, gorushclient.PlatformAndroid)
			platformTokens[gorushclient.PlatformAndroid][token.Lang] = append(platformTokens[gorushclient.PlatformAndroid][token.Lang], token.AppToken)
		case iotconst.PlatformXiaomi:
			//platformTokens[gorushclient.PlatformXiaomi][token.Lang] = append(platformTokens[gorushclient.PlatformXiaomi][token.Lang], token.AppToken)
		case iotconst.PlatformOppo:
			//platformTokens[gorushclient.PlatformOppo][token.Lang] = append(platformTokens[gorushclient.PlatformOppo][token.Lang], token.AppToken)
		case iotconst.PlatformVivo:
			//platformTokens[gorushclient.PlatformVivo][token.Lang] = append(platformTokens[gorushclient.PlatformVivo][token.Lang], token.AppToken)
		case iotconst.PlatformHoner:
			//platformTokens[gorushclient.PlatformHoner][token.Lang] = append(platformTokens[gorushclient.PlatformHoner][token.Lang], token.AppToken)
		case iotconst.PlatformMeizu:
			//platformTokens[gorushclient.PlatformMeizu][token.Lang] = append(platformTokens[gorushclient.PlatformMeizu][token.Lang], token.AppToken)
		}
	}
	req := gorushclient.PushNotificationRequest{[]gorushclient.PushNotification{}}

	var msgInfo map[string]interface{}
	sendMessage := pushModel.MessageSendModel{}.SetModel(message)
	iotutil.StructToStruct(sendMessage, &msgInfo)

	for platform, items := range platformTokens {
		if len(items) > 0 {
			credentials, credentialSecret := "", ""
			if platform == gorushclient.PlatformAndroid {
				//fcmJson: -credentials "D:\code\work\test\go\golang_demo\fcm\push.json"
				credentials = iotutil.MapGetStringVal(g.Cfg.Fcm["fcmJsonContent"], "")
			} else if platform == gorushclient.PlatformIos {
				credentials = iotutil.MapGetStringVal(g.Cfg.Apns["apnsCert"], "")
				credentialSecret = iotutil.MapGetStringVal(g.Cfg.Apns["apnsSecret"], "")
			}
			for lang, tokens := range items {
				if len(tokens) == 0 {
					continue
				}
				m, ok := g.LangCt[lang]
				if !ok {
					continue
				}
				n := gorushclient.PushNotification{
					Tokens:   tokens,
					Platform: platform,
					//Title:    m.Title, //消息他标题
					Message:          m.Content,
					Data:             msgInfo,
					Topic:            g.Cfg.IosPkgName,
					AppKey:           m.AppKey,
					CredentialSecret: credentialSecret,
				}
				if credentials != "" {
					n.Credentials = credentials
				}
				req.Notifications = append(req.Notifications, n)
			}
		}
	}
	if len(req.Notifications) > 0 {
		err := g.GorushCli.SendPush(req)
		if err != nil {
			iotlogger.LogHelper.Errorf("g.GorushCli.SendPush error: %v", err.Error())
			return err
		}
	}
	iotlogger.LogHelper.Infof("GorushClient.PushMessage Gorush推送完成，通知数量：%v", len(req.Notifications))
	return nil
}

func checkPlatformTokens(platformTokens map[int]map[string][]string, lang string, appPushPlatform int) {
	if platformTokens[appPushPlatform] == nil {
		platformTokens[appPushPlatform] = make(map[string][]string)
	}
	if platformTokens[appPushPlatform][lang] == nil {
		platformTokens[appPushPlatform][lang] = make([]string, 0)
	}
}
