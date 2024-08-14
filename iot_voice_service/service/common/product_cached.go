package common

import (
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"cloud_platform/iot_voice_service/rpc/rpcclient"
	"context"
	"errors"
)

type ProductCachedData struct {
	data map[int64]string
}

func setProductCached(item *protosService.OpmProduct) {
	productId := iotutil.ToString(item.Id)
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
		"panelUrl":       item.PanelUrl,
		"panelKey":       item.PanelKey,
		"tenantId":       item.TenantId,
	}
	iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productId, mapSave).Err()
	iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+item.ProductKey, mapSave).Err()
}

func (s *ProductCachedData) GetProductByFieldName(productKey string, fieldName string) (str interface{}, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, fieldName)
	str = strCmd.Val()
	return
}

func (s *ProductCachedData) GetProduct(productKey string) (product *protosService.OpmProduct, err error) {
	if productKey == "" {
		return nil, errors.New("productKey not found")
	}
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey)
	str := strCmd.Val() //.String()
	_ = iotutil.JsonToStruct(iotutil.ToString(str), &product)
	if product == nil {
		rep, psErr := rpcclient.ClientOpmProductService.Find(context.Background(), &protosService.OpmProductFilter{
			ProductKey: productKey,
		})
		if psErr != nil {
			return
		}
		if rep.Code != 200 {
			return
		}
		for _, item := range rep.Data {
			setProductCached(item)
		}
	}
	return
}
