package handler

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

func SetResponse(rsp *protosService.Response, err error) {
	if err != nil {
		rsp.Code = ERROR
		rsp.Message = err.Error()
	} else {
		rsp.Code = SUCCESS
		rsp.Message = "success"
	}
}

func SetResponseCustomCode(rsp *protosService.Response, err error, code int32) {
	if err != nil {
		rsp.Code = code
		rsp.Message = err.Error()
	} else {
		rsp.Code = SUCCESS
		rsp.Message = "success"
	}
}
