package services

import (
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
)

type ProductCachedData struct {
	data map[int64]string
}

func RefreshProductCache() {
	go CacheProduct()
}

func CacheProduct() {
	defer iotutil.PanicHandler()
	rep, err := rpc.ClientOpmProductService.Lists(context.Background(), &protosService.OpmProductListRequest{
		IsPlatform: true,
	})
	if err != nil {
		return
	}
	if rep.Code != 200 {
		return
	}
	for _, item := range rep.Data {
		//productId := iotutil.ToString(item.Id)
		iotredis.GetClient().HMSet(context.Background(), iotconst.HKEY_PRODUCT_DATA+item.ProductKey, item).Err()
	}
}

func (s *ProductCachedData) GetProductName(productKey string) (str string, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, "name")
	str = strCmd.String()
	return
}

func (s *ProductCachedData) GetProductByFieldName(productKey string, fieldName string) (str interface{}, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGet(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey, fieldName)
	str = strCmd.String()
	return
}

func (s *ProductCachedData) GetByProduct(productKey string) (product *protosService.OpmProduct, err error) {
	defer iotutil.PanicHandler()
	strCmd := iotredis.GetClient().HGetAll(context.Background(), iotconst.HKEY_PRODUCT_DATA+productKey)
	str := strCmd.String()
	err = iotutil.JsonToStruct(str, product)
	return
}
