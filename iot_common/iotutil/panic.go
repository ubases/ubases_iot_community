package iotutil

import (
	"cloud_platform/iot_common/iotlogger"
	"encoding/json"
	"fmt"
	"runtime/debug"
)

func PanicHandler(params ...interface{}) {
	if err := recover(); err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf(string(debug.Stack())))
		//iotlogger.LogHelper.Error(err.(error).Error())
		iotlogger.LogHelper.Error(fmt.Sprintf("panic...:%v\r\n", err))

		if params != nil {
			errMsg := make([]string, 0)
			for _, v := range params {
				switch v.(type) {
				case map[string]interface{}, []map[string]interface{}, interface{}, []interface{}:
					strJson, _ := json.Marshal(v)
					errMsg = append(errMsg, string(strJson))
				default:
					errMsg = append(errMsg, ToString(v))
				}
			}
			iotlogger.LogHelper.Error(fmt.Sprintf("error params: %s", errMsg))
		}
	}
}
