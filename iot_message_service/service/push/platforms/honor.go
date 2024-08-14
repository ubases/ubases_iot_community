package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

//{"oppoId":"31644869","oppoKey":"ec97c242b206495c8b5802137da0c10a","oppoMasterSecret":"e60cbc8414054d11b7b9c06af19b858c","oppoSecret":"d105b3e71e794686b2307c867006cea4"}

type HonorPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *HonorPush) GetType() int {
	return gorushclient.PlatformHoner
}

func (s *HonorPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"honorAppId":"31644869","honorAppSecret":"ec97c242b206495c8b5802137da0c10a","honorClientId":"e60cbc8414054d11b7b9c06af19b858c","honorClientSecret":"d105b3e71e794686b2307c867006cea4"}
	honorAppId := iotutil.MapGetStringVal(s.Cfg.Honor["honorAppId"], "")
	honorAppSecret := iotutil.MapGetStringVal(s.Cfg.Honor["honorAppSecret"], "")
	honorClientId := iotutil.MapGetStringVal(s.Cfg.Honor["honorClientId"], "")
	honorClientSecret := iotutil.MapGetStringVal(s.Cfg.Honor["honorClientSecret"], "")
	intentUrl := getIntentUrl(s.Cfg.IosPkgName, msgInfo)
	s.credental = &CredentalInfo{
		AppId:        honorAppId,
		AppSecret:    honorAppSecret,
		ClientId:     honorClientId,
		ClientSecret: honorClientSecret,
		IntentUrl:    intentUrl,
		PkgName:      s.Cfg.AndroidPkgName,
	}
	return s.credental
}

func (s *HonorPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

	if len(tokens) == 0 {
		return nil
	}
	m, ok := s.LangCt[lang]
	if !ok {
		return nil
	}
	n := gorushclient.PushNotification{
		Tokens:       tokens,
		Platform:     s.GetType(),
		Title:        m.Title, //消息标题
		Message:      m.Content,
		AppKey:       m.AppKey,
		Intent:       s.credental.IntentUrl,
		AppID:        s.credental.AppId,
		AppSecret:    s.credental.AppSecret,
		ClientID:     s.credental.ClientId,
		ClientSecret: s.credental.ClientSecret,
		PkgName:      s.credental.PkgName,
		ChannelId:    "DeviceAccount",
	}
	return &n
}
