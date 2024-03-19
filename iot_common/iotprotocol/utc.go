package iotprotocol

var UTC_HEAD_NS = "iot.device.utc"
var UTC_HEAD_NAME = "utc" //UTC时间请求

type PackUTC struct {
	Header Header `json:"header"`
	//无payload部分
}

type PackUTCAck struct {
	Header  Header `json:"header"`
	Payload UtcAck `json:"payload"`
}

type UtcAck struct {
	Code int    `json:"code"`
	Data string `json:"data"` //RFC3339格式
}
