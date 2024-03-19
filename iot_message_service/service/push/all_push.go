package push

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/jpushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"encoding/json"
	"fmt"
)

type AllClient struct {
	Cfg *config.JpushCfg
}

// 推送消息
func (g AllClient) PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	cfg, configErr := getJsPushConfig(message.AppKey)
	if configErr != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", configErr.Error()))
		//return configErr
	}
	langCt := map[string]pushModel.MessageRequestModel{}
	if inputTarget.IsPublic {
		//公共消息无法进行翻译，直接推送
	} else {
		//处理翻译（中文、英文）
		for _, lang := range langs {
			var m pushModel.MessageRequestModel
			iotutil.StructToStruct(message, &m)
			setTranslate(lang, &m)
			langCt[lang] = m
		}
	}
	//gorush厂商推送
	if inputTarget.PushTokens != nil && len(inputTarget.PushTokens) > 0 {
		gorushCli := &GorushClient{Cfg: cfg, LangCt: langCt, GorushCli: gorushclient.NewGorushClient(config.Global.Gorush.Url)}
		go gorushCli.PushMessage(inputTarget, message)
	}
	//极光推送
	jpushCli := &JPushClient{Cfg: cfg, LangCt: langCt}
	go jpushCli.PushMessage(inputTarget, message)
	return nil
}

func (g AllClient) ClearAlias(userId, appKey string) error {
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
