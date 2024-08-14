/**
 * @Author: hogan
 * @Date: 2021/11/17 16:28
 */

package push

import (
	"bytes"
	"cloud_platform/iot_message_service/config"

	"io/ioutil"
	"net/http"
	"text/template"

	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/rpc/rpcclient"
	"cloud_platform/iot_message_service/service/push/jpushclient"
	"cloud_platform/iot_message_service/service/push/pushModel"
	"cloud_platform/iot_message_service/service/push/translate"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type PushMessage struct {
}

// 从字典获取
var langs = []string{"zh", "en"}

// 给目标设置语言
func setLang(lang string, arr []string) []string {
	for i := 0; i < len(arr); i++ {
		arr[i] = fmt.Sprintf("%s_%s", arr[i], lang)
	}
	return arr
}

// 将所有别名转换为多语言版本
func setLangAll(arr []string) []string {
	var newLangs = []string{}
	for i := 0; i < len(arr); i++ {
		for _, lang := range langs {
			newLangs = append(newLangs, fmt.Sprintf("%s_%s", arr[i], lang))
		}
	}
	return newLangs
}

// 内容值设置，将插槽替换为值
func setValues(str string, params []interface{}) string {
	for i := 0; i < len(params); i++ {
		str = strings.Replace(str, "[0]", params[i].(string), 1)
	}
	return str
}

func setValuesExtand(str string, params []string) string {
	for i := 0; i < len(params); i++ {
		str = strings.Replace(str, "[0]", params[i], 1)
	}
	return str
}

// 从文本内容加载
func paramIntoContent(templateContent string, params interface{}) (string, error) {
	tmp, err := template.New("TemplateContent").Parse(templateContent)
	if err != nil {
		return templateContent, err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, params); err != nil {
		return templateContent, err
	}
	return buf.String(), nil
}

// 设置翻译
func setTranslate(lang string, message *pushModel.MessageRequestModel) {
	//todo 初始化Translate
	title, content := translate.Get(lang, message.TplCode, message.Content)
	if message.Title != message.TplCode {
		message.Title = title
	}
	message.Content, _ = paramIntoContent(content, message.Params)
}

func getJsPushConfig(appKey string) (*config.JpushCfg, error) {
	if appKey == "" {
		return nil, errors.New("appKey is null")
	}
	jpush, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_APPPUSH_DATA_PREFIX+appKey).Result()
	if err != nil {
		jpush = make(map[string]string)
		//return nil
	}
	if jpush["jpushKey"] == "" || jpush["jpushSecret"] == "" || jpush["androidPkgName"] == "" {
		oemAppResult, err := rpcclient.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
			AppKey: appKey,
		})
		if err != nil {
			iotlogger.LogHelper.Errorf("ClientOemAppService error %v", err.Error())
			return nil, errors.New("oemAppResult is null")
		}
		if oemAppResult.Code != 200 {
			iotlogger.LogHelper.Errorf("ClientOemAppService code!=200 %v", oemAppResult.Message)
			return nil, errors.New("oemAppResult is null")
		}
		oemAppInfo := oemAppResult.Data[0]
		oemAppPushCert, err := rpcclient.ClientOemAppPushCertService.Find(context.Background(), &protosService.OemAppPushCertFilter{
			AppId:   oemAppInfo.Id,
			Version: oemAppInfo.Version,
		})
		if err != nil {
			iotlogger.LogHelper.Errorf("oemAppPushCert error %v", err.Error())
			return nil, errors.New("oemAppPushCert is null")
		}
		if oemAppPushCert.Code != 200 {
			iotlogger.LogHelper.Errorf(oemAppPushCert.Message)
			return nil, errors.New("oemAppPushCert is null")
		}

		if oemAppPushCert.Data[0].Jpush == "" {
			return nil, errors.New("jpush is null")
		}
		jpushMap := iotutil.JsonToMap(oemAppPushCert.Data[0].Jpush)
		//存储jpush redis
		jpush["androidPkgName"] = oemAppInfo.AndroidPkgName
		jpush["jpushKey"] = iotutil.ToString(jpushMap["jpushKey"])
		jpush["jpushSecret"] = iotutil.ToString(jpushMap["jpushSecret"])
		jpush["androidPkgName"] = oemAppInfo.AndroidPkgName
		jpush["iosPkgName"] = oemAppInfo.IosPkgName

		//推送参数配置
		//fcm: {"fcmId":"44","fcmJson":"https://osspublic.aithinker.com/app/14ab859e-9108-441d-9a1c-7cbd5cdd57f1.json","fcmKey":"33"}
		//apns: {"apnsCert":"https://osspublic.aithinker.com/app/1bd99cc0-e445-46d5-9b87-57c45de56ae6.p12","apnsSecret":"123456","appId":"9052313536742457344","version":"1.0.1"}
		//jpush: {"jpushKey":"fc81299b7e8aa7c8e555784b","jpushSecret":"ffdc934e5927ca7780158f4d"}
		//huawei: {"huaweiClientSecret":"","huaweiId":"","huaweiJson":"","huaweiSecret":""}
		//xiaomi: {"miId":"123456","miKey":"123456","miSecret":"123456"}
		//vivo: {"vivoId":"","vivoKey":"","vivoSecret":""}
		//oppo: {"oppoId":"","oppoKey":"","oppoMasterSecret":"","oppoSecret":""}
		jpush["apns"] = oemAppPushCert.Data[0].Apns
		jpush["jpush"] = oemAppPushCert.Data[0].Jpush
		jpush["fcm"] = oemAppPushCert.Data[0].Fcm
		jpush["huawei"] = oemAppPushCert.Data[0].Huawei
		jpush["xiaomi"] = oemAppPushCert.Data[0].Xiaomi
		jpush["vivo"] = oemAppPushCert.Data[0].Vivo
		jpush["oppo"] = oemAppPushCert.Data[0].Oppo
		jpush["honor"] = oemAppPushCert.Data[0].Honor

		appCachedCmd := iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_APPPUSH_DATA_PREFIX+appKey, jpush)
		if appCachedCmd.Err() != nil {
			iotlogger.LogHelper.Errorf("app推送数据缓存异常，%s", appCachedCmd.Err().Error())
		}
	}
	iotlogger.LogHelper.Info("极光参数：", iotutil.ToString(jpush))
	apnsProduction := true
	if jpush["apnsProduction"] != "" {
		iotlogger.LogHelper.Info("极光参数：", iotutil.ToString(jpush))
		apnsProduction = jpush["apnsProduction"] == "true"
	}
	return &config.JpushCfg{
		AppKey:         jpush["jpushKey"],
		Secret:         jpush["jpushSecret"],
		AndroidPkgName: jpush["androidPkgName"],
		IosPkgName:     jpush["iosPkgName"],
		ApnsProduction: apnsProduction,
		Jpush:          iotutil.ToMap(jpush["jpush"]),
		Apns:           iotutil.ToMap(jpush["apns"]),
		Fcm:            getFcmJson(jpush["fcm"]),
		Huawei:         iotutil.ToMap(jpush["huawei"]),
		Xiaomi:         iotutil.ToMap(jpush["xiaomi"]),
		Vivo:           iotutil.ToMap(jpush["vivo"]),
		Oppo:           iotutil.ToMap(jpush["oppo"]),
		Honor:          iotutil.ToMap(jpush["honor"]),
	}, nil
}

