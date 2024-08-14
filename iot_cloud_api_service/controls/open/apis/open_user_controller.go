package apis

import (
	"cloud_platform/iot_cloud_api_service/controls"
	"cloud_platform/iot_cloud_api_service/controls/open/entitys"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotutil"
	"fmt"
	"strings"

	apiservice "cloud_platform/iot_cloud_api_service/controls/open/services"

	"github.com/gin-gonic/gin"
)

var OpenUsercontroller OpenUserController

var OpenUserService apiservice.OpenUserService

type OpenUserController struct {
} //用户操作控制器

func (OpenUserController) GetTest(c *gin.Context) {
	id := c.Query("id")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "id not found")
		return
	}
	iotgin.ResSuccess(c, id)
}

// Register 开放平台注册
func (OpenUserController) Register(c *gin.Context) {
	var req entitys.OpenUserRegisterReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.IP = c.ClientIP()
	res, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).RegisterUser(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// GuideCheck 引导检查
func (OpenUserController) GuideCheck(c *gin.Context) {
	var req entitys.GuideCheckReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.UserId = controls.GetUserId(c)
	req.TenantId = controls.GetTenantId(c)
	req.AccountType = int32(controls.GetAccountType(c))
	err = OpenUserService.SetContext(controls.WithOpenUserContext(c)).GuideCheck(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// SetHasGuided 设置引导完成
func (OpenUserController) SetHasGuided(c *gin.Context) {
	req := entitys.GuideCheckReq{
		UserId: controls.GetUserId(c),
	}
	err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).SetHasGuided(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

// Login 登录
func (OpenUserController) Login(c *gin.Context) {
	var resq entitys.UserLoginReq
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
	token, refreshToken, expiresAt, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).UserLogin(resq, c.ClientIP())
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

// 注销登录
func (OpenUserController) Logout(c *gin.Context) {
	token := controls.GetToken(c)
	err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).UserLogout(token)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, "ok")
}

// 修改密码（忘记密码）
func (OpenUserController) UpdateOpenUserPwd(c *gin.Context) {
	var req entitys.OpenUserUpdatePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	userId, _ := c.Get("userId")
	if req.NewPassword == "" || userId == nil {
		iotgin.ResFailCode(c, "参数异常", -1)
		return
	}
	err = OpenUserService.UpdatePassword(iotutil.ToString(userId), req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpenUserController) ForgetPassword(c *gin.Context) {
	var req entitys.OpenUserUpdatePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	err = OpenUserService.ForgetPassword(req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccessMsg(c)
}

func (OpenUserController) GetUserProfile(c *gin.Context) {
	id, _ := c.Get("userId")
	if iotutil.IsEmpty(id) {
		iotgin.ResBadRequest(c, "token not found")
		return
	}
	tenantId, _ := c.Get("tenantId")
	if iotutil.IsEmpty(tenantId) {
		iotgin.ResBadRequest(c, "token not found")
		return
	}

	res, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).GetUserProfile(iotutil.ToString(id), iotutil.ToString(tenantId))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 选择切换租户
func (OpenUserController) ChangeTenant(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)

	var req entitys.RefreshToken
	err := c.ShouldBindJSON(&req)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	req.RefreshToken = strings.Split(token, " ")[1]

	token, refreshToken, expiresAt, err := OpenUserService.ChangeTenant(req)
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

func (OpenUserController) RefreshToken(c *gin.Context) {
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
	token, refreshToken, expiresAt, err := OpenUserService.RefreshToken(req, c.ClientIP())
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

// 获取登录菜单
func (OpenUserController) GetRouters(c *gin.Context) {
	res, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).GetRouters()
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

func (OpenUserController) GetVerificationCode(c *gin.Context) {
	var (
		userName = c.Query("userName")
		lang     = controls.GetLang(c)
		tenantId = controls.GetTenantId(c)
		codeType = c.Query("type")
	)

	code, err := iotutil.ToInt32Err(codeType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}

	res, err := OpenUserService.SendVerificationCode(userName, tenantId, lang, code)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// 获取验证码
func (OpenUserController) GetVerificationCodeForExists(c *gin.Context) {
	var (
		userName = c.Query("userName")
		lang     = controls.GetLang(c)
		tenantId = controls.GetTenantId(c)
		codeType = c.Query("type")
	)

	res, err := OpenUserService.SendVerificationCodeForExists(userName, tenantId, lang, iotutil.ToInt32(codeType))
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResSuccess(c, res)
}

// APP用户管理
func (OpenUserController) QueryUserList(c *gin.Context) {
	pageNum := iotutil.ToInt64(c.Query("page"))
	pageSize := iotutil.ToInt64(c.Query("limit"))
	userMobile := c.Query("userMobile")
	userAccount := c.Query("userAccount")
	userEmail := c.Query("userEmail")
	appType := c.Query("appType")
	res, total, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).QueryAppUserList(pageNum, pageSize, userMobile, userAccount, userEmail, appType)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(pageNum))
}

// APP用户绑定设备
func (OpenUserController) QueryUserDeviceList(c *gin.Context) {
	pageNum := iotutil.ToInt64(c.Query("page"))
	pageSize := iotutil.ToInt64(c.Query("limit"))
	customerUserId := c.Query("customerUserId")
	res, total, err := OpenUserService.SetContext(controls.WithOpenUserContext(c)).QueryUserDeviceList(pageNum, pageSize, customerUserId)
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	iotgin.ResPageSuccess(c, res, total, int(pageNum))
}

// LangTypeList 从字典表获取语言类型
func (s OpenUserController) LangOpenTypeList(c *gin.Context) {
	list, err := OpenUserService.QueryLangBaseDataList(entitys.LangBaseDataQuery{
		DictType: "language_type",
	})
	if err != nil {
		iotgin.ResErrCli(c, err)
		return
	}
	var res []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
	for _, data := range list {
		res = append(res, struct {
			Name string `json:"name"`
			Code string `json:"code"`
		}{Name: data.DictLabel, Code: data.DictValue})
	}
	iotgin.ResSuccess(c, res)
}
