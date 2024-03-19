package controls

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/ioterrs"
	"cloud_platform/iot_common/iotgin"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go-micro.dev/v4/metadata"
)

type UserInfo struct {
	UserID      int64    `json:"userId"`
	Nickname    string   `json:"nickName"`
	Avatar      string   `json:"avatar"`
	DeptId      int64    `json:"deptId"`   //云平台用户才有
	RoleIds     []string `json:"roleIds"`  //云平台用户才有，多个用逗号分隔
	PostIds     []string `json:"postIds"`  //云平台用户才有，多个用逗号分隔
	TenantId    string   `json:"tenantId"` //租户ID，开放测试平台使用
	AccountType int32    `json:"accountType"`
}

type OpenUserInfo struct {
	UserID   int64  `json:"userId"`
	Nickname string `json:"nickName"`
	Avatar   string `json:"avatar"`
	TenantId string `json:"tenantId"` //租户ID，开放测试平台使用
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

// AuthCheck 路由过滤器方法，包括token解析和验证
func AuthCheck(c *gin.Context) {
	token := GetToken(c)
	if token == "" {
		iotlogger.LogHelper.Info("token not found")
		iotgin.ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}
	var err error
	var code ioterrs.KVCode = ioterrs.SUCCESS
	var userInfo UserInfo
	//先从缓存拿token
	//iotredis.GetClient().Get(context.Background(),)
	err = cached.RedisStore.Get(token, &userInfo)
	if err != nil {
		iotlogger.LogHelper.Info("cached.RedisStore.Get error " + err.Error())
		code = ioterrs.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
		//缓存拿不到，拿token临时验证
		// resp, err := rpc.ClientCloudAuthService.VerifyToken(context.Background(), &protosService.VerifyTokenRequest{Token: token})
		// if err == nil && resp != nil && resp.Valid {
		// 	code = iotconst.SUCCESS
		// 	userInfo = UserInfo{
		// 		UserID:   resp.UserInfo.UserId,
		// 		Nickname: resp.UserInfo.NickName,
		// 		Avatar:   resp.UserInfo.Avatar,
		// 		DeptId:   resp.UserInfo.DeptId,
		// 		RoleIds:  resp.UserInfo.RoleIds,
		// 		PostIds:  resp.UserInfo.PostIds,
		// 		TenantId: resp.UserInfo.TenantId,
		// 		//TenantId: resp.UserInfo.
		// 	}
		// 	expires := time.Unix(resp.ExpiresAt, 0).Sub(time.Now())
		// 	err = cached.RedisStore.Set(token, userInfo, expires)
		// 	if err != nil {
		// 		iotlogger.LogHelper.Errorf("UserLogin,缓存token失败:%s", err.Error())
		// 	}
		// }
	}
	if code.Code != ioterrs.SUCCESS.Code {
		iotgin.ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}

	c.Set("userId", userInfo.UserID)
	c.Set("nickName", userInfo.Nickname)
	c.Set("avatar", userInfo.Avatar)
	c.Set("deptId", userInfo.DeptId)
	c.Set("roleIds", userInfo.RoleIds)
	c.Set("postIds", userInfo.PostIds)
	c.Set("tenantId", userInfo.TenantId)
	c.Set("accountType", userInfo.AccountType)
	c.Set("Token", token)

	c.Next()
}

func GetUserId(c *gin.Context) int64 {
	return c.GetInt64("userId")
}

func GetTenantId(c *gin.Context) string {
	return c.GetString("tenantId")
}

func GetLang(c *gin.Context) string {
	return c.GetHeader("lang")
}

// 获取userAgent信息
func GetUserAgent(c *gin.Context) (string, string, string) {
	userAgent := c.Request.UserAgent()
	ua := user_agent.New(userAgent)
	os := ua.OS()
	browserName, browserVersion := ua.Browser()
	return os, browserName, browserVersion
}

// 将用户信息和token传给后端微服务
func WithUserContext(c *gin.Context) context.Context {
	ctx := metadata.NewContext(context.Background(),
		map[string]string{
			"userId":   iotutil.ToString(c.GetInt64("userId")),
			"tenantId": c.GetString("tenantId"),
			"lang":     c.GetHeader("lang"),
			"token":    GetToken(c),
		})
	return ctx
}

// 将用户信息和token传给后端微服务
func WithOpenUserContext(c *gin.Context) context.Context {
	ctx := metadata.NewContext(context.Background(),
		map[string]string{
			"userId":   iotutil.ToString(c.GetInt64("userId")),
			"tenantId": c.GetString("tenantId"),
			"lang":     c.GetHeader("lang"),
			"token":    GetToken(c),
		})
	return ctx
}

func GetAppKeyByHost(c *gin.Context) string {
	host := c.Request.Host
	arr := strings.Split(host, ".")
	//三级域名
	appKey := ""
	if len(arr) > 3 {
		appKey = arr[0]
	}
	iotlogger.LogHelper.Info("host: ", host, ", appKey:", appKey)
	return appKey
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
