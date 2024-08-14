package config

import (
	"cloud_platform/iot_common/iotconfig"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
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
	Service    ServiceConfig   `yaml:"service"`            //服务配置
	Jwt        JWTConfig       `yaml:"jwt"`                //JWT配置
	Database   DatabaseConfig  `yaml:"database,omitempty"` //数据库配置
	Redis      iotredis.Config `yaml:"redis,omitempty"`    //redis配置
	Nats       NATSConfig      `yaml:"NATS,omitempty"`     //Nats配置
	Zipkin     Zipkin          `yaml:"zipkin"`             //ZipKin配置
	Etcd       EtcdCfg         `yaml:"etcd"`               //Etcd配置
	Oss        Oss             `yaml:"oss"`                //OSS配置
	AppMQTT    AppMQTT         `yaml:"appmqtt"`
	DeviceMQTT DeviceMQTT      `yaml:"DeviceMQTT"`
	IpService  IpService       `yaml:"ipService"` //获取地理位置服务配置
}
type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
}

type NATSConfig struct {
	Addrs []string `yaml:"addrs"`
}

type AppMQTT struct {
	Addrs []string `yaml:"addrs"`
}

type DeviceMQTT struct {
	Ip            string   `yaml:"ip"`
	Port          int32    `yaml:"port"`
	Nationcode    string   `yaml:"nationcode"`
	NoAclProducts []string `yaml:"noAclProducts"`
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
	HttpAddr             string `yaml:"httpAddr,omitempty"`
	Httpqps              int    `yaml:"httpqps,omitempty"`
	IPLimitRequest       int    `yaml:"IPLimitRequest,omitempty"`
	Logfile              string `yaml:"logfile"`
	Loglevel             string `yaml:"loglevel"`
	ThirdPartyRequirePwd bool   `yaml:"thirdPartyRequirePwd"` //第三方必须密码开关
	ResponseRealError    int    `yaml:"responseRealError"`
}

type JWTConfig struct {
	SigningKey      string `yaml:"SigningKey"`
	AccessTokenTTL  int64  `yaml:"AccessTokenTTL"`
	RefreshTokenTTL int64  `yaml:"RefreshTokenTTL"`
}

type EtcdCfg struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

type OssConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Region          string
}

type Oss struct {
	UseOss   string       `yaml:"useOss,omitempty"`
	ImgStyle *OssImgStyle `yaml:"imgStyle,omitempty"`
	Qiniu    OssConfig    `yaml:"qiniu,omitempty"`
	S3       OssConfig    `yaml:"s3,omitempty"`
	Ali      OssConfig    `yaml:"ali,omitempty"`
}

type OssImgStyle struct {
	ProductImg string `yaml:"productImg,omitempty"`
}

type IpService struct {
	QueryUrl string `yaml:"queryUrl"`
	AppCode  string `yaml:"appCode"`
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
	configFile = "./conf/" + env + "/iot_app_api_service.yml"

	viper.SetConfigFile(configFile)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err)
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
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_app_api_service-%s.yaml", cnf.Config.Env))
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
