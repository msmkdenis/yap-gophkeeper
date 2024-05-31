package grpchandlers

import (
	"context"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/text_data/specification"
)

type TextDataService interface {
	SaveTextData(ctx context.Context, req model.TextDataPostRequest) (model.TextData, error)
	LoadAllTextData(ctx context.Context, spec specification.TextDataSpecification) ([]model.TextData, error)
}

type Validator interface {
	ValidatePostRequest(req *model.TextDataPostRequest) (map[string]string, bool)
}

type CreditCardHandler struct {
	textDataService TextDataService
	pb.UnimplementedTextDataServiceServer
	validator Validator
}

func New(textDataService TextDataService, validator Validator) *CreditCardHandler {
	return &CreditCardHandler{
		textDataService: textDataService,
		validator:       validator,
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
