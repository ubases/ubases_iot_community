package config

import (
	"cloud_platform/iot_common/iotconfig"
	"errors"
	"fmt"
	"os"

	"go-micro.dev/v4/config/reader"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"go-micro.dev/v4/util/log"
)

type Settings struct {
	Service  ServiceConfig  `yaml:"service"`            //服务配置
	Database DatabaseConfig `yaml:"database,omitempty"` //数据库配置
	Nats     NATSConfig     `yaml:"NATS,omitempty"`     //Nats配置
	Zipkin   Zipkin         `yaml:"zipkin"`             //ZipKin配置
	Etcd     EtcdCfg        `yaml:"etcd"`               //Etcd配置
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
	Days     int    `yaml:"days"`
}

type NATSConfig struct {
	Addrs []string `yaml:"addrs"`
}

type Zipkin struct {
	Url string `yaml:"url"`
}

type ServiceConfig struct {
	GrpcAddr string `yaml:"grpcAddr,omitempty"`
	Grpcqps  int    `yaml:"grpcqps,omitempty"`
	HttpAddr string `yaml:"httpAddr,omitempty"`
	Httpqps  int    `yaml:"httpqps,omitempty"`
	Logfile  string `yaml:"logfile"`
	Loglevel string `yaml:"loglevel"`
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
	configFile = "./conf/" + env + "/iot_log_service.yml"

	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err)
		return err
	}
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

func InitTest(path ...string) error {
	pathStr := "../"
	if len(path) > 0 {
		pathStr = path[0]
	}
	var (
		err        error
		configFile string
		env        string
	)
	if err = gotenv.Load(pathStr + "conf/.env"); err != nil {
		env = "local"
		log.Error(err)
	}
	env = os.Getenv("ENVIRONMENT")
	if env == "" {
		configFile = "./conf/local/iot_log_service.yml"
	} else {
		configFile = "./conf/" + env + "/iot_log_service.yml"
	}
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

func Init2() error {
	cnf, err := iotconfig.LoadIotConfig()
	if err != nil {
		return Init()
	}
	if cnf.Config.Location == iotconfig.Location_local {
		return Init()
	}
	if cnf.Config.Location != iotconfig.Location_nacos {
		return errors.New("location unsupported ")
	}
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_log_service-%s.yaml", cnf.Config.Env))
	if err != nil {
		return err
	}
	if err := conf.Scan(Global); err != nil {
		return err
	}
	//开启监听
	iotconfig.Watch(conf, WatchCB)
	return nil
}

func WatchCB(v reader.Value, err error) {
	if err != nil {
		log.Error(err)
		return
	}
	if err = v.Scan(Global); err != nil {
		log.Error(err)
	}
}
