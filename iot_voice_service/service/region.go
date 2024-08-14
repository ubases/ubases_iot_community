package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"context"
	"fmt"
)

type RegionApi struct {
}

func NewRegionApi() *RegionApi {
	s := &RegionApi{}
	return s
}

type SysRegionServerEntitysList struct {
	Id              int64  `json:"id,string"`
	Describe        string `json:"name"`
	IsDefault       int32  `json:"isDefault"`
	AreaPhoneNumber string `json:"areaPhoneNumber"`
}

func SysRegionServer_pb2e(src *protosService.SysRegionServer, lang string) *SysRegionServerEntitysList {
	if src == nil {
		return nil
	}
	var describe string
	if lang == "zh" {
		describe = src.Describe
	} else if lang == "en" {
		describe = src.EnDescribe
	}
	entitysObj := SysRegionServerEntitysList{
		Id:              src.Id,
		Describe:        describe,
		AreaPhoneNumber: src.AreaPhoneNumber,
	}
	return &entitysObj
}

func RegionList(lang, ip string) ([]*SysRegionServerEntitysList, error) {
	//geo, err := iotutil.Geoip(ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode)
	geo, err := Geoip(ip)
	country := geo.Country
	result := []*SysRegionServerEntitysList{}
	//增加状态判断查询（用于启用关闭区域操作）
	rep, err := rpcclient.SysRegionServerService.Lists(context.Background(), &protosService.SysRegionServerListRequest{
		Query: &protosService.SysRegionServer{Enabled: 1},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, err
	}
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_CONFIG_REGION_SERVER).Result()
	if err != nil {
		langMap = make(map[string]string)
		//异常不处理
		iotlogger.LogHelper.Errorf("字典缓存读取异常, %s", err.Error())
	}
	haveDefaultData := false
	for _, v := range rep.Data {
		item := SysRegionServer_pb2e(v, lang)
		fKey := fmt.Sprintf("%s_%d_name", lang, item.Id)
		if country != "" && v.Describe == country {
			item.IsDefault = 1
			haveDefaultData = true
		}
		item.Describe = iotutil.MapGetStringVal(langMap[fKey], item.Describe)
		result = append(result, item)
	}
	if !haveDefaultData {
		result[0].IsDefault = 1
	}
	return result, nil
}

func Geoip(ip string) (iotutil.GeoipInfo, error) {
	geo, err := rpcclient.IPService.GetGeoIPInfo(context.Background(), &protosService.IPRequest{Ip: ip})
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
