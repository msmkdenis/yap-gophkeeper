package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/internal/credentials/specification"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
)

func (h *CredentialsHandler) GetLoadCredentials(ctx context.Context, in *pb.GetCredentialsRequest) (*pb.GetCredentialsResponse, error) {
	spec, err := specification.NewCredentialsSpecification(in)
	if err != nil {
		slog.Error("Error while creating credentials specification: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	creds, err := h.credentialsService.LoadAllCredentials(ctx, spec)
	if err != nil {
		slog.Error("Error while loading credentials: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	credsData := make([]*pb.Credentials, 0, len(creds))
	for _, v := range creds {
		credsData = append(credsData, &pb.Credentials{
			Id:        v.ID,
			OwnerId:   v.OwnerID,
			Login:     v.Login,
			Password:  v.Password,
			Metadata:  v.MetaData,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetCredentialsResponse{Creds: credsData}, nil
}
