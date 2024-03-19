package iotprotocol

import (
	json "github.com/json-iterator/go"
)

var OTA_HEAD_NS = "iot.device.upgrade"
var OTA_HEAD_NAME_OTAINFO = "otaInfo"         //升级信息
var OTA_HEAD_NAME_OTAPROGRESS = "otaProgress" //升级进度

const UPSTATE_DOWNLOADING = "Downloading"
const UPSTATE_INSTALLING = "Installing"

type PackUpgrade struct {
	Header  Header       `json:"header"`
	Payload UpgradeParam `json:"payload"`
}

func (o *PackUpgrade) Encode(gid string, data UpgradeDetailParam) ([]byte, error) {
	obj := PackUpgrade{
		Header:  EncodeHeader(OTA_HEAD_NS, OTA_HEAD_NAME_OTAINFO, gid),
		Payload: UpgradeParam{Param: data},
	}
	*o = obj
	return json.Marshal(obj)
}

type PackUpgradeAck struct {
	Header  Header    `json:"header"`
	Payload NormalAck `json:"payload"`
}

type PackUpgradeReport struct {
	Header  Header        `json:"header"`
	Payload UpgradeReport `json:"payload"`
}

type UpgradeParam struct {
	Param UpgradeDetailParam `json:"param"`
}

type UpgradeDetailParam struct {
	Chanel     int    `json:"chanel"`
	PointVer   string `json:"pointVer"`
	BaseVer    string `json:"baseVer"`
	McuBaseVer string `json:"mcuBaseVer"`
	OtaType    string `json:"otaType"`
	AppURL     string `json:"appUrl"`
	McuURL     string `json:"mcuUrl"`
	Md5        string `json:"md5"`
	PubId      string `json:"pubId"`  //ota发布Id
	OtaVer     string `json:"otaVer"` //ota版本号
	Timeout    int32  `json:"timeout"`
}

type UpgradeReport struct {
	OtaState string `json:"otaState"` //OTA升级状态
	Code     int    `json:"code"`     //升级错误码
	Progress int    `json:"progress"` //进度
	Version  string `json:"otaVer"`   //当前升级版本号
	PubId    string `json:"pubId"`    //升级发送Id
}
