package iotprotocol

import json "github.com/json-iterator/go"

var CONFIG_HEAD_NS = "iot.device.config"
var CONFIG_HEAD_NAME = "config"

type PackConfig struct {
	Header  Header       `json:"header"`
	Payload ConfigParams `json:"payload"`
}

func (o *PackConfig) Encode(gid string, data DetailParam) ([]byte, error) {
	obj := PackConfig{
		Header:  EncodeHeader(CONFIG_HEAD_NS, CONFIG_HEAD_NAME, gid, ""),
		Payload: ConfigParams{Param: data},
	}
	*o = obj
	return json.Marshal(obj)
}

type PackConfigAck struct {
	Header  Header    `json:"header"`
	Payload NormalAck `json:"payload"`
}
func (o *PackConfigAck) Encode(gid, mid string, err error) ([]byte, error) {
	obj := PackConfigAck{
		Header:  EncodeHeader(CONFIG_HEAD_NS, CONFIG_HEAD_NAME, gid, mid),
		Payload: NormalAck{Code: 0, Msg: "ok"},
	}
	if err != nil {
		obj.Payload.Code = 1
		obj.Payload.Msg = err.Error()
	}
	*o = obj
	return json.Marshal(obj)
}

type ConfigParams struct {
	Param DetailParam `json:"param"`
}

type DetailParam struct {
	ReportTimer *int     `json:"reportTimer,omitempty"`
	Keepalive   *int     `json:"keepalive,omitempty"`
	Timeout     *int     `json:"timeout,omitempty"`
	Events      []string `json:"events,omitempty"`
}
