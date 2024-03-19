package iotstruct

type TimerFunctions struct {
	FuncName       string      `json:"funcName"`
	FuncKey        interface{} `json:"funcKey"`
	FuncValue      interface{} `json:"funcValue"`
	FuncDesc       string      `json:"funcDesc"`
	FuncValueDesc  string      `json:"funcValueDesc"`
	FuncIdentifier string      `json:"funcIdentifier"`
}
