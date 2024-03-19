package router

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/user/apis"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotnats/jetstream"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/app"
	r := e.Group(webApiPrefix)
	r.Use(controls.SetParams)
	r.Use(iotgin.AppLogger(jetstream.GetJsPublisherMgr()))
	r.POST("/user/refreshToken", apis.Usercontroller.RefreshToken)
	r.POST("/user/login", apis.Usercontroller.Login)
	r.POST("/user/register", apis.Usercontroller.Register)
	r.POST("/user/registerex", apis.Usercontroller.RegisterEx)
	r.POST("/user/sendEmail", apis.Usercontroller.SendEmail)
	r.POST("/user/sendSms", apis.Usercontroller.SendSms)
	r.POST("/user/checkCode", apis.Usercontroller.CheckCode)
	r.POST("/user/checkAccount", apis.Usercontroller.CheckAccount)
	r.POST("/user/forgetPassword", apis.Usercontroller.ForgetPassword)
	r.POST("/user/channelAuth", apis.Usercontroller.ChannelAuth)
	r.POST("/user/channelBind", apis.Usercontroller.ChannelBind)
	r.POST("/user/getAppId", apis.Usercontroller.GetAppId)
	r.GET("/user/convertMd5", apis.Usercontroller.ConvertToMd5)
	//推送注册
	r.POST("/push/register", apis.Usercontroller.PushRegister)
	r.Any("/user/logout", apis.Usercontroller.Logout)

	r.Use(controls.AuthCheck)
	r.POST("/user/updateUser", apis.Usercontroller.UpdateUser)
	r.POST("/user/setPassword", apis.Usercontroller.SetUserPassword)
	//r.GET("/user/getAddress", apis.Usercontroller.GetAddress)
	//r.POST("/user/setDefaultHome", apis.Usercontroller.SetDefaultHome) todo 合并到修改用户的方法
	r.GET("/user/homeList", apis.Usercontroller.HomeList)
	//用户信息可以缓存，以提高性能
	r.GET("/user/detail" /* cache.CacheByRequestURI(cached.RedisStore, 15*time.Second),*/, apis.Usercontroller.GetUserDetail)
	r.POST("/user/cancelAccount", apis.Usercontroller.CancelAccount)
	r.POST("/user/cancelThirdPartyAccount", apis.Usercontroller.CancelThirdPartyAccount)

	r.POST("/user/addChannelBind", apis.Usercontroller.AddChannelBind)
	r.POST("/user/unbindChannel", apis.Usercontroller.UnbindChannel)
	r.POST("/user/accountBind", apis.Usercontroller.AccountBind)

	r.GET("/user/getFunctionConfig", apis.Usercontroller.GetFunctionConfigVoice)
	r.GET("/functionConfig", apis.Usercontroller.GetFunctionConfig)

	r.GET("/home/details/:id", apis.Homecontroller.Details)
	r.POST("/home/add", apis.Homecontroller.Add)
	r.POST("/home/update/:id", apis.Homecontroller.Update)
	r.POST("/home/delete/:id", apis.Homecontroller.Delete)
	r.POST("/home/sendInvitationCode", apis.Homecontroller.SendInvitationCode)
	r.POST("/home/joinHome", apis.Homecontroller.JoinHome)
	r.POST("/home/setRole", apis.Homecontroller.SetRole)
	r.POST("/home/removeMembers", apis.Homecontroller.RemoveMembers)
	r.POST("/home/transferOwnership", apis.Homecontroller.TransferOwnership)
	r.POST("/home/quit", apis.Homecontroller.Quit)
	r.GET("/home/roomList/:homeId", apis.Homecontroller.RoomList)
	r.GET("/home/deviceList/:homeId", apis.Homecontroller.DeviceList)
	r.GET("/user/deviceList", apis.Homecontroller.UserDeviceList)
	r.GET("/home/serverAlloc/:homeId", apis.Homecontroller.ServerAlloc)
	r.POST("/home/addDev", apis.Homecontroller.AddDev)
	r.POST("/home/setDevSort", apis.Homecontroller.SetDevSort)

	r.POST("/room/add", apis.Roomcontroller.Add)
	r.POST("/room/setSort", apis.Roomcontroller.SetSort)
	r.POST("/room/setDevSort", apis.Roomcontroller.SetDevSort) //TODO app报错增加临时路由，需要前端修改地址
	r.GET("/room/details/:homeId/:roomId", apis.Roomcontroller.Details)
	r.POST("/room/delete", apis.Roomcontroller.Delete)
	r.POST("/room/update", apis.Roomcontroller.Update)

	// 有奖征集
	r.POST("/marketing/prizecollect/add", apis.PrizeCollectcontroller.Add)

	// 区域管理
	r.GET("/area/list", apis.Areacontroller.QueryList)
	r.POST("/area/add", apis.Areacontroller.Add)
	r.PUT("/area/edit", apis.Areacontroller.Edit)
	r.DELETE("/area/delete", apis.Areacontroller.Delete)
	r.POST("/area/delete", apis.Areacontroller.Delete)
	r.GET("/area/detail/:id", apis.Areacontroller.QueryDetail)

	//获取区域的树型数据
	r.GET("/area/treeData/:parentId/:showChild", apis.Areacontroller.GetAreas)

	//切换语言
	r.GET("/changeLang", apis.Usercontroller.ChangeLang)
}
