package convert

type Header struct {
	Ns   string      `json:"ns"`
	Name string      `json:"name"`
	Mid  string      `json:"mid"`
	Ts   interface{} `json:"ts"`
	Ver  interface{} `json:"ver"`
	Gid  string      `json:"gid"`
}

type DeviceInfoReport struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Ap struct {
	Ssid  string `json:"ssid"`
	Bssid string `json:"bssid"`
	Rssi  int    `json:"rssi"`
	//Primary int    `json:"primary"`
}

type Netif struct {
	LocalIP string `json:"localIp"`
	Mask    string `json:"mask"`
	Gw      string `json:"gw"`
}
type Payload struct {
	Error      int    `json:"error"`
	UID        int64  `json:"uid"`
	DeviceId   string `json:"deviceId"`
	ProductKey string `json:"productKey"`
	Token      string `json:"token"`
	SecrtKey   string `json:"secrtKey"`
	FwVer      string `json:"fwVer"`
	McuVer     string `json:"mcuVer"`
	HwVer      string `json:"hwVer"`
	MemFree    int    `json:"memFree"`
	Mac        string `json:"mac"`
	Ap         Ap     `json:"ap"`
	Netif      Netif  `json:"netif"`
}
