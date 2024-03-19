package apis

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/device/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/device/services"
	entitys2 "cloud_platform/iot_cloud_api_service/controls/oem/entitys"
	"cloud_platform/iot_cloud_api_service/controls/oem/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

var DeviceTriadcontroller IotDeviceTriadController

type IotDeviceTriadController struct{} //部门操作控制器

var deviceTriadServices = apiservice.IotDeviceTriadService{}

func (IotDeviceTriadController) QueryList(c *gin.Context) {
	var filter entitys.IotDeviceTriadQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := deviceTriadServices.SetContext(controls.WithUserContext(c)).QueryIotDeviceTriadList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (IotDeviceTriadController) QueryVirtualDeviceList(c *gin.Context) {
	var filter entitys.VirtualDeviceQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.ProductId == 0 {
		iotgin.ResBadRequest(c, "productKey")
		return
	}
	res, total, err := deviceTriadServices.SetContext(controls.WithUserContext(c)).QueryVirtualDeviceTriadList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (IotDeviceTriadController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := deviceTriadServices.SetContext(controls.WithUserContext(c)).GetIotDeviceTriadDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (IotDeviceTriadController) Edit(c *gin.Context) {
	var req entitys.IotDeviceTriadEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceTriadServices.SetContext(controls.WithUserContext(c)).UpdateIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceTriadController) Add(c *gin.Context) {
	var req entitys.IotDeviceTriadEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := deviceTriadServices.SetContext(controls.WithUserContext(c)).AddIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (IotDeviceTriadController) Delete(c *gin.Context) {
	var req entitys.IotDeviceTriadFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).DeleteIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (IotDeviceTriadController) SetStatus(c *gin.Context) {
	var req entitys.IotDeviceTriadFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).SetStatusIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// Generator 生成三元组 （输入数量生成三元组）
func (IotDeviceTriadController) Generator(c *gin.Context) {
	var req entitys.GenerateDeviceTriad
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).GeneratorIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// GeneratorDeviceTriad 生成三元组（导入序列号生成三元组）
func (ct *IotDeviceTriadController) GeneratorDeviceTriad(c *gin.Context) {
	ct.generatorOrImportDeviceTriad(c, false)
}

func (ct *IotDeviceTriadController) ImportDeviceTriad(c *gin.Context) {
	ct.generatorOrImportDeviceTriad(c, true)
}

// 生成或者导入三元组
func (ct *IotDeviceTriadController) generatorOrImportDeviceTriad(c *gin.Context, isImport bool) {
	//参数解析和验证
	productId := c.PostForm("productId")
	numberForm := c.PostForm("number")
	number, err2 := strconv.ParseInt(numberForm, 0, 64)
	if err2 != nil {
		iotgin.ResBadRequest(c, "number")
		return
	}
	batchId := c.PostForm("batchId")
	ctx := controls.WithUserContext(c)
	var (
		productKey            = ""
		deviceNatureKey int32 = 0
	)
	if productId != "" {
		//查询产品信息
		productIdInt, err := iotutil.ToInt64AndErr(productId)
		if err == nil {
			productRes, err := rpc.ClientOpmProductService.FindById(ctx, &protosService.OpmProductFilter{Id: productIdInt})
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			}
			if productRes.Code != 200 {
				iotgin.ResBadRequest(c, productRes.Message)
				return
			}
		}
	}
	req := entitys.GenerateDeviceTriad{
		Number:          iotutil.ToInt32(number),
		ProductKey:      productKey,
		ProductId:       productId,
		Batch:           batchId,
		DeviceNatureKey: deviceNatureKey,
		AccountType:     controls.GetAccountType(c),
	}
	if err := req.CheckGenerateParams(); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	multiFiles, _ := c.MultipartForm()
	file := multiFiles.File["file"][0]
	fileOpen, err := file.Open()
	ext := iotutil.GetExt(file.Filename)
	var rows [][]string
	if ext == ".csv" {
		tmpFile, err := os.CreateTemp("", "uploaded-file")
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		defer tmpFile.Close()
		if _, err := io.Copy(tmpFile, fileOpen); err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		// 读取 CSV 文件内容
		rows, err = parseCSV(tmpFile.Name())
		if len(rows) <= 1 {
			iotgin.ResBadRequest(c, "导入的Csv无任何数据")
			return
		}
	} else {
		f, _ := excelize.OpenReader(fileOpen)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}

		sheetName := f.GetSheetList()[0]
		rows, err = f.GetRows(sheetName)
		if err != nil {
			iotgin.ResBadRequest(c, err.Error())
			return
		}
		if len(rows) <= 1 {
			iotgin.ResBadRequest(c, "导入的Excel无任何数据")
			return
		}
	}
	exportData := make(map[string]interface{})
	repeatData := make([]string, 0)
	if isImport {
		startRow := 0
		importData := make([]entitys.DeviceImportData, 0)
		for i := 2; i <= len(rows); i++ {
			startRow++
			devId := rows[startRow][0]    // f.GetCellValue(sheetName, fmt.Sprintf("A%d", startRow))
			userName := rows[startRow][1] // f.GetCellValue(sheetName, fmt.Sprintf("B%d", startRow))
			password := rows[startRow][2] // f.GetCellValue(sheetName, fmt.Sprintf("C%d", startRow))
			sn := rows[startRow][3]       // f.GetCellValue(sheetName, fmt.Sprintf("D%d", startRow))
			if _, ok := exportData[devId]; !ok {
				exportData[devId] = devId
				importData = append(importData, entitys.DeviceImportData{Sn: sn, DeviceId: devId, UserName: userName, Password: password})
			} else {
				repeatData = append(repeatData, devId)
			}
		}
		if len(repeatData) > 0 {
			iotgin.ResBadRequest(c, "导入数据中有重复的设备Id")
			return
		}
		if len(importData) != int(req.Number) {
			iotgin.ResBadRequest(c, "导入数量与填写数量不一致")
			return
		}
		req.Devices = importData
	} else {
		startRow := 0
		importData := make([]string, 0)
		for i := 2; i <= len(rows); i++ {
			startRow++
			sn := rows[startRow][0] //, _ := f.GetCellValue(sheetName, fmt.Sprintf("A%d", startRow))
			if _, ok := exportData[sn]; !ok {
				exportData[sn] = sn
				importData = append(importData, sn)
			} else {
				repeatData = append(repeatData, sn)
			}
		}
		if len(repeatData) > 0 {
			iotgin.ResBadRequest(c, "导入数据中有重复的设备Id")
			return
		}
		if len(importData) != int(req.Number) {
			iotgin.ResBadRequest(c, "导入序列号与填写数量不一致")
			return
		}
		req.SerialNumbers = importData
	}

	//调用微服务
	err = deviceTriadServices.SetContext(ctx).GeneratorIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 解析CSV文件内容
func parseCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	// 遍历CSV文件的每一行
	var rows [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		newRow := make([]string, 0)
		// 将每行数据写入新的工作表行
		for _, field := range record {
			newRow = append(newRow, field)
		}
		rows = append(rows, newRow)
	}
	return rows, err
}

// CreateDeviceTriad 创建三元组（通过数量创建）
func (ct *IotDeviceTriadController) CreateDeviceTriad(c *gin.Context) {
	//参数解析和验证
	var req entitys.GenerateDeviceTriad
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.AddMode == 1 {
		if err := req.CheckGenerateParams(); err != nil {
			iotgin.ResBadRequest(c, err.Error())
			return
		}
		req.SerialNumbers = []string{}
		for i := 0; i < int(req.Number); i++ {
			req.SerialNumbers = append(req.SerialNumbers, "XN"+iotutil.GetSecret(6))
		}
	}
	//调用微服务
	tenantId, _ := c.Get("tenantId")
	accountType, _ := c.Get("accountType")
	req.TenantId = tenantId.(string)
	req.AccountType = iotutil.ToInt32(accountType)
	req.IsTest = 1
	req.UseType = iotconst.Use_Type_Device_Normal
	//req.ProductId = req.ProductId //接口定义问题
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).GeneratorTestIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// CreateVirtualDeviceTriad 创建三元组（通过数量创建）
func (ct *IotDeviceTriadController) CreateVirtualDeviceTriad(c *gin.Context) {
	//参数解析和验证
	var req entitys.GenerateDeviceTriad
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.RegionServerId == 0 {
		req.RegionServerId = 1 //默认是中国地区服务器
	}
	//调用微服务
	tenantId, _ := c.Get("tenantId")
	accountType, _ := c.Get("accountType")
	req.TenantId = tenantId.(string)
	req.AccountType = iotutil.ToInt32(accountType)
	req.Number = 1
	req.SerialNumbers = []string{"XN" + iotutil.GetSecret(6)}
	req.AddMode = 3 //新增模式
	req.IsTest = 1  //测试新增
	req.UseType = iotconst.Use_Type_Device_Real_Test
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).GeneratorTestIotDeviceTriad(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (ct *IotDeviceTriadController) AddAppAccount(c *gin.Context) {
	//参数解析和验证
	var req entitys.AddAppAccountEntity
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.RegionServerId == 0 {
		req.RegionServerId = 1 //默认是中国地区服务器
	}
	req.TenantId = c.GetString("tenantId")
	err = deviceTriadServices.SetContext(controls.WithUserContext(c)).AddAppAccount(req, nil)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// TODO GetDefaultApp获取默认APP名称和下载二维码
func (ct *IotDeviceTriadController) GetDefaultApp(c *gin.Context) {
	var req entitys2.OemAppCommonReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	mp := services.GetBaseDataValue("oem_app_default_download_url", context.Background())
	if mp == nil {
		iotgin.ResErrCli(c, errors.New("参数配置异常"))
		return
	}
	url := iotutil.ToString(mp[services.GetOemAppEnv()])

	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, map[string]string{
		"url":      url,
		"name":     config.Global.DefaultApp.AppName,
		"tenantId": config.Global.DefaultApp.TenantId,
		"appKey":   config.Global.DefaultApp.AppKey,
	})
}
