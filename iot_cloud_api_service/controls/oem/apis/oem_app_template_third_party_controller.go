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

var OemAppTemplateThirdPartycontroller OemAppTemplateThirdPartyController

// OemAppTemplateThirdPartyController 基础页面配置
type OemAppTemplateThirdPartyController struct {
}

// Get 查询详情
func (s *OemAppTemplateThirdPartyController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateThirdPartyFilter{Id: id}
	resp, err := rpc.ClientOemAppTemplateThirdPartyService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppTemplateThirdParty_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppTemplateThirdPartyController) List(c *gin.Context) {
	var req entitys.OemAppTemplateThirdPartyQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppTemplateThirdPartyListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	reqSvc.Query = &protosService.OemAppTemplateThirdParty{
		AppTemplateId: req.AppTemplateId,
	}
	resp, err := rpc.ClientOemAppTemplateThirdPartyService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.OemAppTemplateThirdPartyEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.OemAppTemplateThirdParty_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppTemplateThirdPartyController) Add(c *gin.Context) {
	var req entitys.OemAppTemplateThirdPartyEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.OemAppTemplateThirdParty_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateThirdPartyService.Create(context.Background(), saveObj)
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
func (s *OemAppTemplateThirdPartyController) Update(c *gin.Context) {
	var req entitys.OemAppTemplateThirdPartyEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppTemplateThirdParty_e2pb(&req)
	resp, err := rpc.ClientOemAppTemplateThirdPartyService.Update(context.Background(), updateObj)
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
func (s *OemAppTemplateThirdPartyController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppTemplateThirdParty{Id: id}
	resp, err := rpc.ClientOemAppTemplateThirdPartyService.DeleteById(context.Background(), reqSvc)
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
