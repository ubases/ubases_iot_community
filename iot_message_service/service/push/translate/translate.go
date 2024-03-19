package translate

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"context"
	"fmt"
)

// 执行翻译（直接指定lang）
func Get(lang, inputTitle, inputContent string) (outTitle, outContent string) {
	if lang == "" {
		lang = "zh"
	}
	outTitle, outContent = GetAndDefault(lang, inputTitle, inputTitle, inputContent)
	return
}

// 转换中文翻译
func GetAndDefault(language, code, inputTitle, inputContent string) (title, content string) {
	title, content = inputTitle, inputContent
	langKey := fmt.Sprintf("%s%s", iotconst.HKEY_LANGUAGE_DATA_PREFIX, iotconst.LANG_MESSAGE_TEMPLATE)
	titleKey := fmt.Sprintf("%s_%s_tplSubject", language, code)
	contentKey := fmt.Sprintf("%s_%s_tplContent", language, code)
	strCmd := iotredis.GetClient().HMGet(context.Background(), langKey, titleKey, contentKey)
	arr := strCmd.Val()
	if len(arr) > 0 {
		title = iotutil.ToString(arr[0])
	}
	if len(arr) > 1 {
		content = iotutil.ToString(arr[1])
	}
	return
}
