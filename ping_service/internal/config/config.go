package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT string
	GRPC_PORT string
	GRPC_HOST string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("load env file: %v", err)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return nil, paramIsEmpty("HTTP_PORT")
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		return nil, paramIsEmpty("GRPC_PORT")
	}

	grpcHost := os.Getenv("GRPC_HOST")
	if grpcPort == "" {
		return nil, paramIsEmpty("GRPC_HOST")
	}

	return &Config{
		HTTP_PORT: httpPort,
		GRPC_PORT: grpcPort,
		GRPC_HOST: grpcHost,
	}, nil
}

func paramIsEmpty(param string) error {
	return fmt.Errorf("%s is empty", param)
}
