package apis

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/intelligence/entitys"
	"cloud_platform/iot_app_api_service/controls/user/services"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
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
// @Summary 查询详情
// @Description
// @Tags intelligence
// @Accept application/json
// @Param id query string true "模板Id"
// @Param homeId query string true "家庭Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /sceneTemplate/detail [get]
func (s *SceneTemplateController) Get(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	homeIdStr := c.Query("homeId")
	homeId, err := strconv.ParseInt(homeIdStr, 0, 64)
	if homeId == 0 || err != nil {
		iotgin.ResBadRequest(c, "homeId")
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
	//查询用户设备列表
	deviceRes, err := rpc.IotDeviceHomeService.HomeDevList(context.Background(), &protosService.IotDeviceHomeHomeId{
		HomeId: homeId,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if deviceRes.Code != 200 {
		iotgin.ResErrCli(c, errors.New(deviceRes.Message))
		return
	}
	//绑定设备所在房间名称
	userServies := services.AppHomeService{Ctx: controls.WithUserContext(c)}
	roomMap, _ := userServies.GetRoomList(iotutil.ToInt64(homeId))
	for i, dev := range deviceRes.DevList {
		if dev.RoomId != "" {
			if v, ok := roomMap[iotutil.ToInt64(dev.RoomId)]; ok {
				deviceRes.DevList[i].RoomName = v
			}
		}
	}
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	ret := entitys.SceneTemplate_pb2e(lang, tenantId, resp.Data[0], deviceRes.DevList)
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
// @Summary 查询列表数据
// @Description
// @Tags intelligence
// @Accept application/json
// @Param data body entitys.SceneTemplateQuery true "模板Id"
// @Success 200 {object} iotgin.ResponseModel 成功返回值
// @Router /sceneTemplate/list [post]
func (s *SceneTemplateController) List(c *gin.Context) {
	var req entitys.SceneTemplateQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}

	appKey := controls.GetAppKey(c)
	lang := controls.GetLang(c)
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		iotgin.ResBadRequest(c, "tenantId")
		return
	}
	reqSvc := &protosService.SceneTemplateListRequest{
		PageSize: int64(req.Limit),
		Page:     int64(req.Page),
	}
	reqSvc.Query = &protosService.SceneTemplate{
		AppList:  []*protosService.SceneTemplateAppRelation{{AppKey: appKey, TenantId: tenantId}},
		Status:   1,
		TenantId: tenantId,
	}
	if req.Query != nil {
		reqSvc.Query.Type = req.Query.Type
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
	var ret []*entitys.SceneTemplateApiEntity
	for _, v := range resp.Data {
		ret = append(ret, entitys.SceneTemplate_pb2e(lang, tenantId, v, nil))
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}
