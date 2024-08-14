/**
 * @Author: hogan
 * @Date: 2022/3/18 9:24
 */

package iotsms

import (
	"cloud_platform/iot_common/iotutil"
)

type client struct {
	cfg *Config
}

var (
	ALI_SMS_PATFORM string = "alisms"
)

// New a client
// smsProvider 短信渠道
// 短信模板
func NewClient(smsProvider string) *client {
	cfg := &Config{
		SMSProvider:     smsProvider,
		RegionId:        "cn-hangzhou",
		AccessKey:       "LTAI4FueswDAttZRWSpC4FXF",
		AccessKeySecret: "och1oekoXH4hzdyA8PShQKzZAyJWiM",
		SignName:        "深圳市安信可科技有限公司",
	}
	return &client{cfg: cfg}
}

// Send message to mobile
func (c *client) Send(mobile string, template string, param map[string]interface{}) error {
	switch c.cfg.SMSProvider {
	case ALI_SMS_PATFORM:
		NewAliyunSms(c.cfg).Send(SmsContent{
			Phone:        mobile,
			TemplateCode: template, //短信模板
			Params:       iotutil.ToString(param),
		})
	}
	return nil
}
