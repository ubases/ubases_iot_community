package iotconst

const (
	//HKEY_DEV_STATUS_PREFIX string = "dev_stat_"
	HKEY_DEV_DATA_PREFIX      string = "dev_data_"        //设备实时数据,key规则:dev_data_{deviceKey}
	FIELD_ONLINE              string = "onlineStatus"     //设备实时数据之在线状态字段，字符串offline\online
	FIELD_ONLINETIME          string = "onlineTime"       //设备实时数据之在线状态时间字段,整数时间戳
	FIELD_TIME                string = "time"             //设备实时数据属性上报时间字段,整数时间戳
	FIELD_UPGRADE_STATE       string = "upgradeState"     //设备固件升级状态（结果状态）
	FIELD_UPGRADE_HAS         string = "hasOtaUpgrade"    //是否存在ota升级（创建ota的时候修改，升级完成ota的时候标记）
	FIELD_UPGRADE_RUNNING     string = "isUpgradeRunning" //是否升级进行中
	FIELD_UPGRADE_MODE        string = "otaUpgradeMode"   //升级方式 1: APP提醒升级, 2: APP强制升级, 3: APP检测升级
	FIELD_UPGRADE_FORCE_VER   string = "forceUpgradeVer"  //强制升级版本
	FIELD_UPGRADE_CODE        string = "upgradeCode"      //设备固件升级编码
	FIELD_UPGRADE_PROGRESS    string = "upgradeRrogress"  //设备固件升级进度
	FIELD_UPGRADE_PUB_ID      string = "upgradePubId"     //设备固件升级发布编号
	FIELD_UPGRADE_OTA_VER     string = "upgradeOtaVer"    //设备固件升级OTA版本
	FIELD_UPGRADE_TIME        string = "upgradeTime"      //设备固件升级时间
	FIELD_UPGRADE_TIMEOUT     string = "upgradeTimeout"   //设备固件升级时间超时
	FIELD_IS_PUSH_OFFLINE_MSG string = "isPushOfflineMsg" //是否推送offline消息
	FIELD_IS_PUSH_ONLINE_MSG  string = "isPushOnlineMsg"  //是否推送online消息
	FIELD_IS_AUTH_UPGRADE     string = "isAuthUpgrade"    //是否自动升级（0不自动升级 1自动升级）
	FIELD_IS_FW_VER           string = "fwVer"            //固件的当前版本
	FIELD_DEVICE_NAME         string = "deviceName"       //设备名称

	//智能场景缓存格式
	HKEY_INTELLIGENCE_DATA   string = "intelligence_data_" //智能场景实时数据,key规则:intelligence_data_{intelligenceId}
	HKEY_PRODUCT_DATA        string = "product_data_"      //产品缓存，key规则： product_data_{productId}
	HKEY_WEATHER_LIST        string = "weather_data_list"  //产品缓存，key规则： weather_data_list
	HKEY_WEATHER_DATA        string = "weather_data_"      //产品缓存，key规则： weather_data_{CITY}
	APP_INVITE_CODE          string = "app_invite_code_"   //加入家庭邀请码 ，key规则： app_invite_code_{code}
	HKEY_APPPUSH_DATA_PREFIX string = "app_push_data_"     //APPPush参数数据,key规则:app_push_data_{deviceKey}
	HKEY_GROUP_DATA          string = "group_data_"        //群组物模型缓存，key规则： group_data_{groupId}
	APP_DEVICE_SHARE_CODE    string = "app_share_code_"    //设备分享邀请码 ，key规则： app_invite_code_{code}
	HKEY_DFAULT_DATA = "app_default" //APP默认数据  { zh_default_home_name:"我的家庭",en_default_home_name:"my family" }
	HKEY_REGION_DATA string = "region_data_" //区域数据，key规则： region_data_{regionId}
)

var (
	FIELD_PREFIX_TLS  string = "tls_"  //物模型缓存Key前缀
	FIELD_PREFIX_DPID string = "dpid_" //群组物模型缓存Key前缀
)

var (
	//翻译
	HKEY_LANGUAGE_DATA_PREFIX string = "language_data_" //Language参数数据,key规则:language_data_{sourceTable}
)

