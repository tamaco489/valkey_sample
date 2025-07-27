package health

import (
	"context"
	"time"
)

// Checker represents a health check interface
type Checker interface {
	Check(ctx context.Context) error
}

// Service represents health check service
type Service struct {
	checkers map[string]Checker
}

// NewService creates a new health check service
func NewService() *Service {
	return &Service{
		checkers: make(map[string]Checker),
	}
}

// AddChecker adds a health checker
func (s *Service) AddChecker(name string, checker Checker) {
	s.checkers[name] = checker
}

// Check performs all health checks
func (s *Service) Check(ctx context.Context) map[string]error {
	results := make(map[string]error)

	for name, checker := range s.checkers {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		results[name] = checker.Check(ctx)
	}

	return results
}
