package repository

import "github.com/msmkdenis/yap-gophkeeper/internal/storage/postgresql"

type PostgresCreditCardRepository struct {
	postgresPool *postgresql.PostgresPool
}

func New(postgresPool *postgresql.PostgresPool) *PostgresCreditCardRepository {
	return &PostgresCreditCardRepository{postgresPool: postgresPool}
}
