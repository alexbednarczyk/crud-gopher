package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCurrentTime(t *testing.T) {

	now := time.Now().UTC()
	currentTime, _ := CurrentTime()

	assert.Equal(t, currentTime.Year(), now.Year())
	assert.Equal(t, currentTime.Month(), now.Month())
	assert.Equal(t, currentTime.Day(), now.Day())
}
