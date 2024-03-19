package entitys

//配网信息数据
type ServerEntity struct {
	MqttIp     string `json:"mqttIp"` //MQTT Serve让服务器地址 mqtts://axydev.aithinker.com:1883
	MqttPort   int32  `json:"mqttPort"`
	Nationcode string `json:"nationcode"`
}
