package config

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/alexbednarczyk/crud-gopher/internal/environment/variables"

	"github.com/jackc/pgx/v5/pgxpool"
)

// BuildMigrationURLConnectionString creates a connection URL used for database migrations.
func BuildMigrationURLConnectionString(e variables.AllEnvironmentVariables) string {
	connURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(e.PostgresUser, e.PostgresPassword),
		Host:   fmt.Sprintf("%s:%d", e.PostgresHost, e.PostgresPort),
		Path:   "/" + e.PostgresDatabase,
	}

	query := url.Values{}
	if e.PostgresSSLMode != "" {
		query.Set("sslmode", e.PostgresSSLMode)
	}
	connURL.RawQuery = query.Encode()

	return connURL.String()
}

func buildPoolConfig(e variables.AllEnvironmentVariables) (*pgxpool.Config, error) {
	connString := BuildMigrationURLConnectionString(e)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse postgres connection string: %w", err)
	}

	if e.PostgresPoolMaxConnections > 0 {
		config.MaxConns = e.PostgresPoolMaxConnections
	}
	if e.PostgresPoolMinConnections > 0 {
		config.MinConns = e.PostgresPoolMinConnections
	}

	if e.PostgresPoolMaxConnectionLifetime != "" {
		if duration, err := time.ParseDuration(e.PostgresPoolMaxConnectionLifetime); err == nil {
			config.MaxConnLifetime = duration
		} else {
			return nil, fmt.Errorf("parse PostgresPoolMaxConnectionLifetime: %w", err)
		}
	}

	if e.PostgresPoolMaxConnectionIdleTime != "" {
		if duration, err := time.ParseDuration(e.PostgresPoolMaxConnectionIdleTime); err == nil {
			config.MaxConnIdleTime = duration
		} else {
			return nil, fmt.Errorf("parse PostgresPoolMaxConnectionIdleTime: %w", err)
		}
	}

	if e.PostgresPoolHealthCheckPeriod > 0 {
		config.HealthCheckPeriod = time.Duration(e.PostgresPoolHealthCheckPeriod) * time.Second
	}

	return config, nil
}

// GetPostgresDBConnectionPool creates a Postgres database connection pool based on the environment arguments passed in.
func GetPostgresDBConnectionPool(e variables.AllEnvironmentVariables) (*pgxpool.Pool, error) {
	poolConfig, err := buildPoolConfig(e)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	if !e.PostgresLazyConnect {
		if err := pool.Ping(ctx); err != nil {
			pool.Close()
			return nil, fmt.Errorf("ping database: %w", err)
		}
	}

	return pool, nil
}
