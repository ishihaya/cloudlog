package cloudlog

// Option gives functions changed loggerOptions member.
// It uses Functional Option Pattern.
type Option func(*loggerOptions)

type loggerOptions struct {
	needErrorReporting bool
	serviceName        string
	logLevel           string
}

const (
	defaultNeedErrorReporting = false
	defaultServiceName        = "golang-app"
	defaultLogLevel           = "debug"
)

// NeedErrorReporting can assign needErrorReporting.
func NeedErrorReporting(needErrorReporting bool) Option {
	return func(lo *loggerOptions) {
		lo.needErrorReporting = needErrorReporting
	}
}

// ServiceName can assign serviceName.
func ServiceName(serviceName string) Option {
	return func(lo *loggerOptions) {
		lo.serviceName = serviceName
	}
}

// LogLevel can assign logLevel.
func LogLevel(logLevel string) Option {
	return func(lo *loggerOptions) {
		lo.logLevel = logLevel
	}
}

func getOptions(options []Option) (
	needErrorReporting bool,
	serviceName string,
	logLevel string,
) {
	loggerOps := &loggerOptions{
		needErrorReporting: defaultNeedErrorReporting,
		serviceName:        defaultServiceName,
		logLevel:           defaultLogLevel,
	}
	for _, option := range options {
		option(loggerOps)
	}
	needErrorReporting = loggerOps.needErrorReporting
	serviceName = loggerOps.serviceName
	logLevel = loggerOps.logLevel
	return
}
