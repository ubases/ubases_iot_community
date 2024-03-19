package services

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

type OemAppVersionRecordService struct {
	Ctx context.Context
}

func (s OemAppVersionRecordService) SetContext(ctx context.Context) OemAppVersionRecordService {
	s.Ctx = ctx
	return s
}

// 更新app版本
func (s OemAppVersionRecordService) GetOemAppVersionRecordList(req *entitys.OemAppVersionRecordQuery) ([]*entitys.OemAppVersionRecordEntitys, error) {
	reqV := &protosService.OemAppVersionRecordListRequest{
		Query: &protosService.OemAppVersionRecord{
			AppId: iotutil.ToInt64(req.Query.AppId),
		},
	}
	resp, err := rpc.ClientOemAppVersionRecordService.Lists(s.Ctx, reqV)
	if err != nil {
		return nil, err
	}
	data := []*entitys.OemAppVersionRecordEntitys{}
	for i := range resp.Data {
		data = append(data, entitys.OemAppVersionRecord_pb2e(resp.Data[i]))
	}
	return data, nil
}
