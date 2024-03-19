package tianmao

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/config"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tidwall/gjson"
)

type TmDeviceReportSvc struct {
	SkillId         string `json:"skill_id"`
	MessageId       string `json:"message_id"`
	DeviceId        string `json:"device_id"`
	ReportType      int8   `json:"report_type"`
	Payload         string `json:"payload"`
	PayloadVersion  int8   `json:"payload_version"`
	AccountType     int8   `json:"account_type"`
	UserAccessToken string `json:"user_access_token"`
	Extension       string `json:"extension"`
	TimeStamp       int64  `json:"time_stamp"`
}

func RequestSync(devId string, userId string, onlineStatus interface{}, product protosService.OpmVoiceProduct) error {
	defer iotutil.PanicHandler(userId, devId)
	token := cached.RedisStore.GetClient().Get(context.Background(), fmt.Sprintf(iotconst.VoiceUserTokenKey, userId)).Val()
	if len(token) == 0 {
		return errors.New("用户语控token为空")
	}
	payload := map[string]interface{}{
		"onlinestate": onlineStatus,
	}
	plDatas, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	svc := &TmDeviceReportSvc{
		SkillId:         gjson.Get(product.VoiceOther, "voiceSkill").String(),
		MessageId:       uuid.New().String(),
		DeviceId:        devId,
		ReportType:      2,
		Payload:         string(plDatas),
		PayloadVersion:  2,
		AccountType:     1,
		UserAccessToken: token,
		TimeStamp:       time.Now().Unix(),
	}
	if err := svc.TmDeviceReport(); err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("device report svc data: ", svc)
	return nil
}

// 请求同步设备状态数据
func RequestSyncDeviceData(devId string, userId string, payload map[string]interface{}, product protosService.OpmVoiceProduct) error {
	token := cached.RedisStore.GetClient().Get(context.Background(), fmt.Sprintf(iotconst.VoiceUserTokenKey, userId)).Val()
	if len(token) == 0 {
		return errors.New("用户语控token为空")
	}
	plDatas, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	svc := &TmDeviceReportSvc{
		SkillId:         gjson.Get(product.VoiceOther, "voiceSkill").String(),
		MessageId:       uuid.New().String(),
		DeviceId:        devId,
		ReportType:      2,
		Payload:         string(plDatas),
		PayloadVersion:  2,
		AccountType:     1,
		UserAccessToken: token,
		TimeStamp:       time.Now().Unix(),
	}
	if err := svc.TmDeviceReport(); err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("device report svc data: ", svc)
	return nil
}

// 请求同步设备列表
func RequestSyncDeviceList(res iotstruct.DeviceRedisUpdate) error {
	defer iotutil.PanicHandler(res)
	userId, err := iotutil.ToInt64AndErr(res.UserId)
	if err != nil {
		return err
	}
	// 通过产品Key获取产品语控配置信息
	ctx := context.Background()
	var homeId string = res.HomeId
	if res.HomeId == "" {
		respUser, err := rpcclient.ClientUcUserService.FindById(ctx, &protosService.UcUserFilter{
			Id:     userId,
			Status: 1,
		})
		if err != nil {
			return err
		}
		homeId = respUser.Data[0].DefaultHomeId
	}
	homeDevList, err := rpcclient.ClientIotDeviceHomeService.Lists(ctx, &protosService.IotDeviceHomeListRequest{
		Query: &protosService.IotDeviceHome{
			HomeId: iotutil.ToInt64(homeId),
		},
	})
	if err != nil {
		return err
	}
	checked := map[string]struct{}{}
	for i := range homeDevList.Data {
		checked[homeDevList.Data[i].ProductKey] = struct{}{}
	}
	for k, _ := range checked {
		opmVoice, err := rpcclient.ClienOpmVoiceProductService.Find(ctx, &protosService.OpmVoiceProductFilter{
			ProductKey: k,
		})
		if err != nil {
			iotlogger.LogHelper.Helper.Error("open voice product service: ", err)
			continue
		}
		if len(opmVoice.Data) == 0 {
			iotlogger.LogHelper.Helper.Error("通过产品key获取语控配置信息为空")
			continue
		}
		token := cached.RedisStore.GetClient().Get(ctx, fmt.Sprintf(iotconst.VoiceUserTokenKey, res.UserId)).Val()
		if len(token) == 0 {
			iotlogger.LogHelper.Helper.Error("用户语控token为空")
			continue
		}
		svc := &TmDeviceReportSvc{
			SkillId:         gjson.Get(opmVoice.Data[0].VoiceOther, "voiceSkill").String(),
			UserAccessToken: token,
		}
		if err := svc.TmDeviceListUpdate(); err != nil {
			iotlogger.LogHelper.Helper.Error("更新天猫设备列表错误: ", err)
			continue
		}
		iotlogger.LogHelper.Helper.Debug("device report svc data: ", svc)
	}
	return nil
}

