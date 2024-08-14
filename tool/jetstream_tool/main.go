package main

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotnatsjs"
	"context"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"strings"
)

// 首次部署时执行或流名称有更改时执行
// 注意根据nats服务的部署配置修改addrs信息。
//var addrs string = "nats://iLmz8sCXjkTYuh@192.168.5.56:4222"

func main() {
	addrs := flag.String("addrs", "nats://{token}@127.0.0.1:4222", "nats连接URL")
	flag.Parse()

	client, err := iotnatsjs.NewJsClient(*addrs)
	if err != nil {
		fmt.Println("NewJsClient: error:", err.Error())
		return
	}
	var streamMap = make(map[string]bool)
	opts := []jetstream.StreamListOpt{}
	sinfo := client.Js.ListStreams(context.Background(), opts...)
	if sinfo.Err() == nil {
		for info := range sinfo.Info() {
			for _, v := range info.Config.Subjects {
				streamMap[v] = true
			}
		}
	}
	var streamList []string
	streamList = append(streamList, iotconst.NATS_STREAM_DEVICE+".>")
	streamList = append(streamList, iotconst.NATS_STREAM_APP+".>")
	streamList = append(streamList, iotconst.NATS_LANGUAGE+".>")
	streamList = append(streamList, iotconst.NATS_MESSAGE+".>")
	streamList = append(streamList, iotconst.NATS_PRODUCT_PUBLISH+".>")
	streamList = append(streamList, iotconst.NATS_BUILDAPP+".>")
	streamList = append(streamList, iotconst.HKEY_CACHED_CLEAR_PUB_PREFIX+".>")
	streamList = append(streamList, iotconst.NATS_STREAM_ORIGINAL_REDIS+".>")
	for _, v := range streamList {
		if _, ok := streamMap[v]; ok {
			continue
		}
		err = client.CreateOrUpdateStream(strings.ReplaceAll(v, ".>", ""), []string{v})
		if err != nil {
			fmt.Println(fmt.Sprintf("CreateOrUpdateStream: stream=%s,error:%s", v, err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("CreateOrUpdateStream: stream=%s,successed.", v))
		}
	}
	fmt.Println("stream update completed")
}
