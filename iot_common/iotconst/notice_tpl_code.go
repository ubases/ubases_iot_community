package iotconst

type NoticeType string
type NoticeMethod string

var (
	NOTICE_VERIFICATION_CODE   NoticeType   // 短信验证码
	NOTICE_VERIFICATION_METHOD NoticeMethod //通知方式
)

type NoticeRequest struct {
	Target string                 `json:"target"` //推送目标
	Type   NoticeType             `json:"type"`   //推送类型
	Method NoticeMethod           `json:"method"` //推送类型
	Data   map[string]interface{} `json:"data"`   //推送参数
}

var (
	SceneIntelligenceNotice = "SceneIntelligenceNotice"
)
