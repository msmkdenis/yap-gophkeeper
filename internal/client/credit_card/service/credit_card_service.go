package service

import (
	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
)

type CreditCardService interface {
	SaveCreditCard(token string, card model.CreditCardPostRequest) (model.CreditCard, error)
	LoadCreditCard(token string, card model.CreditCardLoadRequest) ([]model.CreditCard, error)
}

type CreditCardProvider struct {
	creditCardService CreditCardService
	state             *state.ClientState
}

func NewUserService(u CreditCardService, state *state.ClientState) *CreditCardProvider {
	return &CreditCardProvider{
		creditCardService: u,
		state:             state,
	}
}
