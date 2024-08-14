package service

import (
	"cloud_platform/iot_app_user_service/config"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/gogf/gf/frame/g"
)

var MiniProgramApp *miniProgram.MiniProgram

func InitWechat() {
	var (
		err            error
		appId          = config.Global.ThirdPartyLogin.MiniProgram.AppId
		appSecret      = config.Global.ThirdPartyLogin.MiniProgram.AppSecret
		redis_host     = config.Global.Redis.Addrs
		redis_db       = config.Global.Redis.Database
		redis_password = config.Global.Redis.Password
	)
	//TODO 放入配置文件
	MiniProgramApp, err = miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     appId,     // 小程序appid
		Secret:    appSecret, // 小程序app secret
		HttpDebug: true,
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    redis_host,
			DB:       redis_db,
			Password: redis_password,
		}),
	})
	if err != nil {
		g.Log().Error("MiniProgramApp", err.Error())
	}
}

func GetMiniProgram(request *proto.MiniProgramLoginRequest) (channelUserId string, channelNickname string, msgCode int, msg string) {
	if request.Code == "" {
		msgCode = 100024
		msg = "授权Code为空"
		return
	}

	info, err := MiniProgramApp.Auth.Session(context.Background(), request.Code)
	if err != nil {
		msgCode = 100024
		msg = err.Error()
		return
	}
	if err != nil && info.OpenID == "" && info.UnionID == "" {
		msgCode = 100024
		msg = "获取openId、unionID失败"
		return
	}
	channelUserId = info.OpenID
	if len(info.OpenID) > 5 {
		l := len(info.OpenID)
		channelNickname = "wx-" + info.OpenID[l-5:]
	} else {
		channelNickname = "wx-" + info.OpenID
	}
	return
}
