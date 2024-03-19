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

var OemAppAssistReleasecontroller OemAppAssistReleaseController

// OemAppAssistReleaseController 基础页面配置
type OemAppAssistReleaseController struct {
}

// Get 查询详情
func (s *OemAppAssistReleaseController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppAssistReleaseFilter{Id: id}
	resp, err := rpc.ClientOemAppAssistReleaseService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppAssistRelease_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppAssistReleaseController) List(c *gin.Context) {
	var req entitys.OemAppAssistReleaseQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppAssistReleaseListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	if req.Query != nil {
		reqSvc.Query = &protosService.OemAppAssistRelease{
			DevelopId:          req.Query.DevelopId,
			TenantId:           req.Query.TenantId,
			DevelopPhone:       req.Query.DevelopPhone,
			AppKey:             req.Query.AppKey,
			AppName:            req.Query.AppName,
			AppTemplateId:      req.Query.AppTemplateId,
			AppTemplateVersion: req.Query.AppTemplateVersion,
			AppVersion:         req.Query.AppVersion,
			SkinId:             req.Query.SkinId,
		}
	}
	resp, err := rpc.ClientOemAppAssistReleaseService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	var ret []*entitys.OemAppAssistReleaseEntitys
	for _, v := range resp.Data {
		ret = append(ret, entitys.OemAppAssistRelease_pb2e(v))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppAssistReleaseController) Add(c *gin.Context) {
	var req entitys.OemAppAssistReleaseEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj := entitys.OemAppAssistRelease_e2pb(&req)
	resp, err := rpc.ClientOemAppAssistReleaseService.Create(context.Background(), saveObj)
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
func (s *OemAppAssistReleaseController) Update(c *gin.Context) {
	var req entitys.OemAppAssistReleaseEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppAssistRelease_e2pb(&req)
	resp, err := rpc.ClientOemAppAssistReleaseService.Update(context.Background(), updateObj)
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
func (s *OemAppAssistReleaseController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppAssistRelease{Id: id}
	resp, err := rpc.ClientOemAppAssistReleaseService.DeleteById(context.Background(), reqSvc)
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
func (s *OemAppAssistReleaseController) SetStatus(c *gin.Context) {
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
	reqSvc := &protosService.OemAppAssistRelease{Id: id, Status: int32(status)}
	resp, err := rpc.ClientOemAppAssistReleaseService.Update(context.Background(), reqSvc)
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

// DeveloperAppList 查询开发者的APP信息
func (s *OemAppAssistReleaseController) DeveloperAppList(c *gin.Context) {
	account := c.Query("account")
	if account == "" {
		iotgin.ResBadRequest(c, "account")
		return
	}
	//查询APP
	developers, err := rpc.ClientDeveloperService.BasicList(context.Background(), &protosService.DeveloperListRequest{
		Query: &protosService.DeveloperListSearchInfo{
			Account: account,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if developers.Code != 200 {
		iotgin.ResErrCli(c, errors.New(developers.Message))
		return
	}
	if len(developers.Data) == 0 {
		iotgin.ResBusinessP(c, "未找到开发者信息")
		return
	}
	developer := developers.Data[0]
	reqSvc := &protosService.DeveloperAppListRequest{TenantId: developer.TenantId}
	resp, err := rpc.ClientOemAppAssistReleaseService.DeveloperAppList(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.DeveloperApp_pb2e(developer, resp.AppList)
	iotgin.ResSuccess(c, ret)
}
