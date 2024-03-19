package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_smart_speaker_service/rpc/rpcclient"
	"cloud_platform/iot_smart_speaker_service/service/common"
	"cloud_platform/iot_smart_speaker_service/service/google"
	"cloud_platform/iot_smart_speaker_service/service/google/proto"
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	json "github.com/json-iterator/go"
)

func GoogleMain(c *gin.Context, userId string) {
	google.SetAgentUserIdHeader(c, userId)
	google.GetSmartHome().Handle(c)
}
func LoadDevice(id string) {
	defer iotutil.PanicHandler()
	//TODO 密钥文件需要在开放平台配置

	devList, err := common.DiscoveryDevices(id, "google")
	if err != nil {
		return
	}

	for _, dev := range devList {
		//没有配置语控的不用管
		if dev.VoiceProduct == nil || dev.VoiceProduct.VoiceProductInfo == nil || dev.VoiceProduct.VoiceProductInfo.SecretUrl == "" {
			continue
		}
		typ, ok := proto.IsSupportedType(dev.DeviceType)
		if !ok {
			continue
		}
		traits, cmds, err := initDeviceTrait(dev.VoiceProduct.FunctionMap)
		if err != nil {
			continue
		}
		//过滤掉没有实现任何特征的设备
		if len(traits) == 0 || len(cmds) == 0 {
			continue
		}

		err = google.GetSmartHome().RegisterCredentialsFile(id, dev.VoiceProduct.VoiceProductInfo.SecretUrl)
		if err != nil {
			continue
		}

		var ids []proto.OtherDeviceIds
		ids = append(ids, proto.OtherDeviceIds{DeviceId: dev.DeviceId})
		d := google.BasicDevice{
			Id: dev.DeviceId,
			Name: proto.DeviceName{
				Name:         dev.DeviceName,
				DefaultNames: []string{dev.DeviceName},
				Nicknames:    []string{dev.DeviceName},
			},
			Type:   typ,
			Traits: traits,
			Info: proto.DeviceInfo{
				Manufacturer: dev.Brand,
				Model:        dev.Model,
				HwVersion:    "1.0.0",
				SwVersion:    "1.0.0",
			},
			RoomHint:       dev.Zone, //房间名称
			CustomData:     map[string]interface{}{proto.CUSTOMDATA_PRODUCTKEY: dev.ProductKey},
			CommandInfo:    cmds,
			OtherDeviceIds: ids,
		}
		google.GetSmartHome().RegisterOrUpdateDevice(id, d)
	}
}

type IntTypeConfig struct {
	Min   int64  `json:"min"`
	Max   int64  `json:"max"`
	Step  int64  `json:"step"`
	Unit  string `json:"unit"`
	VUnit string `json:"vUnit"`
	VMin  int64  `json:"vMin"`
	VMax  int64  `json:"vMax"`
	VStep int64  `json:"vStep"`
}

type ValueMap struct {
	//Val             int     `json:"val"`
	Val             interface{} `json:"val"`
	VoiceVal        string      `json:"voiceVal"`
	VoiceValSynonym string      `json:"voiceValSynonym"`
	Min             float64     `json:"min"`  //最小值,爱星物联整数映射到语控枚举时有效
	Max             float64     `json:"max"`  //最大值,爱星物联整数映射到语控枚举时有效
	Step            float64     `json:"step"` //步长,爱星物联整数映射到语控枚举时有效
}

func initTimerTrait(valueJson string) google.Trait {
	t := google.TimerTrait{
		MaxTimerLimitSec:      2 * 3600, //默认2小时
		CommandOnlyTimer:      false,
		OnExecuteChangeStart:  TimerTrait_CommandValue,
		OnExecuteChangeAdjust: TimerTrait_CommandValue,
		OnExecuteChangePause:  TimerTrait_Command,
		OnExecuteChangeResume: TimerTrait_Command,
		OnExecuteChangeCancel: TimerTrait_Command,
		OnStateHandler:        TimerTrait_State,
	}

	var conf IntTypeConfig
	err := json.UnmarshalFromString(valueJson, &conf)
	if err == nil {
		if conf.Max > 0 && conf.Max > conf.Min {
			t.MaxTimerLimitSec = conf.Max * 60 //分钟转为秒钟
		}
	}
	return t
}

