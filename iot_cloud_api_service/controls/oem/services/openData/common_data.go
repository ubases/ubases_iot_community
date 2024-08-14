package openData

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func GetAppMaps(ctx context.Context, tenantId string) (appKeyMap map[string]string, appIdMap map[int64]string) {
	appKeyMap = make(map[string]string)
	appIdMap = make(map[int64]string)
	//查询开发者的APP信息
	apps, err := rpc.ClientOemAppService.Lists(ctx, &protosService.OemAppListRequest{
		Query: &protosService.OemApp{
			TenantId: tenantId,
		},
	})
	if err != nil {
		return
	}
	if apps.Code != 200 {
		return
	}
	for _, app := range apps.Data {
		appIdMap[app.Id] = app.Name
		appKeyMap[app.AppKey] = app.Name
	}
	return
}
