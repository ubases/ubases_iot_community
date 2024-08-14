package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	services "cloud_platform/iot_cloud_api_service/controls/global"
	"cloud_platform/iot_cloud_api_service/controls/system/entitys"
	apiservice "cloud_platform/iot_cloud_api_service/controls/system/services"
	"cloud_platform/iot_common/iotutil"
	"fmt"

	"github.com/gin-gonic/gin"

	"cloud_platform/iot_common/iotgin"
)

var Usercontroller UserController

type UserController struct {
} //用户操作控制器

var userservices = apiservice.SysUserService{}

// 忘记密码获取验证码
func (UserController) SendVerificationCodeForExists(c *gin.Context) {
	userName := c.Query("userName")
	if iotutil.IsEmpty(userName) {
		iotgin.ResBadRequest(c, "userName not found")
		return
	}
	codeType := c.Query("type")
	if iotutil.IsEmpty(codeType) {
		iotgin.ResBadRequest(c, "type not found")
		return
	}
	var (
		lang     = controls.GetLang(c)
		tenantId = controls.GetTenantId(c)
	)
	res, err := userservices.SendVerificationCodeForExists(userName, tenantId, lang, iotutil.ToInt32(codeType))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 忘记密码重置密码
func (UserController) ForgetPassword(c *gin.Context) {
	var req entitys.UserResetPasswordNoTokenReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	res, err := userservices.ForgetPassword(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (UserController) GetUserDetail(c *gin.Context) {
	id := c.Query("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	res, err := userservices.GetUserDetail(id)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (UserController) GetLoginUserInfo(c *gin.Context) {
	id, _ := c.Get("userId")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found "+iotutil.ToString(id))
		return
	}

	res, err := userservices.QueryLoginUserInfo(iotutil.ToString(id))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (UserController) GetUserProfile(c *gin.Context) {
	id, _ := c.Get("userId")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "token not found")
		return
	}
	res, err := userservices.GetUserProfile(iotutil.ToString(id))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (UserController) QueryUserList(c *gin.Context) {
	//userid, _ := c.Get("UserId")
	// var resq = entitys.QueryUser{
	// 	Page:      iotutil.ToInt64(c.DefaultQuery("pageNum", "1")),
	// 	Limit:     iotutil.ToInt64(c.DefaultQuery("pageSize", "10")),
	// 	BeginTime: c.DefaultQuery("beginTime", ""),
	// 	EndTime:   c.DefaultQuery("endTime", ""),
	// }

	var req entitys.QueryUser
	err := c.ShouldBindQuery(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, total, err := userservices.QueryUserList(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(req.Page))
}

func (UserController) AddUser(c *gin.Context) {
	var req entitys.UserCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysUser_SetCommonFiled(&req, iotutil.ToInt64(userid), 1)

	id, err := userservices.AddUser(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
	services.RefreshUserCache()
}

func (UserController) EditUser(c *gin.Context) {
	var req entitys.UserCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.Id == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")
	entitys.SysUser_SetCommonFiled(&req, iotutil.ToInt64(userid), 2)

	id, err := userservices.UpdateUser(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
	services.RefreshUserCache()
}

// 个人中心修改用户自己信息
func (UserController) EditUserCenter(c *gin.Context) {
	var req entitys.UserCenterEditReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//赋值公共字段
	userid, _ := c.Get("userId")

	id, err := userservices.UpdateUserCenter(req, iotutil.ToInt64(userid))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, id)
}

// 重置为默认密码
func (UserController) ResetUserPwd(c *gin.Context) {
	var req entitys.UserResetPasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.UserId == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = userservices.UpdateResetPassword(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// 修改密码
func (UserController) UpdateUserPwd(c *gin.Context) {
	var req entitys.UserUpdatePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId, _ := c.Get("userId")
	if req.OldPassword == "" || req.NewPassword == "" || userId == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = userservices.UpdatePassword(iotutil.ToString(userId), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (UserController) DeleteUser(c *gin.Context) {
	var req entitys.DeleteCommonQuery
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if len(req.Ids) == 0 {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = userservices.DeleteUser(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	//清理删除用户的Token
	for _, userId := range req.Ids {
		controls.ClearTokenByUserId(iotutil.ToInt64(userId))
	}
	iotgin.ResSuccessMsg(c)
	services.RefreshUserCache()
}

func (UserController) Login(c *gin.Context) {
	var resq entitys.UserLogin
	err := c.ShouldBindJSON(&resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if resq.Username == "" || resq.Password == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	//获取浏览器信息
	os, browserName, browserVersion := controls.GetUserAgent(c)
	resq.Explorer = fmt.Sprintf("%v %v", browserName, browserVersion)
	resq.Os = os
	resq.ClientIp = c.ClientIP()
	token, refreshToken, expiresAt, err := userservices.UserLogin(resq)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var res = map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
		"expiresAt":    expiresAt,
	}
	iotgin.ResSuccess(c, res)
}

func (UserController) Logout(c *gin.Context) {
	token := controls.GetToken(c)
	//userId, _ := c.Get("userId")
	err := userservices.UserLogout(token)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, "ok")
}

func (UserController) QueryUserRouters(c *gin.Context) {
	userId, _ := c.Get("userId")
	list, err := authruleServices.QueryUserRouters(iotutil.ToInt64(userId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, list)
}

func (UserController) RefreshToken(c *gin.Context) {
	var req entitys.RefreshToken
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	if req.RefreshToken == "" {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	token, refreshToken, expiresAt, err := userservices.RefreshToken(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var res = map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
		"expiresAt":    expiresAt,
	}
	iotgin.ResSuccess(c, res)
}
