package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var OemAppTemplatecontroller OemAppTemplateController

// OemAppTemplateController 基础页面配置
type OemAppTemplateController struct {
}

// Get 查询详情
func (s *OemAppTemplateController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateFilter{Id: id}
	resp, err := rpc.ClientOemAppTemplateService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppTemplate_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppTemplateController) List(c *gin.Context) {
	var req entitys.OemAppTemplateQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppTemplateListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	if req.Query != nil {
		reqSvc.Query = &protosService.OemAppTemplate{
			Type:          req.Query.Type,
			Version:       req.Query.Version,
			OpenRangeType: req.Query.OpenRangeType,
			Status:        req.Query.Status,
		}
	}
	resp, err := rpc.ClientOemAppTemplateService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.OemAppTemplateEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.OemAppTemplate_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppTemplateController) Add(c *gin.Context) {
	var req entitys.OemAppTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.OemAppTemplate_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateService.Create(context.Background(), saveObj)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccess(c, iotutil.ToString(resp.Data))
}

// Update 修改
func (s *OemAppTemplateController) Update(c *gin.Context) {
	var req entitys.OemAppTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppTemplate_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateService.Update(context.Background(), updateObj)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	iotgin.ResSuccess(c, iotutil.ToString(resp.Data))
}

// Delete 删除
func (s *OemAppTemplateController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplate{Id: id}
	resp, err := rpc.ClientOemAppTemplateService.DeleteById(context.Background(), reqSvc)
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
func (s *OemAppTemplateController) SetStatus(c *gin.Context) {
	var req entitys.OemAppTemplateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.Id == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	if req.Status == 0 {
		iotgin.ResBadRequest(c, "status")
		return
	}
	reqSvc := &protosService.OemAppTemplate{Id: req.Id, Status: req.Status}
	resp, err := rpc.ClientOemAppTemplateService.Update(context.Background(), reqSvc)
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

// Copy 复制
func (s *OemAppTemplateController) Copy(c *gin.Context) {
	var req entitys.OemAppTemplateCopy
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.Id == 0 {
		iotgin.ResBadRequest(c, "id")
	}
	resp, err := rpc.ClientOemAppTemplateService.AppTemplateCopy(context.Background(), &protosService.OemAppTemplate{
		Id:      req.Id,
		Version: req.Version,
	})
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
