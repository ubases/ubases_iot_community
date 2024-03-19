package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/common/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/common"
	r := e.Group(webApiPrefix)

	r.Use(controls.AuthCheck)
	r.POST("/fileStore/getLocalFile", apis.Commoncontroller.GetLocalFile)
	r.POST("/fileStore/uploadLocalFile", apis.Commoncontroller.UploadLocalFile)
	r.POST("/fileStore/uploadOssFile", apis.Filecontroller.UploadFile)
	r.POST("/fileStore/uploadOssImage", apis.Filecontroller.UploadImage)
	r.GET("/regionList", apis.Commoncontroller.RegionList)
}
