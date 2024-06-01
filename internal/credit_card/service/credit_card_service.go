package service

import (
	"context"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

const creditCard = "credit_card"

type DataRepository interface {
	Insert(ctx context.Context, data model.Data) (model.Data, error)
	SelectAll(ctx context.Context, userID, dataType string) ([]model.Data, error)
}

type CryptService interface {
	Encrypt(key, data []byte) ([]byte, error)
	Decrypt(key, data []byte) ([]byte, error)
	GenerateKey() ([]byte, error)
}

type CreditCardService struct {
	repository DataRepository
	crypt      CryptService
	jwtManager *jwtmanager.JWTManager
	dataType   string
}

func New(repository DataRepository, crypt CryptService, jwtManager *jwtmanager.JWTManager) *CreditCardService {
	return &CreditCardService{
		repository: repository,
		crypt:      crypt,
		jwtManager: jwtManager,
		dataType:   creditCard,
	}
}
