package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var CommunityProductcontroller OpmCommunityProductController

type OpmCommunityProductController struct{} //部门操作控制器

var communityProductServices = apiservice.OpmCommunityProductService{}

// 查询社区产品信息
func (OpmCommunityProductController) QueryList(c *gin.Context) {
	var filter entitys.OpmCommunityProductQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.OpmCommunityProductFilter{}
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	res, total, err := communityProductServices.SetContext(controls.WithUserContext(c)).QueryOpmCommunityProductList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 查询社区产品详细信息
func (OpmCommunityProductController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := communityProductServices.SetContext(controls.WithUserContext(c)).GetOpmCommunityProductDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 编辑社区产品信息
func (OpmCommunityProductController) Edit(c *gin.Context) {
	var req entitys.OpmCommunityProductEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UpdatedBy = controls.GetUserId(c)
	req.UpdatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	id, err := communityProductServices.SetContext(controls.WithUserContext(c)).UpdateOpmCommunityProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 新增社区产品信息
func (OpmCommunityProductController) Add(c *gin.Context) {
	var req entitys.OpmCommunityProductEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.CreatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	id, err := communityProductServices.SetContext(controls.WithUserContext(c)).AddOpmCommunityProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 删除社区产品信息
func (OpmCommunityProductController) Delete(c *gin.Context) {
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
	err = communityProductServices.SetContext(controls.WithUserContext(c)).DeleteOpmCommunityProduct(entitys.OpmCommunityProductFilter{Id: idInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmCommunityProductController) SetStatus(c *gin.Context) {
	var req entitys.OpmCommunityProductFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = communityProductServices.SetContext(controls.WithUserContext(c)).SetStatusOpmCommunityProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
