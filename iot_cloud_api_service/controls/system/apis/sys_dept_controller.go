package apis

import (
	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
)

var Deptcontroller DeptController

type DeptController struct{} //部门操作控制器

var deptServices = apiservice.SysDeptService{}

func (DeptController) QueryList(c *gin.Context) {
	var filter entitys.SysDeptQuery
	err := c.BindQuery(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := deptServices.QuerySysDeptList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, 1)
}

func (DeptController) QueryDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := deptServices.GetSysDeptDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (DeptController) Edit(c *gin.Context) {
	var req entitys.SysDeptEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysDept_SetCommonFiled(&req, iotutil.ToInt64(userid), 2)

	id, err := deptServices.UpdateSysDept(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (DeptController) Add(c *gin.Context) {
	var req entitys.SysDeptEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysDept_SetCommonFiled(&req, iotutil.ToInt64(userid), 1)

	id, err := deptServices.AddSysDept(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (DeptController) Delete(c *gin.Context) {
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
	err = deptServices.DeleteSysDept(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
