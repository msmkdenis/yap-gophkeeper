package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/internal/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
)

func (h *BinaryDataHandler) PostSaveBinaryData(ctx context.Context, in *pb.PostBinaryDataRequest) (*pb.PostBinaryDataResponse, error) {
	req := model.BinaryDataPostRequest{
		Name:      in.Name,
		Extension: in.Extension,
		Data:      in.Data,
		MetaData:  in.Metadata,
	}

	report, ok := h.validator.ValidatePostRequest(&req)
	if !ok {
		slog.Info("Unable to register user: invalid user request",
			slog.Any("violated_fields", report))
		return nil, lib.ProcessValidationError("invalid text_data post request", report)
	}

	binary, err := h.binaryDataService.SaveBinaryData(ctx, req)
	if err != nil {
		slog.Error("Unable to save binary_data", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.PostBinaryDataResponse{
		Id:        binary.ID,
		Name:      binary.Name,
		Extension: binary.Extension,
		Metadata:  binary.MetaData,
		CreatedAt: binary.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt: binary.UpdatedAt.Format(time.RFC3339Nano),
	}, nil
}
