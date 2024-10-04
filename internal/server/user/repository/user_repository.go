package repository

import "github.com/msmkdenis/yap-gophkeeper/internal/server/storage/postgresql"

type PostgresUserRepository struct {
	postgresPool *postgresql.PostgresPool
}

func New(postgresPool *postgresql.PostgresPool) *PostgresUserRepository {
	return &PostgresUserRepository{postgresPool: postgresPool}
}
