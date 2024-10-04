package service //nolint:dupl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/lib"
)

func (u *UserProvider) LoginUser() {
	scanner := bufio.NewScanner(os.Stdin)
	red := color.New(color.FgRed).SprintFunc()

	var login, password string

	yellowBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(yellowBold("Input 'login password' to login:"))

	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("Input login as %s: ", yellow("'valid email'"))
	scanner.Scan()
	data := scanner.Text()
	login = strings.TrimSpace(data)
	if len(login) == 0 {
		fmt.Println(red("Login must not be empty please try again"))
		return
	}

	fmt.Printf("Input password as %s: ", yellow("'text'"))
	scanner.Scan()
	data = scanner.Text()
	password = strings.TrimSpace(data)
	if len(password) == 0 {
		fmt.Println(red("Password must not be empty please try again"))
		return
	}

	token, err := u.userService.LoginUser(login, password)
	if err != nil {
		lib.UnpackGRPCError(err)
	}

	u.state.SetToken(token)
	u.state.SetIsAuthorized(true)
	u.state.SetLogin(login)
}
