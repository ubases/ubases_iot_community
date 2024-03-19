package service

import (
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

// CheckUserId 公共方法，检查是否存在用户
func CheckUserId(ctx context.Context) (string, error) {
	tenantId, _ := metadata.Get(ctx, "Userid")
	if tenantId == "" {
		logger.Errorf("Userid not found")
		return "", errors.New("Userid not found")
	}
	return tenantId, nil
}

// CheckUserId 公共方法，检查是否存在用户
func getLang(ctx context.Context) (string, error) {
	lang, _ := metadata.Get(ctx, "lang")
	if lang != "" {
		return lang, nil
	}
	return "zh", nil
}
