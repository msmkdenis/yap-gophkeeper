package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
)

func (h *CreditCardHandler) PostSaveCreditCard(ctx context.Context, in *pb.PostCreditCardRequest) (*pb.PostCreditCardResponse, error) {
	req := model.CreditCardPostRequest{
		Number:    in.Number,
		OwnerName: in.OwnerName,
		ExpiresAt: in.ExpiresAt,
		CVV:       in.CvvCode,
		PinCode:   in.PinCode,
		MetaData:  in.Metadata,
	}

	report, ok := h.validator.ValidatePostRequest(&req)
	if !ok {
		slog.Info("Unable to register user: invalid user request",
			slog.Any("violated_fields", report))
		return nil, lib.ProcessValidationError("invalid credit_card post request", report)
	}

	creditCard, err := h.creditCardService.SaveCreditCard(ctx, req)
	if err != nil {
		slog.Error("Unable to save credit_card", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.PostCreditCardResponse{
		Id:        creditCard.ID,
		OwnerId:   creditCard.OwnerID,
		Number:    creditCard.Number,
		OwnerName: creditCard.OwnerName,
		ExpiresAt: creditCard.ExpiresAt,
		CvvCode:   creditCard.CVV,
		PinCode:   creditCard.PinCode,
		Metadata:  creditCard.MetaData,
		CreatedAt: creditCard.CreatedAt.Format(time.RFC3339),
		UpdatedAt: creditCard.UpdatedAt.Format(time.RFC3339),
	}, nil
}
