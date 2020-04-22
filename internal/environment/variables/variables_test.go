package variables

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentVariables(t *testing.T) {

	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	e := FetchAllEnvironmentVariables()

	assert.NotNil(t, e.Deployed)
	assert.NotEmpty(t, e.LogLevel)

	os.Setenv("DEPLOYED", "true")
	e = FetchAllEnvironmentVariables()

	assert.NotNil(t, e.Deployed)
	assert.Equal(t, true, e.Deployed)
}
