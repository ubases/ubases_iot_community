package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gin-gonic/gin"
)

var Overviewcontroller OverviewController

type OverviewController struct{} //部门操作控制器

func (OverviewController) getAccumulateData(c *gin.Context) {
	start := GetLast12Month()
	end := iotutil.New(time.Now()).BeginningOfMonth()
	req := protosService.DataOverviewMonthListRequest{
		Query: &protosService.DataOverviewMonthListFilter{
			TenantId:  "",
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
	var resp entitys.OverviewEntitys
	for _, v := range rsp.GetData() {
		//总计
		if v.DataTime.AsTime().IsZero() {
			resp.ActiveDevice.Total = v.DeviceActiveSum
			resp.DeviceFault.Total = v.DeviceFaultSum
			resp.AppUser.Total = v.UserRegisterSum
			resp.Developer.Total = v.DeveloperRegisterSum
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
			//开发者注册
			tData = entitys.TimeData{Time: mon, Total: v.DeveloperRegisterSum}
			resp.Developer.Data = append(resp.Developer.Data, tData)
		}
	}
	resp.ActiveDevice.Data = entitys.FillTimeData(resp.ActiveDevice.Data, 1, start, end)
	resp.DeviceFault.Data = entitys.FillTimeData(resp.DeviceFault.Data, 1, start, end)
	resp.AppUser.Data = entitys.FillTimeData(resp.AppUser.Data, 1, start, end)
	resp.Developer.Data = entitys.FillTimeData(resp.Developer.Data, 1, start, end)
	iotgin.ResSuccess(c, resp)
}

func (OverviewController) getTodayData(c *gin.Context) {
	now := time.Now()
	start := iotutil.New(now).BeginningOfDay()
	end := iotutil.New(now).BeginningOfHour()
	req := protosService.DataOverviewHourListRequest{
		Query: &protosService.DataOverviewHourListFilter{
			TenantId:  "",
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
	var resp entitys.OverviewEntitys
	for _, v := range rsp.GetData() {
		//总计
		if v.DataTime.AsTime().IsZero() {
			resp.ActiveDevice.Total = v.DeviceActiveSum
			resp.DeviceFault.Total = v.DeviceFaultSum
			resp.AppUser.Total = v.UserRegisterSum
			resp.Developer.Total = v.DeveloperRegisterSum
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
			//开发者注册
			tData = entitys.TimeData{Time: hour, Total: v.DeveloperRegisterSum}
			resp.Developer.Data = append(resp.Developer.Data, tData)
		}
	}
	resp.ActiveDevice.Data = entitys.FillTimeData(resp.ActiveDevice.Data, 3, start, end)
	resp.DeviceFault.Data = entitys.FillTimeData(resp.DeviceFault.Data, 3, start, end)
	resp.AppUser.Data = entitys.FillTimeData(resp.AppUser.Data, 3, start, end)
	resp.Developer.Data = entitys.FillTimeData(resp.Developer.Data, 3, start, end)
	iotgin.ResSuccess(c, resp)
}

func (OverviewController) getCityDeviceData(c *gin.Context) {
	//todo 当前缺城市代码、经纬度配置，缺集中存储的在线设备数据，暂时不实现
	str := `{"country":"CN","data":[{"cityName":"长沙市","cityCode":"101050211","longitude":112.982279,
"latitude":28.19409,"activeDevice":1000000,"onlineDevice":1000000}]}`
	var resp entitys.DeviceCityEntitys
	_ = json.Unmarshal([]byte(str), &resp)
	iotgin.ResSuccess(c, resp)
}

func GetLast12Month() time.Time {
	t := time.Now()
	if t.Month() == 12 {
		return iotutil.New(t).BeginningOfYear()
	}
	t1 := t.AddDate(-1, 1, 0)
	return iotutil.New(t1).BeginningOfMonth()
}

//func GetLast12Month() time.Time {
//	t := time.Now()
//	if t.Month() == 12 {
//		return iotutil.New(t).BeginningOfYear()
//	}
//	t1 := t.AddDate(-1, 1, 0)
//	return iotutil.New(t1).BeginningOfMonth()
//}
