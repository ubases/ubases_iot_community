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

var OemAppTemplateSkincontroller OemAppTemplateSkinController

// OemAppTemplateSkinController 基础页面配置
type OemAppTemplateSkinController struct {
}

// Get 查询详情
func (s *OemAppTemplateSkinController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateSkinFilter{Id: id}
	resp, err := rpc.ClientOemAppTemplateSkinService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppTemplateSkin_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppTemplateSkinController) List(c *gin.Context) {
	var req entitys.OemAppTemplateSkinQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppTemplateSkinListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	reqSvc.Query = &protosService.OemAppTemplateSkin{
		AppTemplateId: req.AppTemplateId,
	}
	resp, err := rpc.ClientOemAppTemplateSkinService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.OemAppTemplateSkinEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.OemAppTemplateSkin_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppTemplateSkinController) Add(c *gin.Context) {
	var req entitys.OemAppTemplateSkinEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.OemAppTemplateSkin_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateSkinService.Create(context.Background(), saveObj)
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
func (s *OemAppTemplateSkinController) Update(c *gin.Context) {
	var req entitys.OemAppTemplateSkinEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppTemplateSkin_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateSkinService.Update(context.Background(), updateObj)
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
func (s *OemAppTemplateSkinController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateSkin{Id: id}
	resp, err := rpc.ClientOemAppTemplateSkinService.DeleteById(context.Background(), reqSvc)
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
