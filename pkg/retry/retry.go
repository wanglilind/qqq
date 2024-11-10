package retry

import (
	"context"
	"time"

	"github.com/wanglilind/qqq/pkg/errors"
)

type RetryConfig struct {
	MaxAttempts     int
	InitialInterval time.Duration
	MaxInterval     time.Duration
	Multiplier      float64
	Timeout         time.Duration
}

type RetryableFunc func() error

func WithRetry(ctx context.Context, fn RetryableFunc, config RetryConfig) error {
	var lastErr error
	interval := config.InitialInterval

	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), errors.ErrSystem, "context cancelled")
		default:
		}

		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
			// 指数退
			interval = time.Duration(float64(interval) * config.Multiplier)
			if interval > config.MaxInterval {
				interval = config.MaxInterval
			}
				time.Sleep(interval)
		}
	}

	return errors.Wrap(lastErr, errors.ErrSystem, "max retry attempts reached")
}

// 重试策略
type RetryStrategy interface {
	ShouldRetry(err error) bool
	NextInterval(attempt int) time.Duration
}

type ExponentialBackoff struct {
	Config RetryConfig
}

func (eb *ExponentialBackoff) ShouldRetry(err error) bool {
	// 判断错误是否可重
	switch err.(type) {
	case *errors.Error:
		code := err.(*errors.Error).Code
		return code == errors.ErrNetwork || code == errors.ErrDatabase
	default:
		return false
	}
}

func (eb *ExponentialBackoff) NextInterval(attempt int) time.Duration {
	interval := eb.Config.InitialInterval
	for i := 0; i < attempt; i++ {
		interval = time.Duration(float64(interval) * eb.Config.Multiplier)
		if interval > eb.Config.MaxInterval {
			interval = eb.Config.MaxInterval
			break
		}
	}
	return interval
} 
