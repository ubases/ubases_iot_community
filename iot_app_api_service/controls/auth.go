package controls

import (
	"cloud_platform/iot_app_api_service/config"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotstruct"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"go-micro.dev/v4/metadata"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserID         int64       `json:"userId"`
	Nickname       string      `json:"nickname"`
	Account        string      `json:"account"`
	Avatar         string      `json:"avatar"`
	TenantId       string      `json:"tenantId"` //租户ID，开放测试平台使用
	AppKey         string      `json:"appKey"`
	RegionServerId interface{} `json:"regionServerId"`
}

// 系统信息
type SystemInfo struct {
	Os      string `json:"os"`
	Version string `json:"version"`
	Model   string `json:"model"`
}

func GetToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Request.Header.Get("token")
	}
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}
	return parts[1]
}

func SetParams(c *gin.Context) {
	c.Set("IsShowErrMsg", config.Global.Service.ResponseRealError)
	c.Next()
}

// AuthCheck 路由过滤器方法，包括token解析和验证
func AuthCheck(c *gin.Context) {
	token := GetToken(c)
	if token == "" {
		iotlogger.LogHelper.Error("AuthCheck: header缺少token或Authorization.")
		iotgin.ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}
	//var err error
	var code = ioterrs.SUCCESS
	var userInfoStr string
	var err error
	var userInfo UserInfo

	//防止redis偶然错误，导致用户token校验失败，最多重试3次
	for i := 0; i < 3; i++ {
		userInfoStr, err = iotredis.GetClient().Get(context.Background(), token).Result()
		if err == nil || err == redis.Nil {
			break
		}
		//重试越多延时越多
		time.Sleep(time.Duration(1+i) * 500 * time.Millisecond)
	}

	if userInfoStr == "" {
		go Logout("AuthCheck:Token已过期", c)
		iotgin.ResFailCode(c, ioterrs.ERROR_AUTH_CHECK_TOKEN_EXPIRE.Msg, ioterrs.ERROR_AUTH_CHECK_TOKEN_EXPIRE.Code)
		c.Abort()
		return
	}

	err = iotutil.JsonToStruct(userInfoStr, &userInfo)
	if err != nil {
		iotlogger.LogHelper.Errorf("AuthCheck:解析Token对应的userinfo信息错误:%s.", err.Error())
		iotgin.ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}
	if code.Code != ioterrs.SUCCESS.Code {
		iotlogger.LogHelper.Error("AuthCheck:code不等于200")
		iotgin.ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}

	c.Set("UserId", userInfo.UserID)
	c.Set("Nickname", userInfo.Nickname)
	c.Set("Account", userInfo.Account)
	c.Set("tenantId", userInfo.TenantId)
	c.Set("regionServerId", userInfo.RegionServerId)
	//c.Set("avatar", userInfo.Avatar)
	c.Set("Token", token)
	c.Set("IsShowErrMsg", config.Global.Service.ResponseRealError)
	//临时方案，后续引入事件服务再优化
	PushUcUserOperate(&userInfo, c)

	c.Next()
}

// 将用户信息和token传给后端微服务
func WithUserContext(c *gin.Context) context.Context {
	ctx := metadata.NewContext(context.Background(),
		map[string]string{
			"userId":   iotutil.ToString(c.GetInt64("UserId")),
			"tenantId": c.Request.Header.Get("tenantId"),
			"appKey":   c.Request.Header.Get("appKey"),
			"lang":     c.GetHeader("lang"),
			"region":   c.GetHeader("region"),
			"tz":       c.GetHeader("tz"),
			"token":    GetToken(c),
		})
	return ctx
}

func WithUserContextV2(c *gin.Context, userId int64, regionId int64, appKey, tenantId string) context.Context {
	ctx := metadata.NewContext(context.Background(),
		map[string]string{
			"userId":   iotutil.ToString(userId),
			"tenantId": tenantId,
			"appKey":   appKey,
			"lang":     c.GetHeader("lang"),
			"region":   c.GetHeader("region"),
			"token":    GetToken(c),
		})
	return ctx
}

func GetSystemInfoRaw(c *gin.Context) string {
	return c.GetHeader("x-sys-info")
}

func GetSystemInfo(c *gin.Context) SystemInfo {
	var xSysInfo = c.GetHeader("x-sys-info")
	sysArrs := strings.Split(xSysInfo, ",")
	if len(sysArrs) == 4 {
		return SystemInfo{
			Os:      sysArrs[0],
			Model:   sysArrs[1] + "," + sysArrs[2],
			Version: sysArrs[3],
		}
	}
	//最多支取3个
	if len(sysArrs) > 3 {
		sysArrs = sysArrs[0:3]
	}
	sysNewArr := make([]string, 4)
	for i := 0; i < len(sysArrs); i++ {
		sysNewArr[i] = sysArrs[i]
	}
	return SystemInfo{
		Os:      sysNewArr[0],
		Model:   sysNewArr[1],
		Version: sysNewArr[2],
	}
}

func GetUserId(c *gin.Context) int64 {
	return c.GetInt64("UserId")
}