func getFcmJson(s string) map[string]interface{} {
	fcm := make(map[string]interface{})
	if s == "" {
		return fcm
	}
	//{"fcmId":"44","fcmJson":"https://osspublic.aithinker.com/app/14ab859e-9108-441d-9a1c-7cbd5cdd57f1.json","fcmKey":"33"}
	fcm = iotutil.ToMap(s)
	if fcm == nil {
		return fcm
	}
	if url, ok := fcm["fcmServerJson"]; ok {
		// 发起HTTP GET请求
		response, err := http.Get(url.(string))
		if err != nil {
			fmt.Println("Failed to send HTTP request:", err)
			return fcm
		}
		defer response.Body.Close()

		// 检查HTTP响应状态码
		if response.StatusCode != http.StatusOK {
			fmt.Println("HTTP request failed with status code:", response.StatusCode)
			return fcm
		}

		// 读取响应体
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return fcm
		}
		// 打印文件内容
		fcm["fcmJsonContent"] = string(content)
	}
	return fcm
}

func (s *PushMessage) PushMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	return s.SendMessage(inputTarget, message)
}

// 发送极光消息，数据转换翻译处理
func (s *PushMessage) SendMessage(inputTarget pushModel.MessageTarget, message pushModel.MessageRequestModel) error {
	cfg, configErr := getJsPushConfig(message.AppKey)
	if configErr != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", configErr.Error()))
		return configErr
	}
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID) //android平台
	pf.Add(jpushclient.IOS)     //IOS平台
	//pf.Add(jpushclient.WINPHONE)	//WinPhone平台
	//pf.All()						//所有平台
	if inputTarget.IsPublic {
		iotlogger.LogHelper.Info("广播消息")
		var ad jpushclient.Audience
		ad.All()
		RunSend(pf, ad, inputTarget, message, cfg)
		return nil
	} else {
		//指定推送
		langs = []string{"zh", "en"} //排查为什么没有推送en
		for _, lang := range langs {
			//TODO 这里存在复制的问题。
			target := pushModel.MessageTarget{}
			iotutil.StructToStruct(inputTarget, &target)
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
				return errors.New("未设置任何目标对象")
			}
			//复制对象，防止中英文翻译问题
			var newMessage pushModel.MessageRequestModel
			iotutil.StructToStruct(message, &newMessage)
			//翻译处理
			setTranslate(lang, &newMessage)
			//执行推送逻辑
			RunSend(pf, ad, inputTarget, newMessage, cfg)
		}
	}
	return nil
}

// 执行推送
func RunSend(pf jpushclient.Platform, ad jpushclient.Audience, inputTarget pushModel.MessageTarget,
	message pushModel.MessageRequestModel, cfg *config.JpushCfg) error {
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
		"url": fmt.Sprintf("intent:#Intent;action=android.intent.action.MAIN;component=%s/%s.MainActivity;end", cfg.AndroidPkgName, cfg.AndroidPkgName),
	}
	notice.Android.UriAction = "android.intent.action.MAIN"
	notice.Android.UriActivity = fmt.Sprintf("%s.MainActivity", cfg.AndroidPkgName)

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
	payload.SetPlatform(&pf)
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
	options := &jpushclient.Option{ApnsProduction: cfg.ApnsProduction}
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
	//jpush := jpushclient.NewPushClient(config.Global.Jpush.Secret, config.Global.Jpush.AppKey)
	jpush := jpushclient.NewPushClient(cfg.Secret, cfg.AppKey)
	str, err := jpush.Send(bytes)
	if err != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
	} else {
		iotlogger.LogHelper.Info(fmt.Sprintf("ok:%s", str))
	}
	return err
}

func (s *PushMessage) ClearAlias(userId, appKey string) error {
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
