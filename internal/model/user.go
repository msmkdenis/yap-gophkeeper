package model

import "time"

type UserRegisterRequest struct {
	Login    string `validate:"email"`
	Password string `validate:"required"`
}

type UserLoginRequest struct {
	Login    string `validate:"email"`
	Password string `validate:"required"`
}

type User struct {
	ID        string    `db:"id" redis:"id"`
	Login     string    `db:"login" redis:"login"`
	Password  []byte    `db:"password" redis:"password"`
	CryptKey  []byte    `db:"crypt_key" redis:"crypt_key"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
