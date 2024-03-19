package geoip

import (
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"net"

	"cloud_platform/iot_weather_service/config"

	"github.com/oschwald/geoip2-golang"
)

var geoMgr GeoMgr

func InitGeoMgr() error {
	geoServer, err := NewServer(config.Global.Geo.DbPath, config.Global.Geo.LicenseKey)
	if err != nil {
		return err
	}
	geoMgr.server = geoServer
	return nil
}

type GeoMgr struct {
	server *GeoServer
}

func (o *GeoMgr) City(ipAddress net.IP) (*geoip2.City, error) {
	ip := string(ipAddress)
	if ip == "" {
		return nil, errors.New(fmt.Sprintf("缺IP地址"))
	}
	g := get{ip, make(chan *geoip2.City)}
	o.server.cacheGet <- g
	city := <-g.resp
	return city, nil
}

func Test() {
	//city, err := geoMgr.City(net.IP("110.242.68.66"))
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(city)
	//}
	//
	city, err := geoMgr.City(net.IP("20.205.243.166"))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(city)
	}

	svc := IPSvc{Ctx: context.Background()}
	resp, err := svc.GetDataEx(&proto.GeoIpDataRequest{
		Ip:   "199.59.150.39",
		Lang: "zh-CN",
	})
	if err == nil {
		fmt.Println(resp)
	}
}
