package grpchandlers

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/lib"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/model"
	"github.com/msmkdenis/yap-gophkeeper/internal/server/user/cerrors"
)

func (h *UserHandler) PostLoginUser(ctx context.Context, in *pb.PostUserLoginRequest) (*pb.PostUserLoginResponse, error) {
	req := model.UserLoginRequest{
		Login:    in.Login,
		Password: in.Password,
	}

	report, ok := h.validator.ValidateLoginRequest(&req)
	if !ok {
		slog.Info("Unable to login user: invalid user request",
			slog.String("user_login", req.Login),
			slog.Any("violated_fields", report))
		return nil, lib.ProcessValidationError("invalid user request", report)
	}

	token, err := h.userService.Login(ctx, req)
	if errors.Is(err, cerrors.ErrUserNotFound) {
		slog.Info("Unable to login user: user not found",
			slog.String("user_login", req.Login))
		return nil, status.Error(codes.NotFound, "user with this login not found")
	}

	if errors.Is(err, cerrors.ErrInvalidPassword) {
		slog.Info("Unable to login user: invalid password",
			slog.String("user_login", req.Login))
		return nil, status.Error(codes.Unauthenticated, "incorrect password")
	}

	if err != nil {
		slog.Error("Unable to login user: invalid user request",
			slog.String("user_login", req.Login),
			slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &pb.PostUserLoginResponse{Token: token}, nil
}
