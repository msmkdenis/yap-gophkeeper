package grpchandlers

import (
	"context"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/credentials/specification"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
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
