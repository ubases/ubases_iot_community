package iotconfig

import "github.com/spf13/viper"

const Location_local = "local"
const Location_nacos = "nacos"

type IotConfig struct {
	Config Config      `yaml:"config"`
	Nacos  NacosConfig `yaml:"nacos"`
}

type NacosConfig struct {
	Addrs       []string `yaml:"addrs"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	NamespaceId string   `yaml:"namespaceId"`
	Group       string   `yaml:"group"`
}

type Config struct {
	Env      string `yaml:"env"`
	Location string `yaml:"location"`
}

func LoadIotConfig() (*IotConfig, error) {
	configFile := "./conf/config.yml"
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var conf IotConfig
	if err = viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
