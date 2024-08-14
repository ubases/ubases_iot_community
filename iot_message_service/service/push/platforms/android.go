package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

type AndroidPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *AndroidPush) GetType() int {
	return gorushclient.PlatformAndroid
}

func (s *AndroidPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"fcmId":"","fcmJson":"","fcmKey":"","fcmServerJson":""}
	//fcmJson: -credentials "D:\code\work\test\go\golang_demo\fcm\push.json"
	credentials := iotutil.MapGetStringVal(s.Cfg.Fcm["fcmJsonContent"], "")
	s.credental = &CredentalInfo{
		Credentials: credentials,
		PkgName:     s.Cfg.AndroidPkgName,
	}
	return s.credental
}

func (s *AndroidPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {
	if len(tokens) == 0 {
		return nil
	}
	m, ok := s.LangCt[lang]
	if !ok {
		return nil
	}
	n := gorushclient.PushNotification{
		Tokens:      tokens,
		Platform:    s.GetType(),
		Title:       m.Title, //消息标题
		Message:     m.Content,
		Data:        map[string]interface{}{"PushExtra": msgInfo},
		AppKey:      m.AppKey,
		Topic:       s.credental.PkgName,
		PkgName:     s.credental.PkgName,
		Credentials: s.credental.Credentials,
		ChannelId:   "DeviceAccount",
	}
	return &n
}
