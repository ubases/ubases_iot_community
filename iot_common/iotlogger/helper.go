package iotlogger

import (
	//"cloud_platform/iot_common/iotlogger/global_exception"

	"cloud_platform/iot_common/iotlogger/global_exception"
	"go-micro.dev/v4/logger"
)

type LoggerHelper struct {
	*logger.Helper
}

var LogHelper LoggerHelper

func InitLog(filename string, programName string, levelStr string) error {
	level, err := logger.GetLevel(levelStr)
	if err != nil {
		return err
	}
	l, err := NewLogger(
		WithCallerSkip(2),
		WithLogFileName(filename),
		logger.WithLevel(level),
		logger.WithFields(map[string]interface{}{"programName": programName}),
	)
	if err != nil {
		return err
	}
	//加载全局异常捕获
	global_exception.RedirectStderr(filename)
	//设置为go-micro默认日志
	logger.DefaultLogger = l
	LogHelper = LoggerHelper{logger.NewHelper(l)}
	return nil
}

func (lh *LoggerHelper) WithTag(key, value string) *LoggerHelper {
	return &LoggerHelper{lh.WithFields(map[string]interface{}{key: value})}
}
