package common

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

// SetLogLevel function sets the log level
func SetLogLevel(logLevel string) {
	switch logLevel {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

// StandardFormatLogs sets the log formatting
func StandardFormatLogs(deployed bool) bool {

	if deployed {
		formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
		formatter.Line = true
		formatter.Package = true
		formatter.File = true
		log.SetFormatter(&formatter)
		log.SetOutput(os.Stdout)
	} else {
		formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{}}
		formatter.Line = true
		formatter.Package = true
		formatter.File = true
		log.SetFormatter(&formatter)
		log.SetOutput(os.Stdout)
	}
	return deployed
}
