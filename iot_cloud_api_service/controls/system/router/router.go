package router

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/system/apis"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	webApiPrefix := "/v1/platform/web/system"
	admin := e.Group(webApiPrefix)
	admin.POST("/refreshToken", apis.Usercontroller.RefreshToken)
	admin.POST("/login", apis.Usercontroller.Login)
	//重置密码
	admin.POST("/auth/resetPwd", apis.Usercontroller.ForgetPassword)
	//获取验证码
	admin.GET("/auth/getVerificationCodeForExists", apis.Usercontroller.SendVerificationCodeForExists)

	admin.Use(controls.AuthCheck)
	admin.POST("/logout", apis.Usercontroller.Logout)
	admin.GET("/auth/userList", apis.Usercontroller.QueryUserList)
	admin.POST("/auth/addUser", apis.Usercontroller.AddUser)
	admin.PUT("/auth/editUser", apis.Usercontroller.EditUser)
	admin.POST("/auth/editUser", apis.Usercontroller.EditUser)
	admin.PUT("/auth/editUserCenter", apis.Usercontroller.EditUserCenter)
	admin.POST("/auth/editUserCenter", apis.Usercontroller.EditUserCenter)

	admin.DELETE("/auth/deleteUser", apis.Usercontroller.DeleteUser)
	admin.POST("/auth/deleteUser", apis.Usercontroller.DeleteUser)
	admin.GET("/auth/getEditUser", apis.Usercontroller.GetUserDetail)
	admin.PUT("/auth/resetUserPwd", apis.Usercontroller.ResetUserPwd)
	admin.POST("/auth/resetUserPwd", apis.Usercontroller.ResetUserPwd)
	admin.GET("/user/profile", apis.Usercontroller.GetUserProfile)
	//用户信息可以缓存，以提高性能
	admin.PUT("/user/edit", apis.Usercontroller.EditUser)
	admin.POST("/user/edit", apis.Usercontroller.EditUser)
	admin.PUT("/user/updatePwd", apis.Usercontroller.UpdateUserPwd)
	admin.POST("/user/updatePwd", apis.Usercontroller.UpdateUserPwd)
	admin.GET("/user/getInfo" /* cache.CacheByRequestURI(redisStore, 15*time.Second),*/, apis.Usercontroller.GetLoginUserInfo)
	admin.GET("/user/logout", apis.Usercontroller.Logout)
	admin.GET("/user/getRouters", apis.Usercontroller.QueryUserRouters)

	//webApiPrefix := "/v1/platform/web"
	//admin := e.Group(webApiPrefix)
	///admin.Use(iotgin.AuthCheck)
	admin.GET("/auth/menuList", apis.Authrulecontroller.QueryList)
	admin.GET("/auth/menu/", apis.Authrulecontroller.QueryDetail)
	admin.POST("/auth/addMenu", apis.Authrulecontroller.Add)
	admin.PUT("/auth/editMenu", apis.Authrulecontroller.Edit)
	admin.POST("/auth/editMenu", apis.Authrulecontroller.Edit)
	admin.DELETE("/auth/deleteMenu", apis.Authrulecontroller.Delete)
	admin.POST("/auth/deleteMenu", apis.Authrulecontroller.Delete)

	admin.GET("/auth/openmenuList", apis.OpenAuthrulecontroller.QueryList)
	admin.GET("/auth/openmenu", apis.OpenAuthrulecontroller.QueryDetail)
	admin.POST("/auth/openaddMenu", apis.OpenAuthrulecontroller.Add)
	admin.PUT("/auth/openeditMenu", apis.OpenAuthrulecontroller.Edit)
	admin.POST("/auth/openeditMenu", apis.OpenAuthrulecontroller.Edit)
	admin.DELETE("/auth/opendeleteMenu", apis.OpenAuthrulecontroller.Delete)
	admin.POST("/auth/opendeleteMenu", apis.OpenAuthrulecontroller.Delete)

	//开发者认证审核和列表
	admin.GET("/opendev/list", apis.OpenDevcontroller.QueryList)
	admin.GET("/opendev/detail", apis.OpenDevcontroller.QueryDetail)
	admin.POST("/opendev/auth", apis.OpenDevcontroller.OpenDevAuth)

	//webApiPrefix := "/v1/platform/web"
	//admin := e.Group(webApiPrefix)
	//admin.Use(iotgin.AuthCheck)
	admin.GET("/auth/roleList", apis.Rolecontroller.QueryList)
	//角色修改的时候获取角色信息
	admin.GET("/auth/editRole", apis.Rolecontroller.QueryDetail)
	//角色新增时获取菜单列表
	admin.GET("/auth/addRole", apis.Rolecontroller.AddRoleByMenuList)
	//角色数据权限查询
	admin.GET("/auth/roleDeptTreeSelect", apis.Rolecontroller.RoleDeptTreeSelect)
	//角色新增
	admin.POST("/auth/addRole", apis.Rolecontroller.Add)
	//角色修改
	admin.PUT("/auth/editRole", apis.Rolecontroller.Edit)
	admin.POST("/auth/editRole", apis.Rolecontroller.Edit)
	//角色数据权限修改
	admin.PUT("/auth/roleDataScope", apis.Rolecontroller.RoleDataScope)
	admin.POST("/auth/roleDataScope", apis.Rolecontroller.RoleDataScope)
	//角色状态修改
	admin.PUT("/auth/statusSetRole", apis.Rolecontroller.StatusSetRole)
	admin.POST("/auth/statusSetRole", apis.Rolecontroller.StatusSetRole)
	//删除角色
	admin.DELETE("/auth/deleteRole", apis.Rolecontroller.Delete)
	admin.POST("/auth/deleteRole", apis.Rolecontroller.Delete)

	//webApiPrefix := "/v1/platform/web"
	//admin := e.Group(webApiPrefix)
	//admin.Use(iotgin.AuthCheck)
	admin.GET("/auth/postList", apis.Postcontroller.QueryList)
	admin.GET("/auth/postGet", apis.Postcontroller.QueryDetail)
	admin.POST("/auth/postAdd", apis.Postcontroller.Add)
	admin.PUT("/auth/postEdit", apis.Postcontroller.Edit)
	admin.POST("/auth/postEdit", apis.Postcontroller.Edit)
	admin.DELETE("/auth/postDelete", apis.Postcontroller.Delete)
	admin.POST("/auth/postDelete", apis.Postcontroller.Delete)

	//webApiPrefix := "/v1/platform/web"
	//admin := e.Group(webApiPrefix)
	//admin.Use(iotgin.AuthCheck)
	admin.GET("/auth/deptList", apis.Deptcontroller.QueryList)
	admin.GET("/auth/deptGet", apis.Deptcontroller.QueryDetail)
	admin.POST("/auth/deptAdd", apis.Deptcontroller.Add)
	admin.PUT("/auth/deptEdit", apis.Deptcontroller.Edit)
	admin.POST("/auth/deptEdit", apis.Deptcontroller.Edit)
	admin.DELETE("/auth/deptDelete", apis.Deptcontroller.Delete)
	admin.POST("/auth/deptDelete", apis.Deptcontroller.Delete)

	//webApiPrefix := "/v1/platform/web"
	//admin := e.Group(webApiPrefix)
	//admin.Use(iotgin.AuthCheck)

	webApiPrefix = "/v1/platform/web"
	basedata := e.Group(webApiPrefix)
	basedata.GET("/basedata/dictdata/list", apis.BaseDatacontroller.QueryBaseDataList)
	basedata.Use(controls.AuthCheck)
	basedata.GET("/basedata/dictdata/detail/:id", apis.BaseDatacontroller.GetBaseDataDetail)
	basedata.POST("/basedata/dictdata/add", apis.BaseDatacontroller.AddTConfigDictData)
	basedata.POST("/basedata/dictdata/edit", apis.BaseDatacontroller.EditBaseData)
	basedata.POST("/basedata/dictdata/delete/:id", apis.BaseDatacontroller.DeleteBaseData)
	basedata.GET("/basedata/dicttype/detail/:id", apis.BaseDatacontroller.GetBaseTypeDetail)
	basedata.GET("/basedata/dicttype/list", apis.BaseDatacontroller.QueryBaseDataTypeList)
	basedata.POST("/basedata/dicttype/add", apis.BaseDatacontroller.AddTConfigDictDataType)
	basedata.POST("/basedata/dicttype/edit", apis.BaseDatacontroller.EditBaseDataType)
	basedata.POST("/basedata/dicttype/delete/:id", apis.BaseDatacontroller.DeleteBaseDataType)

	basedata.GET("/basedata/translate/detail", apis.BaseDatacontroller.GetTConfigTranslateDetail)
	basedata.POST("/basedata/translate/add", apis.BaseDatacontroller.AddTConfigTranslate)
	basedata.POST("/basedata/translate/edit", apis.BaseDatacontroller.EditTConfigTranslate)

	basedata.GET("/basedata/translate/language/list", apis.BaseDatacontroller.QueryTranslateLanguageList)

	//开发者账户相关 /v1/platform/web/system/deve/status
	admin.GET("/deve/list", apis.Developercontroller.List)
	admin.GET("/deve/basicList", apis.Developercontroller.BasicList)
	admin.GET("/deve/companyList", apis.Developercontroller.ListCompany)
	admin.GET("/deve/detail", apis.Developercontroller.Detail)
	admin.POST("/deve/add", apis.Developercontroller.Add)
	admin.PUT("/deve/status", apis.Developercontroller.SetStatus)
	admin.POST("/deve/status", apis.Developercontroller.SetStatus)
	admin.PUT("/deve/edit", apis.Developercontroller.Edit)
	admin.POST("/deve/edit", apis.Developercontroller.Edit)
	admin.PUT("/deve/resetPassword", apis.Developercontroller.ResetPassword)
	admin.POST("/deve/resetPassword", apis.Developercontroller.ResetPassword)
	admin.DELETE("/deve/delete", apis.Developercontroller.Delete)
	admin.POST("/deve/delete", apis.Developercontroller.Delete)

	//Oem 底部菜单
	admin.GET("/oemapp/menu", apis.OemAppDefMenucontroller.QueryList)
	admin.POST("/oemapp/menu", apis.OemAppDefMenucontroller.Add)
	admin.PUT("/oemapp/menu", apis.OemAppDefMenucontroller.Edit)
	admin.POST("/oemapp/menu/update", apis.OemAppDefMenucontroller.Edit)

	//SysApp文档目录
	admin.POST("/app/doc/directory", apis.SysAppDocDircontroller.CreateDir)
	admin.PUT("/app/doc/directory", apis.SysAppDocDircontroller.UpdateDir)
	admin.POST("/app/doc/directory/update", apis.SysAppDocDircontroller.UpdateDir)
	admin.DELETE("/app/doc/directory", apis.SysAppDocDircontroller.DeleteDir)
	admin.POST("/app/doc/directory/delete", apis.SysAppDocDircontroller.DeleteDir)
	admin.GET("/app/doc/directory/detail", apis.SysAppDocDircontroller.DetailDir)
	admin.GET("/app/doc/directory", apis.SysAppDocDircontroller.ListDir)

	//SysApp文档词条
	admin.POST("/app/doc/entry", apis.SysAppEntrycontroller.EntrySave)
	admin.GET("/app/doc/entry/detail", apis.SysAppEntrycontroller.EntryDetail)
	admin.POST("/app/doc/entry/seting", apis.SysAppEntrycontroller.EntrySetingSave)
	admin.GET("/app/doc/entry/seting", apis.SysAppEntrycontroller.EntrySetingDetail)
	admin.DELETE("/app/doc/entry", apis.SysAppEntrycontroller.EntryDelete)
	admin.POST("/app/doc/entry/delete", apis.SysAppEntrycontroller.EntryDelete)
	admin.POST("/app/doc/entry/list", apis.SysAppEntrycontroller.EntryList)

	//SysApp帮助中心
	admin.POST("/app/helpCenter/add", apis.SysAppHelpCentercontroller.CreateHelpCenter)
	admin.POST("/app/helpCenter/copy", apis.SysAppHelpCentercontroller.CopyHelpCenter)
	admin.POST("/app/helpCenter/edit", apis.SysAppHelpCentercontroller.UpdateHelpCenter)
	admin.POST("/app/helpCenter/del", apis.SysAppHelpCentercontroller.DeleteHelpCenter)
	admin.GET("/app/helpCenter/detail", apis.SysAppHelpCentercontroller.GetHelpCenter)
	admin.POST("/app/helpCenter/list", apis.SysAppHelpCentercontroller.GetHelpCenterList)

	//APP服务器管理
	admin.POST("/appService/add", apis.SysRegionServercontroller.Add)
	admin.POST("/appService/edit", apis.SysRegionServercontroller.Edit)
	admin.POST("/appService/delete", apis.SysRegionServercontroller.Delete)
	admin.GET("/appService/detail/:id", apis.SysRegionServercontroller.QueryDetail)
	admin.POST("/appService/setStatus", apis.SysRegionServercontroller.SetStatus)
	admin.POST("/appService/list", apis.SysRegionServercontroller.QueryList)

	msApiPrefix := "/v1/platform/web/data"
	ms := e.Group(msApiPrefix)
	ms.Use(controls.AuthCheck)
	//发送短信/邮件记录 t_ms_notice_record
	ms.POST("/noticeInfo/sendRecord", apis.MsNoticeRecordcontroller.QueryList)
	ms.GET("/noticeInfo/detail/:id", apis.MsNoticeRecordcontroller.QueryDetail)
	ms.POST("/noticeInfo/delete", apis.MsNoticeRecordcontroller.Delete)
}
