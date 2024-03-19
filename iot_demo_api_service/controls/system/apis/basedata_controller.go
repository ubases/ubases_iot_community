package apis

import (
	"cloud_platform/iot_demo_api_service/controls/system/entitys"
	"cloud_platform/iot_demo_api_service/controls/system/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
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
