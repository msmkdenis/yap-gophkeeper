package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (u *UserService) Register(ctx context.Context, req model.UserRegisterRequest) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("register create uuid: %w", err)
	}

	userToSave := model.User{
		ID:       id.String(),
		Login:    req.Login,
		Password: req.Password,
	}

	user, err := u.repository.Insert(ctx, userToSave)
	if err != nil {
		return "", fmt.Errorf("register insert user: %w", err)
	}

	token, err := u.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return "", fmt.Errorf("register build jwt: %w", err)
	}

	return token, nil
}
