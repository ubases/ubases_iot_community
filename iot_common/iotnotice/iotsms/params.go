/**
 * @Author: hogan
 * @Date: 2022/3/23 20:02
 */

package iotsms

type Config struct {
	SMSProvider     string
	RegionId        string
	AccessKey       string
	AccessKeySecret string
	SignName        string
}

type SmsContent struct {
	Phone        string
	TemplateCode string // 短信模板
	Params       string // 短信变量
}