func initHumidityTrait(valueJson string) google.Trait {
	t := google.HumidityTrait{
		HumiditySetpointRange:      google.HumiditySetpointRange{MinPercent: 30, MaxPercent: 80},
		CommandOnlyHumiditySetting: false,
		QueryOnlyHumiditySetting:   false,
		OnHumidityChange:           HumidityTrait_Command,
		OnHumidityRelativeCommand:  HumidityTrait_HumidityRelativeCommand,
		OnStateHandler:             HumidityTrait_State,
	}
	var conf IntTypeConfig
	err := json.UnmarshalFromString(valueJson, &conf)
	if err == nil {
		if conf.Max > 0 && conf.Max > conf.Min {
			t.HumiditySetpointRange.MinPercent = float32(conf.Min)
			t.HumiditySetpointRange.MaxPercent = float32(conf.Max)
		}
	}
	return t
}

func initVolumeTrait(valueJson string) google.Trait {
	t := google.VolumeTrait{
		VolumeMaxLevel:                0,
		VolumeCanMuteAndUnmute:        false,
		VolumeDefaultPercentage:       67,
		LevelStepSize:                 1,
		CommandOnlyVolume:             false,
		OnExecuteChangeMute:           VolumeTrait_CommandMute,
		OnExecuteChangeSetVolume:      VolumeTrait_CommandSetVolume,
		OnExecuteChangeVolumeRelative: VolumeTrait_CommandVolumeRelative,
		OnStateHandler:                VolumeTrait_State,
	}
	var conf IntTypeConfig
	err := json.UnmarshalFromString(valueJson, &conf)
	if err == nil {
		if conf.Max > 0 && conf.Max > conf.Min {
			t.VolumeMaxLevel = int(conf.Max)
			t.LevelStepSize = int(conf.Step)
		}
	}
	return t
}

func initModeTrait(attr string, attrSynonym []string, valueJson string, lang string, dataType string) (google.Modes, map[string]int, map[string]google.NumberRange, string, error) {
	if len(attrSynonym) == 0 {
		attrSynonym = append(attrSynonym, attr)
	}
	nameValues := google.NameValues{
		NameSynonym: attrSynonym,
		Lang:        lang,
	}
	var vm []ValueMap
	if err := json.UnmarshalFromString(valueJson, &vm); err != nil {
		return google.Modes{}, nil, nil, "", err
	}

	var Default string
	var mapValue map[string]int = nil
	var mapNumberRange map[string]google.NumberRange = nil
	if dataType == "ENUM" {
		mapValue = make(map[string]int)
		for i, v := range vm {
			nVal, err := iotutil.ToInt64AndErr(v.Val)
			if err == nil {
				mapValue[v.VoiceVal] = int(nVal)
			}
			if i == 0 {
				Default = v.VoiceVal
			}
		}
	} else {
		mapNumberRange = make(map[string]google.NumberRange)
		for i, v := range vm {
			mapNumberRange[v.VoiceVal] = google.NumberRange{Min: v.Min, Max: v.Max, Step: v.Step}
			if i == 0 {
				Default = v.VoiceVal
			}
		}
	}

	var settings []google.Settings
	for _, v := range vm {
		settingSynonym := strings.Split(v.VoiceValSynonym, ";")
		if len(settingSynonym) == 0 {
			settingSynonym = append(settingSynonym, v.VoiceVal)
		}
		s := google.Settings{
			SettingName:   v.VoiceVal,
			SettingValues: []google.SettingValues{{SettingSynonym: settingSynonym, Lang: lang}},
		}
		settings = append(settings, s)
	}
	m := google.Modes{
		Name:       attr,
		NameValues: []google.NameValues{nameValues},
		Settings:   settings,
		Ordered:    true,
	}
	return m, mapValue, mapNumberRange, Default, nil
}

func initTogglesTrait(attr string, attrSynonym []string, lang string) google.Toggles {
	if len(attrSynonym) == 0 {
		attrSynonym = append(attrSynonym, attr)
	}
	nameValues := google.NameValues{
		NameSynonym: attrSynonym,
		Lang:        lang,
	}
	t := google.Toggles{
		Name:       attr,
		NameValues: []google.NameValues{nameValues},
	}
	return t
}

