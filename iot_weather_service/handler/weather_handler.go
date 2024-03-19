package handler

import (
	"cloud_platform/iot_weather_service/service"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type WeatherServiceHandler struct{}

func (WeatherServiceHandler) CurrentByCity(ctx context.Context, request *proto.CityRequest, response *proto.CurrentData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.CurrentByCity(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) ForecasByCity(ctx context.Context, request *proto.CityRequest, response *proto.ForecastData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.ForecasByCity(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) CurrentByIP(ctx context.Context, request *proto.IPRequest, response *proto.CurrentData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.CurrentByIP(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) ForecastByIP(ctx context.Context, request *proto.IPRequest, response *proto.ForecastData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.ForecastByIP(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) CurrentByCoordinates(ctx context.Context, request *proto.CoordinatesRequest, response *proto.CurrentData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.CurrentByCoordinates(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) ForecastByCoordinates(ctx context.Context, request *proto.CoordinatesRequest, response *proto.ForecastData) error {
	s := service.WeatherSvc{Ctx: ctx}
	resp, err := s.ForecastByCoordinates(request)
	if err == nil {
		*response = *resp
	}
	return nil
}

func (WeatherServiceHandler) Subscribe(ctx context.Context, request *proto.SubscribeRequest, response *proto.Response) error {
	s := service.WeatherSvc{Ctx: ctx}
	_, err := s.Subscribe(request)
	SetResponse(response, err)
	return nil
}

func (WeatherServiceHandler) UnSubscribe(ctx context.Context, request *proto.UnSubscribeRequest, response *proto.Response) error {
	s := service.WeatherSvc{Ctx: ctx}
	_, err := s.UnSubscribe(request)
	SetResponse(response, err)
	return nil
}
