package platform

import (
	"context"
	"errors"
	"fmt"

	"github.com/esferachill/mini/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	DatabaseURL string

	// Pool sizing
	// MinConns int32
	// MaxConns int32

	// Lifecycle
	// MaxConnLifetime   time.Duration
	// MaxConnIdleTime   time.Duration
	// HealthCheckPeriod time.Duration

	// Timeouts
	// ConnectTimeout time.Duration
	// QueryTimeout   time.Duration

	// Displays in pg_stat_activity
	// AppName string
}

type DatabaseClient struct {
	pool    *pgxpool.Pool
	Queries *repo.Queries
	config  DatabaseConfig
}

func NewDatabaseClient(config DatabaseConfig) (*DatabaseClient, error) {
	if config.DatabaseURL == "" {
		return nil, errors.New("DatabaseURL is required")
	}
	conn, err := pgxpool.ParseConfig(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("Parse config: %w", err)
	}

	// Pool sizing
	// if config.MinConns > 0 {
	// 	pc.MinConns = config.MinConns
	// }
	// if config.MaxConns > 0 {
	// 	pc.MaxConns = config.MaxConns
	// }

	// Lifetimes & health
	// if config.MaxConnLifetime > 0 {
	// 	pc.MaxConnLifetime = config.MaxConnLifetime
	// }
	// if config.MaxConnIdleTime > 0 {
	// 	pc.MaxConnIdleTime = config.MaxConnIdleTime
	// }
	// if config.HealthCheckPeriod > 0 {
	// 	pc.HealthCheckPeriod = config.HealthCheckPeriod
	// }

	// AppName is useful in monitoring (pg_stat_activity)
	// if config.AppName != "" {
	// 	pc.ConnConfig.RuntimeParams["application_name"] = config.AppName
	// }

	// Example of AfterConnect hook (fast, idempotent operations only)
	// pc.AfterConnect = func(ctx context.Context, c *pgx.Conn) error {
	// 	// Ensure UTC timestamps; tweak to your needs.
	// 	_, err := c.Exec(ctx, `SET TIME ZONE 'UTC'`)
	// 	return err
	// }

	// Use a bounded timeout on the initial connect
	// connectCtx := ctx
	// if config.ConnectTimeout > 0 {
	// 	var cancel context.CancelFunc
	// 	connectCtx, cancel = context.WithTimeout(ctx, config.ConnectTimeout)
	// 	defer cancel()
	// }

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("Connect: %w", err)
	}

	queries := repo.New(pool)

	client := &DatabaseClient{
		pool:    pool,
		Queries: queries,
		config:  config,
	}

	// Readiness ping (ensures creds/network correct). Short timeout.
	// pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	// if err := pool.Ping(pingCtx); err != nil {
	// 	_ = pool.Close()
	// 	return nil, fmt.Errorf("pg: ping: %w", err)
	// }

	return client, nil
}
