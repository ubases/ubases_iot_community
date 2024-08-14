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

var PublicAppVersioncontroller PublicAppVersionController

// PublicAppVersionController 基础页面配置
type PublicAppVersionController struct {
}

// Get 查询详情
func (s *PublicAppVersionController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.PublicAppVersionFilter{Id: id}
	resp, err := rpc.ClientPublicAppVersionService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.PublicAppVersion_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *PublicAppVersionController) List(c *gin.Context) {
	var req entitys.PublicAppVersionQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.PublicAppVersionListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	if req.Query != nil {
		reqSvc.Query = &protosService.PublicAppVersion{}
	}
	resp, err := rpc.ClientPublicAppVersionService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.PublicAppVersionEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.PublicAppVersion_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *PublicAppVersionController) Add(c *gin.Context) {
	var req entitys.PublicAppVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.PublicAppVersion_e2pb(&req)
	resp, err := rpc.ClientPublicAppVersionService.Create(context.Background(), saveObj)
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
func (s *PublicAppVersionController) Update(c *gin.Context) {
	var req entitys.PublicAppVersionEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.PublicAppVersion_e2pb(&req)
	resp, err := rpc.ClientPublicAppVersionService.Update(context.Background(), updateObj)
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
func (s *PublicAppVersionController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.PublicAppVersion{Id: id}
	resp, err := rpc.ClientPublicAppVersionService.DeleteById(context.Background(), reqSvc)
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

// SetStatus 修改状态
func (s *PublicAppVersionController) SetStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	statusStr := c.Query("status")
	status, err := strconv.Atoi(statusStr)
	if status == 0 || err != nil {
		iotgin.ResBadRequest(c, "status")
		return
	}
	reqSvc := &protosService.PublicAppVersion{Id: id, Status: int32(status)}
	resp, err := rpc.ClientPublicAppVersionService.Update(context.Background(), reqSvc)
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
