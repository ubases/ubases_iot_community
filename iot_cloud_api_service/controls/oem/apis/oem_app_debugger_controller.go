package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/config/services"
	"cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	services2 "cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var OemAppDebuggercontroller OemAppDebuggerController

// OemAppDebuggerController 基础页面配置
type OemAppDebuggerController struct {
}

// Get 查询详情
func (s *OemAppDebuggerController) Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppDebuggerFilter{Id: id}
	resp, err := rpc.ClientOemAppDebuggerService.FindById(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OemAppDebugger_pb2e(resp.Data[0])
	iotgin.ResSuccess(c, ret)
}

// List 查询列表数据
func (s *OemAppDebuggerController) List(c *gin.Context) {
	var req entitys.OemAppDebuggerQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	reqSvc := &protosService.OemAppDebuggerListRequest{
		PageSize:  int64(req.Limit),
		Page:      int64(req.Page),
		OrderKey:  req.SortField,
		OrderDesc: req.Sort,
	}
	if req.Query != nil {
		reqSvc.Query = &protosService.OemAppDebugger{
			UserName: req.Query.UserName,
			RegionId: req.Query.RegionId,
			AppKey:   req.Query.AppKey,
			TenantId: controls.GetTenantId(c),
		}
	}
	resp, err := rpc.ClientOemAppDebuggerService.Lists(context.Background(), reqSvc)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	appSvc := services2.OemAppService{Ctx: context.Background()}
	appMaps, _ := appSvc.GetAppKeyMap(controls.GetTenantId(c))

	regionSvc := services.SysAreaService{}
	regionMap, _ := regionSvc.GetRegionMap()
	var ret []*entitys.OemAppDebuggerEntitys
	for _, v := range resp.Data {
		regionName := ""
		if v, ok := regionMap[v.RegionId]; ok {
			regionName = v.Describe
		}
		d := entitys.OemAppDebugger_pb2e(v)
		d.RegionName = regionName
		if a, ok := appMaps[v.AppKey]; ok {
			d.AppImg = a.AppIconUrl
			d.AppName = a.Name
		}
		ret = append(ret, d)
	}
	iotgin.ResPageSuccess(c, ret, resp.Total, int(req.Page))
}

// Add 新增
func (s *OemAppDebuggerController) Add(c *gin.Context) {
	var req entitys.OemAppDebuggerEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if req.AppKey == "" {
		iotgin.ResBadRequest(c, "appKey")
		return
	}
	if req.RegionId == 0 {
		iotgin.ResBadRequest(c, "regionId")
		return
	}
	if req.UserName == "" {
		iotgin.ResBadRequest(c, "userName")
		return
	}
	saveObj := entitys.OemAppDebugger_e2pb(&req)
	saveObj.Id = iotutil.GetNextSeqInt64()
	userReq, err := rpc.UcUserService.Find(context.Background(), &protosService.UcUserFilter{
		UserName:       req.UserName,
		RegionServerId: req.RegionId,
		AppKey:         req.AppKey,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(userReq.Data) == 0 {
		iotgin.ResBadRequest(c, "未找到用户")
		return
	}
	userId := userReq.Data[0].Id
	saveObj.UserId = userId
	saveObj.Status = 1
	resp, err := rpc.ClientOemAppDebuggerService.Create(context.Background(), saveObj)
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
func (s *OemAppDebuggerController) Update(c *gin.Context) {
	var req entitys.OemAppDebuggerEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	updateObj := entitys.OemAppDebugger_e2pb(&req)
	resp, err := rpc.ClientOemAppDebuggerService.Update(context.Background(), updateObj)
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
func (s *OemAppDebuggerController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if id == 0 || err != nil {
		iotgin.ResBadRequest(c, "id")
		return
	}
	reqSvc := &protosService.OemAppDebugger{Id: id}
	resp, err := rpc.ClientOemAppDebuggerService.DeleteById(context.Background(), reqSvc)
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
