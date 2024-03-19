package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
)

var ProductTypecontroller ProductTypeController

// ProductTypeController ctrl
type ProductTypeController struct{}

var ProductTypeService services.ProductTypeService = services.ProductTypeService{}

func (ct *ProductTypeController) Create(c *gin.Context) {
	var (
		req  = new(entitys.CreateProductTypeForm)
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
	c.ShouldBind(&req)

	if data, err = ProductTypeService.CreateProductType(req); err != nil {
		return
	}
}

func (ct *ProductTypeController) Update(c *gin.Context) {
	var (
		req = new(entitys.UpProductTypeForm)
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

	if err = ProductTypeService.UpdateProductType(req); err != nil {
		return
	}
}

func (ct *ProductTypeController) Get(c *gin.Context) {
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
	if data, err = ProductTypeService.GetProductType(key); err != nil {
		return
	}
}

func (ct *ProductTypeController) GetProductTypeList(c *gin.Context) {
	var (
		data  interface{}
		total int
		req   = new(entitys.QueryProductTypeForm)
		err   error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResPageSuccess(c, data, int64(total), req.Page)
		}
	}()

	c.ShouldBind(&req)

	if data, total, err = ProductTypeService.GetProductTypeList(req); err != nil {
		return
	}
}

func (ct *ProductTypeController) Delete(c *gin.Context) {
	var (
		req  = new(entitys.UpProductTypeForm)
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
	if err = ProductTypeService.DelProductType(key); err != nil {
		return
	}
}

func (ct *ProductTypeController) GetTypeAndProductList(c *gin.Context) {
	var (
		data  interface{}
		total int
		req   = new(entitys.QueryProductTypeForm)
		err   error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResPageSuccess(c, data, int64(total), req.Page)
		}
	}()

	c.ShouldBind(&req)

	if data, total, err = ProductTypeService.GetProductTree(req); err != nil {
		return
	}
}

// GetExport 导出的get方法
func (ct *ProductTypeController) GetModelTemplate(c *gin.Context) {
	fileName := "modelTemplate.xlsx"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	tempPathFile := strings.Join([]string{iotconst.GetWorkTempDir(), fileName}, string(filepath.Separator))
	c.File(tempPathFile)
}
