package grpchandlers

import (
	"context"

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

type TextDataHandler struct {
	textDataService TextDataService
	pb.UnimplementedTextDataServiceServer
	validator Validator
}

func New(textDataService TextDataService, validator Validator) *TextDataHandler {
	return &TextDataHandler{
		textDataService: textDataService,
		validator:       validator,
	}
}
