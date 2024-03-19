package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotutil"
	"errors"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var OtaPkgcontroller OpmOtaPkgController

type OpmOtaPkgController struct{} //部门操作控制器

var otaPkgServices = apiservice.OpmOtaPkgService{}

func (OpmOtaPkgController) QueryList(c *gin.Context) {
	var filter entitys.OpmOtaPkgQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmOtaPkgFilter)
	}
	res, total, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmOtaPkgList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmOtaPkgController) QueryDropDownList(c *gin.Context) {
	var filter entitys.OpmOtaPkgQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmOtaPkgList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmOtaPkgController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).GetOpmOtaPkgDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmOtaPkgController) Edit(c *gin.Context) {
	var req entitys.OpmOtaPkgEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).UpdateOpmOtaPkg(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmOtaPkgController) Add(c *gin.Context) {
	var req entitys.OpmOtaPkgEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).AddOpmOtaPkg(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmOtaPkgController) Delete(c *gin.Context) {
	var req entitys.OpmOtaPkgFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, _ := iotutil.ToInt64AndErr(req.Id)
	if id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = otaPkgServices.SetContext(controls.WithOpenUserContext(c)).DeleteOpmOtaPkg(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmOtaPkgController) SetStatus(c *gin.Context) {
	var req entitys.OpmOtaPkgFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, _ := iotutil.ToInt64AndErr(req.Id)
	if id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	otaPkgServices.Ctx = controls.WithOpenUserContext(c)
	err = otaPkgServices.SetStatusOpmOtaPkg(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetShelf 设置状态
func (OpmOtaPkgController) SetShelf(c *gin.Context, status int32) {
	var req entitys.OpmOtaPkgFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, _ := iotutil.ToInt64AndErr(req.Id)
	if id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	req.Status = status // 已上架
	otaPkgServices.Ctx = controls.WithOpenUserContext(c)
	err = otaPkgServices.SetStatusOpmOtaPkg(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

//发布

// OtaPublish 固件OTA发布
func (OpmOtaPkgController) OtaPublish(c *gin.Context) {
	var req entitys.OtaReleaseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.ProductID == 0 {
		iotgin.ResFailCode(c, "产品编号不能为空", -1)
		return
	}
	if req.OtaPkgId == 0 {
		iotgin.ResFailCode(c, "固件OTA编号不能为空", -1)
		return
	}
	missDids, err := otaPkgServices.SetContext(controls.WithUserContext(c)).OtaPublish(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, missDids)
}

// OtaPublishStop OTA暂停
func (OpmOtaPkgController) OtaPublishStop(c *gin.Context) {
	var req entitys.OpmOtaPublishFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	err = otaPkgServices.SetContext(controls.WithUserContext(c)).OtaPublishStop(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpmOtaPkgController) OtaPublishQueryList(c *gin.Context) {
	var filter entitys.OpmOtaPublishQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmOtaPublishFilter)
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	res, total, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).QueryOpmOtaPublishList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// OtaRecoveryPublish 设置状态
func (OpmOtaPkgController) OtaRecoveryPublish(c *gin.Context) {
	var req entitys.OpmOtaPublishFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	otaPkgServices.Ctx = controls.WithOpenUserContext(c)
	err = otaPkgServices.OtaRecoveryPublish(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpmOtaPkgController) QueryOtaVersions(c *gin.Context) {
	productIdInt, err := iotutil.ToInt64AndErr(c.Query("productId"))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	firmwareIdInt, _ := iotutil.ToInt64AndErr(c.Query("firmwareId"))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).QueryOtaVersions(productIdInt, firmwareIdInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmOtaPkgController) QueryOtaAreas(c *gin.Context) {
	productIdInt, err := iotutil.ToInt64AndErr(c.Query("productId"))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	firmwareIdInt, _ := iotutil.ToInt64AndErr(c.Query("firmwareId"))
	res, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).QueryOtaAreas(productIdInt, firmwareIdInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmOtaPkgController) OtaResultList(c *gin.Context) {
	publishIdInt64, err := iotutil.ToInt64AndErr(c.Param("publishId"))
	if err != nil {
		iotgin.ResBadRequest(c, "publishId")
		return
	}
	page, err := iotutil.ToInt64AndErr(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := iotutil.ToInt64AndErr(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	status, err := iotutil.ToInt32Err(c.Query("status"))
	if err != nil {
		status = 0
	}
	result, err := iotutil.ToInt32Err(c.Query("result"))
	if err != nil {
		result = 0
	}
	area := c.Query("area")

	publishList, _, err := otaPkgServices.SetContext(controls.WithUserContext(c)).QueryOpmOtaPublishList(entitys.OpmOtaPublishQuery{
		Query: &entitys.OpmOtaPublishFilter{Id: publishIdInt64},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(publishList) == 0 {
		iotgin.ResErrCli(c, errors.New("发布记录不存在"))
		return
	}
	deviceId := c.DefaultQuery("deviceId", "")
	isGray := publishList[0].IsGray
	version := publishList[0].Version
	//如果不是灰度，则
	//if isGray == 1 {
	//	publishIdInt64 = 0
	//}
	resMap, err := otaPkgServices.SetContext(controls.WithOpenUserContext(c)).
		QueryOtaResultList(deviceId, version, isGray, publishIdInt64, status, result, area, page, limit)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	resMap["details"] = publishList[0]
	//返回值
	iotgin.ResSuccess(c, resMap)
}
