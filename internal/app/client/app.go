package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	binarypb "github.com/msmkdenis/yap-gophkeeper/internal/client/binary_data/pbclient"
	binaryservice "github.com/msmkdenis/yap-gophkeeper/internal/client/binary_data/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/config"
	credentialspb "github.com/msmkdenis/yap-gophkeeper/internal/client/credentials/pbclient"
	credentialsservice "github.com/msmkdenis/yap-gophkeeper/internal/client/credentials/service"
	creditcardpb "github.com/msmkdenis/yap-gophkeeper/internal/client/credit_card/pbclient"
	creditcardservice "github.com/msmkdenis/yap-gophkeeper/internal/client/credit_card/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/client/state"
	textdatapb "github.com/msmkdenis/yap-gophkeeper/internal/client/text_data/pbclient"
	textdataservice "github.com/msmkdenis/yap-gophkeeper/internal/client/text_data/service"
	userpb "github.com/msmkdenis/yap-gophkeeper/internal/client/user/pbclient"
	userservice "github.com/msmkdenis/yap-gophkeeper/internal/client/user/service"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/binary_data"
	credGrpc "github.com/msmkdenis/yap-gophkeeper/internal/proto/credentials"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/credit_card"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/text_data"
	"github.com/msmkdenis/yap-gophkeeper/internal/proto/user"
	"github.com/msmkdenis/yap-gophkeeper/internal/tlsconfig"
)

func Run() { //nolint:cyclop
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

	creditCardClient := creditcardpb.NewCreditCardPBClient(credit_card.NewCreditCardServiceClient(grpcClient))
	creditCardService := creditcardservice.NewUserService(creditCardClient, clientState)

	textDataClient := textdatapb.NewCreditCardPBClient(text_data.NewTextDataServiceClient(grpcClient))
	textDataService := textdataservice.NewTextDataService(textDataClient, clientState)

	credentialsClient := credentialspb.NewCredentialsPBClient(credGrpc.NewCredentialsServiceClient(grpcClient))
	credentialsService := credentialsservice.NewCredentialsService(credentialsClient, clientState)

	binaryClient := binarypb.NewBinaryDataPBClient(binary_data.NewBinaryDataServiceClient(grpcClient))
	binaryService := binaryservice.NewBinaryDataService(binaryClient, clientState)

	scanner := bufio.NewScanner(os.Stdin)

	blue := color.New(color.FgBlue).SprintFunc()

	for {
		if clientState.IsAuthorized() {
			fmt.Printf("You are authorized as %s\n", blue(clientState.GetLogin()))
		} else {
			fmt.Printf("You are not authorized, please login or register\n")
		}

		if clientState.GetDirPath() == "" {
			fmt.Printf("Working directory is not set \n")
		} else {
			fmt.Printf("Working directory is set to %s\n", blue(clientState.GetDirPath()))
		}

		fmt.Println("Input command number to proceed")
		fmt.Println("[0] - quit")
		fmt.Println("[1] - login")
		fmt.Println("[2] - register")
		fmt.Println("[3] - save credit card")
		fmt.Println("[4] - load credit cards")
		fmt.Println("[5] - save text data")
		fmt.Println("[6] - load text data")
		fmt.Println("[7] - save credentials")
		fmt.Println("[8] - load credentials")
		fmt.Println("[9] - save binary file")
		fmt.Println("[10] - load binary files")
		fmt.Println("[11] - set working directory")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "0":
			fmt.Println("Application shutdown.")
			return
		case "2":
			userService.RegisterUser()
		case "1":
			userService.LoginUser()
		case "3":
			creditCardService.Save()
		case "4":
			creditCardService.Load()
		case "5":
			textDataService.Save()
		case "6":
			textDataService.Load()
		case "7":
			credentialsService.Save()
		case "8":
			credentialsService.Load()
		case "9":
			binaryService.Save()
		case "10":
			binaryService.Load()
		case "11":
			clientState.SetWorkingDirectory()
		default:
			fmt.Println("Unknown command, please try again")
		}
	}
}
