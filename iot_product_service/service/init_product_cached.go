package service

import (
	iotmodel "cloud_platform/iot_model"
	"cloud_platform/iot_model/db_product/model"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
	"math"
	"time"

	"go-micro.dev/v4/logger"

	"cloud_platform/iot_model/db_product/orm"
)

// 缓存更新时间
var updateTime time.Time

// ProductCached 语言缓存处理（定时更新缓存、初始化设置缓存）
func ProductCached() {
	//初始化语言缓存
	go initCached()
	//初始化更新时间为当前时间
	updateTime = time.Now()
	//定时更新语言缓存
	timeUpdateCached()
}

func initCached() {
	t := orm.Use(iotmodel.GetDB()).TOpmProduct
	do := t.WithContext(context.Background())

	totalCount, err := do.Count()
	if err != nil {
		logger.Errorf("initCached cached error : %s", err.Error())
		return
	}
	limit := 20
	pageCount := int(math.Ceil(float64(totalCount) / float64(limit)))
	for page := 1; page <= pageCount; page++ {
		offset := limit * (page - 1)
		list, err := do.Offset(offset).Limit(limit).Find()
		if err != nil {
			continue
		}
		if len(list) == 0 {
			continue
		}
		setCached(list)
		//时间
		time.Sleep(3 * time.Second)
	}
}

func timeUpdateCached() {
	//十分钟同步一次
	ticket := time.NewTicker(10 * time.Minute)
	go func() {
		for _ = range ticket.C {
			updateCached()
			//更新为最新的时间
			updateTime = time.Now()
		}
	}()
}

func updateCached() {
	defer iotutil.PanicHandler()
	t := orm.Use(iotmodel.GetDB()).TOpmProduct
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

// 缓存初始化
func setCached(list []*model.TOpmProduct) {
	for _, item := range list {
		cacheKey := iotconst.HKEY_PRODUCT_DATA + iotutil.ToString(item.ProductKey)
		mapSave := map[string]interface{}{
			"name":           item.Name,
			"productName":    item.Name,
			"productKey":     item.ProductKey,
			"nameEn":         item.NameEn,
			"imageUrl":       item.ImageUrl,
			"wifiFlag":       item.WifiFlag,
			"networkType":    item.NetworkType,
			"controlPanelId": item.ControlPanelId,
			"moduleId":       item.ModuleId,
		}
		thingModelSvc := OpmThingModelSvc{Ctx: context.Background()}
		thingModelData, err := thingModelSvc.GetOpmThingModelByProduct(&proto.OpmThingModelByProductRequest{
			ProductId: item.Id,
			Custom:    -1,
		})
		if err == nil {
			for _, property := range thingModelData.Properties {
				//缓存物模型的内容
				mapSave[iotconst.FIELD_PREFIX_TLS+iotutil.ToString(property.Dpid)] = iotutil.ToString(map[string]interface{}{
					"identifier":    property.Identifier,
					"dataType":      property.DataType,
					"name":          property.Name,
					"rwFlag":        property.RwFlag,
					"dataSpecs":     property.DataSpecs,
					"dataSpecsList": property.DataSpecsList,
					"custom":        property.Custom,
					"dpid":          property.Dpid,
				})
			}
		}
		iotredis.GetClient().HMSet(context.Background(), cacheKey, mapSave).Err()

		//刷新面板与产品的关系，用于面板更新之后刷新缓存；
		langKeysCached := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG_KEYS, item.ControlPanelId)
		cachedKey := persist.GetRedisKey(iotconst.APP_PRODUCT_PANEL_LANG, item.ProductKey)
		iotredis.GetClient().HMSet(context.Background(), langKeysCached, cachedKey, map[string]interface{}{cachedKey: "1"})
	}
}