func (tr *TmDeviceReportSvc) TmDeviceReport() error {
	var v string = "2.0"
	var format string = "json"
	var signMethod string = "md5"
	var tmUrl = "https://eco.taobao.com/router/rest"
	var method string = "alibaba.ailabs.iot.cloud.device.report"

	appKey := config.Global.Aligenie.AppKey
	secret := config.Global.Aligenie.Secret
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	paramDatas, err := json.Marshal(*tr)
	if err != nil {
		return err
	}

	signParam := map[string]interface{}{
		"method":             method,
		"app_key":            appKey,
		"timestamp":          ts,
		"v":                  v,
		"sign_method":        signMethod,
		"format":             format,
		"cloud_report_param": string(paramDatas),
	}

	sign := TmSignToRequest(signParam, secret)

	dataUrlVal := url.Values{}
	dataUrlVal.Add("method", method)
	dataUrlVal.Add("app_key", appKey)
	dataUrlVal.Add("timestamp", ts)
	dataUrlVal.Add("v", v)
	dataUrlVal.Add("sign_method", signMethod)
	dataUrlVal.Add("format", format)
	dataUrlVal.Add("sign", sign)
	dataUrlVal.Add("cloud_report_param", string(paramDatas))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", tmUrl, strings.NewReader(dataUrlVal.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//resbody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}
	//iotlogger.LogHelper.Helper.Debug("向天猫精灵推送设备report报文响应: ", string(resbody))
	return nil
}

func (tr *TmDeviceReportSvc) TmDeviceListUpdate() error {
	var v string = "2.0"
	var typeVal string = "1"
	var format string = "json"
	var signMethod string = "md5"
	var tmUrl = "https://eco.taobao.com/router/rest"
	var method string = "alibaba.ailabs.iot.device.list.update.notify"

	appKey := config.Global.Aligenie.AppKey
	secret := config.Global.Aligenie.Secret
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	signParam := map[string]interface{}{
		"method":      method,
		"app_key":     appKey,
		"timestamp":   ts,
		"v":           v,
		"sign_method": signMethod,
		"format":      format,
		"token":       tr.UserAccessToken,
		"skill_id":    tr.SkillId,
		"type":        typeVal,
	}

	sign := TmSignToRequest(signParam, secret)

	dataUrlVal := url.Values{}
	dataUrlVal.Add("method", method)
	dataUrlVal.Add("app_key", appKey)
	dataUrlVal.Add("timestamp", ts)
	dataUrlVal.Add("v", v)
	dataUrlVal.Add("sign_method", signMethod)
	dataUrlVal.Add("format", format)
	dataUrlVal.Add("sign", sign)
	dataUrlVal.Add("token", tr.UserAccessToken)
	dataUrlVal.Add("skill_id", tr.SkillId)
	dataUrlVal.Add("type", typeVal)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", tmUrl, strings.NewReader(dataUrlVal.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("向天猫精灵推送设备列表更新响应: ", string(resbody))
	return nil
}

func TmSignToRequest(signParam map[string]interface{}, secret string) string {
	var ksort []string
	for k, _ := range signParam {
		ksort = append(ksort, k)
	}
	sort.Strings(ksort)
	strBuilder := secret
	for i := 0; i < len(ksort); i++ {
		tmp := iotutil.ToString(signParam[ksort[i]])
		if ksort[i] != "" && tmp != "" {
			strBuilder += ksort[i] + tmp
		}
	}
	strBuilder += secret
	return strings.ToUpper(iotutil.EncodeMD5(strBuilder))
}
