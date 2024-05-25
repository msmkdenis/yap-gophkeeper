package service

import (
	"context"
	"fmt"
	"time"

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
		return "", fmt.Errorf("login password compare hash: %w", err)
	}

	token, err := u.jwtManager.BuildJWTString(user.ID)
	if err != nil {
		return "", fmt.Errorf("login build jwt: %w", err)
	}

	st := u.redis.Client.Set(ctx, user.ID, user.CryptKey, 24*time.Hour)
	if st.Err() != nil {
		return "", fmt.Errorf("login redis set: %w", st.Err())
	}

	return token, nil
}
