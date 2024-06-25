package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/msmkdenis/yap-gophkeeper/internal/user/cerrors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (r *PostgresUserRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	rows, err := r.postgresPool.WriteDB.Query(ctx,
		`
			insert into gophkeeper.user
				(id, login, password, crypt_key, created_at, updated_at)
			values
				($1, $2, $3, $4, NOW(), NOW())
			returning id, login, password, crypt_key, created_at, updated_at;
			`,
		user.ID,
		user.Login,
		user.Password,
		user.CryptKey)
	if err != nil {
		return model.User{}, fmt.Errorf("make query: %w", err)
	}

	savedUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.User])
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		return model.User{}, fmt.Errorf("collect row: %w", cerrors.ErrUserAlreadyExists)
	}

	if err != nil {
		return model.User{}, fmt.Errorf("collect row: %w", err)
	}

	return savedUser, nil
}
