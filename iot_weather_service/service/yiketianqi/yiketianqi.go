package yiketianqi

import (
	"cloud_platform/iot_weather_service/config"
	"strconv"

	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
)

const yiketianqiURL = "https://v0.yiketianqi.com/api"

type TianqiWeather struct {
	City      string `json:"city"`
	CityEn    string `json:"cityEn"`
	Country   string `json:"country"`
	CountryEn string `json:"countryEn"`
	Air       string `json:"air"`
	AirPm25   string `json:"air_pm25"`
	AirLevel  string `json:"air_level"`
}

// 根据城市名获取天气
func GetTianqiWeatherByCity(city string) (*TianqiWeather, error) {
	paras := make(map[string]string)
	paras["city"] = city
	return getTianqiWeather(paras)
}

// 根据IP地址获取天气
func GetTianqiWeatherByIP(ip string) (*TianqiWeather, error) {
	paras := make(map[string]string)
	paras["ip"] = ip
	return getTianqiWeather(paras)
}

func getTianqiWeather(paras map[string]string) (*TianqiWeather, error) {
	paras["unescape"] = "1"
	paras["version"] = "v61"
	paras["appid"] = strconv.Itoa(config.Global.Yiketianqi.Appid)
	paras["appsecret"] = config.Global.Yiketianqi.Appsecret
	resp, err := resty.New().R().SetQueryParams(paras).Get(yiketianqiURL)
	if err != nil {
		return nil, err
	}
	var ret TianqiWeather
	err = json.Unmarshal(resp.Body(), &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
