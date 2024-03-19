package ioterrs

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgincache/persist"
)

const (
	ErrRecordNotFound = "record not found"
)

const (
	empty       string = ""   // lang字段为空时，默认返回中文信息
	zh          string = "zh" // 中文
	en          string = "en" // 英文
	codeKeyTemp string = "code.to.message.%s"
)

type ResponseModel struct {
	Code    int32       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponsePageModel struct {
	Code    int32            `json:"code"`
	Message string           `json:"msg"`
	Data    ResponsePageData `json:"data,omitempty"`
}

type ResponsePageData struct {
	List    interface{} `json:"list,omitempty"`
	Total   int64       `json:"total"`
	Current int         `json:"page"`
}

var (
	successMsg = map[string]string{
		"zh": "成功",
		"en": "success",
	}
)

func GetCodeMsgKey(lang string) string {
	if lang == "" {
		lang = "zh"
	}
	return fmt.Sprintf(codeKeyTemp, lang)
}

func CodeToMessage(store *persist.RedisStoreEx, code int32, lang string) string {
	key := GetCodeMsgKey(lang)
	if code == 0 {
		return successMsg[lang]
	}
	switch lang {
	case empty:
		return store.HGetCodeMsg(key, strconv.Itoa(int(code)))
	case zh:
		return store.HGetCodeMsg(key, strconv.Itoa(int(code)))
	case en:
		return store.HGetCodeMsg(key, strconv.Itoa(int(code)))
	}
	return ""
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
}

// 响应成功
func Response(c *gin.Context, store *persist.RedisStoreEx, code int32, v interface{}) {
	ret := ResponseModel{
		Code:    code,
		Message: CodeToMessage(store, code, c.GetHeader("lang")),
		Data:    v,
	}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResponsePage(c *gin.Context, store *persist.RedisStoreEx, code int32, v interface{}, total int64, current int) {
	ret := ResponsePageModel{
		Code:    code,
		Message: CodeToMessage(store, code, c.GetHeader("lang")),
		Data: ResponsePageData{
			List:    v,
			Total:   total,
			Current: current,
		},
	}
	ResJSON(c, http.StatusOK, &ret)
}

// ResponseV2 响应成功（通过配置文件是否返回真实错误信息）
func ResponseV2(c *gin.Context, store *persist.RedisStoreEx, err error, code int32, v interface{}) {
	var msg string
	isShowErrMsg := c.GetInt("isShowErrMsg")
	if isShowErrMsg == 1 && err != nil {
		msg = err.Error()
	} else {
		if err != nil && code == 0 {
			msg = err.Error()
			code = int32(ERROR_FAIL.Code)
		} else {
			msg = CodeToMessage(store, code, c.GetHeader("lang"))
		}
	}
	ret := ResponseModel{
		Code:    code,
		Message: msg,
		Data:    v,
	}
	ResJSON(c, http.StatusOK, &ret)
}

// ResponsePageV2 响应成功
func ResponsePageV2(c *gin.Context, store *persist.RedisStoreEx, err error, code int32, v interface{}, total int64, current int) {
	var msg string
	isShowErrMsg := c.GetInt("isShowErrMsg")
	if isShowErrMsg == 1 {
		msg = err.Error()
	} else {
		msg = CodeToMessage(store, code, c.GetHeader("lang"))
	}
	ret := ResponsePageModel{
		Code:    code,
		Message: msg,
		Data: ResponsePageData{
			List:    v,
			Total:   total,
			Current: current,
		},
	}
	ResJSON(c, http.StatusOK, &ret)
}
