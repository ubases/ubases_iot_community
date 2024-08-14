package gorushclient

import (
	"bytes"
	"cloud_platform/iot_common/iotutil"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
)

type GorushClient struct {
	BaseURL string
}

const (
	// PlatFormIos constant is 1 for iOS
	PlatformIos = 1
	// PlatFormAndroid constant is 2 for Android
	PlatformAndroid = 2
	// PlatFormHuawei constant is 3 for Huawei
	PlatformHuawei = 3
	// PlatFormXiaomi constant is 3 for Xaomi
	PlatformXiaomi = 4
	// PlatFormOppo constant is 3 for Oppo
	PlatformOppo = 5
	// PlatFormVivo constant is 3 for Vivo
	PlatformVivo = 6
	// PlatFormHoner constant is 3 for Honer
	PlatformHoner = 7
	// PlatFormMeizu constant is 3 for Meizu
	PlatformMeizu = 8
)

type PushNotificationRequest struct {
	Notifications []PushNotification `json:"notifications"`
}

type PushNotification struct {
	ID               string   `json:"notif_id,omitempty"`
	Tokens           []string `json:"tokens" binding:"required"`
	Platform         int      `json:"platform" binding:"required"`
	Message          string   `json:"message,omitempty"`
	Topic            string   `json:"topic,omitempty"`
	Title            string   `json:"title,omitempty"`
	Credentials      string   `json:"credentials,omitempty"` //证书文件
	CredentialSecret string   `json:"credentialSecret"`      //证书密钥
	AppKey           string   `json:"appKey,omitempty"`
	Data             D        `json:"data,omitempty"`
	AppID            string   `json:"app_id,omitempty"`
	AppSecret        string   `json:"app_secret,omitempty"`
	HuaweiData       string   `json:"huawei_data,omitempty"`
	Intent           string   `json:"intent,omitempty"` //自定义的intent参数
	Color            string   `json:"color,omitempty"`
	ClickAction      string   `json:"click_action,omitempty"` //%v.MainActivity
	PkgName          string   `json:"pkg_name,omitempty"`     //包名

	Category  string `json:"category,omitempty"`
	ChannelId string `json:"channel_id"`

	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`

	//小米
	XiaomiChanelId     int `json:"xiaomi_chanel_id,omitempty"`
	XiaomiNotifyEffect int `json:"extra_notify_effect"` //预定义通知栏消息的点击行为 https://dev.mi.com/distribute/doc/details?pId=1559
	XiaomiNotifyType   int `json:"xiaomi_notify_type"`  // notify_type的值可以是DEFAULT_ALL或者以下其他几种的OR组合： DEFAULT_ALL = -1;DEFAULT_SOUND = 1; // 使用默认提示音提示；DEFAULT_VIBRATE = 2; // 使用默认振动提示；DEFAULT_LIGHTS = 4; // 使用默认呼吸灯提示。

	//VIVO
	VivoAddBadge bool `json:"vivo_add_badge"` //角标
}

// D provide string array
type D map[string]interface{}

type PushResponse struct {
	Counts  int           `json:"counts"`
	Logs    []interface{} `json:"logs"`
	Success string        `json:"success"`
}

func NewGorushClient(baseURL string) *GorushClient {
	return &GorushClient{BaseURL: baseURL}
}

// 发送推送消息
func (gc *GorushClient) SendPush(notification PushNotificationRequest) error {
	payload, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/api/push", gc.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var pushResp PushResponse
	err = json.NewDecoder(resp.Body).Decode(&pushResp)
	if err != nil {
		return err
	}

	// 处理推送响应4
	if pushResp.Success != "ok" {
		return errors.New(iotutil.ToString(pushResp.Logs))
	}
	return nil
}
