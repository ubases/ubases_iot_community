package extract

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

func GetAppInfo(ctx context.Context, appKey string) (*protosService.OemApp, error) {
	res, err := rpc.ClientOemAppService.Find(ctx, &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Message)
	}
	if res.Data == nil || len(res.Data) <= 0 {
		return nil, errors.New("参数错误,未找到记录")
	}
	return res.Data[0], nil
}
