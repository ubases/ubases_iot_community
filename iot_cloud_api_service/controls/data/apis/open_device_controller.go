package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

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

func (OpenDeviceController) ExportActiveStatistics(c *gin.Context) {
	productKey := c.Query("productKey")
	dataType := c.Query("dataType")
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
	//last30days:近30日每日激活；last12months:最近12月每月激活
	var tempfile string
	var fileName string
	if dataType == "last30days" {
		var dataList []entitys.TimeData
		for _, v := range resp.Data.DeviceDayActive.GetData() {
			dataList = append(dataList, entitys.TimeData{Time: v.Time, Total: v.Total})
		}
		tempfile, err = GenExportFile([]string{"日期", "激活数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, errors.New(resp.Message))
			return
		}
		fileName = fmt.Sprintf("产品(%s)最近30天设备激活数", productKey) + time.Now().Format("20060102150400") + ".xlsx"
	} else {
		var dataList []entitys.TimeData
		for _, v := range resp.Data.DeviceMonActive.GetData() {
			dataList = append(dataList, entitys.TimeData{Time: v.Time, Total: v.Total})
		}
		tempfile, err = GenExportFile([]string{"月分", "激活数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, errors.New(resp.Message))
			return
		}
		fileName = fmt.Sprintf("产品(%s)最近12月设备激活数", productKey) + time.Now().Format("20060102150400") + ".xlsx"
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempfile)
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

func (OpenDeviceController) ExportFaultStatistics(c *gin.Context) {
	productKey := c.Query("productKey")
	dataType := c.Query("dataType")
	//num_last12months:最近12月每月故障数；type_all:故障类型累计故障数
	var tempfile string
	var fileName string
	if dataType == "num_last12months" {
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
		var dataList []entitys.TimeData
		for i := range resp.Data {
			dataList = append(dataList, entitys.TimeData{
				Time:  resp.Data[i].Month.AsTime().Local().Format("2006-01"),
				Total: resp.Data[i].Total,
			})
		}
		tempfile, err = GenExportFile([]string{"月分", "故障数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, errors.New(resp.Message))
			return
		}
		fileName = fmt.Sprintf("产品(%s)最近12月故障数", productKey) + time.Now().Format("20060102150400") + ".xlsx"
	} else if dataType == "type_all" {
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
			iotgin.ResErrCli(c, errors.New(respType.Message))
			return
		}
		var dataList []entitys.TimeData
		for i := range respType.Data {
			dataList = append(dataList, entitys.TimeData{
				Time:  respType.Data[i].FaultType,
				Total: respType.Data[i].Total,
			})
		}
		tempfile, err = GenExportFile([]string{"故障类型", "故障数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		fileName = fmt.Sprintf("产品(%s)故障类型统计", productKey) + time.Now().Format("20060102150400") + ".xlsx"
	}
	if fileName != "" {
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		//发送文件
		c.File(tempfile)
	}
}
