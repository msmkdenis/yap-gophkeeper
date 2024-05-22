package service

import (
	"context"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) (model.User, error)
}

type UserService struct {
	repository UserRepository
	jwtManager *jwtmanager.JWTManager
}

func New(repository UserRepository, jwtManager *jwtmanager.JWTManager) *UserService {
	return &UserService{
		repository: repository,
		jwtManager: jwtManager,
	}
}
