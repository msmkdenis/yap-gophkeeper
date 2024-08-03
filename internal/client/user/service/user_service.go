package service

import "strings"

type UserService interface {
	RegisterUser(login, password string) (string, error)
	LoginUser(login, password string) (string, error)
}

type UserProvider struct {
	userService UserService
}

func NewUserService(u UserService) *UserProvider {
	return &UserProvider{
		userService: u,
	}
}

func (u *UserProvider) RegisterUser(data string) string {
	args := strings.Fields(data)
	token, err := u.userService.RegisterUser(args[0], args[1])
	if err != nil {
		return err.Error()
	}

	return token
}

func (u *UserProvider) LoginUser(data string) string {
	args := strings.Fields(data)
	token, err := u.userService.LoginUser(args[0], args[1])
	if err != nil {
		return err.Error()
	}

	return token
}
