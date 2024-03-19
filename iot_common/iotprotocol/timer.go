package iotprotocol

var TIMER_HEAD_NS = "iot.device.timer"
var TIMER_HEAD_NAME = "timer"
var TIMER_HEAD_NAME_COUNTDOWN = "countdown"

//定时器
type PackTimer struct {
	Header  Header    `json:"header"`
	Payload TimerData `json:"payload"`
}

type PackTimerAck struct {
	Header  Header    `json:"header"`
	Payload NormalAck `json:"payload"`
}

type TimerData struct {
	Timer     []DetailTimer `json:"timer,omitempty"`     //定时器
	Countdown *DetailTimer  `json:"countdown,omitempty"` //倒计时
}

type DetailTimer struct {
	Enabled int                    `json:"enabled"`
	Type    string                 `json:"type"`
	At      string                 `json:"at"` //格式2016-11-29T02:49:00Z
	Do      map[string]interface{} `json:"do"`
}