var (
	RED_DOT_DATA                string = "red_dot_data_%d"                //用户红点数据
	CONTROL_PANEL_IS_UPDATE     string = "control_panel_is_update_%d"     //控制面板更新
	PRODUCT_TYPE_ID_DATA        string = "product_type_id_data_%s"        //产品分类id
	PRODUCT_TYPE_DATA           string = "product_type_data_%s"           //产品类型
	PRODUCT_TYPE_ALL_DATA       string = "product_type_data_all"          //产品类型
	APP_HOME_DETAIL_DATA        string = "app_home_detail_data_%v_%v"     //app家庭数据  _家庭Id_用户Id
	REGION_LIST_DATA            string = "region_list_data"               //区域服务列表数据
	APP_HOME_LIST_DATA          string = "app_home_list_data_%v"          //家庭列表数据
	APP_HOME_ROOM_LIST_DATA     string = "app_home_room_list_data_%v_%v"  //家庭房间列表数据 app_home_room_list_data_{lang}_{home}
	OPEN_USER_PROFILE_DATA      string = "open_user_profile_data_%s_%s"   //开放平台用户profile数据
	OPEN_PRODUCT_TREE_DATA      string = "open_product_tree_data"         //开放平台产品树数据
	OPEN_PRODUCT_FUNC_LIST_DATA string = "open_product_func_list_data_%s" //开放平台功能列表数据
	DICT_DATA                   string = "dict_data_%s"                   //字典数据
	DEVELOPER_TENANT_ID_DATA    string = "open_developer_data_%s"         //开发平台开发者数据
	SYSTEM_USER_DATA            string = "system_user_data_%s"            //西格用户数据
	PRODUCT_MATERIAL_DATA       string = "product_material_%s_%s_%s_%s"   //精油缓存，前缀_租户id_品牌_香型_语言
	APP_PRODUCT_PANEL_LANG      string = "app_product_panel_lang_%s"      //产品面板的翻译（app_product_panel_lang_{产品Key})
	APP_PRODUCT_PANEL_LANG_KEYS string = "app_product_panel_lang_keys_%v" //产品面板的翻译（app_product_panel_lang_keys_{面板Id})
	APP_CUSTOM_LANG             string = "app_lang_custom_%v"             //APP的翻译（app_lang_custom_{产品Key})
	APP_COMMON_LANG             string = "app_lang_common_%v"             //APP的翻译（app_lang_common_{产品Key})
	APP_HOME_INTELLIGENCE_DATA  string = "app_home_intelligence_data_%v"  //家庭场景数据 app_home_intelligence_data_{homeId}
	PANEL_RULE_SETTINGS_DATA    string = "panel_rule_settings_%v_%v"      //面板规则配置缓存 panel_rule_settings_{productKey}_{panelId}
)

// 语控缓存
var (
	VOICE_PRODUCT_SHILL_CACHED      string = "voice_product_skill_%v_%v"      //产品缓存 voice_product_skill_{语控类型，例如：xiaoai}_{productKey}
	VOICE_PRODUCT_LIST_SHILL_CACHED string = "voice_product_list_skill_%v_%v" //产品缓存 voice_product_list_skill_{语控类型，例如：xiaoai}_{productKey}
	VOICE_PRODUCT_DATA_CACHED       string = "voice_product_data_%v_%v"       //产品语音配置 voice_product_data_{语控类型，例如：xiaoai}_{productKey}
)

var (
	APPOPERATELIST string = "AppOperate" //缓存app操作
)

var (
	USERTOKENPREFIX string = "utokens:"      //用户tokens集合前缀，规则“前缀:userid”
	USERLASTLOGIN   string = "openLastLogin" //开放平台用户最近登录信息前缀,固定值
)

//var (
//	HKEY_REPORT_DATA_PUB_PREFIX    string = "report"   //设备数据发布到redis的前缀
//	HKEY_ONLINE_DATA_PUB_PREFIX    string = "online"   //设备数据发布到redis的前缀
//	HKEY_UPDATE_DATA_PUB_PREFIX    string = "update"   //设备数据发布到redis的前缀
//	HKEY_ACK_DATA_PUB_PREFIX       string = "ack"      //设备数据发布到redis的前缀
//	HKEY_QUERY_ACK_DATA_PUB_PREFIX string = "queryAck" //设备数据发布到redis的前缀
//)

//var (
//	HKEY_CACHED_CLEAR_PUB_PREFIX string = "clearCached" //OK 清理缓存
//)
