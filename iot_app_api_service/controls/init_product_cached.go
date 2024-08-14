package controls

import (
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"errors"
)

type ProductCachedData struct {
	data map[int64]string
}

func RefreshProductCache() {
	go CacheProduct()
}

func CacheProduct() {
	defer iotutil.PanicHandler()
	rep, err := rpc.ProductService.Lists(context.Background(), &protosService.OpmProductListRequest{
		IsPlatform: true,
	})
	if err != nil {
		return
	}
	if rep.Code != 200 {
		return
	}
	for _, item := range rep.Data {
		setProductCached(item)
	}
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
		"productTypeId":  item.ProductTypeId,
		"productId":      item.ProductId,
		"tenantId":       item.TenantId,
	}
	iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productId, mapSave).Err()
	iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+item.ProductKey, mapSave).Err()
}

func (s *ProductCachedData) GetProductName(productKey string) (str string, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, "name")
	str = strCmd.Val()
	if str == "" {
		pro, err := s.GetProduct(productKey)
		if err == nil {
			str = pro.NameEn
		}
	}
	return
}

func (s *ProductCachedData) GetProductByFieldName(productKey string, fieldName string) (str interface{}, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, fieldName)
	str = strCmd.Val()
	return
}

func (s *ProductCachedData) GetProductByFieldNames(productKey string, fieldName ...string) (res []interface{}, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HMGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, fieldName...)
	res = strCmd.Val()
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
		rep, psErr := rpc.ProductService.Lists(context.Background(), &protosService.OpmProductListRequest{
			IsPlatform: true,
			Query: &protosService.OpmProduct{
				ProductKey: productKey,
			},
		})
		if psErr != nil {
			return
		}
		if rep.Code != 200 {
			return
		}
		for _, item := range rep.Data {
			setProductCached(item)
			product = item
		}
	}
	return
}
