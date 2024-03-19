package task

import (
	"cloud_platform/iot_common/iotlogger"

	"github.com/xxl-job/xxl-job-executor-go"
)

// xxl.Logger接口实现
type XxlLogger struct{}

func (l *XxlLogger) Info(format string, a ...interface{}) {
	iotlogger.LogHelper.Infof(format, a...)
}

func (l *XxlLogger) Error(format string, a ...interface{}) {
	iotlogger.LogHelper.Infof(format, a...)
}

func XxlLogHandler(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "这个是自定义日志handler",
		IsEnd:       true,
	}}
}
