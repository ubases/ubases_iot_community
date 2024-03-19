package services

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"
	"time"
)

// GenerateNetworkToken 将生成的Token缓存到Redis中，key=token， value= { UserId: userid, Devices: [] }
func GenerateNetworkToken(cacheData iotstruct.DeviceNetworkTokenCacheModel) (token string, err error) {
	token = iotutil.Uuid()
	iotlogger.LogHelper.Info("token---" + token)
	cacheData.Devices = []string{}
	expireTime := 600 * time.Second                                                                       //十分钟的有效时间
	res := iotredis.GetClient().Set(context.Background(), token, iotutil.ToString(cacheData), expireTime) //有效期10分钟
	iotlogger.LogHelper.Info("res---" + iotutil.ToString(res))
	if res.Err() != nil {
		iotlogger.LogHelper.Errorf("SendSms,缓存smsCodeInt失败:%s", res.Err().Error())
		return "", res.Err()
	}
	return
}

// CheckNetworkToken 将生成的Token缓存到Redis中，key=token， value= { UserId: userid, Devices: [] }
func CheckNetworkToken(token, devId string) (resInt int32, msg string, err error) {
	cacheInfo, err := GetTokenCacheInfo(token)
	if err != nil {
		return 0, "", err
	}
	resInt = 0
	if cacheInfo.DevicesMap != nil {
		if res, ok := cacheInfo.DevicesMap[devId]; ok {
			resInt = iotutil.ToInt32(res.Code)
			msg = res.Msg
		}
	}
	return
}

// GetTokenCacheInfo 从token中获取用户名称（redis，token作为key）
func GetTokenCacheInfo(token string) (cacheInfo *iotstruct.DeviceNetworkTokenCacheModel, err error) {
	var cacheData iotstruct.DeviceNetworkTokenCacheModel
	valueCmd := iotredis.GetClient().Get(context.Background(), token)
	val := valueCmd.Val()
	if val == "" {
		return nil, errors.New("无任何数据")
	}
	err = iotutil.JsonToStruct(val, &cacheData)
	if err != nil {
		err = errors.New("redis cache convert error " + val)
		return nil, nil
	}
	return &cacheData, nil
}
