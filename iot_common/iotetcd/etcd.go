package iotetcd

import (
	"errors"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/etcd/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/encoder"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
)

type EtcdConfig struct {
	Conf       config.Config
	Enc        encoder.Encoder
	Addrs      []string
	Username   string
	Password   string
	CommonConf string
	CustomConf string
}

var _etcdConf *EtcdConfig

func NewConfig(addrs []string, username, password, commonConf, customConf string) (*EtcdConfig, error) {
	enc := yaml.NewEncoder()
	// new config
	c, err := config.NewConfig(
		config.WithReader(
			json.NewReader( // json reader for internal config merge
				reader.WithEncoder(enc),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	_etcdConf = &EtcdConfig{
		Conf:       c,
		Enc:        enc,
		Addrs:      addrs,
		Username:   username,
		Password:   password,
		CommonConf: commonConf,
		CustomConf: customConf,
	}
	return _etcdConf, nil
}

func (ec *EtcdConfig) Load() error {
	common := etcd.NewSource(
		etcd.WithAddress(ec.Addrs...),
		etcd.WithPrefix(ec.CommonConf),
		etcd.StripPrefix(true),
		etcd.Auth(ec.Username, ec.Password),
		source.WithEncoder(ec.Enc),
	)

	custom := etcd.NewSource(
		etcd.WithAddress(ec.Addrs...),
		etcd.WithPrefix(ec.CustomConf),
		etcd.StripPrefix(true),
		etcd.Auth(ec.Username, ec.Password),
		source.WithEncoder(ec.Enc),
	)

	// load the config from a etcd source
	err := _etcdConf.Conf.Load(common, custom)
	if err != nil {
		return err
	}
	return nil
}

func GetEtcdConf() *EtcdConfig {
	if _etcdConf == nil {
		panic(errors.New("请先调用iotetcd.NewConfig创建Config"))
	}
	return _etcdConf
}
