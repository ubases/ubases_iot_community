package iotprotocol

import (
	"time"

	json "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

var CONTROL_HEAD_NS = "iot.device.control"
var CONTROL_HEAD_NAME = "control"         //设备控制
var CONTROL_HEAD_NAME_RESTORE = "restore" //恢复出厂设置
var CONTROL_HEAD_NAME_REBOOT = "reboot"   //设备重启

type PackControl struct {
	Header  Header  `json:"header"`
	Payload Control `json:"payload"`
}

func (o *PackControl) Encode(gid, name string, data map[string]interface{}) ([]byte, error) {
	obj := PackControl{
		Header:  EncodeHeader(CONTROL_HEAD_NS, name, gid),
		Payload: Control{CtrlData: data},
	}
	*o = obj
	return json.Marshal(obj)
}

type PackControlAck struct {
	Header  Header    `json:"header"`
	Payload NormalAck `json:"payload"`
}

type Control struct {
	CtrlData map[string]interface{} `json:"control,omitempty"`
}

func (o *PackControlAck) Encode(gid, name, mid string, err error) ([]byte, error) {
	if mid == "" {
		mid = uuid.NewV4().String()
	}
	obj := PackControlAck{
		Header:  Header{Ns: CONTROL_HEAD_NS, Name: name, Mid: mid, Ts: time.Now().UTC().Unix(), Ver: "1.0.0", Gid: gid},
		Payload: NormalAck{Code: 0},
	}
	if err != nil {
		obj.Payload.Code = 1
	}
	*o = obj
	return json.Marshal(obj)
}
