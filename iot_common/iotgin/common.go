package iotgin

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE       = 0     //成功的状态码
	BAD_REQUEST        = 1     //参数错误
	BUSINESS_ERR       = 2     //业务错误
	PERMISSION_ERR     = 4     //权限错误
	FAIL_CODE          = -1    //失败的状态码
	USER_UID_KEY       = "UID" //页面UUID键名
	SUPER_ADMIN_ID     = "1"   //超级管理员
	FAIL_TENANTID_CODE = -2    //失败的状态码
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponsePageModel struct {
	Code    int              `json:"code"`
	Message string           `json:"msg"`
	Data    ResponsePageData `json:"data,omitempty"`
}

type ResponsePageData struct {
	List    interface{} `json:"list,omitempty"`
	Total   int64       `json:"total"`
	Current int         `json:"page"`
}

// 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
}

// 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
	//c.JSON(ret.Code, ret)
	//c.JSON(200, gin.H{"code": SUCCESS_CODE, "": "success", "data": v})
	//c.Writer()
}

// 响应成功,指定code
func ResSuccessWithCode(c *gin.Context, code int, v interface{}) {
	ret := ResponseModel{Code: code, Message: "ok", Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResPageSuccess(c *gin.Context, v interface{}, total int64, current int) {
	ret := ResponsePageModel{Code: SUCCESS_CODE, Message: "ok", Data: ResponsePageData{
		List:    v,
		Total:   total,
		Current: current,
	}}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResSuccessMsg(c *gin.Context) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: "ok"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应成功
func ResSuccessDataAndMsg(c *gin.Context, v interface{}, msg string) {
	ret := ResponseModel{Code: SUCCESS_CODE, Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 参数错误
func ResBadRequest(c *gin.Context, msg string) {
	ret := ResponseModel{Code: BAD_REQUEST, Message: "参数绑定错误: " + msg}
	ResJSON(c, http.StatusOK, &ret)
}

func ResErrParams(c *gin.Context) {
	ret := ResponseModel{Code: BAD_REQUEST, Message: "参数错误"}
	ResJSON(c, http.StatusOK, &ret)
}

// 业务错误
func ResBusinessP(c *gin.Context, msg string) {
	ret := ResponseModel{Code: BUSINESS_ERR, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-服务端故障
func ResErrSrv(c *gin.Context) {
	ret := ResponseModel{Code: FAIL_CODE, Message: "服务端故障"}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应失败
func ResFailCode(c *gin.Context, msg string, code int) {
	ret := ResponseModel{Code: code, Message: msg}
	ResJSON(c, http.StatusOK, &ret)
}

func ResFailCustomCode(c *gin.Context, msg string, code int, v interface{}) {
	ret := ResponseModel{Code: code, Message: msg, Data: v}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCli(c *gin.Context, err error) {
	var ret ResponseModel
	// TODO 临时处理方案，将go.micro.client 相关错误的detail提取返回前端。
	if strings.Index(err.Error(), "go.micro.client") != -1 {
		iotlogger.LogHelper.Error(err)
		//ret = ResponseModel{Code: FAIL_CODE, Message: "内部服务器错误"}
		errMap := iotutil.JsonToMap(err.Error())
		if msg, ok := errMap["detail"]; ok {
			ret = ResponseModel{Code: FAIL_CODE, Message: msg.(string)}
		} else {
			ret = ResponseModel{Code: FAIL_CODE, Message: err.Error()}
		}
	} else {
		ret = ResponseModel{Code: FAIL_CODE, Message: err.Error()}
	}
	ResJSON(c, http.StatusOK, &ret)
}

// 响应错误-用户端故障
func ResErrCliCustomCode(c *gin.Context, err error, code int) {
	var ret ResponseModel
	// TODO 临时处理方案，将go.micro.client 相关错误的detail提取返回前端。
	if strings.Index(err.Error(), "go.micro.client") != -1 {
		iotlogger.LogHelper.Error(err)
		//ret = ResponseModel{Code: FAIL_CODE, Message: "内部服务器错误"}
		errMap := iotutil.JsonToMap(err.Error())
		if msg, ok := errMap["detail"]; ok {
			ret = ResponseModel{Code: code, Message: msg.(string)}
		} else {
			ret = ResponseModel{Code: code, Message: err.Error()}
		}
	} else {
		ret = ResponseModel{Code: code, Message: err.Error()}
	}
	ResJSON(c, http.StatusOK, &ret)
}

// ResErrCliExt 响应错误-用户端故障
func ResErrCliExt(c *gin.Context, err error, msg string) {
	if err == nil {
		err = errors.New(msg)
	}
	var ret ResponseModel
	// TODO 临时处理方案，将go.micro.client 相关错误的detail提取返回前端。
	if strings.Index(err.Error(), "go.micro.client") != -1 {
		iotlogger.LogHelper.Error(err)
		//ret = ResponseModel{Code: FAIL_CODE, Message: "内部服务器错误"}
		errMap := iotutil.JsonToMap(err.Error())
		if msg, ok := errMap["detail"]; ok {
			ret = ResponseModel{Code: FAIL_CODE, Message: msg.(string)}
		} else {
			ret = ResponseModel{Code: FAIL_CODE, Message: err.Error()}
		}
	} else {

	}
	ResJSON(c, http.StatusOK, &ret)
}
