package iotconst

var (
	FIRMWARE_KEY        = "fws" //云端固件Key前缀 fws{10位}
	CUSTOM_FIRMWARE_KEY = "fwc" //开发者固件Key前缀 fwc{10位}
)

// 固件类型
var (
	FIRMWARE_TYPE_MODULE int32 = 2 //模组固件（通讯模组）
	FIRMWARE_TYPE_BLE    int32 = 4 //蓝牙固件
	FIRMWARE_TYPE_ZIGBEE int32 = 5 //Zigbee固件
	FIRMWARE_TYPE_EXTAND int32 = 6 //扩展固件
	FIRMWARE_TYPE_MCU    int32 = 3 //MCU固件

)
