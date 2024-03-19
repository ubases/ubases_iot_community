package service

import (
	"cloud_platform/iot_proto/protos/protosService"
)

const (
	ERROR   = 400
	SUCCESS = 200
)

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

func (err *Response) Error() string {
	return err.Message
}

//
//func FailWithMessage(message string) *Response {
//	return &Response{Code: ERROR, Message:message, Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
//}
//
//func Success() *Response {
//	return &Response{Code: SUCCESS, Message:"请求成功", Data: map[string]interface{}{}, Meta:map[string]interface{}{}}
//}
//
//func FailAny(data protobuf.Message) *proto.Response {
//	any, _ := anypb.New(data)
//	meta := &proto.Meta{Total: 1} // 需要进行初始化,初始化为0显示不出来？？？
//
//	rsp := &proto.Response{
//		Code: ERROR,
//		Message: "请求出错",
//		Data: any,
//		Meta: meta,
//	}
//	return rsp
//}
//
//func FailAnyWithError(data protobuf.Message,err error) *proto.Response {
//	any, _ := anypb.New(data)
//
//	meta := &proto.Meta{Total: 1} // 需要进行初始化,初始化为0显示不出来？？？
//
//	rsp := &proto.Response{
//		Code: ERROR,
//		Message: err.Error(),
//		Data: any,
//		Meta: meta,
//	}
//	return rsp
//}
//
//
//func FailAnyWithErrors(data protobuf.Message,err error) *proto.Responses {
//	any, _ := anypb.New(data)
//
//	meta := &proto.Meta{Total: 1}
//
//	rsp := &proto.Responses{
//		Code: ERROR,
//		Message: err.Error(),
//		Data: []*anypb.Any{any},
//		Meta: meta,
//	}
//	return rsp
//}
//
//
//func SuccessAny(data protobuf.Message) *proto.Response {
//	any, _ := anypb.New(data)
//
//	meta := &proto.Meta{Total: 1}
//
//	rsp := &proto.Response{
//		Code: SUCCESS,
//		Message: "请求成功",
//		Data: any,
//		Meta: meta,
//	}
//	return rsp
//}
//
//func SuccesssAny(any []*anypb.Any, total int64) *proto.Responses {
//	meta := &proto.Meta{Total: total}
//
//	rsps := &proto.Responses{
//		Code: SUCCESS,
//		Message: "请求成功",
//		Data: any,
//		Meta: meta,
//	}
//	return rsps
//}

func SetResponse(rsp *protosService.Response, err error) {
	if err != nil {
		rsp.Code = ERROR
		rsp.Message = err.Error()
	} else {
		rsp.Code = SUCCESS
		rsp.Message = "success"
	}
}
