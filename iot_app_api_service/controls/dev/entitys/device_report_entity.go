package entitys

type DeviceLogDateRequest struct {
	ProductKey string `json: "productKey"` //产品model
	DevId      string `json: "devid"`      //设备编号
	Date       string `json:"date"`        //月时间
}
