package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

//{"vivoId":"105733260","vivoKey":"a51a6f773749bc24e488fb25e336da19","vivoSecret":"fc19f82b-89ee-44fb-b35a-0eea63aa0672"}

type VivoPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *VivoPush) GetType() int {
	return gorushclient.PlatformVivo
}

func (s *VivoPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"vivoId":"105733260","vivoKey":"a51a6f773749bc24e488fb25e336da19","vivoSecret":"fc19f82b-89ee-44fb-b35a-0eea63aa0672"}
	vivoId := iotutil.MapGetStringVal(s.Cfg.Vivo["vivoId"], "")
	vivoKey := iotutil.MapGetStringVal(s.Cfg.Vivo["vivoKey"], "")
	vivoSecret := iotutil.MapGetStringVal(s.Cfg.Vivo["vivoSecret"], "")
	intentUrl := getIntentUrl(s.Cfg.IosPkgName, msgInfo)
	s.credental = &CredentalInfo{
		ClientId:  vivoId,
		AppId:     vivoKey,
		AppSecret: vivoSecret,
		IntentUrl: intentUrl,
		PkgName:   s.Cfg.AndroidPkgName,
	}
	return s.credental
}

func (s *VivoPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

	if len(tokens) == 0 {
		return nil
	}
	m, ok := s.LangCt[lang]
	if !ok {
		return nil
	}
	n := gorushclient.PushNotification{
		ID:        iotutil.UUID(),
		Tokens:    tokens,
		Platform:  s.GetType(),
		Title:     m.Title, //消息标题
		Message:   m.Content,
		AppKey:    m.AppKey,
		Intent:    s.credental.IntentUrl,
		AppID:     s.credental.AppId,
		ClientID:  s.credental.ClientId,
		AppSecret: s.credental.AppSecret,
		PkgName:   s.credental.PkgName,
		Category:  "DEVICE_REMINDER",
		ChannelId: "DeviceAccount",
	}
	return &n
}
