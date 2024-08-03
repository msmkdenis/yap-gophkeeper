package grpchandlers

import (
	"context"

	"github.com/msmkdenis/yap-gophkeeper/internal/credit_card/specification"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
)

type CreditCardService interface {
	SaveCreditCard(ctx context.Context, req model.CreditCardPostRequest) (model.CreditCard, error)
	LoadAllCreditCard(ctx context.Context, spec specification.CreditCardSpecification) ([]model.CreditCard, error)
}

type Validator interface {
	ValidatePostRequest(req *model.CreditCardPostRequest) (map[string]string, bool)
}

type CreditCardHandler struct {
	creditCardService CreditCardService
	pb.UnimplementedCreditCardServiceServer
	validator Validator
}

func New(creditCardService CreditCardService, validator Validator) *CreditCardHandler {
	return &CreditCardHandler{
		creditCardService: creditCardService,
		validator:         validator,
	}
}
