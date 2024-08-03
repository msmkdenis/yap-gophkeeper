package grpchandlers

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/user/cerrors"
)

func (h *UserHandler) PostRegisterUser(ctx context.Context, in *pb.PostUserRegisterRequest) (*pb.PostUserRegisterResponse, error) {
	req := model.UserRegisterRequest{
		Login:    in.Login,
		Password: in.Password,
	}

	report, ok := h.validator.ValidateRegisterRequest(&req)
	if !ok {
		slog.Info("Unable to register user: invalid user request",
			slog.String("user_login", req.Login),
			slog.Any("violated_fields", report))
		return nil, processValidationError("invalid user request", report)
	}

	token, err := h.userService.Register(ctx, req)
	if errors.Is(err, cerrors.ErrUserAlreadyExists) {
		slog.Info("Unable to register user: user already exists",
			slog.String("user_login", req.Login))
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	if err != nil {
		slog.Error("Unable to register user",
			slog.String("user_login", req.Login),
			slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.PostUserRegisterResponse{Token: token}, nil
}
