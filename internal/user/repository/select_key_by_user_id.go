package repository

import (
	"context"
	"fmt"
)

func (r *PostgresUserRepository) SelectKeyByID(ctx context.Context, userID string) ([]byte, error) {
	var userKey []byte
	err := r.postgresPool.WriteDB.QueryRow(ctx,
		`
			select 
				crypt_key
			from gophkeeper.user
			where id = $1;
			`,
		userID).Scan(&userKey)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	return userKey, nil
}
