package entitys

import (
	"cloud_platform/iot_proto/protos/protosService"
	"strings"
)

type CheckOtaVersionResponse struct {
	Code            int32  `json:"code,omitempty"`
	Message         string `json:"message,omitempty"`
	Custom          bool   `json:"custom"`
	ProductKey      string `json:"productKey"`
	FirmwareName    string `json:"firmwareName"`
	Remark          string `json:"remark"`
	Version         string `json:"version"`
	McuVersion      string `json:"mcuVersion"`
	UpgradeMode     int32  `json:"upgradeMode"`
	IsAuto          int32  `json:"isAuto"`          // 是否自动升级
	IsAutoUpgrade   bool   `json:"isAutoUpgrade"`   // 是否自动升级
	UpgradeTimeMode int32  `json:"upgradeTimeMode"` // 升级时间模式 =1 全天 =2 指定时间
	UpgradeOvertime int32  `json:"upgradeOvertime"` // 超时时间
	OtaState        string `json:"otaState"`        //ota升级进度状态
	Progress        int32  `json:"progress"`        //ota升级进度

	Payload PayloadParams `json:"payload"`
}

type CheckOtaVersionSimpleResponse struct {
	IsAuto        int32 `json:"isAuto"`        // 是否自动升级
	IsAutoUpgrade bool  `json:"isAutoUpgrade"` // 是否自动升级
}

type PayloadParams struct {
	Chanel     int    `json:"chanel"`
	Version    string `json:"otaVer"`
	PubId      int64  `json:"pubId,string"`
	PointVer   string `json:"pointVer"`
	BaseVer    string `json:"baseVer"`
	McuBaseVer string `json:"mcuBaseVer"`
	OtaType    string `json:"otaType"`
	AppURL     string `json:"appUrl"`
	McuURL     string `json:"mcuUrl"`
	Md5        string `json:"md5"`
	Timeout    int32  `json:"timeout"` //超时时间
}

type SetAutoUpgradeRequest struct {
	DeviceId      string `json:"deviceId"`      //设备编号
	IsAutoUpgrade bool   `json:"isAutoUpgrade"` // 是否自动升级
}

func CheckOtaVersion_Pd2E(src *protosService.CheckOtaVersionResponse, lang string) *CheckOtaVersionResponse {
	otaParams := &CheckOtaVersionResponse{
		Code:            src.Code,
		Message:         src.Message,
		Custom:          src.OtaPkg.Custom,
		ProductKey:      src.OtaPkg.ProductKey,
		FirmwareName:    src.OtaPkg.FirmwareName,
		Remark:          src.UpgradePublish.UpgradeDesc,
		Version:         src.OtaPkg.Version,
		UpgradeMode:     src.UpgradePublish.UpgradeMode,
		IsAuto:          src.UpgradePublish.IsAuto,
		UpgradeTimeMode: src.UpgradePublish.UpgradeTimeMode,
		UpgradeOvertime: src.OtaPkg.UpgradeOvertime,
		OtaState:        src.OtaPkg.OtaState,
		Progress:        src.OtaPkg.Progress,
	}
	if lang == "en" {
		otaParams.Remark = src.UpgradePublish.UpgradeDescEn
	}
	//todo 转换http
	appUrl := strings.Replace(src.OtaPkg.AppUrl, "https://", "http://", 1)
	mcuUrl := strings.Replace(src.OtaPkg.McuUrl, "https://", "http://", 1)
	otaParams.Payload = PayloadParams{
		Chanel:   2,
		PubId:    src.UpgradePublish.PubId,
		PointVer: src.UpgradePublish.PointVer,
		BaseVer:  src.UpgradePublish.BaseVer,
		OtaType:  src.OtaPkg.OtaType,
		AppURL:   appUrl,
		McuURL:   mcuUrl,
		Md5:      src.OtaPkg.Md5,
		Timeout:  src.OtaPkg.UpgradeOvertime,
	}
	otaParams.Payload.Version = src.OtaPkg.Version
	if src.OtaPkg.OtaType == "module_mcu_all" {
		otaParams.McuVersion = src.OtaPkg.Version
		//"module_mcu_all"
		//otaParams.Payload.McuBaseVer = src.OtaPkg.Version
	}
	//else {
	//	//"module_ota_all"
	//	otaParams.Payload.Version = src.OtaPkg.Version
	//}
	return otaParams
}
