package ioterrs

import (
	"bytes"
	"cloud_platform/iot_common/iotlogger"
	"context"
	"runtime"

	"go-micro.dev/v4/server"
)

// 异常拦截器
func BatPanicHandler() server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			defer func() {
				if err := recover(); err != nil {
					iotlogger.LogHelper.Errorf("error:%v", err)
					errorMsg := PanicTrace(1)
					iotlogger.LogHelper.Error(errorMsg)
				}
			}()
			return h(ctx, req, rsp)
		}
	}
}
func PanicTrace(kb int) string {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return string(stack)
}
