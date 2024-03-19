package entitys

type FlashScreen struct {
	Id           string   `json:"id"`
	StartTime    string   `json:"startTime"`
	EndTime      string   `json:"endTime"`
	PutinUser    int32    `json:"putinUser"`
	Accounts     []string `json:"accounts"`
	OpenPageType int      `json:"openPageType"`
	AppPageType  int      `json:"appPageType"`
	OpenPageUrl  string   `json:"openPageUrl"`
	ShowImageUrl string   `json:"showImageUrl"`
	ShowImageMd5 string   `json:"showImageMd5"`
	ShowTime     int      `json:"showTime"`
}

type ImageInfo struct {
	ImageUrl string `json:"imageUrl"`
	ImageMd5 string `json:"imageMd5"`
}
