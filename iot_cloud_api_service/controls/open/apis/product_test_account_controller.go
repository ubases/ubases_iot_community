package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"github.com/gin-gonic/gin"
	goerrors "go-micro.dev/v4/errors"
)

var ProductTestAccountcontroller OpmProductTestAccountController

type OpmProductTestAccountController struct{} //部门操作控制器

var productTestAccountServices = apiservice.OpmProductTestAccountService{}

// QueryList APP测试账号列表
// @Summary APP测试账号列表
// @Description
// @Tags APP测试账号
// @Accept application/json
// @Param data body entitys.OpmProductTestAccountQuery true "请求参数"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/productTestAccount/list [post]
func (s OpmProductTestAccountController) QueryList(c *gin.Context) {
	var filter entitys.OpmProductTestAccountQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind query ProductTestAccount param error: ", err)
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	res, total, err := productTestAccountServices.SetContext(controls.WithUserContext(c)).QueryOpmProductTestAccountList(filter)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("query ProductTestAccount error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// Add 新增APP测试账号
// @Summary 新增APP测试账号
// @Description
// @Tags APP测试账号
// @Accept application/json
// @Param data form entitys.OpmProductTestAccountEntitys true "请求参数"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/productTestAccount/add [post]
func (s *OpmProductTestAccountController) Add(c *gin.Context) {
	var req entitys.OpmProductTestAccountEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("bind create ProductTestAccount param error: ", err)
		ioterrs.ResponseV2(c, cached.RedisStore, err, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	if err := req.AddCheck(); err != nil {
		iotlogger.LogHelper.Helper.Error("check create ProductTestAccount param error: ", err)
		ioterrs.ResponseV2(c, cached.RedisStore, err, goerrors.FromError(err).GetCode(), nil)
		return
	}

	req.CreatedBy = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	_, err = productTestAccountServices.SetContext(controls.WithOpenUserContext(c)).AddOpmProductTestAccount(req)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("create ProductTestAccount error: ", err)
		ioterrs.ResponseV2(c, cached.RedisStore, err, goerrors.FromError(err).GetCode(), nil)
		return
	}
	ioterrs.Response(c, cached.RedisStore, ioterrs.Success, nil)
}

// Delete 删除APP测试账号
// @Summary 删除APP测试账号
// @Description
// @Tags APP测试账号
// @Accept application/json
// @Param id path string true "APP测试账号Id"
// @Param data body entitys.OpmProductTestAccountEntitys true "请求参数"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /v1/platform/web/open/productTestAccount/delete/{id} [post]
func (OpmProductTestAccountController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotlogger.LogHelper.Helper.Error("query details ProductTestAccount param error: id")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	idInt, err := iotutil.ToInt64AndErr(id)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("query details ProductTestAccount param error: id format")
		ioterrs.Response(c, cached.RedisStore, ioterrs.ErrShouldBindJSON, nil)
		return
	}
	err = productTestAccountServices.SetContext(controls.WithUserContext(c)).DeleteOpmProductTestAccount(entitys.OpmProductTestAccountEntitys{Id: idInt})
	if err != nil {
		iotlogger.LogHelper.Helper.Error("delete ProductTestAccount error: ", err)
		ioterrs.Response(c, cached.RedisStore, goerrors.FromError(err).GetCode(), nil)
		return
	}
	iotgin.ResSuccessMsg(c)
}
