package services

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"fmt"
)

var (
	ZhDefaultHomeName = "我的家庭"
	EnDefaultHomeName = "My family"
)

// 家庭语言
func HomeLanguage(lang, homeName interface{}) string {
	name := iotutil.ToString(homeName)
	if name == ZhDefaultHomeName && lang == iotconst.LANG_ENGLISH {
		return EnDefaultHomeName
	} else if name == EnDefaultHomeName && lang == iotconst.LANG_CHINA {
		return ZhDefaultHomeName
	}
	return name
}

// 获取APP默认家庭名称（通过缓存获取，待定，考虑到调用会比较频繁，以下方式待定）
func getAppDefaultHomeName(lang string) string {
	if lang == "" {
		return ZhDefaultHomeName
	}
	key := fmt.Sprintf("%s_%s", iotconst.HKEY_LANGUAGE_DATA_PREFIX, iotconst.Dict_app_default)
	datas := iotredis.GetClient().HGetAll(context.Background(), key)
	if datas.Err() != nil {
		return ZhDefaultHomeName
	}
	langKey := fmt.Sprintf("%s_default_home_name", iotconst.HKEY_LANGUAGE_DATA_PREFIX)
	name := datas.Val()[langKey]
	if name == "" {
		return ZhDefaultHomeName
	}
	return name
}
