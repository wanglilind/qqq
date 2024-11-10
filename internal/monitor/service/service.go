package service

import (
	"context"
	"sync"
	"time"

	"github.com/wanglilind/qqq/internal/monitor/performance"
	"github.com/wanglilind/qqq/internal/monitor/security"
	"github.com/wanglilind/qqq/internal/monitor/alert"
	"github.com/wanglilind/qqq/pkg/config"
)

type MonitorService struct {
	performanceMonitor *performance.Monitor
	securityAuditor   *security.Auditor
	alertManager      *alert.Manager
	config           *config.Config
	mu               sync.RWMutex
	isCollecting     bool
	metricsChan      chan Metric
}

func NewMonitorService(cfg *config.Config) *MonitorService {
	return &MonitorService{
		performanceMonitor: performance.NewMonitor(cfg),
		securityAuditor:   security.NewAuditor(cfg),
		alertManager:      alert.NewManager(cfg),
		config:           cfg,
		metricsChan:      make(chan Metric, 1000),
	}
}

// StartMetricsCollection å¯å¨ææ æ¶é
func (s *MonitorService) StartMetricsCollection() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isCollecting {
		return nil
	}

	// å¯å¨æ§è½çæ§
	if err := s.performanceMonitor.Start(); err != nil {
		return err
	}

	// å¯å¨å®å¨å®¡è®¡
	if err := s.securityAuditor.Start(); err != nil {
		s.performanceMonitor.Stop()
		return err
	}

	// å¯å¨åè­¦ç®¡ç
	if err := s.alertManager.Start(); err != nil {
		s.performanceMonitor.Stop()
		s.securityAuditor.Stop()
		return err
	}

	// å¯å¨ææ å¤çåç¨
	go s.processMetrics()

	s.isCollecting = true
	return nil
}

// StopMetricsCollection åæ­¢ææ æ¶é
func (s *MonitorService) StopMetricsCollection() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isCollecting {
		return nil
	}

	// åæ­¢ææçæ§ç»ä»?
	s.performanceMonitor.Stop()
	s.securityAuditor.Stop()
	s.alertManager.Stop()

	close(s.metricsChan)
	s.isCollecting = false
	return nil
}

// GetSystemMetrics è·åç³»ç»ææ 
func (s *MonitorService) GetSystemMetrics(ctx context.Context, req *GetSystemMetricsRequest) (*GetSystemMetricsResponse, error) {
	metrics := s.performanceMonitor.GetMetrics()
	return &GetSystemMetricsResponse{
		Metrics: metrics,
	}, nil
}

// GetSecurityAuditLog è·åå®å¨å®¡è®¡æ¥å¿
func (s *MonitorService) GetSecurityAuditLog(ctx context.Context, req *GetSecurityAuditLogRequest) (*GetSecurityAuditLogResponse, error) {
	logs := s.securityAuditor.GetAuditLogs(req.StartTime, req.EndTime)
	return &GetSecurityAuditLogResponse{
		Logs: logs,
	}, nil
}

// processMetrics å¤çæ¶éå°çææ 
func (s *MonitorService) processMetrics() {
	ticker := time.NewTicker(s.config.MetricsProcessInterval)
	defer ticker.Stop()

	for {
		select {
		case metric, ok := <-s.metricsChan:
			if !ok {
				return
			}
			// å¤çææ 
			s.handleMetric(metric)
		case <-ticker.C:
			// å®æå¤çç§¯ç´¯çææ ?
			s.processAccumulatedMetrics()
		}
	}
}

// handleMetric å¤çåä¸ªææ 
func (s *MonitorService) handleMetric(metric Metric) {
	// æ£æ¥æ¯å¦éè¦è§¦ååè­?
	if s.shouldTriggerAlert(metric) {
		s.alertManager.TriggerAlert(metric)
	}

	// å­å¨ææ 
	s.storeMetric(metric)
}

// shouldTriggerAlert å¤æ­æ¯å¦éè¦è§¦ååè­?
func (s *MonitorService) shouldTriggerAlert(metric Metric) bool {
	// å®ç°åè­¦è§¦åé»è¾
	return metric.Value > s.config.AlertThreshold
}

// storeMetric å­å¨ææ 
func (s *MonitorService) storeMetric(metric Metric) {
	// å®ç°ææ å­å¨é»è¾
}

// processAccumulatedMetrics å¤çç´¯ç§¯çææ ?
func (s *MonitorService) processAccumulatedMetrics() {
	// å®ç°æ¹éææ å��çé»è¾
} 
