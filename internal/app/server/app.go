package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/msmkdenis/yap-gophkeeper/internal/encryption"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/msmkdenis/yap-gophkeeper/internal/config"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/storage/postgresql"
	"github.com/msmkdenis/yap-gophkeeper/internal/tlsconfig"
	"github.com/msmkdenis/yap-gophkeeper/internal/user/api/v1/grpchandlers"
	"github.com/msmkdenis/yap-gophkeeper/internal/user/api/v1/validation"
	"github.com/msmkdenis/yap-gophkeeper/internal/user/repository"
	"github.com/msmkdenis/yap-gophkeeper/internal/user/service"
	"github.com/msmkdenis/yap-gophkeeper/pkg/jwtmanager"
)

func Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		slog.Error("Failed to initialize config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	cryptService, err := encryption.New([]byte("master-key"))
	if err != nil {
		slog.Error("Failed to initialize crypt service", slog.String("error", err.Error()))
		os.Exit(1)
	}

	postgresPool, err := initPostgresPool(cfg.DatabaseURI)
	if err != nil {
		slog.Error("Failed to initialize postgres pool", slog.String("error", err.Error()))
		os.Exit(1)
	}

	jwtManager := jwtmanager.New(cfg.TokenName, cfg.TokenSecret, cfg.TokenExpHours)

	userRepository := repository.NewUserRepository(postgresPool)

	userService := service.New(userRepository, cryptService, jwtManager)

	tls, err := tlsconfig.NewTLS(cfg.ServerCert, cfg.ServerKey, cfg.ServerCa)
	if err != nil {
		slog.Error("Failed to initialize tls", slog.String("error", err.Error()))
		os.Exit(1)
	}

	validate := validator.New()

	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tls)))

	user.RegisterUserServiceServer(grpcServer, grpchandlers.NewUserHandler(userService, validation.New(validate)))

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost:3300")
	if err != nil {
		slog.Error("Unable to create listener", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := grpcServer.Serve(listener); err != nil {
		slog.Error("Unable to start gRPC server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func initPostgresPool(databaseURI string) (*postgresql.PostgresPool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	postgresPool, err := postgresql.NewPool(ctx, databaseURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	migrations, err := postgresql.NewMigrations(postgresPool)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize migrations: %w", err)
	}

	err = migrations.Up()
	if err != nil {
		return nil, fmt.Errorf("failed to up migrations: %w", err)
	}
	slog.Info("Connected to database")

	return postgresPool, nil
}
