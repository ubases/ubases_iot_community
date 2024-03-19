package service

import (
	"cloud_platform/iot_common/iotlogger"
	"cloud_platform/iot_common/iotutil"
	"context"
	"errors"

	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

// GetUserId 公共方法，检查是否存在用户Id
func GetUserId(ctx context.Context) (int64, error) {
	userId, _ := metadata.Get(ctx, "Userid")
	if userId == "" {
		logger.Errorf("Userid not found")
		return 0, errors.New("Userid not found")
	}
	return iotutil.ToInt64(userId), nil
}

// CheckTenantId 公共方法，检查是否存在租户Id
func CheckTenantId(ctx context.Context) (string, error) {
	tenantId, _ := metadata.Get(ctx, "tenantid")
	iotlogger.LogHelper.Info("tenantId---" + tenantId)
	if tenantId == "" {
		return "", errors.New("tenantId not found")
	}
	return tenantId, nil
}

// CheckTenantId 公共方法，检查是否存在租户Id
func CheckAppKey(ctx context.Context) (string, error) {
	appKey, _ := metadata.Get(ctx, "appkey")
	if appKey == "" {
		return "", errors.New("appKey not found")
	}
	return appKey, nil
}

// CheckLang 公共方法，检查是否存在lang
func CheckLang(ctx context.Context) (string, error) {
	lang, _ := metadata.Get(ctx, "lang")
	if lang == "" {
		return "", errors.New("lang not found")
	}
	return lang, nil
}
