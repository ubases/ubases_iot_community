package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

type LoginInput struct {
	Password string `json: "Password"` //密码
	Account  string `json: "Acount"`   //登录参数
	Type     int
}

var jwtSecret []byte

func NewUserApi() *UserApi {
	s := &UserApi{}
	return s
}

// 调用rpc
func (this *UserApi) Auth(c *gin.Context) (map[string]interface{}, error) {
	username := c.Request.Form.Get("username")
	password := c.Request.Form.Get("password")
	regionServerId := c.Request.Form.Get("region")
	if regionServerId == "" {
		regionServerId = "1"
	}
	appkey := strings.Split(c.Request.Host, ".")[0]
	iotlogger.LogHelper.Helper.Debugf("username: %v, password: %v, appkey: %v, region: %v", username, password, appkey, regionServerId)
	if !(iotutil.CheckAllPhone("", username) || iotutil.IsEmail(username)) || len(password) == 0 {
		return nil, errors.New("账号和密码不合法或者为空")
	}
	res, err := rpcclient.AppAuthService.PasswordLogin(context.Background(), &protosService.PasswordLoginRequest{
		Channel:        "1",
		LoginName:      username,
		Password:       password,
		AppKey:         appkey,
		RegionServerId: iotutil.ToInt64(regionServerId),
	})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"access_token":  res.Token,
		"token_type":    "Bearer",
		"expires_in":    res.ExpiresAt,
		"refresh_token": res.RefreshToken,
		"scope":         "all",
		"userid":        iotutil.ToString(res.UserInfo.Id),
	}, nil
}
