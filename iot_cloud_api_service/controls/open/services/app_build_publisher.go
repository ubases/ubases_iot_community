package services

import (
	"cloud_platform/iot_cloud_api_service/config"
	"cloud_platform/iot_common/iotnats/jetstream"
	"cloud_platform/iot_common/iotstruct"
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go"
)

// 延用APP构建的topic
const BUILD_PANEL = "buildapp"

func PulisherPanelBuildMessage(bi iotstruct.BuildInfo) error {

	addrs := config.Global.Nats.Addrs //[]string{"nats://mytoken@120.77.96.118:4222", "nats://mytoken@120.77.64.252:4222", "nats://mytoken@120.77.61.157:4222"}
	p, err := jetstream.NewJSPublisher("build_app_message", BUILD_PANEL, BUILD_PANEL+".android", connerrhandler, addrs...)
	if err != nil {
		return err
	}
	//2 发布消息
	buf, err := json.Marshal(bi)
	if err != nil {
		return err
	}

	err = p.PublishEx([]byte(buf), handler)
	if err != nil {
		//log.Println(err.Error())
	} else {
		//log.Println("消息发送成功:", buf)
	}
	time.Sleep(5 * time.Second)
	//3 当不需要时关闭
	p.Close()
	return nil
}

func connerrhandler(conn *nats.Conn, err error) {

}
func handler(msg *nats.Msg, err error) {

}
