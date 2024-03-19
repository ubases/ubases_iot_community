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

func RefreshUserCache() {
	go CacheUser()
}

func CacheUser() {
	defer iotutil.PanicHandler()
	resp, err := rpc.ClientSysUserService.Lists(context.Background(), &proto.SysUserListRequest{
		Query: nil,
	})
	if err != nil {
		return
	}
	result := make(map[string]*proto.SysUser, 0)
	for _, item := range resp.Data {
		userId := iotutil.ToString(item.Id)
		if _, ok := result[userId]; !ok {
			result[userId] = &proto.SysUser{}
		}
		result[userId] = item
	}
	for k, m := range result {
		iotredis.GetClient().Set(context.Background(), k, iotutil.ToString(m), 0)
	}
}

type UserCachedData struct {
	data map[string]*proto.SysUser
}

func (s *UserCachedData) GetByUserId(userId string) (res *proto.SysUser, err error) {
	res = &proto.SysUser{}
	userIdInt, err := iotutil.ToInt64AndErr(userId)
	if err != nil {
		return res, err
	}
	//系统管理员默认判断
	if userIdInt == 0 {
		return &proto.SysUser{
			UserName:     "admin",
			UserNickname: "超级管理员",
			Id:           userIdInt,
		}, nil
	}

	if err := cached.RedisStore.Get(persist.GetRedisKey(iotconst.SYSTEM_USER_DATA, userId), res); err == nil {
		return res, nil
	}
	resp, err := rpc.ClientSysUserService.Lists(context.Background(), &proto.SysUserListRequest{
		Query: &proto.SysUser{Id: userIdInt},
	})
	if err != nil {
		return res, err
	}
	if len(resp.Data) == 0 {
		return res, err
	}
	err = cached.RedisStore.Set(persist.GetRedisKey(iotconst.SYSTEM_USER_DATA, userId), resp.Data[0], 0)
	if err != nil {
		return res, err
	}
	return
}
