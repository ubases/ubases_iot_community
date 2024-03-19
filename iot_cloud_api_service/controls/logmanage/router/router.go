package router

import (
	"cloud_platform/iot_cloud_api_service/controls/logmanage/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/log"
	admin := e.Group(webApiPrefix)
	admin.POST("/app/userList", apis.GetAppLogUserList)
	admin.POST("/app/recordsList", apis.GetAppLogRecordsList)
}
