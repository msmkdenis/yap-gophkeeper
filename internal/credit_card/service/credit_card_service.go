package service

import (
	"context"

	"github.com/msmkdenis/yap-gophkeeper/internal/cache"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

type CreditCardRepository interface {
	Insert(ctx context.Context, card model.CreditCard) (model.CreditCard, error)
	SelectAll(ctx context.Context, userID string) ([]model.CreditCard, error)
}

type CryptService interface {
	EncryptWithMasterKey(data []byte) ([]byte, error)
	DecryptWithMasterKey(data []byte) ([]byte, error)
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
	GenerateKey() ([]byte, error)
}

type CreditCardService struct {
	repository CreditCardRepository
	crypt      CryptService
	jwtManager *jwtmanager.JWTManager
	redis      *cache.Redis
}

func New(repository CreditCardRepository, crypt CryptService, jwtManager *jwtmanager.JWTManager, redis *cache.Redis) *CreditCardService {
	return &CreditCardService{
		repository: repository,
		crypt:      crypt,
		jwtManager: jwtManager,
		redis:      redis,
	}
}
