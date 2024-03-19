package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/document/services"
	"cloud_platform/iot_common/iotgin"
	"strings"

	"github.com/gin-gonic/gin"
)

var ProductHelpDoccontroller ProductHelpDocController

type ProductHelpDocController struct {
}

var productHelpDocService = services.ProductHelpDocService{}

// GetProductHelpDoc 获取产品帮助文档
// @Summary 获取问题类型
// @Description
// @Tags Document
// @Accept application/json
// @Param productKey path string true "产品Key"
// @Param tenantId header string true "租户Id"
// @Param lang header string true "所属语言"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /productHelpDoc/list/{productKey} [get]
func (ProductHelpDocController) GetProductHelpDoc(c *gin.Context) {
	productKey := c.Param("productKey")
	if strings.TrimSpace(productKey) == "" {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	tenantId := c.GetHeader("tenantId")
	lang := c.GetHeader("lang")
	res, err := productHelpDocService.SetContext(controls.WithUserContext(c)).GetProductHelpDoc(productKey, tenantId, lang)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
