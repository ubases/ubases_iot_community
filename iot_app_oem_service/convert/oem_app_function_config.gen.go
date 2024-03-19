// Code generated by sgen.exe,2022-11-04 09:15:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_app_oem/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func OemAppFunctionConfig_pb2db(src *proto.OemAppFunctionConfig) *model.TOemAppFunctionConfig {
	if src == nil {
		return nil
	}
	dbObj := model.TOemAppFunctionConfig{
		Id:            src.Id,
		AppId:         src.AppId,
		Version:       src.Version,
		AboutUs:       src.AboutUs,
		Eula:          src.Eula,
		PrivacyPolicy: src.PrivacyPolicy,
		Weather:       src.Weather,
		Voices:        src.Voices,
		Thirds:        src.Thirds,
		AutoUpgrade:   src.AutoUpgrade,
		Geo:           src.Geo,
	}
	return &dbObj
}

func OemAppFunctionConfig_db2pb(src *model.TOemAppFunctionConfig) *proto.OemAppFunctionConfig {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppFunctionConfig{
		Id:            src.Id,
		AppId:         src.AppId,
		Version:       src.Version,
		AboutUs:       src.AboutUs,
		Eula:          src.Eula,
		PrivacyPolicy: src.PrivacyPolicy,
		Weather:       src.Weather,
		Voices:        src.Voices,
		Thirds:        src.Thirds,
		AutoUpgrade:   src.AutoUpgrade,
		Geo:           src.Geo,
	}
	return &pbObj
}
