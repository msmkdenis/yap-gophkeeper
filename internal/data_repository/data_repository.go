package repository

import "github.com/msmkdenis/yap-gophkeeper/internal/storage/postgresql"

type PostgresDataRepository struct {
	postgresPool *postgresql.PostgresPool
}

func New(postgresPool *postgresql.PostgresPool) *PostgresDataRepository {
	return &PostgresDataRepository{postgresPool: postgresPool}
}
