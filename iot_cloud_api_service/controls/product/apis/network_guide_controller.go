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

var NetworkGuidecontroller NetworkGuideController

// NetworkGuideController ctrl
type NetworkGuideController struct{}

var NetworkGuideService services.NetworkGuideService = services.NetworkGuideService{}

func (ct *NetworkGuideController) Create(c *gin.Context) {
	var (
		req  = new(entitys.PmNetworkGuideEntitys)
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

	if data, err = NetworkGuideService.CreateNetworkGuide(req); err != nil {
		return
	}
}

func (ct *NetworkGuideController) Update(c *gin.Context) {
	var (
		req = new(entitys.PmNetworkGuideEntitys)
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

	if err = NetworkGuideService.UpdateNetworkGuide(req); err != nil {
		return
	}
}

func (ct *NetworkGuideController) Get(c *gin.Context) {
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
	if data, err = NetworkGuideService.GetNetworkGuide(key); err != nil {
		return
	}
}

func (ct *NetworkGuideController) GetNetworkGuideList(c *gin.Context) {
	var (
		data  interface{}
		total int64
		req   = new(entitys.PmNetworkGuideQuery)
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

	if data, total, err = NetworkGuideService.GetNetworkGuideList(req); err != nil {
		return
	}
}

func (ct *NetworkGuideController) Delete(c *gin.Context) {
	var (
		req  = new(entitys.PmNetworkGuideEntitys)
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
	if err = NetworkGuideService.DelNetworkGuide(key); err != nil {
		return
	}
}
