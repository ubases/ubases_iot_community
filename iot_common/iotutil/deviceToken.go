package iotutil

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

var deviceTokenKey string = "12345678"

type DeviceTokenInfo struct {
	UID string
	ET  int64
}

func EncryptDeviceToken(userId string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * 24 * time.Hour)
	data := DeviceTokenInfo{
		UID: userId,
		ET:  expireTime.Unix(),
	}
	str, _ := json.Marshal(data)
	enStr, _ := AES_CBC_EncryptBase64([]byte(str), []byte(deviceTokenKey))
	return enStr
}

func DecryptDeviceToken(token string) *DeviceTokenInfo {
	str, _ := base64.StdEncoding.DecodeString(token)
	data, _ := AES_CBC_EncryptBase64([]byte(str), []byte(deviceTokenKey))
	dataStr := string(data)
	if dataStr == "" {
		return nil
	}
	devInfo := &DeviceTokenInfo{}
	err := json.Unmarshal([]byte(dataStr), &devInfo)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return devInfo
}
