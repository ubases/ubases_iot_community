package service

import (
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"

	"go-micro.dev/v4/metadata"
)

// CheckTenantId 公共方法，检查是否存在租户Id
func CheckTenantId(ctx context.Context) (string, error) {
	tenantId, _ := metadata.Get(ctx, "tenantid")
	if tenantId == "" {
		tenantId, _ := metadata.Get(ctx, "tenantid")
		if tenantId == "" {
			return "", errors.New("tenantId not found")
		}
	}
	return tenantId, nil
}

// CheckUserId 公共方法，检查是否存在用户
func CheckUserId(ctx context.Context) (string, int64, error) {
	userId, _ := metadata.Get(ctx, "Userid")
	if userId == "" {
		return "", 0, errors.New("Userid not found")
	}
	userIdInt, err := iotutil.ToInt64AndErr(userId)
	if err != nil {
		return "", 0, err
	}
	return userId, userIdInt, nil
}

// 获取区域Id，默认区域为1
func GetRegionInt(ctx context.Context) int64 {
	region, _ := metadata.Get(ctx, "region")
	var defaultRegion int64 = 1
	if region == "" {
		return defaultRegion
	} else {
		regionInt, err := iotutil.ToInt64AndErr(region)
		if err != nil {
			return defaultRegion
		}
		return regionInt
	}
}
