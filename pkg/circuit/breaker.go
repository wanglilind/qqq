package circuit

import (
	"context"
	"sync"
	"time"

	"github.com/sony/gobreaker"
)

type CircuitBreaker struct {
	name    string
	breaker *gobreaker.CircuitBreaker
}

func NewCircuitBreaker(name string, settings gobreaker.Settings) *CircuitBreaker {
	return &CircuitBreaker{
		name:    name,
		breaker: gobreaker.NewCircuitBreaker(settings),
	}
}

func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	return cb.breaker.Execute(req)
}

type BreakerManager struct {
	mu       sync.RWMutex
	breakers map[string]*CircuitBreaker
}

func NewBreakerManager() *BreakerManager {
	return &BreakerManager{
		breakers: make(map[string]*CircuitBreaker),
	}
}

func (bm *BreakerManager) GetBreaker(name string) *CircuitBreaker {
	bm.mu.RLock()
	if breaker, exists := bm.breakers[name]; exists {
		bm.mu.RUnlock()
		return breaker
	}
	bm.mu.RUnlock()

	bm.mu.Lock()
	defer bm.mu.Unlock()

	breaker := NewCircuitBreaker(name, gobreaker.Settings{
		Name:        name,
		MaxRequests: 3,
		Interval:    10 * time.Second,
		Timeout:     60 * time.Second,
	})
	bm.breakers[name] = breaker
	return breaker
} 
