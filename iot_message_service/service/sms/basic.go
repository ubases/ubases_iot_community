package sms

import "fmt"

const (
	Aliyun       = "Aliyun"
	TencentCloud = "Tencent"
	VolcEngine   = "VolcEngine"
	Huyi         = "Huyi"
	HuaweiCloud  = "Huawei"
)

type SmsClient interface {
	SendMessage(template string, param map[string]string, targetPhoneNumber ...string) error
}

func NewSmsClient(provider string, accessId, accessKey string, sign map[string]string, other ...string) (SmsClient, error) {
	switch provider {
	case Aliyun:
		return GetAliyunClient(accessId, accessKey, sign, other)
	case TencentCloud:
		return GetTencentClient(accessId, accessKey, sign, other)
	case VolcEngine:
		return GetVolcClient(accessId, accessKey, sign, other)
	case Huyi:
		return GetHuyiClient(accessId, accessKey)
	case HuaweiCloud:
		return GetHuaweiClient(accessId, accessKey, sign, other)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}
}
