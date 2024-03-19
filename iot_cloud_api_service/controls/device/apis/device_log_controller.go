package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/device/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

var DeviceLogcontroller IotDeviceLogController

type IotDeviceLogController struct{} //设备信息操作控制器

var deviceLogServices = apiservice.IotDeviceLogService{}

func (IotDeviceLogController) QueryList(c *gin.Context) {
	var filter entitys.IotDeviceLogQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil || filter.Query.Did == "" {
		iotgin.ResBadRequest(c, "did")
		return
	}
	res, total, err := deviceLogServices.QueryIotDeviceLogList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// QueryCount 查询导出数据统计
func (IotDeviceLogController) QueryCount(c *gin.Context) {
	var filter entitys.IotDeviceLogQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		iotgin.ResBadRequest(c, "query not found")
		return
	}
	if filter.Query.StartTime == 0 || filter.Query.EndTime == 0 {
		iotgin.ResBadRequest(c, "startTime or endTime not found")
		return
	}
	// 设置只是查询统计数量
	filter.Query.IsOnlyCount = 1
	res, total, err := deviceLogServices.QueryIotDeviceLogList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (IotDeviceLogController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := deviceLogServices.GetIotDeviceLogDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (IotDeviceLogController) Edit(c *gin.Context) {
	var req entitys.IotDeviceLogEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceLogServices.UpdateIotDeviceLog(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceLogController) Add(c *gin.Context) {
	var req entitys.IotDeviceLogEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceLogServices.AddIotDeviceLog(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceLogController) Delete(c *gin.Context) {
	var req entitys.IotDeviceLogFilter
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = deviceLogServices.DeleteIotDeviceLog(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// Export 导出
func (s *IotDeviceLogController) Export(c *gin.Context) {
	filter := entitys.IotDeviceLogQueryObj{}
	filter.Did = c.Query("did")
	if filter.Did == "" {
		iotgin.ResBadRequest(c, "did")
		return
	}
	eventType := c.Query("eventType")
	if eventType != "" {
		filter.EventType, _ = iotutil.ToInt32Err(eventType)
	}
	eventKey := c.Query("eventKey")
	if eventKey != "" {
		filter.EventKey = eventKey
	}
	origin := c.Query("origin")
	if origin != "" {
		filter.Origin, _ = iotutil.ToInt32Err(origin)
	}
	endTime := c.Query("endTime")
	if endTime != "" {
		filter.EndTime, _ = iotutil.ToInt64AndErr(endTime)
	}
	startTime := c.Query("startTime")
	if startTime != "" {
		filter.StartTime, _ = iotutil.ToInt64AndErr(startTime)
	}
	isOnlyCount := c.Query("isOnlyCount")
	if isOnlyCount != "" {
		filter.IsOnlyCount, _ = iotutil.ToIntErr(isOnlyCount)
	}

	fileName, tempPathFile, err := deviceLogServices.Export(entitys.IotDeviceLogQuery{Query: &filter})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// ExportPostMethod 导出
func (s *IotDeviceLogController) ExportPostMethod(c *gin.Context) {
	var filter entitys.IotDeviceLogQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	fileName, tempPathFile, err := deviceLogServices.Export(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}
