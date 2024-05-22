package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (r *PostgresUserRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	rows, err := r.postgresPool.DB.Query(ctx,
		`
			insert into gophkeeper.user
				(id, login, password, created_at, updated_at)
			values
				($1, $2, $3, NOW(), NOW())
			on conflict (login) do update set
				password = excluded.password,
				updated_at = now()
			returning id, login, password, created_at, updated_at;
			`,
		user.ID,
		user.Login,
		user.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("insert user: %w", err)
	}

	savedUser, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.User])
	if err != nil {
		return model.User{}, fmt.Errorf("insert user: %w", err)
	}

	return savedUser, nil
}
