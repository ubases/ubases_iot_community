package iotprotocol

import json "github.com/json-iterator/go"

var INFO_HEAD_NS = "iot.device.info"
var INFO_HEAD_NAME = "info"

type PackInfo struct {
	Header Header `json:"header"`
	//无payload部分
}

func (o *PackInfo) Encode(gid string) ([]byte, error) {
	obj := PackInfo{
		Header: EncodeHeader(INFO_HEAD_NS, INFO_HEAD_NAME, gid)}
	*o = obj
	return json.Marshal(obj)
}

type PackInfoReport struct {
	Header  Header     `json:"header"`
	Payload InfoReport `json:"payload"`
}

type InfoReport struct {
	Error      int    `json:"error"`
	UID        int64  `json:"uid"`
	DeviceId   string `json:"deviceId"`
	ProductKey string `json:"productKey"`
	Token      string `json:"token"`
	SecrtKey   string `json:"secrtKey"`
	FwVer      string `json:"fwVer"`
	McuVer     string `json:"mcuVer"`
	HwVer      string `json:"hwVer"`
	MemFree    int    `json:"memFree"`
	Mac        string `json:"mac"`
	Ap         Ap     `json:"ap"`
	Netif      Netif  `json:"netif"`
}

type Ap struct {
	Ssid  string `json:"ssid"`
	Bssid string `json:"bssid"`
	Rssi  int    `json:"rssi"`
}

type Netif struct {
	LocalIP string `json:"localIp"`
	Mask    string `json:"mask"`
	Gw      string `json:"gw"`
}

//fixme 这个用于测试，请勿在正式产品中使用
func (o *PackInfoReport) Encode(gid, DeviceId, ProductKey string) ([]byte, error) {
	obj := PackInfoReport{
		Header:  EncodeHeader(INFO_HEAD_NS, INFO_HEAD_NAME, gid),
		Payload: InfoReport{DeviceId: DeviceId, ProductKey: ProductKey},
	}
	*o = obj
	return json.Marshal(obj)
}
