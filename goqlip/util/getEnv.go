package util

import (
	"os"
	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load("./config/.env")
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

func GetEnvCloud(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}

