package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/message/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	app := e.Group(webApiPrefix)
	app.Use(controls.AuthCheck)

	//获取消息统计数据（红点数据）
	app.GET("/message/reddot", apis.Messagecontroller.GetMessageRedDot)
	//获取设备消息统计列表

	//清空家庭消息、系统消息、设备消息
	app.POST("/message/allDel/:type", apis.Messagecontroller.ClearMessage)
	app.POST("/message/del/:id", apis.Messagecontroller.DeleteMessage)

	//获取家庭消息、系统消息、设备消息列表
	app.GET("/message/grouplist/:type", apis.Messagecontroller.GetMessageGroupList)
	app.GET("/message/list/:type", apis.Messagecontroller.GetMessageList)
}
