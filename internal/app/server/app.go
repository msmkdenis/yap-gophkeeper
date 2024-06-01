package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	tlsCreds "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/msmkdenis/yap-gophkeeper/internal/cache"
	"github.com/msmkdenis/yap-gophkeeper/internal/config"
	credentialsGRPCHandlers "github.com/msmkdenis/yap-gophkeeper/internal/credentials/api/v1/grpchandlers"
	credentialsValidation "github.com/msmkdenis/yap-gophkeeper/internal/credentials/api/v1/validation"
	credentialsService "github.com/msmkdenis/yap-gophkeeper/internal/credentials/service"
	creditCardGRPCHandlers "github.com/msmkdenis/yap-gophkeeper/internal/credit_card/api/v1/grpchandlers"
	creditCardValidation "github.com/msmkdenis/yap-gophkeeper/internal/credit_card/api/v1/validation"
	creditCardService "github.com/msmkdenis/yap-gophkeeper/internal/credit_card/service"
	repository "github.com/msmkdenis/yap-gophkeeper/internal/data_repository"
	"github.com/msmkdenis/yap-gophkeeper/internal/encryption"
	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/auth"
	"github.com/msmkdenis/yap-gophkeeper/internal/interceptors/keyextraction"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/storage/postgresql"
	textDataGRPCHandlers "github.com/msmkdenis/yap-gophkeeper/internal/text_data/api/v1/grpchandlers"
	textDataValidation "github.com/msmkdenis/yap-gophkeeper/internal/text_data/api/v1/validation"
	textDataService "github.com/msmkdenis/yap-gophkeeper/internal/text_data/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/tlsconfig"
	userGRPCHandlers "github.com/msmkdenis/yap-gophkeeper/internal/user/api/v1/grpchandlers"
	userValidation "github.com/msmkdenis/yap-gophkeeper/internal/user/api/v1/validation"
	userRepository "github.com/msmkdenis/yap-gophkeeper/internal/user/repository"
	userService "github.com/msmkdenis/yap-gophkeeper/internal/user/service"
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

	redis, err := cache.NewRedis(cfg.RedisURL, cfg.RedisPassword, cfg.RedisDB, cfg.RedisTimeoutSec)
	if err != nil {
		slog.Error("Failed to initialize redis", slog.String("error", err.Error()))
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

	userRepo := userRepository.New(postgresPool)
	dataRepo := repository.New(postgresPool)

	userServ := userService.New(userRepo, cryptService, jwtManager, redis)
	creditCardServ := creditCardService.New(dataRepo, cryptService, jwtManager)
	textDataServ := textDataService.New(dataRepo, cryptService, jwtManager)
	credentialServ := credentialsService.New(dataRepo, cryptService, jwtManager)

	tls, err := tlsconfig.NewTLS(cfg.ServerCert, cfg.ServerKey, cfg.ServerCa)
	if err != nil {
		slog.Error("Failed to initialize tls", slog.String("error", err.Error()))
		os.Exit(1)
	}

	validate := validator.New()
	creditCardValidator, err := creditCardValidation.New(validate)
	if err != nil {
		slog.Error("Failed to initialize credit card validator", slog.String("error", err.Error()))
		os.Exit(1)
	}
	textDataValidator, err := textDataValidation.New(validate)
	if err != nil {
		slog.Error("Failed to initialize text data validator", slog.String("error", err.Error()))
		os.Exit(1)
	}
	credentialsValidator, err := credentialsValidation.New(validate)
	if err != nil {
		slog.Error("Failed to initialize credentials validator", slog.String("error", err.Error()))
		os.Exit(1)
	}

	jwtAuth := auth.New(jwtManager)
	userKeyExtractor := keyextraction.New(cryptService, userRepo, redis)

	grpcServer := grpc.NewServer(grpc.Creds(tlsCreds.NewTLS(tls)),
		grpc.ChainUnaryInterceptor(jwtAuth.GRPCJWTAuth, userKeyExtractor.ExtractUserKey))

	user.RegisterUserServiceServer(grpcServer, userGRPCHandlers.New(userServ, userValidation.New(validate)))
	credit_card.RegisterCreditCardServiceServer(grpcServer, creditCardGRPCHandlers.New(creditCardServ, creditCardValidator))
	text_data.RegisterTextDataServiceServer(grpcServer, textDataGRPCHandlers.New(textDataServ, textDataValidator))
	credentials.RegisterCredentialsServiceServer(grpcServer, credentialsGRPCHandlers.New(credentialServ, credentialsValidator))

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", cfg.GRPCServer)
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
