package handler

import (
	"cloud_platform/iot_statistics_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type DeviceActiveHandler struct{}

func (DeviceActiveHandler) Lists(ctx context.Context, filter *proto.DeviceActiveListFilter, response *proto.DeviceActiveResponse) error {
	s := service.DeviceActiveSvc{Ctx: ctx}
	rsp, err := s.Lists(filter)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
		response.Data = rsp.Data
	}
	return nil
}
