package router

import (
	"cloud_platform/iot_demo_api_service/controls/system/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web"
	basedata := e.Group(webApiPrefix)
	//basedata.Use(controls.AuthCheck)
	basedata.GET("/basedata/dicttype/detail/:id", apis.BaseDatacontroller.GetBaseTypeDetail)
	basedata.GET("/basedata/dicttype/list", apis.BaseDatacontroller.QueryBaseDataTypeList)
	basedata.POST("/basedata/dicttype/add", apis.BaseDatacontroller.AddTConfigDictDataType)
	basedata.POST("/basedata/dicttype/edit", apis.BaseDatacontroller.EditBaseDataType)
	basedata.POST("/basedata/dicttype/delete/:id", apis.BaseDatacontroller.DeleteBaseDataType)
}
