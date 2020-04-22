package variables

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var e AllEnvironmentVariables

// AllEnvironmentVariables is a struct that defines all of the available environment variables
type AllEnvironmentVariables struct {
	Deployed bool   `envconfig:"DEPLOYED" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
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
