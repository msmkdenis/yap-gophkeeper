package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (r *PostgresCreditCardRepository) SelectAll(ctx context.Context, userID string) ([]model.CreditCard, error) {
	rows, err := r.postgresPool.DB.Query(ctx,
		`
			select
			    id, owner_id, data, created_at, updated_at, metadata 
			from gophkeeper.credit_card
			where owner_id = $1;
			`,
		userID)
	if err != nil {
		return nil, fmt.Errorf("make query: %w", err)
	}

	cards, err := pgx.CollectRows(rows, pgx.RowToStructByPos[model.CreditCard])
	if err != nil {
		return nil, fmt.Errorf("collect row: %w", err)
	}

	return cards, nil
}
