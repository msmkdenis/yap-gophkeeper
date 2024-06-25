package grpchandlers

import (
	"context"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/internal/credentials/specification"
	"github.com/msmkdenis/yap-gophkeeper/internal/model"
	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
)

type CredentialsService interface {
	SaveCredentials(ctx context.Context, req model.CredentialsPostRequest) (model.Credentials, error)
	LoadAllCredentials(ctx context.Context, spec specification.CredentialsSpecification) ([]model.Credentials, error)
}

type Validator interface {
	ValidatePostRequest(req *model.CredentialsPostRequest) (map[string]string, bool)
}

type CredentialsHandler struct {
	credentialsService CredentialsService
	pb.UnimplementedCredentialsServiceServer
	validator Validator
}

func New(textDataService CredentialsService, validator Validator) *CredentialsHandler {
	return &CredentialsHandler{
		credentialsService: textDataService,
		validator:          validator,
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
