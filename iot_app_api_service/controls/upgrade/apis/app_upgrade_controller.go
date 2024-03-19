package apis

import (
	"cloud_platform/iot_app_api_service/controls/upgrade/entitys"
	"cloud_platform/iot_app_api_service/controls/upgrade/services"
	"cloud_platform/iot_common/iotgin"

	"github.com/gin-gonic/gin"
)

var AppUpgradecontroller AppUpgradeController

type AppUpgradeController struct {
} //用户操作控制器

var AppUpgradeService = services.AppUpgradeService{}

// @Summary 根据app类型获取最新app版本
// @Description
// @Tags APP
// @Accept application/json
// @Param data body entitys.AppQueryAppUpgradeForm true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/appVersion/get [post]
func (AppUpgradeController) GetLatestApp(c *gin.Context) {
	appKey := c.GetHeader("appKey")
	res, err := AppUpgradeService.GetLatestApp(appKey)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GetFunctionConfigAutoUpgrade 获取APP自动升级功能配置
// @Summary 获取APP自动升级功能配置
// @Description
// @Tags APP
// @Accept application/json
// @Param data body entitys.OemAppCommonReq true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/functionConfig/autoUpgrade [get]
func (AppUpgradeController) GetFunctionConfigAutoUpgrade(c *gin.Context) {
	var req entitys.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	appKey := c.GetHeader("appKey")
	req.AppKey = appKey
	res, err := AppUpgradeService.GetFunctionConfigAutoUpgrade(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
