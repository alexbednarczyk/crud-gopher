package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadDefaults(t *testing.T) {
	env, err := Load()
	require.NoError(t, err)

	assert.False(t, env.Deployed)
	assert.Equal(t, "info", env.LogLevel)
	assert.Equal(t, "localhost", env.PostgresHost)
	assert.Equal(t, int32(5432), env.PostgresPort)
	assert.Equal(t, "disable", env.PostgresSSLMode)
}

func TestLoadOverrides(t *testing.T) {
	t.Setenv("DEPLOYED", "true")
	t.Setenv("POSTGRES_HOST", "db.local")
	t.Setenv("POSTGRES_PORT", "15432")
	t.Setenv("POSTGRES_DATABASE", "sample")
	t.Setenv("LOG_LEVEL", "debug")

	env, err := Load()
	require.NoError(t, err)

	assert.True(t, env.Deployed)
	assert.Equal(t, "db.local", env.PostgresHost)
	assert.Equal(t, int32(15432), env.PostgresPort)
	assert.Equal(t, "sample", env.PostgresDatabase)
	assert.Equal(t, "debug", env.LogLevel)
}
