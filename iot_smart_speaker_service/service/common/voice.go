package common

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/cached"
	"cloud_platform/iot_smart_speaker_service/entitys"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Voice struct {
}

func NewVoiceApi() *Voice {
	s := &Voice{}
	return s
}

func (v *Voice) GetList() ([]map[string]interface{}, error) {
	res, err := rpcclient.ClientOemAppFunctionConfigService.Lists(context.Background(), &protosService.OemAppFunctionConfigListRequest{})
	if err != nil {
		return nil, err
	}
	if res.Code != 200 && res.Message != "record not found" {
		return nil, errors.New(res.Message)
	}

	var list []map[string]interface{}
	for i := range res.Data {
		var items []map[string]interface{}
		if res.Data[i].Voices != "" {
			if err := json.Unmarshal([]byte(res.Data[i].Voices), &items); err != nil {
				iotlogger.LogHelper.Helper.Error("json unmarshal voices error: ", err)
				continue
			}
			list = append(list, items...)
		}
	}
	return list, nil
}

func GetDetail(req protosService.OpmVoiceProductDetailReq) (*protosService.OpmVoiceProductDetailRes, error) {
	if req.SkillId == "" {
		return nil, errors.New("SkillI不能为空")
	}
	if req.VoiceNo == "" {
		return nil, errors.New("VoiceNo不能为空")
	}
	var res protosService.OpmVoiceProductDetailRes
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.VOICE_PRODUCT_SHILL_CACHED, req.VoiceNo, req.SkillId), res)
	if err == nil {
		return &res, nil
	}

	resObj, err := rpcclient.ClienOpmVoiceProductService.GetDetail(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resObj.Code != 200 {
		return nil, errors.New(resObj.Message)
	}
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.VOICE_PRODUCT_SHILL_CACHED, req.VoiceNo, req.SkillId), &resObj, 12000*time.Second)
	if err != nil {
		iotlogger.LogHelper.Error("Voice GetDetail error:", err)
		return resObj, nil
	}
	return resObj, nil
}

// 缓存语音产品信息
func GetDetailList(req protosService.OpmVoiceProductDetailReq) ([]*protosService.OpmVoiceProductDetailRes, error) {
	if req.SkillId == "" {
		return nil, errors.New("SkillI不能为空")
	}
	if req.VoiceNo == "" {
		return nil, errors.New("VoiceNo不能为空")
	}
	var res []*protosService.OpmVoiceProductDetailRes
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.VOICE_PRODUCT_LIST_SHILL_CACHED, req.VoiceNo, req.SkillId), res)
	if err == nil {
		return res, nil
	}

	resObj, err := rpcclient.ClienOpmVoiceProductService.GetDetailList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resObj.Code != 200 {
		return nil, errors.New(resObj.Message)
	}
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.VOICE_PRODUCT_LIST_SHILL_CACHED, req.VoiceNo, req.SkillId), &resObj, 12000*time.Second)
	if err != nil {
		iotlogger.LogHelper.Error("Voice GetDetail error:", err)
		return resObj.Details, nil
	}
	return resObj.Details, nil
}

// 缓存语音产品信息
func GetVoiceProductDetails(ctx context.Context, productKey, voiceNo string) (*entitys.VoiceProductCached, error) {
	var res entitys.VoiceProductCached
	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, voiceNo, productKey), &res)
	if err == nil {
		return &res, nil
	}

	res = entitys.VoiceProductCached{}
	// 通过产品Key获取产品语控配置信息
	opmVoice, err := rpcclient.ClienOpmVoiceProductService.Find(ctx, &protosService.OpmVoiceProductFilter{
		ProductKey: productKey,
		VoiceNo:    voiceNo,
	})
	if err != nil {
		return nil, err
	}
	if len(opmVoice.Data) == 0 {
		return nil, fmt.Errorf("产品：%s 的语控配置信息为空", productKey)
	}
	res.VoiceProductInfo = opmVoice.Data[0]
	//{"voiceBrand":"13123","voiceModel":"23123","voiceSkill":"222"}
	voiceOther, err := iotutil.JsonToMapErr(res.VoiceProductInfo.VoiceOther)
	if err != nil {
		return nil, errors.New("语控配置信息异常")
	}
	res.VoiceBrand = iotutil.ToString(voiceOther["voiceBrand"])
	res.VoiceModel = iotutil.ToString(voiceOther["voiceModel"])
	res.VoiceSkill = iotutil.ToString(voiceOther["voiceSkill"])
	// 设备属性状态需要从哪里获取？redis缓存dev_data_设备id，通过dpid从redis hash里面取? 非必填，暂时不填，防止转换出错，天猫获取设备失败
	voiceMapList, err := rpcclient.ClienOpmVoiceProductMapService.Lists(ctx, &protosService.OpmVoiceProductMapListRequest{
		Query: &protosService.OpmVoiceProductMap{
			VoiceProductId: res.VoiceProductInfo.Id,
			VoiceNo:        voiceNo,
		},
	})
	if err != nil {
		return nil, err
	}
	res.FunctionMap = voiceMapList.Data
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, voiceNo, productKey), res, 600*time.Second)
	if err != nil {
		iotlogger.LogHelper.Error("Voice GetVoiceProductDetails cached error:", err)
	}

	return &res, nil
}

// GetDeviceInfo 获取设备信息
func GetDeviceInfo(devId string) (map[string]string, error) {
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

func GetDeviceOnline(devId string) bool {
	deviceStatus, redisErr := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+devId, "onlineStatus").Result()
	if redisErr != nil {
		return false
	}
	if deviceStatus == "online" {
		return true
	}
	return false
}

func getDeviceTriad(devId string) (triadInfo *protosService.IotDeviceTriad, err error) {
	devTriad, err := rpcclient.ClientIotDeviceTriadService.Find(context.Background(), &protosService.IotDeviceTriadFilter{
		Did: devId,
	})
	if err != nil {
		return nil, err
	}
	if len(devTriad.Data) == 0 {
		return nil, errors.New("未获取到三元组数据")
	}
	return devTriad.Data[0], nil
}

//缓存语音产品信息
//func getVoiceProduct(ctx context.Context, productKey string, voiceNo string) ([]*protosService.OpmVoiceProduct, error) {
//	var res []*protosService.OpmVoiceProduct
//	err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, voiceNo, productKey), res)
//	if err == nil {
//		return res, nil
//	}
//	// 通过产品Key获取产品语控配置信息
//	opmVoice, err := rpcclient.ClienOpmVoiceProductService.Find(ctx, &protosService.OpmVoiceProductFilter{
//		ProductKey: productKey,
//	})
//	if err != nil {
//		return nil, err
//	}
//	if len(opmVoice.Data) == 0 {
//		return nil, fmt.Errorf("产品：%s 的语控配置信息为空", productKey)
//	}
//	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.VOICE_PRODUCT_DATA_CACHED, voiceNo, productKey), &opmVoice.Data, 12000*time.Second)
//	if err != nil {
//		iotlogger.LogHelper.Error("Voice getVoiceProduct cached error:", err)
//	}
//	return opmVoice.Data, nil
//}
