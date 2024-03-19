package service

import (
	"reflect"
)

//变量定义
//变量来自于爱星云的各个业务逻辑中，需要新增带模板变量的消息类别，则需要增加对应的变量

// 短信和邮件用，验证码支持的变量
type CodeInput struct {
	UserName  string `json:"userName"`
	Code      string `json:"code"`
	PhoneType string `json:"phoneType"`
	Lang      string `json:"lang"`
	Template  string `json:"template"`
}

// 短信和邮件用，异地登录支持的变量
type LoggedInInput struct {
	UserName string `json:"userName"`
	IP       string `json:"IP"`
}

// 短信和邮件用，用户注册支持的变量
type RegisterInput struct {
	UserName string `json:"userName"`
}

// 仅短信用
func VariablesToMap(dat interface{}) map[string]string {
	obj := reflect.TypeOf(dat)
	val := reflect.ValueOf(dat)
	data := make(map[string]string)
	for i := 0; i < obj.NumField(); i++ {
		if val.Field(i).String() != "" {
			if obj.Field(i).Tag.Get("json") != "" {
				data[obj.Field(i).Tag.Get("json")] = val.Field(i).String()
			} else {
				data[obj.Field(i).Name] = val.Field(i).String()
			}
		}
	}
	return data
}
