package iotprotocol

import json "github.com/json-iterator/go"

var QUERY_HEAD_NS = "iot.device.query"
var QUERY_HEAD_NAME = "query"        //获取设备状态信息
var QUERY_HEAD_NAME_ALL = "queryAll" //获取所有属性

type PackQuery struct {
	Header  Header `json:"header"`
	Payload Query  `json:"payload"`
}

func (o *PackQuery) Encode(gid, name string, paras []string) ([]byte, error) {
	obj := PackQuery{
		Header:  EncodeHeader(QUERY_HEAD_NS, name, gid, ""),
		Payload: Query{Param: QueryParams{Props: paras}},
	}
	*o = obj
	return json.Marshal(obj)
}

type PackQueryAck struct {
	Header  Header   `json:"header"`
	Payload QueryAck `json:"payload"`
}

type Query struct {
	Param QueryParams `json:"param"`
}
type QueryParams struct {
	Props []string `json:"props,omitempty"`
}

type QueryAck struct {
	Code       int                    `json:"code"`
	DeviceData map[string]interface{} `json:"device,omitempty"`
}
