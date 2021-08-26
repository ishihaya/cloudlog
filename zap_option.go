package cloudlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type serviceContext struct {
	Service string
}

// MarshalLogObject add value in object.
// see: https://pkg.go.dev/go.uber.org/zap#Object
func (s serviceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("service", s.Service)
	return nil
}

// GetCloudServiceContextOption get zap.Option about serviceName.
// You have to call this function when you use Error Reporting.
// e.g. cloud-run-backend-api
func GetCloudServiceContextOption(serviceName string) zap.Option {
	return zap.Fields(zap.Object("serviceContext", serviceContext{Service: serviceName}))
}

// AddCloudErrorReportingOption add zap option about Error Reporting.
// You can use Error Reporting in GCP when you use this function.
// see: https://cloud.google.com/error-reporting/docs/formatting-error-messages?hl=ja#json_representation
func AddCloudErrorReportingOption(zapLogger *zap.Logger) *zap.Logger {
	typeKey := "@type"
	typeValue := "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent"
	return zapLogger.With(zap.String(typeKey, typeValue))
}
