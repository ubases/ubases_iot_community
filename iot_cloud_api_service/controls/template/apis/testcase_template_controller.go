package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	"cloud_platform/iot_cloud_api_service/controls/template/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/template/services"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"

	"cloud_platform/iot_common/iotgin"
)

var TestcaseTemplatecontroller TplTestcaseTemplateController

type TplTestcaseTemplateController struct{} //部门操作控制器

var testcaseTemplateServices = apiservice.TplTestcaseTemplateService{}

func (TplTestcaseTemplateController) QueryList(c *gin.Context) {
	var filter entitys.TplTestcaseTemplateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := testcaseTemplateServices.QueryTplTestcaseTemplateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (TplTestcaseTemplateController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := testcaseTemplateServices.GetTplTestcaseTemplateDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GetTestReportTemplate 获取测试用例模板
func (TplTestcaseTemplateController) GetTestReportTemplate(c *gin.Context) {
	productTypeId := c.Query("productTypeId")
	if productTypeId == "" {
		iotgin.ResBadRequest(c, "productTypeId")
		return
	}
	var status int32 = 1
	res, _, err := testcaseTemplateServices.QueryTplTestcaseTemplateList(entitys.TplTestcaseTemplateQuery{
		Page: 1, Limit: 1,
		Sort:      "desc",
		SortField: "version",
		Query: &entitys.TplTestcaseTemplateFilter{
			ProductTypeId: productTypeId,
			Status:        &status,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(res) == 0 {
		iotgin.ResErrCli(c, errors.New("未获取到模板数据"))
		return
	}
	iotgin.ResSuccess(c, res[0])
}

func (TplTestcaseTemplateController) Edit(c *gin.Context) {
	//参数解析和验证
	id := c.PostForm("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	productTypeIdStr := c.PostForm("productTypeId")
	_, err2 := strconv.ParseInt(productTypeIdStr, 0, 64)
	if err2 != nil {
		iotgin.ResBadRequest(c, "productTypeId")
		return
	}
	req := entitys.TplTestcaseTemplateEntitys{
		Id:            iotutil.ToString(id),
		TplName:       c.PostForm("tplName"),
		Lang:          c.PostForm("lang"),
		TplDesc:       c.PostForm("tplDesc"),
		Version:       c.PostForm("version"),
		ProductTypeId: productTypeIdStr,
		UpdatedBy:     controls.GetUserId(c),
	}
	if err := req.UpdateCheck(); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	//测试用例文档文件
	//file, err := c.FormFile("tplFile")
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["tplFile"]
		for _, file := range files {
			f, err := apis.SaveFileToOSS(c, file, apis.TestCaseTempPath, "xlsx", "xls")
			if err != nil {
				iotgin.ResErrCli(c, err)
				return
			} else {
				req.TplFile = f.FullPath
				req.TplFileName = file.Filename
				req.TplFileSize = file.Size
				break
			}
		}
	}
	_, err = testcaseTemplateServices.UpdateTplTestcaseTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (ct *TplTestcaseTemplateController) Add(c *gin.Context) {
	//参数解析和验证
	productTypeIdStr := c.PostForm("productTypeId")
	_, err2 := strconv.ParseInt(productTypeIdStr, 0, 64)
	if err2 != nil {
		iotgin.ResBadRequest(c, "productTypeId")
		return
	}
	req := entitys.TplTestcaseTemplateEntitys{
		TplName:       c.PostForm("tplName"),
		Lang:          c.PostForm("lang"),
		TplDesc:       c.PostForm("tplDesc"),
		Version:       c.PostForm("version"),
		ProductTypeId: productTypeIdStr,
		CreatedBy:     controls.GetUserId(c),
	}
	if err := req.UpdateCheck(); err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	//测试用例文档文件
	file, err := c.FormFile("tplFile")
	if err != nil {
		iotgin.ResErrCli(c, err)
		//	iotgin.ResErrCli(c, errors.New("请确保测试用例模板文件正确上传"))
		return
	}
	f, err := apis.SaveFileToOSS(c, file, apis.TestCaseTempPath, "xlsx", "xls")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TplFile = f.FullPath
	req.TplFileSize = f.Size
	req.TplFileName = f.Name
	id, err := testcaseTemplateServices.AddTplTestcaseTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (TplTestcaseTemplateController) Delete(c *gin.Context) {
	var req entitys.TplTestcaseTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = testcaseTemplateServices.DeleteTplTestcaseTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (TplTestcaseTemplateController) SetStatus(c *gin.Context) {
	var req entitys.TplTestcaseTemplateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" || req.Status == nil {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = testcaseTemplateServices.SetStatusTplTestcaseTemplate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (s *TplTestcaseTemplateController) GetTplFile(c *gin.Context) {
	code := c.Param("tplcode")
	fileName, tempPathFile, err := s.GetImportTemplate(code)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)

	//删除临时文件
	go func() {
		//延时3面删除临时文件
		defer iotutil.PanicHandler()
		time.Sleep(3 * time.Second)
		os.Remove(tempPathFile)
	}()
}

// TODO excel模板需要修改到模板管理
func (s *TplTestcaseTemplateController) GetImportTemplate(code string) (string, string, error) {
	//导出生成excel导出模板
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	headerRow := sheet.AddRow()
	cell := headerRow.AddCell()
	cell.Value = "序列号"

	row1 := sheet.AddRow()
	cell1 := row1.AddCell()
	cell1.Value = "A1000001"

	tempPathFile := strings.Join([]string{iotconst.GetWorkTempDir(), iotutil.Uuid() + ".xlsx"}, string(filepath.Separator))
	err := file.Save(tempPathFile)
	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("save file %s error:%s", tempPathFile, err.Error()))
		return "", "", err
	}
	fileName := code + ".xlsx"

	return fileName, tempPathFile, nil
}
