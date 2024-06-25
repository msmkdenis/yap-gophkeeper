package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MasterDatabaseURI string
	SlaveDatabaseURI  string
	GRPCServer        string
	TokenName         string
	TokenSecret       string
	TokenExpHours     int
	ServerCert        string
	ServerKey         string
	ServerCa          string
	RedisURL          string
	RedisPassword     string
	RedisDB           int
	RedisTimeoutSec   int
}

func New() (*Config, error) {
	err := godotenv.Load("gophkeeper.env")
	if err != nil {
		return nil, fmt.Errorf("new load .env: %w", err)
	}

	config := &Config{}
	config.MasterDatabaseURI = os.Getenv("MASTER_DATABASE_URI")
	config.SlaveDatabaseURI = os.Getenv("SLAVE_DATABASE_URI")
	config.GRPCServer = os.Getenv("GRPC_SERVER")
	config.TokenName = os.Getenv("TOKEN_NAME")
	expHours, err := strconv.Atoi(os.Getenv("TOKEN_EXP_HOURS"))
	if err != nil {
		return nil, fmt.Errorf("atoi TOKEN_EXP_HOURS: %w", err)
	}
	config.TokenExpHours = expHours
	config.TokenSecret = os.Getenv("TOKEN_SECRET")
	config.ServerCert = os.Getenv("SERVER_CERT_FILE")
	config.ServerKey = os.Getenv("SERVER_KEY_FILE")
	config.ServerCa = os.Getenv("SERVER_CA_FILE")

	config.RedisURL = os.Getenv("REDIS_URL")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, fmt.Errorf("atoi REDIS_DB: %w", err)
	}
	config.RedisDB = db

	config.RedisTimeoutSec, err = strconv.Atoi(os.Getenv("REDIS_TIMEOUT_SEC"))
	if err != nil {
		return nil, fmt.Errorf("atoi REDIS_TIMEOUT_SEC: %w", err)
	}

	return config, nil
}
