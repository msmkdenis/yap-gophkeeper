package auth

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

var authMandatoryMethods = map[string]struct{}{
	"/proto.CreditCardService/PostSaveCreditCard": {},
	"/proto.CreditCardService/GetLoadCreditCard":  {},
}

type UserIDContextKey string

type JWTAuth struct {
	jwtManager *jwtmanager.JWTManager
}

func New(jwtManager *jwtmanager.JWTManager) *JWTAuth {
	return &JWTAuth{jwtManager: jwtManager}
}

// GRPCJWTAuth checks token from gRPC metadata and sets userID in the context. otherwise returns 401.
func (j *JWTAuth) GRPCJWTAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if _, ok := authMandatoryMethods[info.FullMethod]; !ok {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.Info("Authentication failed: missing metadata")
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	c := md.Get(j.jwtManager.TokenName)
	if len(c) < 1 {
		slog.Info("Authentication failed: token not found")
		return nil, status.Errorf(codes.Unauthenticated, "token not found")
	}

	userID, err := j.jwtManager.GetUserID(c[0])
	if err != nil {
		slog.Info("Authentication failed: unable to get userID from token", slog.String("error", err.Error()))
		return nil, status.Errorf(codes.Unauthenticated, "authentification by UserID failed")
	}

	ctx = context.WithValue(ctx, UserIDContextKey("userID"), userID)
	return handler(ctx, req)
}
