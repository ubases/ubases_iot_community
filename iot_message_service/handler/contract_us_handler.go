package handler

import (
	"cloud_platform/iot_message_service/config"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"

	"github.com/cyanBone/dingtalk_robot"
	"github.com/cyanBone/dingtalk_robot/message"
)

type ContractUsServiceHandler struct{}

func (c ContractUsServiceHandler) PushContractUsNotice(ctx context.Context, request *proto.ContractUsNoticeRequest, response *proto.Response) error {
	//推送通知
	var title = fmt.Sprintf("【%s】收到官网发来的联系通知!", config.Global.Notice.DingTalk.Keyword)
	var text = fmt.Sprintf(`##### 【官网联系通知】您好！收到官网发来的联系通知，请尽快跟进！  
**姓名：**%s  
**电话：**%s   
**公司名称：**%s  
**邮件：**%s  
**内容：** %s`, request.Name, request.Phone, request.CompanyName, request.Email, request.Content)

	err := sendDingtalkMarkdownMsg(title, text)
	if err != nil {
		response.Code = ERROR
		response.Message = err.Error()
	} else {
		response.Code = SUCCESS
		response.Message = "success"
	}
	return nil
}

// 联系我们钉钉通知
func sendDingtalkMarkdownMsg(title, text string) error {
	//webhook 机器人地址
	webhook := config.Global.Notice.DingTalk.Webhook
	//机器人密钥，也可以设置为关键字
	secert := config.Global.Notice.DingTalk.Secert

	//初始化客户端
	client, err := dingtalk_robot.New(webhook, secert)
	if err != nil {
		return err
	}

	//普通文本信息
	textMessage := message.NewMarkdownMessage()
	textMessage.Title = title
	textMessage.Text = text
	//@对应的人
	//textMessage.AtMobiles([]string{"1111111111"})
	//@所有人
	textMessage.AtAll(true)

	//发送信息
	err = client.Send(textMessage)
	if err != nil {
		return err
	}
	return nil
}
