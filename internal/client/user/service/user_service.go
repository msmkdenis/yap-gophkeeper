package service

import (
	"fmt"
	"strings"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/lib"
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

func (u *UserProvider) RegisterUser(data string) {
	args := strings.Split(data, " ")
	if len(args) != 2 {
		fmt.Println("Please enter 'username password'")
		return
	}

	token, err := u.userService.RegisterUser(args[0], args[1])
	if err != nil {
		lib.UnpackGRPCError(err)
	}

	u.state.SetToken(token)
	u.state.SetIsAuthorized(true)
}

func (u *UserProvider) LoginUser(data string) {
	args := strings.Split(data, " ")
	if len(args) != 2 {
		fmt.Println("Please enter 'username password'")
		return
	}

	token, err := u.userService.LoginUser(args[0], args[1])
	if err != nil {
		lib.UnpackGRPCError(err)
	}

	u.state.SetToken(token)
	u.state.SetIsAuthorized(true)
}
