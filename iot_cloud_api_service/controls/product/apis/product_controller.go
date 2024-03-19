package apis

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
)

var Productcontroller ProductController

// ProductController ctrl
type ProductController struct{}

var productService services.ProductService = services.ProductService{}

func (ct *ProductController) Create(c *gin.Context) {
	var (
		req  = new(entitys.CreateProductForm)
		data interface{}
		err  error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, iotutil.ToString(data))
		}
	}()
	c.ShouldBindJSON(&req)

	if data, err = productService.CreateProduct(req); err != nil {
		return
	}
}

func (ct *ProductController) Update(c *gin.Context) {
	var (
		req = new(entitys.UpProductForm)
		err error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccessMsg(c)
		}
	}()

	c.ShouldBind(&req)
	if iotutil.IsEmpty(req.ModelId) {
		err = fmt.Errorf("%s 不能为空", "modelId")
		return
	}
	req.Status = req.IsPublish

	if err = productService.UpdateProduct(req); err != nil {
		return
	}
}

// 发布/停用产品
func (ct *ProductController) Status(c *gin.Context) {
	var (
		req = new(entitys.UpProductForm)
		nId int
		err error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccessMsg(c)
		}
	}()

	c.ShouldBind(&req)

	nId, err = strconv.Atoi(req.Id)
	if err != nil {
		return
	}
	_, err = rpc.ClientProductService.UpdateStatus(context.Background(), &protosService.TPmProductStatusRequest{
		Id:     int64(nId),
		Status: req.Status,
	})
	if err == nil {
		cached.RedisStore.Delete(iotconst.OPEN_PRODUCT_TREE_DATA)
	}
	//if err = productService.UpdateProduct(req); err != nil {
	//	return
	//}
}

func (ct *ProductController) Get(c *gin.Context) {
	var (
		data interface{}
		key  string
		err  error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, data)
		}
	}()
	if key = c.Query("id"); len(key) == 0 {
		err = fmt.Errorf("%s 不能为空", "id")
		return
	}
	if data, err = productService.GetProduct(key); err != nil {
		return
	}
}

func (ct *ProductController) GetProductList(c *gin.Context) {
	var req entitys.QueryProductForm
	err := c.ShouldBind(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	data, total, err := productService.GetProductList(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, data, total, req.Page)
}

func (ct *ProductController) Delete(c *gin.Context) {
	var (
		req  = new(entitys.UpProductForm)
		data interface{}
		err  error
		key  int64
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, data)
		}
	}()

	c.ShouldBind(&req)
	if key = iotutil.ToInt64(req.Id); key <= 0 {
		err = fmt.Errorf("%s 不能为空", "id")
		return
	}
	if err = productService.DelProduct(key); err != nil {
		return
	}
}

func (ct *ProductController) GetStandardThingModelDetail(c *gin.Context) {
	var (
		data          interface{}
		productTypeId string
		ProductKey    string
		err           error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, data)
		}
	}()

	if productTypeId = c.Query("productTypeId"); len(productTypeId) == 0 {
		err = fmt.Errorf("%s 不能为空", "productTypeId")
		return
	}
	ProductKey = c.Query("productKey")

	if len(ProductKey) == 0 {
		if data, err = productService.GetStandardThingModelDetail(productTypeId); err != nil {
			return
		}
	} else {
		if data, err = productService.GetProductThingModelDetail(ProductKey); err != nil {
			return
		}
	}
}

// GetDefaultNetworkGuides
func (ct *ProductController) GetDefaultNetworkGuides(c *gin.Context) {
	var (
		data interface{}
		err  error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, data)
		}
	}()

	if data, err = productService.GetDefaultNetworkGuides(); err != nil {
		return
	}
}

func (ct *ProductController) CheckExists(c *gin.Context) {
	var (
		data   interface{}
		name   string
		nameEn string
		id     string
		err    error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResSuccess(c, data)
		}
	}()
	name = c.Query("name")
	nameEn = c.Query("nameEn")
	if len(name) == 0 && len(nameEn) == 0 {
		err = fmt.Errorf("%s 不能为空", "name")
		return
	}
	if id = c.Query("id"); len(id) == 0 {
		err = fmt.Errorf("%s 不能为空", "id")
		return
	}
	if data, err = productService.CheckExists(id, name, nameEn); err != nil {
		return
	}
}
