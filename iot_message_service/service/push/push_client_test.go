package push

import (
	"cloud_platform/iot_common/iotlogger"
	"testing"

	"github.com/goccy/go-json"
)

func TestPush(t *testing.T) {
	//统一日志到服务的日志
	err := iotlogger.InitLog("./logs/iot_message_service.log", "iot_message_service", "debug")
	if err != nil {
		panic(err)
	}

	client, _ := NewPushClient(All)
	var pushInfo PushInfo
	err = json.Unmarshal([]byte("{\"InputTarget\":{\"isPublic\":false,\"tags\":null,\"alias\":[\"9177311063388028928\"],\"regIds\":null,\"pushTokens\":null},\"Message\":{\"messageId\":\"5598316778158456832\",\"type\":\"home\",\"childtype\":\"10\",\"objecttype\":\"\",\"objectid\":\"\",\"model\":\"\",\"devimg\":\"\",\"homename\":\"\",\"title\":\"家庭被移除（不删）\",\"content\":\"“hogan”移除了家庭“uieieiei”\",\"devids\":\"\",\"extands\":null,\"url\":\"\",\"params\":null,\"isread\":false,\"unsetexpire\":false,\"ispublic\":0,\"tagtype\":0,\"homeid\":\"5581377335392632832\",\"userid\":\"[9177311063388028928]\",\"userids\":null,\"result\":\"\",\"createtime\":1692762522,\"expiretime\":1692766122,\"created_at\":1692762522,\"creator_id\":\"\",\"creator_name\":\"\",\"appKey\":\"J509kB6Dw3AWxJxTpfxyP01e\",\"tenantId\":\"ioqp4r\",\"tplCode\":\"RemoveHomeMessage\"}}"), &pushInfo)
	if err != nil {
		t.Logf("PushClientMgr.SendPush error:%s", err.Error())
	}
	err = client.PushMessage(pushInfo.InputTarget, pushInfo.Message)
	if err != nil {
		t.Logf("PushClientMgr.SendPush error:%s", err.Error())
	}
}
