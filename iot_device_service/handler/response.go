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

func SetResponseNew(rsp *protosService.Response, codes int32, err error) {
	if err != nil {
		rsp.Code = codes
		rsp.Message = err.Error()
	} else {
		rsp.Code = codes
		rsp.Message = "success"
	}
}
