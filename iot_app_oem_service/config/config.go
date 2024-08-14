package config

import (
	"cloud_platform/iot_common/iotconfig"
	"cloud_platform/iot_common/iotredis"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"go-micro.dev/v4/config/reader"

	"gopkg.in/yaml.v3"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"go-micro.dev/v4/util/log"
)

type Settings struct {
	Service   ServiceConfig   `yaml:"service"`            //服务配置
	Database  DatabaseConfig  `yaml:"database,omitempty"` //数据库配置
	Redis     iotredis.Config `yaml:"redis,omitempty"`    //redis配置
	Nats      NATSConfig      `yaml:"NATS,omitempty"`     //Nats配置
	Zipkin    Zipkin          `yaml:"zipkin"`             //ZipKin配置
	Etcd      EtcdCfg         `yaml:"etcd"`               //Etcd配置
	Oss       Oss             `yaml:"oss"`                //OSS配置
	AppBuild  AppBuildConfig  `yaml:"appBuild"`           //AppBuild配置
	GlobalApp GlobalApp       `yaml:"globalApp"`          //构建APP全局配置

}

type GlobalApp struct {
	Color GlobalAppColor `yaml:"color"`
}

type GlobalAppColor struct {
	LightColour map[string]interface{} `yaml:"lightColour"`
	Dark        map[string]interface{} `yaml:"dark"`
}

type ColorSet struct {
	MainColor            string `json:"mainColor" yaml:"mainColor"`                       //重要文字
	PlainColor           string `json:"plainColor" yaml:"plainColor"`                     //普通文字
	SecondaryColor       string `json:"secondaryColor" yaml:"secondaryColor"`             //次要信息
	CardBgColor          string `json:"cardBgColor" yaml:"cardBgColor"`                   //卡片底色
	DialogBgColor        string `json:"dialogBgColor" yaml:"dialogBgColor"`               //弹窗底色
	DialogBtnBgColor     string `json:"dialogBtnBgColor" yaml:"dialogBtnBgColor"`         //弹窗操作按钮底色
	DividerColor         string `json:"dividerColor" yaml:"dividerColor"`                 //分割线
	ArrowRightColor      string `json:"arrowRightColor" yaml:"arrowRightColor"`           //列表右侧箭头
	ArrowLeftColor       string `json:"arrowLeftColor" yaml:"arrowLeftColor"`             //导航栏左侧返回箭头
	LineColor            string `json:"lineColor" yaml:"lineColor"`                       //
	RingColor            string `json:"ringColor" yaml:"ringColor"`                       //
	SolidColor           string `json:"solidColor" yaml:"solidColor"`                     //
	NoSelectedBgColor    string `json:"noSelectedBgColor" yaml:"noSelectedBgColor"`       //
	CircleColor          string `json:"circleColor" yaml:"circleColor"`                   //新建家庭未选择房间圆圈颜色
	ShareColor           string `json:"shareColor" yaml:"shareColor"`                     //图片
	GuanjiColor          string `json:"guanjiColor" yaml:"guanjiColor"`                   //首页关机颜色
	GuanjiBgColor        string `json:"guanjiBgColor" yaml:"guanjiBgColor"`               //首页关机背景颜色
	DashedColor          string `json:"dashedColor" yaml:"dashedColor"`                   //
	OffLineCardBgColor   string `json:"offLineCardBgColor" yaml:"offLineCardBgColor"`     //
	OnLineCardBgColor    string `json:"onLineCardBgColor" yaml:"onLineCardBgColor"`       //
	CardShadow           string `json:"cardShadow" yaml:"cardShadow"`                     //卡片阴影
	BottomShadow         string `json:"bottomShadow" yaml:"bottomShadow"`                 //底部阴影
	NoSelectedBntBgColor string `json:"noSelectedBntBgColor" yaml:"noSelectedBntBgColor"` //群组设备，虚拟体验，未选中时背景颜色
	DialogSpecialColor   string `json:"dialogSpecialColor" yaml:"dialogSpecialColor"`     //特殊颜色，弹窗深色是计算
	SpecialColor         string `json:"specialColor" yaml:"specialColor"`                 // 特殊颜色
}

type AppBuildConfig struct {
	AssociatedDomains string `json:"associatedDomains" yaml:"associatedDomains"`
	RegionServerUrl   string `json:"regionServerUrl" yaml:"regionServerUrl"` //APP通过获取区域服务器列表服务地址
	BuildNotify       string `json:"buildNotify" yaml:"buildNotify"`         //APP构建回调通知地址，将地址推送给构建服务，构建服务通过该地址上报构建状态
	BuildMode         int    `json:"buildMode" yaml:"buildMode"`             //构建生成的方式（=1 自有打包机生成， 否则：为官方打包机）
	BuildKey          string `json:"buildKey" yaml:"buildKey"`               //APP构建Key，公版平台云打包服务提供
	BuildServerUrl    string `json:"buildServerUrl" yaml:"buildServerUrl"`   //APP构建云打包平台地址，公版平台云打包服务提供
}

type Oss struct {
	UseOss string    `yaml:"useOss,omitempty"`
	Qiniu  OssConfig `yaml:"qiniu,omitempty"`
	Ali    OssConfig `yaml:"ali,omitempty"`
	S3     OssConfig `yaml:"s3,omitempty"`
}

type OssConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Region          string
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
	configFile = "./conf/" + env + "/iot_app_oem_service.yml"

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
	globalFile := "./conf/global.yml"
	err = readGlobalYml(globalFile)
	if err != nil {
		return err
	}
	log.Info("setting init success !")
	return err
}

func readGlobalYml(file string) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var global GlobalAppColor
	err = yaml.Unmarshal(buf, &global)
	if err != nil {
		return err
	}
	Global.GlobalApp.Color = global
	return nil
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
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_app_oem_service-%s.yaml", cnf.Config.Env))
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