func GetNickName(c *gin.Context) string {
	return c.GetString("Nickname")
}

func GetAccount(c *gin.Context) string {
	return c.GetString("Account")
}

func GetTenantId(c *gin.Context) string {
	tenantId := c.GetHeader("tenantId")
	if tenantId == "" {
		tenantId = c.GetString("tenantId")
	}
	return tenantId
}

func GetAppKey(c *gin.Context) string {
	return c.GetHeader("appKey")
}
func GetAppPushId(c *gin.Context) string {
	return c.GetHeader("appPushId")
}

// 获取语言
func GetLang(c *gin.Context) string {
	return c.GetHeader("lang")
}

// 获取区域Id
func GetRegion(c *gin.Context) string {
	return c.GetHeader("region")
}

// 获取区域Id，默认区域为1
func GetRegionInt(c *gin.Context) int64 {
	region := c.GetHeader("region")
	var defaultRegion int64 = 1
	if region == "" {
		return defaultRegion
	} else {
		regionInt, err := iotutil.ToInt64AndErr(region)
		if err != nil {
			return defaultRegion
		}
		return regionInt
	}
}

// 获取时区
func GetTimezone(c *gin.Context) string {
	return c.GetHeader("tz")
}

func PushUcUserOperate(u *UserInfo, c *gin.Context) {
	uc := iotstruct.UcUserOperate{
		Id:         iotutil.GetNextSeqInt64(),
		TenantId:   u.TenantId,
		AppKey:     u.AppKey,
		UserId:     u.UserID,
		Account:    u.Account,
		RequestUri: c.Request.RequestURI,
		Ip:         c.ClientIP(),
		OptTime:    time.Now(),
	}
	_, err := iotredis.GetClient().LPush(context.Background(), iotconst.APPOPERATELIST, uc).Result()
	if err != nil {
		iotlogger.LogHelper.Errorf("PushUcUserOperate,error.%s", err.Error())
	}
}

// 修改密码后清除该用户的token
func ClearTokenByUserId(id int64) {
	ClearToken(id, true)
}

func ClearToken(id int64, isAll bool) {
	key := iotconst.USERTOKENPREFIX + strconv.Itoa(int(id))
	mapToken, err := iotredis.GetClient().HGetAll(context.Background(), key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			iotlogger.LogHelper.Errorf("ClearTokenByUserId,error:%s", err.Error())
		}
		return
	}
	var tokenList []string
	for k, v := range mapToken {
		if isAll { //清除所有
			tokenList = append(tokenList, k)
		} else { //清除过期
			expireAt, err1 := strconv.Atoi(v)
			if err1 == nil {
				if time.Now().After(time.Unix(int64(expireAt), 0)) {
					//已过期
					tokenList = append(tokenList, k)
				}
			}
		}
	}
	if len(tokenList) > 0 {
		_, err = iotredis.GetClient().Del(context.Background(), tokenList...).Result()
		if err != nil {
			iotlogger.LogHelper.Errorf("ClearTokenByUserId,error:%s", err.Error())
		}
	}
}

// 插入token
func CacheTokenByUserId(id int64, token string, expiresAt int64) {
	key := iotconst.USERTOKENPREFIX + strconv.Itoa(int(id))
	_, err := iotredis.GetClient().HSet(context.Background(), key, token, expiresAt).Result()
	if err != nil {
		return
	}
	//清除过期token
	ClearToken(id, false)
}

// 通过token获取用户信息
func TokenGetUserInfo(token string) (*UserInfo, error) {
	var userInfo UserInfo
	if token == "" {
		return nil, errors.New("token is empty")
	}
	cacheData := iotredis.GetClient().Get(context.Background(), token)
	if cacheData.Err() == nil && cacheData.Val() != "" {
		err := iotutil.JsonToStruct(cacheData.Val(), &userInfo)
		if err != nil {
			iotlogger.LogHelper.Errorf("TokenGetUserInfo,error:%s", err.Error())
			return nil, err
		}
	}
	return &userInfo, nil
}

func Logout(msg string, c *gin.Context) error {
	defer iotutil.PanicHandler(msg)
	iotlogger.LogHelper.Infof("Controls.Logout,msg:%s", msg)
	//退出登录，清理push token，清理推送
	token := GetToken(c)
	//var userId int64
	////验证解析token
	//res, err := rpc.AppAuthService.VerifyToken(context.Background(), &proto.VerifyTokenRequest{Token: token})
	//if err == nil {
	//	userId = res.UserInfo.Id
	//}
	if token != "" {
		//清理token
		iotredis.GetClient().Del(context.Background(), token)
	}
	//
	appPushId := GetAppPushId(c)
	if appPushId != "" {
		iotlogger.LogHelper.Error("Logout.AppPushTokenUserService.appPushId is empty")
		return errors.New("appPushId is empty")
	}
	appKey := GetAppKey(c)
	_, err := rpc.AppPushTokenUserService.Delete(context.Background(), &proto.AppPushTokenUser{
		AppPushId: appPushId,
		AppKey:    appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("Logout.AppPushTokenUserService.Delete err:%s", err.Error())
	}
	return err
}
