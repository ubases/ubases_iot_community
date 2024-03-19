package service

import (
	"cloud_platform/iot_weather_service/config"
	"cloud_platform/iot_weather_service/service/aqicn"
	"cloud_platform/iot_weather_service/service/cache"
	owm "cloud_platform/iot_weather_service/service/openweathermap"
	"cloud_platform/iot_weather_service/service/yiketianqi"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/mozillazg/go-pinyin"
)

var _suberonce sync.Once
var _subersingle *Subscriber

func GetSubscriber() *Subscriber {
	_suberonce.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		_subersingle = &Subscriber{
			CityMap:  make(map[string]CityStore),
			Interval: time.Duration(config.Global.Weather.Interval) * time.Second,
			ctx:      ctx,
			cancel:   cancel,
		}
	})
	return _subersingle
}

type Subscriber struct {
	mu sync.RWMutex
	//订阅的城市列表,城市名->最近更新时间
	//todo 持久化支持
	CityMap  map[string]CityStore
	Interval time.Duration
	ctx      context.Context
	cancel   context.CancelFunc
}

// 城市存储
type CityStore struct {
	UpdateTime time.Time
	Province   string
}

func (s *Subscriber) UpdateCityMap(city, province string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CityMap[city] = CityStore{UpdateTime: time.Now(), Province: province}
}

func (s *Subscriber) CurrentByCity(cityName, province string) (*proto.WeatherData, error) {
	//处理城市名字
	var data proto.WeatherData
	err := cache.RepoCache.Get(cityName, &data)
	if err == nil && time.Now().Sub(data.UpdatedAt.AsTime()) <= s.Interval {
		return &data, nil
	}
	pinyinCityName := cityName
	if HasHans(cityName) { //去掉市县区
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "市")
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "县")
		pinyinCityName = strings.TrimSuffix(pinyinCityName, "区")
		pinyinCityName = strings.Join(pinyin.LazyConvert(pinyinCityName, nil), "")
	}
	dataPtr, err := s.Current(pinyinCityName, cityName, province)
	if err != nil {
		dataPtr, err = s.Current(cityName, cityName, province)
		if err != nil {
			return nil, err
		}
	}
	return dataPtr, nil
}

func (s *Subscriber) Current(cityName string, cacheCityName, province string) (*proto.WeatherData, error) {
	cwd, err := owm.NewCurrent("F", "EN", config.Global.Weather.ApiKey)
	if err != nil {
		return nil, err
	}
	if err = cwd.CurrentByName(cityName); err != nil {
		return nil, err
	}
	data, _ := WeatherData_topb(cwd)

	//获取空气质量
	airQualityData := aqicn.CurrentAriQualityData{
		Settings: aqicn.NewSettings(),
	}
	if airQualityData.CurrentByName(cityName) == nil {
		data.Aqi, _ = iotutil.ToFloat64Err(airQualityData.Aqi)
		if airQualityData.IAqi != nil {
			data.Pm10 = airQualityData.IAqi.Pm10.V
			data.Pm25 = airQualityData.IAqi.Pm25.V
		}
	} else {
		//使用省份名称重试
		if province != "" && airQualityData.CurrentByName(province) == nil {
			data.Aqi, _ = iotutil.ToFloat64Err(airQualityData.Aqi)
			if airQualityData.IAqi != nil {
				data.Pm10 = airQualityData.IAqi.Pm10.V
				data.Pm25 = airQualityData.IAqi.Pm25.V
			}
		}
	}

	if data.Aqi < 0.1 || data.Pm25 < 0.0000001 {
		w, _ := yiketianqi.GetTianqiWeatherByCity(cacheCityName)
		if w != nil {
			if data.Aqi < 0.1 { //补AQI
				faqi, _ := strconv.ParseFloat(w.Air, 32)
				data.Aqi = faqi
				data.Quality = w.AirLevel
			}
			if data.Pm25 < 0.0000001 { //补PM2.5
				fpm25, _ := strconv.ParseFloat(w.AirPm25, 32)
				data.Pm25 = fpm25
			}
		}
	}

	err = cache.RepoCache.Set(cacheCityName, data)
	s.UpdateCityMap(cityName, province)

	return data, nil
}

func (s *Subscriber) WatchWeather() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			s.CheckUpdate()
			time.Sleep(time.Second * 5)
		}
	}
}

func (s *Subscriber) CheckUpdate() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.CityMap {
		if time.Now().Sub(v.UpdateTime) > s.Interval {
			go s.Current(k, k, v.Province)
		}
	}
}

func (s *Subscriber) AddSubscribe(city, province string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.CityMap[city]; !ok {
		s.CityMap[city] = CityStore{Province: province, UpdateTime: time.Time{}}
	}
}

func (s *Subscriber) UnSubscribe(citys []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range citys {
		delete(s.CityMap, v)
	}
}

func (s *Subscriber) Close() {
	s.cancel()
}

func HasHans(str string) bool {
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			return true
		}
	}
	return false
}
