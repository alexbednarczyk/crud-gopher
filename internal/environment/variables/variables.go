package variables

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// AllEnvironmentVariables is a struct that defines all of the available environment variables.
// Sensible defaults are provided to make local development and tests frictionless.
type AllEnvironmentVariables struct {
	Deployed                          bool   `envconfig:"DEPLOYED" default:"false"`
	LogLevel                          string `envconfig:"LOG_LEVEL" default:"info"`
	PostgresHost                      string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresUser                      string `envconfig:"POSTGRES_USER" default:"postgres"`
	PostgresPassword                  string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresPort                      int32  `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresSSLMode                   string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
	PostgresDatabase                  string `envconfig:"POSTGRES_DATABASE" default:"postgres"`
	PostgresLazyConnect               bool   `envconfig:"POSTGRES_LAZY_CONNECT" default:"false"`
	PostgresPoolMaxConnectionLifetime string `envconfig:"POSTGRES_POOL_MAX_CONNECTION_LIFETIME" default:""`
	PostgresPoolMaxConnectionIdleTime string `envconfig:"POSTGRES_POOL_MAX_CONNECTION_IDLE_TIME" default:""`
	PostgresPoolMaxConnections        int32  `envconfig:"POSTGRES_POOL_MAX_CONNECTIONS" default:"0"`
	PostgresPoolMinConnections        int32  `envconfig:"POSTGRES_POOL_MIN_CONNECTIONS" default:"0"`
	PostgresPoolHealthCheckPeriod     int32  `envconfig:"POSTGRES_POOL_HEALTH_CHECK_PERIOD" default:"0"`
}

// Load reads environment variables. Optional env files can be provided to
// mirror production config locally.
func Load(envFiles ...string) (AllEnvironmentVariables, error) {
	if len(envFiles) > 0 {
		if err := godotenv.Load(envFiles...); err != nil {
			return AllEnvironmentVariables{}, fmt.Errorf("load env file: %w", err)
		}
	} else {
		// Ignore missing .env during tests/local runs.
		_ = godotenv.Load()
	}

	var env AllEnvironmentVariables
	if err := envconfig.Process("", &env); err != nil {
		return AllEnvironmentVariables{}, fmt.Errorf("process env vars: %w", err)
	}

	log.WithFields(log.Fields{
		"DEPLOYED":  env.Deployed,
		"LOG_LEVEL": env.LogLevel,
	}).Info("environment loaded")

	return env, nil
}

// MustLoad is a convenience wrapper that terminates on error.
func MustLoad(envFiles ...string) AllEnvironmentVariables {
	env, err := Load(envFiles...)
	if err != nil {
		log.Fatal(err)
	}
	return env
}
