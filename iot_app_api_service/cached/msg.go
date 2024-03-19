package cached

import (
	"cloud_platform/iot_common/ioterrs"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-micro.dev/v4/util/log"
)

var (
	configFile = "./conf/app_api_msg.ini"
	langs      = []string{"zh", "en"}
)

func InitMsg() error {
	var (
		err error
	)
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("viper.ReadInConfig: ", configFile, err)
		return err
	}
	for i := range langs {
		if err := saveCodeMsg(langs[i], viper.GetStringMap(langs[i])); err != nil {
			return err
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file: ", e.Name, " Op: ", e.Op)
		for i := range langs {
			if err := saveCodeMsg(langs[i], viper.GetStringMap(langs[i])); err != nil {
				log.Error("save code msg to redis: ", err)
			}
		}
	})
	log.Info("setting code msg to redis success !")
	return err
}

func saveCodeMsg(lang string, msg map[string]interface{}) error {
	return RedisStore.HMSet(ioterrs.GetCodeMsgKey(lang), msg)
}
