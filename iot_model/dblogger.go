package iotmodel

import (
	"cloud_platform/iot_common/iotlogger"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

var SlowThreshold = 200 * time.Millisecond

type gormLogger struct {
	LogLevel logger.LogLevel
}

func New(level logger.LogLevel) logger.Interface {
	return &gormLogger{
		LogLevel: level,
	}
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		iotlogger.LogHelper.Info(msg, data)
	}
}

func (l gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		iotlogger.LogHelper.Warn(msg, data)
	}
}

func (l gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		iotlogger.LogHelper.Error(msg, data)
	}
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	elapsed := time.Since(begin)
	printelapsed := strconv.Itoa(int(elapsed.Milliseconds())) + "ms"
	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		iotlogger.LogHelper.WithTag("elapsedTime", printelapsed).WithTag("sql", sql).WithTag("rows", strconv.Itoa(int(rows))).Error(err)
	case elapsed > SlowThreshold && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", SlowThreshold)
		iotlogger.LogHelper.WithTag("elapsedTime", printelapsed).WithTag("sql", sql).WithTag("rows", strconv.Itoa(int(rows))).Warn(slowLog)
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		iotlogger.LogHelper.WithTag("elapsedTime", printelapsed).WithTag("sql", sql).WithTag("rows", strconv.Itoa(int(rows))).Debug("trace")
	default:
		sql, rows := fc()
		iotlogger.LogHelper.WithTag("elapsedTime", printelapsed).WithTag("sql", sql).WithTag("rows", strconv.Itoa(int(rows))).Debug("trace")
	}
}
