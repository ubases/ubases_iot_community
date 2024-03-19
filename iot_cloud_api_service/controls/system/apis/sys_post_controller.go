package apis

import (
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
)

var Postcontroller SysPostController

type SysPostController struct{} //部门操作控制器

var postServices = apiservice.SysPostService{}

func (SysPostController) QueryList(c *gin.Context) {
	var filter entitys.SysPostQuery
	err := c.BindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := postServices.QuerySysPostList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (SysPostController) QueryDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := postServices.GetSysPostDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (SysPostController) Edit(c *gin.Context) {
	var req entitys.SysPostEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysPost_SetCommonFiled(&req, iotutil.ToInt64(userid), 2)
	id, err := postServices.UpdateSysPost(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (SysPostController) Add(c *gin.Context) {
	var req entitys.SysPostEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysPost_SetCommonFiled(&req, iotutil.ToInt64(userid), 1)
	id, err := postServices.AddSysPost(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (SysPostController) Delete(c *gin.Context) {
	var req entitys.DeleteCommonQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(req.Ids) == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	err = postServices.DeleteSysPost(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
