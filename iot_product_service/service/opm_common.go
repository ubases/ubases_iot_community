package service

import (
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

// CheckTenantId 公共方法，检查是否存在租户Id
func CheckTenantId(ctx context.Context) (string, error) {
	tenantId, _ := metadata.Get(ctx, "tenantid")
	if tenantId == "" {
		logger.Errorf("tenantId not found")
		return "", errors.New("tenantId not found")
	}
	return tenantId, nil
}

// GetUserId 公共方法，检查是否存在用户Id
func GetUserId(ctx context.Context) (string, error) {
	userId, _ := metadata.Get(ctx, "Userid")
	if userId == "" {
		logger.Errorf("Userid not found")
		return "", errors.New("Userid not found")
	}
	return userId, nil
}

func GetUserIdInt64(ctx context.Context) (int64, error) {
	userId, _ := metadata.Get(ctx, "Userid")
	if userId == "" {
		logger.Errorf("Userid not found")
		return 0, errors.New("Userid not found")
	}
	return iotutil.ToInt64(userId), nil
}
