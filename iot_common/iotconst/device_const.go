package iotconst

var (
	Use_Type_Device_Real_Test int32 = 1 //虚拟设备
	Use_Type_Device_Normal    int32 = 0 //正常设备

	//灰度发布模式
	OTA_GRAY_TYPE_SCALE  int32 = 1 //比例发布
	OTA_GRAY_TYPE_NUMBER int32 = 2 //数量发布
	OTA_GRAY_TYPE_DEVICE int32 = 3 //指定设备

	//发布模式
	OTA_PUBLISH_ALL  int32 = 1 //全量发布
	OTA_PUBLISH_GRAY int32 = 2 //灰度发布

	//自定义功能名称类型（属性名称自定义， 属性值名称自定义）
	FUNCTION_CUSTOM_PROPERTY_SET       int32 = 1
	FUNCTION_CUSTOM_PROPERTY_VALUE_SET int32 = 2
	//产品的功能值类型
	FUNCTION_DATA_TYPE_ENUM   = "ENUM"
	FUNCTION_DATA_TYPE_INT    = "INT"
	FUNCTION_DATA_TYPE_TEXT   = "TEXT"
	FUNCTION_DATA_TYPE_BOOL   = "BOOL"
	FUNCTION_DATA_TYPE_FAULT  = "FAULT"
	FUNCTION_DATA_TYPE_FLOAT  = "FLOAT"
	FUNCTION_DATA_TYPE_DOUBLE = "DOUBLE"
	FUNCTION_DATA_TYPE_DATE   = "DATE"
	FUNCTION_DATA_TYPE_JSON   = "JSON"
)
