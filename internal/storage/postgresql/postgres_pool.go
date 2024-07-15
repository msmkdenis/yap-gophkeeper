package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	DB *pgxpool.Pool
}

func NewPool(ctx context.Context, connection string) (*PostgresPool, error) {
	dbPool, err := pgxpool.New(ctx, connection)
	if err != nil {
		return nil, fmt.Errorf("newPostgresPool %w", err)
	}
	slog.Info("Successful connection", slog.String("database", dbPool.Config().ConnConfig.Database))

	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping postgresql %w", err)
	}
	slog.Info("Successful ping", slog.String("database", dbPool.Config().ConnConfig.Database))

	return &PostgresPool{DB: dbPool}, nil
}
