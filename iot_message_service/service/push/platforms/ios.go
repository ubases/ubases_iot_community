package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

type IosPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *IosPush) GetType() int {
	return gorushclient.PlatformIos
}

func (s *IosPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"apnsCert":"https://osspublic.iot-aithings.com/app/57ab90cd-a819-4c00-90db-7ff20c274f1f.p12","apnsSecret":"123456","appId":"9052313536742457344","version":"1.0.5"}
	credentials := iotutil.MapGetStringVal(s.Cfg.Apns["apnsCert"], "")
	credentialSecret := iotutil.MapGetStringVal(s.Cfg.Apns["apnsSecret"], "")
	s.credental = &CredentalInfo{
		Credentials:      credentials,
		CredentialSecret: credentialSecret,
		PkgName:          s.Cfg.IosPkgName,
	}
	return s.credental
}

func (s *IosPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

	if len(tokens) == 0 {
		return nil
	}
	m, ok := s.LangCt[lang]
	if !ok {
		return nil
	}
	n := gorushclient.PushNotification{
		Tokens:    tokens,
		Platform:  s.GetType(),
		Title:     m.Title, //消息标题
		Message:   m.Content,
		Data:      msgInfo,
		AppKey:    m.AppKey,
		Topic:     s.credental.PkgName,
		AppID:     s.credental.Credentials, //兼容华为和其它平台写法
		AppSecret: s.credental.CredentialSecret,
	}
	if s.credental.Credentials != "" {
		n.Credentials = s.credental.Credentials
	}
	return &n
}
