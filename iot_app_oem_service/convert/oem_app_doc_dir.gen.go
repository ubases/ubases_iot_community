// Code generated by sgen.exe,2022-07-14 15:09:42. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_app_oem/model"
	proto "cloud_platform/iot_proto/protos/protosService"
)

func OemAppDocDir_pb2db(src *proto.OemAppDocDir) *model.TOemAppDocDir {
	if src == nil {
		return nil
	}
	dbObj := model.TOemAppDocDir{
		Id:       src.Id,
		DocId:    src.DocId,
		ParentId: src.ParentId,
		DirName:  src.DirName,
		DirImg:   src.DirImg,
		Sort:     src.Sort,
	}
	return &dbObj
}

func OemAppDocDir_db2pb(src *model.TOemAppDocDir) *proto.OemAppDocDir {
	if src == nil {
		return nil
	}
	pbObj := proto.OemAppDocDir{
		Id:       src.Id,
		DocId:    src.DocId,
		ParentId: src.ParentId,
		DirName:  src.DirName,
		DirImg:   src.DirImg,
		Sort:     src.Sort,
	}
	return &pbObj
}
