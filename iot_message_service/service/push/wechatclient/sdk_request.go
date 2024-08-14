package wechatclient

type SdkRequestMap map[string]string

type SdkRequest struct {
	title string
	data  map[string]SdkRequestMap
	gMap  map[string]interface{}
}

func NewSdkRequest() *SdkRequest {
	return &SdkRequest{
		data: make(map[string]SdkRequestMap),
	}
}

func (m *SdkRequest) GetGMap() map[string]interface{} {
	return m.gMap
}

func (m *SdkRequest) SetGMap(gMap map[string]interface{}) {
	m.gMap = gMap
}

func (m *SdkRequest) Set(outerKey, innerKey, value string) {
	innerMap, ok := m.data[outerKey]
	if !ok {
		innerMap = make(map[string]string)
		m.data[outerKey] = innerMap
	}
	innerMap[innerKey] = value
}

func (m *SdkRequest) SetMap(value map[string]string) {
	for outerKey, val := range value {
		innerMap, ok := m.data[outerKey]
		if !ok {
			innerMap = make(map[string]string)
			m.data[outerKey] = innerMap
		}
		innerMap["value"] = val
	}
}

func (m *SdkRequest) Get(outerKey, innerKey string) (string, bool) {
	innerMap, ok := m.data[outerKey]
	if !ok {
		return "", false
	}
	value, ok := innerMap[innerKey]
	return value, ok
}

func (m *SdkRequest) GetVal(outerKey, innerKey string) string {
	innerMap, ok := m.data[outerKey]
	if !ok {
		return ""
	}
	value, _ := innerMap[innerKey]
	return value
}

func (m *SdkRequest) SetTitle(title string) {
	m.title = title
}
func (m *SdkRequest) GetTitle() string {
	return m.title
}
