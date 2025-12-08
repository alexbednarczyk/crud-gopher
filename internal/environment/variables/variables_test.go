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

	// Set required environment variables for testing
	os.Setenv("DEPLOYED", "false")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_SSLMODE", "disable")
	os.Setenv("POSTGRES_DATABASE", "testdb")

	e := FetchAllEnvironmentVariables()

	assert.NotNil(t, e.Deployed)
	assert.NotEmpty(t, e.LogLevel)

	os.Setenv("DEPLOYED", "true")
	e = FetchAllEnvironmentVariables()

	assert.NotNil(t, e.Deployed)
	assert.Equal(t, true, e.Deployed)
}
