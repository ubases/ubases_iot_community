package controls

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/controls/common/commonGlobal"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
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

// GetDevicePowerstate 获取设备开关状态
func GetDevicePowerstate(deviceId string) string {
	powerstate, _ := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_DEV_DATA_PREFIX+deviceId, "1").Result()
	return powerstate
}

type SysRegionServerEntitysList struct {
	Id              int64  `json:"id,string"`
	Sid             int64  `json:"sid,string"`
	HttpServer      string `json:"host"`
	Describe        string `json:"name"`
	IsDefault       int32  `json:"isDefault"`
	AreaPhoneNumber string `json:"areaPhoneNumber"`
	MqttServer      string `json:"mqttServer"`
}

func RegionIdToServerId(regionId string) (int64, error) {
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_REGION_DATA+regionId, "sid")
	res := strCmd.Val()
	serverId, _ := iotutil.ToInt64AndErr(res)
	if serverId == 0 {
		regionInfo, err := commonGlobal.GetRegionInfo("zh", regionId)
		if err != nil {
			return 0, err
		}
		iotredis.GetClient().HSet(context.Background(), iotconst.HKEY_REGION_DATA+regionId, map[string]interface{}{
			"id":              regionInfo.Id,
			"sid":             regionInfo.Sid,
			"host":            regionInfo.HttpServer,
			"name":            regionInfo.Describe,
			"isDefault":       regionInfo.IsDefault,
			"areaPhoneNumber": regionInfo.AreaPhoneNumber,
			"mqttServer":      regionInfo.MqttServer,
		})
		serverId = regionInfo.Sid
	}
	return serverId, nil
}

// 获取产品面板信息
func GetProductPanelInfo(poductIds ...int64) (map[int64]*proto.ProductPanelInfoItem, error) {
	var proMap = make(map[int64]*proto.ProductPanelInfoItem)
	if len(poductIds) == 0 {
		return proMap, nil
	}
	pros, err := rpc.ProductService.GetProductPanelInfo(context.Background(), &proto.ListsByProductIdsRequest{
		ProductIds: poductIds,
	})
	if err != nil {
		return proMap, err
	}
	for _, p := range pros.Data {
		proMap[p.ProductId] = p
	}
	return proMap, nil
}
