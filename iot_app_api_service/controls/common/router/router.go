package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/common/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app/common"
	r := e.Group(webApiPrefix)
	r.Use(controls.SetParams)
	//区域列表
	r.GET("/regionList", apis.Commoncontroller.RegionList)
	//区域详情
	r.GET("/region/:id", apis.Commoncontroller.RegionInfo)
	//APP的语言
	r.GET("/customLang/list", apis.Commoncontroller.CustomResourceExport)
	//r.GET("/customLang/panel", apis.Commoncontroller.PanelCustomResourceExport)
	//面板的语言
	r.GET("/customLang/panel", apis.Commoncontroller.QueryPanelResourceV2)
	r.GET("/panelLang/list", apis.Commoncontroller.QueryPanelResourceV2)
	//字典列表
	r.POST("/dic/list", apis.Commoncontroller.DictList)
	//房间配置（默认房间、房间图标）
	r.GET("/room/config/:code", apis.Commoncontroller.RoomConfigList)
	//天气接口
	r.GET("/weather", apis.Commoncontroller.GetWeather)

	//闪屏推送
	r.GET("/flashscreen", apis.Commoncontroller.GetFlashScreen)

	//APP信息、APP版本信息
	r.GET("/appInfo", apis.Appcontroller.GetAppDetailByApp)

	//APP临时皮肤
	r.GET("/checkSkin", apis.AssistReleasecontroller.CheckSkin)

	//上传图片
	r.POST("/uploadPic", apis.Filecontroller.UploadFile)

	r.Use(controls.AuthCheck)

	//第三方语音
	r.GET("/voice/config/:voiceCode", apis.Commoncontroller.VoiceService)

	//获取当前时间
	r.GET("/nowtime", apis.Commoncontroller.GetNowTime)

}
