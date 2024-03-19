package services

import (
	"cloud_platform/iot_app_api_service/cached"
	"cloud_platform/iot_app_api_service/rpc"
	"cloud_platform/iot_common/iotconst"
	"cloud_platform/iot_common/iotgincache/persist"
	"cloud_platform/iot_common/iotutil"
	"cloud_platform/iot_proto/protos/protosService"
	"context"
	"fmt"
)

// 家庭详细缓存Keys
func ReadHomeDetailsCachedKey(homeId, userId string) []string {
	keys := []string{}
	keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(homeId), iotutil.ToString(userId)))
	return keys
}

// 家庭家庭房间列表
func ReadHomeRoomListsCachedKey(homeId int64) []string {
	keys := make([]string, 0)
	for _, l := range iotconst.APP_SUPPORT_LANGUAGE {
		keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_ROOM_LIST_DATA, l, iotutil.ToString(homeId)))
	}
	return keys
}

func ClearHomeCached(userId int64, clearRoom bool, keys ...string) error {
	// 删除家庭详情缓存
	if keys == nil {
		keys = []string{}
	}
	keys = append(keys, persist.GetRedisKey(iotconst.APP_HOME_LIST_DATA, userId))

	ctx := context.Background()
	resp, err := rpc.UcHomeUserService.Lists(context.Background(), &protosService.UcHomeUserListRequest{
		Query: &protosService.UcHomeUser{
			UserId: iotutil.ToInt64(userId),
		},
	})
	if err != nil {
		return err
	}
	for i := range resp.Data {
		keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_DETAIL_DATA, iotutil.ToString(resp.Data[i].HomeId), userId))
		if clearRoom {
			keys = append(keys, fmt.Sprintf(iotconst.APP_HOME_ROOM_LIST_DATA, iotutil.ToString(resp.Data[i].HomeId)))
		}
	}
	if len(keys) != 0 {
		pipe := cached.RedisStore.Pipeline()
		pipe.Del(ctx, keys...)
		if _, err := pipe.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
