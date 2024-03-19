package main

import (
	"flag"
	"fmt"
	"strconv"

	"go-micro.dev/v4/registry"
)

func main() {
	var addr string
	var user string
	var password string
	var serviceName string
	flag.StringVar(&addr, "addr", "127.0.0.1:2379", "etcd地址，格式类似127.0.0.1:2379")
	flag.StringVar(&user, "user", "", "etcd用户名")
	flag.StringVar(&password, "password", "", "etcd用户密码")
	flag.StringVar(&serviceName, "name", "", "要查询的服务名称")
	flag.Parse()

	var svrlist []*registry.Service
	var err error
	reg := GetRegistry(user, password, addr)
	if serviceName == "" {
		svrlist, err = reg.ListServices()
	} else {
		svrlist, err = reg.GetService(serviceName)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mapAddress := make(map[string]string)
	for i, v := range svrlist {
		for _, node := range v.Nodes {
			mapAddress[v.Name+"\t"+strconv.Itoa(i)] = node.Address
		}
	}
	for k, v := range mapAddress {
		fmt.Println(k, "\t\t", v)
	}
}
