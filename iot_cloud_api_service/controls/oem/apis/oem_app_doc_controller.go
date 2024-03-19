package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	entitys2 "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var OemAppDoccontroller OemAppDocController

var docApp apiservice.OemAppDocService

type OemAppDocController struct {
} //用户操作控制器

func (OemAppDocController) DocList(c *gin.Context) {
	tenantId, _ := c.Get("tenantId")
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).DocList(iotutil.ToString(tenantId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取公版文档列表
func (OemAppDocController) GetPubDocList(c *gin.Context) {
	var req entitys2.SysAppHelpCenterQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	tenantId, _ := c.Get("tenantId")
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).GetHelpCenterListForOpen(iotutil.ToString(tenantId), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, int64(len(res)), int(req.Page))
}

// 创建app文档
func (OemAppDocController) CreateDoc(c *gin.Context) {
	var req entitys.OemAppDocSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := docApp.SetContext(controls.WithOpenUserContext(c)).CreateDoc(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 创建app文档
func (OemAppDocController) UpdateDoc(c *gin.Context) {
	var req entitys.OemAppDocSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := docApp.SetContext(controls.WithOpenUserContext(c)).UpdateDoc(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 创建app文档
func (OemAppDocController) DeleteDoc(c *gin.Context) {
	var req entitys.OemAppDocSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = docApp.SetContext(controls.WithOpenUserContext(c)).DeleteDoc(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, nil)
}

func (OemAppDocController) GetApps(c *gin.Context) {
	tenantId, _ := c.Get("tenantId")
	docId := c.Query("docId")
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).GetApps(iotutil.ToString(tenantId), docId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppDocController) GetPubLangs(c *gin.Context) {
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).GetPubLangs()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppDocController) GetDocSupportLangs(c *gin.Context) {

	docId := c.Query("docId")
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).GetDocSupportLangs(docId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OemAppDocController) DetailDoc(c *gin.Context) {

	docId := c.Query("docId")
	res, err := docApp.SetContext(controls.WithOpenUserContext(c)).DetailDoc(docId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
