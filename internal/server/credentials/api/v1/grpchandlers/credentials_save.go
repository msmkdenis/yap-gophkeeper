package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
)

func (h *CredentialsHandler) PostSaveCredentials(ctx context.Context, in *pb.PostCredentialsRequest) (*pb.PostCredentialsResponse, error) {
	req := model.CredentialsPostRequest{
		Login:    in.Login,
		Password: in.Password,
		MetaData: in.Metadata,
	}

	report, ok := h.validator.ValidatePostRequest(&req)
	if !ok {
		slog.Info("Unable to register user: invalid credentials request",
			slog.Any("violated_fields", report))
		return nil, lib.ProcessValidationError("invalid credentials post request", report)
	}

	cred, err := h.credentialsService.SaveCredentials(ctx, req)
	if err != nil {
		slog.Error("Unable to save credentials", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &pb.PostCredentialsResponse{
		Id:        cred.ID,
		Login:     cred.Login,
		Password:  cred.Password,
		Metadata:  cred.MetaData,
		CreatedAt: cred.CreatedAt.Format(time.RFC3339),
		UpdatedAt: cred.UpdatedAt.Format(time.RFC3339),
	}, nil
}
