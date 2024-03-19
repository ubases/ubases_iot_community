package iotmodel

import (
	"errors"
	"fmt"

	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type DBConfig struct {
	DBType     string //mysql，clickhouse
	ConnectStr string
	LogLevel   string
}

func GetDB(dbs ...*gorm.DB) (res *gorm.DB) {
	if len(dbs) > 0 {
		return dbs[0]
	}
	return db
}

func SetDB(database *gorm.DB) {
	db = database
}

func InitDB(cnf DBConfig) error {
	var err error
	switch cnf.DBType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(cnf.ConnectStr), &gorm.Config{Logger: getLogger(cnf), PrepareStmt: false})
	case "clickhouse":
		db, err = gorm.Open(clickhouse.Open(cnf.ConnectStr), &gorm.Config{Logger: getLogger(cnf)})
	default:
		panic(errors.New(fmt.Sprintf("不支持的数据库:%s", cnf.DBType)))
	}
	return err
}

func InitDBEx(cnf DBConfig) (*gorm.DB, error) {
	var err error
	var retDB *gorm.DB
	switch cnf.DBType {
	case "mysql":
		retDB, err = gorm.Open(mysql.Open(cnf.ConnectStr), &gorm.Config{Logger: getLogger(cnf), PrepareStmt: false})
	case "clickhouse":
		retDB, err = gorm.Open(clickhouse.Open(cnf.ConnectStr), &gorm.Config{Logger: getLogger(cnf)})
	default:
		panic(errors.New(fmt.Sprintf("不支持的数据库:%s", cnf.DBType)))
	}
	return retDB, err
}

func getLogger(cnf DBConfig) logger.Interface {
	var lg logger.Interface
	if cnf.LogLevel == "debug" {
		lg = New(logger.Info)
	} else {
		lg = New(logger.Warn)
	}
	return lg
}
