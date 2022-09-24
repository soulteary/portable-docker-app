package config

import (
	"os"
	"strings"
)

func SetDataFromEnv(envKey string, defaultValue string) string {
	dataFromEnv := strings.ToLower(strings.TrimSpace(os.Getenv(envKey)))
	if dataFromEnv != "" {
		return dataFromEnv
	}
	return defaultValue
}
