package util

import (
	"os"
)

func GetEnvCloud(key string) string {
	return os.Getenv(key)
}

