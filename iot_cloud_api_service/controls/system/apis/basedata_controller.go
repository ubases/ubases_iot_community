package apis

import (
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	"cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"

	"github.com/gin-gonic/gin"
)

var BaseDatacontroller BaseDataController

type BaseDataController struct {
} //用户操作控制器

var basedataservices services.BaseDataService = services.BaseDataService{}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param id path  string true "id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dictdata/detail/{id} [get]
func (BaseDataController) GetBaseDataDetail(c *gin.Context) {
	id := c.Param("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	res, err := basedataservices.GetBaseDataDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	result := make([]map[string]interface{}, 0)
	for _, item := range res.Data {
		obj := map[string]interface{}{
			"dictId":      iotutil.ToString(item.DictCode),
			"dictLabel":   item.DictLabel,
			"dictValue":   iotutil.ToInterface(item.DictValue, int(item.ValueType)),
			"dictType":    item.DictType,
			"valueType":   item.ValueType,
			"dictSort":    item.DictSort,
			"pinyin":      item.Pinyin,
			"firstletter": item.Firstletter,
			"listimg":     item.Listimg,
		}
		result = append(result, obj)
	}
	iotgin.ResSuccess(c, result)
}

// @Summary basedata
// @Description query basedata list
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseDataQuery true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dictdata/list [get]
func (BaseDataController) QueryBaseDataList(c *gin.Context) {
	var resq entitys.BaseDataQuery
	err := c.BindQuery(&resq)
	res, err := basedataservices.QueryBaseDataList(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	result := make([]map[string]interface{}, 0)
	for _, item := range res {
		obj := map[string]interface{}{
			"dictId":      iotutil.ToString(item.DictCode),
			"dictLabel":   item.DictLabel,
			"dictValue":   iotutil.ToInterface(item.DictValue, int(item.ValueType)),
			"dictType":    item.DictType,
			"valueType":   item.ValueType,
			"dictSort":    item.DictSort,
			"pinyin":      item.Pinyin,
			"firstletter": item.Firstletter,
			"listimg":     item.Listimg,
		}
		result = append(result, obj)
	}
	iotgin.ResSuccess(c, result)
}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseData true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dictdata/add [post]
func (BaseDataController) AddTConfigDictData(c *gin.Context) {
	var req entitys.BaseData
	err := c.ShouldBindJSON(&req)
	err = basedataservices.AddBaseData(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description edit basedata
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseData true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dictdata/edit [post]
func (BaseDataController) EditBaseData(c *gin.Context) {
	var resq entitys.BaseData
	err := c.ShouldBindJSON(&resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resq.DictId == "" {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	err = basedataservices.UpdateBaseData(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description delete basedata
// @Tags basedata
// @Accept application/json
// @Param id path string true "id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dictdata/delete/{id} [post]
func (BaseDataController) DeleteBaseData(c *gin.Context) {
	id := c.Param("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	err := basedataservices.DeleteBaseData(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param id path  string true "id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dicttype/detail/{id} [get]
func (BaseDataController) GetBaseTypeDetail(c *gin.Context) {
	id := c.Param("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	res, err := basedataservices.GetBaseDataTypeDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// @Summary basedata
// @Description query basedata list
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseDataTypeQuery true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dicttype/list [get]
func (BaseDataController) QueryBaseDataTypeList(c *gin.Context) {
	iotlogger.LogHelper.Info("enter QueryBaseDataTypeList")
	var resq entitys.BaseDataTypeQuery
	err := c.BindQuery(&resq)
	//iotlogger.LogHelper.Info("TConfigDictTypeFilterPage request -- " + resq.String())
	res, err := basedataservices.QueryBaseDataTypeList(resq)
	iotlogger.LogHelper.Info("QueryBaseDataTypeList end")
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	result := make([]map[string]interface{}, 0)
	for _, item := range res.Data {
		obj := map[string]interface{}{
			"dictId":    iotutil.ToString(item.DictId),
			"dictName":  item.DictName,
			"dictType":  item.DictType,
			"status":    item.Status,
			"valueType": item.ValueType,
		}
		result = append(result, obj)
	}
	iotgin.ResPageSuccess(c, result, res.Total, 0)
}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseDataType true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dicttype/add [post]
func (BaseDataController) AddTConfigDictDataType(c *gin.Context) {
	var resq entitys.BaseDataType
	err := c.ShouldBindJSON(&resq)
	err = basedataservices.AddBaseDataType(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description edit basedata
// @Tags basedata
// @Accept application/json
// @Param data body entitys.BaseDataType true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dicttype/edit [post]
func (BaseDataController) EditBaseDataType(c *gin.Context) {
	var resq entitys.BaseDataType
	err := c.ShouldBindJSON(&resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resq.DictID == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = basedataservices.UpdateBaseDataType(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description delete basedata
// @Tags basedata
// @Accept application/json
// @Param id path  string true "id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/dicttype/delete/{id} [post]
func (BaseDataController) DeleteBaseDataType(c *gin.Context) {
	id := c.Param("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	err := basedataservices.DeleteBaseDataType(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param id path  string true "id"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/translate/detail [get]
func (BaseDataController) GetTConfigTranslateDetail(c *gin.Context) {
	code := c.Query("code")
	if iotutil.IsEmpty(code) {
		iotgin.ResBadRequest(c, "code not found")
		return
	}
	res, err := basedataservices.GetTConfigTranslateDetail(code)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var configTranslateObj *proto.ConfigTranslate
	if len(res.Data) > 0 {
		configTranslateObj = res.Data[0]
	} else {
		iotgin.ResSuccess(c, res.Data)
		return
	}

	result := make([]map[string]interface{}, 0)
	result = append(result, map[string]interface{}{
		"id":   iotutil.ToString(configTranslateObj.Id),
		"code": configTranslateObj.Code,
		"zh":   configTranslateObj.Zh,
		"en":   configTranslateObj.En,
		"jp":   configTranslateObj.Jp,
	})
	iotgin.ResSuccess(c, result)
}

// @Summary basedata
// @Description
// @Tags basedata
// @Accept application/json
// @Param data body entitys.TranslateParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/translate/add [post]
func (BaseDataController) AddTConfigTranslate(c *gin.Context) {
	//id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//rep1, _ := svrcli.ClientBaseDataServerService.GetTConfigDictData(context.Background(), &protosService.TConfigDictDataFilter{ Id: id})
	//fmt.Println(rep1)
	//iotgin.ResSuccess(c,rep1)
	var resq entitys.TranslateParam
	err := c.ShouldBindJSON(&resq)
	id, err := basedataservices.AddTConfigTranslate(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// @Summary basedata
// @Description edit basedata
// @Tags basedata
// @Accept application/json
// @Param data body entitys.TranslateParam true "请求参数结构体"
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Failure 400 object iotgin.ResponseModel 失败返回
// @Router /basedata/translate/edit [post]
func (BaseDataController) EditTConfigTranslate(c *gin.Context) {
	var resq entitys.TranslateParam
	err := c.ShouldBindJSON(&resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if iotutil.ToInt64(resq.ID) == 0 {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	err = basedataservices.UpdateTConfigTranslate(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// @Summary basedata
// @Description query basedata list
// @Tags basedata
// @Accept application/json
// @Success 200 object iotgin.ResponseModel 成功返回值
// @Router /basedata/translate/language/list [get]
func (BaseDataController) QueryTranslateLanguageList(c *gin.Context) {
	res, err := basedataservices.QueryTranslateLanguageList()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res.Data)
}
