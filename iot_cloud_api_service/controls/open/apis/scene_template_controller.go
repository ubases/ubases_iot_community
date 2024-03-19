package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var SceneTemplatecontroller SceneTemplateController

// SceneTemplateController 基础页面配置
type SceneTemplateController struct {
}

// Get 查询详情
func (s *SceneTemplateController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.SceneTemplateFilter{Id: id}
	resp, err := rpc.ClientSceneTemplateService.FindById(controls.WithUserContext(c), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.SceneTemplate_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *SceneTemplateController) List(c *gin.Context) {
	var req entitys.SceneTemplateQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.SceneTemplateListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	tenantId := controls.GetTenantId(c)
	reqSvc.Query = &protosService.SceneTemplate{TenantId: tenantId}
	if req.Query != nil {
		reqSvc.Query.Type = req.Query.Type
		reqSvc.Query.Title = req.Query.Title
		reqSvc.Query.Status = req.Query.Status
		reqSvc.Query.ConditionMode = req.Query.ConditionMode
	}
	resp, err := rpc.ClientSceneTemplateService.Lists(controls.WithUserContext(c), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	appMap := s.getAppMaps(tenantId)

	var ret []*entitys.SceneTemplateApiEntity
	for _, v := range resp.Data {
		r := entitys.SceneTemplate_pb2e(v)
		for i, app := range r.AppList {
			if v, ok := appMap[app.AppKey]; ok {
				r.AppList[i].AppName = v
			}
		}
		ret = append(ret, r)
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

func (s SceneTemplateController) getAppMaps(tenantId string) map[string]string {
	//查询开发者的APP信息
	apps, err := rpc.ClientOemAppService.Lists(context.Background(), &protosService.OemAppListRequest{
		Query: &protosService.OemApp{
			TenantId: tenantId,
		},
	})
	if err != nil {
		return nil
	}
	if apps.Code != 200 {
		return nil
	}
	var appMap = make(map[string]string)
	for _, app := range apps.Data {
		appMap[app.AppKey] = app.Name
	}
	return appMap
}

// Add 新增
func (s *SceneTemplateController) Add(c *gin.Context) {
	var req entitys.SceneTemplateApiEntity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	saveObj, err := entitys.SceneTemplate_e2pb(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	resp, err := rpc.ClientSceneTemplateService.Create(controls.WithUserContext(c), saveObj)
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
func (s *SceneTemplateController) Update(c *gin.Context) {
	var req entitys.SceneTemplateApiEntity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj, err := entitys.SceneTemplate_e2pb(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	resp, err := rpc.ClientSceneTemplateService.Update(controls.WithUserContext(c), updateObj)
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
func (s *SceneTemplateController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.SceneTemplate{Id: id}
	resp, err := rpc.ClientSceneTemplateService.DeleteById(controls.WithUserContext(c), reqSvc)
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
func (s *SceneTemplateController) SetStatus(c *gin.Context) {
	var req entitys.SceneTemplateFilter
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
	reqSvc := &protosService.SceneTemplateUpdateFieldsRequest{
		Fields: []string{"status"},
		Data:   &protosService.SceneTemplate{Id: req.Id, Status: req.Status}}
	resp, err := rpc.ClientSceneTemplateService.UpdateFields(controls.WithUserContext(c), reqSvc)
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
