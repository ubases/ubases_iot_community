// Code generated by sgen.exe,2022-04-27 10:55:25. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

package convert

import (
	"cloud_platform/iot_model/db_open_system/model"
	proto "cloud_platform/iot_proto/protos/protosService"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func OpenOperLog_pb2db(src *proto.OpenOperLog) *model.TOpenOperLog {
	if src == nil {
		return nil
	}
	dbObj := model.TOpenOperLog{
		OperId:        src.OperId,
		Title:         src.Title,
		BusinessType:  src.BusinessType,
		Method:        src.Method,
		RequestMethod: src.RequestMethod,
		OperatorType:  src.OperatorType,
		OperName:      src.OperName,
		DeptName:      src.DeptName,
		OperUrl:       src.OperUrl,
		OperIp:        src.OperIp,
		OperLocation:  src.OperLocation,
		OperParam:     src.OperParam,
		JsonResult:    src.JsonResult,
		Status:        src.Status,
		ErrorMsg:      src.ErrorMsg,
		OperTime:      src.OperTime.AsTime(),
	}
	return &dbObj
}

func OpenOperLog_db2pb(src *model.TOpenOperLog) *proto.OpenOperLog {
	if src == nil {
		return nil
	}
	pbObj := proto.OpenOperLog{
		OperId:        src.OperId,
		Title:         src.Title,
		BusinessType:  src.BusinessType,
		Method:        src.Method,
		RequestMethod: src.RequestMethod,
		OperatorType:  src.OperatorType,
		OperName:      src.OperName,
		DeptName:      src.DeptName,
		OperUrl:       src.OperUrl,
		OperIp:        src.OperIp,
		OperLocation:  src.OperLocation,
		OperParam:     src.OperParam,
		JsonResult:    src.JsonResult,
		Status:        src.Status,
		ErrorMsg:      src.ErrorMsg,
		OperTime:      timestamppb.New(src.OperTime),
	}
	return &pbObj
}
