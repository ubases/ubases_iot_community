package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"

	"github.com/gin-gonic/gin"
)

var ProductVoicecontroller ProductVoiceController

var serviceApp apiservice.ProductViceService

type ProductVoiceController struct{}

// 保存产品语控配置
func (s ProductVoiceController) Save(c *gin.Context) {
	var req protosService.OpmVoiceProductSaveReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).Save(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 获取产品语控列表
func (s ProductVoiceController) GetList(c *gin.Context) {
	var req protosService.OpmVoiceProductListReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).GetList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取产品语控配置详细
func (s ProductVoiceController) GetDetail(c *gin.Context) {
	var req protosService.OpmVoiceProductDetailReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).GetDetail(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 产品语控配置发布
func (s ProductVoiceController) Publish(c *gin.Context) {
	var req protosService.OpmVoiceProductPublishReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).Publish(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 产品语控配置发布
func (s ProductVoiceController) GetVoiceDoc(c *gin.Context) {
	voiceNo := c.Query("voiceNo")

	if voiceNo == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).GetVoiceDoc(voiceNo)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 发布记录列表查询接口
func (s ProductVoiceController) GetVoicePublishRecord(c *gin.Context) {
	voiceNo := c.Query("voiceNo")
	productKey := c.Query("productKey")

	if voiceNo == "" && productKey == "" {
		iotgin.ResErrCli(c, errors.New("参数错误"))
		return
	}
	id, err := serviceApp.SetContext(controls.WithOpenUserContext(c)).GetVoicePublishRecord(voiceNo, productKey)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

type VoiceUnitReq struct {
	Unit string `json:"unit"`
}

// 获取语控单位列表
func (s ProductVoiceController) GetVoiceUnitList(c *gin.Context) {
	var req VoiceUnitReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	list, _ := iotconst.VoiceDataUnitList[req.Unit]
	iotgin.ResPageSuccess(c, list, 0, 0)
}
