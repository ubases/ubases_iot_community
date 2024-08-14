package wechatclient

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_message_service/config"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
)

var MiniProgramApp *miniProgram.MiniProgram

func init() {
	var (
		err            error
		appId          = config.Global.ThirdPartyLogin.MiniProgram.AppId
		appSecret      = config.Global.ThirdPartyLogin.MiniProgram.AppId
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
		iotlogger.LogHelper.Error("MiniProgramApp", err.Error())
	}
}

type WechatClient struct {
	BaseURL string
}

type PushNotificationRequest struct {
	Notifications []PushNotification `json:"notifications"`
}

type PushNotification struct {
	Tokens           []string `json:"tokens" binding:"required"`
	Platform         int      `json:"platform" binding:"required"`
	Message          string   `json:"message,omitempty"`
	Topic            string   `json:"topic,omitempty"`
	Title            string   `json:"title,omitempty"`
	Credentials      string   `json:"credentials,omitempty"` //证书文件
	CredentialSecret string   `json:"credentialSecret"`      //证书密钥
	AppKey           string   `json:"appKey,omitempty"`
	Data             D        `json:"data,omitempty"`
}

// D provide string array
type D map[string]interface{}

type PushResponse struct {
	Counts  int           `json:"counts"`
	Logs    []interface{} `json:"logs"`
	Success string        `json:"success"`
}

type NoticeRequest struct {
	Inner      string //内部应用标识
	OpenId     string
	TemplateId string
	PageCode   string
	UserType   int32
	Data       *SdkRequest
}

func NewWechatClient(baseURL string) *WechatClient {
	return &WechatClient{BaseURL: baseURL}
}

// 发送推送消息req NoticeRequest
func (gc *WechatClient) SendPush(req []NoticeRequest) error {
	for _, d := range req {
		resData := make(power.HashMap)
		for k, v := range d.Data.GetGMap() {
			resData[k] = &power.HashMap{
				"value": v,
				"color": "#000",
			}
		}
		sendData := &request.RequestSubscribeMessageSend{
			ToUser:           d.OpenId,
			TemplateID:       d.TemplateId,
			Page:             d.PageCode,
			MiniProgramState: "trial",
			Lang:             "zh_CN",
			Data:             &resData,
		}
		iotlogger.LogHelper.Info("MiniProgramApp.SubscribeMessage.Send", iotutil.ToString(sendData))
		resp, err := MiniProgramApp.SubscribeMessage.Send(context.Background(), sendData)
		if err != nil {
			iotlogger.LogHelper.Info(err.Error())
			return err
		}
		if resp.ErrCode != 0 {
			iotlogger.LogHelper.Info(iotutil.ToString(resp))
			return err
		}
	}
	return nil
}
