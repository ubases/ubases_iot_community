package iotutil

import (
	"bytes"
	"html/template"
	txttemplate "text/template"
)

func RenderHtmpTemplateByPath(templatePath string, data interface{}) (string, error) {
	tmp, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderHtmpTemplate(templateStr string, data interface{}) (string, error) {
	tmp, err := template.New("").Parse(templateStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderTextTemplate(templateStr string, data interface{}) (string, error) {
	tmp, err := txttemplate.New("").Parse(templateStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderTextTemplateByPath(templatePath string, data interface{}) (string, error) {
	tmp, err := txttemplate.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = tmp.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
