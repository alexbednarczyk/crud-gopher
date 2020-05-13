package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexbednarczyk/crud-gopher/internal/environment/variables"

	"github.com/jackc/pgx/v4/pgxpool"
)

func buildURLConnectionString(e variables.AllEnvironmentVariables) string {

	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", e.PostgresUser, e.PostgresPassword, e.PostgresHost, e.PostgresPort, e.PostgresDatabase, e.PostgresSSLMode))

	if len(e.PostgresPoolMaxConnectionLifetime) > 0 {
		builder.WriteString(fmt.Sprintf("&pool_max_conn_lifetime=%s", e.PostgresPoolMaxConnectionLifetime))
	}

	if len(e.PostgresPoolMaxConnectionIdleTime) > 0 {
		builder.WriteString(fmt.Sprintf("&pool_max_conn_idle_time=%s", e.PostgresPoolMaxConnectionIdleTime))
	}

	if e.PostgresPoolMaxConnections > 0 {
		builder.WriteString(fmt.Sprintf("&pool_max_conns=%d", e.PostgresPoolMaxConnections))
	}

	if e.PostgresPoolMinConnections >= 0 {
		builder.WriteString(fmt.Sprintf("&pool_min_conns=%d", e.PostgresPoolMinConnections))
	}

	if e.PostgresPoolHealthCheckPeriod > 0 {
		builder.WriteString(fmt.Sprintf("&pool_health_check_period=%d", e.PostgresPoolHealthCheckPeriod))
	}

	return builder.String()
}

// GetPostgresDBConnectionPool creates a Postgres database connection pool based on the environment arguments passed in
func GetPostgresDBConnectionPool(e variables.AllEnvironmentVariables) (*pgxpool.Pool, error) {

	connString := buildURLConnectionString(e)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse connString (%v)", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	return pool, nil
}

// BuildMigrationURLConnectionString creates a connection URL used for database migrations
func BuildMigrationURLConnectionString(e variables.AllEnvironmentVariables) string {

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", e.PostgresUser, e.PostgresPassword, e.PostgresHost, e.PostgresPort, e.PostgresDatabase, e.PostgresSSLMode)

}
