package config

import (
	"os"
)

// Config represents application configuration
type Config struct {
	Server ServerConfig
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
	}
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
