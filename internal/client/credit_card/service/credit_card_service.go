package service

import (
	"bufio"
	"fmt"
	"os"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
)

type CreditCardService interface {
	SaveCreditCard(token string, card model.CreditCardPostRequest) (model.CreditCard, error)
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

func (p *CreditCardProvider) Save() {
	if !p.state.IsAuthorized() {
		fmt.Println("You are not authorized, please use 'login' or 'register'")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	req := model.CreditCardPostRequest{}
	fmt.Println("Input credit card data 'number owner expires cvv pin metadata':")

	fmt.Print("Input number in format 'dddd dddd dddd dddd':")
	scanner.Scan()
	data := scanner.Text()
	req.Number = data

	fmt.Print("Input owner in format 'name surname':")
	scanner.Scan()
	data = scanner.Text()
	req.OwnerName = data

	fmt.Print("Input expiry date in format 'dd-mm-yyyy':")
	scanner.Scan()
	data = scanner.Text()
	req.ExpiresAt = data

	fmt.Print("Input pin on format 'dddd':")
	scanner.Scan()
	data = scanner.Text()
	req.PinCode = data

	fmt.Print("Input cvv in format 'ddd':")
	scanner.Scan()
	data = scanner.Text()
	req.CVV = data

	fmt.Print("Input metadata:")
	scanner.Scan()
	data = scanner.Text()
	req.MetaData = data

	card, err := p.creditCardService.SaveCreditCard(p.state.GetToken(), req)
	if err != nil {
		lib.UnpackGRPCError(err)
		return
	}

	fmt.Println(card)
}
