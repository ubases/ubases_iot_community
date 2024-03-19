package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var ThingModelRulecontroller OpmThingModelRuleController

type OpmThingModelRuleController struct{} //部门操作控制器

var thingModelRuleServices = apiservice.OpmThingModelRuleService{}

// 查询物模型规则设置列表
func (OpmThingModelRuleController) QueryList(c *gin.Context) {
	var filter entitys.OpmThingModelRuleQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.OpmThingModelRuleFilter{}
	}
	res, total, err := thingModelRuleServices.SetContext(controls.WithUserContext(c)).QueryOpmThingModelRuleList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 查询物模型规则设置列表
func (OpmThingModelRuleController) QueryListByProductId(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	dataOrigin := c.Query("dataOrigin")
	if dataOrigin == "" {
		iotgin.ResBadRequest(c, "dataOrigin")
		return
	}
	dataOriginInt, err := iotutil.ToInt32Err(dataOrigin)
	if err != nil {
		iotgin.ResBadRequest(c, "dataOrigin")
		return
	}
	res, _, err := thingModelRuleServices.SetContext(controls.WithUserContext(c)).QueryOpmThingModelRuleListByProductId(productId, dataOriginInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 查询规则设置详情
func (OpmThingModelRuleController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := thingModelRuleServices.SetContext(controls.WithUserContext(c)).GetOpmThingModelRuleDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 编辑规则设置信息
func (OpmThingModelRuleController) Edit(c *gin.Context) {
	var req entitys.OpmThingModelRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.UpdatedBy = controls.GetUserId(c)
	id, err := thingModelRuleServices.SetContext(controls.WithUserContext(c)).UpdateOpmThingModelRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 新增规则设置信息
func (OpmThingModelRuleController) Add(c *gin.Context) {
	var req entitys.OpmThingModelRuleEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.CreatedBy = controls.GetUserId(c)

	//dataOrigin不能为空
	//if req.DataOrigin == 0 {
	//	//iotgin.ResBadRequest(c, "dataOrigin")
	//	//return
	//}

	id, err := thingModelRuleServices.SetContext(controls.WithUserContext(c)).AddOpmThingModelRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 删除规则设置
func (OpmThingModelRuleController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	idInt, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		iotgin.ResBadRequest(c, "id format")
		return
	}
	err = thingModelRuleServices.SetContext(controls.WithUserContext(c)).DeleteOpmThingModelRule(entitys.OpmThingModelRuleFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmThingModelRuleController) SetStatus(c *gin.Context) {
	var req entitys.OpmThingModelRuleFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = thingModelRuleServices.SetContext(controls.WithUserContext(c)).SetStatusOpmThingModelRule(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
