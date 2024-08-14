package platforms

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"cloud_platform/iot_message_service/service/push/gorushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"fmt"
)

type IPlatform interface {
	GetType() int
	GetCredentialInfo(msgInfo map[string]interface{}) *CredentalInfo
	SetPushNotification(msgInfo map[string]interface{}, tokens []string, lang string, commonFunc func()) *gorushclient.PushNotification
}

type CredentalInfo struct {
	Credentials      string
	CredentialSecret string
	AppId            string
	AppSecret        string
	ClientId         string
	ClientSecret     string
	IntentUrl        string
	PkgName          string

	//小米
	MiDeviceChanelId  string
	MiAccountChanelId string
	//oppo
	OppoChanelId   string
	OppoChanelName string
}

func PlatformSelect(platformType int, Cfg *config.JpushCfg, LangCt map[string]pushModel.MessageRequestModel) IPlatform {
	switch platformType {
	case gorushclient.PlatformHuawei:
		return &HuaweiPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformIos:
		return &IosPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformAndroid:
		return &AndroidPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformXiaomi:
		return &XiaomiPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformOppo:
		return &OppoPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformVivo:
		return &VivoPush{Cfg: Cfg, LangCt: LangCt}
	case gorushclient.PlatformHoner:
		return &HonorPush{Cfg: Cfg, LangCt: LangCt}
	}
	return nil
}

func getIntentUrl(pkgName string, data map[string]interface{}) string {
	return fmt.Sprintf("intent:#Intent;action=android.intent.action.MAIN;component=%v/%v.MainActivity;S.PushExtra=%v;end",
		pkgName, pkgName, iotutil.ToString(data),
	)
}

func checkPlatformTokens(platformTokens map[int]map[string][]string, lang string, appPushPlatform int) {
	if platformTokens[appPushPlatform] == nil {
		platformTokens[appPushPlatform] = make(map[string][]string)
	}
	if platformTokens[appPushPlatform][lang] == nil {
		platformTokens[appPushPlatform][lang] = make([]string, 0)
	}
}

func TokenGroup(tokens []pushModel.PushTokenItem) map[int]map[string][]string {
	platformTokens := make(map[int]map[string][]string)
	for _, token := range tokens {
		if token.Lang == "" {
			iotlogger.LogHelper.WithTag("mode", "gorush").Errorf("GorushClient.PushMessage 未设置语言")
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
			platformType := gorushclient.PlatformXiaomi
			checkPlatformTokens(platformTokens, token.Lang, platformType)
			platformTokens[platformType][token.Lang] = append(platformTokens[platformType][token.Lang], token.AppToken)
		case iotconst.PlatformOppo:
			platformType := gorushclient.PlatformOppo
			checkPlatformTokens(platformTokens, token.Lang, platformType)
			platformTokens[platformType][token.Lang] = append(platformTokens[platformType][token.Lang], token.AppToken)
		case iotconst.PlatformVivo:
			platformType := gorushclient.PlatformVivo
			checkPlatformTokens(platformTokens, token.Lang, platformType)
			platformTokens[platformType][token.Lang] = append(platformTokens[platformType][token.Lang], token.AppToken)
		case iotconst.PlatformHoner:
			platformType := gorushclient.PlatformHoner
			checkPlatformTokens(platformTokens, token.Lang, platformType)
			platformTokens[platformType][token.Lang] = append(platformTokens[platformType][token.Lang], token.AppToken)
		}
	}
	return platformTokens
}