func initDeviceTrait(funcMap []*protosService.OpmVoiceProductMap) ([]google.Trait, map[string]*google.CommandInfo, error) {
	if len(funcMap) == 0 {
		return nil, nil, errors.New("no voice map")
	}
	//待加入缓存，读取配置
	config, err := rpcclient.ClienOpmVoiceProductMapService.Lists(context.Background(), &protosService.OpmVoiceProductMapListRequest{
		Query: &protosService.OpmVoiceProductMap{
			VoiceProductId: funcMap[0].VoiceProductId,
		},
	})
	if err != nil {
		return nil, nil, err
	}
	if config == nil || len(config.Data) == 0 {
		return nil, nil, errors.New("no google voice map ")
	}
	mapConfig := make(map[string]*protosService.OpmVoiceProductMap)
	for _, v := range config.GetData() {
		mapConfig[v.VoiceCode] = v
	}

	mapCmd := make(map[string]*google.CommandInfo)

	var ret []google.Trait
	var modes []google.Modes
	var toggles []google.Toggles
	for _, v := range funcMap {
		var t google.Trait
		vv, ok := mapConfig[v.VoiceCode]
		if !ok {
			continue
		}
		mapCmd[v.VoiceCode] = &google.CommandInfo{
			Dpid:      uint8(v.AttrDpid),
			Name:      v.AttrCode,
			DataType:  v.DataType,
			TraitName: vv.Trait,
		}
		var voiceSynonym []string
		if strings.TrimSpace(vv.VoiceSynonym) == "" {
			voiceSynonym = nil
		} else {
			voiceSynonym = strings.Split(vv.VoiceSynonym, ";")
		}

		switch vv.Trait {
		case proto.ACTION_DEVICES_TRAITS_ONOFF:
			t = google.OnOffTrait{
				CommandOnlyOnOff: false,
				QueryOnlyOnOff:   false,
				OnExecuteChange:  OnOffTrait_Command,
				OnStateHandler:   OnOffTrait_State,
			}
		case proto.ACTION_DEVICES_TRAITS_MODES:
			if m, mapValue, mapNumberRange, def, err := initModeTrait(vv.VoiceCode, voiceSynonym, vv.ValueMap, "en", v.DataType); err == nil {
				mapCmd[v.VoiceCode].MapValue = mapValue
				mapCmd[v.VoiceCode].MapNumberRange = mapNumberRange
				mapCmd[v.VoiceCode].Default = def
				modes = append(modes, m)
			}
		case proto.ACTION_DEVICES_TRAITS_TOGGLES:
			tt := initTogglesTrait(v.AttrCode, voiceSynonym, "en")
			toggles = append(toggles, tt)
		case proto.ACTION_DEVICES_TRAITS_TIMER:
			t = initTimerTrait(v.ValueMap)
		case proto.ACTION_DEVICES_TRAITS_OPENCLOSE:
			t = google.OpenCloseTrait{
				DiscreteOnlyOpenClose: false,
				QueryOnlyOpenClose:    false,
				OnExecuteChange:       OpenClose_Command,
				OnStateHandler:        OpenClose_State,
			}
		case proto.ACTION_DEVICES_TRAITS_VOLUME:
			t = initVolumeTrait(v.ValueMap)
		case proto.ACTION_DEVICES_TRAITS_BRIGHTNESS:
			t = google.BrightnessTrait{
				CommandOnlyBrightness:      false,
				OnBrightnessChange:         BrightnessTrait_BrightnessAbsoluteCommand,
				OnBrightnessRelativeChange: BrightnessTrait_BrightnessRelativeCommand,
				OnStateHandler:             BrightnessTrait_State,
			}
		case proto.ACTION_DEVICES_TRAITS_HUMIDITYSETTING:
			t = initHumidityTrait(v.ValueMap)
		default:
			iotlogger.LogHelper.Errorf("voice code  [%s] not supported.", v.VoiceCode)
		}
		if t != nil {
			ret = append(ret, t)
		}
	}

	if len(modes) > 0 {
		t := google.ModesTrait{
			AvailableModes:   modes,
			CommandOnlyModes: false,
			QueryOnlyModes:   false,
			OnExecuteChange:  ModesTrait_Command,
			OnStateHandler:   ModesTrait_State,
		}
		ret = append(ret, t)
	}
	if len(toggles) > 0 {
		t := google.TogglesTrait{
			AvailableToggles:   toggles,
			CommandOnlyToggles: false,
			QueryOnlyToggles:   false,
			OnExecuteChange:    TogglesTrait_Command,
			OnStateHandler:     TogglesTrait_State,
		}
		ret = append(ret, t)
	}
	return ret, mapCmd, nil
}

func init() {
	google.GetSmartHome().AddAspects(google.IntentAspect{
		Intent: google.IntentSync,
		Func:   LoadDevice,
	})
}
