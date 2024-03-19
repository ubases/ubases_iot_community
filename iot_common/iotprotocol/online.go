package iotprotocol

import json "github.com/json-iterator/go"

var ONLINE_HEAD_NS = "iot.device.report"
var ONLINE_HEAD_NAME = "online"

type PackOnline struct {
	Header  Header `json:"header"`
	Payload Online `json:"payload"`
}

type Online struct {
	OnlineStatus string `json:"onlineStatus"`
}

func (o *PackOnline) Encode(gid string, online bool) ([]byte, error) {
	var status string
	if online {
		status = "online"
	} else {
		status = "offline"
	}
	obj := PackOnline{
		Header:  EncodeHeader(ONLINE_HEAD_NS, ONLINE_HEAD_NAME, gid),
		Payload: Online{OnlineStatus: status},
	}
	*o = obj
	return json.Marshal(obj)
}
