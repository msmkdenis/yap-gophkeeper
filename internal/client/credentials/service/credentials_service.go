package service

import (
	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
)

type CredentialsService interface {
	SaveCredentials(token string, cred model.CredentialsPostRequest) (model.Credentials, error)
	LoadCredentials(token string, cred model.CredentialsLoadRequest) ([]model.Credentials, error)
}

type CredentialsProvider struct {
	credentialsService CredentialsService
	state              *state.ClientState
}

func NewCredentialsService(u CredentialsService, state *state.ClientState) *CredentialsProvider {
	return &CredentialsProvider{
		credentialsService: u,
		state:              state,
	}
}
