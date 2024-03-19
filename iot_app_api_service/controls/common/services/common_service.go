package services

import (
	"cloud_platform/iot_app_api_service/controls"
	"cloud_platform/iot_app_api_service/controls/common/entitys"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"go-micro.dev/v4/metadata"
)

type CommonService struct {
	Ctx context.Context
}

func (s CommonService) SetContext(ctx context.Context) CommonService {
	s.Ctx = ctx
	return s
}

// DictList 字典列表查询
func (s CommonService) DictList(dictTypeList []string) ([]*entitys.TConfigDictData, error) {
	lang, _ := metadata.Get(s.Ctx, "lang")
	rep, err := rpc.ConfigDictDataService.Lists(context.Background(), &protosService.ConfigDictDataListRequest{
		Query: &protosService.ConfigDictData{
			DictTypeList: dictTypeList,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, err
	}

	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_CONFIG_DICT_DATA).Result()
	if err != nil {
		langMap = make(map[string]string)
		//异常不处理
		iotlogger.LogHelper.Errorf("字典缓存读取异常, %s", err.Error())
	}
	result := []*entitys.TConfigDictData{}
	for _, v := range rep.Data {
		item := entitys.ConfigDictData_pb2db(v)
		fKey := fmt.Sprintf("%s_%s-%d_name", lang, item.DictType, item.DictValue)
		//兼容逻辑，字典key修改为DictValue
		if _, ok := langMap[fKey]; ok {
			item.DictLabel = iotutil.MapGetStringVal(langMap[fKey], item.DictLabel)
		} else {
			fKey = fmt.Sprintf("%s_%s-%d_name", lang, item.DictType, item.DictCode)
			item.DictLabel = iotutil.MapGetStringVal(langMap[fKey], item.DictLabel)
		}

		result = append(result, item)
	}
	return result, nil
}

// CustomLangList 自定义翻译资源数据
func (s CommonService) CustomLangList(appKey string) (interface{}, error) {
	rep, err := rpc.LangCustomResourcesService.Lists(s.Ctx, &protosService.LangCustomResourcesListRequest{
		Query: &protosService.LangCustomResources{
			AppKey: appKey,
		},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, err
	}
	result := s.LangCustomResultConvert(rep.Data)
	return result, nil
}

// LangCustomResultConvert 翻译数据格式转换
func (s CommonService) LangCustomResultConvert(res []*protosService.LangCustomResources) map[string]map[string]interface{} {
	dateList := map[string]map[string]interface{}{}
	for _, v := range res {
		lang := v.Lang
		_, ok := dateList[lang]
		if !ok {
			dateList[lang] = map[string]interface{}{}
		}
		//list = append(list, map[string]interface{}{
		//	v.Code:v.Value,
		//})
		dateList[lang][v.Code] = v.Value
	}
	return dateList
}

// LangResultConvert 翻译数据格式转换
func (s CommonService) LangResultConvert(res []*protosService.LangResources) map[string]map[string]interface{} {
	dateList := map[string]map[string]interface{}{}
	for _, v := range res {
		lang := v.Lang
		_, ok := dateList[lang]
		if !ok {
			dateList[lang] = map[string]interface{}{}
		}
		//list = append(list, map[string]interface{}{
		//	v.Code:v.Value,
		//})
		dateList[lang][v.Code] = v.Value
	}
	return dateList
}

// RegionList 区域数据列表
func (s CommonService) RegionList(lang, ip string) ([]*entitys.SysRegionServerEntitysList, error) {
	//geo, err := iotutil.Geoip(ip, config.Global.IpService.QueryUrl, config.Global.IpService.AppCode)
	geo, err := controls.Geoip(ip)
	country := geo.Country
	result := []*entitys.SysRegionServerEntitysList{}
	//增加状态判断查询（用于启用关闭区域操作）
	rep, err := rpc.SysRegionServerService.Lists(context.Background(), &protosService.SysRegionServerListRequest{
		Query: &protosService.SysRegionServer{Enabled: 1},
	})
	if err != nil {
		return nil, err
	}
	if rep.Code != 200 {
		return nil, err
	}
	if len(rep.Data) == 0 {
		return nil, err
	}
	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_CONFIG_REGION_SERVER).Result()
	if err != nil {
		langMap = make(map[string]string)
		//异常不处理
		iotlogger.LogHelper.Errorf("字典缓存读取异常, %s", err.Error())
	}
	haveDefaultData := false
	for _, v := range rep.Data {
		item := entitys.SysRegionServer_pb2e(v, lang)
		fKey := fmt.Sprintf("%s_%d_name", lang, item.Id)
		if country != "" && v.Describe == country {
			item.IsDefault = 1
			haveDefaultData = true
		}
		item.Describe = iotutil.MapGetStringVal(langMap[fKey], item.Describe)
		result = append(result, item)
	}
	if !haveDefaultData {
		result[0].IsDefault = 1
	}
	return result, nil
}

// GetWeather 获取天气
func (s CommonService) GetWeather(city, province string, lang string) (entitys.WeatherData, error) {
	result, err := rpc.WeatherService.CurrentByCity(context.Background(), &protosService.CityRequest{
		CityCode: "",
		CityName: city,
		Province: province,
	})
	if err != nil {
		return entitys.WeatherData{}, err
	}
	data := result.Data
	iotlogger.LogHelper.Info(data)
	if data == nil {
		return entitys.WeatherData{}, errors.New("天气数据获取异常")
	}

	//todo 空气质量等级，1 优；  2 良；  3 差
	airDegree := iotutil.GetDegreeByAqi(data.Aqi)
	//
	quality := "良"
	if lang != "zh" {
		quality = "good"
	}

	langMap, err := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_WEATHER).Result()
	if err != nil {
		langMap = make(map[string]string)
		//异常不处理
		iotlogger.LogHelper.Errorf("天气缓存读取异常, %s", err.Error())
	}
	fKey := fmt.Sprintf("%s_%s_name", lang, data.Weather)
	data.Weather = iotutil.MapGetStringVal(langMap[fKey], data.Weather)

	//翻译处理
	fKey = fmt.Sprintf("%s_air_quality_%v_name", lang, airDegree)
	qualityDesc, err := iotredis.GetClient().HGet(context.Background(),
		iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_T_CONFIG_DICT_DATA, fKey).Result()
	if err == nil && qualityDesc != "" {
		quality = qualityDesc
	}

	return entitys.WeatherData{
		Weather:  data.Weather,
		Humidity: fmt.Sprintf("%d", data.Humidity) + "%",
		Temp:     iotutil.ToString(math.Round(iotutil.HToSTemperature(data.Temperature))), //华氏温度转摄氏温度
		Quality:  quality,                                                                 //todo api取不到值
		PicType:  convertIconToPicType(data.Icon),
		Pm25:     data.Pm25,
		Pm10:     data.Pm10,
		Aqi:      data.Aqi,
	}, nil
}

