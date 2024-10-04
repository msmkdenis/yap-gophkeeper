package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
)

func (h *TextDataHandler) PostSaveTextData(ctx context.Context, in *pb.PostTextDataRequest) (*pb.PostTextDataResponse, error) {
	req := model.TextDataPostRequest{
		Text:     in.Text,
		MetaData: in.Metadata,
	}

	report, ok := h.validator.ValidatePostRequest(&req)
	if !ok {
		slog.Info("Unable to register user: invalid user request",
			slog.Any("violated_fields", report))
		return nil, lib.ProcessValidationError("invalid text_data post request", report)
	}

	text, err := h.textDataService.SaveTextData(ctx, req)
	if err != nil {
		slog.Error("Unable to save text_data", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.PostTextDataResponse{
		Id:        text.ID,
		Text:      text.Text,
		Metadata:  text.MetaData,
		CreatedAt: text.CreatedAt.Format(time.RFC3339),
		UpdatedAt: text.UpdatedAt.Format(time.RFC3339),
	}, nil
}
