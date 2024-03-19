package scene_executor

import (
	"cloud_platform/iot_intelligence_service/service/scene_executor/valscene"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"math"
)

// ObserverWeatherItems 天气管理
type ObserverWeatherItems struct {
	observerList []observer
}

func (s *ObserverWeatherItems) initSub() {
	//订阅天气变化
	//分发天气订阅
	//sub ret city weather
	go func() {
		for {
			select {
			case job := <-WeatherChan:
				iotlogger.LogHelper.Info("weather:", job)
				//s.notifyByCity(job.CityName, &job)
				s.notifyByCity(iotutil.ToString(job["city"]), job)
			}
		}
	}()

}
func (s *ObserverWeatherItems) register(o observer) (bool, error) {
	if s.observerList == nil {
		s.observerList = make([]observer, 0)
	}
	s.observerList = append(s.observerList, o)
	return true, nil
}
func (s *ObserverWeatherItems) deregister(o observer) (bool, error) {
	s.removeFormSlice(o)
	return true, nil
}
func (s *ObserverWeatherItems) removeFormSlice(o observer) {
	olen := len(s.observerList)
	for i, obs := range s.observerList {
		if obs.getRuleId() == o.getRuleId() {
			//替换数组对象的位置
			s.observerList[olen-1], s.observerList[i] = s.observerList[i], s.observerList[olen-1]
			//移除当前对象
			s.observerList = s.observerList[:olen-1]
			return
		}
	}
	return
}
func (s *ObserverWeatherItems) notifyByCity(city string, weather map[string]interface{}) { // *protosService.WeatherData) {
	if !s.check(city, weather) {
		//当前天气无变化
		return
	}
	//天气发送变化通知所有观察者
	for _, o := range s.observerList {
		if city == o.getKey() {
			o.run(weather)
		}
	}
}
func (s *ObserverWeatherItems) check(city string, weather map[string]interface{}) bool {
	//当前天气数据是否发生了改变
	//与历史上一次天气进行对比，如果上次天气为空，则直接推送，如天气发生变化推送
	//天气更新时间（暂无考虑）
	//cachedMaps := map[string]interface{}{
	//	"time":        ret.Data.UpdatedAt,
	//	"weather":     weather,
	//	"sun":         isSunrise,
	//	"temperature": ret.Data.Temperature,
	//	"humidity":    ret.Data.Humidity,
	//	"pm_2_5":      ret.Data.Pm25,
	//	"quality":     ret.Data.Quality,
	//	"windspeed":   ret.Data.WindSpeed,
	//}
	return true
}

// 天气变化
type observer interface {
	run(weather map[string]interface{}) bool
	getRuleId() string
	getKey() string
}

// WeatherObserver 天气观察者
type WeatherObserver struct {
	id             string
	city           string
	weatherType    int32
	weatherCompare int32
	weatherValue   interface{}
}

// weather map[string]interface{}
func (w WeatherObserver) run(weather map[string]interface{}) bool {
	if valscene.Gengine == nil {
		return false
	}
	if valscene.WeatherRuleBuilder == nil {
		return false
	}
	weatherKey := valscene.WeatherType[w.weatherType]
	isPush := false
	weatherType := iotconst.WeatherType(w.weatherType)
	switch weatherType {
	case iotconst.WEATHER_TYPE_WEATHER, iotconst.WEATHER_TYPE_SUN:
		switch w.weatherCompare {
		case 2: //等于
			if v, ok := weather[weatherKey]; ok {
				isPush = v == w.weatherValue
			}
		}
	case iotconst.WEATHER_TYPE_TEMPERATURE, iotconst.WEATHER_TYPE_HUMIDITY, iotconst.WEATHER_TYPE_PM25, iotconst.WEATHER_TYPE_WINDSPEED:
		if v, ok := weather[weatherKey]; ok {
			val, err1 := iotutil.ToFloat64Err(w.weatherValue)
			cVal, err2 := iotutil.ToFloat64Err(v)
			if err1 == nil && err2 == nil {
				return false
			}
			val = math.Round(val)
			cVal = math.Round(cVal)
			switch w.weatherCompare {
			case 1: //小于
				isPush = cVal < val
			case 2: //等于
				isPush = cVal == val
			case 3: //大于
				isPush = cVal > val
			}
		}
	}
	if isPush {
		err := valscene.Gengine.ExecuteSelectedRules(valscene.WeatherRuleBuilder, []string{w.getRuleId()})
		if err != nil {
			return false
		}
	}
	return true
}

func (w WeatherObserver) getRuleId() string {
	return w.id
}
func (w WeatherObserver) getKey() string {
	//天气中的Key是城市
	return w.city
}
