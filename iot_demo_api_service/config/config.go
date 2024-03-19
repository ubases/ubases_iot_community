package config

import (
	"cloud_platform/iot_common/iotutil"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"go-micro.dev/v4/util/log"
)

type Settings struct {
	Service ServiceConfig `yaml:"service"`         //服务配置
	Redis   RedisConfig   `yaml:"redis,omitempty"` //redis配置
	Zipkin  Zipkin        `yaml:"zipkin"`          //ZipKin配置
	Etcd    EtcdCfg       `yaml:"etcd"`            //Etcd配置
}

type Zipkin struct {
	Url string `yaml:"url"`
}

type RedisConfig struct {
	Cluster      bool     `yaml:"Cluster"`
	Addrs        []string `yaml:"Addrs"`
	Username     string   `yaml:"Username,omitempty"`
	Password     string   `yaml:"Password"`
	Database     int      `yaml:"Database"`
	MinIdleConns int      `yaml:"MinIdleConns"` // 最小空闲连接
	IdleTimeout  int      `yaml:"IdleTimeout"`  // 空闲时间
	PoolSize     int      `yaml:"PoolSize"`     // 连接池大小
	MaxConnAge   int      `yaml:"MaxConnAge"`   // 连接最大可用时间
}

type ServiceConfig struct {
	HttpAddr       string `yaml:"httpAddr,omitempty"`
	Httpqps        int    `yaml:"httpqps,omitempty"`
	IPLimitRequest int    `yaml:"IPLimitRequest,omitempty"`
	ReadTimeout    int    `yaml:"readTimeout,omitempty"`
	WriteTimeout   int    `yaml:"writeTimeout,omitempty"`
	Logfile        string `yaml:"logfile"`
	Loglevel       string `yaml:"loglevel"`
}

type EtcdCfg struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

var (
	Global = new(Settings)
)

func Init() error {
	var (
		err        error
		configFile string
		env        string
	)
	if err = gotenv.Load("./conf/.env"); err != nil {
		env = "local"
		log.Error(err)
	}
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	configFile = "./conf/" + env + "/iot_demo_api_service.yml"

	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("viper.ReadInConfig: ", configFile, err)
		return err
	}
	if err = viper.Unmarshal(Global); err != nil {
		log.Error("viper.Unmarshal", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file: ", e.Name, " Op: ", e.Op)
		if err = viper.Unmarshal(Global); err != nil {
			log.Error(err)
		}
	})
	log.Info("setting init success !" + iotutil.ToString(Global))
	return err
}
