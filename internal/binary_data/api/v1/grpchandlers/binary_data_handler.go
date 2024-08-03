package grpchandlers

import (
	"context"

	"github.com/msmkdenis/yap-gophkeeper/internal/binary_data/specification"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
)

type BinaryDataService interface {
	SaveBinaryData(ctx context.Context, req model.BinaryDataPostRequest) (model.BinaryData, error)
	LoadAllBinaryData(ctx context.Context, spec specification.BinaryDataSpecification) ([]model.BinaryData, error)
}

type Validator interface {
	ValidatePostRequest(req *model.BinaryDataPostRequest) (map[string]string, bool)
}

type BinaryDataHandler struct {
	binaryDataService BinaryDataService
	pb.UnimplementedBinaryDataServiceServer
	validator Validator
}

func New(binaryDataService BinaryDataService, validator Validator) *BinaryDataHandler {
	return &BinaryDataHandler{
		binaryDataService: binaryDataService,
		validator:         validator,
	}
}
