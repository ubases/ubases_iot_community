package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	langEntitys "cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	dictEntitys "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

// QueryProductThingModel 基础物模型数据
func (OpmProductController) QueryProductThingModel(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	isCustomStr := c.Query("custom")
	var isCustom int32 = 0
	if isCustomStr != "" {
		isCustom = iotutil.ToInt32(isCustomStr)
	}

	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryProductThingModel(productId, isCustom)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// QueryProductFaultThingModel 故障物模型数据
func (OpmProductController) QueryProductFaultThingModel(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	var isCustom int32 = -1
	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryProductFaultThingModel(productId, isCustom)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// QueryProductThingModelAndLang 基础物模型数据
func (s OpmProductController) QueryProductThingModelAndLang(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}

	ctx := controls.WithUserContext(c)
	productInfo, err := productServices.SetContext(ctx).GetOpmProductDetail(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := productServices.SetContext(ctx).QueryProductThingModel(productId, -1)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	resultMap, err := s.convertDataTranslate(ctx, productInfo.ProductKey, res, langTypes)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
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

// 将物模型数据转换为语言key，并查询出翻译内容
func (s OpmProductController) getTranslateData(ctx context.Context, productKey string, res *entitys.OpmThingModelList) (map[string][]*langEntitys.LangTranslateEntitys, error) {
	sourceIds := []string{}
	for _, data := range res.List {
		sourceIds = append(sourceIds, fmt.Sprintf("%s_%s", productKey, data.Identifier))
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err := json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					dataType := iotutil.ToString(spec["dataType"])
					if dataType == "BOOL" {
						if val == "1" || val == "true" {
							val = "true"
						} else {
							val = "false"
						}
					}
					langKey := fmt.Sprintf("%s_%s_%s", productKey, data.Identifier, val)
					//[{"custom":1,"dataType":"ENUM","name":"打开","value":1},{"custom":1,"dataType":"ENUM","name":"关闭","value":0}]
					sourceIds = append(sourceIds, langKey)
				}
			}
		}
	}
	langList, err := langServices.SetContext(ctx).QueryLangList(iotconst.LANG_PRODUCT_THINGS_MODEL, sourceIds)
	if err != nil {
		return nil, err
	}
	return langList, nil
}

// 数据列表转换为翻译数据
func (s OpmProductController) convertDataTranslate(ctx context.Context, productKey string, res *entitys.OpmThingModelList, langTypes []*dictEntitys.DictKeyVal) ([]map[string]interface{}, error) {
	//将物模型数据转换为语言key，并查询出翻译内容
	langList, err := s.getTranslateData(ctx, productKey, res)
	if err != nil {
		return nil, err
	}
	resultMap := make([]map[string]interface{}, 0)

	for _, data := range res.List {
		newRow := make(map[string]interface{})
		newRow["id"] = data.Id
		newRow["name"] = data.Name
		newRow["isChild"] = false
		newRow["funcValue"] = data.Name
		newRow["identifier"] = data.Identifier
		newRow["zh"] = ""
		newRow["en"] = ""
		langKey := fmt.Sprintf("%s_%s", productKey, data.Identifier)
		newRow["langKey"] = langKey
		var langs []*langEntitys.LangTranslateEntitys
		if val, ok := langList[langKey]; ok {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "identifier", data.Name, "", langTypes, val)
		} else {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, langKey, "identifier", data.Name, "", langTypes, []*langEntitys.LangTranslateEntitys{})
		}
		for _, lang := range langs {
			newRow[lang.Lang] = lang
		}
		resultMap = append(resultMap, newRow)
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err = json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				//[{"custom":1,"dataType":"ENUM","name":"打开","value":1},{"custom":1,"dataType":"ENUM","name":"关闭","value":0}]
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					desc := spec["desc"]
					name := spec["name"]
					if desc == "" || desc == nil {
						desc = name
					}
					dataType := iotutil.ToString(spec["dataType"])
					//数值转换（BOOL类型特殊处理）
					if dataType == "BOOL" {
						if val == "1" || val == "true" {
							val = "true"
						} else {
							val = "false"
						}
					}
					funcLangKey := fmt.Sprintf("%s_%s_%v", productKey, data.Identifier, val)
					newRowSpec := make(map[string]interface{})
					newRowSpec["id"] = funcLangKey
					newRowSpec["isChild"] = true
					newRowSpec["name"] = ""
					newRowSpec["identifier"] = ""
					newRowSpec["funcValue"] = desc
					newRowSpec["funcValueKey"] = name
					newRowSpec["langKey"] = funcLangKey
					newRowSpec["zh"] = ""
					newRowSpec["en"] = ""
					var langSpecs []*langEntitys.LangTranslateEntitys
					if val, ok := langList[funcLangKey]; ok {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, funcLangKey, "identifier_value", "", "", langTypes, val)
					} else {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_THINGS_MODEL, funcLangKey, "identifier_value", "", "", langTypes, []*langEntitys.LangTranslateEntitys{})
					}
					for _, lang := range langSpecs {
						newRowSpec[lang.Lang] = lang
					}
					resultMap = append(resultMap, newRowSpec)
				}
			}
		}
	}
	return resultMap, nil
}

