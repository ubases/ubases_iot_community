package sms

import (
	"bytes"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliyunClient struct {
	sendId string
	sign   map[string]string
	core   *dysmsapi.Client
}

func GetAliyunClient(accessId, accessKey string, sign map[string]string, other []string) (*AliyunClient, error) {
	var sendId string
	if len(other) != 0 {
		sendId = other[0]
	}
	region := "cn-hangzhou"
	client, err := dysmsapi.NewClientWithAccessKey(region, accessId, accessKey)
	if err != nil {
		return nil, err
	}

	aliyunClient := &AliyunClient{
		core:   client,
		sendId: sendId,
		sign:   sign,
	}

	return aliyunClient, nil
}

func (c *AliyunClient) SendMessage(templateCode string, param map[string]string, targetPhoneNumber ...string) error {
	lang := param["lang"]
	iotlogger.LogHelper.Helper.Debugf("templateCode: %v, param: %v, targetphone: %v, sign: %v", templateCode, param, targetPhoneNumber, c.sign[lang])
	if param["phoneType"] == "1" {
		delete(param, "template")
		delete(param, "phoneType")
		delete(param, "lang")
		delete(param, "recordId")
		requestParam, err := json.Marshal(param)
		if err != nil {
			return err
		}

		if len(targetPhoneNumber) < 1 {
			return fmt.Errorf("missing parameter: targetPhoneNumber")
		}

		phoneNumbers := bytes.Buffer{}
		phoneNumbers.WriteString(targetPhoneNumber[0])
		for _, s := range targetPhoneNumber[1:] {
			phoneNumbers.WriteString(",")
			phoneNumbers.WriteString(s)
		}

		request := dysmsapi.CreateSendSmsRequest()
		request.Scheme = "https"
		request.PhoneNumbers = phoneNumbers.String()
		request.TemplateCode = templateCode
		request.TemplateParam = string(requestParam)
		request.SignName = c.sign[lang]

		res, err := c.core.SendSms(request)
		if err != nil {
			return err
		}
		iotlogger.LogHelper.Info("SendMessage: ", iotutil.ToString(res))
		if res.Code != "OK" {
			return errors.New(res.Message + ", 错误码为" + res.Code)
		}
		//成功
		//{\"RequestId\":\"C9D6B9EF-F24B-58D7-81D5-B175AECA3A2F\",\"BizId\":\"660019421897908596^0\",
		//\"Code\":\"OK\",\"Message\":\"OK\"}
		//失败
		//{\"RequestId\":\"A8D01564-9E63-57D7-801E-5F93100F76A3\",
		//\"BizId\":\"\",\"Code\":\"isv.SMS_TEMPLATE_ILLEGAL\",\"Message\":\"该账号下找不到对应模板\"}
	} else {
		tmpl, err := template.New("test").Parse(param["template"])
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		codeParam := struct {
			Code string `json:"code"`
		}{
			Code: param["code"],
		}
		if err := tmpl.Execute(&buf, codeParam); err != nil {
			return err
		}
		tempStr := buf.String()
		request := requests.NewCommonRequest()
		request.Method = "POST"
		request.Scheme = "https" // https | http
		request.Domain = "dysmsapi.aliyuncs.com"
		request.Version = "2017-05-25"
		request.ApiName = "SendMessageToGlobe"
		request.QueryParams["RegionId"] = "cn-hangzhou"
		request.QueryParams["To"] = targetPhoneNumber[0]
		if v, ok := c.sign[lang]; ok && v != "" {
			//国际短信的签名配置需要增加【】、[] ，因申请模板可根据用户自行选择
			request.QueryParams["Message"] = c.sign[lang] + tempStr //"[" + c.sign[lang] + "] "
		} else {
			//阿里云那边提供的10dlc，无法使用签名进行发送
			request.QueryParams["Message"] = tempStr //"【" + c.sign[lang] + "】" +
		}
		request.QueryParams["From"] = c.sendId // 阿里云10DLC认证后，分配给你的11位数字长码，用于表明短信是谁发送的
		request.QueryParams["TaskId"] = iotutil.GetRandomNumber(11)
		request.QueryParams["Type"] = "OTP" // 短信类型，验证码OTP，通知NOTIFY，营销MKT，这个必须要写。

		res, err := c.core.ProcessCommonRequest(request)
		if err != nil {
			return err
		}
		iotlogger.LogHelper.Helper.Debugf("sign: %v, message: %v, res: %v", c.sign[lang], request.QueryParams["Message"], iotutil.ToString(res))
	}

	return nil
}
