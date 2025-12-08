package config

import (
	"testing"
	"time"

	"github.com/alexbednarczyk/crud-gopher/internal/environment/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildMigrationURLConnectionString(t *testing.T) {
	env := variables.AllEnvironmentVariables{
		PostgresUser:     "user",
		PostgresPassword: "pass",
		PostgresHost:     "db",
		PostgresPort:     5432,
		PostgresDatabase: "service",
		PostgresSSLMode:  "verify-full",
	}

	conn := BuildMigrationURLConnectionString(env)
	assert.Equal(t, "postgres://user:pass@db:5432/service?sslmode=verify-full", conn)
}

func TestBuildPoolConfig(t *testing.T) {
	env := variables.AllEnvironmentVariables{
		PostgresUser:                      "user",
		PostgresPassword:                  "pass",
		PostgresHost:                      "db",
		PostgresPort:                      5432,
		PostgresDatabase:                  "service",
		PostgresSSLMode:                   "disable",
		PostgresPoolMaxConnections:        10,
		PostgresPoolMinConnections:        2,
		PostgresPoolMaxConnectionLifetime: "1h",
		PostgresPoolMaxConnectionIdleTime: "10m",
		PostgresPoolHealthCheckPeriod:     5,
	}

	cfg, err := buildPoolConfig(env)
	require.NoError(t, err)

	assert.Equal(t, int32(10), cfg.MaxConns)
	assert.Equal(t, int32(2), cfg.MinConns)
	assert.Equal(t, time.Hour, cfg.MaxConnLifetime)
	assert.Equal(t, 10*time.Minute, cfg.MaxConnIdleTime)
	assert.Equal(t, 5*time.Second, cfg.HealthCheckPeriod)
}
