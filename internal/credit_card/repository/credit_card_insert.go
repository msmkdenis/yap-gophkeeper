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

func (r *PostgresCreditCardRepository) Insert(ctx context.Context, card model.CreditCard) (model.CreditCard, error) {
	rows, err := r.postgresPool.DB.Query(ctx,
		`
			insert into gophkeeper.credit_card
			    (id, owner_id, data, created_at, updated_at, metadata) 
			values
				($1, $2, $3, now(), now(), $4)
			returning id, owner_id, data, created_at, updated_at, metadata;
			`,
		card.ID,
		card.OwnerID,
		card.CryptData,
		card.MetaData)
	if err != nil {
		return model.CreditCard{}, fmt.Errorf("make query: %w", err)
	}

	savedCreditCard, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[model.CreditCard])
	var e *pgconn.PgError
	if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
		return model.CreditCard{}, fmt.Errorf("collect row: %w", cerrors.ErrUserAlreadyExists)
	}

	if err != nil {
		return model.CreditCard{}, fmt.Errorf("collect row: %w", err)
	}

	return savedCreditCard, nil
}
