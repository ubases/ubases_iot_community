package main

import (
	"time"

	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4/registry"
)

func GetRegistry(Username string, Password string, Address ...string) registry.Registry {
	var addrs []string
	var name, password string
	if len(Address) == 0 {
		addrs = []string{"127.0.0.1:2379"}
		name = ""
		password = ""
	} else {
		addrs = Address
		name = Username
		password = Password
	}
	reg := etcd.NewRegistry(
		registry.Addrs(addrs...),
		registry.Timeout(5*time.Second),
		etcd.Auth(name, password),
	)
	return reg
}
