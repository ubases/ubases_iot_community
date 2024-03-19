package alexa

import (
	"bytes"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func RequestSync(userId, deviceId string, payload map[string]interface{}) error {
	defer iotutil.PanicHandler(userId, deviceId)
	tokenInfo, err := getCachedAlexaTokenInfo(userId)
	if err != nil {
		return err
	}
	//token := tokenInfo["token"]
	alexaToken := tokenInfo["access_token"]
	correlationToken := tokenInfo["correlationToken"]
	//alexaAuthTokenUrl := tokenInfo["alexaAuthTokenUrl"]
	alexaEventUrl := tokenInfo["alexaEventUrl"]

	//发现设备
	devices, err := common.DiscoveryDevices(userId, "alexa")
	if err != nil {
		return err
	}

	miDevs := make([]entitys.XiaomiSubscribe, 0)
	for _, device := range devices {
		if deviceId != "" && device.DeviceId != deviceId {
			continue
		}
		miDevs = append(miDevs, entitys.XiaomiSubscribe{
			Did:            device.DeviceId,
			SubscriptionId: device.SubscriptionId,
		})
		properties := []entitys.AlexaControlProperties{}
		devInfo, err := getDeviceInfo(device.DeviceId)
		if err != nil {
			break
		}
		for _, v := range device.VoiceProduct.FunctionMap {
			interfaceStr, instanceStr := "", ""
			arr := strings.Split(v.VoiceCode, "#")
			interfaceStr = arr[0]
			if len(arr) > 1 {
				instanceStr = arr[1]
			}
			if instanceStr == "" {
				instanceStr = v.AttrCode
			}

			var attrVal map[string]string = make(map[string]string)
			attrVal["val"] = "0"
			if val, ok := devInfo[iotutil.ToString(v.AttrDpid)]; ok {
				attrVal["val"] = val
			}
			attrValNew := iotutil.MapStringToInterface(attrVal)
			properties = append(properties, entitys.AlexaControlProperties{
				Namespace:                 interfaceStr,
				Name:                      instanceStr,
				Value:                     attrValNew["val"],
				TimeOfSample:              time.Now().Format(time.RFC3339),
				UncertaintyInMilliseconds: 0,
			})
		}
		requestData := entitys.EventRequet{
			Event: entitys.CommonVoiceRequest{
				Header: entitys.CommonHeader{
					Namespace:        "Alexa.Discovery",
					Name:             "AddOrUpdateReport",
					MessageID:        iotutil.UUID(),
					PayloadVersion:   "3",
					CorrelationToken: correlationToken,
				},
				Endpoint: entitys.AlexaVoicePayload{
					Scope: entitys.AlexaVoiceScope{
						Type:  "BearerToken",
						Token: alexaToken,
					},
					EndpointId: device.DeviceId,
				},
				Payload: map[string]interface{}{},
			},
			Context: &entitys.AlexaControlPropertiesContext{Properties: properties},
		}
		for i := 0; i < 3; i++ {
			code, _ := pushToAlexa(userId, alexaToken, alexaEventUrl, requestData, i)
			if code != "INTERNAL_SERVICE_EXCEPTION" {
				break
			}
			time.Sleep(time.Duration(3) * time.Second)
		}
	}
	return nil
}

func RequestDeviceListSync(res iotstruct.DeviceRedisUpdate) error {
	defer iotutil.PanicHandler(res)
	userId := res.UserId
	tokenInfo, err := getCachedAlexaTokenInfo(userId)
	if err != nil {
		return err
	}
	alexaToken := tokenInfo["access_token"]
	alexaEventUrl := tokenInfo["alexaEventUrl"]

	//发现设备
	requestData := &entitys.DirectiveRequet{
		Directive: entitys.CommonVoiceRequest{
			Header: entitys.CommonHeader{
				Namespace:      "Alexa.Discovery",
				Name:           "AddOrUpdateReport",
				MessageID:      iotutil.UUID(),
				PayloadVersion: "3",
			},
		},
	}
	err = DiscoverDevices(requestData, res.UserId, "")
	if err != nil {
		return err
	}
	r := requestData.Directive.Payload.(entitys.AlexaPayload)
	r.Scope = &entitys.AlexaVoiceScope{
		Type:  "BearerToken",
		Token: alexaToken,
	}
	rd := &entitys.EventRequet{
		Event: entitys.CommonVoiceRequest{
			Header:  requestData.Directive.Header,
			Payload: r,
		},
	}

	for i := 0; i < 3; i++ {
		code, _ := pushToAlexa(userId, alexaToken, alexaEventUrl, rd, i)
		if code != "INTERNAL_SERVICE_EXCEPTION" {
			break
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
	return nil
}

func RequestDeviceSync(res iotstruct.DeviceRedisUpdate) error {
	iotlogger.LogHelper.Helper.Infof("开始同步alexa设备数据: %s ", iotutil.ToString(res))

	defer iotutil.PanicHandler(res)
	userId := res.UserId
	tokenInfo, err := getCachedAlexaTokenInfo(userId)
	if err != nil {
		return err
	}
	alexaToken := tokenInfo["access_token"]
	alexaEventUrl := tokenInfo["alexaEventUrl"]

	endpoints := []entitys.AlexaEndpoints{}
	for _, devId := range res.DeviceIds {
		endpoints = append(endpoints, entitys.AlexaEndpoints{
			EndpointID: devId,
		})
	}

	//发现设备
	requestData := &entitys.EventRequet{
		Event: entitys.CommonVoiceRequest{
			Header: entitys.CommonHeader{
				Namespace:      "Alexa.Discovery",
				Name:           "DeleteReport",
				MessageID:      iotutil.UUID(),
				PayloadVersion: "3",
			},
		},
	}
	requestData.Event.Payload = entitys.AlexaPayload{
		Scope: &entitys.AlexaVoiceScope{
			Type:  "BearerToken",
			Token: alexaToken,
		},
		Endpoints: endpoints,
	}
	for i := 0; i < 3; i++ {
		code, err := pushToAlexa(userId, alexaToken, alexaEventUrl, requestData, i)
		if err != nil {
			iotlogger.LogHelper.Helper.Errorf("开始同步alexa设备数据: %s ", iotutil.ToString(res))
		}
		if code != "INTERNAL_SERVICE_EXCEPTION" {
			break
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
	return nil
}

// 获取设备信息
func getDeviceInfo(devId string) (map[string]string, error) {
	deviceStatus, redisErr := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+devId).Result()
	if redisErr != nil {
		return nil, errors.New("设备信息获取失败")
	}
	//if deviceStatus["onlineStatus"] != "online" {
	//	return nil, errors.New("设备不在线")
	//}
	productKey := deviceStatus["productKey"]
	if productKey == "" {
		return nil, errors.New("设备信息异常")
	}
	return deviceStatus, nil
}

func pushToAlexa(userId string, alexaToken string, alexaEventUrl string, requestData interface{}, againCount int) (string, error) {
	//北美：https://api.amazonalexa.com/v3/events
	//欧洲和印度：https://api.eu.amazonalexa.com/v3/events
	//远东和澳大利亚：https://api.fe.amazonalexa.com/v3/events
	if alexaEventUrl == "" {
		return "", errors.New("alexaEventUrl参数异常")
	}
	if alexaToken == "" {
		return "", errors.New("alexaToken参数异常")
	}
	var apiUrl = alexaEventUrl //"https://api.amazonalexa.com/v3/events"

	reqData, _ := json.Marshal(requestData)
	iotlogger.LogHelper.Helper.Debug("向Alexa推送的数据: ", string(reqData))
	// 创建 HTTP POST 请求
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(reqData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", alexaToken))

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	// 读取响应体并处理
	defer resp.Body.Close()

	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		iotlogger.LogHelper.Helper.Debug("向Alexa推送报错: err： ", err.Error())
		return "", err
	}
	iotlogger.LogHelper.Helper.Debugf("向Alexa推送设备列表更新响应（%v）:%v ", againCount, string(resbody))

	resMap := gjson.GetBytes(resbody, "payload").Map()
	if code, ok := resMap["code"]; ok {
		//if code.String() != "INTERNAL_SERVICE_EXCEPTION" {
		//	break
		//}
		return code.String(), nil
	}

	return "", nil
}
