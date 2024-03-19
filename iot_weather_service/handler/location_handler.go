package handler

import (
	"cloud_platform/iot_weather_service/service/geoip"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type IPServiceHandler struct{}

func (IPServiceHandler) GetData(ctx context.Context, request *proto.GeoIpDataRequest, response *proto.GeoIpDataResponse) error {
	s := geoip.IPSvc{Ctx: ctx}
	resp, err := s.GetData(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (IPServiceHandler) GetDataEx(ctx context.Context, request *proto.GeoIpDataRequest, response *proto.GeoIpResp) error {
	s := geoip.IPSvc{Ctx: ctx}
	resp, err := s.GetDataEx(request)
	if err == nil {
		*response = *resp
	}
	return err
}

func (h IPServiceHandler) GetGeoIPInfo(ctx context.Context, request *proto.IPRequest, response *proto.GeoIPInfo) error {
	s := geoip.IPSvc{Ctx: ctx}
	resp, err := s.GetGeoIPInfo(request)
	if err == nil {
		*response = *resp
	}
	return err
}
