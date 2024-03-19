package iotconst

var (
	NATS_STREAM_DEVICE          = "device"
	NATS_SUBJECT_INFO           = NATS_STREAM_DEVICE + ".info"       //激活或信息上报,payload有token则为激活
	NATS_SUBJECT_REPORT         = NATS_STREAM_DEVICE + ".report"     //定时上报
	NATS_SUBJECT_ONLINE         = NATS_STREAM_DEVICE + ".online"     //上线或离线
	NATS_SUBJECT_DEVICE_JOB     = NATS_STREAM_DEVICE + ".job"        //定时发送控制指令到设备
	NATS_SUBJECT_AUTH           = NATS_STREAM_DEVICE + ".auth"       //三元组
	NATS_SUBJECT_UPGRADE_REPORT = NATS_STREAM_DEVICE + ".upgrade"    //OTA升级
	NATS_SUBJECT_CONTROL        = NATS_STREAM_DEVICE + ".control"    //控制消息
	NATS_SUBJECT_CONTROL_ACK    = NATS_STREAM_DEVICE + ".controlAck" //控制结果反馈
)

var (
	NATS_STREAM_APP      = "app"
	NATS_SUBJECT_RECORDS = NATS_STREAM_APP + ".records"
	// NATS_APP_LOGIN_LOG   = NATS_STREAM_LOG + ".login"
	// NATS_APP_LOG_RECORDS = NATS_STREAM_LOG + ".records"
)

var (
	NATS_LANGUAGE                = "language"                //翻译新增
	NATS_SUBJECT_LANGUAGE_UPDATE = NATS_LANGUAGE + ".update" //定时发送控制指令到设备
)

var (
	NATS_MESSAGE                = "message"              //消息推送
	NATS_SUBJECT_MESSAGE_UPDATE = NATS_MESSAGE + ".push" //定时发送控制指令到设备
)

var (
	NATS_WEATHER              = "weather"              //消息推送
	NATS_SUBJECT_WEATHER_DATA = NATS_MESSAGE + ".data" //定时发送控制指令到设备
)

//APP Name
var NATS_APPNAME_PRODUCT = "iot_product_service"
var (
	NATS_PRODUCT_PUBLISH         = "product"
	NATS_SUBJECT_PRODUCT_PUBLISH = NATS_PRODUCT_PUBLISH + ".datasync"
)
