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
	ID        string    `postgresql:"id" redis:"id"`
	Login     string    `postgresql:"login" redis:"login"`
	Password  []byte    `postgresql:"password" redis:"password"`
	CryptKey  []byte    `postgresql:"crypt_key" redis:"crypt_key"`
	CreatedAt time.Time `postgresql:"created_at" json:"created_at"`
	UpdatedAt time.Time `postgresql:"updated_at" json:"updated_at"`
}
