package sms

import (
	"fmt"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

type TencentClient struct {
	core  *sms.Client
	appId string
	sign  map[string]string
}

func GetTencentClient(accessId string, accessKey string, sign map[string]string, appId []string) (*TencentClient, error) {
	if len(appId) < 1 {
		return nil, fmt.Errorf("missing parameter: appId")
	}

	credential := common.NewCredential(accessId, accessKey)
	config := profile.NewClientProfile()
	config.HttpProfile.ReqMethod = "POST"

	region := "ap-guangzhou"
	client, err := sms.NewClient(credential, region, config)
	if err != nil {
		return nil, err
	}

	tencentClient := &TencentClient{
		core:  client,
		appId: appId[0],
		sign:  sign,
	}

	return tencentClient, nil
}

func (c *TencentClient) SendMessage(template string, param map[string]string, targetPhoneNumber ...string) error {
	var paramArray []string
	index := 0
	for {
		value := param[strconv.Itoa(index)]
		if len(value) == 0 {
			break
		}
		paramArray = append(paramArray, value)
		index++
	}

	request := sms.NewSendSmsRequest()
	request.SmsSdkAppid = common.StringPtr(c.appId)
	request.Sign = common.StringPtr(c.sign["zh"])
	request.TemplateParamSet = common.StringPtrs(paramArray)
	request.TemplateID = common.StringPtr(template)
	request.PhoneNumberSet = common.StringPtrs(targetPhoneNumber)

	_, err := c.core.SendSms(request)
	return err
}
