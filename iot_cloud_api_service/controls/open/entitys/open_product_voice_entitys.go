package entitys

type VoicePublishRecordRes struct {
	ProductName string   `json:"productName"`
	VoiceName   string   `json:"voiceName"`
	AttrText    []string `json:"attrText"`
	CreatedAt   int64    `json:"createdAt"`
	Id          string   `json:"id"`
}
