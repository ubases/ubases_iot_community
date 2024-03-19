package entitys

import proto "cloud_platform/iot_proto/protos/protosService"

type SysRegionServerEntitysList struct {
	Id              int64  `json:"id,string"`
	HttpServer      string `json:"host"`
	Describe        string `json:"name"`
	IsDefault       int32  `json:"isDefault"`
	AreaPhoneNumber string `json:"areaPhoneNumber"`
}

// pb对象转实体
func SysRegionServer_pb2e(src *proto.SysRegionServer, lang string) *SysRegionServerEntitysList {
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
		HttpServer:      src.HttpServer,
		Describe:        describe,
		AreaPhoneNumber: src.AreaPhoneNumber,
	}
	return &entitysObj
}
