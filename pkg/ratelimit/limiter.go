package ratelimit

import (
	"context"
	"math"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(r, b),
	}
}

func (l *RateLimiter) Allow() bool {
	return l.limiter.Allow()
}

func (l *RateLimiter) Wait(ctx context.Context) error {
	return l.limiter.Wait(ctx)
}

type TokenBucketLimiter struct {
	capacity  int64
	rate      float64
	tokens    float64
	lastCheck time.Time
	mu        sync.Mutex
}

func NewTokenBucketLimiter(capacity int64, rate float64) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity:  capacity,
		rate:      rate,
		tokens:    float64(capacity),
		lastCheck: time.Now(),
	}
}

func (t *TokenBucketLimiter) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(t.lastCheck).Seconds()
	t.tokens = math.Min(float64(t.capacity), t.tokens+elapsed*t.rate)
	t.lastCheck = now

	if t.tokens >= 1 {
		t.tokens--
		return true
	}
	return false
} 
