package alert

import (
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/config"
)

type AlertLevel string

const (
	InfoLevel     AlertLevel = "INFO"
	WarningLevel  AlertLevel = "WARNING"
	ErrorLevel    AlertLevel = "ERROR"
	CriticalLevel AlertLevel = "CRITICAL"
)

type Alert struct {
	ID          string
	Level       AlertLevel
	Source      string
	Message     string
	Timestamp   time.Time
	Status      string
	Metadata    map[string]interface{}
}

type Manager struct {
	config      *config.Config
	mu          sync.RWMutex
	alerts      []Alert
	handlers    map[AlertLevel][]AlertHandler
	stopChan    chan struct{}
}

type AlertHandler interface {
	Handle(alert Alert) error
}

func NewManager(cfg *config.Config) *Manager {
	m := &Manager{
		config:   cfg,
		alerts:   make([]Alert, 0),
		handlers: make(map[AlertLevel][]AlertHandler),
		stopChan: make(chan struct{}),
	}

	// æ³¨åé»è®¤åè­¦å¤çå?
	m.registerDefaultHandlers()
	return m
}

func (m *Manager) Start() error {
	go m.processAlerts()
	return nil
}

func (m *Manager) Stop() {
	close(m.stopChan)
}

func (m *Manager) TriggerAlert(level AlertLevel, source, message string, metadata map[string]interface{}) {
	alert := Alert{
		ID:        generateAlertID(),
		Level:     level,
		Source:    source,
		Message:   message,
		Timestamp: time.Now(),
		Status:    "NEW",
		Metadata:  metadata,
	}

	m.mu.Lock()
	m.alerts = append(m.alerts, alert)
	m.mu.Unlock()

	// å¼æ­¥å¤çåè­¦
	go m.handleAlert(alert)
}

func (m *Manager) handleAlert(alert Alert) {
	handlers, exists := m.handlers[alert.Level]
	if !exists {
		return
	}

	for _, handler := range handlers {
		go func(h AlertHandler) {
			if err := h.Handle(alert); err != nil {
				// å¤çéè¯¯
			}
		}(handler)
	}
}

func (m *Manager) RegisterHandler(level AlertLevel, handler AlertHandler) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.handlers[level]; !exists {
		m.handlers[level] = make([]AlertHandler, 0)
	}
	m.handlers[level] = append(m.handlers[level], handler)
}

func (m *Manager) processAlerts() {
	ticker := time.NewTicker(m.config.AlertProcessInterval)
	defer ticker.Stop()

	for {
		select {
		case <-m.stopChan:
			return
		case <-ticker.C:
			m.cleanupOldAlerts()
		}
	}
}

func (m *Manager) cleanupOldAlerts() {
	m.mu.Lock()
	defer m.mu.Unlock()

	cutoff := time.Now().Add(-m.config.AlertRetentionPeriod)
	var activeAlerts []Alert

	for _, alert := range m.alerts {
		if alert.Timestamp.After(cutoff) {
			activeAlerts = append(activeAlerts, alert)
		}
	}

	m.alerts = activeAlerts
}

func (m *Manager) registerDefaultHandlers() {
	// é®ä»¶åè­¦å¤çå?
	m.RegisterHandler(CriticalLevel, &EmailAlertHandler{})
	
	// SMSåè­¦å¤çå?
	m.RegisterHandler(CriticalLevel, &SMSAlertHandler{})
	
	// æ¥å¿åè­¦å¤çå?
	m.RegisterHandler(InfoLevel, &LogAlertHandler{})
}

func generateAlertID() string {
	// çæå¯ä¸çåè­¦ID
	return time.Now().Format("20060102150405") + "-" + randomString(6)
}

// åè­¦å¤çå¨å®ç?
type EmailAlertHandler struct{}

func (h *EmailAlertHandler) Handle(alert Alert) error {
	// å®ç°é®ä»¶åéé»è¾
	return nil
}

type SMSAlertHandler struct{}

func (h *SMSAlertHandler) Handle(alert Alert) error {
	// å®ç°SMSåéé»è¾
	return nil
}

type LogAlertHandler struct{}

func (h *LogAlertHandler) Handle(alert Alert) error {
	// å®ç°æ¥å¿è®°å½é»è¾
	return nil
} 
