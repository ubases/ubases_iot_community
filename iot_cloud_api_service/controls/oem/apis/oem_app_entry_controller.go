package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/oem/services"
	entSys "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiPubService "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var OemAppEntrycontroller OemAppEntryController

var entryApp apiservice.OemAppEntryService

var entryPubApp apiPubService.SysAppEntryService

type OemAppEntryController struct {
} //用户操作控制器

// 词条列表
func (OemAppEntryController) EntryList(c *gin.Context) {
	var req entitys.OemAppEntryListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntryList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(req.Page))
}

// 公版词条列表
func (OemAppEntryController) EntryPubList(c *gin.Context) {
	var req entSys.SysAppEntryListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := entryPubApp.SetContext(controls.WithOpenUserContext(c)).EntryList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(req.Page))
}

// 保存词条
func (OemAppEntryController) EntrySave(c *gin.Context) {
	var req entitys.OemAppEntrySaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntrySave(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 词条详细
func (OemAppEntryController) EntryDetail(c *gin.Context) {
	setingId := c.Query("setingId")
	lang := c.Query("lang")
	if setingId == "" || lang == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}

	id, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntryDetail(setingId, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 公版词条详细
func (OemAppEntryController) EntryPubDetail(c *gin.Context) {
	setingId := c.Query("setingId")
	lang := c.Query("lang")
	if setingId == "" || lang == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}

	id, err := entryPubApp.SetContext(controls.WithOpenUserContext(c)).EntryDetail(setingId, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 保存词条设置
func (OemAppEntryController) EntrySetingSave(c *gin.Context) {
	var req entitys.OemAppEntrySetingSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntrySetingSave(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OemAppEntryController) EntrySetingDetail(c *gin.Context) {
	setingId := c.Query("setingId")
	if setingId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	id, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntrySetingDetail(setingId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OemAppEntryController) EntryDelete(c *gin.Context) {
	setingId := c.Query("setingId")
	if setingId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	id, err := entryApp.SetContext(controls.WithOpenUserContext(c)).EntryDelete(setingId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}
