package model

import "time"

type UserRegisterRequest struct {
	Login    string `validate:"email"`
	Password []byte `validate:"required"`
}

type UserLoginRequest struct {
	Login    string `validate:"email"`
	Password []byte `validate:"required"`
}

type User struct {
	ID        string    `postgresql:"id"`
	Login     string    `postgresql:"login"`
	Password  []byte    `postgresql:"password"`
	CreatedAt time.Time `postgresql:"created_at"`
	UpdatedAt time.Time `postgresql:"updated_at"`
}