// 新增功能定义
func (OpmProductController) EditThingModel(c *gin.Context) {
	var req entitys.OpmThingModelPropertiesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).UpdateOpmThingModel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmProductController) AddThingModel(c *gin.Context) {
	var req entitys.OpmThingModelPropertiesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).AddOpmThingModel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmProductController) DeleteThingModel(c *gin.Context) {
	var req entitys.OpmThingModelPropertiesEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	res, err := productServices.SetContext(controls.WithOpenUserContext(c)).DeleteOpmThingModel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpmProductController) AddStandThingModel(c *gin.Context) {
	var req entitys.AddStandardThingModelRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).AddOpmThingModelByStandard(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// QueryStandardThingModel 查询基础物模型数据
func (OpmProductController) QueryStandardThingModel(c *gin.Context) {
	productId := c.Query("baseProductId")
	if productId == "" {
		iotgin.ResBadRequest(c, "baseProductId")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryStandardThingModel(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// QueryControlPanelLang 面板翻译数据
func (s OpmProductController) QueryControlPanelLang(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	ctx := controls.WithUserContext(c)
	productInfo, err := productServices.SetContext(ctx).GetOpmProductDetail(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := productServices.SetContext(ctx).QueryProductThingModel(productId, -1)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	controlPanelId := productInfo.ControlPanelId
	//读取翻译
	sourceIds := []string{}
	for _, data := range res.List {
		sourceIds = append(sourceIds, fmt.Sprintf("panel_%d_%s_%s", controlPanelId, productInfo.ProductKey, data.Identifier))
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err = json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					langKey := fmt.Sprintf("panel_%d_%s_%s_%s", controlPanelId, productInfo.ProductKey, data.Identifier, val)
					sourceIds = append(sourceIds, langKey)
				}
			}
		}
	}
	langList, err := langServices.SetContext(ctx).QueryLangList(iotconst.LANG_PRODUCT_CONTROL_PANEL, sourceIds)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()

	resultMap := make([]map[string]interface{}, 0)
	for _, data := range res.List {
		newRow := make(map[string]interface{})
		newRow["id"] = data.Id
		newRow["name"] = data.Name
		newRow["nameEn"] = data.Identifier
		langKey := fmt.Sprintf("panel_%d_%s_%s", controlPanelId, productInfo.ProductKey, data.Identifier)
		newRow["langKey"] = langKey
		newRow["showLangKey"] = data.Identifier
		var langs []*langEntitys.LangTranslateEntitys
		if val, ok := langList[langKey]; ok {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_CONTROL_PANEL, langKey, "identifier", data.Name, "", langTypes, val)
		} else {
			langs = s.SetLangs(iotconst.LANG_PRODUCT_CONTROL_PANEL, langKey, "identifier", data.Name, "", langTypes, []*langEntitys.LangTranslateEntitys{})
		}
		for _, lang := range langs {
			if lang.Lang == "zh" {
				newRow["name"] = lang.FieldValue
			} else if lang.Lang == "en" {
				newRow["nameEn"] = lang.FieldValue
			}
		}
		resultMap = append(resultMap, newRow)
		if data.DataSpecsList != "" {
			mapSpecs := []map[string]interface{}{}
			err = json.Unmarshal([]byte(data.DataSpecsList), &mapSpecs)
			if err == nil {
				//[{"custom":1,"dataType":"ENUM","name":"打开","value":1},{"custom":1,"dataType":"ENUM","name":"关闭","value":0}]
				for _, spec := range mapSpecs {
					val := iotutil.ToString(spec["value"])
					funcLangKey := fmt.Sprintf("panel_%d_%s_%s_%s", controlPanelId, productInfo.ProductKey, data.Identifier, val)
					funcShowLangKey := fmt.Sprintf("%s_%s", data.Identifier, val)
					newRowSpec := make(map[string]interface{})
					newRowSpec["id"] = funcLangKey
					newRowSpec["langKey"] = funcLangKey
					newRowSpec["showLangKey"] = funcShowLangKey
					newRowSpec["name"] = ""
					newRowSpec["nameEn"] = ""
					var langSpecs []*langEntitys.LangTranslateEntitys
					if val, ok := langList[funcLangKey]; ok {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_CONTROL_PANEL, funcLangKey, "identifier_value", "", "", langTypes, val)
					} else {
						langSpecs = s.SetLangs(iotconst.LANG_PRODUCT_CONTROL_PANEL, funcLangKey, "identifier_value", "", "", langTypes, []*langEntitys.LangTranslateEntitys{})
					}
					for _, lang := range langSpecs {
						if lang.Lang == "zh" {
							newRowSpec["name"] = lang.FieldValue
						} else if lang.Lang == "en" {
							newRowSpec["nameEn"] = lang.FieldValue
						}
					}
					resultMap = append(resultMap, newRowSpec)
				}
			}
		}
	}
	iotgin.ResSuccess(c, resultMap)
}

