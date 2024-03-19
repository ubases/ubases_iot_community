package controls

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

func Geoip(ip string) (iotutil.GeoipInfo, error) {
	geo, err := rpc.IPService.GetGeoIPInfo(context.Background(), &protosService.IPRequest{Ip: ip})
	if err != nil {
		return iotutil.GeoipInfo{}, err
	}
	ret := iotutil.GeoipInfo{
		EnCode:   geo.EnCode,
		EnName:   geo.EnName,
		Country:  geo.Country,
		Province: geo.Province,
		City:     geo.City,
		District: geo.District,
		Adcode:   int(geo.Adcode),
		Lat:      geo.Lat,
		Lng:      geo.Lng,
	}
	return ret, nil
}
