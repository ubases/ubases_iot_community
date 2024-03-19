package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/community/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"errors"

	"github.com/gin-gonic/gin"
)

var CommunityProductcontroller CommunityProductController

type CommunityProductController struct{} //部门操作控制器

// GetCommunityProductDetail 获取社区产品详情
// @Summary 获取社区产品详情
// @Description 获取社区产品详情
// @Tags 通用
// @Param id query string true "产品Id"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app//community/product/detail [get]
func (s CommunityProductController) GetCommunityProductDetail(c *gin.Context) {
	productId := c.Query("id")
	if productId != "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	productIdInt, err := iotutil.ToInt64AndErr(productId)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}

	res, err := rpc.ClientOpmCommunityProductService.FindById(controls.WithUserContext(c), &proto.OpmCommunityProductFilter{
		Id: productIdInt,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}

	resData := entitys.OpmCommunityProduct_pb2e(res.Data[0])
	lang := controls.GetLang(c)
	for _, language := range res.Data[0].Langs {
		if lang == language.Lang {
			resData.ProductName = language.Name
			resData.ProductDesc = language.Description
			break
		}
	}
	iotgin.ResSuccess(c, resData)
}

// GetCommunityProductList 获取社区产品列表
// @Summary 获取社区产品列表
// @Description 获取社区产品列表
// @Tags 通用
// @Param tenantId query string true "租户Id"
// @Success 200 {object} iotgin.ResponseModel "{"code": 0, "data": [...]}"
// @Router /v1/platform/app/community/product/list [post]
func (s CommunityProductController) GetCommunityProductList(c *gin.Context) {
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
	res, err := rpc.ClientOpmCommunityProductService.Lists(controls.WithUserContext(c), &proto.OpmCommunityProductListRequest{
		Page:      int64(filter.Page),
		PageSize:  int64(filter.Limit),
		OrderDesc: filter.Sort,
		OrderKey:  filter.SortField,
		Query:     filter.OpmCommunityProductQuery_e2pb(),
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}

	lang := controls.GetLang(c)
	resList := make([]*entitys.OpmCommunityProductEntitys, 0)
	for _, d := range res.Data {
		resData := entitys.OpmCommunityProduct_pb2e(d)
		for _, language := range d.Langs {
			if lang == language.Lang {
				resData.ProductName = language.Name
				resData.ProductDesc = language.Description
				break
			}
		}
		resList = append(resList, resData)
	}
	iotgin.ResPageSuccess(c, resList, res.Total, int(filter.Page))
}
