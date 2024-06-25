package postgresql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPool struct {
	WriteDB *pgxpool.Pool
	ReadDB  *pgxpool.Pool
}

func NewPool(ctx context.Context, masterConn, slaveConn string) (*PostgresPool, error) {
	writeDBPool, err := pgxpool.New(ctx, masterConn)
	if err != nil {
		return nil, fmt.Errorf("newPostgresPool %w", err)
	}
	slog.Info("Successful masterConn", slog.String("database", writeDBPool.Config().ConnConfig.Database))

	err = writeDBPool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping postgresql %w", err)
	}
	slog.Info("Successful ping to master", slog.String("database", writeDBPool.Config().ConnConfig.Database))

	readDBPool, err := pgxpool.New(ctx, slaveConn)
	if err != nil {
		return nil, fmt.Errorf("newPostgresPool %w", err)
	}
	slog.Info("Successful slaveConn", slog.String("database", writeDBPool.Config().ConnConfig.Database))

	err = readDBPool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping postgresql %w", err)
	}
	slog.Info("Successful ping to slave", slog.String("database", readDBPool.Config().ConnConfig.Database))

	return &PostgresPool{
		WriteDB: writeDBPool,
		ReadDB:  readDBPool,
	}, nil
}
