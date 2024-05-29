package model

import "time"

type CreditCardPostRequest struct {
	Number    string `validate:"required,card_number"`
	OwnerName string `validate:"required,owner"`
	ExpiresAt string `validate:"expires_at"`
	CVV       string `validate:"cvv"`
	PinCode   string `validate:"pin"`
	MetaData  string
}

type CreditCardPostResponse struct {
	ID        string
	OwnerID   string
	Number    string
	OwnerName string
	ExpiresAt string
	CVV       string
	PinCode   string
	MetaData  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreditCard struct {
	ID        string    `db:"id" redis:"id"`
	OwnerID   string    `db:"owner_id" redis:"owner_id"`
	Number    []byte    `db:"number" redis:"number"`
	OwnerName []byte    `db:"owner_name" redis:"owner_name"`
	ExpiresAt []byte    `db:"expires_at" redis:"expires_at"`
	CVV       []byte    `db:"cvv_code" redis:"cvv_code"`
	PinCode   []byte    `db:"pin_code" redis:"pin_code"`
	CreatedAt time.Time `db:"created_at" redis:"created_at"`
	UpdatedAt time.Time `db:"updated_at" redis:"updated_at"`
	MetaData  string    `db:"meta_data" redis:"meta_data"`
}
