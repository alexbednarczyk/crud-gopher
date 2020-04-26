package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestSetLogLevel(t *testing.T) {

	SetLogLevel("foo")
	assert.Equal(t, log.DebugLevel, log.GetLevel())

	SetLogLevel("debug")
	assert.Equal(t, log.DebugLevel, log.GetLevel())

	SetLogLevel("trace")
	assert.Equal(t, log.TraceLevel, log.GetLevel())

	SetLogLevel("info")
	assert.Equal(t, log.InfoLevel, log.GetLevel())

	SetLogLevel("warn")
	assert.Equal(t, log.WarnLevel, log.GetLevel())

	SetLogLevel("error")
	assert.Equal(t, log.ErrorLevel, log.GetLevel())

	SetLogLevel("panic")
	assert.Equal(t, log.PanicLevel, log.GetLevel())

	SetLogLevel("fatal")
	assert.Equal(t, log.FatalLevel, log.GetLevel())
}

func TestStandardFormatLogs(t *testing.T) {

	assert.Equal(t, false, StandardFormatLogs(false))

	assert.Equal(t, true, StandardFormatLogs(true))

}
