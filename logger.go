package cloudlog

import (
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

// Logger provides structured logging in Go / Zap.
type Logger struct {
	zapLogger      *zap.Logger
	sugar          *zap.SugaredLogger
	errorZapLogger *zap.Logger
	errorSugar     *zap.SugaredLogger
}

// NewCloudLogger is logger constructor in Cloud Environment.
// It has optional arguments about logger options written on options.go.
// It has been set Error Reporting and Cloud Logging in GCP by default.
func NewCloudLogger(options ...Option) (logger *Logger, err error) {
	needErrorReporting, serviceName, logLevel := getOptions(options)
	config := NewCloudZapConfig(logLevel, "json")
	option := GetCloudServiceContextOption(serviceName)
	zapLogger, err := config.Build(option)
	if err != nil {
		return nil, xerrors.Errorf("init logging cloud configs error: %w", err)
	}
	var errorZapLogger *zap.Logger
	if needErrorReporting {
		errorZapLogger = AddCloudErrorReportingOption(zapLogger)
	} else {
		errorZapLogger = zapLogger
	}
	logger = &Logger{
		zapLogger:      zapLogger,
		sugar:          zapLogger.Sugar(),
		errorZapLogger: errorZapLogger,
		errorSugar:     errorZapLogger.Sugar(),
	}
	return
}

// NewLocalLogger is logger constructor in Local Environment.
// It has optional arguments about logger options written on options.go.
func NewLocalLogger(options ...Option) (logger *Logger, err error) {
	_, _, logLevel := getOptions(options)
	config := NewLocalZapConfig(logLevel, "console")
	zapLogger, err := config.Build()
	if err != nil {
		return nil, xerrors.Errorf("init logging local configs error: %w", err)
	}
	errorZapLogger := zapLogger
	logger = &Logger{
		zapLogger:      zapLogger,
		sugar:          zapLogger.Sugar(),
		errorZapLogger: errorZapLogger,
		errorSugar:     errorZapLogger.Sugar(),
	}
	return
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
// e.g. Warnf("something went wrong: %+v", err)
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
func (l *Logger) Warnf(template string, args ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Warnf(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
// e.g. Errorf("something went wrong: %+v", err)
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
func (l *Logger) Errorf(template string, args ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.errorSugar.Errorf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.errorSugar.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
// e.g. WarnW("something went wrong", "key", "value", "sum", 10)
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.sugar.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
// e.g. WarnW("something went wrong", "key", "value", "sum", 10)
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.errorSugar.Errorw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	defer func() {
		_ = l.sugar.Sync()
	}()
	l.errorSugar.Fatalw(msg, keysAndValues)
}

// Debug logs a message at DebugLevel.
func (l *Logger) Debug(msg string, field ...zap.Field) {
	l.zapLogger.Debug(msg, field...)
}

// Info logs a message at InfoLevel.
func (l *Logger) Info(msg string, field ...zap.Field) {
	l.zapLogger.Info(msg, field...)
}

// Warn logs a message at WarnLevel.
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
// You shoud use this method when you want faster logging, but if you don't want to use zap.Field and depend them, you can use Warnf or Warnw.
func (l *Logger) Warn(msg string, field ...zap.Field) {
	l.zapLogger.Warn(msg, field...)
}

// Error logs a message at ErrorLevel.
// If you don't need these wrapper methods, you can override them by parent interface defined on your own.
// You shoud use this method when you want faster logging, but if you don't want to use zap.Field and depend them, you can use Warnf or Warnw.
func (l *Logger) Error(msg string, field ...zap.Field) {
	l.errorZapLogger.Error(msg, field...)
}

// Fatal logs a message at FatalLevel.
func (l *Logger) Fatal(msg string, field ...zap.Field) {
	l.errorZapLogger.Fatal(msg, field...)
}
