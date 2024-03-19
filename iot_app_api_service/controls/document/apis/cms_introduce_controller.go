package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/document/services"
	"cloud_platform/iot_common/iotgin"
	"github.com/gin-gonic/gin"
)

var Introducecontroller IntroduceController

type IntroduceController struct {
} //用户操作控制器

var introduceService = services.IntroduceService{}

// GetIntroduceDetailByApp 获取协议文档详情
// @Summary 获取协议文档详情
// @Description
// @Tags Document
// @Accept application/json
// @Param id path string true "文档编码"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /introduce/detail/{id} [post]
//func (IntroduceController) GetIntroduceDetailByApp(c *gin.Context) {
//	id := c.Param("id")
//	if id == "" {
//		iotgin.ResBadRequest(c, "id")
//		return
//	}
//	idi, err := strconv.ParseInt(id, 10, 64)
//	if err != nil {
//		iotgin.ResErrCli(c, err)
//		return
//	}
//	res, err := introduceService.SetContext(controls.WithUserContext(c)).GetIntroduceDetail(idi)
//	if err != nil {
//		iotgin.ResErrCli(c, err)
//		return
//	}
//	iotgin.ResSuccess(c, res)
//}

// GetIntroduceByApp 获取协议文档列表
// @Summary 获取协议文档详情
// @Description
// @Tags Document
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /introduce [get]
//func (IntroduceController) GetIntroduceByApp(c *gin.Context) {
//	res, err := introduceService.SetContext(controls.WithUserContext(c)).GetAppIntroduceByCode([]string{"AboutUs", "PrivacyPolicy", "UserAgreement"})
//	if err != nil {
//		iotgin.ResErrCli(c, err)
//		return
//	}
//	iotgin.ResSuccess(c, res)
//}

// GetIntroduceDetailByCode 获取协议文档详情
// @Summary 获取协议文档详情
// @Description
// @Tags Document
// @Accept application/json
// @Param id path string true "文档编码"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /introduce/{code} [get]
func (IntroduceController) GetIntroduceDetailByCode(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		iotgin.ResBadRequest(c, "code")
		return
	}
	lang := c.GetHeader("lang")
	appKey := c.GetHeader("appKey")
	res, _ := introduceService.SetContext(controls.WithUserContext(c)).GetIntroduceDetailByCode(appKey, code, lang)
	iotgin.ResSuccess(c, res)
}
