package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURI   string
	GRPCServer    string
	TokenName     string
	TokenSecret   string
	TokenExpHours int
	ServerCert    string
	ServerKey     string
	ServerCa      string
}

func New() (*Config, error) {
	err := godotenv.Load("gophkeeper.env")
	if err != nil {
		return nil, fmt.Errorf("new load .env: %w", err)
	}

	config := &Config{}
	config.DatabaseURI = os.Getenv("DATABASE_URI")
	config.GRPCServer = os.Getenv("GRPC_SERVER")
	config.TokenName = os.Getenv("TOKEN_NAME")
	expHours, err := strconv.Atoi(os.Getenv("TOKEN_EXP_HOURS"))
	if err != nil {
		return nil, fmt.Errorf("new Atoi TOKEN_EXP_HOURS: %w", err)
	}
	config.TokenExpHours = expHours
	config.TokenSecret = os.Getenv("TOKEN_SECRET")
	config.ServerCert = os.Getenv("SERVER_CERT_FILE")
	config.ServerKey = os.Getenv("SERVER_KEY_FILE")
	config.ServerCa = os.Getenv("SERVER_CA_FILE")

	return config, nil
}
