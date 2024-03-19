package config

import (
	"strconv"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-micro.dev/v4/util/log"
)

var (
	configFile  = "./conf/device_fault.ini"
	FaultConfig DevFaultConfig
)

type DevFaultConfig struct {
	mu   sync.RWMutex
	data map[int32]string
}

func (o *DevFaultConfig) Init() error {
	o.data = make(map[int32]string)
	var err error
	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("viper.ReadInConfig: ", configFile, err)
		return err
	}
	mapConfig := viper.GetStringMapString("default")
	for k, v := range mapConfig {
		nkey, err := strconv.Atoi(k)
		if err == nil {
			o.data[int32(nkey)] = v
		}
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file: ", e.Name, " Op: ", e.Op)
		mapConfig := viper.GetStringMapString("default")
		data := make(map[int32]string)
		for k, v := range mapConfig {
			nkey, err := strconv.Atoi(k)
			if err == nil {
				data[int32(nkey)] = v
			}
		}
		if len(data) > 0 {
			o.mu.Lock()
			defer o.mu.Unlock()
			o.data = data
		}
	})
	log.Info("load devfault.ini success !")
	return err
}

func (o *DevFaultConfig) GetValue(key int32) string {
	o.mu.RLock()
	o.mu.RUnlock()
	if o.data != nil {
		return o.data[key]
	}
	return ""
}
