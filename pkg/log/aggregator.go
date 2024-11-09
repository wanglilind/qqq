package log

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

type LogAggregator struct {
	esClient    *elasticsearch.Client
	indexPrefix string
	batchSize   int
	batchDelay  time.Duration
	logChan     chan LogEntry
}

type LogEntry struct {
	Timestamp   time.Time              `json:"timestamp"`
	Level       string                 `json:"level"`
	Service     string                 `json:"service"`
	Message     string                 `json:"message"`
	TraceID     string                 `json:"trace_id,omitempty"`
	SpanID      string                 `json:"span_id,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

func NewLogAggregator(config *Config) (*LogAggregator, error) {
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.ElasticsearchURLs,
	})
	if err != nil {
		return nil, err
	}

	aggregator := &LogAggregator{
		esClient:    esClient,
		indexPrefix: config.IndexPrefix,
		batchSize:   config.BatchSize,
		batchDelay:  config.BatchDelay,
		logChan:     make(chan LogEntry, 10000),
	}

	go aggregator.processBatch()
	return aggregator, nil
}

func (la *LogAggregator) Log(entry LogEntry) {
	la.logChan <- entry
}

func (la *LogAggregator) processBatch() {
	var batch []LogEntry
	ticker := time.NewTicker(la.batchDelay)
	defer ticker.Stop()

	for {
		select {
		case entry := <-la.logChan:
			batch = append(batch, entry)
			if len(batch) >= la.batchSize {
				la.sendBatch(batch)
				batch = nil
			}
		case <-ticker.C:
			if len(batch) > 0 {
				la.sendBatch(batch)
				batch = nil
			}
		}
	}
}

func (la *LogAggregator) sendBatch(batch []LogEntry) {
	index := fmt.Sprintf("%s-%s", la.indexPrefix, time.Now().Format("2006.01.02"))
	
	for _, entry := range batch {
		data, err := json.Marshal(entry)
		if err != nil {
			continue
		}

		_, err = la.esClient.Index(
			index,
			strings.NewReader(string(data)),
			la.esClient.Index.WithContext(context.Background()),
		)
		if err != nil {
			// 处理错误，可能需要重试或记录失败
			continue
		}
	}
} 
