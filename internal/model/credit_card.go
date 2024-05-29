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

type CreditCardCryptData struct {
	Number    string
	OwnerName string
	ExpiresAt string
	CVV       string
	PinCode   string
}

type CreditCard struct {
	ID        string    `db:"id" redis:"id"`
	OwnerID   string    `db:"owner_id" redis:"owner_id"`
	CryptData []byte    `db:"data" redis:"number"`
	CreatedAt time.Time `db:"created_at" redis:"created_at"`
	UpdatedAt time.Time `db:"updated_at" redis:"updated_at"`
	MetaData  string    `db:"meta_data" redis:"meta_data"`
}
