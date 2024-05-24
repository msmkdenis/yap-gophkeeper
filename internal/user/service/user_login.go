package service

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (u *UserService) Login(ctx context.Context, req model.UserLoginRequest) (string, error) {
	user, err := u.repository.SelectByLogin(ctx, req.Login)
	if err != nil {
		return "", fmt.Errorf("login SelectByLogin: %w", err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("password compare hash: %w", err)
	}

	token, err := u.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return "", fmt.Errorf("register build jwt: %w", err)
	}

	return token, nil
}
