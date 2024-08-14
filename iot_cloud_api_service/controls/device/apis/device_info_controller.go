package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/device/services"
	"cloud_platform/iot_common/iotgin"
)

var DeviceInfocontroller IotDeviceInfoController

type IotDeviceInfoController struct{} //设备信息操作控制器

var deviceInfoServices = apiservice.IotDeviceInfoService{}

func (IotDeviceInfoController) QueryDetail(c *gin.Context) {
	did := c.Param("did")
	if did == "" {
		iotgin.ResBadRequest(c, "did")
		return
	}
	lang := controls.GetLang(c)
	if lang == "" {
		lang = "zh"
	}
	res, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceDetails(lang, did)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (IotDeviceInfoController) Edit(c *gin.Context) {
	var req entitys.IotDeviceInfoEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).UpdateIotDeviceInfo(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceInfoController) Add(c *gin.Context) {
	var req entitys.IotDeviceInfoEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).AddIotDeviceInfo(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceInfoController) Delete(c *gin.Context) {
	var req entitys.IotDeviceInfoFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = deviceInfoServices.SetContext(controls.WithUserContext(c)).DeleteIotDeviceInfo(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (IotDeviceInfoController) Count(c *gin.Context) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.IotDeviceInfoQueryObj{}
	}
	filter.Query.IsQueryTriadData = false
	res, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceCount(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (IotDeviceInfoController) QueryList(c *gin.Context) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = &entitys.IotDeviceInfoQueryObj{}
	}
	filter.Query.IsQueryTriadData = false
	res, total, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 获取开发者生产管理数据
func (IotDeviceInfoController) QueryProduceList(c *gin.Context) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	filter.IsPlatform = false
	if filter.Query == nil {
		filter.Query = &entitys.IotDeviceInfoQueryObj{}
	}
	filter.Query.IsQueryTriadData = true
	res, total, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 获取云管平台生产管理数据
func (IotDeviceInfoController) QueryProducePlatformList(c *gin.Context) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	filter.IsPlatform = true
	if filter.Query == nil {
		filter.Query = &entitys.IotDeviceInfoQueryObj{}
	}
	filter.Query.IsQueryTriadData = true

	res, total, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (IotDeviceInfoController) PlatformQueryList(c *gin.Context) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	filter.IsPlatform = true
	res, total, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(filter, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// Export 导出
func (this *IotDeviceInfoController) Export(c *gin.Context, mode int) {
	var filter entitys.IotDeviceInfoQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	fileName, tempPathFile, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).Export(mode, filter) //entitys.IotDeviceInfoQuery{SearchKey: filter.SearchKey, Query: filter}
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// GetExport 导出的get方法
func (this *IotDeviceInfoController) GetExport(c *gin.Context, mode int) {
	filter := entitys.IotDeviceInfoQueryObj{}
	if isActive := c.Query("isActive"); isActive != "" {
		isActiveInt := iotutil.ToInt32(isActive)
		filter.IsActive = &isActiveInt
	}
	if isOnline := c.Query("isOnline"); isOnline != "" {
		isOnlineInt := iotutil.ToInt32(isOnline)
		filter.IsOnline = &isOnlineInt
	}
	if startTime := c.Query("startTime"); startTime != "" {
		filter.StartTime = iotutil.ToInt64(startTime)
	}
	if endTime := c.Query("endTime"); endTime != "" {
		filter.EndTime = iotutil.ToInt64(endTime)
	}
	if deviceNature := c.Query("deviceNature"); deviceNature != "" {
		filter.DeviceNature, _ = iotutil.ToInt32Err(deviceNature)
	}
	if searchKey := c.Query("searchKey"); searchKey != "" {
		filter.SearchKey = searchKey
	}
	if productId := c.Query("productId"); productId != "" {
		filter.ProductId, _ = iotutil.ToInt64AndErr(productId)
	}
	fileName, tempPathFile, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).Export(mode,
		entitys.IotDeviceInfoQuery{SearchKey: filter.SearchKey, Query: &filter}) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// GetExport 导出的get方法
func (this *IotDeviceInfoController) GetExportTriad(c *gin.Context) {
	productId, err := iotutil.ToInt64AndErr(c.DefaultQuery("productId", "0"))
	activeStatus, _ := iotutil.ToInt32Err(c.DefaultQuery("activeStatus", "-1"))
	serialNumber := c.Query("serialNumber")
	platformCode := c.Query("platformCode")
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		if platformCode != "" {
			iotgin.ResBadRequest(c, "platformCode")
			return
		}
	}
	var filter = entitys.IotDeviceInfoQueryObj{
		ProductId:    productId,
		ActiveStatus: &activeStatus,
		BatchId:      c.Query("batch"),
		SerialNumber: serialNumber,
		PlatformCode: platformCode,
	}
	if filter.ActiveStatus != nil {
		filter.IsActive = filter.ActiveStatus
	}
	filter.IsQueryTriadData = true
	filter.IsExport = true
	userId := controls.GetUserId(c)
	fileName, tempPathFile, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).ExportCsvTriad(userId,
		entitys.IotDeviceInfoQuery{SearchKey: filter.SearchKey, Query: &filter}, nil) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// GetExportTriadCount 导出数量的get方法
func (this *IotDeviceInfoController) GetExportTriadCount(c *gin.Context) {
	var inputFilter entitys.IotDeviceTriadExportQuery
	err := c.ShouldBindQuery(&inputFilter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	productId, err := iotutil.ToInt64AndErr(c.DefaultQuery("productId", "0"))
	//deviceName := c.DefaultQuery("deviceName", "")
	activeStatus, _ := iotutil.ToInt32Err(c.DefaultQuery("activeStatus", "-1"))
	serialNumber := c.Query("serialNumber")
	platformCode := c.Query("platformCode")
	tenantId := controls.GetTenantId(c)
	if tenantId == "" {
		if platformCode != "" {
			iotgin.ResBadRequest(c, "platformCode")
			return
		}
	}
	var filter = entitys.IotDeviceInfoQueryObj{
		ProductId: productId,
		//DeviceName:   deviceName,
		ActiveStatus: &activeStatus,
		BatchId:      c.Query("batch"),
		SerialNumber: serialNumber,
		PlatformCode: platformCode,
	}
	if filter.ActiveStatus != nil {
		filter.IsActive = filter.ActiveStatus
	}
	filter.IsQueryTriadData = true
	_, count, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(entitys.IotDeviceInfoQuery{
		SearchKey: filter.SearchKey, Query: &filter, IsOnlyCount: 1}, nil) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, count)
}

// GetExportTriadPlatform 云管平台导出的get方法
func (this *IotDeviceInfoController) GetExportTriadPlatform(c *gin.Context) {
	productId, err := iotutil.ToInt64AndErr(c.DefaultQuery("productId", "0"))
	activeStatus, _ := iotutil.ToInt32Err(c.DefaultQuery("activeStatus", "-1"))
	isQueryExport, _ := iotutil.ToInt32Err(c.DefaultQuery("isQueryExport", "0"))
	serialNumber := c.Query("serialNumber")
	platformCode := c.Query("platformCode")
	if platformCode == "" {
		iotgin.ResBadRequest(c, "platformCode")
		return
	}
	var filter = entitys.IotDeviceInfoQueryObj{
		ProductId:     productId,
		ActiveStatus:  &activeStatus,
		BatchId:       c.Query("batch"),
		SerialNumber:  serialNumber,
		PlatformCode:  platformCode,
		IsQueryExport: isQueryExport,
	}
	if filter.ActiveStatus != nil {
		filter.IsActive = filter.ActiveStatus
	}
	filter.IsQueryTriadData = true
	filter.IsExport = true
	userId := controls.GetUserId(c)
	fileName, tempPathFile, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).ExportCsvTriad(userId,
		entitys.IotDeviceInfoQuery{SearchKey: filter.SearchKey, Query: &filter, IsPlatform: true}, nil) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}

// GetExportTriadCountPlatform 云管平台导出数量的get方法
func (this *IotDeviceInfoController) GetExportTriadCountPlatform(c *gin.Context) {
	var inputFilter entitys.IotDeviceTriadExportQuery
	err := c.ShouldBindQuery(&inputFilter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	productId, err := iotutil.ToInt64AndErr(c.DefaultQuery("productId", "0"))
	activeStatus, _ := iotutil.ToInt32Err(c.DefaultQuery("activeStatus", "-1"))
	isQueryExport, _ := iotutil.ToInt32Err(c.DefaultQuery("isQueryExport", "0"))
	serialNumber := c.Query("serialNumber")
	platformCode := c.Query("platformCode")
	if platformCode == "" {
		iotgin.ResBadRequest(c, "platformCode")
		return
	}
	var filter = entitys.IotDeviceInfoQueryObj{
		ProductId: productId,
		//DeviceName:   deviceName,
		ActiveStatus:  &activeStatus,
		BatchId:       c.Query("batch"),
		SerialNumber:  serialNumber,
		PlatformCode:  platformCode,
		IsQueryExport: isQueryExport,
	}
	if filter.ActiveStatus != nil {
		filter.IsActive = filter.ActiveStatus
	}
	filter.IsQueryTriadData = true
	_, count, _, err := deviceInfoServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceInfoList(entitys.IotDeviceInfoQuery{
		SearchKey: filter.SearchKey, Query: &filter, IsOnlyCount: 1, IsPlatform: true}, nil) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, count)
}
