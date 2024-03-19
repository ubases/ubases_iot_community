package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gin-gonic/gin"
)

var Openoverviewcontroller OpenOverviewController

type OpenOverviewController struct{} //部门操作控制器

func (OpenOverviewController) getAccumulateData(c *gin.Context) {
	tenantId := c.GetString("tenantId")
	if tenantId == "" {
		iotgin.ResErrCli(c, errors.New("缺tenantId"))
		return
	}
	start := GetLast12Month()
	end := iotutil.New(time.Now()).BeginningOfMonth()
	req := protosService.DataOverviewMonthListRequest{
		Query: &protosService.DataOverviewMonthListFilter{
			TenantId:  tenantId,
			StartTime: timestamppb.New(start),
			EndTime:   timestamppb.New(end),
		},
	}
	rsp, err := rpc.ClientDataOverviewMonthService.Lists(context.Background(), &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rsp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rsp.Message))
		return
	}
	var resp entitys.OpenOverviewEntitys
	for _, v := range rsp.GetData() {
		//总计
		if v.DataTime.AsTime().IsZero() {
			resp.ActiveDevice.Total = v.DeviceActiveSum
			resp.DeviceFault.Total = v.DeviceFaultSum
			resp.AppUser.Total = v.UserRegisterSum
		} else {
			//分月
			mon := v.DataTime.AsTime().Local().Format("2006-01")
			//设备激活
			tData := entitys.TimeData{Time: mon, Total: v.DeviceActiveSum}
			resp.ActiveDevice.Data = append(resp.ActiveDevice.Data, tData)
			//设备故障
			tData = entitys.TimeData{Time: mon, Total: v.DeviceFaultSum}
			resp.DeviceFault.Data = append(resp.DeviceFault.Data, tData)
			//用户注册
			tData = entitys.TimeData{Time: mon, Total: v.UserRegisterSum}
			resp.AppUser.Data = append(resp.AppUser.Data, tData)
		}
	}
	resp.ActiveDevice.Data = entitys.FillTimeData(resp.ActiveDevice.Data, 1, start, end)
	resp.DeviceFault.Data = entitys.FillTimeData(resp.DeviceFault.Data, 1, start, end)
	resp.AppUser.Data = entitys.FillTimeData(resp.AppUser.Data, 1, start, end)
	iotgin.ResSuccess(c, resp)
}

func (OpenOverviewController) getTodayData(c *gin.Context) {
	tenantId := c.GetString("tenantId")
	if tenantId == "" {
		iotgin.ResErrCli(c, errors.New("缺tenantId"))
		return
	}
	now := time.Now()
	start := iotutil.New(now).BeginningOfDay()
	end := iotutil.New(now).BeginningOfHour()
	req := protosService.DataOverviewHourListRequest{
		Query: &protosService.DataOverviewHourListFilter{
			TenantId:  tenantId,
			StartTime: timestamppb.New(start),
			EndTime:   timestamppb.New(end),
		},
	}
	rsp, err := rpc.ClientDataOverviewHourService.Lists(context.Background(), &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if rsp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(rsp.Message))
		return
	}
	var resp entitys.OpenOverviewEntitys
	for _, v := range rsp.GetData() {
		//总计
		if v.DataTime.AsTime().IsZero() {
			resp.ActiveDevice.Total = v.DeviceActiveSum
			resp.DeviceFault.Total = v.DeviceFaultSum
			resp.AppUser.Total = v.UserRegisterSum
		} else {
			//小时
			hour := v.DataTime.AsTime().Local().Format("15:04")
			//设备激活
			tData := entitys.TimeData{Time: hour, Total: v.DeviceActiveSum}
			resp.ActiveDevice.Data = append(resp.ActiveDevice.Data, tData)
			//设备故障
			tData = entitys.TimeData{Time: hour, Total: v.DeviceFaultSum}
			resp.DeviceFault.Data = append(resp.DeviceFault.Data, tData)
			//用户注册
			tData = entitys.TimeData{Time: hour, Total: v.UserRegisterSum}
			resp.AppUser.Data = append(resp.AppUser.Data, tData)
		}
	}
	resp.ActiveDevice.Data = entitys.FillTimeData(resp.ActiveDevice.Data, 3, start, end)
	resp.DeviceFault.Data = entitys.FillTimeData(resp.DeviceFault.Data, 3, start, end)
	resp.AppUser.Data = entitys.FillTimeData(resp.AppUser.Data, 3, start, end)
	iotgin.ResSuccess(c, resp)
}
