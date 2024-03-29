// Code generated by sgen.exe,2022-06-02 11:15:12. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_app_oem/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func OemAppPushCert_pb2db(src *proto.OemAppPushCert) *model.TOemAppPushCert {
	if src == nil {
		return nil
	}
	dbObj := model.TOemAppPushCert{
		Id:      src.Id,
		AppId:   src.AppId,
		Version: src.Version,
		Apns:    src.Apns,
		Jpush:   src.Jpush,
		Fcm:     src.Fcm,
		Huawei:  src.Huawei,
		Xiaomi:  src.Xiaomi,
		Vivo:    src.Vivo,
		Oppo:    src.Oppo,
	}
	return &dbObj
}

func OemAppPushCert_db2pb(src *model.TOemAppPushCert) *proto.OemAppPushCert {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppPushCert{
		Id:      src.Id,
		AppId:   src.AppId,
		Version: src.Version,
		Apns:    src.Apns,
		Jpush:   src.Jpush,
		Fcm:     src.Fcm,
		Huawei:  src.Huawei,
		Xiaomi:  src.Xiaomi,
		Vivo:    src.Vivo,
		Oppo:    src.Oppo,
	}
	return &pbObj
}
