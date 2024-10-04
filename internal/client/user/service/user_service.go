package service

import (
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
)

type UserService interface {
	RegisterUser(login, password string) (string, error)
	LoginUser(login, password string) (string, error)
}

type UserProvider struct {
	userService UserService
	state       *state.ClientState
}

func NewUserService(u UserService, state *state.ClientState) *UserProvider {
	return &UserProvider{
		userService: u,
		state:       state,
	}
}
