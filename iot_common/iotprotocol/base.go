package iotprotocol

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const FROM_APP = "app"
const FROM_DEVICE = "device"
const FROM_CLOUD = "cloud"
type Header struct {
	Ns   string `json:"ns"`
	Name string `json:"name"`
	Mid  string `json:"mid"`
	Ts   int64  `json:"ts"`
	Ver  string `json:"ver"`
	Gid  string `json:"gid"`
	From string `json:"from"`
}

type NormalAck struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}

func EncodeHeader(ns, name, gid, mid string) Header {
	if mid == "" {
		mid = uuid.NewV4().String()
	}
	return Header{Ns: ns, Name: name, Mid: mid, Ts: time.Now().UTC().Unix(), Ver: "1.0.0", Gid: gid, From: FROM_DEVICE}
}

// 响应时拷贝上行的head，目前只需要修改时间戳
// todo 需要和嵌入式、app协商，ts改为整数时间戳
func (h *Header) CopyHeader() Header {
	hh := Header{
		Ns:   h.Ns,
		Name: h.Name,
		Mid:  h.Mid,
		Ts:   time.Now().UTC().Unix(),
		Ver:  h.Ver,
		Gid:  h.Gid,
		From: h.From,
	}
	return hh
}
type PackHeader struct {
	Header Header `json:"header"`
}
