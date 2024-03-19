package push

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/jpushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"encoding/json"
	"errors"
	"fmt"
)

type JPushClient struct {
	Cfg         *config.JpushCfg
	LangCt      map[string]pushModel.MessageRequestModel
	inputTarget pushModel.MessageTarget
	message     pushModel.MessageRequestModel
	pf          *jpushclient.Platform
}

// 推送公共消息
func (j *JPushClient) pushPublicMsg(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	if j.pf == nil {
		j.setJPushClient()
	}
	iotlogger.LogHelper.Info("广播消息")
	var ad jpushclient.Audience
	ad.All()
	j.runJPush(ad, message)
	return nil
}

// 设置推送平台
func (j *JPushClient) setJPushClient() {
	j.pf = &jpushclient.Platform{}
	j.pf.Add(jpushclient.ANDROID) //android平台
	j.pf.Add(jpushclient.IOS)     //IOS平台
	//pf.Add(jpushclient.WINPHONE)	//WinPhone平台
	//pf.All()						//所有平台
}

// 检查推送目录
func (j *JPushClient) checkPushAudience(lang string) (jpushclient.Audience, error) {
	target := pushModel.MessageTarget{}
	iotutil.StructToStruct(j.inputTarget, &target)
	//设置推送对象
	var ad jpushclient.Audience
	iotlogger.LogHelper.Info("指定用户推送消息", lang)
	hasTarget := false
	if target.Tags != nil && len(target.Tags) > 0 {
		hasTarget = true
		ad.SetTag(setLang(lang, target.Tags)) //标签
	}
	if target.Alias != nil && len(target.Alias) > 0 {
		hasTarget = true
		ad.SetAlias(setLang(lang, target.Alias)) //别名
	}
	if target.RegIds != nil && len(target.RegIds) > 0 {
		hasTarget = true
		ad.SetID(setLang(lang, target.RegIds)) //regid
	}
	if !hasTarget {
		iotlogger.LogHelper.Error("没有设置tags")
		return ad, errors.New("未设置任何目标对象")
	}
	return ad, nil
}

// 推送消息
func (j *JPushClient) PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	iotlogger.LogHelper.Infof("JPushClient.PushMessage 极光开始推送，inputTarget: %v", iotutil.ToString(inputTarget))
	j.inputTarget = inputTarget
	j.message = message
	defer iotutil.PanicHandler()
	j.checkConfig(message.AppKey)
	j.setLang()
	j.setJPushClient()

	if inputTarget.IsPublic {
		return j.pushPublicMsg(inputTarget, message)
	} else {
		for _, lang := range langs {
			ad, err := j.checkPushAudience(lang)
			if err != nil {
				continue
			}
			//执行推送逻辑
			j.runJPush(ad, j.LangCt[lang])
		}
	}

	iotlogger.LogHelper.Infof("JPushClient.PushMessage 极光推送完成")
	return nil
}

// 清理标签
func (J JPushClient) ClearAlias(userId, appKey string) error {
	cfg, configErr := getJsPushConfig(appKey)
	if configErr != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", configErr.Error()))
		return configErr
	}
	jpush := jpushclient.NewPushClient(cfg.Secret, cfg.AppKey)
	for _, lang := range langs {
		aliasName := fmt.Sprintf("%s_%s", userId, lang)
		iotlogger.LogHelper.Info("clear alias ", aliasName)
		//查询别名
		resStr, err := jpush.SendGetAliasesRequest(aliasName)
		if err != nil {
			iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
			continue
		}
		var res jpushclient.AliasesResponse
		err = json.Unmarshal([]byte(resStr), &res)
		if err != nil {
			iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
			continue
		}
		if res.Data != nil {
			if len(res.Data) > 9 {
				newData := res.Data[8:len(res.Data)]
				if len(newData) == 0 {
					continue
				}
				//清除别名
				alias := jpushclient.NewAliases()
				for _, d := range newData {
					alias.SetRegIds(d.RegistrationId)
				}
				bytes, _ := alias.ToBytes()
				jpush.BaseUrl = fmt.Sprintf("%s/%s", jpushclient.HOST_ALIASES_SSL, aliasName)
				jpush.Send(bytes)
			}
		}
	}
	return nil
}

