package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	apisPubService "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var OemAppDocDircontroller OemAppDocDirController

var docDirApp apiservice.OemAppDocDirService

var docPubDirApp apisPubService.SysAppDocDirService

type OemAppDocDirController struct {
} //用户操作控制器

// 创建文档目录
func (OemAppDocDirController) CreateDir(c *gin.Context) {
	var req entitys.OemAppDocDirSaveReq
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
func (OemAppDocDirController) UpdateDir(c *gin.Context) {
	var req entitys.OemAppDocDirSaveReq
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
func (OemAppDocDirController) DetailDir(c *gin.Context) {
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

// 公版文档目录详细
func (OemAppDocDirController) DetailPubDir(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := docPubDirApp.SetContext(controls.WithOpenUserContext(c)).DetailDir(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppDocDirController) DeleteDir(c *gin.Context) {
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

func (OemAppDocDirController) ListDir(c *gin.Context) {
	docId := c.Query("docId")
	lang := c.Query("lang")

	if docId == "" || lang == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	res, err := docDirApp.SetContext(controls.WithOpenUserContext(c)).ListDir(docId, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 公版文档目录列表
func (OemAppDocDirController) ListPubDir(c *gin.Context) {
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
	res, err := docPubDirApp.SetContext(controls.WithOpenUserContext(c)).ListDir(lang, helpId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
