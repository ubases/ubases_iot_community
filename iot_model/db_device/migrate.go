package db_device

import (
	"cloud_platform/iot_model/db_device/model"
	"cloud_platform/iot_common/iotlogger"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	//数据库升级，现存表更新
	if err := db.AutoMigrate(
		model.TIotDeviceHome{},
		model.TIotDeviceInfo{},
		model.TIotDeviceLog{},
		model.TIotDeviceTriad{},
	); err != nil {
		iotlogger.LogHelper.Errorf("iot_device数据库表结构更新失败:%s", err.Error())
	}
	//新表
	if !db.Migrator().HasTable(model.TIotDeviceHome{}) {
		err := db.Migrator().CreateTable(model.TIotDeviceHome{})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if !db.Migrator().HasTable(model.TIotDeviceInfo{}) {
		err := db.Migrator().CreateTable(model.TIotDeviceInfo{})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if !db.Migrator().HasTable(model.TIotDeviceLog{}) {
		err := db.Migrator().CreateTable(model.TIotDeviceLog{})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if !db.Migrator().HasTable(model.TIotDeviceTriad{}) {
		err := db.Migrator().CreateTable(model.TIotDeviceTriad{})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
