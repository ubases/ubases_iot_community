package service

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInitCasbin(t *testing.T) {
	type args struct {
		mysqlConnect   string
		casbinConfPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCasbin(tt.args.mysqlConnect, tt.args.casbinConfPath)
		})
	}
}

func TestCansbin(t *testing.T) {
	fmt.Println("test start.............")
	// if err := config.Init(); err != nil {
	// 	fmt.Println("加载配置文件发生错误:", err)
	// 	return
	// }
	InitCasbin("proxy:proxy649@aixingyun.com@tcp(120.77.96.118:3306)/iot_system?charset=utf8mb4&parseTime=True&loc=Local", "D:\\aa-works\\cloud_platform\\conf\\rbac_model.conf")

	//根据角色id 找用户
	res := Casbin_Enforcer.GetFilteredGroupingPolicy(1, "4947433231210151936")

	for k, v := range res {

		fmt.Println(k)
		fmt.Println(v)
	}

}
