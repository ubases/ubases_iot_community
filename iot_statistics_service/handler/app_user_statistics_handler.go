package handler

import (
	"cloud_platform/iot_statistics_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type AppUserStatisticsHandler struct{}

func (AppUserStatisticsHandler) GetAppUserStatistics(ctx context.Context, filter *proto.AppUserStatisticsFilter, response *proto.AppUserStatisticsResponse) error {
	s := service.AppUserStatisticsSvc{Ctx: ctx}
	rsp, err := s.GetAppUserStatistics(filter)
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
