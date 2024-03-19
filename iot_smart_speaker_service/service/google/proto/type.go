package proto

var mapDeviceType map[string]DeviceType

// DeviceType 设备类别
type DeviceType string

// 设备类别表，参考链接 https://developers.google.com/assistant/smarthome/guides
const (
	NotSupported                     DeviceType = "NotSupported"
	DeviceTypeAirConditioningUnit    DeviceType = "action.devices.types.AC_UNIT"
	DeviceTypeAirCooler              DeviceType = "action.devices.types.AIRCOOLER"
	DeviceTypeAirFreshener           DeviceType = "action.devices.types.AIRFRESHENER"
	DeviceTypeAirPurifier            DeviceType = "action.devices.types.AIRPURIFIER"
	DeviceTypeAudioVideoReceiver     DeviceType = "action.devices.types.AUDIO_VIDEO_RECEIVER"
	DeviceTypeAwning                 DeviceType = "action.devices.types.AWING"
	DeviceTypeBathtub                DeviceType = "action.devices.types.BATHTUB"
	DeviceTypeBed                    DeviceType = "action.devices.types.BED"
	DeviceTypeBlender                DeviceType = "action.devices.types.BLENDER"
	DeviceTypeBlinds                 DeviceType = "action.devices.types.BLINDS"
	DeviceTypeBoiler                 DeviceType = "action.devices.types.BOILER"
	DeviceTypeCamera                 DeviceType = "action.devices.types.CAMERA"
	DeviceTypeCarbonMonoxideDetector DeviceType = "action.devices.types.CARBON_MONOXIDE_DETECTOR"
	DeviceTypeCharger                DeviceType = "action.devices.types.CHARGER"
	DeviceTypeCloset                 DeviceType = "action.devices.types.CLOSET"
	DeviceTypeCoffeeMaker            DeviceType = "action.devices.types.COFFEE_MAKER"
	DeviceTypeCooktop                DeviceType = "action.devices.types.COOKTOP"
	DeviceTypeCurtain                DeviceType = "action.devices.types.CURTAIN"
	DeviceTypeDehumidifier           DeviceType = "action.devices.types.DEHUMIDIFIER"
	DeviceTypeDehydrator             DeviceType = "action.devices.types.DEHYDRATOR"
	DeviceTypeDishwasher             DeviceType = "action.devices.types.DISHWASHER"
	DeviceTypeDoor                   DeviceType = "action.devices.types.DOOR"
	DeviceTypeDoorBell               DeviceType = "action.devices.types.DOORBELL"
	DeviceTypeDrawer                 DeviceType = "action.devices.types.DRAWER"
	DeviceTypeDryer                  DeviceType = "action.devices.types.DRYER"
	DeviceTypeFan                    DeviceType = "action.devices.types.FAN"
	DeviceTypeFaucet                 DeviceType = "action.devices.types.FAUCET"
	DeviceTypeFireplace              DeviceType = "action.devices.types.FIREPLACE"
	DeviceTypeFreezer                DeviceType = "action.devices.types.FREEZER"
	DeviceTypeFryer                  DeviceType = "action.devices.types.FRYER"
	DeviceTypeGarage                 DeviceType = "action.devices.types.GARAGE"
	DeviceTypeGate                   DeviceType = "action.devices.types.GATE"
	DeviceTypeGrill                  DeviceType = "action.devices.types.GRILL"
	DeviceTypeHeater                 DeviceType = "action.devices.types.HEATER"
	DeviceTypeHood                   DeviceType = "action.devices.types.HOOD"
	DeviceTypeHumidifier             DeviceType = "action.devices.types.HUMIDIFIER"
	DeviceTypeKettle                 DeviceType = "action.devices.types.KETTLE"
	DeviceTypeLight                  DeviceType = "action.devices.types.LIGHT"
	DeviceTypeLock                   DeviceType = "action.devices.types.LOCK"
	DeviceTypeMicrowave              DeviceType = "action.devices.types.MICROWAVE"
	DeviceTypeMop                    DeviceType = "action.devices.types.MOP"
	DeviceTypeMower                  DeviceType = "action.devices.types.MOWER"
	DeviceTypeMulticooker            DeviceType = "action.devices.types.MULTICOOKER"
	DeviceTypeNetwork                DeviceType = "action.devices.types.NETWORK"
	DeviceTypeOutlet                 DeviceType = "action.devices.types.OUTLET"
	DeviceTypeOven                   DeviceType = "action.devices.types.OVEN"
	DeviceTypePergola                DeviceType = "action.devices.types.PERGOLA"
	DeviceTypePetFeeder              DeviceType = "action.devices.types.PETFEEDER"
	DeviceTypePressureCooker         DeviceType = "action.devices.types.PRESSURECOOKER"
	DeviceTypeRadiator               DeviceType = "action.devices.types.RADIATOR"
	DeviceTypeRefrigerator           DeviceType = "action.devices.types.REFRIGERATOR"
	DeviceTypeRemoteControl          DeviceType = "action.devices.types.REMOTECONTROL"
	DeviceTypeRouter                 DeviceType = "action.devices.types.ROUTER"
	DeviceTypeScene                  DeviceType = "action.devices.types.SCENE"
	DeviceTypeSecuritySystem         DeviceType = "action.devices.types.SECURITYSYSTEM"
	DeviceTypeSensor                 DeviceType = "action.devices.types.SENSOR"
	DeviceTypeSettop                 DeviceType = "action.devices.types.SETTOP"
	DeviceTypeShower                 DeviceType = "action.devices.types.SHOWER"
	DeviceTypeShutter                DeviceType = "action.devices.types.SHUTTER"
	DeviceTypeSmokeDetector          DeviceType = "action.devices.types.SMOKE_DETECTOR"
	DeviceTypeSoundBar               DeviceType = "action.devices.types.SOUNDBAR"
	DeviceTypeSousVide               DeviceType = "action.devices.types.SOUSVIDE"
	DeviceTypeSpeaker                DeviceType = "action.devices.types.SPEAKER"
	DeviceTypeSprinkler              DeviceType = "action.devices.types.SPRINKLER"
	DeviceTypeStandMixer             DeviceType = "action.devices.types.STANDMIXER"
	DeviceTypeStreamingBox           DeviceType = "action.devices.types.STREAMING_BOX"
	DeviceTypeStreamingSoundBar      DeviceType = "action.devices.types.STREAMING_SOUNDBAR"
	DeviceTypeStreamingStick         DeviceType = "action.devices.types.STREAMING_STICK"
	DeviceTypeSwitch                 DeviceType = "action.devices.types.SWITCH"
	DeviceTypeThermostat             DeviceType = "action.devices.types.THERMOSTAT"
	DeviceTypeTV                     DeviceType = "action.devices.types.TV"
	DeviceTypeVacuum                 DeviceType = "action.devices.types.VACUUM"
	DeviceTypeValve                  DeviceType = "action.devices.types.VALVE"
	DeviceTypeWasher                 DeviceType = "action.devices.types.WASHER"
	DeviceTypeWaterHeater            DeviceType = "action.devices.types.WATERHEATER"
	DeviceTypeWaterPurifier          DeviceType = "action.devices.types.WATERPURIFIER"
	DeviceTypeWaterSoftener          DeviceType = "action.devices.types.WATERSOFTENER"
	DeviceTypeWindow                 DeviceType = "action.devices.types.WINDOW"
	DeviceTypeYogurtMaker            DeviceType = "action.devices.types.YOGURTMAKER"
)

