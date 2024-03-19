package rpc

import (
	"cloud_platform/iot_system_service/config"
	"sync"
	"time"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4/registry"
)

var reg registry.Registry
var registryOnce sync.Once

func getRegistry() registry.Registry {
	registryOnce.Do(func() {
		var addrs []string
		var name, password string
		etcdCfg := config.Global.Etcd
		if len(etcdCfg.Address) == 0 {
			addrs = []string{"127.0.0.1:2379"}
			name = ""
			password = ""
		} else {
			addrs = etcdCfg.Address
			name = etcdCfg.Username
			password = etcdCfg.Password
		}
		reg = etcd.NewRegistry(
			registry.Addrs(addrs...),
			registry.Timeout(5*time.Second),
			etcd.Auth(name, password),
		)
	})
	return reg
}
