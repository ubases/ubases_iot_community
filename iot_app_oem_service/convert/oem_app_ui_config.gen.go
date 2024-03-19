// Code generated by sgen.exe,2022-06-02 11:15:12. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_app_oem/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func OemAppUiConfig_pb2db(src *proto.OemAppUiConfig) *model.TOemAppUiConfig {
	if src == nil {
		return nil
	}
	dbObj := model.TOemAppUiConfig{
		Id:                  src.Id,
		AppId:               src.AppId,
		Version:             src.Version,
		IconUrl:             src.IconUrl,
		IosLaunchScreen:     src.IosLaunchScreen,
		AndroidLaunchScreen: src.AndroidLaunchScreen,
		ThemeColors:         src.ThemeColors,
		BottomMenu:          src.BottomMenu,
		Personalize:         src.Personalize,
		Room:                src.Room,
		RoomIcons:           src.RoomIcons,
	}
	return &dbObj
}

func OemAppUiConfig_db2pb(src *model.TOemAppUiConfig) *proto.OemAppUiConfig {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppUiConfig{
		Id:                  src.Id,
		AppId:               src.AppId,
		Version:             src.Version,
		IconUrl:             src.IconUrl,
		IosLaunchScreen:     src.IosLaunchScreen,
		AndroidLaunchScreen: src.AndroidLaunchScreen,
		ThemeColors:         src.ThemeColors,
		BottomMenu:          src.BottomMenu,
		Personalize:         src.Personalize,
		Room:                src.Room,
		RoomIcons:           src.RoomIcons,
	}
	return &pbObj
}
