package pbclient

import (
	"context"

	"google.golang.org/grpc/metadata"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
)

func (u *CreditCardPBClient) LoadCreditCard(token string, card model.CreditCardLoadRequest) ([]model.CreditCard, error) {
	req := &pb.GetCreditCardRequest{
		Number:        card.Number,
		Owner:         card.Owner,
		CvvCode:       card.CvvCode,
		PinCode:       card.PinCode,
		Metadata:      card.Metadata,
		ExpiresAfter:  card.ExpiresAfter,
		ExpiresBefore: card.ExpiresBefore,
	}

	md := metadata.New(map[string]string{"token": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := u.creditCardService.GetLoadCreditCard(ctx, req)
	if err != nil {
		return nil, err
	}

	cards := make([]model.CreditCard, 0, len(resp.Cards))
	for _, card := range resp.Cards {
		cards = append(cards, model.CreditCard{
			Number:    card.Number,
			OwnerName: card.OwnerName,
			ExpiresAt: card.ExpiresAt,
			CVV:       card.CvvCode,
			PinCode:   card.PinCode,
			MetaData:  card.Metadata,
		})
	}

	return cards, nil
}
