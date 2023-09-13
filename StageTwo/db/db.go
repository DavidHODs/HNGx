package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv returns the value of the key as specified in the env file
func LoadEnv(key string) (string, error) {
	err := godotenv.Load("stageTwo.env")
	if err != nil {
		return "", errors.New("error: could not load .env file")
	}

	keyValue, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("error: %s not found in env file", key)
	}

	return keyValue, nil
}
