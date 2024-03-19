package push

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_message_service/service/push/jpushclient"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPushClear(t *testing.T) {
	jpush := jpushclient.NewPushClient("b25919ee51cf1f82f544b0c7", "2db568da131d6c904c31174f")
	//url := fmt.Sprintf("https://device.jpush.cn/v3/aliases/%s?platform=android,ios,quickapp&new_format=true", "1540232140371886080_zh")
	//查询别名
	resStr, err := jpush.SendGetAliasesRequest("1540232140371886080_zh")
	if err != nil {
		t.Error(fmt.Sprintf("err:%s", err.Error()))
	}
	if err != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
		return
	}
	var res jpushclient.AliasesResponse
	err = json.Unmarshal([]byte(resStr), &res)
	if err != nil {
		iotlogger.LogHelper.Info(fmt.Sprintf("err:%s", err.Error()))
		return
	}
	if res.Data != nil {
		if len(res.Data) > 6 {
			newData := res.Data[5:len(res.Data)]
			if len(newData) == 0 {
				return
			}
			//清除别名
			alias := jpushclient.NewAliases()
			for _, d := range newData {
				alias.SetRegIds(d.RegistrationId)
			}
			bytes, _ := alias.ToBytes()
			url := fmt.Sprintf("https://device.jpush.cn/v3/aliases/%s", "1540232140371886080_zh")
			jpush.BaseUrl = url
			str, err := jpush.Send(bytes)
			fmt.Println(str, err)
		}
	}
	t.Log(resStr)
}
