package cloudlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewCloudZapConfig provides default zap config in cloud environment.
func NewCloudZapConfig(level, encoding string) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(unmarshalLogLevel(level)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    1,
			Thereafter: 1,
		},
		Encoding: encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "eventTime",
			LevelKey:       "severity",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stack_trace",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    encodeLevel,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// NewLocalZapConfig provides default zap config in local environment.
func NewLocalZapConfig(level, encoding string) zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(unmarshalLogLevel(level)),
		Development:      true,
		Encoding:         encoding,
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func encodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(marshalLogLevel(l))
}

// see: https://github.com/uber-go/zap/blob/425214515ff452748375576b20c82524849177c6/zapcore/level.go#L126-L146
func unmarshalLogLevel(text string) zapcore.Level {
	switch text {
	case "debug", "DEBUG":
		return zap.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		return zap.InfoLevel
	case "warn", "WARN":
		return zap.WarnLevel
	case "error", "ERROR":
		return zap.ErrorLevel
	case "dpanic", "DPANIC":
		return zap.DPanicLevel
	case "panic", "PANIC":
		return zap.PanicLevel
	case "fatal", "FATAL":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry?hl=ja#LogSeverity
func marshalLogLevel(l zapcore.Level) string {
	switch l {
	case zapcore.DebugLevel:
		return "DEBUG"
	case zapcore.InfoLevel:
		return "INFO"
	case zapcore.WarnLevel:
		return "WARNING"
	case zapcore.ErrorLevel:
		return "ERROR"
	case zapcore.DPanicLevel:
		return "CRITICAL"
	case zapcore.PanicLevel:
		return "ALERT"
	case zapcore.FatalLevel:
		return "EMERGENCY"
	default:
		return "INFO"
	}
}
