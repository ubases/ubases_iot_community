package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

type XiaomiPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *XiaomiPush) GetType() int {
	return gorushclient.PlatformXiaomi
}

func (s *XiaomiPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"miId":"2882303761520304506","miKey":"5172030463506","miSecret":"d/IfLeS/XTk8mDZsrQjj2g=="}
	miKey := iotutil.MapGetStringVal(s.Cfg.Xiaomi["miKey"], "")
	miSecret := iotutil.MapGetStringVal(s.Cfg.Xiaomi["miSecret"], "")
	miDeviceChanelId := iotutil.MapGetStringVal(s.Cfg.Xiaomi["miDevChanelId"], "")
	miAccountChanelId := iotutil.MapGetStringVal(s.Cfg.Xiaomi["miChannelId"], "")
	intentUrl := getIntentUrl(s.Cfg.IosPkgName, msgInfo)
	s.credental = &CredentalInfo{
		AppId:             miKey,
		AppSecret:         miSecret,
		IntentUrl:         intentUrl,
		PkgName:           s.Cfg.AndroidPkgName,
		MiDeviceChanelId:  miDeviceChanelId,
		MiAccountChanelId: miAccountChanelId,
	}
	return s.credental
}

func (s *XiaomiPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

	if len(tokens) == 0 {
		return nil
	}
	m, ok := s.LangCt[lang]
	if !ok {
		return nil
	}

	//TODO 根据消息的类型设置ChanelId

	n := gorushclient.PushNotification{
		Tokens:             tokens,
		Platform:           s.GetType(),
		Title:              m.Title, //消息标题
		Message:            m.Content,
		Data:               msgInfo,
		AppKey:             m.AppKey,
		PkgName:            s.credental.PkgName,
		Intent:             s.credental.IntentUrl,
		AppID:              s.credental.AppId, //兼容华为和其它平台写法
		AppSecret:          s.credental.AppSecret,
		XiaomiChanelId:     117741,
		XiaomiNotifyEffect: 2,
	}
	if s.credental.Credentials != "" {
		n.Credentials = s.credental.Credentials
	}
	return &n
}
