package iotgin

import (
	"cloud_platform/iot_common/ioterrs"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserClaims struct {
	UserID          string `json:"UserID"`
	HomeID          string `json:"HomeID"`
	Account         string `json:"Account"`
	Name            string `json:"Name"`
	DevControlSecrt string `json:"DevControlSecrt"`
	//jwt.StandardClaims
}

// 路由过滤器方法，包括token解析和验证
func AuthCheck(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Request.Header.Get("token")
	}
	if authHeader == "" {
		ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}

	token := parts[1]
	if token == "" {
		ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}
	var code ioterrs.KVCode
	//解析token，调用gRPC授权服务方法 ParseToken(token) jwtserct string,
	claims := UserClaims{
		UserID: "test",
	}
	var err error
	if err != nil {
		// 错误返回
		code = ioterrs.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	} else {
		c.Set("UserID", claims.UserID)
		c.Set("Account", claims.Account)
		c.Set("HomeID", claims.HomeID)
		c.Set("Name", claims.Name)
		c.Set("DevSecrt", claims.DevControlSecrt)
		c.Set("Token", token)
	}

	if code.Code == ioterrs.SUCCESS.Code {
		ResFailCode(c, ioterrs.INVALID_PARAMS.Msg, ioterrs.INVALID_PARAMS.Code)
		c.Abort()
		return
	}

	c.Next()
}
