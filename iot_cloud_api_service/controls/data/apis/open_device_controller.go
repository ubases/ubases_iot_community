package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"

	"github.com/gin-gonic/gin"
)

var Opendevicecontroller OpenDeviceController

type OpenDeviceController struct {
}

func (OpenDeviceController) getActiveStatistics(c *gin.Context) {
	productKey := c.Query("productKey")
	//c.Set("tenantId", "ioqp4r") //for debug
	req := protosService.DeviceActiveListFilter{
		ProductKey: productKey,
		TenantId:   c.GetString("tenantId"),
	}
	resp, err := rpc.ClientDeviceActiveService.Lists(c, &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Data == nil || resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OpenActiveEntitys{
		DeviceTodayActive: resp.Data.DeviceTodayActive,
		Device7DayActive:  resp.Data.Device7DayActive,
		DeviceActiveAll:   resp.Data.DeviceActiveAll,
		DeviceMonActive:   entitys.Data{Total: resp.Data.DeviceMonActive.Total},
		DeviceDayActive:   entitys.Data{Total: resp.Data.DeviceDayActive.Total},
	}

	for _, v := range resp.Data.DeviceMonActive.GetData() {
		ret.DeviceMonActive.Data = append(ret.DeviceMonActive.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}

	for _, v := range resp.Data.DeviceDayActive.GetData() {
		ret.DeviceDayActive.Data = append(ret.DeviceDayActive.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}

	iotgin.ResSuccess(c, ret)
}

func (OpenDeviceController) getFaultStatistics(c *gin.Context) {
	productKey := c.Query("productKey")
	data := entitys.OpenFaultEntitys{}
	// 聚合月故障数据
	req := &protosService.ProductFaultMonthListRequest{
		Query: &protosService.ProductFaultMonth{
			ProductKey: productKey,
		},
	}
	resp, err := rpc.ClientProductFaultMonthService.Lists(context.Background(), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	for i := range resp.Data {
		timeData := entitys.TimeData{
			Time:  resp.Data[i].Month.AsTime().Local().Format("2006-01"),
			Total: resp.Data[i].Total,
		}
		data.DeviceMonFault.Data = append(data.DeviceMonFault.Data, timeData)
	}
	data.DeviceMonFault.Total = resp.Total
	// 聚合月故障类型数据
	reqType := &protosService.ProductFaultTypeListRequest{
		Query: &protosService.ProductFaultType{
			ProductKey: productKey,
		},
	}
	respType, err := rpc.ClientProductFaultTypeService.Lists(context.Background(), reqType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if respType.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	if respType.Data == nil {
		data.DeviceFaultType.Data = []entitys.FaultTypeData{}
	}
	for i := range respType.Data {
		timeData := entitys.FaultTypeData{
			FaultType: respType.Data[i].FaultType,
			Total:     respType.Data[i].Total,
		}
		data.DeviceFaultType.Data = append(data.DeviceFaultType.Data, timeData)
		data.DeviceFaultType.Total += respType.Data[i].Total
	}
	iotgin.ResSuccess(c, data)
}
