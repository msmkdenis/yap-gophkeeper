package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/user/cerrors"
)

func (r *PostgresDataRepository) Insert(ctx context.Context, data model.Data) (model.Data, error) {
	rows, err := r.postgresPool.WriteDB.Query(ctx,
		`
			insert into gophkeeper.data
			    (id, owner_id, type, data, metadata, created_at, updated_at) 
			values
				($1, $2, $3, $4, $5, now(), now())
			returning id, owner_id, type, data, metadata, created_at, updated_at;
			`,
		data.ID,
		data.OwnerID,
		data.Type,
		data.Data,
		data.MetaData)
	if err != nil {
		return model.Data{}, fmt.Errorf("make query: %w", err)
	}

	savedData, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.Data])
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		return model.Data{}, fmt.Errorf("collect row: %w", cerrors.ErrUserAlreadyExists)
	}

	if err != nil {
		return model.Data{}, fmt.Errorf("collect row: %w", err)
	}

	return savedData, nil
}
