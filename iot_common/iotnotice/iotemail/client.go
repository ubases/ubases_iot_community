/**
 * @Author: hogan
 * @Date: 2022/3/18 9:24
 */

package iotemail

import (
	"errors"
	"github.com/go-gomail/gomail"
)

type client struct {
	EmailServer string
	EmailPort   int
	FromEmail   string
	FromPasswd  string
	ToEmail     string
}

var (
	prefix string = "【爱星云】" // 邮件前缀标识
)

// New a client
func NewClient(smsProvider, appCode, templateCode string) *client {
	return &client{
		EmailServer: "imap.exmail.qq.com", // 腾讯企业邮箱服务器 地址
		EmailPort:   465,                  // 腾讯企业邮箱服务器 端口
		FromEmail:   "hejy@tech-now.com",
		FromPasswd:  "xxxxxx",
	}
}

// Send message to mobile
func (c *client) Send(email, subject, body string) error {
	c.SendEmail(email, subject, body)
	return nil
}

// execute mail sending
// subject：邮件主题
// body：邮件内容，支持html格式字符串
func (c *client) SendEmail(email, subject, body string) error {
	//初始化调剖箭头
	m, err := c.SetToUsers(prefix, email)
	if err != nil {
		return err
	}
	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(c.EmailServer, c.EmailPort, c.FromEmail, c.FromPasswd)
	// 发送
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}

// set email recipient and email sender
func (c *client) SetToUsers(templateCode string, toUsers ...string) (m *gomail.Message, err error) {
	m = gomail.NewMessage()
	if len(toUsers) == 0 {
		err = errors.New("no users")
		return
	}
	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toUsers...)
	// 发件人
	m.SetAddressHeader("From", c.FromEmail, templateCode)
	return
}
