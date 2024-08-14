package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

//{"oppoId":"31644869","oppoKey":"ec97c242b206495c8b5802137da0c10a","oppoMasterSecret":"e60cbc8414054d11b7b9c06af19b858c","oppoSecret":"d105b3e71e794686b2307c867006cea4"}

type OppoPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *OppoPush) GetType() int {
	return gorushclient.PlatformOppo
}

func (s *OppoPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"oppoId":"31644869","oppoKey":"ec97c242b206495c8b5802137da0c10a","oppoMasterSecret":"e60cbc8414054d11b7b9c06af19b858c","oppoSecret":"d105b3e71e794686b2307c867006cea4"}
	oppoKey := iotutil.MapGetStringVal(s.Cfg.Oppo["oppoKey"], "")
	oppoSecret := iotutil.MapGetStringVal(s.Cfg.Oppo["oppoMasterSecret"], "")
	oppoChanelId := iotutil.MapGetStringVal(s.Cfg.Oppo["oppoChanelId"], "")
	oppoChanelName := iotutil.MapGetStringVal(s.Cfg.Oppo["oppoChanelName"], "")
	//intentUrl := getIntentUrl(s.Cfg.IosPkgName, msgInfo)
	s.credental = &CredentalInfo{
		AppId:          oppoKey,
		AppSecret:      oppoSecret,
		PkgName:        s.Cfg.AndroidPkgName,
		OppoChanelId:   oppoChanelId,
		OppoChanelName: oppoChanelName,
		//Credentials:      s.Cfg.Oppo["oppoSecret"],
		//CredentialSecret: s.Cfg.Oppo["oppoSecret"],
		//IntentUrl:        intentUrl,
	}
	return s.credental
}

func (s *OppoPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

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
		AppKey:    m.AppKey,
		Data:      map[string]interface{}{"PushExtra": msgInfo},
		Intent:    s.credental.IntentUrl,
		AppID:     s.credental.AppId,
		AppSecret: s.credental.AppSecret,
		PkgName:   s.credental.PkgName,
		ChannelId: "DeviceAccount",
	}
	if s.credental.Credentials != "" {
		n.Credentials = s.credental.Credentials
	}
	return &n
}
