package etl

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_statistics_service/config"
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"

	appModel "cloud_platform/iot_model/db_app/model"
	appOrm "cloud_platform/iot_model/db_app/orm"

	"github.com/go-redis/redis/v8"
)

type UcUserOperateETL struct {
	Cache []*appModel.TUcUserOperate
}

func (o *UcUserOperateETL) Handler(args ...interface{}) interface{} {
	for {
		list, err := iotredis.GetClient().BRPop(context.Background(), 5*time.Second, iotconst.APPOPERATELIST).Result()
		if err != nil {
			if err == redis.Nil {
				o.Save()
				time.Sleep(100 * time.Millisecond)
			} else {
				time.Sleep(60 * time.Second)
				iotlogger.LogHelper.Errorf("UcUserOperateHandler,error.%s", err.Error())
			}
		} else {
			if len(list) == 2 {
				o.PushCache(list[1])
			}
			if len(o.Cache) > 100 {
				o.Save()
			}
		}
	}
}
func (o *UcUserOperateETL) PushCache(str string) {
	var opt appModel.TUcUserOperate
	err := jsoniter.UnmarshalFromString(str, &opt)
	if err != nil {
		iotlogger.LogHelper.Errorf("PushCache,error.%s", err.Error())
		return
	}
	o.Cache = append(o.Cache, &opt)
}

func (o *UcUserOperateETL) Save() {
	if len(o.Cache) == 0 {
		return
	}
	appDB, ok := config.DBMap["iot_app"]
	t := appOrm.Use(appDB).TUcUserOperate
	if !ok {
		return
	}
	err := t.WithContext(context.Background()).CreateInBatches(o.Cache, len(o.Cache))
	if err != nil {
		iotlogger.LogHelper.Errorf("UcUserOperateHandler保存数据到数据库错误:%s", err.Error())
	}
	o.Cache = nil
}
