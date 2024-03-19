package iotprotocol

import json "github.com/json-iterator/go"

//APP通知
var NOTICE_HEAD_NS = "iot.app.notice"

var NOTICE_HEAD_UPGRADE_NOTICE_NAME = "otaUpgradeNotice"

type PackNotice struct {
	Header  Header                 `json:"header"`
	Payload map[string]interface{} `json:"payload"`
}

func (o *PackNotice) Encode(name string, data map[string]interface{}) ([]byte, error) {
	obj := PackNotice{
		Header:  EncodeHeader(NOTICE_HEAD_NS, name, ""),
		Payload: data,
	}
	*o = obj
	return json.Marshal(obj)
}
