package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/iotgin"
	"github.com/gin-gonic/gin"
)

var ProductAppRelationcontroller OpmProductAppRelationController

type OpmProductAppRelationController struct{} //部门操作控制器

var productAppServices = apiservice.OpmProductAppRelationService{}

// 查询信息
func (OpmProductAppRelationController) BindAppList(c *gin.Context) {
	var (
		productId = c.Query("productId")
		appKey    = c.Query("appKey")
	)
	res, err := productAppServices.SetContext(controls.WithUserContext(c)).QueryList(productId, appKey, controls.GetTenantId(c))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 编辑信息
func (OpmProductAppRelationController) ProductBindApp(c *gin.Context) {
	var req entitys.OpmProductAppRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productAppServices.SetContext(controls.WithUserContext(c)).BindRelation(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}
