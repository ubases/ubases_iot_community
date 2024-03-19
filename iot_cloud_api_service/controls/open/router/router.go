package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/apis"
	apis2 "cloud_platform/iot_cloud_api_service/controls/product/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/open"
	admin := e.Group(webApiPrefix)

	admin.POST("/register", apis.OpenUsercontroller.Register)
	admin.POST("/forgetPassword", apis.OpenUsercontroller.ForgetPassword)
	admin.POST("/login", apis.OpenUsercontroller.Login)
	admin.GET("/logout", apis.OpenUsercontroller.Logout)
	admin.GET("/getVerificationCode", apis.OpenUsercontroller.GetVerificationCode)
	//通过用户名查询用户是否存在
	admin.GET("/getVerificationCodeForExists", apis.OpenUsercontroller.GetVerificationCodeForExists)

	admin.Use(controls.AuthCheck)
	admin.POST("/refreshToken", apis.OpenUsercontroller.RefreshToken)
	admin.GET("/user/profile", apis.OpenUsercontroller.GetUserProfile)
	admin.GET("/user/getRouters", apis.OpenUsercontroller.GetRouters)
	admin.PUT("/user/updatePwd", apis.OpenUsercontroller.UpdateOpenUserPwd)
	admin.POST("/user/updatePwd", apis.OpenUsercontroller.UpdateOpenUserPwd)

	admin.GET("/baseinfo", apis.OpenCompanycontroller.GetBaseInfo)
	admin.PUT("/baseinfo", apis.OpenCompanycontroller.UpdateBaseInfo)
	admin.POST("/baseinfo", apis.OpenCompanycontroller.UpdateBaseInfo)
	admin.GET("/getCompanyInfo", apis.OpenCompanycontroller.GetCompanyInfo)

	admin.POST("/company/connect", apis.OpenCompanyConnectcontroller.AddConnect)
	admin.PUT("/company/connect", apis.OpenCompanyConnectcontroller.UpdateConnect)
	admin.POST("/company/connect/update", apis.OpenCompanyConnectcontroller.UpdateConnect)
	admin.DELETE("/company/connect", apis.OpenCompanyConnectcontroller.DeleteConnect)
	admin.POST("/company/connect/delete", apis.OpenCompanyConnectcontroller.DeleteConnect)

	admin.GET("/company/auth", apis.OpenCompanycontroller.GetCompanyAuth)

	admin.POST("/role", apis.OpenRolecontroller.AddRole)
	admin.PUT("/role", apis.OpenRolecontroller.EditRole)
	admin.POST("/role/update", apis.OpenRolecontroller.EditRole)
	admin.DELETE("/role", apis.OpenRolecontroller.DeleteRole)
	admin.POST("/role/delete", apis.OpenRolecontroller.DeleteRole)
	admin.GET("/role", apis.OpenRolecontroller.RoleDetail)
	admin.GET("/role/list", apis.OpenRolecontroller.RoleList)
	admin.POST("/role/user", apis.OpenRolecontroller.RoleSetUser)

	admin.POST("/company/auth", apis.OpenCompanycontroller.CompanyAuth)
	admin.POST("/company/changeName", apis.OpenCompanycontroller.CompanyChangeName)
	admin.POST("/auth/add", apis.OpenUserCompanycontroller.UserCompanyAuth)
	admin.POST("/auth/remark", apis.OpenUserCompanycontroller.UserCompanyUpdateReamk)
	admin.DELETE("/auth", apis.OpenUserCompanycontroller.UserCompanyDelete)
	admin.POST("/auth/delete", apis.OpenUserCompanycontroller.UserCompanyDelete)
	admin.GET("/auth/list", apis.OpenUserCompanycontroller.UserCompanyAuthList)
	admin.POST("/chooseCompany", apis.OpenUsercontroller.ChangeTenant)
	//支持语言列表
	admin.GET("/lang/langTypeList", apis.OpenUsercontroller.LangOpenTypeList)

	// admin.POST("/refreshToken", apis.Usercontroller.RefreshToken)
	// admin.POST("/login", apis.Usercontroller.Login)
	// admin.Use(controls.AuthCheck)
	// admin.POST("/logout", apis.Usercontroller.Logout)
	// admin.GET("/auth/userList", apis.Usercontroller.QueryUserList)
	// admin.POST("/auth/addUser", apis.Usercontroller.AddUser)
	// admin.PUT("/auth/editUser", apis.Usercontroller.EditUser)
	// admin.PUT("/auth/editUserCenter", apis.Usercontroller.EditUserCenter)

	//=====================================================
	//OTA升级选择固件数据
	admin.GET("/product/firmwareList", apis.ProductFirmwarecontroller.QueryDropDownList)
	admin.GET("/product/firmwareVersionList", apis.ProductFirmwarecontroller.QueryDropDownVersionList)

	//固件管理
	admin.POST("/firmware/list", apis.Firmwarecontroller.QueryList)
	admin.POST("/firmware/add", apis.Firmwarecontroller.Add)
	admin.PUT("/firmware/edit", apis.Firmwarecontroller.Edit)
	admin.POST("/firmware/edit", apis.Firmwarecontroller.Edit)
	admin.POST("/firmware/setStatus", apis.Firmwarecontroller.SetStatus)
	admin.DELETE("/firmware/delete", apis.Firmwarecontroller.Delete)
	admin.POST("/firmware/delete", apis.Firmwarecontroller.Delete)
	admin.GET("/firmware/detail/:id", apis.Firmwarecontroller.QueryDetail)

	//固件版本管理
	admin.POST("/firmwareVersion/list", apis.FirmwareVersioncontroller.QueryList)
	admin.POST("/firmwareVersion/enableList", apis.FirmwareVersioncontroller.QueryEnableList)
	admin.POST("/firmwareVersion/add", apis.FirmwareVersioncontroller.Add)
	admin.PUT("/firmwareVersion/edit", apis.FirmwareVersioncontroller.Edit)
	admin.PUT("/firmwareVersion/setStatus", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, -1)
	})
	admin.POST("/firmwareVersion/edit", apis.FirmwareVersioncontroller.Edit)
	admin.POST("/firmwareVersion/setStatus", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, -1)
	})

	admin.PUT("/firmwareVersion/onShelf", apis.FirmwareVersioncontroller.OnShelf)
	admin.PUT("/firmwareVersion/unShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 2)
	})
	admin.POST("/firmwareVersion/onShelf", apis.FirmwareVersioncontroller.OnShelf)
	admin.POST("/firmwareVersion/unShelf", func(context *gin.Context) {
		apis.FirmwareVersioncontroller.SetStatus(context, 2)
	})
	admin.DELETE("/firmwareVersion/delete", apis.FirmwareVersioncontroller.Delete)
	admin.POST("/firmwareVersion/delete", apis.FirmwareVersioncontroller.Delete)
	admin.GET("/firmwareVersion/detail/:id", apis.FirmwareVersioncontroller.QueryDetail)

	//我的产品列表
	admin.POST("/product/list", apis.Productcontroller.QueryList)
	admin.GET("/product/langList", apis.Productcontroller.QueryLangList)
	//产品详情
	admin.GET("/product/baseDetail/:id", apis.Productcontroller.QueryDetail)
	admin.GET("/product/detail/:id", apis.Productcontroller.QueryAllDetail)
	//查询产品分类列表
	admin.GET("/productType/get", apis2.ProductTypecontroller.GetTypeAndProductList)
	//查询开发方案列表（云管理平台接口）
	admin.GET("/product/getBaseProduct", apis2.Productcontroller.GetProductList)
	//删除产品
	admin.DELETE("/product/delete", apis.Productcontroller.Delete)
	admin.POST("/product/delete", apis.Productcontroller.Delete)
	//查询配网引导详情 （GetDefaultNetworkGuides）
	admin.GET("/product/getNetworkGuide", apis.Productcontroller.QueryProductNetworkGuide)
	admin.GET("/product/getNetworkGuideLang", apis.Productcontroller.QueryProductNetworkGuideLang)
	admin.GET("/product/getDefaultNetworkGuide", apis.Productcontroller.QueryProductDefaultNetworkGuide)
	//保存配网引导步骤 (修改为save)
	admin.POST("/product/saveNetworkGuide", apis.Productcontroller.SaveProductNetworkGuide)
	admin.POST("/product/changeNetworkGuide", apis.Productcontroller.SetProductNetworkGuideType)

	//第一步：
	//	创建产品
	admin.POST("/product/create", apis.Productcontroller.Add)
	admin.POST("/product/save", apis.Productcontroller.Add)
	//  编辑产品
	admin.PUT("/product/edit", apis.Productcontroller.Edit)
	admin.POST("/product/edit", apis.Productcontroller.Edit)
	admin.POST("/product/editPanelInfo", apis.Productcontroller.EditPanelInfo)
	//第二步：
	//	查询标准功能/自定义功能
	admin.GET("/product/funcList", apis.Productcontroller.QueryProductThingModel)
	admin.GET("/product/langFuncList", apis.Productcontroller.QueryProductThingModelAndLang)
	//admin.GET("/product/controlPanelLang", apis.Productcontroller.QueryControlPanelLang)
	admin.GET("/controlPanel/langList", apis.Productcontroller.ControlPanelCustomResource)
	//	添加标准 - 标准功能列表 (未定义）
	admin.GET("/product/standardFuncList", apis.Productcontroller.QueryStandardThingModel)
	admin.GET("/product/functions", apis.Productcontroller.GetTaskOrWhereByProduct)
	//	添加标准 - 提交
	admin.POST("/product/addStandardFunc", apis.Productcontroller.AddStandThingModel)
	//	功能 - 重置
	admin.GET("/product/resetStandardFunc", apis.Productcontroller.ResetStandThingFunc)
	admin.POST("/product/setSceneFunc", apis.Productcontroller.SetThingsModelSceneFunc)
	//  预约功能设置
	admin.POST("/product/appointmentFunc/setLevel", apis.Productcontroller.SetFuncLevel)
	admin.POST("/product/appointmentFunc/select", apis.Productcontroller.SetAppointmentFunc)
	admin.POST("/product/appointmentFunc/moveUp", apis.Productcontroller.SetFuncMoveUp)
	admin.POST("/product/appointmentFunc/moveDown", apis.Productcontroller.SetFuncMoveDown)
	admin.GET("/product/appointmentFunc/list", apis.Productcontroller.QueryAppointmentFuncList)
	//  新增自定义功能
	admin.POST("/product/addFunc", apis.Productcontroller.AddThingModel)
	//  编辑自定义功能 + 编辑标准功能
	admin.PUT("/product/editFunc", apis.Productcontroller.EditThingModel)
	admin.POST("/product/editFunc", apis.Productcontroller.EditThingModel)
	//  编辑自定义功能 + 编辑标准功能
	admin.DELETE("/product/deleteFunc", apis.Productcontroller.DeleteThingModel)
	admin.POST("/product/deleteFunc", apis.Productcontroller.DeleteThingModel)
	//第三步：
	// 查询可选择的模组
	admin.GET("/product/queryModules", apis.Productcontroller.QueryModuleList)
	// 查询可选择的自定义固件
	admin.GET("/product/queryCustomFirmware", apis.Firmwarecontroller.QueryCustomFirmwareList)
	// 选择模组固件
	admin.POST("/product/selectModule", apis.Productcontroller.SaveProductModule)
	admin.POST("/product/selectCustomerFirmware", apis.Productcontroller.SaveProductFirmware)
	admin.POST("/product/removeCustomerFirmware", apis.Productcontroller.RemoveProductFirmware)
	// 获取固件类型
	admin.GET("/product/firmwareTypeList", apis.Productcontroller.QueryProductFirmwareType)
	// 更换固件版本
	admin.POST("/product/changeVersionSubmit", apis.Productcontroller.ChangeVersionSubmit)
	// MCU SDK下载
	admin.GET("/product/sdk/download", apis.Productcontroller.DownloadMcuSdk)
	//第四步：
	// 查询面板列表
	admin.GET("/product/controlPanelList", apis.Productcontroller.QueryControlPanelList)
	// 选择面板
	admin.POST("/product/selectControlPanel", apis.Productcontroller.SaveProductControlPanel)
	//第五步：
	//开发完成查询
	admin.GET("/product/completeDevelopDetailed", apis.Productcontroller.QueryCompleteDevelopDetail)
	//	开发完成
	admin.POST("/product/completeDevelop", func(context *gin.Context) {
		apis.Productcontroller.SetShelf(context, 1)
	})
	//	返回开发
	admin.POST("/product/returnDevelop", func(context *gin.Context) {
		apis.Productcontroller.SetShelf(context, 2)
	})
	//	上传测试报告
	admin.POST("/product/uploadTestReport", apis.Productcontroller.UploadTestReport)
	admin.GET("/product/getTestReportFile", apis.Productcontroller.GetTestReport)
	//测试用例模板下载

	admin.GET("/firmware/changeVersionList", apis2.FirmwareVersioncontroller.QueryChangeVersionList)
	admin.GET("/firmware/changeCustomVersionList", apis.FirmwareVersioncontroller.QueryCustomEnableList)

	//OTA升级相关接口
	admin.POST("/firmware/otaList", apis.OtaPkgcontroller.QueryList)
	//固件OTA详细信息查询
	admin.GET("/firmware/otaDetail/:id", apis.OtaPkgcontroller.QueryDetail)
	//固件OTA新增
	admin.POST("/firmware/otaAdd", apis.OtaPkgcontroller.Add)
	//固件OTA修改
	admin.PUT("/firmware/otaEdit", apis.OtaPkgcontroller.Edit)
	admin.POST("/firmware/otaEdit", apis.OtaPkgcontroller.Edit)
	//固件OTA发布
	admin.POST("/firmware/otaRelease", apis.OtaPkgcontroller.OtaPublish)
	//固件OTA暂停
	admin.POST("/firmware/otaStop", apis.OtaPkgcontroller.OtaPublishStop)
	//固件OTA恢复发布
	admin.POST("/firmware/otaRecoveryRelease", apis.OtaPkgcontroller.OtaRecoveryPublish)
	//固件OTA发布记录列表查询
	admin.POST("/firmware/releaseRecord", apis.OtaPkgcontroller.OtaPublishQueryList)
	//固件OTA删除
	admin.DELETE("/firmware/otaDelete", apis.OtaPkgcontroller.Delete)
	admin.POST("/firmware/otaDelete", apis.OtaPkgcontroller.Delete)
	//设备OTA升级记录
	admin.GET("/firmware/otaResult/:publishId", apis.OtaPkgcontroller.OtaResultList)
	//指定版本列表
	admin.GET("/firmware/upgradableVersions", apis.OtaPkgcontroller.QueryOtaVersions)
	admin.GET("/firmware/upgradableAreas", apis.OtaPkgcontroller.QueryOtaAreas)

	//APP用户管理
	admin.GET("/appUser/userList", apis.OpenUsercontroller.QueryUserList)
	//APP用户绑定设备
	admin.GET("/appUser/deviceList", apis.OpenUsercontroller.QueryUserDeviceList)

	//产品帮助中心配置相关接口
	admin.POST("/marketing/product/help/conf/edit", apis.ProHelpConfController.EditProductHelpConf)
	admin.POST("/marketing/product/help/conf/set", apis.ProHelpConfController.SetProductHelpConf)
	admin.GET("/marketing/product/help/conf/detail", apis.ProHelpConfController.GetProductHelpConf)
	admin.POST("/marketing/product/help/conf/list", apis.ProHelpConfController.GetProductHelpConfList)

	//产品帮助中心文档相关接口
	admin.POST("/marketing/product/help/doc/add", apis.ProHelpDocController.AddProductHelpDoc)
	admin.POST("/marketing/product/help/doc/edit", apis.ProHelpDocController.EditProductHelpDoc)
	admin.POST("/marketing/product/help/doc/set", apis.ProHelpDocController.SetProductHelpDoc)
	admin.POST("/marketing/product/help/doc/del", apis.ProHelpDocController.DeleteProductHelpDoc)
	admin.GET("/marketing/product/help/doc/detail", apis.ProHelpDocController.GetProductHelpDoc)
	admin.POST("/marketing/product/help/doc/list", apis.ProHelpDocController.GetProductHelpDocList)

	//闪屏
	admin.POST("/marketing/flashscreen/add", apis.OemAppFlashScreenControl.CreateFlashScreen)
	admin.POST("/marketing/flashscreen/edit", apis.OemAppFlashScreenControl.UpdateFlashScreen)
	admin.POST("/marketing/flashscreen/set", apis.OemAppFlashScreenControl.SetFlashScreen)
	admin.GET("/marketing/flashscreen/detail", apis.OemAppFlashScreenControl.GetFlashScreen)
	admin.POST("/marketing/flashscreen/list", apis.OemAppFlashScreenControl.GetFlashScreenList)

	//产品语控配置相关
	admin.POST("/product/voice/save", apis.ProductVoicecontroller.Save)
	admin.POST("/product/voice/list", apis.ProductVoicecontroller.GetList)
	admin.POST("/product/voice/detail", apis.ProductVoicecontroller.GetDetail)
	admin.POST("/product/voice/publish", apis.ProductVoicecontroller.Publish)
	admin.GET("/product/voice/getDoc", apis.ProductVoicecontroller.GetVoiceDoc)
	admin.GET("/product/voice/publish/record", apis.ProductVoicecontroller.GetVoicePublishRecord)
	admin.POST("/product/voice/unitList", apis.ProductVoicecontroller.GetVoiceUnitList)

	//场景模板 /v1/platform/web/open/sceneTemplate/list
	admin.POST("/sceneTemplate/list", apis.SceneTemplatecontroller.List)
	admin.GET("/sceneTemplate/detail/:id", apis.SceneTemplatecontroller.Get)
	admin.POST("sceneTemplate/add", apis.SceneTemplatecontroller.Add)
	admin.PUT("/sceneTemplate/edit", apis.SceneTemplatecontroller.Update)
	admin.POST("/sceneTemplate/edit", apis.SceneTemplatecontroller.Update)
	admin.DELETE("/sceneTemplate/delete/:id", apis.SceneTemplatecontroller.Delete)
	admin.POST("/sceneTemplate/delete/:id", apis.SceneTemplatecontroller.Delete)
	admin.POST("/sceneTemplate/setStatus", apis.SceneTemplatecontroller.SetStatus)

	//产品说明书
	admin.POST("/manual/add", apis.ProductManualControl.CreateProductManual)
	admin.POST("/manual/edit", apis.ProductManualControl.UpdateProductManual)
	admin.POST("/manual/del", apis.ProductManualControl.DeleteProductManual)
	admin.GET("/manual/detail", apis.ProductManualControl.GetProductManual)
	admin.POST("/manual/list", apis.ProductManualControl.GetProductManualList)

	//物理模型导出
	admin.GET("/product/thingModels/export", apis.Productcontroller.ExportThingsModel)

	//面板设计器
	admin.POST("/panel/add", apis.Panelcontroller.Add)
	admin.POST("/panel/edit", apis.Panelcontroller.EditStudio)
	admin.POST("/panel/editInfo", apis.Panelcontroller.Edit)
	admin.POST("/panel/setStatus", apis.Panelcontroller.SetStatus)
	admin.POST("/panel/delete/:id", apis.Panelcontroller.Delete)
	admin.GET("/panel/detail/:id", apis.Panelcontroller.QueryDetail)
	admin.POST("/panel/list", apis.Panelcontroller.QueryList)
	//社区产品
	admin.POST("/community/product/add", apis.CommunityProductcontroller.Add)
	admin.POST("/community/product/edit", apis.CommunityProductcontroller.Edit)
	admin.POST("/community/product/setStatus", apis.CommunityProductcontroller.SetStatus)
	admin.POST("/community/product/delete/:id", apis.CommunityProductcontroller.Delete)
	admin.GET("/community/product/detail/:id", apis.CommunityProductcontroller.QueryDetail)
	admin.POST("/community/product/list", apis.CommunityProductcontroller.QueryList)

	//产品规则设置
	admin.POST("/product/ruleSetting/add", apis.ThingModelRulecontroller.Add)
	admin.POST("/product/ruleSetting/edit", apis.ThingModelRulecontroller.Edit)
	admin.POST("/product/ruleSetting/setStatus", apis.ThingModelRulecontroller.SetStatus)
	admin.POST("/product/ruleSetting/delete/:id", apis.ThingModelRulecontroller.Delete)
	admin.GET("/product/ruleSetting/detail/:id", apis.ThingModelRulecontroller.QueryDetail)
	admin.POST("/product/ruleSetting/list", apis.ThingModelRulecontroller.QueryList)
	admin.GET("/product/ruleSetting/listByProductId", apis.ThingModelRulecontroller.QueryListByProductId)

	//文档中心
	admin.POST("/product/doc/add", apis.Documentscontroller.Add)
	admin.POST("/product/doc/edit", apis.Documentscontroller.Edit)
	admin.POST("/product/doc/delete/:id", apis.Documentscontroller.Delete)
	admin.GET("/product/doc/detail/:id", apis.Documentscontroller.QueryDetail)
	admin.POST("/product/doc/list", apis.Documentscontroller.QueryList)

	//APP
	admin.POST("/productTestAccount/add", apis.ProductTestAccountcontroller.Add)
	admin.POST("/productTestAccount/list", apis.ProductTestAccountcontroller.QueryList)
	admin.POST("/productTestAccount/delete/:id", apis.ProductTestAccountcontroller.Delete)
}
