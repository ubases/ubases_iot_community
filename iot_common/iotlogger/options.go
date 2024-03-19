package iotlogger

import (
	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Options struct {
	logger.Options
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) logger.Option {
	return logger.SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) logger.Option {
	return logger.SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) logger.Option {
	return logger.SetOption(namespaceKey{}, namespace)
}

type optionsKey struct{}

func WithOptions(opts ...zap.Option) logger.Option {
	return logger.SetOption(optionsKey{}, opts)
}

type callerSkipKey struct{}

func WithCallerSkip(i int) logger.Option {
	return logger.SetOption(callerSkipKey{}, i)
}

type logFileNameKey struct{}

// WithLogFileName pass log file name to logger
func WithLogFileName(name string) logger.Option {
	return logger.SetOption(logFileNameKey{}, name)
}
