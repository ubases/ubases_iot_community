package commonGlobal

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func GetRegionInfo(lang string, regionId string) (*protosService.SysRegionServer, error) {
	regionIdInt, _ := iotutil.ToInt64AndErr(regionId)
	rep, err := rpc.SysRegionServerService.FindById(context.Background(), &protosService.SysRegionServerFilter{Id: regionIdInt})
	if err != nil {
		return &protosService.SysRegionServer{}, nil
	}
	if rep.Code != 200 {
		return &protosService.SysRegionServer{}, nil
	}
	if len(rep.Data) == 0 {
		return &protosService.SysRegionServer{}, nil
	}
	return rep.Data[0], nil
}

func RegionIdToServerId(regionId string) (int64, error) {
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_REGION_DATA+regionId, "sid")
	res := strCmd.Val()
	serverId, _ := iotutil.ToInt64AndErr(res)
	if serverId == 0 {
		regionInfo, err := GetRegionInfo("zh", regionId)
		if err != nil {
			return 0, err
		}
		iotredis.GetClient().HSet(context.Background(), iotconst.HKEY_REGION_DATA+regionId, map[string]interface{}{
			"id":              regionInfo.Id,
			"sid":             regionInfo.Sid,
			"host":            regionInfo.HttpServer,
			"name":            regionInfo.Describe,
			"areaPhoneNumber": regionInfo.AreaPhoneNumber,
			"mqttServer":      regionInfo.MqttServer,
		})
		serverId, _ = iotutil.ToInt64AndErr(regionInfo.Sid)
	}
	return serverId, nil
}
