package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"github.com/gin-gonic/gin"
)

var FailLogcontroller FailLogController

type FailLogController struct {
}

// ReportOperationFailLog 设备操作失败日志上报
// @Summary 设备操作失败日志上报
// @Description 设备操作失败日志上报
// @Tags 通用
// @Param data body entitys.DeviceOperationFailLogRequest true "上报数据"
// @Param appKey header string true "APP Key"
// @Param tenantId header string true "租户Id"
// @Param token header string true "token"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/community/product/list [post]
func (s *FailLogController) ReportOperationFailLog(c *gin.Context) {
	req := entitys.DeviceOperationFailLogRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var (
		userId   = controls.GetUserId(c)
		account  = controls.GetAccount(c)
		tenantId = controls.GetTenantId(c)
		appKey   = controls.GetAppKey(c)
	)
	//如果存在token，则解析获取token详情
	token := controls.GetToken(c)
	if token != "" {
		userInfo, err := controls.TokenGetUserInfo(token)
		if err == nil {
			userId = userInfo.UserID
			account = userInfo.Account
			tenantId = userInfo.TenantId
			appKey = userInfo.AppKey
		}
	}

	res, err := rpc.IotDeviceLogService.DeviceOperationFailLogReport(context.Background(), &protosService.DeviceOperationFailLogRequest{
		DeviceId:    req.DeviceId,
		Type:        req.Type,
		Content:     req.Content,
		UserId:      iotutil.ToString(userId),
		UserAccount: account,
		TenantId:    tenantId,
		AppKey:      appKey,
	})
	if err != nil || res.Code != 200 {
		iotgin.ResErrCliExt(c, err, res.Message)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// ReportOperationFailLogEx 设备操作失败日志上报
// @Summary 设备操作失败日志上报
// @Description 设备操作失败日志上报
// @Tags 通用
// @Param data body entitys.AppFailLog true "上报数据"
// @Param appKey header string true "APP Key"
// @Param tenantId header string true "租户Id"
// @Param tz header string true "时区"
// @Param region header string true "登录的区域，区域id"
// @Param lang header string true "app语言"
// @Param x-sys-info header string true "手机信息"
// @Param token header string true "token"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/community/product/list [post]
func (s *FailLogController) ReportOperationFailLogEx(c *gin.Context) {
	req := entitys.AppFailLog{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var (
		userId     = controls.GetUserId(c)        //用户id
		account    = controls.GetAccount(c)       //用户账号
		tenantId   = controls.GetTenantId(c)      //租户id，app开发者公司的租户id
		appKey     = controls.GetAppKey(c)        //app key
		Tz         = controls.GetTimezone(c)      //用户app 时区
		Region     = controls.GetRegionInt(c)     //用户app 登录的区域，区域id
		Lang       = controls.GetLang(c)          //用户app app语言
		sysinfo    = controls.GetSystemInfoRaw(c) //用户手机系统信息
		content, _ = jsoniter.Marshal(req.Content)
	)

	//读取token用户信息
	if userId == 0 || account == "" {
		token := controls.GetToken(c)
		if token != "" {
			var userInfo controls.UserInfo
			cacheData := iotredis.GetClient().Get(context.Background(), token)
			if cacheData.Err() == nil && cacheData.Val() != "" {
				userId = userInfo.UserID
				account = userInfo.Account
				appKey = userInfo.AppKey
				tenantId = userInfo.TenantId
			}
		}
	}

	var datalist []*protosService.DeviceOperationFailLogRequestEx
	obj := protosService.DeviceOperationFailLogRequestEx{
		DeviceId:     req.DeviceId,
		Type:         int32(req.Type),
		Content:      string(content),
		UserId:       strconv.Itoa(int(userId)),
		UserAccount:  account,
		TenantId:     tenantId,
		AppKey:       appKey,
		ProductKey:   req.ProductKey,
		FailTime:     req.Time,
		Code:         int32(req.Code),
		Timezone:     Tz,
		Region:       strconv.Itoa(int(Region)),
		Lang:         Lang,
		Sysinfo:      sysinfo,
		UploadFrom:   "app",
		UploadMethod: "http",
	}
	datalist = append(datalist, &obj)

	res, err := rpc.IotDeviceLogService.OperationFailLogReport(context.Background(), &protosService.OperationFailLogRequest{Data: datalist})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	iotgin.ResSuccessMsg(c)
}
