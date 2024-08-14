package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/common"
	"cloud_platform/iot_cloud_api_service/controls/data/entitys"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/gin-gonic/gin"
)

var Openappusercontroller OpenAppUserController

type OpenAppUserController struct {
}

func (OpenAppUserController) getUserAppStatistics(c *gin.Context) {
	appKey := c.Query("appKey")
	req := protosService.AppUserStatisticsFilter{AppKey: appKey}
	resp, err := rpc.ClientAppUserStatisticsService.GetAppUserStatistics(c, &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Data == nil || resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	ret := entitys.OpenAppUserEntitys{
		AppUserTodayActive: resp.Data.AppUserTodayActive,
		AppUserToday:       resp.Data.AppUserToday,
		AppUserAll:         resp.Data.AppUserAll,
		AppUser:            entitys.Data{Total: resp.Data.AppUser.Total},
		ActiveUser:         entitys.Data{Total: resp.Data.ActiveUser.Total},
	}
	for _, v := range resp.Data.AppUser.GetData() {
		ret.AppUser.Data = append(ret.AppUser.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}
	for _, v := range resp.Data.ActiveUser.GetData() {
		ret.ActiveUser.Data = append(ret.ActiveUser.Data, entitys.TimeData{Time: v.Time, Total: v.Total})
	}
	iotgin.ResSuccess(c, ret)
}

func (OpenAppUserController) ExportUserAppStatistics(c *gin.Context) {
	appKey := c.Query("appKey")
	dataType := c.Query("dataType")
	req := protosService.AppUserStatisticsFilter{AppKey: appKey}
	resp, err := rpc.ClientAppUserStatisticsService.GetAppUserStatistics(c, &req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resp.Data == nil || resp.Code != 200 {
		iotgin.ResErrCli(c, errors.New(resp.Message))
		return
	}
	//active:活跃用户；register:注册用户
	var tempfile string
	var fileName string
	if dataType == "active" {
		var dataList []entitys.TimeData
		for _, v := range resp.Data.ActiveUser.GetData() {
			dataList = append(dataList, entitys.TimeData{Time: v.Time, Total: v.Total})
		}
		tempfile, err = GenExportFile([]string{"月分", "活跃用户数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, errors.New(resp.Message))
			return
		}
		fileName = fmt.Sprintf("APP(%s)最近30天活跃用户数", appKey) + time.Now().Format("20060102150400") + ".xlsx"
	} else {
		var dataList []entitys.TimeData
		for _, v := range resp.Data.AppUser.GetData() {
			dataList = append(dataList, entitys.TimeData{Time: v.Time, Total: v.Total})
		}
		tempfile, err = GenExportFile([]string{"月分", "注册用户数"}, dataList)
		if err != nil {
			iotgin.ResErrCli(c, errors.New(resp.Message))
			return
		}
		fileName = fmt.Sprintf("APP(%s)最近12月注册用户数", appKey) + time.Now().Format("20060102150400") + ".xlsx"
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName)))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempfile)
}

func GenExportFile(headerList []string, data []entitys.TimeData) (string, error) {
	headerStyle := common.ExcelHeaderStyle()
	contentStyle := common.ExcelContentStyle()
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	headerRow := sheet.AddRow()
	for _, col := range headerList {
		cell := headerRow.AddCell()
		cell.SetStyle(headerStyle)
		cell.Value = col
	}
	for _, row := range data {
		headerRow := sheet.AddRow()
		cell := headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = row.Time
		cell = headerRow.AddCell()
		cell.SetStyle(contentStyle)
		cell.Value = iotutil.ToString(row.Total)
	}
	var tempPathFile = iotconst.GetWorkTempDir() + string(filepath.Separator)
	os.MkdirAll(tempPathFile, os.ModePerm)
	tempPathFile = tempPathFile + iotutil.Uuid() + ".xlsx"
	err := file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		return "", err
	}
	return tempPathFile, nil
}
