package xiaomi

import (
	"bytes"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RequestSync(userId string, deviceId string, data map[string]interface{}) error {
	defer iotutil.PanicHandler(userId, deviceId)
	//params := make(map[string]gjson.Result)
	//for k, _ := range data {
	//	params[k] = gjson.Parse("{\"value\": \"\"}").Get("value")
	//}
	//funcData, err := common.VoiceGetFuncData(context.Background(), deviceId, params, nil, false, "alexa")
	//if err != nil {
	//	iotlogger.LogHelper.Helper.Error("xiaomi.RequestSync error: ", err)
	//	return err
	//}
	devInfo, err := common.GetDeviceInfo(deviceId)
	if err != nil {
		return err
	}
	//productKey := devInfo["productKey"]
	subscriptionId := devInfo["subscriptionId"]

	miDevs := make([]entitys.XiaomiSubscribe, 0)
	miDevs = append(miDevs, entitys.XiaomiSubscribe{
		Did:            deviceId,
		Status:         0,
		SubscriptionId: subscriptionId,
	})
	requestData := map[string]interface{}{
		"requestId": iotutil.UUID(),
		"devices":   miDevs,
	}
	return pushToXiaomi(userId, requestData)
}

func RequestDeviceListSync(res iotstruct.DeviceRedisUpdate) error {
	defer iotutil.PanicHandler(res)
	userId, err := iotutil.ToInt64AndErr(res.UserId)
	if err != nil {
		return err
	}
	if res.DeviceIds == nil || len(res.DeviceIds) == 0 {
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
		res.DeviceIds = []string{}
		for _, d := range homeDevList.Data {
			res.DeviceIds = append(res.DeviceIds, d.DeviceId)
		}
	}
	for _, devId := range res.DeviceIds {
		devInfo, err := common.GetDeviceInfo(devId)
		if err != nil {
			return err
		}
		//productKey := devInfo["productKey"]
		subscriptionId := devInfo["subscriptionId"]

		miDevs := make([]entitys.XiaomiSubscribeNotify, 0)
		miDevs = append(miDevs, entitys.XiaomiSubscribeNotify{
			Did:            devId,
			SubscriptionId: subscriptionId,
		})
		requestData := map[string]interface{}{
			"requestId": iotutil.UUID(),
			"topic":     "device-status-changed",
			"devices":   miDevs,
		}
		pushToXiaomi(res.UserId, requestData)
	}
	return nil
}

func pushToXiaomi(userId string, requestData map[string]interface{}) error {
	var apiUrl = "https://api.home.mi.com/api/xiot/notify"

	jsonStr, _ := json.Marshal(requestData)

	// 创建 HTTP POST 请求
	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	// 读取响应体并处理
	defer resp.Body.Close()

	resbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Helper.Debug("向小米推送设备列表更新响应: ", string(resbody))
	return nil
}
