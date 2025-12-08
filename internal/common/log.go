package common

import (
	"os"
	"strings"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

// SetLogLevel configures the global log level, accepting case-insensitive
// values. Unknown levels fall back to info to avoid noisy debug output in
// production environments.
func SetLogLevel(logLevel string) {
	level, err := log.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

// ConfigureLogger applies a consistent formatter and output destination.
// JSON is used when deployed for easier ingestion; text is used locally for
// readability. The runtime formatter preserves caller information.
func ConfigureLogger(deployed bool) {
	var child log.Formatter
	if deployed {
		child = &log.JSONFormatter{}
	} else {
		child = &log.TextFormatter{FullTimestamp: true}
	}

	formatter := runtime.Formatter{ChildFormatter: child}
	formatter.Line = true
	formatter.Package = true
	formatter.File = true

	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
}
