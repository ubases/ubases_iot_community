package iotconst

var (
	SVR_PREFIX                = "ai."                                    //单机部署多个云平台使用
	IOT_APP_USER_SERVICE      = SVR_PREFIX + "iot.app.user.service"      //用户管理服务
	IOT_PRODUCT_SERVICE       = SVR_PREFIX + "iot.product.service"       //产品管理服务
	IOT_SYSTEM_SERVICE        = SVR_PREFIX + "iot.system.service"        //系统管理服务
	IOT_DEVICE_SERVICE        = SVR_PREFIX + "iot.device.service"        //物联管理服务
	IOT_BASIC_SERVICE         = SVR_PREFIX + "iot.basic.service"         //基础数据管理服务
	IOT_MQTT_SERVICE          = SVR_PREFIX + "iot.mqtt.service"          //MQTT服务
	IOT_MQTT_DATA_SERVICE     = SVR_PREFIX + "iot.mqtt.data.service"     //MQTT服务 数据服务
	IOT_OPEN_SYSTEM_SERVICE   = SVR_PREFIX + "iot.open.system.service"   //开放平台服务
	IOT_WEATHER_SERVICE       = SVR_PREFIX + "iot.weather.service"       //天气服务
	IOT_INTELLIGENCE_SERVICE  = SVR_PREFIX + "iot.intelligence.service"  //智能场景
	IOT_MESSAGE_SERVICE       = SVR_PREFIX + "iot.message.service"       //APP消息（厂商推送）
	IOT_APP_NOTIFIER_SERVICE  = SVR_PREFIX + "iot.app.notifier.service"  //APP通知消息
	IOT_APP_OEM_SERVICE       = SVR_PREFIX + "iot.app.oem.service"       //OEMapp服务
	IOT_STATISTICS_SERVICE    = SVR_PREFIX + "iot.statistics.service"    //数据统计分析服务
	IOT_LOG_SERVICE           = SVR_PREFIX + "iot.log.service"           //日志服务
	IOT_SMART_SPEAKER_SERVICE = SVR_PREFIX + "iot.smart.speaker.service" //oauth服务
	IOT_LVGL_SERVICE          = SVR_PREFIX + "iot.lvgl.service"          //oauth服务
	IOT_PANEL_DESIGN_SERVICE  = SVR_PREFIX + "iot.panel.design.service"  //面板生成服务
	IOT_APP_API_SERVICE       = SVR_PREFIX + "iot.app.api.service"
	IOT_CLOUD_API_SERVICE     = SVR_PREFIX + "iot.cloud.api.service"
	IOT_DEMO_SERVICE          = SVR_PREFIX + "iot.demo.service"
	IOT_DEMO_API_SERVICE      = SVR_PREFIX + "iot.demo.api.service"
)
