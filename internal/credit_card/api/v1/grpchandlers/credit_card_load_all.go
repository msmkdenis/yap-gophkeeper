package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"github.com/msmkdenis/yap-gophkeeper/internal/credit_card/specification"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
)

func (h *CreditCardHandler) GetLoadCreditCard(ctx context.Context, in *pb.GetCreditCardRequest) (*pb.GetCreditCardResponse, error) {
	spec, err := specification.NewCreditCardSpecification(in)
	if err != nil {
		slog.Error("Error while creating credit card specification: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	cards, err := h.creditCardService.LoadAllCreditCard(ctx, spec)
	if err != nil {
		slog.Error("Error while loading credit cards: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	creditCards := make([]*pb.CreditCard, 0, len(cards))
	for _, v := range cards {
		creditCards = append(creditCards, &pb.CreditCard{
			Id:        v.ID,
			OwnerId:   v.OwnerID,
			Number:    v.Number,
			OwnerName: v.OwnerName,
			ExpiresAt: v.ExpiresAt,
			CvvCode:   v.CVV,
			PinCode:   v.PinCode,
			Metadata:  v.MetaData,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetCreditCardResponse{Cards: creditCards}, nil
}
