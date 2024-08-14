package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"
	services2 "cloud_platform/iot_cloud_api_service/controls/global"
	langEntitys "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	apilangservice "cloud_platform/iot_cloud_api_service/controls/lang/services"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"
	dictEntitys "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Productcontroller OpmProductController

type OpmProductController struct{} //部门操作控制器

var productServices = apiservice.OpmProductService{}
var langServices = apilangservice.LangTranslateService{}

func (OpmProductController) QueryList(c *gin.Context) {
	var filter entitys.OpmProductQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmProductFilter)
	}
	filter.Query.TenantId = controls.GetTenantId(c)
	res, total, err := productServices.SetContext(controls.WithUserContext(c)).QueryOpmProductList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

// 平台查询的产品列表
func (OpmProductController) QueryListToPlatform(c *gin.Context) {
	var filter entitys.OpmProductQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if filter.Query == nil {
		filter.Query = new(entitys.OpmProductFilter)
	}
	filter.IsPlatform = true
	res, total, err := productServices.SetContext(controls.WithUserContext(c)).QueryOpmProductList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (OpmProductController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).GetOpmProductDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmProductController) QueryAllDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).GetOpmProductAllDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmProductController) Edit(c *gin.Context) {
	var req entitys.OpmProductEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := productServices.SetContext(controls.WithUserContext(c)).UpdateOpmProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 修改面板信息
func (OpmProductController) EditPanelInfo(c *gin.Context) {
	var req entitys.OpmProductEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := productServices.SetContext(controls.WithUserContext(c)).UpdateOpmProductPanelInfo(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmProductController) Add(c *gin.Context) {
	var req entitys.OpmProductEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.TenantId = controls.GetTenantId(c)
	id, err := productServices.SetContext(controls.WithUserContext(c)).AddOpmProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmProductController) Delete(c *gin.Context) {
	var req entitys.OpmProductFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = productServices.SetContext(controls.WithUserContext(c)).DeleteOpmProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetStatus 设置状态
func (OpmProductController) SetStatus(c *gin.Context) {
	var req entitys.OpmProductFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 || req.Status == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = productServices.SetContext(controls.WithUserContext(c)).SetStatusOpmProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetShelf 设置状态
func (OpmProductController) SetShelf(c *gin.Context, status int32) {
	var req entitys.OpmProductFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	req.Status = status // 已上架
	err = productServices.SetContext(controls.WithUserContext(c)).SetStatusOpmProduct(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// QueryProductNetworkGuide 产品配网引导数据查询
func (OpmProductController) QueryProductNetworkGuide(c *gin.Context) {
	id := c.Query("productId")
	if id == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryProductNetworkGuide(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (s OpmProductController) QueryProductNetworkGuideLang(c *gin.Context) {
	id := c.Query("productId")
	if id == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	ctx := controls.WithUserContext(c)
	res, err := productServices.SetContext(ctx).QueryProductNetworkGuide(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//读取翻译
	sourceIds := []string{}
	for _, data := range res {
		for _, step := range data.Steps {
			langKey := fmt.Sprintf("%d", step.Id)
			sourceIds = append(sourceIds, langKey)
		}
	}
	langList, err := langServices.SetContext(ctx).QueryLangList(iotconst.LANG_NETWORK_GUIDE, sourceIds)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	resultMap := make([]map[string]interface{}, 0)
	for _, data := range res {
		for j, step := range data.Steps {
			stepIndex := j + 1
			langKey := fmt.Sprintf("%d", step.Id)
			newRow := make(map[string]interface{})
			newRow["id"] = langKey
			newRow["type"] = data.Type
			newRow["step"] = stepIndex
			newRow["instruction"] = step.Instruction
			newRow["instructionEn"] = step.InstructionEn
			newRow["langKey"] = langKey
			var langs []*langEntitys.LangTranslateEntitys
			if val, ok := langList[iotutil.ToString(langKey)]; ok {
				langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "instruction", step.Instruction, step.InstructionEn, langTypes, val)
			} else {
				langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "instruction", step.Instruction, step.InstructionEn, langTypes, []*langEntitys.LangTranslateEntitys{})
			}
			for _, lang := range langs {
				if lang.Lang != "zh" && lang.Lang != "en" {
					newRow[lang.Lang] = lang
				}
			}
			resultMap = append(resultMap, newRow)
		}
	}

	var newLangTypes []*dictEntitys.DictKeyVal
	for _, lang := range langTypes {
		if lang.Code != "zh" && lang.Code != "en" {
			newLangTypes = append(newLangTypes, lang)
		}
	}

	//根据了type进行排序
	sort.Slice(resultMap, func(i, j int) bool {
		return resultMap[i]["type"].(int32) > resultMap[j]["type"].(int32)
	})

	iotgin.ResSuccess(c, map[string]interface{}{
		"list":      resultMap,
		"langTypes": newLangTypes,
	})
}

// QueryProductDefaultNetworkGuide 产品配网引导数据查询
func (OpmProductController) QueryProductDefaultNetworkGuide(c *gin.Context) {
	id := c.Query("productId")
	if id == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	networkGuideTypeStr := c.Query("type")
	if networkGuideTypeStr == "" {
		iotgin.ResBadRequest(c, "type")
		return
	}
	networkGuideType := iotutil.ToInt32(networkGuideTypeStr)
	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryProductDefaultNetworkGuide(id, networkGuideType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// SaveProductNetworkGuide 产品配网引导数据提交
func (OpmProductController) SaveProductNetworkGuide(c *gin.Context) {
	var req entitys.OpmNetworkGuideEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//后台增加步骤限制
	if len(req.Steps) > 4 {
		iotgin.ResBusinessP(c, "步骤不能超过4步")
		return
	}
	id, err := productServices.SetContext(controls.WithUserContext(c)).SaveProductNetworkGuide(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// SetProductNetworkGuideType 设置产品配网类型
func (OpmProductController) SetProductNetworkGuideType(c *gin.Context) {
	var req entitys.ChangeNetworkGuideRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = productServices.SetContext(controls.WithUserContext(c)).SetProductNetworkGuideType(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 查询可选择的模组
func (OpmProductController) QueryModuleList(c *gin.Context) {
	baseProductId := c.Query("baseProductId")
	if baseProductId == "" {
		iotgin.ResBadRequest(c, "品类编号不能为空")
		return
	}
	intBaseProductId := iotutil.ToInt64(baseProductId)
	if intBaseProductId == 0 {
		iotgin.ResBadRequest(c, "品类编号不能为空")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).GetOpenModuleListByProductId(intBaseProductId)
	if intBaseProductId == 0 {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 提交选择模组
func (OpmProductController) SaveProductModule(c *gin.Context) {
	var req entitys.OpmProductModuleRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.IsCustom = 2 //非自定义固件
	err = productServices.SetContext(controls.WithUserContext(c)).SaveOpenProductAndModuleRelation(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 提交选择固件
func (OpmProductController) SaveProductFirmware(c *gin.Context) {
	var req entitys.OpmProductModuleRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.IsCustom = 1
	req.ModuleId = 0
	err = productServices.SetContext(controls.WithUserContext(c)).SaveOpenProductAndModuleRelation(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// RemoveProductFirmware 解除自定义固件
func (OpmProductController) RemoveProductFirmware(c *gin.Context) {
	var req entitys.OpmProductModuleRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResBadRequest(c, "id")
		return
	}
	//if req.ProductId == 0 {
	//	iotgin.ResBadRequest(c, "productId")
	//	return
	//}
	//if req.FirmwareId == 0 {
	//	iotgin.ResBadRequest(c, "firmwareId")
	//	return
	//}
	err = productServices.SetContext(controls.WithUserContext(c)).RemoveProductFirmwareRelation(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpmProductController) ChangeVersionSubmit(c *gin.Context) {
	var req entitys.OpmProductFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	err = productServices.SetContext(controls.WithUserContext(c)).ChangeVersionSubmit(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 下载MCU SDK
func (OpmProductController) DownloadMcuSdk(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResErrCli(c, errors.New("productId is empty"))
		return
	}
	// 根据产品key获取物模型，并根据MCU SDK模板生成MCU SDK压缩包，并以二进制文件流的形式返回给前端
	ctx := controls.WithUserContext(c)
	filename, err := productServices.SetContext(ctx).GenerateMcuSdkCode(ctx, productId)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("generate mcu sdk code error: ", err)
		iotgin.ResErrCli(c, err)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=mcu_sdk.zip")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filename)
}

// 查询面板列表
func (OpmProductController) QueryControlPanelList(c *gin.Context) {
	baseProductId := c.Query("baseProductId")
	if baseProductId == "" {
		iotgin.ResBadRequest(c, "品类编号不能为空")
		return
	}
	intBaseProductId := iotutil.ToInt64(baseProductId)
	if intBaseProductId == 0 {
		iotgin.ResBadRequest(c, "品类编号不能为空")
		return
	}

	var intProductId int64
	productId := c.Query("productId")
	if productId != "" {
		intProductId, _ = iotutil.ToInt64AndErr(productId)
	}

	res, err := productServices.SetContext(controls.WithUserContext(c)).GetOpenControlPanelsListByProductId(intBaseProductId, intProductId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 提交面板
func (OpmProductController) SaveProductControlPanel(c *gin.Context) {
	var req entitys.OpmProductPanelRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	panel, err := productServices.SetContext(controls.WithUserContext(c)).SaveOpenProductAndControlPanelRelation(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, panel)
}

// 提交面板
func (OpmProductController) CancelReminder(c *gin.Context) {
	var req entitys.OpmProductPanelRelationEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = productServices.SetContext(controls.WithUserContext(c)).ControlPanelRelationUpdateCreatedAt(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// QueryAllDetail 开发完成查询
func (OpmProductController) QueryCompleteDevelopDetail(c *gin.Context) {
	id := c.Query("productId")
	if id == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	//CompleteDevelopDetail
	res, err := productServices.SetContext(controls.WithUserContext(c)).GetCompleteDevelopDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 上传测试报告
func (OpmProductController) UploadTestReport(c *gin.Context) {
	var err error
	//参数解析和验证
	productIdStr := c.PostForm("productId")
	productId, err := iotutil.ToInt64AndErr(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	testResult := c.PostForm("testResult")
	testResultInt, err := iotutil.ToInt32Err(testResult)
	if err != nil {
		iotgin.ResBadRequest(c, "testResult")
		return
	}
	//测试用例文档文件
	file, err := c.FormFile("file")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	f, err := apis.SaveFileToOSS(c, file, apis.TestCaseTempPath, "xlsx", "xls")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	req := &entitys.OpmProductTestReportEntitys{
		ProductId:  productId,
		TestType:   1,
		IsValid:    1,
		FilePath:   f.FullPath,
		FileName:   f.Name,
		FileSize:   int32(f.Size),
		FileKey:    f.Key,
		TestOrigin: 1,
		TestResult: testResultInt,
	}
	ctx := controls.WithUserContext(c)
	err = productServices.SetContext(ctx).UploadTestReport(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//刷新产品更新时间
	rpc.ClientOpmProductService.Update(ctx, &protosService.OpmProduct{Id: productId})
	iotgin.ResSuccessMsg(c)
}

// 获取测试报告
func (OpmProductController) GetTestReport(c *gin.Context) {
	//参数解析和验证
	productIdStr := c.Query("productId")
	productId, err2 := strconv.ParseInt(productIdStr, 0, 64)
	if err2 != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}

	rep, err := productServices.SetContext(controls.WithUserContext(c)).GetTestReport(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, rep)
}

func (s OpmProductController) QueryLangList(c *gin.Context) {
	productId := c.Query("productId")
	productIdInt, err := iotutil.ToInt64AndErr(productId)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	filter := entitys.OpmProductQuery{
		Query: &entitys.OpmProductFilter{
			Id:       productIdInt,
			TenantId: controls.GetTenantId(c),
		},
	}
	ctx := controls.WithUserContext(c)
	res, _, err := productServices.SetContext(ctx).QueryOpmProductList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//读取翻译
	sourceIds := []string{}
	for _, data := range res {
		sourceIds = append(sourceIds, data.ProductKey)
	}
	langList, err := langServices.SetContext(ctx).QueryLangList(iotconst.LANG_PRODUCT_NAME, sourceIds)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	resultMap := make([]map[string]interface{}, 0)

	for _, data := range res {
		newRow := make(map[string]interface{})
		newRow["name"] = data.Name
		newRow["nameEn"] = data.NameEn
		newRow["productKey"] = data.ProductKey
		newRow["productName"] = data.Name
		newRow["id"] = data.Id
		var langs []*langEntitys.LangTranslateEntitys
		if val, ok := langList[data.ProductKey]; ok {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_NAME, data.ProductKey, "name", data.Name, data.NameEn, langTypes, val)
		} else {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_NAME, data.ProductKey, "name", data.Name, data.NameEn, langTypes, []*langEntitys.LangTranslateEntitys{})
		}
		for _, lang := range langs {
			newRow[lang.Lang] = lang
			if lang.Lang == "en" {
				newRow["nameEn"] = lang.FieldValue
			} else if lang.Lang == "zh" {
				newRow["name"] = lang.FieldValue
			}
		}
		resultMap = append(resultMap, newRow)
	}
	var newLangTypes []*dictEntitys.DictKeyVal
	for _, lang := range langTypes {
		if lang.Code != "zh" && lang.Code != "en" {
			newLangTypes = append(newLangTypes, lang)
		}
	}
	iotgin.ResSuccess(c, map[string]interface{}{
		"list":      resultMap,
		"langTypes": newLangTypes,
	})
}

func (s OpmProductController) SetLangs(sourceTable, sourceRowId, fieldName, fieldValue, fieldValueEn string, langTypes []*dictEntitys.DictKeyVal, data []*langEntitys.LangTranslateEntitys) []*langEntitys.LangTranslateEntitys {
	defaultData := []*langEntitys.LangTranslateEntitys{}
	for _, langType := range langTypes {
		var currData *langEntitys.LangTranslateEntitys
		for _, d := range data {
			if langType.Code == d.Lang {
				currData = d
			}
		}
		defaultItem := &langEntitys.LangTranslateEntitys{
			//SourceTable: sourceTable,
			//SourceRowId: sourceRowId,
			Lang:       langType.Code,
			FieldName:  fieldName,
			FieldValue: "",
		}
		if langType.Code == "zh" {
			if currData == nil || currData.FieldValue == "" {
				defaultItem.FieldValue = fieldValue
			}
		} else if langType.Code == "en" {
			if currData == nil || currData.FieldValue == "" {
				defaultItem.FieldValue = fieldValueEn
			}
		}
		if currData != nil {
			defaultItem = currData
		}
		defaultData = append(defaultData, defaultItem)
	}
	return defaultData
}

// QueryProductFirmwareType 查询可选择的模组
func (OpmProductController) QueryProductFirmwareType(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "产品编号不能为空")
		return
	}
	productIdInt, _ := iotutil.ToInt64AndErr(productId)
	rep, err := rpc.ClientOpmProductModuleRelationService.Lists(context.Background(),
		&protosService.OpmProductModuleRelationListRequest{Query: &protosService.OpmProductModuleRelation{ProductId: productIdInt}})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	firmwareType, _ := new(services2.DictTempData).GetDictByCode(iotconst.Dict_type_firmware_type)
	res := make([]*dictEntitys.DictKeyVal, 0)
	repeatMap := make(map[int32]int32)
	for _, d := range rep.Data {
		//去重
		if _, ok := repeatMap[d.FirmwareType]; ok {
			continue
		}
		repeatMap[d.FirmwareType] = d.FirmwareType
		res = append(res, &dictEntitys.DictKeyVal{
			Name: firmwareType.Value(d.FirmwareType),
			Code: iotutil.ToString(d.FirmwareType),
		})
	}
	iotgin.ResSuccess(c, res)
}

// 获取产品对应的智能条件和智能任务
func (OpmProductController) GetTaskOrWhereByProduct(c *gin.Context) {
	productIdStr := c.Query("productId")
	if productIdStr == "" {
		iotgin.ResBadRequest(c, "产品编号不能为空")
		return
	}
	productId, err := iotutil.ToInt64AndErr(productIdStr)
	if err != nil {
		iotgin.ResBadRequest(c, "产品编号不合法")
		return
	}
	condType := c.Query("condType")
	if condType == "" {
		iotgin.ResBadRequest(c, "条件类型不能为空")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).GetTaskOrWhereByProduct(productId, condType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res.Properties)
}

// 导出物模型
func (s OpmProductController) ExportThingsModel(c *gin.Context) {
	productIdStr := c.Query("productId")
	if productIdStr == "" {
		iotgin.ResBadRequest(c, "产品编号不能为空")
		return
	}
	setContext := productServices.SetContext(controls.WithUserContext(c))
	fileName, tempPathFile, err := setContext.Export(productIdStr) //
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	//发送文件
	c.File(tempPathFile)
}
