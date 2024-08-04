package pbclient

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
)

func (u *CreditCardPBClient) SaveCreditCard(token string, card model.CreditCardPostRequest) (model.CreditCard, error) {
	req := &pb.PostCreditCardRequest{
		Number:    card.Number,
		OwnerName: card.OwnerName,
		ExpiresAt: card.ExpiresAt,
		CvvCode:   card.CVV,
		PinCode:   card.PinCode,
		Metadata:  card.MetaData,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.creditCardService.PostSaveCreditCard(ctx, req)
	if err != nil {
		return model.CreditCard{}, err
	}

	creditCard := model.CreditCard{
		Number:    resp.Number,
		OwnerName: resp.OwnerName,
		ExpiresAt: resp.ExpiresAt,
		CVV:       resp.CvvCode,
		PinCode:   resp.PinCode,
		MetaData:  resp.Metadata,
	}

	return creditCard, nil
}