// RegionInfo 区域详情
func (s CommonService) RegionInfo(lang string, regionId int64) (*entitys.SysRegionServerEntitysList, error) {
	rep, err := rpc.SysRegionServerService.FindById(context.Background(), &protosService.SysRegionServerFilter{Id: regionId})
	if err != nil {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	if rep.Code != 200 {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	if len(rep.Data) == 0 {
		return &entitys.SysRegionServerEntitysList{}, nil
	}
	return entitys.SysRegionServer_pb2e(rep.Data[0], lang), nil
}

// RoomConfigList OEM APP配置的房间数据
func (s CommonService) RoomConfigList(lang, tenantId, appKey, code string) (int, string, []entitys.TConfigDictData) {
	resultList := []entitys.TConfigDictData{}
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("RoomConfigList error")
		return -1, err.Error(), resultList
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return -1, oemAppResult.Message, resultList
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppFunctionConfig, err := rpc.ClientOemAppUiConfigService.Find(context.Background(), &protosService.OemAppUiConfigFilter{
		AppId:   oemAppInfo.Id,
		Version: oemAppInfo.Version,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("RoomConfigList error")
		return -1, err.Error(), resultList
	}
	if oemAppFunctionConfig.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppFunctionConfig.Message)
		return -1, oemAppFunctionConfig.Message, resultList
	}

	//room 和 icons
	if code == "room" {
		cacheKey := fmt.Sprintf("%s_%s_%s", tenantId,
			iotconst.HKEY_LANGUAGE_DATA_PREFIX+iotconst.LANG_OEM_APP_ROOMS, appKey)
		langMap, err := iotredis.GetClient().HGetAll(context.Background(), cacheKey).Result()
		if err != nil {
			langMap = make(map[string]string)
		}
		if oemAppFunctionConfig.Data[0].Room == "" {
			return -1, "roomConfig is empty", resultList
		}
		roomConfigs := iotutil.JsonToMapArray(oemAppFunctionConfig.Data[0].Room)
		for _, roomConfig := range roomConfigs {
			roomId := iotutil.ToInt64(roomConfig["roomId"])
			roomName := iotutil.ToString(roomConfig["roomName"])
			roomName = iotutil.MapGetStringVal(langMap[fmt.Sprintf("%s_%v_name", lang, roomId)], roomName)
			resultList = append(resultList, entitys.TConfigDictData{
				DictCode:  roomId,
				DictLabel: roomName,
				DictValue: iotutil.ToString(roomConfig["roomSort"]),
				DictType:  "",
				Listimg:   iotutil.ToString(roomConfig["roomImage"]),
			})
		}
		//根据sort进行排序
		sort.Slice(resultList, func(i, j int) bool {
			return resultList[i].DictValue < resultList[j].DictValue
		})
	} else if code == "icons" {
		if oemAppFunctionConfig.Data[0].RoomIcons == "" {
			return -1, "roomConfig is empty", resultList
		}
		roomIconsConfigs := iotutil.JsonToStringArray(oemAppFunctionConfig.Data[0].RoomIcons)
		for _, roomIconsConfig := range roomIconsConfigs {
			resultList = append(resultList, entitys.TConfigDictData{
				Listimg: roomIconsConfig,
			})
		}
	}

	return 0, "", resultList
}

// VoiceService 语音服务配置
func (s CommonService) VoiceService(appKey, voiceCode, lang string) (int, string, map[string]interface{}) {
	result := map[string]interface{}{}
	oemAppResult, err := rpc.ClientOemAppService.Find(context.Background(), &protosService.OemAppFilter{
		AppKey: appKey,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetVoice error")
		return -1, err.Error(), result
	}
	if oemAppResult.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppResult.Message)
		return -1, oemAppResult.Message, result
	}
	oemAppInfo := oemAppResult.Data[0]
	oemAppIntroduce, err := rpc.ClientOemAppIntroduceService.Find(context.Background(), &protosService.OemAppIntroduceFilter{
		AppId: oemAppInfo.Id,
		// Version:   oemAppInfo.Version,
		VoiceCode: voiceCode,
		//Status: 1,
		Lang:        lang,
		ContentType: 4,
	})
	if err != nil {
		iotlogger.LogHelper.Errorf("GetVoice error")
		return -1, err.Error(), result
	}
	if oemAppIntroduce.Code != 200 {
		iotlogger.LogHelper.Errorf(oemAppIntroduce.Message)
		return -1, oemAppIntroduce.Message, result
	}

	if oemAppIntroduce.Data == nil {
		return -1, "", result
	}
	result["content"] = oemAppIntroduce.Data[0].Content

	return 0, "", result
}

