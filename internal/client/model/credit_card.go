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

type CreditCard struct {
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