func IsSupportedType(typ string) (DeviceType, bool) {
	if typ, ok := mapDeviceType[typ]; ok {
		return typ, true
	}
	return NotSupported, false
}

func init() {
	mapDeviceType = make(map[string]DeviceType)
	mapDeviceType["NotSupported"] = NotSupported
	mapDeviceType["action.devices.types.AC_UNIT"] = DeviceTypeAirConditioningUnit
	mapDeviceType["action.devices.types.AIRCOOLER"] = DeviceTypeAirCooler
	mapDeviceType["action.devices.types.AIRFRESHENER"] = DeviceTypeAirFreshener
	mapDeviceType["action.devices.types.AIRPURIFIER"] = DeviceTypeAirPurifier
	mapDeviceType["action.devices.types.AUDIO_VIDEO_RECEIVER"] = DeviceTypeAudioVideoReceiver
	mapDeviceType["action.devices.types.AWING"] = DeviceTypeAwning
	mapDeviceType["action.devices.types.BATHTUB"] = DeviceTypeBathtub
	mapDeviceType["action.devices.types.BED"] = DeviceTypeBed
	mapDeviceType["action.devices.types.BLENDER"] = DeviceTypeBlender
	mapDeviceType["action.devices.types.BLINDS"] = DeviceTypeBlinds
	mapDeviceType["action.devices.types.BOILER"] = DeviceTypeBoiler
	mapDeviceType["action.devices.types.CAMERA"] = DeviceTypeCamera
	mapDeviceType["action.devices.types.CARBON_MONOXIDE_DETECTOR"] = DeviceTypeCarbonMonoxideDetector
	mapDeviceType["action.devices.types.CHARGER"] = DeviceTypeCharger
	mapDeviceType["action.devices.types.CLOSET"] = DeviceTypeCloset
	mapDeviceType["action.devices.types.COFFEE_MAKER"] = DeviceTypeCoffeeMaker
	mapDeviceType["action.devices.types.COOKTOP"] = DeviceTypeCooktop
	mapDeviceType["action.devices.types.CURTAIN"] = DeviceTypeCurtain
	mapDeviceType["action.devices.types.DEHUMIDIFIER"] = DeviceTypeDehumidifier
	mapDeviceType["action.devices.types.DEHYDRATOR"] = DeviceTypeDehydrator
	mapDeviceType["action.devices.types.DISHWASHER"] = DeviceTypeDishwasher
	mapDeviceType["action.devices.types.DOOR"] = DeviceTypeDoor
	mapDeviceType["action.devices.types.DOORBELL"] = DeviceTypeDoorBell
	mapDeviceType["action.devices.types.DRAWER"] = DeviceTypeDrawer
	mapDeviceType["action.devices.types.DRYER"] = DeviceTypeDryer
	mapDeviceType["action.devices.types.FAN"] = DeviceTypeFan
	mapDeviceType["action.devices.types.FAUCET"] = DeviceTypeFaucet
	mapDeviceType["action.devices.types.FIREPLACE"] = DeviceTypeFireplace
	mapDeviceType["action.devices.types.FREEZER"] = DeviceTypeFreezer
	mapDeviceType["action.devices.types.FRYER"] = DeviceTypeFryer
	mapDeviceType["action.devices.types.GARAGE"] = DeviceTypeGarage
	mapDeviceType["action.devices.types.GATE"] = DeviceTypeGate
	mapDeviceType["action.devices.types.GRILL"] = DeviceTypeGrill
	mapDeviceType["action.devices.types.HEATER"] = DeviceTypeHeater
	mapDeviceType["action.devices.types.HOOD"] = DeviceTypeHood
	mapDeviceType["action.devices.types.HUMIDIFIER"] = DeviceTypeHumidifier
	mapDeviceType["action.devices.types.KETTLE"] = DeviceTypeKettle
	mapDeviceType["action.devices.types.LIGHT"] = DeviceTypeLight
	mapDeviceType["action.devices.types.LOCK"] = DeviceTypeLock
	mapDeviceType["action.devices.types.MICROWAVE"] = DeviceTypeMicrowave
	mapDeviceType["action.devices.types.MOP"] = DeviceTypeMop
	mapDeviceType["action.devices.types.MOWER"] = DeviceTypeMower
	mapDeviceType["action.devices.types.MULTICOOKER"] = DeviceTypeMulticooker
	mapDeviceType["action.devices.types.NETWORK"] = DeviceTypeNetwork
	mapDeviceType["action.devices.types.OUTLET"] = DeviceTypeOutlet
	mapDeviceType["action.devices.types.OVEN"] = DeviceTypeOven
	mapDeviceType["action.devices.types.PERGOLA"] = DeviceTypePergola
	mapDeviceType["action.devices.types.PETFEEDER"] = DeviceTypePetFeeder
	mapDeviceType["action.devices.types.PRESSURECOOKER"] = DeviceTypePressureCooker
	mapDeviceType["action.devices.types.RADIATOR"] = DeviceTypeRadiator
	mapDeviceType["action.devices.types.REFRIGERATOR"] = DeviceTypeRefrigerator
	mapDeviceType["action.devices.types.REMOTECONTROL"] = DeviceTypeRemoteControl
	mapDeviceType["action.devices.types.ROUTER"] = DeviceTypeRouter
	mapDeviceType["action.devices.types.SCENE"] = DeviceTypeScene
	mapDeviceType["action.devices.types.SECURITYSYSTEM"] = DeviceTypeSecuritySystem
	mapDeviceType["action.devices.types.SENSOR"] = DeviceTypeSensor
	mapDeviceType["action.devices.types.SETTOP"] = DeviceTypeSettop
	mapDeviceType["action.devices.types.SHOWER"] = DeviceTypeShower
	mapDeviceType["action.devices.types.SHUTTER"] = DeviceTypeShutter
	mapDeviceType["action.devices.types.SMOKE_DETECTOR"] = DeviceTypeSmokeDetector
	mapDeviceType["action.devices.types.SOUNDBAR"] = DeviceTypeSoundBar
	mapDeviceType["action.devices.types.SOUSVIDE"] = DeviceTypeSousVide
	mapDeviceType["action.devices.types.SPEAKER"] = DeviceTypeSpeaker
	mapDeviceType["action.devices.types.SPRINKLER"] = DeviceTypeSprinkler
	mapDeviceType["action.devices.types.STANDMIXER"] = DeviceTypeStandMixer
	mapDeviceType["action.devices.types.STREAMING_BOX"] = DeviceTypeStreamingBox
	mapDeviceType["action.devices.types.STREAMING_SOUNDBAR"] = DeviceTypeStreamingSoundBar
	mapDeviceType["action.devices.types.STREAMING_STICK"] = DeviceTypeStreamingStick
	mapDeviceType["action.devices.types.SWITCH"] = DeviceTypeSwitch
	mapDeviceType["action.devices.types.THERMOSTAT"] = DeviceTypeThermostat
	mapDeviceType["action.devices.types.TV"] = DeviceTypeTV
	mapDeviceType["action.devices.types.VACUUM"] = DeviceTypeVacuum
	mapDeviceType["action.devices.types.VALVE"] = DeviceTypeValve
	mapDeviceType["action.devices.types.WASHER"] = DeviceTypeWasher
	mapDeviceType["action.devices.types.WATERHEATER"] = DeviceTypeWaterHeater
	mapDeviceType["action.devices.types.WATERPURIFIER"] = DeviceTypeWaterPurifier
	mapDeviceType["action.devices.types.WATERSOFTENER"] = DeviceTypeWaterSoftener
	mapDeviceType["action.devices.types.WINDOW"] = DeviceTypeWindow
	mapDeviceType["action.devices.types.YOGURTMAKER"] = DeviceTypeYogurtMaker
}
