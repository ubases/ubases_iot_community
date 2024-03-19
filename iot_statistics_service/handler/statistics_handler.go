package handler

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_statistics_service/service"
	"context"
)

type StatisticsHandler struct{}

func (h StatisticsHandler) GetDeveloperList(ctx context.Context, request *proto.DeveloperStatListRequest, response *proto.DeveloperStatListResponse) error {
	s := service.StatisticsSvc{Ctx: ctx}
	rsp, err := s.GetDeveloperList(request)
	if err != nil {
		return err
	}
	*response = *rsp
	return nil
}

func (h StatisticsHandler) GetDeveloperDetail(ctx context.Context, request *proto.DeveloperDetailFilter, response *proto.DeveloperDetailResponse) error {
	s := service.StatisticsSvc{Ctx: ctx}
	rsp, err := s.GetDeveloperDetail(request)
	if err != nil {
		return err
	}
	*response = *rsp
	return nil
}

func (h StatisticsHandler) GetDeveloperStatistics(ctx context.Context, null *proto.NULL, response *proto.DeveloperStatisticsResponse) error {
	s := service.StatisticsSvc{Ctx: ctx}
	rsp, err := s.GetDeveloperStatistics(null)
	if err != nil {
		return err
	}
	*response = *rsp
	return nil
}

func (h StatisticsHandler) GetAppDataDetail(ctx context.Context, request *proto.AppDataDetailFilter, response *proto.AppDataDetailResponse) error {
	s := service.StatisticsSvc{Ctx: ctx}
	rsp, err := s.GetAppDataDetail(request)
	if err != nil {
		return err
	}
	*response = *rsp
	return nil
}

func (h StatisticsHandler) GetDeviceTotalStatistics(ctx context.Context, null *proto.NULL, response *proto.DeviceStatisticsResponse) error {
	s := service.StatisticsSvc{Ctx: ctx}
	rsp, err := s.GetDeviceTotalStatistics(null)
	if err != nil {
		return err
	}
	*response = *rsp
	return nil
}
