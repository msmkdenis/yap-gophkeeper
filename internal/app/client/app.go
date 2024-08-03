package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	pbclient "github.com/msmkdenis/yap-gophkeeper/internal/client/user/pbclient"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/user/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/config"
	"github.com/msmkdenis/yap-gophkeeper/internal/tlsconfig"
)

func Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.New()
	if err != nil {
		slog.Error("Failed to initialize config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	tls, err := tlsconfig.NewClientTLS(cfg.ClientCert, cfg.ClientKey, cfg.ClientCa)
	if err != nil {
		slog.Error("Failed to initialize tls", slog.String("error", err.Error()))
		os.Exit(1)
	}

	grpcClient, err := grpc.NewClient(cfg.GRPCServer, grpc.WithTransportCredentials(credentials.NewTLS(tls)))
	if err != nil {
		slog.Error("Failed to initialize grpcClient", slog.String("error", err.Error()))
		os.Exit(1)
	}

	userClient := pbclient.NewUserPBClient(user.NewUserServiceClient(grpcClient))
	userService := service.NewUserService(userClient)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите команду: ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "quit":
			fmt.Println("Программа завершает работу.")
			return
		case "register":
			fmt.Print("Введите данные для обработки: login, password: ")
			scanner.Scan()
			data := scanner.Text()
			fmt.Println(userService.RegisterUser(data))
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
