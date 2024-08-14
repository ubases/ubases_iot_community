package iotprotocol

import json "github.com/json-iterator/go"

var REPORT_HEAD_NS = "iot.device.report"
var REPORT_HEAD_NAME = "report"          //�1�7�1�7�0�2�1�7�1�7�1�7�1�7�1�7�0�9�1�7
var REPORT_HEAD_NAME_CONTROL = "control" //�1�7�1�7�1�7�0�0�1�7�1�7�1�7�1�7�0�9�1�7
var REPORT_HEAD_NAME_QUERY = "query"     //�1�7�1�7�0�9�1�7�1�7�1�7�1�7�1�7�0�9�1�7

type PackReport struct {
	Header  Header     `json:"header"`
	Payload DeviceData `json:"payload"`
}

type DeviceData struct {
	DeviceData map[string]interface{} `json:"device,omitempty"`
}

func (o *PackReport) Encode(gid, name, mid string, data map[string]interface{}) ([]byte, error) {
	obj := PackReport{
		Header:  EncodeHeader(REPORT_HEAD_NS, name, gid, mid),
		Payload: DeviceData{data},
	}
	*o = obj
	return json.Marshal(obj)
}
