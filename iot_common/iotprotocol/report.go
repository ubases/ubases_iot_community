package iotprotocol

import json "github.com/json-iterator/go"

var REPORT_HEAD_NS = "iot.device.report"
var REPORT_HEAD_NAME = "report"

type PackReport struct {
	Header  Header     `json:"header"`
	Payload DeviceData `json:"payload"`
}

type DeviceData struct {
	DeviceData map[string]interface{} `json:"device,omitempty"`
}

func (o *PackReport) Encode(gid string, data map[string]interface{}) ([]byte, error) {
	obj := PackReport{
		Header:  EncodeHeader(REPORT_HEAD_NS, REPORT_HEAD_NAME, gid),
		Payload: DeviceData{data},
	}
	*o = obj
	return json.Marshal(obj)
}
