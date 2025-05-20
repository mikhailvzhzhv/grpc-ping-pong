package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("load env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, paramIsEmpty("PORT")
	}

	return &Config{
		PORT: port,
	}, nil
}

func paramIsEmpty(param string) error {
	return fmt.Errorf("%s is empty", param)
}
