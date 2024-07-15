package grpchandlers

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/internal/binary_data/specification"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
)

func (h *BinaryDataHandler) GetLoadBinaryData(ctx context.Context, in *pb.GetBinaryDataRequest) (*pb.GetBinaryDataResponse, error) {
	spec, err := specification.NewTextDataSpecification(in)
	if err != nil {
		slog.Error("Error while creating text data specification: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	binaries, err := h.binaryDataService.LoadAllBinaryData(ctx, spec)
	if err != nil {
		slog.Error("Error while loading binary data: ", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	binaryData := make([]*pb.BinaryData, 0, len(binaries))
	for _, v := range binaries {
		binaryData = append(binaryData, &pb.BinaryData{
			Id:        v.ID,
			OwnerId:   v.OwnerID,
			Data:      v.Data,
			Name:      v.Name,
			Extension: v.Extension,
			Metadata:  v.MetaData,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &pb.GetBinaryDataResponse{Binaries: binaryData}, nil
}
