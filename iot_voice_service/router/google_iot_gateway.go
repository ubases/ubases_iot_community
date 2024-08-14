package routers

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_voice_service/service"
	"context"
	"fmt"
	"strings"

	"github.com/qiniu/x/errors"

	"github.com/gin-gonic/gin"
)

func GoogleFulfillment(c *gin.Context) {
	//验证参数
	token, err := CheckRequest(c)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("check request parameters error: ", err)
		c.Abort()
		return
	}
	//验证token
	Id, err := CheckToken(token)
	if err != nil {
		iotlogger.LogHelper.Helper.Error("check token error: ", err)
		c.Abort()
		return
	}
	service.GoogleMain(c, Id)
}

func CheckRequest(c *gin.Context) (string, error) {
	//内容格式检查
	contentType := c.Request.Header.Get("content-type")
	if !strings.Contains(contentType, "application/json") {
		return "", errors.New("Request not JSON")
	}
	//认证头检查
	authHeader := c.Request.Header.Get("Authorization")
	if len(authHeader) < 1 {
		return "", errors.New("Access Token Required")
	}
	authTokenParts := strings.Split(authHeader, " ")
	if len(authTokenParts) != 2 || strings.ToLower(authTokenParts[0]) != "bearer" {
		return "", errors.New("Access Token Must Be Bearer")
	}
	return authTokenParts[1], nil
}

func CheckToken(token string) (string, error) {
	ti, err := manager.LoadAccessToken(context.Background(), token)
	if err != nil {
		return "", fmt.Errorf("load access token error: %v", err.Error())
	}
	return ti.GetUserID(), nil
}
