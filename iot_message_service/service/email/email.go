package email

import (
	"bytes"
	"errors"
	"html/template"
)

// 邮件接口
type Email interface {
	Send(input SendEmailInput) (bool, error)
}

// 邮件详细输入参数
type SendEmailInput struct {
	To      string
	Subject string
	Body    string
}

// 从模板文件htmp文件加载
func (s *SendEmailInput) GenerateBodyFromFile(templateFileName string, data interface{}) error {
	tmp, err := template.ParseGlob(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return err
	}

	s.Body = buf.String()

	return nil
}

// 从文本内容加载
func (s *SendEmailInput) GenerateBodyFromContent(templateContent string, data interface{}) error {
	tmp, err := template.New("Email").Parse(templateContent)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return err
	}
	s.Body = buf.String()
	return nil
}

// 邮件输入参数验证
func (e *SendEmailInput) Validate() error {
	if e.To == "" {
		return errors.New("empty to")
	}
	if e.Subject == "" || e.Body == "" {
		return errors.New("empty subject/body")
	}
	if !IsEmailValid(e.To) {
		return errors.New("invalid to email")
	}
	return nil
}
