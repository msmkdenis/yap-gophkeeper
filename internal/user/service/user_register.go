package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
)

func (u *UserService) Register(ctx context.Context, req model.UserRegisterRequest) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("register create uuid: %w", err)
	}

	userKey, err := u.crypt.GenerateKey()
	if err != nil {
		return "", fmt.Errorf("register generate key: %w", err)
	}

	cryptUserKey, err := u.crypt.EncryptWithMasterKey(userKey)
	if err != nil {
		return "", fmt.Errorf("register crypt key: %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("register hash: %w", err)
	}

	userToSave := model.User{
		ID:       id.String(),
		Login:    req.Login,
		Password: passwordHash,
		CryptKey: cryptUserKey,
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
