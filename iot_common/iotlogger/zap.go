package iotlogger

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go-micro.dev/v4/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zaplog struct {
	cfg  zap.Config
	zap  *zap.Logger
	opts logger.Options
	sync.RWMutex
	fields map[string]interface{}
}

func (l *zaplog) Init(opts ...logger.Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	zapConfig := zap.NewProductionConfig()
	if zconfig, ok := l.opts.Context.Value(configKey{}).(zap.Config); ok {
		zapConfig = zconfig
	}

	if zcconfig, ok := l.opts.Context.Value(encoderConfigKey{}).(zapcore.EncoderConfig); ok {
		zapConfig.EncoderConfig = zcconfig
	}
	zapConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	skip, ok := l.opts.Context.Value(callerSkipKey{}).(int)
	if !ok || skip < 1 {
		skip = 2
	}

	zapConfig.Level = zap.NewAtomicLevel()
	if l.opts.Level != logger.InfoLevel {
		zapConfig.Level.SetLevel(loggerToZapLevel(l.opts.Level))
	}
	var allCore []zapcore.Core

	//输出到文件
	if logFileName, ok := l.opts.Context.Value(logFileNameKey{}).(string); ok {
		hook := &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     10,
			LocalTime:  true,
			Compress:   true,
		}
		jsonEncoder := zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
		fileWriter := zapcore.AddSync(hook)
		core := zapcore.NewCore(jsonEncoder, fileWriter, zapConfig.Level)
		allCore = append(allCore, core)
	}

	//输出到控制台
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consolecore := zapcore.NewCore(consoleEncoder, consoleDebugging, zapConfig.Level)
	allCore = append(allCore, consolecore)

	core := zapcore.NewTee(allCore...)

	//log := zap.New(core,
	//	zap.AddCaller(),
	//	zap.AddCallerSkip(skip),
	//	zap.AddStacktrace(zapcore.ErrorLevel))

	log := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(skip),
		zap.AddStacktrace(zapcore.PanicLevel))

	if l.opts.Fields != nil {
		data := []zap.Field{}
		for k, v := range l.opts.Fields {
			data = append(data, zap.Any(k, v))
		}
		log = log.With(data...)
	}

	if namespace, ok := l.opts.Context.Value(namespaceKey{}).(string); ok {
		log = log.With(zap.Namespace(namespace))
	}

	if options, ok := l.opts.Context.Value(optionsKey{}).([]zap.Option); ok {
		log = log.WithOptions(options...)
	}

	l.cfg = zapConfig
	l.zap = log
	l.fields = make(map[string]interface{})

	return nil
}

func (l *zaplog) Fields(fields map[string]interface{}) logger.Logger {
	l.Lock()
	nfields := make(map[string]interface{}, len(l.fields))
	for k, v := range l.fields {
		nfields[k] = v
	}
	l.Unlock()
	for k, v := range fields {
		nfields[k] = v
	}

	data := make([]zap.Field, 0, len(nfields))
	for k, v := range fields {
		data = append(data, zap.Any(k, v))
	}

	zl := &zaplog{
		cfg:    l.cfg,
		zap:    l.zap.With(data...),
		opts:   l.opts,
		fields: make(map[string]interface{}),
	}

	return zl
}

func (l *zaplog) Error(err error) logger.Logger {
	return l.Fields(map[string]interface{}{"error": err})
}

func (l *zaplog) Log(level logger.Level, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zaplog) Logf(level logger.Level, format string, args ...interface{}) {
	l.RLock()
	data := make([]zap.Field, 0, len(l.fields))
	for k, v := range l.fields {
		data = append(data, zap.Any(k, v))
	}
	l.RUnlock()

	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	switch lvl {
	case zap.DebugLevel:
		l.zap.Debug(msg, data...)
	case zap.InfoLevel:
		l.zap.Info(msg, data...)
	case zap.WarnLevel:
		l.zap.Warn(msg, data...)
	case zap.ErrorLevel:
		l.zap.Error(msg, data...)
	case zap.FatalLevel:
		l.zap.Fatal(msg, data...)
	}
}

func (l *zaplog) String() string {
	return "zap"
}

func (l *zaplog) Options() logger.Options {
	return l.opts
}

func NewLogger(opts ...logger.Option) (logger.Logger, error) {
	options := logger.Options{
		Level:           logger.InfoLevel,
		Fields:          make(map[string]interface{}),
		Out:             os.Stderr,
		Context:         context.Background(),
		CallerSkipCount: 2,
	}

	l := &zaplog{opts: options}
	if err := l.Init(opts...); err != nil {
		return nil, err
	}

	return l, nil
}

func loggerToZapLevel(level logger.Level) zapcore.Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return zap.DebugLevel
	case logger.InfoLevel:
		return zap.InfoLevel
	case logger.WarnLevel:
		return zap.WarnLevel
	case logger.ErrorLevel:
		return zap.ErrorLevel
	case logger.FatalLevel:
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func zapToLoggerLevel(level zapcore.Level) logger.Level {
	switch level {
	case zap.DebugLevel:
		return logger.DebugLevel
	case zap.InfoLevel:
		return logger.InfoLevel
	case zap.WarnLevel:
		return logger.WarnLevel
	case zap.ErrorLevel:
		return logger.ErrorLevel
	case zap.FatalLevel:
		return logger.FatalLevel
	default:
		return logger.InfoLevel
	}
}