// GetFlashScreen 获取闪屏数据
func (s CommonService) GetFlashScreen(appVersion, account string, sizeType int) ([]*entitys.FlashScreen, error) {
	data := []*entitys.FlashScreen{}
	appKey, _ := metadata.Get(s.Ctx, "appKey")
	tenantId, _ := metadata.Get(s.Ctx, "tenantId")
	//无账号，不需要返回内容
	if account == "" {
		return data, nil
	}
	resp, err := rpc.ClientOemAppFlashScreenService.GetFlashScreen(context.Background(), &protosService.OemAppFlashScreenRequest{
		AppKey:   appKey,
		TenantId: tenantId,
		Version:  appVersion,
		Account:  account,
	})
	if err != nil {
		return data, err
	}
	if resp.Data == nil || len(resp.Data) == 0 {
		return data, nil
	}
	for i := range resp.Data {
		accounts := []string{}
		//如果是指定用户推送
		if resp.Data[i].PutinUser == 2 {
			//账号等于空
			if account != "" {
				respUser, err := rpc.ClientOemAppFlashScreenUserService.Lists(s.Ctx, &protosService.OemAppFlashScreenUserListRequest{
					Query: &protosService.OemAppFlashScreenUser{
						FlashScreenId: resp.Data[i].Id,
						Account:       account,
					},
				})
				if err != nil {
					return nil, err
				}
				if len(respUser.Data) == 0 {
					continue
				}
				for j := range respUser.Data {
					accounts = append(accounts, respUser.Data[j].Account)
				}
			}
		}
		imageInfo := []entitys.ImageInfo{}
		if err := json.Unmarshal([]byte(resp.Data[i].PutinImgUrls), &imageInfo); err != nil {
			return nil, err
		}
		if len(imageInfo) < sizeType-1 {
			return nil, errors.New("the size image not exist")
		}
		if resp.Data[i].PutinUser == 2 && len(accounts) == 0 {
			continue
		}
		item := &entitys.FlashScreen{
			Id:           iotutil.ToString(resp.Data[i].Id),
			StartTime:    iotutil.GetLocalTimeStr(resp.Data[i].StartTime.AsTime()),
			EndTime:      iotutil.GetLocalTimeStr(resp.Data[i].EndTime.AsTime()),
			PutinUser:    resp.Data[i].PutinUser,
			Accounts:     accounts,
			OpenPageType: int(resp.Data[i].OpenPageType),
			AppPageType:  int(resp.Data[i].AppPageType),
			OpenPageUrl:  resp.Data[i].OpenPageUrl,
			ShowImageUrl: imageInfo[sizeType-1].ImageUrl,
			ShowImageMd5: imageInfo[sizeType-1].ImageMd5,
			ShowTime:     int(resp.Data[i].ShowTime),
		}
		data = append(data, item)
	}

	return data, nil
}

// convertIconToPicType OpenWeather天气接口返回数据，天气图标去除后缀
func convertIconToPicType(icon string) string {
	if icon == "" {
		return ""
	}
	//d 白天 n代表晚上上
	icon = strings.Replace(icon, "d", "", 1)
	icon = strings.Replace(icon, "n", "", 1)
	return icon
}
