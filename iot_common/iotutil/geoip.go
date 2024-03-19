package iotutil

import (
	"cloud_platform/iot_common/iotlogger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

// 位置信息接口数据
type GeoipInfo struct {
	EnCode   string  `json:"en_code"`
	EnName   string  `json:"en_name"`
	Country  string  `json:"country"`
	Province string  `json:"province"`
	City     string  `json:"city"`
	District string  `json:"district"`
	Adcode   int     `json:"adcode"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

// 通过IP获取地址位置信息
func Geoip(ip, queryUrl, appCodeE string) (geo GeoipInfo, err error) {
	//初始化
	geo = GeoipInfo{}
	//测试代码，如果是局域网位置，默认绑定配网文件ip
	var ipaddress string = ip
	if strings.Contains(ip, "127.0.0") || strings.Contains(ip, "192.168") {
		return
	}
	queryUrl = queryUrl + ipaddress
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", queryUrl, nil)

	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("通过IP获取地理位置失败[connect], ip:%s, err: %s", ip, err.Error()))
		return
	}
	reqest.Header.Add("Authorization", appCodeE)
	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		iotlogger.LogHelper.Error(fmt.Sprintf("通过IP获取地理位置失败[body], ip:%s, err: %s", ip, err.Error()))
		return
	}
	info := make(map[string]interface{})

	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(bodyByte, &info)
	}
	//if err == nil {
	//	global.GVA_LOG.Error(fmt.Sprintf("get address by ip[%s], error:%s", ip, err.Error()))
	//}
	if ToString(info["code"]) != "100" {
		return
	}

	result := info["result"].(map[string]interface{})
	//TODO 实体
	geo.EnCode = ToString(result["en_short"])
	geo.EnName = ToString(result["en_name"])
	geo.Country = ToString(result["nation"])
	geo.Province = ToString(result["province"])
	geo.City = ToString(result["city"])
	geo.District = ToString(result["district"])
	geo.Lat = ToFloat64(result["lat"])
	geo.Lng = ToFloat64(result["lng"])
	return
}

func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}
