package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/product/entitys"
	"cloud_platform/iot_cloud_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
)

var NetworkGuideStepcontroller NetworkGuideStepController

// NetworkGuideStepController ctrl
type NetworkGuideStepController struct{}

var NetworkGuideStepService services.NetworkGuideStepService = services.NetworkGuideStepService{}

func (ct *NetworkGuideStepController) Create(c *gin.Context) {
	var (
		req  = new(entitys.PmNetworkGuideStepEntitys)
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

	if data, err = NetworkGuideStepService.CreateNetworkGuideStep(req); err != nil {
		return
	}
}

func (ct *NetworkGuideStepController) Update(c *gin.Context) {
	var (
		req = new(entitys.PmNetworkGuideStepEntitys)
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

	if err = NetworkGuideStepService.UpdateNetworkGuideStep(req); err != nil {
		return
	}
}

func (ct *NetworkGuideStepController) Get(c *gin.Context) {
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
	if data, err = NetworkGuideStepService.GetNetworkGuideStep(key); err != nil {
		return
	}
}

func (ct *NetworkGuideStepController) GetNetworkGuideStepList(c *gin.Context) {
	var (
		data  interface{}
		total int64
		req   = new(entitys.PmNetworkGuideStepQuery)
		err   error
	)
	defer func() {
		if err != nil {
			log.Error(err)
			iotgin.ResErrCli(c, err)
		} else {
			iotgin.ResPageSuccess(c, data, total, int(req.Page))
		}
	}()

	c.ShouldBind(&req)

	if data, total, err = NetworkGuideStepService.GetNetworkGuideStepList(req); err != nil {
		return
	}
}

func (ct *NetworkGuideStepController) Delete(c *gin.Context) {
	var (
		req  = new(entitys.PmNetworkGuideStepEntitys)
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
	if err = NetworkGuideStepService.DelNetworkGuideStep(key); err != nil {
		return
	}
}
