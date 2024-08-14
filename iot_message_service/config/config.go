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
	Service         ServiceConfig          `yaml:"service"`            //服务配置
	Database        DatabaseConfig         `yaml:"database,omitempty"` //数据库配置
	Gorush          GorushCfg              `yaml:"gorush"`
	Jpush           JpushCfg               `yaml:"jpush"`
	SMS             SMSConfig              `yaml:"SMS"`  //SMS配置
	SMTP            SMTPConfig             `yaml:"smtp"` //SMTP邮箱配置
	Notice          NoticeConfig           `yaml:"notice"`
	Redis           iotredis.Config        `yaml:"redis,omitempty"` //redis配置
	Nats            NATSConfig             `yaml:"NATS,omitempty"`  //Nats配置
	Zipkin          Zipkin                 `yaml:"zipkin"`          //ZipKin配置
	Etcd            EtcdCfg                `yaml:"etcd"`            //Etcd配置
	ThirdPartyLogin ThirdPartyLoginTypeCfg `yaml:"thirdPartyLogin"`
}

type ThirdPartyLoginTypeCfg struct {
	MiniProgram ThirdPartyLoginConfig `yaml:"miniProgram"`
}

type ThirdPartyLoginConfig struct {
	AppId      string `yaml:"appId"`
	AppSecret  string `yaml:"appSecret"`
	TemplateId string `yaml:"templateId"`
	Page       string `yaml:"page"`
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
}

type SMSConfig struct {
	Provider  string
	AccessId  string
	AccessKey string
	Sign      map[string]string
	Other     []string
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
	Exchange       bool   `yaml:"exchange,omitempty"`
	AuthType       int    `yaml:"authType"` //0:AuthPlain; 1:AuthLogin; 3:AuthCRAMMD5; 4:AuthNone
	Ssl            int    `yaml:"ssl"`
	From           string `yaml:"from"` //发件人邮箱地址
}

type NoticeConfig struct {
	DingTalk DingTalkCfg `json:"dingtalk"`
}

type DingTalkCfg struct {
	Keyword string `yaml:"keyword"`
	Webhook string `yaml:"webhook"`
	Secert  string `yaml:"secert"`
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

type EtcdCfg struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username,omitempty"`
	Password string   `yaml:"password,omitempty"`
}

type JpushCfg struct {
	AppKey         string `yaml:"appKey"`
	Secret         string `yaml:"secret"`
	AndroidPkgName string `yaml:"androidPkgName"` //android包名
	IosPkgName     string `yaml:"iosPkgName"`     //ios包名
	ApnsProduction bool   `yaml:"apnsProduction"`

	Jpush  map[string]interface{} `json:"jpush"`
	Apns   map[string]interface{} `json:"apns"`
	Fcm    map[string]interface{} `json:"fcm"`
	Huawei map[string]interface{} `json:"huawei"`
	Xiaomi map[string]interface{} `json:"xiaomi"`
	Vivo   map[string]interface{} `json:"vivo"`
	Oppo   map[string]interface{} `json:"oppo"`
	Honor  map[string]interface{} `json:"honor"`
}

type GorushCfg struct {
	Url string `yaml:"url"`
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
	configFile = "./conf/" + env + "/iot_message_service.yml"

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
		configFile = "./conf/local/iot_message_service.yml"
	} else {
		configFile = "./conf/" + env + "/iot_message_service.yml"
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
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_message_service-%s.yaml", cnf.Config.Env))
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
