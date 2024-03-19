package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"cloud_platform/iot_weather_service/config"
	owm "cloud_platform/iot_weather_service/service/openweathermap"
	proto "cloud_platform/iot_proto/protos/protosService"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type WeatherSvc struct {
	Ctx context.Context
}

func (s *WeatherSvc) CurrentByCity(request *proto.CityRequest) (*proto.CurrentData, error) {
	data, err := GetSubscriber().CurrentByCity(request.CityName, request.Province)
	if err != nil {
		return nil, err
	}
	ret := &proto.CurrentData{
		CityCode:  request.CityCode,
		CityName:  request.CityName,
		UpdatedAt: timestamppb.Now(),
		Data:      data,
	}
	return ret, nil
}

func (s *WeatherSvc) ForecasByCity(request *proto.CityRequest) (*proto.ForecastData, error) {
	c5, err := owm.NewForecast("5", "F", "EN", config.Global.Weather.ApiKey)
	if err != nil {
		return nil, err
	}
	if err = c5.DailyByName(request.CityName, 5); err != nil {
		return nil, err
	}
	c5data := c5.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	data, _ := Forecast5WeatherData_topb(c5data)
	ret := &proto.ForecastData{
		CityCode:  request.CityCode,
		CityName:  request.CityName,
		UpdatedAt: timestamppb.Now(),
		Data:      data,
	}
	return ret, nil
}

func (s *WeatherSvc) CurrentByIP(request *proto.IPRequest) (*proto.CurrentData, error) {
	//cwd, err := owm.NewCurrent("F", "EN", config.Global.Weather.ApiKey)
	//if err != nil {
	//	return nil, err
	//}
	//if err = cwd.CurrentByID(); err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (s *WeatherSvc) ForecastByIP(request *proto.IPRequest) (*proto.ForecastData, error) {
	return nil, nil
}

func (s *WeatherSvc) CurrentByCoordinates(request *proto.CoordinatesRequest) (*proto.CurrentData, error) { //nolint:lll
	cwd, err := owm.NewCurrent("F", "EN", config.Global.Weather.ApiKey)
	if err != nil {
		return nil, err
	}
	if err = cwd.CurrentByCoordinates(&owm.Coordinates{
		Longitude: request.Longitude,
		Latitude:  request.Latitude}); err != nil {
		return nil, err
	}
	data, _ := WeatherData_topb(cwd)
	ret := &proto.CurrentData{
		CityCode:  strconv.Itoa(cwd.ID),
		CityName:  cwd.Name,
		UpdatedAt: timestamppb.Now(),
		Data:      data,
	}
	return ret, nil
}

func (s *WeatherSvc) ForecastByCoordinates(request *proto.CoordinatesRequest) (*proto.ForecastData, error) { //nolint:lll
	c5, err := owm.NewForecast("5", "F", "EN", config.Global.Weather.ApiKey)
	if err != nil {
		return nil, err
	}
	if err = c5.DailyByCoordinates(&owm.Coordinates{
		Longitude: request.Longitude,
		Latitude:  request.Latitude}, 5); err != nil {
		return nil, err
	}
	c5data := c5.ForecastWeatherJson.(*owm.Forecast5WeatherData)
	data, _ := Forecast5WeatherData_topb(c5data)
	ret := &proto.ForecastData{
		CityCode:  strconv.Itoa(c5data.City.ID),
		CityName:  c5data.City.Name,
		UpdatedAt: timestamppb.Now(),
		Data:      data,
	}
	return ret, nil
}

func (s *WeatherSvc) Subscribe(request *proto.SubscribeRequest) (*proto.Response, error) {
	GetSubscriber().AddSubscribe(request.CityName, request.Province)
	return nil, nil
}

func (s *WeatherSvc) UnSubscribe(request *proto.UnSubscribeRequest) (*proto.Response, error) {
	GetSubscriber().UnSubscribe(request.CityName)
	return nil, nil
}

func WeatherData_topb(src *owm.CurrentWeatherData) (*proto.WeatherData, error) {
	if src == nil {
		return nil, errors.New("src is nil")
	}
	dst := proto.WeatherData{
		CityCode: strconv.Itoa(src.ID),
		CityName: src.Name,
		Date:     time.Now().Format("2006-01-02"),
		Time:     time.Now().Format("15:04"),
		Sunrise:  timestamppb.New(time.Unix(int64(src.Sys.Sunrise), 0)),
		Sunset:   timestamppb.New(time.Unix(int64(src.Sys.Sunset), 0)),
		//Weather:    src.Weather[0].Main,
		Temperature:     src.Main.Temp,
		TemperatureHigh: src.Main.TempMax,
		TemperatureLow:  src.Main.TempMin,
		Humidity:        int32(src.Main.Humidity),
		//Pm25:            ,//天气缓存
		//Pm10:            "",
		WindSpeed: src.Wind.Speed,
		WindGrade: src.Wind.Deg,
		//Aqi:        "",
		//Quality:    "",
		//WindDir:    ,
		//UvIndex:    ,
		Pressure:   src.Main.Pressure,
		Visibility: int32(src.Visibility),
		Source:     "openweathermap",
		UpdatedAt:  timestamppb.Now(),
	}
	if len(src.Weather) > 0 {
		dst.Weather = src.Weather[0].Main
		dst.Icon = src.Weather[0].Icon
		dst.WeatherDesc = src.Weather[0].Description
	}
	return &dst, nil
}

func Forecast5WeatherData_topb(src *owm.Forecast5WeatherData) ([]*proto.WeatherData, error) {
	if src == nil {
		return nil, errors.New("src is nil")
	}

	var list []*proto.WeatherData = nil
	for _, v := range src.List {
		dst := &proto.WeatherData{
			CityCode: strconv.Itoa(src.City.ID),
			CityName: src.City.Name,
			Date:     v.DtTxt.Format("2006-01-02"),
			Time:     v.DtTxt.Format("15:04"),
			//Sunrise:         ,
			//Sunset:         ,
			//Weather:         ,
			Temperature:     v.Main.Temp,
			TemperatureHigh: v.Main.TempMax,
			TemperatureLow:  v.Main.TempMin,
			Humidity:        int32(v.Main.Humidity),
			//Pm25:            ,
			//Pm10:            0,
			WindSpeed: v.Wind.Speed,
			WindGrade: v.Wind.Deg,
			//Aqi:             "",
			//Quality:         "",
			//WindDir:         "",
			//UvIndex:         0,
			Pressure: v.Main.Pressure,
			//Visibility:      ,
			Source:    "openweathermap",
			UpdatedAt: timestamppb.Now(),
		}
		list = append(list, dst)
	}
	return list, nil
}
