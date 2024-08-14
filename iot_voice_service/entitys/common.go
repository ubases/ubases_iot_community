package entitys

import "cloud_platform/iot_proto/protos/protosService"

type VoiceProductCached struct {
	VoiceProductInfo *protosService.OpmVoiceProduct      //产品语控信息
	FunctionMap      []*protosService.OpmVoiceProductMap //产品语控制功能配置
	VoiceBrand       string                              `json:"voiceBrand"` //品牌名称
	VoiceModel       string                              `json:"voiceModel"` //产品型号
	VoiceSkill       string                              `json:"voiceSkill"` //技能ID
}

type ControlData struct {
	ProductKey     string
	DeviceId       string
	Data           map[string]interface{} //Key为FunName
	DataDpid       map[int32]interface{}  //Key为Dpid
	DataVoiceCode  map[string]interface{} //Key为VoiceCode
	DataVoice      map[string]interface{}
	SubscriptionId string //小米专用
}
