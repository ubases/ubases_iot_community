package scene_executor

import (
	"cloud_platform/iot_intelligence_service/rpc/rpcclient"
	cron2 "cloud_platform/iot_intelligence_service/service/scene_executor/cron"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"time"
	"unicode"

	"github.com/robfig/cron/v3"
)

//var expireTIme = time.Duration(40 * time.Minute)

func ClearWeatherRedisKey() {
	weatherCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_WEATHER_LIST)
	if weatherCmd.Err() != nil {
		iotlogger.LogHelper.Error("缓存天气获取异常", weatherCmd.Err())
		return
	}
	//移除所有定时任务
	for _, val := range weatherCmd.Val() {
		if val != "" && val != "0" {
			cron2.CronCtx.Remove(cron.EntryID(iotutil.ToInt32(val)))
		}
	}
}

var cityKeys = []string{
	"长沙市",
	"深圳市",
	"广州市",
}

func ConvertCityName(cityName string) string {
	pinyinCityName := cityName
	if hasHans(cityName) { //去掉市县区
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "市")
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "县")
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "区")
		pinyinCityName = strings.Join(pinyin.LazyConvert(pinyinCityName, nil), "")
	}
	return cityName
}

func InitMonitorWeatherData() {
	//添加默认
	iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_WEATHER_LIST, map[string]interface{}{
		"长沙": "",
		"深圳": "",
		"佛山": "",
		"广州": "",
	})

	weatherCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_WEATHER_LIST)
	if weatherCmd.Err() != nil {
		iotlogger.LogHelper.Error("缓存天气获取异常", weatherCmd.Err())
		return
	}
	//重启所有
	for key, val := range weatherCmd.Val() {
		if val != "" && val != "0" {
			cron2.CronCtx.Remove(cron.EntryID(iotutil.ToInt32(val)))
		}
		monitorWeatherData(ConvertCityName(key), true)
	}
}

func MonitorWeatherChange(city string) {
	cityName := ConvertCityName(city)
	monitorWeatherData(cityName, false)
}

func hasHans(str string) bool {
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			return true
		}
	}
	return false
}

// MonitorWeatherChange 监控规则
func monitorWeatherData(cityName string, isReset bool) {
	defer iotutil.PanicHandler(cityName)

	if !isReset {
		//是否重新开始任务
		weatherCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_WEATHER_LIST, cityName)
		if weatherCmd.Val() != "" {
			return
		}
	}

	//默认读取一次
	reloadWeather(cityName)

	id, _ := cron2.CronCtx.AddFunc("0 */10 * * * *", func() {
		reloadWeather(cityName)
	})

	weatherSaveCmd := iotredis.GetClient().HSet(context.Background(), iotconst.HKEY_WEATHER_LIST, cityName, iotutil.ToString(id))
	if weatherSaveCmd.Err() != nil {
		iotlogger.LogHelper.Error("天气缓存异常", weatherSaveCmd.Err())
		return
	}
	iotlogger.LogHelper.Infof("创建了天气缓存定时任务：id:%d", id)
}

func reloadWeather(cityName string) {
	defer iotutil.PanicHandler(cityName)
	ret, err := rpcclient.ClientWeatherService.CurrentByCity(context.Background(), &protosService.CityRequest{
		CityName: cityName,
	})
	if err != nil {
		iotlogger.LogHelper.Error("天气获取异常", err)
		return
	}
	//var weatherMap map[string]interface{}
	//iotutil.StructToStruct(ret.Data, &weatherMap)
	//if weatherMap == nil {
	//	iotlogger.LogHelper.Error("天气存储异常", err)
	//	return
	//}
	if ret.Data == nil {
		iotlogger.LogHelper.Error("天气获取异常")
		return
	}
	weatherKey := iotconst.HKEY_WEATHER_DATA + cityName

	//转换天气值 (日出1， 日落2）
	isSunrise := 0
	nowTime := time.Now()
	sunriseTime := ret.Data.Sunrise.AsTime()
	if sunriseTime.Add(60*time.Minute).After(nowTime) && sunriseTime.Before(nowTime) {
		isSunrise = 1
	}
	sunsetTime := ret.Data.Sunset.AsTime()
	if sunsetTime.Add(60*time.Minute).After(nowTime) && sunsetTime.Before(nowTime) {
		isSunrise = 2
	}
	weather := iotconst.WEATHER_VALUE_0
	weatherStr := strings.ToLower(ret.Data.Weather)
	//https://openweathermap.org/weather-conditions
	switch weatherStr {
	case "clear":
		weather = iotconst.WEATHER_VALUE_1
	case "clouds":
		weather = iotconst.WEATHER_VALUE_2
	case "thunderstorm", "rain", "drizzle": //雷雨 雨 下蒙蒙细雨
		weather = iotconst.WEATHER_VALUE_3
	case "snow":
		weather = iotconst.WEATHER_VALUE_4
	case "haze", "mist", "smoke", "dust", "fo", "sand", "ash", "squall":
		weather = iotconst.WEATHER_VALUE_5
		//701	Mist 薄雾	薄雾	50天
		//711	Smoke 抽烟	抽烟	50天
		//721	Haze 阴霾	阴霾	50天
		//731	Dust  灰尘	沙尘漩涡	50天
		//741	Fo 多雾路段	多雾路段	50天
		//751	Sand 沙	沙	50天
		//761	Dust 灰尘	灰尘	50天
		//762	Ash  灰	火山灰	50天
		//771	Squall 飑	暴风雨	50天
		//781	Tornado 龙卷风	龙卷风	50天
	}
	//开氏度 转 摄氏度
	temp := iotutil.HToSTemperature(ret.Data.Temperature) // 5 * (iotutil.ToInt(ret.Data.Temperature) - 32) / 9
	cachedMaps := map[string]interface{}{
		"city":        cityName,
		"time":        ret.Data.UpdatedAt,
		"weather":     weather,
		"sun":         isSunrise,
		"temperature": temp,
		"humidity":    ret.Data.Humidity,
		"pm_2_5":      ret.Data.Pm25,
		"quality":     ret.Data.Quality,
		"windspeed":   ret.Data.WindSpeed,
	}
	/*
		{
		    "cityCode": "1796989",
		    "cityName": "Changsha",
		    "date": "2022-06-26",
		    "time": "11:56",
		    "sunrise": {
		        "seconds": 1656193617
		    },
		    "sunset": {
		        "seconds": 1656242214
		    },
		    "weather": "Clouds",
		    "temperature": 91.22,
		    "temperatureHigh": 91.22,
		    "temperatureLow": 91.22,
		    "humidity": 63,
		    "windSpeed": 9.42,
		    "windGrade": 198,
		    "pressure": 1009,
		    "visibility": 10000,
		    "source": "openweathermap",
		    "updatedAt": {
		        "seconds": 1656215810,
		        "nanos": 252467600
		    }
		}
	*/
	weatherCmd := iotredis.GetClient().Set(context.Background(), weatherKey, iotutil.ToString(cachedMaps), 0)
	if weatherCmd.Err() != nil {
		iotlogger.LogHelper.Error("天气存储异常", weatherCmd.Err())
		return
	}
	//推送天气检查
	//WeatherChan <- *ret.Data
	WeatherChan <- cachedMaps
}
