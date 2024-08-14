package commonGlobal

import (
	"cloud_platform/iot_app_api_service/controls/common/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func GetRegionInfo(lang string, regionId string) (*entitys.SysRegionServerEntitysList, error) {
	regionIdInt, _ := iotutil.ToInt64AndErr(regionId)
	rep, err := rpc.SysRegionServerService.FindById(context.Background(), &protosService.SysRegionServerFilter{Id: regionIdInt})
	if err != nil {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	if rep.Code != 200 {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	if len(rep.Data) == 0 {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	return entitys.SysRegionServer_pb2e(rep.Data[0], lang), nil
}