// ControlPanelCustomResource 查询面板自定义资源
func (s *OpmProductController) ControlPanelCustomResource(c *gin.Context) {
	//默认是APP的资源下载
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResErrCli(c, errors.New("参数错误 productId"))
		return
	}
	thisContext := controls.WithUserContext(c)
	productIdInt := iotutil.ToInt64(productId)
	ctx := controls.WithUserContext(c)
	products, err := rpc.ClientOpmProductService.FindById(ctx, &proto.OpmProductFilter{Id: productIdInt})
	if err != nil {
		iotgin.ResBadRequest(c, err.Error())
		return
	}
	if products.Code != 200 {
		iotgin.ResBadRequest(c, products.Message)
		return
	}
	if len(products.Data) == 0 {
		iotgin.ResBadRequest(c, "产品不存在")
		return
	}
	productInfo := products.Data[0]
	//productInfo.ControlPanelId
	//获取资源部信息，读取支持的语言
	res, err := rpc.ClientLangResourcesPackageService.Lists(thisContext, &proto.LangResourcePackageListRequest{
		Query: &proto.LangResourcePackage{
			BelongType: 4,
			BelongId:   productInfo.ControlPanelId,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if res.Code != 200 {
		iotgin.ResErrCli(c, errors.New(res.Message))
		return
	}
	if len(res.Data) == 0 {
		iotgin.ResErrCli(c, errors.New("未找到任何可导出的资源"))
		return
	}

	//优先下载自己的资源
	crs, err := rpc.ClientLangCustomResourceService.Lists(thisContext, &proto.LangCustomResourcesListRequest{
		Query: &proto.LangCustomResources{
			BelongType: 4,
			BelongId:   productInfo.ControlPanelId,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if crs.Code != 200 {
		iotgin.ResErrCli(c, errors.New(crs.Message))
		return
	}

	resultArr := map[string]map[string]string{}
	//如果自定义为0，则读取基础的翻译
	if len(crs.Data) == 0 {
		rep, err := rpc.ClientLangResourcesService.Lists(thisContext, &proto.LangResourcesListRequest{
			Query: &proto.LangResources{
				BelongType: 4,
				BelongId:   productInfo.ControlPanelId,
			},
		})
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		if rep.Code != 200 {
			iotgin.ResErrCli(c, errors.New(rep.Message))
			return
		}
		for _, d := range rep.Data {
			if _, ok := resultArr[d.Code]; !ok {
				resultArr[d.Code] = map[string]string{"code": d.Code, d.Lang: d.Value}
			} else {
				resultArr[d.Code][d.Lang] = d.Value
			}
		}
	} else {
		for _, d := range crs.Data {
			if _, ok := resultArr[d.Code]; !ok {
				resultArr[d.Code] = map[string]string{"code": d.Code, d.Lang: d.Value}
			} else {
				resultArr[d.Code][d.Lang] = d.Value
			}
		}
	}
	newResultArr := []map[string]interface{}{}
	for _, row := range resultArr {
		rowMap := map[string]interface{}{}
		rowMap["id"] = row["code"]
		rowMap["langKey"] = row["code"]
		rowMap["showLangKey"] = row["code"]
		for _, lang := range res.Data[0].Langs {
			switch lang {
			case "zh":
				rowMap["name"] = row[lang]
			case "en":
				rowMap["nameEn"] = row[lang]
			default:
				rowMap[lang] = row[lang]
			}
		}
		newResultArr = append(newResultArr, rowMap)
	}
	iotgin.ResSuccess(c, newResultArr)
}

func (OpmProductController) ResetStandThingFunc(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	productIdInt, err := iotutil.ToInt64AndErr(productId)
	if err != nil {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).ResetStandardFunc(productIdInt)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (OpmProductController) SetThingsModelSceneFunc(c *gin.Context) {
	var req entitys.AddStandardThingModelRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).SetThingsModelSceneFunc(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 设置预约功能
func (OpmProductController) SetAppointmentFunc(c *gin.Context) {
	var req entitys.AddStandardThingModelRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).SetAppointmentFunc(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 设置功能层级
func (OpmProductController) SetFuncLevel(c *gin.Context) {
	var req entitys.SetFuncLevelRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).SetFuncLevel(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 上移
func (OpmProductController) SetFuncMoveUp(c *gin.Context) {
	var req entitys.SetFuncSortRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.Sort = 1 //上移
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).SetFuncSort(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 下移
func (OpmProductController) SetFuncMoveDown(c *gin.Context) {
	var req entitys.SetFuncSortRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.Sort = 0 //下移
	id, err := productServices.SetContext(controls.WithOpenUserContext(c)).SetFuncSort(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// QueryAppointmentFuncList 查询预约物模型数据
func (OpmProductController) QueryAppointmentFuncList(c *gin.Context) {
	productId := c.Query("productId")
	if productId == "" {
		iotgin.ResBadRequest(c, "productId")
		return
	}
	res, err := productServices.SetContext(controls.WithUserContext(c)).QueryAppointmentFuncList(productId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}
