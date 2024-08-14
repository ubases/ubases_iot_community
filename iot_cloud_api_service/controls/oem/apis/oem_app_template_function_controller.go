package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var OemAppTemplateFunctioncontroller OemAppTemplateFunctionController

// OemAppTemplateFunctionController 基础页面配置
type OemAppTemplateFunctionController struct {
}

// Get 查询详情
func (s *OemAppTemplateFunctionController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateFunctionFilter{Id: id}
	resp, err := rpc.ClientOemAppTemplateFunctionService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppTemplateFunction_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppTemplateFunctionController) List(c *gin.Context) {
	var req entitys.OemAppTemplateFunctionQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppTemplateFunctionListRequest{}
	if req.Query != nil {
		reqSvc.Query = &protosService.OemAppTemplateFunction{}
	}
	resp, err := rpc.ClientOemAppTemplateFunctionService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.OemAppTemplateFunctionEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.OemAppTemplateFunction_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppTemplateFunctionController) Add(c *gin.Context) {
	var req entitys.OemAppTemplateFunctionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.OemAppTemplateFunction_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateFunctionService.Create(context.Background(), saveObj)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccess(c, resp.Data)
}

// Update 修改
func (s *OemAppTemplateFunctionController) Update(c *gin.Context) {
	var req entitys.OemAppTemplateFunctionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppTemplateFunction_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateFunctionService.Update(context.Background(), updateObj)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccess(c, resp.Data)
}

// Delete 删除
func (s *OemAppTemplateFunctionController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateFunction{Id: id}
	resp, err := rpc.ClientOemAppTemplateFunctionService.DeleteById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccessMsg(c)
}
