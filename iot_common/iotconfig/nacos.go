package iotconfig

import (
	"time"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/nacos/v4"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
)

func NewNacosConfig(cnf *NacosConfig, dataId string) (config.Config, error) {
	clientConfig := constant.ClientConfig{
		NamespaceId: cnf.NamespaceId, NotLoadCacheAtStart: true,
		Username: cnf.Username, Password: cnf.Password, LogLevel: "debug"}
	nacosSource := nacos.NewSource(
		nacos.WithAddress(cnf.Addrs), nacos.WithClientConfig(clientConfig), nacos.WithDataId(dataId),
		nacos.WithGroup(cnf.Group), source.WithEncoder(yaml.NewEncoder()))
	var conf config.Config
	var err error
	for retry := 3; retry > 0; retry-- {
		conf, err = config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(yaml.NewEncoder()))), config.WithSource(nacosSource))
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}
		config.DefaultConfig = conf
	}

	return conf, err
}

func Watch(conf config.Config, fun func(reader.Value, error)) {
	go func() {
		watcher, err := conf.Watch()
		if err != nil {
			fun(nil, err)
			return
		}
		for {
			v, err := watcher.Next()
			if err != nil {
				fun(nil, err)
			} else {
				fun(v, nil)
			}
		}
	}()
}
