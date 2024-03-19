package controls

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"context"
	"fmt"
	"strings"
)

func ConvertProImg(proImg string) string {
	if proImg != "" && config.Global.Oss.ImgStyle != nil && config.Global.Oss.ImgStyle.ProductImg != "" {
		if strings.Index(proImg, config.Global.Oss.ImgStyle.ProductImg) == -1 {
			proImg = proImg + config.Global.Oss.ImgStyle.ProductImg
		}
	}
	return proImg
}

// GetProductTslCached 获取产品Tsl缓存
func GetProductTslCached(tenantId string) map[string]string {
	cacheKey := fmt.Sprintf("%s_%s", tenantId, iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_PRODUCT_THINGS_MODEL)
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	if err != nil {
		iotlogger.LogHelper.Debugf("获取产品TSL缓存异常, CachedKey:%v", cacheKey)
		return make(map[string]string)
	}
	if langMap == nil {
		return make(map[string]string)
	}
	return langMap
}

// GetDeviceInfoCached 获取设备的详细信息
func GetDeviceInfoCached(deviceId string) map[string]string {
	cacheKey := iotconst.HKEY_DEV_DATA_PREFIX + deviceId
	deviceInfo, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
	if err != nil {
		iotlogger.LogHelper.Debugf("获取设备缓存异常, CachedKey:%v", cacheKey)
		return nil
	}
	if deviceInfo == nil {
		return nil
	}
	return deviceInfo
}

// GetDeviceName 获取用户设备的名称
func GetDeviceName(deviceId string) string {
	onlineStatus, _ := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId, iotconst.FIELD_DEVICE_NAME).Result()
	return onlineStatus
}
