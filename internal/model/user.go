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
	ID        string    `postgresql:"id"`
	Login     string    `postgresql:"login"`
	Password  []byte    `postgresql:"password"`
	CryptKey  []byte    `postgresql:"crypt_key"`
	CreatedAt time.Time `postgresql:"created_at"`
	UpdatedAt time.Time `postgresql:"updated_at"`
}
