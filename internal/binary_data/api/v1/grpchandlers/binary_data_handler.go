package grpchandlers

import (
	"context"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func processValidationError(msg string, report map[string]string) error {
	st := status.New(codes.InvalidArgument, msg)
	details := make([]*errdetails.BadRequest_FieldViolation, 0, len(report))
	for field, message := range report {
		details = append(details, &errdetails.BadRequest_FieldViolation{
			Field:       field,
			Description: message,
		})
	}
	br := &errdetails.BadRequest{}
	br.FieldViolations = append(br.FieldViolations, details...)
	st, err := st.WithDetails(br)
	if err != nil {
		slog.Error("Internal error: failed to set details", slog.String("error", err.Error()))
		return status.Error(codes.Internal, "internal error")
	}
	return st.Err()
}
