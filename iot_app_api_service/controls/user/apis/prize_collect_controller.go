package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"

	"github.com/gin-gonic/gin"
)

var PrizeCollectcontroller PrizeCollectController

type PrizeCollectController struct {
} //家庭操作控制器

var prizeServices = services.AppPrizeCollectService{}

// @Summary 添加App有奖征集
// @Description
// @Tags APP
// @Accept application/json
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/app/marketing/prizecollect/add [post]
func (PrizeCollectController) Add(c *gin.Context) {
	req := entitys.UcUserPrizeCollectEntitys{}
	if err := c.BindJSON(&req); err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId := controls.GetUserId(c)
	err := prizeServices.SetContext(controls.WithUserContext(c)).AddPrizeCollect(req, iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
