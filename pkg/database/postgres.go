package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type PostgresDB struct {
	pool *pgxpool.Pool
}

func NewPostgresDB(config *PostgresConfig) (*PostgresDB, error) {
	ctx := context.Background()
	
	// Build connection string
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=verify-full",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	// Configure connection pool
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %v", err)
	}

	// Set pool parameters
	poolConfig.MaxConns = 100
	poolConfig.MinConns = 10
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	// Create connection pool
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return &PostgresDB{pool: pool}, nil
} 
