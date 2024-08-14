/**
 * @Author: hogan
 * @Date: 2022/3/23 20:04
 */

package iotsms

import (
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliyunSms struct {
	cfg    *Config
	client *dysmsapi.Client
}

func NewAliyunSms(cfg *Config) *AliyunSms {
	client, err := dysmsapi.NewClientWithAccessKey(cfg.RegionId, cfg.AccessKey, cfg.AccessKeySecret)
	if err != nil {
		return nil
	}
	return &AliyunSms{cfg, client}
}

// 发送短信
func (sms *AliyunSms) Send(data SmsContent) map[string]interface{} {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = data.Phone
	req.SignName = sms.cfg.SignName
	req.TemplateCode = data.TemplateCode
	req.TemplateParam = data.Params
	resp, _ := sms.client.SendSms(req)
	var m map[string]interface{}
	json.Unmarshal(resp.GetHttpContentBytes(), &m)
	return m
}
