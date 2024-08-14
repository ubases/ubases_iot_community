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
	IpService  IpService       `yaml:"ipService"`          //获取地理位置服务配置
	Jwt        JWTConfig       `yaml:"jwt"`                //JWT配置
	Database   DatabaseConfig  `yaml:"database,omitempty"` //数据库配置
	Redis      iotredis.Config `yaml:"redis,omitempty"`    //redis配置
	Nats       NATSConfig      `yaml:"NATS,omitempty"`     //Nats配置
	Zipkin     Zipkin          `yaml:"zipkin"`             //ZipKin配置
	Etcd       EtcdCfg         `yaml:"etcd"`               //Etcd配置
	Oss        Oss             `yaml:"oss"`                //OSS配置
	DefaultApp DefaultAppCfg   `yaml:"defaultapp"`         //OSS配置
	WebMQTT    WebMQTT         `yaml:"webmqtt"`
	WorkOrder  WorkOrderConfig `yaml:"workOrder"` //工单配置
}

type WebMQTT struct {
	Addr string `yaml:"addr"`
}

type DatabaseConfig struct {
	Database string `yaml:"database"`
	Driver   string `yaml:"driver"`
	Connstr  string `yaml:"connstr"`
}

type DefaultAppCfg struct {
	AppName     string `yaml:"appName"`
	AppKey      string `yaml:"appKey"`
	TenantId    string `yaml:"tenantId"`
	IosPkgName  string `yaml:"iosPkgName"`
	DownloadUrl string `yaml:"downloadUrl"`
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
	HttpAddr                  string `yaml:"httpAddr,omitempty"`
	Httpqps                   int    `yaml:"httpqps,omitempty"`
	IPLimitRequest            int    `yaml:"IPLimitRequest,omitempty"`
	ReadTimeout               int    `yaml:"readTimeout,omitempty"`
	WriteTimeout              int    `yaml:"writeTimeout,omitempty"`
	Logfile                   string `yaml:"logfile"`
	Loglevel                  string `yaml:"loglevel"`
	McuSdkDir                 string `yaml:"mcuSdkDir"`
	McuSdkScript              string `yaml:"mcuSdkScript"`
	PanelDir                  string `yaml:"panelDir"`
	PanelScript               string `yaml:"panelScript"`
	PanelBuildNotifyUrl       string `yaml:"panelBuildNotifyUrl"`
	ThirdDomain               string `yaml:"thirdDomain"`
	DefaultPassword           string `yaml:"defaultPassword"`
	PlatformCode              string `yaml:"platformCode"`
	ResponseRealError         int    `yaml:"responseRealError"`         //请求是否返回真实错误
	IsGenTestData             bool   `yaml:"isGenTestData"`             //注册是否生成测试数据
	GenTestDataProductId      int64  `yaml:"genTestDataProductId"`      //注册生成测试数据的产品类型
	GenTestDataControlId      int64  `yaml:"genTestDataControlId"`      //注册生成测试数据的产品面板Id
	VirtualDeviceNumber       int32  `yaml:"virtualDeviceNumber"`       //每个产品可授权虚拟设备数量
	OemAppPackageDomain       string `yaml:"oemAppPackageDomain"`       //构建二维码链接地址、APP构建中隐私政策、用户协议、关于我们的地址前缀
	OemAppPackageDomainSimple string `yaml:"oemAppPackageDomainSimple"` //构建二维码链接地址、APP构建中隐私政策、用户协议、关于我们的地址前缀
	TempDir                   string `yaml:"tempDir"`                   //临时保存文件的路径
	TestVerifTyCode           string `yaml:"testVerifyCode"`            //测试验证码
}

type WorkOrderConfig struct {
	WebsocketUrl string `yaml:"websocketUrl"`
}

type IpService struct {
	QueryUrl string `yaml:"queryUrl"`
	AppCode  string `yaml:"appCode"`
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
	UseOss string    `yaml:"useOss,omitempty"`
	Qiniu  OssConfig `yaml:"qiniu,omitempty"`
	Ali    OssConfig `yaml:"ali,omitempty"`
	S3     OssConfig `yaml:"s3,omitempty"`
}

type NATSConfig struct {
	Addrs []string `yaml:"addrs"`
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
	configFile = "./conf/" + env + "/iot_cloud_api_service.yml"

	viper.SetConfigFile(configFile)

	err = viper.ReadInConfig()
	if err != nil {
		log.Error(err)
		return err
	}
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
	conf, err := iotconfig.NewNacosConfig(&cnf.Nacos, fmt.Sprintf("iot_cloud_api_service-%s.yaml", cnf.Config.Env))
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
