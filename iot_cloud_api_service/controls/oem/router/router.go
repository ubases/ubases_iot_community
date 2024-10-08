package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/oem/apis"
	sysApis "cloud_platform/iot_cloud_api_service/controls/system/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/open/oem"
	admin := e.Group(webApiPrefix)
	admin.GET("/app/introduce/detail/:id/:lang", apis.OemAppIntroducecontroller.GetIntroduceByAppHtml)
	admin.GET("/app/introduce/template/detail/:id/:lang", apis.OemAppIntroducecontroller.GetIntroduceTtemplateByAppHtml)

	admin.POST("/app/buildFinishNotify", apis.OemAppBuildRecordcontroller.BuildFinishNotify)     //新版也在用，构建失败时调用，不传文件，后续会迁移完成并删除
	admin.POST("/app/buildFinishNotifyEx", apis.OemAppBuildRecordcontroller.BuildFinishNotifyEx) //新版打包服务
	admin.GET("/app/build/appIconUrl", apis.OemAppBuildRecordcontroller.GetBuildAppIconUrl)      //新版打包服务
	admin.GET("/app/build/package/qrcode", apis.OemAppBuildRecordcontroller.BuildPackageQrCode)
	//获取自定义app包下载页面
	admin.GET("/app/custom/package/qrcode", apis.OemAppCustomRecordControl.CustomPackageQrCode)
	admin.Use(controls.AuthCheck)

	admin.GET("/app", apis.OemAppcontroller.QueryList)
	admin.GET("/app/getByTenantId", apis.OemAppcontroller.QueryListByTenantId)
	admin.GET("/app/detail", apis.OemAppcontroller.GetOemAppDetail)
	admin.POST("/app", apis.OemAppcontroller.Add)
	admin.POST("/app/updateName", apis.OemAppcontroller.ChangeName)
	admin.POST("/app/updateTemplate", apis.OemAppcontroller.UpdateTemplate)
	admin.POST("/app/currentStep", apis.OemAppcontroller.UpdateCurrentStep)
	admin.DELETE("/app", apis.OemAppcontroller.DeleteOemApp)
	admin.POST("/app/delete", apis.OemAppcontroller.DeleteOemApp)
	admin.POST("/app/editTeamId", apis.OemAppcontroller.ChangeOemAppTeamId)
	admin.POST("/app/setStatus", apis.OemAppcontroller.ChangeOemAppTeamId)

	//自定义app版本记录管理
	admin.POST("/app/custom/version/add", apis.OemAppCustomRecordControl.CreateOemAppCustomRecord)
	admin.POST("/app/custom/version/edit", apis.OemAppCustomRecordControl.UpdateOemAppCustomRecord)
	admin.POST("/app/custom/version/set", apis.OemAppCustomRecordControl.SetOemAppCustomRecord)
	admin.POST("/app/custom/version/del", apis.OemAppCustomRecordControl.DeleteOemAppCustomRecord)
	admin.GET("/app/custom/version/detail", apis.OemAppCustomRecordControl.GetOemAppCustomRecord)
	admin.POST("/app/custom/version/list", apis.OemAppCustomRecordControl.GetOemAppCustomRecordList)

	admin.POST("/app/custom/version/editRemark", apis.OemAppCustomRecordControl.SetRemark)
	admin.POST("/app/custom/version/editLaunchMarkets", apis.OemAppCustomRecordControl.SetLaunchMarkets)
	admin.POST("/app/custom/version/editRemindMode", apis.OemAppCustomRecordControl.SetRemindMode)

	//获取自定义app包二维码链接
	admin.GET("/app/custom/qrCodeUrl", apis.OemAppcontroller.OemAppCustomPackageQrCodeUrl)

	admin.GET("/app/icon", apis.OemAppUiConfigcontroller.GetIcon)
	admin.POST("/app/icon", apis.OemAppUiConfigcontroller.SaveIcon)
	admin.GET("/app/iosLaunchScreen", apis.OemAppUiConfigcontroller.GetIosLaunchScreen)
	admin.POST("/app/iosLaunchScreen", apis.OemAppUiConfigcontroller.SaveIosLaunchScreen)
	admin.GET("/app/androidLaunchScreen", apis.OemAppUiConfigcontroller.GetAndroidLaunchScreen)
	admin.POST("/app/androidLaunchScreen", apis.OemAppUiConfigcontroller.SaveAndroidLaunchScreen)
	admin.GET("/app/themeColors", apis.OemAppUiConfigcontroller.GetThemeColors)
	admin.POST("/app/themeColors", apis.OemAppUiConfigcontroller.SaveThemeColors)
	admin.GET("/app/personalize", apis.OemAppUiConfigcontroller.GetPersonalize)
	admin.POST("/app/personalize", apis.OemAppUiConfigcontroller.SavePersonalize)
	admin.GET("/app/functionConfig", apis.OemAppUiConfigcontroller.GetFunctionConfig)
	admin.POST("/app/functionConfig", apis.OemAppUiConfigcontroller.SaveFunctionConfig)

	admin.GET("/app/functionConfig/voice", apis.OemAppUiConfigcontroller.GetFunctionConfigVoice)
	admin.POST("/app/functionConfig/voice", apis.OemAppUiConfigcontroller.SaveFunctionConfigVoice)
	//语音帮助文档直接调佣协议文档的方法
	admin.POST("/app/functionConfig/voice/doc", apis.OemAppIntroducecontroller.OemAppIntroduceAdd)
	admin.PUT("/app/functionConfig/voice/doc", apis.OemAppIntroducecontroller.OemAppIntroduceUpdate)
	admin.POST("/app/functionConfig/voice/doc/update", apis.OemAppIntroducecontroller.OemAppIntroduceUpdate)
	admin.GET("/app/functionConfig/voice/doc", apis.OemAppIntroducecontroller.OemAppVoiceIntroduceDetail)

	admin.GET("/app/functionConfig/thirdService", apis.OemAppUiConfigcontroller.GetFunctionConfigThird)
	admin.POST("/app/functionConfig/thirdService", apis.OemAppUiConfigcontroller.SaveFunctionConfigThird)

	admin.GET("/app/functionConfig/autoUpgrade", apis.OemAppUiConfigcontroller.GetFunctionConfigAutoUpgrade)
	admin.POST("/app/functionConfig/autoUpgrade", apis.OemAppUiConfigcontroller.SaveFunctionConfigAutoUpgrade)

	//默认房间
	admin.GET("/app/room/list", apis.OemAppUiConfigcontroller.GetRoomList)
	admin.POST("/app/room/save", apis.OemAppUiConfigcontroller.SaveRoom)
	admin.GET("/app/roomIcons/list", apis.OemAppUiConfigcontroller.GetRoomIconList)
	admin.POST("/app/roomIcons/save", apis.OemAppUiConfigcontroller.SaveRoomIconsList)
	admin.DELETE("/app/room/delete", apis.OemAppUiConfigcontroller.DeleteRoom)
	admin.POST("/app/room/delete", apis.OemAppUiConfigcontroller.DeleteRoom)
	admin.GET("/app/room/default", apis.OemAppUiConfigcontroller.RecoverDefaultRoom)

	admin.GET("/app/map", apis.OemAppcontroller.GetMap)
	admin.POST("/app/map", apis.OemAppcontroller.SaveMap)

	admin.GET("/app/menu", apis.OemAppUiConfigcontroller.GetButtonMenu)
	admin.POST("/app/menu/updateFontColor", apis.OemAppUiConfigcontroller.SaveButoomMenuFontColor)
	admin.POST("/app/menu", apis.OemAppUiConfigcontroller.AddButoomMenu)
	admin.PUT("/app/menu", apis.OemAppUiConfigcontroller.UpdateButoomMenu)
	admin.POST("/app/menu/update", apis.OemAppUiConfigcontroller.UpdateButoomMenu)
	admin.DELETE("/app/menu", apis.OemAppUiConfigcontroller.DeleteButoomMenu)
	admin.POST("/app/menu/delete", apis.OemAppUiConfigcontroller.DeleteButoomMenu)
	admin.GET("/app/menu/detail", apis.OemAppUiConfigcontroller.GetButoomMenuDetail)
	admin.POST("/app/ui/default", apis.OemAppUiConfigcontroller.RecoverDefault)
	//协议,隐私,关于我们
	admin.POST("/app/introduce", apis.OemAppIntroducecontroller.OemAppIntroduceAdd)
	admin.PUT("/app/introduce", apis.OemAppIntroducecontroller.OemAppIntroduceUpdate)
	admin.POST("/app/introduce/update", apis.OemAppIntroducecontroller.OemAppIntroduceUpdate)
	admin.POST("/app/introduce/copy", apis.OemAppIntroducecontroller.OemAppIntroduceCopy)
	admin.PUT("/app/introduce/statusEnable", apis.OemAppIntroducecontroller.OemAppIntroduceStatusEnable)
	admin.POST("/app/introduce/statusEnable", apis.OemAppIntroducecontroller.OemAppIntroduceStatusEnable)
	admin.GET("/app/introduce/detail", apis.OemAppIntroducecontroller.OemAppIntroduceDetail)
	admin.POST("/app/introduce/checkVersion", apis.OemAppIntroducecontroller.OemAppIntroduceCheckVersion)
	admin.GET("/app/introduce", apis.OemAppIntroducecontroller.OemAppIntroduceList)
	//admin.GET("/app/introduce/:contentType/:appId", apis.OemAppIntroducecontroller.OemAppIntroduceUrlList)
	admin.GET("/app/introduce/link", apis.OemAppIntroducecontroller.OemAppIntroduceUrlList)
	admin.GET("/app/introduce/template/link", apis.OemAppIntroducecontroller.OemAppIntroduceTemplateUrlList)

	//证书
	admin.POST("/app/iosCert", apis.OemAppCertcontroller.SaveIosCert)
	admin.GET("/app/iosCert", apis.OemAppCertcontroller.GetIosCert)
	admin.POST("/app/androidCert", apis.OemAppCertcontroller.SaveAndroidCert)
	admin.GET("/app/androidCert", apis.OemAppCertcontroller.GetAndroidCert)
	admin.GET("/app/downloadSignCert", apis.OemAppCertcontroller.DownloadSignCert)
	admin.POST("/app/iosPush", apis.OemAppCertcontroller.SaveIosPushCert)
	admin.GET("/app/iosPush", apis.OemAppCertcontroller.GetIosPushCert)
	admin.POST("/app/androidPush", apis.OemAppCertcontroller.SaveAndroidPushCert)
	admin.GET("/app/androidPush", apis.OemAppCertcontroller.GetAndroidPushCert)
	admin.GET("/app/androidCert/regenerate", apis.OemAppCertcontroller.Regenerate)

	//检查
	admin.GET("/app/check", apis.OemAppcontroller.OemAppCheck)
	//开始构建
	admin.POST("/app/build", apis.OemAppcontroller.OemAppBuild)
	admin.GET("/app/buildPackage", apis.OemAppcontroller.OemAppBuildPackage)
	admin.GET("/app/publish", apis.OemAppcontroller.OemAppPublish)
	admin.GET("/app/publishing", apis.OemAppcontroller.OemAppPublishing)
	admin.GET("/app/updateVersion", apis.OemAppcontroller.OemAppUpdateVersion)
	admin.GET("/app/build/qrCodeUrl", apis.OemAppcontroller.OemAppBuildPackageQrCodeUrl)
	admin.POST("/app/buildCancel", apis.OemAppcontroller.OemAppCancelBuild)
	// oem app版本列表
	admin.POST("/app/version/list", apis.OemAppVersionRecordcontroller.GetOemAppVersionRecordList)

	//oemapp文档
	admin.GET("/marketing/appDoc/getApp", apis.OemAppDoccontroller.GetApps)
	admin.GET("/marketing/appDoc/getPubLangs", apis.OemAppDoccontroller.GetPubLangs)
	admin.GET("/marketing/appDoc", apis.OemAppDoccontroller.DocList)
	admin.POST("/marketing/appDoc", apis.OemAppDoccontroller.CreateDoc)
	admin.PUT("/marketing/appDoc", apis.OemAppDoccontroller.UpdateDoc)
	admin.POST("/marketing/appDoc/update", apis.OemAppDoccontroller.UpdateDoc)
	admin.POST("/marketing/appDoc/del", apis.OemAppDoccontroller.DeleteDoc)
	admin.GET("/marketing/appDoc/getSupportLangs", apis.OemAppDoccontroller.GetDocSupportLangs)
	admin.GET("/marketing/appDoc/detail", apis.OemAppDoccontroller.DetailDoc)

	//oemapp文档目录
	admin.POST("/marketing/doc/directory", apis.OemAppDocDircontroller.CreateDir)
	admin.PUT("/marketing/doc/directory", apis.OemAppDocDircontroller.UpdateDir)
	admin.POST("/marketing/doc/directory/update", apis.OemAppDocDircontroller.UpdateDir)
	admin.DELETE("/marketing/doc/directory", apis.OemAppDocDircontroller.DeleteDir)
	admin.POST("/marketing/doc/directory/delete", apis.OemAppDocDircontroller.DeleteDir)
	admin.GET("/marketing/doc/directory/detail", apis.OemAppDocDircontroller.DetailDir)
	admin.GET("/marketing/doc/directory", apis.OemAppDocDircontroller.ListDir)

	//oemapp文档词条
	admin.POST("/marketing/doc/entry", apis.OemAppEntrycontroller.EntrySave)
	admin.GET("/marketing/doc/entry/detail", apis.OemAppEntrycontroller.EntryDetail)
	admin.POST("/marketing/doc/entry/seting", apis.OemAppEntrycontroller.EntrySetingSave)
	admin.GET("/marketing/doc/entry/seting", apis.OemAppEntrycontroller.EntrySetingDetail)
	admin.DELETE("/marketing/doc/entry", apis.OemAppEntrycontroller.EntryDelete)
	admin.POST("/marketing/doc/entry/delete", apis.OemAppEntrycontroller.EntryDelete)
	admin.POST("/marketing/doc/entry/list", apis.OemAppEntrycontroller.EntryList)

	//公版文档目录
	admin.GET("/marketing/public/doc/directory/detail", sysApis.SysAppDocDircontroller.DetailDir)
	admin.GET("/marketing/public/doc/directory", sysApis.SysAppDocDircontroller.ListDir)

	//公版文档词条
	admin.GET("/marketing/public/doc/entry/detail", sysApis.SysAppEntrycontroller.EntryDetail)
	admin.POST("/marketing/public/doc/entry/list", sysApis.SysAppEntrycontroller.EntryList)

	//公版帮助中心
	admin.POST("/marketing/public/helpCenter/copy", sysApis.SysAppHelpCentercontroller.CopyHelpCenter)
	admin.GET("/marketing/public/helpCenter/detail", sysApis.SysAppHelpCentercontroller.GetHelpCenter)
	//admin.POST("/marketing/public/helpCenter/list", sysApis.SysAppHelpCentercontroller.GetHelpCenterListForOpen)
	admin.POST("/marketing/public/helpCenter/list", apis.OemAppDoccontroller.GetPubDocList)

	//第三方登录相关
	e.GET("/.well-known/apple-app-site-association", apis.OemAppcontroller.GetThirdLoginJson)

	//反馈问题类型 /v1/platform/web/open/oem/feedback/problemType/list
	admin.POST("/feedback/problemType/list", apis.OemAppFeedbackProblemTypecontroller.List)
	admin.GET("/feedback/problemType/detail/:id", apis.OemAppFeedbackProblemTypecontroller.Get)
	admin.POST("/feedback/problemType/add", apis.OemAppFeedbackProblemTypecontroller.Add)
	admin.PUT("/feedback/problemType/edit", apis.OemAppFeedbackProblemTypecontroller.Update)
	admin.POST("/feedback/problemType/edit", apis.OemAppFeedbackProblemTypecontroller.Update)
	admin.DELETE("/feedback/problemType/del/:id", apis.OemAppFeedbackProblemTypecontroller.Delete)
	admin.POST("/feedback/problemType/del/:id", apis.OemAppFeedbackProblemTypecontroller.Delete)

	webApiPrefix = "v1/platform/web/app"
	admin = e.Group(webApiPrefix)
	admin.Use(controls.AuthCheck)

	//APP基础界面配置/v1/platform/web/app/basicUISetting/list
	admin.POST("/basicUISetting/list", apis.OemAppBasicUiSettingcontroller.List)
	admin.GET("/basicUISetting/detail", apis.OemAppBasicUiSettingcontroller.Get)
	admin.POST("basicUISetting/add", apis.OemAppBasicUiSettingcontroller.Add)
	admin.PUT("/basicUISetting/update", apis.OemAppBasicUiSettingcontroller.Update)
	admin.POST("/basicUISetting/update", apis.OemAppBasicUiSettingcontroller.Update)
	admin.DELETE("/basicUISetting/delete", apis.OemAppBasicUiSettingcontroller.Delete)
	admin.POST("/basicUISetting/delete", apis.OemAppBasicUiSettingcontroller.Delete)

	//APP模板 /v1/platform/web/app/appTemplate/list
	admin.POST("/appTemplate/list", apis.OemAppTemplatecontroller.List)
	admin.GET("/appTemplate/detail/:id", apis.OemAppTemplatecontroller.Get)
	admin.POST("appTemplate/add", apis.OemAppTemplatecontroller.Add)
	admin.POST("appTemplate/copy", apis.OemAppTemplatecontroller.Copy)
	admin.PUT("/appTemplate/edit", apis.OemAppTemplatecontroller.Update)
	admin.POST("/appTemplate/edit", apis.OemAppTemplatecontroller.Update)
	admin.DELETE("/appTemplate/delete/:id", apis.OemAppTemplatecontroller.Delete)
	admin.POST("/appTemplate/delete/:id", apis.OemAppTemplatecontroller.Delete)
	admin.POST("/appTemplate/setStatus", apis.OemAppTemplatecontroller.SetStatus)

	//APP模板-界面配置 /v1/platform/web/app/appTemplate/uiSettingList
	admin.POST("/appTemplate/uiSettingList", apis.OemAppTemplateUicontroller.List)
	admin.GET("/appTemplate/uiSettingDetail/:id", apis.OemAppTemplateUicontroller.Get)
	admin.POST("appTemplate/uiSettingAdd", apis.OemAppTemplateUicontroller.Add)
	admin.PUT("/appTemplate/uiSettingEdit", apis.OemAppTemplateUicontroller.Update)
	admin.POST("/appTemplate/uiSettingEdit", apis.OemAppTemplateUicontroller.Update)
	admin.PUT("/appTemplate/uiSettingEditContent", apis.OemAppTemplateUicontroller.Update)
	admin.POST("/appTemplate/uiSettingEditContent", apis.OemAppTemplateUicontroller.Update)
	admin.DELETE("/appTemplate/uiSettingDel/:id", apis.OemAppTemplateUicontroller.Delete)
	admin.POST("/appTemplate/uiSettingDel/:id", apis.OemAppTemplateUicontroller.Delete)

	//APP模板-功能配置 /v1/platform/web/app/appTemplate/functionList
	admin.POST("/appTemplate/functionList", apis.OemAppTemplateFunctioncontroller.List)
	admin.GET("/appTemplate/functionDetail/:id", apis.OemAppTemplateFunctioncontroller.Get)
	admin.POST("appTemplate/functionAdd", apis.OemAppTemplateFunctioncontroller.Add)
	admin.PUT("/appTemplate/functionEdit", apis.OemAppTemplateFunctioncontroller.Update)
	admin.POST("/appTemplate/functionEdit", apis.OemAppTemplateFunctioncontroller.Update)
	admin.DELETE("/appTemplate/functionDel/:id", apis.OemAppTemplateFunctioncontroller.Delete)
	admin.POST("/appTemplate/functionDel/:id", apis.OemAppTemplateFunctioncontroller.Delete)

	//APP模板-菜单配置 /v1/platform/web/app/appTemplate/menuList
	admin.POST("/appTemplate/menuList", apis.OemAppTemplateMenucontroller.List)
	admin.GET("/appTemplate/menuDetail/:id", apis.OemAppTemplateMenucontroller.Get)
	admin.POST("appTemplate/menuAdd", apis.OemAppTemplateMenucontroller.Add)
	admin.PUT("/appTemplate/menuEdit", apis.OemAppTemplateMenucontroller.Update)
	admin.POST("/appTemplate/menuEdit", apis.OemAppTemplateMenucontroller.Update)
	admin.DELETE("/appTemplate/menuDel/:id", apis.OemAppTemplateMenucontroller.Delete)
	admin.POST("/appTemplate/menuDel/:id", apis.OemAppTemplateMenucontroller.Delete)

	//APP模板-第三方配置 /v1/platform/web/app/appTemplate/thirdPartyList
	admin.POST("/appTemplate/thirdPartyList", apis.OemAppTemplateThirdPartycontroller.List)
	admin.GET("/appTemplate/thirdPartyDetail/:id", apis.OemAppTemplateThirdPartycontroller.Get)
	admin.POST("appTemplate/thirdPartyAdd", apis.OemAppTemplateThirdPartycontroller.Add)
	admin.PUT("/appTemplate/thirdPartyEdit", apis.OemAppTemplateThirdPartycontroller.Update)
	admin.POST("/appTemplate/thirdPartyEdit", apis.OemAppTemplateThirdPartycontroller.Update)
	admin.DELETE("/appTemplate/thirdPartyDel/:id", apis.OemAppTemplateThirdPartycontroller.Delete)
	admin.POST("/appTemplate/thirdPartyDel/:id", apis.OemAppTemplateThirdPartycontroller.Delete)

	//APP模板-皮肤配置 /v1/platform/web/app/appTemplate/skinList
	admin.POST("/appTemplate/skinList", apis.OemAppTemplateSkincontroller.List)
	admin.GET("/appTemplate/skinDetail/:id", apis.OemAppTemplateSkincontroller.Get)
	admin.POST("appTemplate/skinAdd", apis.OemAppTemplateSkincontroller.Add)
	admin.PUT("/appTemplate/skinEdit", apis.OemAppTemplateSkincontroller.Update)
	admin.POST("/appTemplate/skinEdit", apis.OemAppTemplateSkincontroller.Update)
	admin.DELETE("/appTemplate/skinDel/:id", apis.OemAppTemplateSkincontroller.Delete)
	admin.POST("/appTemplate/skinDel/:id", apis.OemAppTemplateSkincontroller.Delete)

	//公版app版本管理 /v1/platform/web/app/publicVersion/add
	admin.POST("/publicVersion/list", apis.PublicAppVersioncontroller.List)
	admin.GET("/publicVersion/detail/:id", apis.PublicAppVersioncontroller.Get)
	admin.POST("publicVersion/add", apis.PublicAppVersioncontroller.Add)
	admin.PUT("/publicVersion/edit", apis.PublicAppVersioncontroller.Update)
	admin.POST("/publicVersion/edit", apis.PublicAppVersioncontroller.Update)
	admin.DELETE("/publicVersion/del/:id", apis.PublicAppVersioncontroller.Delete)
	admin.POST("/publicVersion/del/:id", apis.PublicAppVersioncontroller.Delete)
	admin.POST("/publicVersion/setStatus", apis.PublicAppVersioncontroller.SetStatus)

	//APP辅助上架 /v1/platform/web/app/assistRelease/list
	admin.POST("/assistRelease/list", apis.OemAppAssistReleasecontroller.List)
	admin.GET("/assistRelease/detail/:id", apis.OemAppAssistReleasecontroller.Get)
	admin.POST("assistRelease/add", apis.OemAppAssistReleasecontroller.Add)
	admin.PUT("/assistRelease/edit", apis.OemAppAssistReleasecontroller.Update)
	admin.POST("/assistRelease/edit", apis.OemAppAssistReleasecontroller.Update)
	admin.DELETE("/assistRelease/del/:id", apis.OemAppAssistReleasecontroller.Delete)
	admin.POST("/assistRelease/del/:id", apis.OemAppAssistReleasecontroller.Delete)
	admin.POST("/assistRelease/setStatus", apis.OemAppAssistReleasecontroller.SetStatus)

	//APP辅助上架 查询开发者信息 v1/platform/web/app/developer/appList
	admin.GET("/developer/appList", apis.OemAppAssistReleasecontroller.DeveloperAppList)

	//APP调试管理
	admin.POST("/appDebugger/list", apis.OemAppDebuggercontroller.List)
	admin.GET("/appDebugger/detail/:id", apis.OemAppDebuggercontroller.Get)
	admin.POST("appDebugger/add", apis.OemAppDebuggercontroller.Add)
	admin.POST("/appDebugger/edit", apis.OemAppDebuggercontroller.Update)
	admin.POST("/appDebugger/delete/:id", apis.OemAppDebuggercontroller.Delete)

}
