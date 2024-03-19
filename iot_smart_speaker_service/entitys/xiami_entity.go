package entitys

//"{\"requestId\":\"0265f4b6-7dd0-4221-a15c-e78f851b4da0\",\"intent\":\"get-devices\"}
type XiaomiRequest struct {
	RequestId  string        `json:"requestId"`
	Intent     string        `json:"intent"`
	Devices    []interface{} `json:"devices,omitempty"`
	Properties []interface{} `json:"properties,omitempty"`
}

type XiaomiDevices struct {
	Did  string `json:"did"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type XiaomiSubscribe struct {
	Did            string `json:"did"`
	SubscriptionId string `json:"subscriptionId,omitempty"`
	Status         int    `json:"status"`
	Description    string `json:"description,omitempty"` //如果status<0，此字段必须存在，用于描述失败的原因
}

type XiaomiSubscribeNotify struct {
	Did            string `json:"did"`
	SubscriptionId string `json:"subscriptionId"`
	Status         int    `json:"status"`
}

type XiaomiDeviceStatus struct {
	Did         string `json:"did"`
	Status      int    `json:"status,omitempty"`
	Online      bool   `json:"online,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"` //如果status<0，此字段必须存在，用于描述失败的原因
}

type XiaomiProperties struct {
	Did         string      `json:"did"`
	Siid        interface{} `json:"siid"`
	Piid        interface{} `json:"piid"`
	PidName     string      `json:"pidName,omitempty"`
	Status      interface{} `json:"status,omitempty"`      // 0态码，0代表成功，负值代表失败
	Value       interface{} `json:"value,omitempty"`       // 属性值，格式必须是产品功能定义中定义数据格式
	Description string      `json:"description,omitempty"` //如果status<0，此字段必须存在，用于描述失败的原因
}
