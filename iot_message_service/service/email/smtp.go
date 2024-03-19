package email

import (
	"strings"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// SMTP客户端结构.
type Client struct {
	client *mail.SMTPClient
	cfg    *SMTPConfig
}

// SMTP配置
type SMTPConfig struct {
	Host           string
	Port           int
	Username       string
	Password       string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
	Helo           string
	KeepAlive      bool
	Exchange       bool
	AuthType       int
	Ssl            int
	From           string
}

// 创建SMTP客户端.
func NewClient(cfg *SMTPConfig) (*Client, error) {
	server := &mail.SMTPServer{
		Host:           cfg.Host,
		Port:           cfg.Port,
		Username:       cfg.Username,
		Password:       cfg.Password,
		ConnectTimeout: cfg.ConnectTimeout,
		SendTimeout:    cfg.SendTimeout,
		Helo:           cfg.Helo,
		KeepAlive:      cfg.KeepAlive,
		Authentication: mail.AuthType(cfg.AuthType),
		Encryption:     mail.Encryption(cfg.Ssl),
	}
	if cfg.Exchange { //如果配置的是邮箱地址
		list := strings.Split(cfg.Username, "@")
		if len(list) == 2 {
			cfg.Username = list[0]
		}
		server.Helo = ""
	}
	client, err := server.Connect()
	if err != nil {
		return nil, err
	}
	if err = client.Noop(); err != nil {
		return nil, err
	}
	return &Client{client: client, cfg: cfg}, nil
}

// 发送邮件
func (c *Client) Send(input SendEmailInput) (bool, error) {
	if err := input.Validate(); err != nil {
		return false, err
	}
	msg := mail.NewMSG()
	from := c.cfg.From
	if c.cfg.Exchange && c.cfg.Helo != "" {
		from = c.cfg.Helo + "<" + c.cfg.From + ">"
	}
	msg.SetFrom(from)
	msg.AddTo(input.To)
	msg.SetSubject(input.Subject)
	msg.SetBody(mail.TextHTML, input.Body)
	if msg.Error != nil {
		return false, msg.Error
	}
	if err := msg.Send(c.client); err != nil {
		return false, err
	}
	return true, nil
}

func (c *Client) Close() {
	c.client.Quit()
}

func (c *Client) Noop() error {
	return c.client.Noop()
}
