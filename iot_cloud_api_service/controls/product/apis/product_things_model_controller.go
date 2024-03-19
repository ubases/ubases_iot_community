package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"

	"github.com/gin-gonic/gin"
)

// productThingsModel
var ProductThingsModelcontroller ProductThingsModelController

type ProductThingsModelController struct{}

func (ProductThingsModelController) ResetThingsModel(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	productIdInt, err := iotutil.ToInt64AndErr(productId)
	if err != nil {
		iotgin.ResBadRequest(c, "productId is "+productId)
		return
	}
	res, err := rpc.ClientProductService.ResetProductThingModels(controls.WithUserContext(c), &protosService.TPmProductRequest{
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
	iotgin.ResSuccessMsg(c)
}
