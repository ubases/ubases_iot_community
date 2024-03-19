package config

import (
	"os"

	"gorm.io/gorm"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"go-micro.dev/v4/util/log"
)

type Settings struct {
	Service  ServiceConfig    `yaml:"service"`         //服务配置
	Database []DatabaseConfig `yaml:"database"`        //数据库配置
	Redis    RedisConfig      `yaml:"redis,omitempty"` //redis配
	XxlJob   XxlJobConfig     `yaml:"xxlJob"`
	Zipkin   Zipkin           `yaml:"zipkin"` //ZipKin配置
	Etcd     EtcdCfg          `yaml:"etcd"`   //Etcd配置
}
type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
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
	GrpcAddr string `yaml:"grpcAddr,omitempty"`
	Grpcqps  int    `yaml:"grpcqps,omitempty"`
	Logfile  string `yaml:"logfile"`
	Loglevel string `yaml:"loglevel"`
}

type XxlJobConfig struct {
	Enable       bool   `yaml:"enable"`
	ServerAddr   string `yaml:"serverAddr"`
	AccessToken  string `yaml:"accessToken"`
	ExecutorIp   string `yaml:"executorIp"`
	ExecutorPort string `yaml:"executorPort"`
	RegistryKey  string `yaml:"registryKey"`
}

type Zipkin struct {
	Url string `yaml:"url"`
}

type EtcdCfg struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

var (
	Global = new(Settings)
)

var DBMap = make(map[string]*gorm.DB)

func Init() error {
	var (
		err        error
		configFile string
		env        string
	)
	if err = gotenv.Load("./conf/.env"); err != nil {
		env = "local"
		log.Error(err)
		return err
	}
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	configFile = "./conf/" + env + "/iot_statistics_service.yml"

	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err = viper.Unmarshal(Global); err != nil {
		log.Error(err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file: ", e.Name, " Op: ", e.Op)
		if err = viper.Unmarshal(Global); err != nil {
			log.Error(err)
		}
	})
	log.Info("setting init success !")
	return err
}
