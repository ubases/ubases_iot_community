package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var ProductFirmwarecontroller ProductFirmwareController

type ProductFirmwareController struct{} //部门操作控制器

// 查询固件列表
func (ProductFirmwareController) QueryDropDownList(c *gin.Context) {
	productIdInt, err := iotutil.ToInt64AndErr(c.Query("productId"))
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	firmwareTypeInt, err := iotutil.ToInt32Err(c.Query("type"))
	if err != nil {
		iotgin.ResBadRequest(c, "type")
		return
	}
	productRes, err := rpc.ClientOpmProductModuleRelationService.
		QueryProductFirmwareList(context.Background(), &protosService.ProductFirmwareFilter{
			ProductId: productIdInt, FirmwareType: firmwareTypeInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if productRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(productRes.Message))
		return
	}
	iotgin.ResSuccess(c, productRes.Data)
}

// 查询固件版本列表
func (ProductFirmwareController) QueryDropDownVersionList(c *gin.Context) {
	productIdInt, err := iotutil.ToInt64AndErr(c.Query("productId"))
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	isCustomInt, err := iotutil.ToInt32Err(c.Query("isCustom"))
	if err != nil {
		iotgin.ResBadRequest(c, "isCustom")
		return
	}
	moduleIdInt, _ := iotutil.ToInt64AndErr(c.Query("moduleId"))
	firmwareIdInt, err := iotutil.ToInt64AndErr(c.Query("firmwareId"))
	if err != nil {
		iotgin.ResBadRequest(c, "firmwareId")
		return
	}
	if moduleIdInt == 0 && isCustomInt == 2 {
		//如果没有ModuleId就从关联记录中查询，一般用于云模组固件
		pmRels, err := rpc.ClientOpmProductModuleRelationService.Lists(context.Background(), &protosService.OpmProductModuleRelationListRequest{
			Query: &protosService.OpmProductModuleRelation{ProductId: productIdInt, IsCustom: isCustomInt},
		})
		if err == nil {
			for _, pmRel := range pmRels.Data {
				if pmRel.FirmwareId == firmwareIdInt {
					moduleIdInt = pmRel.ModuleId
					break
				}
			}
		}
	}
	productRes, err := rpc.ClientOpmProductModuleRelationService.
		QueryProductFirmwareVersionList(context.Background(), &protosService.ProductFirmwareVersionFilter{
			ProductId: productIdInt, ModuleId: moduleIdInt, FirmwareId: firmwareIdInt, IsCustom: isCustomInt})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if productRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(productRes.Message))
		return
	}
	//返回前端
	res := make([]*entitys.ProductFirmwareItemRes, 0)
	for _, p := range productRes.Data {
		res = append(res, &entitys.ProductFirmwareItemRes{
			Id:        p.Id,
			Name:      p.Name,
			IsCurrent: p.IsCurrent,
			IsCustom:  p.IsCustom,
			ModuleId:  p.ModuleId,
			IsMust:    p.IsMust,
		})
	}
	iotgin.ResSuccess(c, res)
}
