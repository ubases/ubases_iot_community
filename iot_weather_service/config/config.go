package config

import (
	"cloud_platform/iot_common/iotconfig"
	"cloud_platform/iot_common/iotredis"
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
	Service    ServiceConfig `yaml:"service"`         //服务配置
	Geo        GeoConfig     `yaml:"geo"`             //获取地理位置服务配置(geoip)
	IpService  IpService     `yaml:"ipService"`       //获取地理位置服务配置（阿里云）
	Redis     iotredis.Config `yaml:"redis,omitempty"` //redis配置
	//Nats       NATSConfig      `yaml:"NATS,omitempty"`  //Nats配置
	Zipkin     Zipkin        `yaml:"zipkin"`          //ZipKin配置
	Etcd       EtcdCfg       `yaml:"etcd"`            //Etcd配置
	Weather    Weather       `yaml:"weather"`         //天气1配置 openweathermap
	Yiketianqi Yiketianqi    `yaml:"yiketianqi"`      //天气2配置 易客云,www.tianqiapi.com
}

type GeoConfig struct {
	DbPath     string `yaml:"dbpath"`
	LicenseKey string `yaml:"licensekey"`
	Interval   int    `yaml:"interval"`
}

type IpService struct {
	QueryUrl string `yaml:"queryUrl"`
	AppCode  string `yaml:"appCode"`
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
}

type NATSConfig struct {
	Addrs []string `yaml:"addrs"`
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
	GrpcAddr string `yaml:"grpcAddr,omitempty"`
	Grpcqps  int    `yaml:"grpcqps,omitempty"`
	HttpAddr string `yaml:"httpAddr,omitempty"`
	Httpqps  int    `yaml:"httpqps,omitempty"`
	Logfile  string `yaml:"logfile"`
	Loglevel string `yaml:"loglevel"`
}

type JWTConfig struct {
	SigningKey      string `yaml:"SigningKey"`
	AccessTokenTTL  int64  `yaml:"AccessTokenTTL"`
	RefreshTokenTTL int64  `yaml:"RefreshTokenTTL"`
}

type SMSConfig struct {
	Provider  string
	AccessId  string
	AccessKey string
	Sign      string
	Other     []string
}

type Weather struct {
	Provider   string `yaml:"provider"`
	ApiKey     string `yaml:"apikey"` //658c3bed42f60a847640dfc7a5fa95c4
	AqicnToken string `yaml:"aqicnToken"`
	Qps        int    `yaml:"qps"`
	Interval   int    `yaml:"interval"`
}

type Yiketianqi struct {
	Enabled   bool   `yaml:"enabled"`
	Appid     int    `yaml:"appid"`
	Appsecret string `yaml:"appsecret"`
}

type EtcdCfg struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

type SMTPConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Username       string `yaml:"userName"`
	Password       string `yaml:"password"`
	ConnectTimeout int64  `yaml:"connectTimeout,omitempty"`
	SendTimeout    int64  `yaml:"sendTimeout,omitempty"`
	Helo           string `yaml:"helo,omitempty"`
	KeepAlive      bool   `yaml:"keepAlive,omitempty"`
}

type NoticeConfig struct {
	DingTalk DingTalkCfg `json:"dingtalk"`
}

type DingTalkCfg struct {
	Keyword string `yaml:"keyword"`
	Webhook string `yaml:"webhook"`
	Secert  string `yaml:"secert"`
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
	configFile = "./conf/" + env + "/iot_weather_service.yml"

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
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_weather_service-%s.yaml", cnf.Config.Env))
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
