package config

import (
	"os"
)

type Config struct {
	Environment string
	LogLevel    string
	Context     struct {
		Timeout string
	}
	Server struct {
		Protocol string
		Host     string
		Port     string
	}
	PostService struct {
		Host string
		Port string
	}

	PostImportService struct {
		Host string
		Port string
	}
}

func New() *Config {
	var config Config

	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "debug")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	// initialization server
	config.Server.Protocol = getEnv("SERVER_PROTOCOL", "http")
	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":9000")

	// initialization post service
	config.PostService.Host = getEnv("POST_SERVICE_HOST", "localhost")
	config.PostService.Port = getEnv("POST_SERVICE_PORT", ":5001")

	// initialization post import service
	config.PostImportService.Host = getEnv("POST_IMPORT_SERVICE_HOST", "localhost")
	config.PostImportService.Port = getEnv("POST_IMPORT_SERVICE_PORT", ":500")

	return &config
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}
