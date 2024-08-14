package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_common/iotgin"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var OemAppCertcontroller OemAppCertController

var serviceCert apiservice.OemAppCertService

type OemAppCertController struct {
}

// 保存ios证书
func (OemAppCertController) SaveIosCert(c *gin.Context) {
	var req entitys.OemAppIosCertReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).SaveIosCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取ios证书
func (OemAppCertController) GetIosCert(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).GetIosCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 保存android证书
func (OemAppCertController) SaveAndroidCert(c *gin.Context) {
	var req entitys.OemAppAndroidCertSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).SaveAndroidCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取android证书
func (OemAppCertController) GetAndroidCert(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).GetAndroidCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 保存android证书
func (OemAppCertController) SaveIosPushCert(c *gin.Context) {
	var req entitys.OemAppIosPushCertSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).SaveIosPushCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取android证书
func (OemAppCertController) GetIosPushCert(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).GetIosPushCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 保存android证书
func (OemAppCertController) SaveAndroidPushCert(c *gin.Context) {
	var req entitys.OemAppAndroidPushCertSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).SaveAndroidPushCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取android证书
func (OemAppCertController) GetAndroidPushCert(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).GetAndroidPushCert(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取android证书
func (OemAppCertController) Regenerate(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	id, err := serviceCert.SetContext(controls.WithOpenUserContext(c)).Regenerate(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// DownloadSignCert 通过签名地址下载本地签名证书 app/downloadSignCert
func (OemAppCertController) DownloadSignCert(c *gin.Context) {
	signCertUrl := c.DefaultQuery("signCertUrl", "")
	if strings.TrimSpace(signCertUrl) == "" {
		iotgin.ResBadRequest(c, "signCertUrl")
		return
	}
	signCerts := strings.Split(signCertUrl, "/")
	fileName := signCerts[len(signCerts)-1:][0]
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(signCertUrl)
}
