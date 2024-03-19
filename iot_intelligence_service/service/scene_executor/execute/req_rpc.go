package execute

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_intelligence_service/rpc/rpcclient"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

func GetUsersByHomeId(homeId int64) (*protosService.UcHomeUsersResponse, error) {
	//TODO 推送给家庭成员，需要通过家庭ID获取家庭成员人员列表
	iotlogger.LogHelper.Info("==========>", homeId)
	users, err := rpcclient.ClientAppHomeUser.GetUsersByHomeId(context.Background(), &protosService.UcHomeUserPrimarykey{
		Id: homeId,
	})
	if err != nil {
		iotlogger.LogHelper.Error("==========>", err)
		return nil, err
	}
	if users.Code != 200 {
		iotlogger.LogHelper.Error("==========>未获取到家庭用户")
		return nil, errors.New("未获取到家庭用户")
	}
	return users, nil
}

func SendMessage(req *protosService.SendMessageRequest) error {
	//发送消息
	ret, err := rpcclient.ClientAppMessage.SendMessage(context.Background(), req)
	if err != nil {
		return err
	}
	if ret.Code == 200 {
		return errors.New(ret.Message)
	}
	return nil
}
