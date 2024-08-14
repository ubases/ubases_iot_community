package platforms

import (
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
)

type HuaweiPush struct {
	Cfg       *config.JpushCfg
	LangCt    map[string]pushModel.MessageRequestModel
	credental *CredentalInfo
}

func (s *HuaweiPush) GetType() int {
	return gorushclient.PlatformHuawei
}

func (s *HuaweiPush) GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo {
	//{"huaweiClientSecret":"","huaweiId":"","huaweiJson":"","huaweiSecret":""}
	huaweiId := iotutil.MapGetStringVal(s.Cfg.Huawei["huaweiId"], "")
	huaweiSecret := iotutil.MapGetStringVal(s.Cfg.Huawei["huaweiSecret"], "")
	intentUrl := getIntentUrl(s.Cfg.IosPkgName, msgInfo)
	s.credental = &CredentalInfo{
		AppId:     huaweiId,
		AppSecret: huaweiSecret,
		IntentUrl: intentUrl,
		PkgName:   s.Cfg.AndroidPkgName,
	}
	return s.credental
}

func (s *HuaweiPush) SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification {

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
		Intent:    s.credental.IntentUrl,
		AppID:     s.credental.AppId, //兼容华为和其它平台写法
		AppSecret: s.credental.AppSecret,
		PkgName:   s.credental.PkgName,
		ChannelId: "DeviceAccount",
	}
	//设备消息：DEVICE_REMINDER  账号：ACCOUNT
	if m.Type == "device" {
		n.Category = "DEVICE_REMINDER"
	} else {
		n.Category = "ACCOUNT"
	}
	return &n
}
