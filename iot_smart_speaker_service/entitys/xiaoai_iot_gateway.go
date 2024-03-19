package entitys

//=============================================
//小爱请求参数
type XiaoaiRequest struct {
	Version string        `json:"version"`
	Session SessionReqObj `json:"session"`
	Request RequestObj    `json:"request"`
	Query   string        `json:"query"`
	Context Context       `json:"context"`
}
type SessionReqObj struct {
	IsNew       bool           `json:"is_new"`
	SessionID   string         `json:"session_id"`
	Application ApplicationObj `json:"application"`
	User        UserObj        `json:"user"`
}
type ApplicationObj struct {
	AppID string `json:"app_id"`
}
type UserObj struct {
	UserID      string `json:"user_id"`
	IsUserLogin bool   `json:"is_user_login"`
	Gender      string `json:"gender"`
	AccessToken string `json:"access_token"`
}
type IntentObj struct {
	Query          string  `json:"query"`
	Score          float64 `json:"score"`
	Complete       bool    `json:"complete"`
	Domain         string  `json:"domain"`
	Confidence     int     `json:"confidence"`
	SkillType      string  `json:"skillType"`
	SubDomain      string  `json:"sub_domain"`
	AppID          string  `json:"app_id"`
	RequestType    string  `json:"request_type"`
	NeedFetchToken bool    `json:"need_fetch_token"`
	IsDirectWakeup bool    `json:"is_direct_wakeup"`
	Slots          string  `json:"slots"`
}
type Slots struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	RawValue string `json:"raw_value"`
}
type SlotInfoObj struct {
	IntentName string  `json:"intent_name"`
	Slots      []Slots `json:"slots"`
}
type RequestObj struct {
	Type      int         `json:"type"`
	RequestID string      `json:"request_id"`
	Timestamp int64       `json:"timestamp"`
	Intent    IntentObj   `json:"intent"`
	Locale    string      `json:"locale"`
	SlotInfo  SlotInfoObj `json:"slot_info"`
}
type Context struct {
	DeviceID string `json:"device_id"`
	InExp    bool   `json:"in_exp"`
}

//=============================================
//开发者响应返回数据结构
type XiaoaiResponse struct {
	IsSessionEnd      bool              `json:"is_session_end"`
	Version           string            `json:"version"`
	Response          ResponseObj       `json:"response"`
	SessionAttributes SessionAttributes `json:"session_attributes"`
}
type ToSpeak struct {
	Type int    `json:"type"`
	Text string `json:"text"`
}
type ToDisplay struct {
	Type int    `json:"type"`
	Text string `json:"text"`
}
type LogInfo struct {
}
type ResponseObj struct {
	Confidence float64   `json:"confidence"`
	OpenMic    bool      `json:"open_mic"`
	ToSpeak    ToSpeak   `json:"to_speak"`
	ToDisplay  ToDisplay `json:"to_display"`
	LogInfo    LogInfo   `json:"log_info"`
}
type MiniSkillInfoObj struct {
	Name string `json:"name"`
}
type SessionData struct {
	Query  string `json:"query"`
	Reply  string `json:"reply"`
	Engine string `json:"engine"`
}
type SessionObj struct {
	SessionID    string        `json:"sessionID"`
	SkillName    string        `json:"skillName"`
	SkillSubName string        `json:"skillSubName"`
	Turn         int           `json:"turn"`
	Data         []SessionData `json:"data"`
}
type SessionAttributes struct {
	NoticeFlag     int              `json:"noticeFlag"`
	ReplyKeyWord   string           `json:"replyKeyWord"`
	Turn           int              `json:"turn"`
	MiniSkillInfo  MiniSkillInfoObj `json:"miniSkillInfo"`
	Session        SessionObj       `json:"session"`
	LongtailEngine string           `json:"longtailEngine"`
	Replace        bool             `json:"replace"`
	Latitude       float64          `json:"latitude"`
	Longtitude     float64          `json:"longtitude"`
}
