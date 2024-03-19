package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"errors"

	"github.com/gin-gonic/gin"
)

var SysAppEntrycontroller SysAppEntryController

var entryApp apiservice.SysAppEntryService

type SysAppEntryController struct {
} //用户操作控制器

// 保存词条
func (SysAppEntryController) EntryList(c *gin.Context) {
	var req entitys.SysAppEntryListReq
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

// 保存词条
func (SysAppEntryController) EntrySave(c *gin.Context) {
	var req entitys.SysAppEntrySaveReq
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
func (SysAppEntryController) EntryDetail(c *gin.Context) {
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

// 保存词条设置
func (SysAppEntryController) EntrySetingSave(c *gin.Context) {
	var req entitys.SysAppEntrySetingSaveReq
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

func (SysAppEntryController) EntrySetingDetail(c *gin.Context) {
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

func (SysAppEntryController) EntryDelete(c *gin.Context) {
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
