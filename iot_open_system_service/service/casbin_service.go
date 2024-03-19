package service

import (
	"cloud_platform/iot_common/iotlogger"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

type CasbinService struct{}

var Casbin_Enforcer *casbin.Enforcer

// 初始化casbin
func InitCasbin(mysqlConnect string, casbinConfPath string) {

	a, _ := gormadapter.NewAdapter("mysql", mysqlConnect, "iot_open_system", "t_open_casbin_rule", true) // Your driver and data source.
	//a, _ := gormadapter.NewAdapter("mysql", mysqlConnect,true) // Your driver and data source.
	var err error
	Casbin_Enforcer, err = casbin.NewEnforcer(casbinConfPath, a)
	if err != nil {
		iotlogger.LogHelper.Error("init casbin error--->" + err.Error())
	}
}
