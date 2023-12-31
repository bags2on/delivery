package utils

import (
	"log"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("env variable %s: %s\n", key, value)
	}
	return value
}
