package monitor

import (
	"context"
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/contract"
	"github.com/wanglilind/qqq/pkg/contract/event"
)

// åçº¦çæ§å?
type ContractWatcher struct {
	eventEmitter  *event.EventEmitter
	metrics       map[string]*ContractMetrics
	alerts        []Alert
	mu            sync.RWMutex
	config        *Config
	stopChan      chan struct{}
}

type ContractMetrics struct {
	CallCount       uint64
	FailureCount    uint64
	AverageGasUsed  uint64
	LastCallTime    time.Time
	ErrorRate       float64
	MemoryUsage     uint64
	StorageUsage    uint64
}

type Alert struct {
	ContractAddress string
	Type           string
	Severity       string
	Message        string
	Timestamp      time.Time
	Data           map[string]interface{}
}

func NewContractWatcher(config *Config) *ContractWatcher {
	return &ContractWatcher{
		eventEmitter: event.NewEventEmitter(1000),
		metrics:     make(map[string]*ContractMetrics),
		alerts:      make([]Alert, 0),
		config:      config,
		stopChan:    make(chan struct{}),
	}
}

// å¯å¨çæ§
func (cw *ContractWatcher) Start(ctx context.Context) error {
	// è®¢éåçº¦äºä»¶
	if err := cw.subscribeToEvents(ctx); err != nil {
		return err
	}

	// å¯å¨ææ æ¶é
	go cw.collectMetrics(ctx)

	// å¯å¨åè­¦æ£æ?
	go cw.checkAlerts(ctx)

	return nil
}

// åæ­¢çæ§
func (cw *ContractWatcher) Stop() {
	close(cw.stopChan)
}

// è·ååçº¦ææ 
func (cw *ContractWatcher) GetMetrics(contractAddr string) (*ContractMetrics, error) {
	cw.mu.RLock()
	defer cw.mu.RUnlock()

	metrics, exists := cw.metrics[contractAddr]
	if !exists {
		return nil, ErrContractNotFound
	}
	return metrics, nil
}

// è®¢éåçº¦äºä»¶
func (cw *ContractWatcher) subscribeToEvents(ctx context.Context) error {
	eventCh, err := cw.eventEmitter.Subscribe(ctx, event.EventFilter{})
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case evt := <-eventCh:
				cw.handleEvent(evt)
			}
		}
	}()

	return nil
}

// æ¶éææ 
func (cw *ContractWatcher) collectMetrics(ctx context.Context) {
	ticker := time.NewTicker(cw.config.MetricsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-cw.stopChan:
			return
		case <-ticker.C:
			cw.updateMetrics()
		}
	}
}

// æ£æ¥åè­?
func (cw *ContractWatcher) checkAlerts(ctx context.Context) {
	ticker := time.NewTicker(cw.config.AlertCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-cw.stopChan:
			return
		case <-ticker.C:
			cw.checkAlertConditions()
		}
	}
}

// å¤çäºä»¶
func (cw *ContractWatcher) handleEvent(evt event.Event) {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	// æ´æ°ææ 
	metrics, exists := cw.metrics[evt.ContractAddress]
	if !exists {
		metrics = &ContractMetrics{}
		cw.metrics[evt.ContractAddress] = metrics
	}

	// æ´æ°ç»è®¡ä¿¡æ¯
	metrics.CallCount++
	metrics.LastCallTime = evt.Timestamp
	// æ´æ°å¶ä»ææ ...
}

// æ´æ°ææ 
func (cw *ContractWatcher) updateMetrics() {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	// æ´æ°ææåçº¦çææ 
	for _, metrics := range cw.metrics {
		metrics.ErrorRate = float64(metrics.FailureCount) / float64(metrics.CallCount)
		// æ´æ°å¶ä»ææ ...
	}
}

// æ£æ¥åè­¦æ¡ä»?
func (cw *ContractWatcher) checkAlertConditions() {
	cw.mu.Lock()
	defer cw.mu.Unlock()

	for addr, metrics := range cw.metrics {
		// æ£æ¥éè¯¯ç
		if metrics.ErrorRate > cw.config.ErrorRateThreshold {
			cw.createAlert(addr, "HIGH_ERROR_RATE", "CRITICAL")
		}

		// æ£æ¥èµæºä½¿ç?
		if metrics.MemoryUsage > cw.config.MemoryThreshold {
			cw.createAlert(addr, "HIGH_MEMORY_USAGE", "WARNING")
		}

		// æ£æ¥å¶ä»æ¡ä»?..
	}
}

// åå»ºåè­¦
func (cw *ContractWatcher) createAlert(contractAddr, alertType, severity string) {
	alert := Alert{
		ContractAddress: contractAddr,
		Type:           alertType,
		Severity:       severity,
		Timestamp:      time.Now(),
	}
	cw.alerts = append(cw.alerts, alert)
} 
