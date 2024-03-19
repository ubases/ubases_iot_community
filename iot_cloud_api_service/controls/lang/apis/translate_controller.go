package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/lang/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/lang/services"
	services2 "cloud_platform/iot_cloud_api_service/controls/oem/services"
	sysEntitys "cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Translatecontroller LangTranslateController

type LangTranslateController struct{} //部门操作控制器

var translateServices = apiservice.LangTranslateService{}

// LangTypeList 从字典表获取语言类型
func (s LangTranslateController) LangTypeList(c *gin.Context) {
	dicService := services.BaseDataService{}
	list, err := dicService.QueryBaseDataList(sysEntitys.BaseDataQuery{
		DictType: "language_type",
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var res []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
	for _, data := range list {
		res = append(res, struct {
			Name string `json:"name"`
			Code string `json:"code"`
		}{Name: data.DictLabel, Code: data.DictValue})
	}
	iotgin.ResSuccess(c, res)
}

// TranslateGetV2 获取指定的翻译详细（业务翻译获取）
func (s LangTranslateController) TranslateGetV2(c *gin.Context) {
	sourceTable := c.Query("sourceTable")
	sourceRowId := c.Query("sourceRowId")
	fieldName := c.Query("fieldName")
	name := c.Query("defaultVal")
	nameEn := c.Query("defaultValEn")
	thisContext := controls.WithUserContext(c)
	ret, err := rpc.ClientLangTranslateService.Lists(thisContext, &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable: sourceTable,
			SourceRowId: sourceRowId,
			FieldName:   fieldName,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if ret.Code != 200 {
		iotgin.ResErrCli(c, errors.New(ret.Message))
		return
	}

	dicService := services.BaseDataService{}
	langTypes := dicService.GetLangType()
	//sort.Slice(&langTypes, func(i, j int) bool {
	//	return langTypes[i].Code > langTypes[j].Code
	//})
	//sort.Slice(&langTypes, func(i, j int) bool {
	//	return langTypes[i].Code < langTypes[j].Code
	//})

	var res []*entitys.LangTranslateEntitys
	for _, langType := range langTypes {
		var currData *protosService.LangTranslate
		for _, d := range ret.Data {
			if langType.Code == d.Lang {
				currData = d
			}
		}
		resItem := entitys.LangTranslateEntitys{
			SourceTable: sourceTable,
			SourceRowId: sourceRowId,
			Lang:        langType.Code,
			FieldName:   fieldName,
		}
		if currData != nil {
			resItem.FieldValue = currData.FieldValue
			if currData.Id != 0 {
				resItem.Id = iotutil.ToString(currData.Id)
			}
		} else {
			switch resItem.Lang {
			case "zh":
				resItem.FieldValue = name
			case "en":
				resItem.FieldValue = nameEn
			}
		}
		switch resItem.Lang {
		case "zh":
			resItem.Sort = 1
		case "en":
			resItem.Sort = 2
		default:
			resItem.Sort = 3
		}
		res = append(res, &resItem)
	}

	//将zh排列在en后方
	//sort.Slice(res, func(i, j int) bool {
	//	return res[i].Lang < res[j].Lang
	//})
	//排序将zh、en排在最前方
	sort.Slice(res, func(i, j int) bool {
		if res[j].Lang == "zh" || res[j].Lang == "en" {
			return res[i].Sort < res[j].Sort
		} else {
			return res[i].Sort > res[j].Sort || res[i].Lang < res[j].Lang
		}
	})

	iotgin.ResSuccess(c, res)
}

// TranslateGet 获取指定的翻译详细（业务翻译获取）
func (s LangTranslateController) TranslateGet(c *gin.Context) {
	sourceTable := c.Query("sourceTable")
	sourceRowId := c.Query("sourceRowId")
	thisContext := controls.WithUserContext(c)
	ret, err := rpc.ClientLangTranslateService.Lists(thisContext, &protosService.LangTranslateListRequest{
		Query: &protosService.LangTranslate{
			SourceTable: sourceTable,
			SourceRowId: sourceRowId,
		},
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if ret.Code != 200 {
		iotgin.ResErrCli(c, errors.New(ret.Message))
		return
	}
	var res []*entitys.LangTranslateEntitys
	for _, data := range ret.Data {
		resItem := entitys.LangTranslateEntitys{
			//SourceTable: data.SourceTable,
			//SourceRowId: data.SourceRowId,
			Lang:       data.Lang,
			FieldName:  data.FieldName,
			FieldType:  data.FieldType,
			FieldValue: data.FieldValue,
		}
		if data.Id != 0 {
			resItem.Id = iotutil.ToString(data.Id)
		}
		res = append(res, &resItem)
	}
	iotgin.ResSuccess(c, res)
}

// TranslateSave 翻译保存（业务翻译保存）
func (s LangTranslateController) TranslateSave(platformType int32, c *gin.Context) {
	var req entitys.LangTranslateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	translateList := []*protosService.BatchSaveTranslateItem{}
	for _, item := range req.TranslateList {
		translateItem := protosService.BatchSaveTranslateItem{
			Lang:       item.Lang,
			FieldName:  item.FieldName,
			FieldType:  item.FieldType,
			FieldValue: item.FieldValue,
		}
		if item.Id != "" {
			translateItem.Id = iotutil.ToInt64(item.Id)
		}
		translateList = append(translateList, &translateItem)
	}
	thisContext := controls.WithUserContext(c)
	ret, err := rpc.ClientLangTranslateService.BatchCreate(thisContext, &protosService.BatchSaveTranslate{
		SourceRowId:  req.SourceRowId,
		SourceTable:  req.SourceTable,
		PlatformType: platformType,
		List:         translateList,
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if ret.Code != 200 {
		iotgin.ResErrCli(c, errors.New(ret.Message))
		return
	}
	s.ClearCached(req.SourceRowId, req.SourceTable)
	iotgin.ResSuccessMsg(c)
}

func (s LangTranslateController) ClearCached(sourceRowId, sourceTable string) {
	defer iotutil.PanicHandler(sourceRowId, sourceTable)
	if strings.Index(sourceTable, iotconst.LANG_OEM_APP_ROOMS) != -1 {
		uiConfig := services2.OemAppUiConfigService{Ctx: context.Background()}
		appKey := strings.ReplaceAll(sourceTable, iotconst.LANG_OEM_APP_ROOMS+"_", "")
		uiConfig.ClearLangCachedByAppKey(appKey)
	}
}

// 基础添加删改查
func (LangTranslateController) QueryList(c *gin.Context) {
	var filter entitys.LangTranslateQuery
	err := c.ShouldBindJSON(&filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, total, err := translateServices.SetContext(controls.WithOpenUserContext(c)).QueryLangTranslateList(filter)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(filter.Page))
}

func (LangTranslateController) QueryDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		iotgin.ResBadRequest(c, "id")
		return
	}
	res, err := translateServices.SetContext(controls.WithOpenUserContext(c)).GetLangTranslateDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (LangTranslateController) Edit(c *gin.Context) {
	var req entitys.LangTranslateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := translateServices.SetContext(controls.WithOpenUserContext(c)).UpdateLangTranslate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangTranslateController) Add(c *gin.Context) {
	var req entitys.LangTranslateEntitys
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	id, err := translateServices.SetContext(controls.WithOpenUserContext(c)).AddLangTranslate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

func (LangTranslateController) Delete(c *gin.Context) {
	var req entitys.LangTranslateFilter
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = translateServices.SetContext(controls.WithOpenUserContext(c)).DeleteLangTranslate(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}
