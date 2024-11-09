package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Collector struct {
	requestCounter   *prometheus.CounterVec
	requestDuration  *prometheus.HistogramVec
	activeGoroutines prometheus.Gauge
	errorCounter     *prometheus.CounterVec
}

func NewCollector(serviceName string) *Collector {
	return &Collector{
		requestCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "requests_total",
				Help: "Total number of requests processed",
			},
			[]string{"service", "method", "status"},
		),
		requestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "request_duration_seconds",
				Help:    "Request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"service", "method"},
		),
		activeGoroutines: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "goroutines_active",
				Help: "Number of active goroutines",
			},
		),
		errorCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "errors_total",
				Help: "Total number of errors",
			},
			[]string{"service", "type"},
		),
	}
}

func (c *Collector) RecordRequest(service, method, status string) {
	c.requestCounter.WithLabelValues(service, method, status).Inc()
}

func (c *Collector) RecordDuration(service, method string, duration float64) {
	c.requestDuration.WithLabelValues(service, method).Observe(duration)
}

func (c *Collector) RecordError(service, errorType string) {
	c.errorCounter.WithLabelValues(service, errorType).Inc()
}

func (c *Collector) UpdateGoroutines(count int) {
	c.activeGoroutines.Set(float64(count))
} 
