package grpchandlers

import (
	"context"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
)

type UserService interface {
	Register(ctx context.Context, user model.UserRegisterRequest) (string, error)
	Login(ctx context.Context, user model.UserLoginRequest) (string, error)
}

type Validator interface {
	ValidateLoginRequest(req *model.UserLoginRequest) (map[string]string, bool)
	ValidateRegisterRequest(req *model.UserRegisterRequest) (map[string]string, bool)
}

type UserHandler struct {
	userService UserService
	pb.UnimplementedUserServiceServer
	validator Validator
}

func New(userService UserService, validator Validator) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator,
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
