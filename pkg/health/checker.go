package health

import (
	"context"
	"sync"
	"time"
)

type Status string

const (
	StatusUp   Status = "UP"
	StatusDown Status = "DOWN"
)

type HealthChecker struct {
	mu      sync.RWMutex
	checks  map[string]Check
	status  Status
}

type Check struct {
	Name     string
	Check    func(context.Context) error
	Timeout  time.Duration
	Required bool
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		checks: make(map[string]Check),
		status: StatusUp,
	}
}

func (h *HealthChecker) AddCheck(name string, check Check) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.checks[name] = check
}

func (h *HealthChecker) RunChecks(ctx context.Context) map[string]error {
	h.mu.Lock()
	defer h.mu.Unlock()

	results := make(map[string]error)
	for name, check := range h.checks {
		checkCtx, cancel := context.WithTimeout(ctx, check.Timeout)
		err := check.Check(checkCtx)
		cancel()

		results[name] = err
		if err != nil && check.Required {
			h.status = StatusDown
		}
	}

	return results
} 
