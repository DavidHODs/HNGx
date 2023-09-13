package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv returns the value of the key as specified in the env file
func LoadEnv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error: could not load .env file")
	}

	keyValue, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("error: %s not found in env file", key)
	}

	return keyValue, nil
}

var DB *sql.DB

// InitDB initializes the database connection.
func InitDB() {
	connStr, err := LoadEnv("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error: could not connect to database ==> %v", err)
	}

	// defer DB.Close()
	// leaving database connection open instead of opening and closing a new connection for every database interraction

	err = DB.Ping()
	if err != nil {
		log.Fatalf("error: could not connect to database ==>  %v", err)
	} else {
		fmt.Println("database connection successful")
	}
}
