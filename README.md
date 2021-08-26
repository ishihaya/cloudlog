# cloudlog

[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/ishihaya/cloudlog)](https://github.com/ishihaya/cloudlog/releases)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/ishihaya/cloudlog)

## Comments

**If you have any feature requests, please feel free to create an issue. We're looking forward to communicate with you!**

## Summary

cloudlog provides structured loggging in Go.

We recommend to use cloudlog with GCP, 
support 
- **[Cloud Logging](https://cloud.google.com/logging/docs/how-to)**(**Stackdriver Logging**) display log level, structured logging, etc...
- **[Error Reporting](https://cloud.google.com/error-reporting/docs/formatting-error-messages)**

format.

you can use logging features **immediately** when you only need default log config.

It uses [zap](https://github.com/uber-go/zap) internal.

## Installation

```
$ go get -u github.com/ishihaya/cloudlog
```

## Usage

You can define methods that you only use in `Log` interface.

[logging methods list](./logger.go)

```go
package log

import (
	"os"
	"sync"

	"github.com/ishihaya/cloudlog"
)

// You can define methods that you only use.
type Log interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

type log struct {
	*cloudlog.Logger
}

var sharedInstance Log
var once sync.Once

// You should call this if you use logger.
func GetInstance() Log {
	once.Do(func() {
		sharedInstance = new()
	})
	return sharedInstance
}

func new() Log {
	var logger *cloudlog.Logger
	var err error
	// serviceName is displayed in Error Reporting.
	serviceName := "backend-api"
	switch os.Getenv("APP_ENV") {
	// List runnning environments on cloud, such as GCP.
	case "production", "staging":
		logger, err = cloudlog.NewCloudLogger(
			cloudlog.NeedErrorReporting(true),
			cloudlog.ServiceName(serviceName),
		)
	default:
		logger, err = cloudlog.NewLocalLogger()
	}
	if err != nil {
		panic(err)
	}
	return &log{logger}
}
```
