package iotstruct

type OtaPublishLog struct {
	PubId        int64  `json:"pubId"`
	SuccessCount int64  `json:"successCount"`
	TotalCount   int64  `json:"totalCount"`
	Message      string `json:"message"`
	Status       int32  `json:"status"`
}
