package service

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_config/model"
	"context"
	"fmt"
	"math"
	"time"

	"go-micro.dev/v4/logger"

	"cloud_platform/iot_model/db_config/orm"
)

// 缓存更新时间
var updateTime time.Time

// 语言缓存处理（定时更新缓存、初始化设置缓存）
func LangCached() {
	//初始化语言缓存
	go initCached()
	//初始化更新时间为当前时间
	updateTime = time.Now()
	//定时更新语言缓存
	timeUpdateCached()
}

// 启动定时器执行
func initCached() {
	t := orm.Use(iotmodel.GetDB()).TLangTranslate
	do := t.WithContext(context.Background())

	totalCount, err := do.Count()
	if err != nil {
		logger.Errorf("setCahced cached error : %s", err.Error())
		return
	}
	limit := 1000
	pageCount := int(math.Ceil(float64(totalCount) / float64(limit)))
	for page := 1; page <= pageCount; page++ {
		offset := limit * (page - 1)
		list, err := do.Order(t.Id).Offset(offset).Limit(limit).Find()
		if err != nil {
			continue
		}
		if len(list) == 0 {
			continue
		}
		setCached(list)
	}
}

func timeUpdateCached() {
	ticket := time.NewTicker(5 * time.Minute)
	go func() {
		for _ = range ticket.C {
			updateCached()
			//更新为最新的时间
			updateTime = time.Now()
		}
	}()
}

func updateCached() {
	t := orm.Use(iotmodel.GetDB()).TLangTranslate
	do := t.WithContext(context.Background())
	//考虑10s左右的同步误差时间
	list, err := do.Where(t.UpdatedAt.Gte(updateTime.Add(-10 * time.Second))).Find()
	if err != nil {
		return
	}
	if len(list) == 0 {
		return
	}
	setCached(list)
}

func setCached(list []*model.TLangTranslate) {
	cachedData := map[string]map[string]interface{}{}
	for _, item := range list {
		cacheKey := iotconst.HKEY_LANGUAGE_DATA_PREFIX + item.SourceTable
		//开放平台缓存
		if item.PlatformType == int32(iotconst.OPEN_USER) {
			cacheKey = fmt.Sprintf("%s_%s", item.TenantId, cacheKey)
		}
		if _, ok := cachedData[cacheKey]; !ok {
			cachedData[cacheKey] = map[string]interface{}{}
		}
		key := item.Lang + "_" + item.SourceRowId + "_" + item.FieldName
		cachedData[cacheKey][key] = item.FieldValue
	}
	for k, v := range cachedData {
		resCmd := iotredis.GetClient().HMSet(context.Background(), k, v)
		if resCmd.Err() != nil {
			logger.Errorf("setCahced cached error : %s", resCmd.Err().Error())
		}
	}
}
