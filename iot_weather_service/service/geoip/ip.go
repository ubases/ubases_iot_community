package geoip

import (
	"cloud_platform/iot_weather_service/config"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"net"
	"strconv"
)

type IPSvc struct {
	Ctx context.Context
}

func (s *IPSvc) GetData(req *proto.GeoIpDataRequest) (*proto.GeoIpDataResponse, error) {
	data, err := geoMgr.City(net.IP(req.Ip))
	if err != nil {
		return nil, err
	}
	var res proto.GeoIpDataResponse
	res.City = &proto.GeoIpCity{GeoNameID: uint32(data.City.GeoNameID), Names: data.City.Names}
	res.Country = &proto.GeoIpCountry{GeoNameID: uint32(data.Country.GeoNameID), IsoCode: data.Country.IsoCode,
		IsInEuropeanUnion: data.Country.IsInEuropeanUnion, Names: data.Country.Names}
	res.Location = &proto.GeoIpLocation{AccuracyRadius: uint32(data.Location.AccuracyRadius),
		Latitude: data.Location.Latitude, Longitude: data.Location.Longitude,
		MetroCode: uint32(data.Location.MetroCode), TimeZone: data.Location.TimeZone}

	return &res, nil
}

func (s *IPSvc) GetDataEx(req *proto.GeoIpDataRequest) (*proto.GeoIpResp, error) {
	data, err := geoMgr.City(net.IP(req.Ip))
	if err != nil {
		return nil, err
	}
	res := proto.GeoIpResp{
		Code:      strconv.Itoa(int(data.City.GeoNameID)),
		Name:      getValue(req.GetLang(), data.City.Names),
		Country:   getValue(req.GetLang(), data.Country.Names),
		City:      getValue(req.GetLang(), data.City.Names),
		District:  getValue(req.GetLang(), data.City.Names),
		Adcode:    int64(data.City.GeoNameID),
		Latitude:  data.Location.Latitude,
		Longitude: data.Location.Longitude,
	}
	if len(data.Subdivisions) > 0 {
		res.Province = getValue(req.GetLang(), data.Subdivisions[0].Names)
	}
	return &res, nil
}

func (s *IPSvc) GetGeoIPInfo(req *proto.IPRequest) (*proto.GeoIPInfo, error) {
	geo, err := iotutil.Geoip(req.GetIp(), config.Global.IpService.QueryUrl, config.Global.IpService.AppCode)
	if err != nil {
		return nil, err
	}
	ret := proto.GeoIPInfo{
		EnCode:   geo.EnCode,
		EnName:   geo.EnName,
		Country:  geo.Country,
		Province: geo.Province,
		City:     geo.City,
		District: geo.District,
		Adcode:   int32(geo.Adcode),
		Lat:      geo.Lat,
		Lng:      geo.Lng,
	}
	return &ret, nil
}

func getValue(lang string, mapdata map[string]string) string {
	var val string
	if len(mapdata) > 0 {
		var ok bool
		val, ok = mapdata[lang]
		if !ok {
			val, _ = mapdata["en"]
		}
	}
	return val
}