// 检查配置文件
func (j *JPushClient) checkConfig(appKey string) error {
	var err error
	if j.Cfg == nil {
		j.Cfg, err = getJsPushConfig(appKey)
		if err != nil {
			iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
			return err
		}
	}
	return nil
}

// 消息内容转换为语言
func (j *JPushClient) setLang() error {
	if j.LangCt == nil {
		j.LangCt = map[string]pushModel.MessageRequestModel{}
		if j.inputTarget.IsPublic {
			//公共消息无法进行翻译，直接推送
		} else {
			//处理翻译（中文、英文）
			for _, lang := range langs {
				var m pushModel.MessageRequestModel
				iotutil.StructToStruct(j.message, &m)
				setTranslate(lang, &m)
				j.LangCt[lang] = m
			}
		}
	}
	return nil
}

// 执行推送
func (j *JPushClient) runJPush(ad jpushclient.Audience, message pushModel.MessageRequestModel) error {
	defer iotutil.PanicHandler()
	//Alert消息内容
	alertMsg := fmt.Sprintf("%s", message.Title)
	iotlogger.LogHelper.Info("发送内容：" + alertMsg)
	var msgInfo map[string]interface{}
	sendMessage := pushModel.MessageSendModel{}.SetModel(message)
	iotutil.StructToStruct(sendMessage, &msgInfo)

	//Notice
	var notice jpushclient.Notice
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: alertMsg})
	notice.Android.Extras = msgInfo

	//cfg.AndroidPacketName
	notice.Android.Intent = map[string]interface{}{
		"url": fmt.Sprintf("intent:#Intent;action=android.intent.action.MAIN;component=%s/%s.MainActivity;end", j.Cfg.AndroidPkgName, j.Cfg.AndroidPkgName),
	}
	notice.Android.UriAction = "android.intent.action.MAIN"
	notice.Android.UriActivity = fmt.Sprintf("%s.MainActivity", j.Cfg.AndroidPkgName)

	if sendMessage.Sound != "" {
		notice.Android.Sound = sendMessage.Sound
	}
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: alertMsg})
	notice.IOS.Extras = msgInfo
	notice.IOS.Sound = "default"
	if sendMessage.Sound != "" {
		notice.IOS.Sound = sendMessage.Sound + ".mp3"
	}
	notice.IOS.ContentAvailable = false

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(j.pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)
	//
	//如果需要message，也就是收到消息，自动调用app的js实现页面调整
	//if message.ChildType == "2" {
	//	var msg jpushclient.Message
	//	msg.Title = message.Title
	//	msg.Content = message.Content
	//	msg.Extras = msgInfo
	//	payload.SetMessage(&msg)
	//}
	//参数配置，上线配置ApnsProduction
	//options := &jpushclient.Option{ApnsProduction: config.Global.Jpush.ApnsProduction}
	options := &jpushclient.Option{ApnsProduction: j.Cfg.ApnsProduction}
	//消息默认为一小时
	//如果不设置过期时间，则为60分钟，如果设置了过期时间，则消息为3天
	//推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天），
	//普通用户最长 3天， VIP 用户最长 10天。设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到。该字段对 iOS 的 Notification 消息无效。
	if message.UnSetExpire {
		options.SetTimelive(3600)
	} else {
		options.SetTimelive(86400 * 3)
	}

	payload.SetOptions(options)

	bytes, _ := payload.ToBytes()
	strJson := string(bytes)

	iotlogger.LogHelper.Info(fmt.Sprintf("%s\r\n", strJson))
	//推送消息
	jpush := jpushclient.NewPushClient(j.Cfg.Secret, j.Cfg.AppKey)
	str, err := jpush.Send(bytes)
	if err != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("JPushClient.runJPush err:%s", err.Error()))
	} else {
		iotlogger.LogHelper.Info(fmt.Sprintf("JPushClient.runJPush ok:%s", str))
	}
	return err
}
