package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

// Factory creates Redis clients
type Factory struct {
	config *Config
}

// NewFactory creates a new Redis factory
func NewFactory(config *Config) *Factory {
	return &Factory{
		config: config,
	}
}

// CreateClient creates a new Redis client using the factory configuration
func (f *Factory) CreateClient() (Client, error) {
	client, err := NewClient(
		f.config.GetAddr(),
		f.config.Password,
		f.config.DB,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Redis client: %w", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping Redis: %w", err)
	}

	slog.Info("Redis client created successfully", 
		"host", f.config.Host, 
		"port", f.config.Port, 
		"db", f.config.DB)

	return client, nil
}

// CreateClientWithConfig creates a new Redis client with custom configuration
func (f *Factory) CreateClientWithConfig(addr, password string, db int) (Client, error) {
	client, err := NewClient(addr, password, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create Redis client with custom config: %w", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping Redis with custom config: %w", err)
	}

	return client, nil
}
