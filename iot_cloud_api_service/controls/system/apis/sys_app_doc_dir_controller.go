package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var SysAppDocDircontroller SysAppDocDirController

var docDirApp apiservice.SysAppDocDirService

type SysAppDocDirController struct {
} //用户操作控制器

// 创建文档目录
func (SysAppDocDirController) CreateDir(c *gin.Context) {
	var req entitys.SysAppDocDirSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).CreateDir(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 创建文档目录
func (SysAppDocDirController) UpdateDir(c *gin.Context) {
	var req entitys.SysAppDocDirSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).UpdateDir(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 创建文档目录
func (SysAppDocDirController) DetailDir(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).DetailDir(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (SysAppDocDirController) DeleteDir(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).DeleteDir(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (SysAppDocDirController) ListDir(c *gin.Context) {
	lang := c.Query("lang")
	if lang == "" {
		lang = c.GetHeader("lang")
	}
	if lang == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	helpId := c.Query("helpId")
	if helpId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).ListDir(lang, helpId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
