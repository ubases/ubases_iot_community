package cached

import (
	"cloud_platform/iot_common/iotredis"
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-micro.dev/v4/util/log"
)

var (
	translateConfigFile        = "./conf/panel_translate.ini"
	translateLangs             = []string{"zh", "en"}
	TranslateKeys       string = "app_panel_def_translate_%s"
)

func GetTranslateKeys(lang string) string {
	if lang == "" {
		lang = "zh"
	}
	return fmt.Sprintf(TranslateKeys, lang)
}

// 初始化APP面板翻译模板
func InitTranslate() error {
	var (
		err error
	)
	viper.SetConfigFile(translateConfigFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("viper.ReadInConfig: ", translateConfigFile, err)
		return err
	}
	for i := range translateLangs {
		if err := saveLangTranslate(translateLangs[i], viper.GetStringMap(translateLangs[i])); err != nil {
			return err
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file: ", e.Name, " Op: ", e.Op)
		for i := range translateLangs {
			if err := saveLangTranslate(translateLangs[i], viper.GetStringMap(translateLangs[i])); err != nil {
				log.Error("save code msg to redis: ", err)
			}
		}
	})
	log.Info("setting app panel default transalte redis success !")
	return err
}

func saveLangTranslate(lang string, msg map[string]interface{}) error {
	return iotredis.GetClient().HMSet(context.Background(), GetTranslateKeys(lang), msg).Err()
}

func GetLangTranslate(lang string) (map[string]string, error) {
	res := iotredis.GetClient().HGetAll(context.Background(), GetTranslateKeys(lang))
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res.Val(), nil
}
