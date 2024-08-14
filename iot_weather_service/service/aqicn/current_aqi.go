package aqicn

import (
	"cloud_platform/iot_weather_service/config"
	"cloud_platform/iot_common/iotutil"
	"encoding/json"
	"errors"
	"fmt"
)

var AqiError string = "error"

// CurrentWeatherData struct contains an aggregate view of the structs
// defined above for JSON to be unmarshaled into.
type CurrentAriQualityData struct {
	Aqi  interface{} `json:"aqi"`
	IAqi *IAqi       `json:"iaqi"`
	*Settings
}

type CurrentAriQualityDataResponse struct {
	//"status":"error","data":"Unknown station"
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type IAqi struct {
	Pm10 IAQIValue `json:"pm10"`
	Pm25 IAQIValue `json:"pm25"`
}

type IAQIValue struct {
	V float64 `json:"v"`
}

// CurrentByName will provide the current weather with the provided
// location name.

func (w *CurrentAriQualityData) CurrentByName(location string) error {
	response, err := w.client.Get(fmt.Sprintf(baseURL, iotutil.UrlEncode(location), config.Global.Weather.AqicnToken))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var res CurrentAriQualityDataResponse
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return err
	}
	if res.Status == AqiError {
		return errors.New(iotutil.ToString(res.Data))
	}
	var aqiData CurrentAriQualityData
	err = iotutil.StructToStructErr(res.Data, &aqiData)
	if err != nil {
		return err
	}
	w.IAqi = aqiData.IAqi
	w.Aqi = aqiData.Aqi
	return nil
}
