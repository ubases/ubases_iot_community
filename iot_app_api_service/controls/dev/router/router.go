package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/dev/apis"
	apis2 "cloud_platform/iot_app_api_service/controls/intelligence/apis"
	productService "cloud_platform/iot_app_api_service/controls/product/services"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotnatsjs"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	r := e.Group(webApiPrefix)

	//设备操作失败日志上报1.0.5
	r.POST("/dev/operation/failLog", apis.FailLogcontroller.ReportOperationFailLog)
	r.Use(controls.AuthCheck)
	r.Use(controls.SetParams)
	r.Use(iotgin.AppLogger(iotnatsjs.GetJsClientPub()))
	r.POST("/dev/operation/failLogEx", apis.FailLogcontroller.ReportOperationFailLogEx)

	r.GET("/dev/deviceInfo/:devId", apis.Devcontroller.DeviceInfo)
	r.POST("/dev/removeDev", apis.Devcontroller.RemoveDev)
	r.POST("/dev/removeRoomDev", apis.Devcontroller.RemoveRoomDev)
	r.POST("/dev/update/:devId", apis.Devcontroller.UpdateDev)
	r.POST("/dev/addDev", apis.Devcontroller.AddDev)
	r.GET("/dev/functions/:devId", apis2.SceneIntelligencecontroller.GetProductFunctions)
	r.GET("/dev/functionsV2/:devId", apis2.SceneIntelligencecontroller.GetAppointmentFunctions)

	//配网检查接口（配网缓存信息有product中创建，这里转调到productService调用）
	r.GET("/dev/check/:did", func(c *gin.Context) {
		devId := c.Param("did")
		token := c.Query("token")
		resInt, _, err := productService.CheckNetworkToken(token, devId)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		//保持与原接口返回格式
		iotgin.ResSuccess(c, resInt)
	})
	//扩展检查接口，增加配网错误信息返回
	r.GET("/dev/checkMsg/:did", func(c *gin.Context) {
		devId := c.Param("did")
		token := c.Query("token")
		resInt, resMsg, err := productService.CheckNetworkToken(token, devId)
		if err != nil {
			iotgin.ResErrCli(c, err)
			return
		}
		//保持与原接口返回格式
		iotgin.ResSuccess(c, map[string]interface{}{"code": resInt, "errMsg": resMsg})
	})

	r.POST("/dev/timer/add", apis.Timercontroller.AddTimer)
	r.POST("/dev/timer/update/:id", apis.Timercontroller.UpdateTimer)
	r.POST("/dev/timer/delete/:id", apis.Timercontroller.RemoveTimer)
	r.POST("/dev/timer/close/:id", apis.Timercontroller.DisabledTimer)
	r.GET("/dev/timer/info/:id", apis.Timercontroller.TimerInfo)
	r.POST("/dev/timer/list", apis.Timercontroller.TimerList)
	r.POST("/dev/timer/open/:id", apis.Timercontroller.EnabledTimer)

	r.POST("/dev/countdown/add", apis.Countdowncontroller.AddCountdown)
	r.GET("/dev/countdown/info/:devId", apis.Countdowncontroller.CountdownInfo)
	r.POST("/dev/countdown/close/:devId", apis.Countdowncontroller.DisabledCountdown)
	r.POST("/dev/countdown/open/:devId", apis.Countdowncontroller.EnabledCountdown)
	r.POST("/dev/countdown/delete/:devId", apis.Countdowncontroller.RemoveCountdown)

	r.GET("dev/runRecord/daysGroupDetail", apis.DeviceReportsapis.GetDaysDetail)
	r.GET("dev/runRecord/daysGroupCount", apis.DeviceReportsapis.GetDaysHourCount)
	r.GET("dev/runRecord/clearDetail", apis.DeviceReportsapis.ClearDetail)
	r.GET("/dev/ota/checkVersion", apis.Otacontroller.CheckOtaVersion)
	r.GET("/dev/ota/checkUpgradeList", apis.Otacontroller.CheckOtaUpgradeList)

	r.GET("/dev/shareDeviceList/:homeId", apis.ShareDevicecontroller.ShareDeviceList)
	r.GET("/dev/shareUserList/:devId", apis.ShareDevicecontroller.ShareUserList)
	r.POST("/dev/addShared", apis.ShareDevicecontroller.AddShared)
	r.GET("/dev/receiveSharedList", apis.ShareDevicecontroller.ReceiveSharedList)
	r.POST("/dev/cancelShare", apis.ShareDevicecontroller.CancelShare)
	r.POST("/dev/cancelReceiveShared", apis.ShareDevicecontroller.CancelReceiveShared)
	r.POST("/dev/receiveShare/:id", apis.ShareDevicecontroller.ReceiveShare)

	//小程序生成分享码
	r.POST("/dev/miniProgram/genShareCode", apis.ShareDevicecontroller.GenShareCode)
	r.POST("/dev/miniProgram/receiveShare", apis.ShareDevicecontroller.ReceiveShareByCode)

	r.GET("/dev/group/info/:groupId", apis.DeviceGroupcontroller.DevGroupInfo)
	r.GET("/dev/group/devListByProductKey", apis.DeviceGroupcontroller.DevListByProductKey)
	r.GET("/dev/group/devList/:groupId", apis.DeviceGroupcontroller.DevGroupDevList)
	r.POST("/dev/group/upsert", apis.DeviceGroupcontroller.UpsertGroup)
	r.POST("/dev/group/remove/:groupId", apis.DeviceGroupcontroller.RemoveGroup)
	r.POST("/dev/group/execute", apis.DeviceGroupcontroller.Execute)
	r.GET("/dev/group/tsl/:groupId", apis.DeviceGroupcontroller.DevGroupTsl)
	r.POST("/dev/group/executeSwitch", apis.DeviceGroupcontroller.ExecuteSwitch)

	//设置自动升级授权1.0.5
	//POST /v1/platform/app/dev/ota/setAutoUpgrade
	r.POST("/dev/ota/setAutoUpgrade", apis.Otacontroller.SetAutoUpgradeSwitch)

	//更新客户自定义功能描述，用于面板中对功能的名称进行自定义
	r.POST("/dev/funtionDesc/update", apis.DeviceFunctionSetcontroller.UpdateFunction)
	r.POST("/dev/funtionDesc/batchUpdate", apis.DeviceFunctionSetcontroller.BatchUpdateFunction)
	r.GET("/dev/functionDesc/list", apis.DeviceFunctionSetcontroller.GetDeviceFunctionSetList)
}
