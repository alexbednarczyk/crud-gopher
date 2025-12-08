package common

import (
	"testing"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/stretchr/testify/assert"

	log "github.com/sirupsen/logrus"
)

func TestSetLogLevel(t *testing.T) {

	SetLogLevel("DEBUG")
	assert.Equal(t, log.DebugLevel, log.GetLevel())

	SetLogLevel("trace")
	assert.Equal(t, log.TraceLevel, log.GetLevel())

	SetLogLevel("warn")
	assert.Equal(t, log.WarnLevel, log.GetLevel())

	SetLogLevel("error")
	assert.Equal(t, log.ErrorLevel, log.GetLevel())

	SetLogLevel("unknown")
	assert.Equal(t, log.InfoLevel, log.GetLevel())
}

func TestConfigureLogger(t *testing.T) {
	ConfigureLogger(false)
	_, ok := log.StandardLogger().Formatter.(*runtime.Formatter)
	assert.True(t, ok)
}
