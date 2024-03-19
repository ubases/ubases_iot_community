package gorushclient

import (
	"bytes"
	"cloud_platform/iot_common/iotlogger"
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

func NewGorushClient(baseURL string) *GorushClient {
	return &GorushClient{BaseURL: baseURL}
}

// 发送推送消息
func (gc *GorushClient) SendPush(notification PushNotificationRequest) error {
	payload, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	iotlogger.LogHelper.Info("Gorush参数：", iotutil.ToString(payload))
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
