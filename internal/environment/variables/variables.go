package variables

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var e AllEnvironmentVariables

// AllEnvironmentVariables is a struct that defines all of the available environment variables
type AllEnvironmentVariables struct {
	Deployed                          bool   `envconfig:"DEPLOYED" required:"true"`
	LogLevel                          string `envconfig:"LOG_LEVEL" required:"true"`
	PostgresHost                      string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresUser                      string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword                  string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresPort                      int32  `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresSSLMode                   string `envconfig:"POSTGRES_SSLMODE" required:"true"`
	PostgresDatabase                  string `envconfig:"POSTGRES_DATABASE" required:"true"`
	PostgresLazyConnect               bool   `envconfig:"POSTGRES_LAZY_CONNECT" required:"false"`
	PostgresPoolMaxConnectionLifetime string `envconfig:"POSTGRES_POOL_MAX_CONNECTION_LIFETIME" required:"false"`
	PostgresPoolMaxConnectionIdleTime string `envconfig:"POSTGRES_POOL_MAX_CONNECTION_IDLE_TIME" required:"false"`
	PostgresPoolMaxConnections        int32  `envconfig:"POSTGRES_POOL_MAX_CONNECTIONS" required:"false"`
	PostgresPoolMinConnections        int32  `envconfig:"POSTGRES_POOL_MIN_CONNECTIONS" required:"false"`
	PostgresPoolHealthCheckPeriod     int32  `envconfig:"POSTGRES_POOL_HEALTH_CHECK_PERIOD" required:"false"`
}

// FetchAllEnvironmentVariables returns all environment variables defined for the service
func FetchAllEnvironmentVariables() AllEnvironmentVariables {
	if err := envconfig.Process(".env", &e); err != nil {
		log.Fatal(err)
	}

	if !e.Deployed {
		log.WithFields(log.Fields{
			"DEPLOYED":  e.Deployed,
			"LOG_LEVEL": e.LogLevel,
		}).Info("All Environment Variables")
	} else {
		log.WithFields(log.Fields{
			"DEPLOYED":  e.Deployed,
			"LOG_LEVEL": e.LogLevel,
		}).Info("All Environment Variables")
	}

	return e
}
