package services

import (
	"cloud_platform/iot_cloud_api_service/cached"
	"cloud_platform/iot_cloud_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotredis"
	"cloud_platform/iot_common/iotutil"
	proto "cloud_platform/iot_proto/protos/protosService"
	"context"
)

type DeveloperCachedData struct {
	data map[string]*proto.OpenCompany
}

func RefreshDevelopCache() {
	go CacheDeveloper()
}

func CacheDeveloper() {
	defer iotutil.PanicHandler()
	resp, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &proto.OpenCompanyListRequest{
		Query: nil,
	})
	if err != nil {
		return
	}
	result := make(map[string]*proto.OpenCompany, 0)
	for _, item := range resp.Data {
		tenantId := item.TenantId
		if _, ok := result[tenantId]; !ok {
			result[tenantId] = &proto.OpenCompany{}
		}
		result[tenantId] = item
	}
	for k, m := range result {
		iotredis.GetClient().Set(context.Background(), k, iotutil.ToString(m), 0)
	}
}

func (s *DeveloperCachedData) GetByTenantId(tenantId string) (res *proto.OpenCompany, err error) {
	res = &proto.OpenCompany{}
	if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.DEVELOPER_TENANT_ID_DATA, tenantId), res); err == nil {
		return res, nil
	}
	resp, err := rpc.ClientOpenCompanyService.Lists(context.Background(), &proto.OpenCompanyListRequest{
		Query: &proto.OpenCompany{TenantId: tenantId},
	})
	if err != nil {
		return res, err
	}
	if len(resp.Data) == 0 {
		return res, err
	}
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.DEVELOPER_TENANT_ID_DATA, tenantId), resp.Data[0], 0)
	if err != nil {
		return res, err
	}
	return
}
