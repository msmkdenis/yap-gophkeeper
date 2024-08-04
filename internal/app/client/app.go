package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/msmkdenis/yap-gophkeeper/internal/client/config"
	creaditcardb "github.com/msmkdenis/yap-gophkeeper/internal/client/credit_card/pbclient"
	creditcardservice "github.com/msmkdenis/yap-gophkeeper/internal/client/credit_card/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
	userpb "github.com/msmkdenis/yap-gophkeeper/internal/client/user/pbclient"
	userservice "github.com/msmkdenis/yap-gophkeeper/internal/client/user/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
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

	clientState := state.NewClientState()

	userClient := userpb.NewUserPBClient(user.NewUserServiceClient(grpcClient))
	userService := userservice.NewUserService(userClient, clientState)

	creditCardClient := creaditcardb.NewCreditCardPBClient(credit_card.NewCreditCardServiceClient(grpcClient))
	creditCardService := creditcardservice.NewUserService(creditCardClient, clientState)

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
			fmt.Print("Введите данные для регистрации 'login password':")
			scanner.Scan()
			data := scanner.Text()
			userService.RegisterUser(data)
		case "login":
			fmt.Print("Введите данные для авторизации 'login password':")
			scanner.Scan()
			data := scanner.Text()
			userService.LoginUser(data)
		case "save credit card":
			creditCardService.Save()
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
