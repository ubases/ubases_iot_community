package iotprotocol

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

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
}

func EncodeHeader(ns, name, gid string) Header {
	return Header{Ns: ns, Name: name, Mid: uuid.NewV4().String(), Ts: time.Now().UTC().Unix(), Ver: "1.0.0", Gid: gid, From: "cloud"}
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
