package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// Config represents application configuration
type Config struct {
	Server ServerConfig
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Load .env file if it exists
	loadEnvFile()

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
	}
}

// loadEnvFile loads environment variables from .env file
func loadEnvFile() {
	// Try to load .env file from current directory
	if err := godotenv.Load(); err != nil {
		// Try to load from project root
		projectRoot := findProjectRoot()
		if projectRoot != "" {
			envPath := filepath.Join(projectRoot, ".env")
			if err := godotenv.Load(envPath); err != nil {
				// .env file not found, use default values
				return
			}
		}
	}
}

// findProjectRoot finds the project root directory by looking for go.mod
func findProjectRoot() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			break
		}
		currentDir = parent
	}

	return ""
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return strings.TrimSpace(value)
	}
	return fallback
}
