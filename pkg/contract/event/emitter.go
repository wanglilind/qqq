package event

import (
	"context"
	"sync"
	"time"
)

// äºä»¶å®ä¹
type Event struct {
	ContractAddress string
	EventType      string
	Data           map[string]interface{}
	BlockNumber    uint64
	TransactionHash string
	Timestamp      time.Time
}

// äºä»¶è¿æ»¤å?
type EventFilter struct {
	ContractAddress string
	EventTypes     []string
	FromBlock      uint64
	ToBlock        uint64
}

// äºä»¶åå°å?
type EventEmitter struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Event
	events      []Event
	maxEvents   int
}

func NewEventEmitter(maxEvents int) *EventEmitter {
	return &EventEmitter{
		subscribers: make(map[string][]chan Event),
		events:     make([]Event, 0, maxEvents),
		maxEvents:  maxEvents,
	}
}

// ååºäºä»¶
func (ee *EventEmitter) Emit(event Event) {
	ee.mu.Lock()
	// å­å¨äºä»¶
	ee.events = append(ee.events, event)
	if len(ee.events) > ee.maxEvents {
		ee.events = ee.events[1:]
	}
	
	// è·åè®¢éè?
	subscribers := ee.subscribers[event.EventType]
	ee.mu.Unlock()

	// éç¥è®¢éè?
	for _, ch := range subscribers {
		select {
		case ch <- event:
		default:
			// å¦æchannelå·²æ»¡ï¼è·³è¿?
		}
	}
}

// è®¢éäºä»¶
func (ee *EventEmitter) Subscribe(ctx context.Context, filter EventFilter) (<-chan Event, error) {
	ch := make(chan Event, 100)
	
	ee.mu.Lock()
	for _, eventType := range filter.EventTypes {
		ee.subscribers[eventType] = append(ee.subscribers[eventType], ch)
	}
	ee.mu.Unlock()

	// åéåå²äºä»?
	go ee.sendHistoricalEvents(ctx, filter, ch)

	return ch, nil
}

// åéåå²äºä»?
func (ee *EventEmitter) sendHistoricalEvents(ctx context.Context, filter EventFilter, ch chan Event) {
	ee.mu.RLock()
	defer ee.mu.RUnlock()

	for _, event := range ee.events {
		if ee.matchesFilter(event, filter) {
			select {
			case <-ctx.Done():
				return
			case ch <- event:
			}
		}
	}
}

// æ£æ¥äºä»¶æ¯å¦å¹éè¿æ»¤å¨
func (ee *EventEmitter) matchesFilter(event Event, filter EventFilter) bool {
	if filter.ContractAddress != "" && event.ContractAddress != filter.ContractAddress {
		return false
	}

	if filter.FromBlock > event.BlockNumber || (filter.ToBlock != 0 && filter.ToBlock < event.BlockNumber) {
		return false
	}

	if len(filter.EventTypes) > 0 {
		matched := false
		for _, t := range filter.EventTypes {
			if t == event.EventType {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	return true
}

// åæ¶è®¢é
func (ee *EventEmitter) Unsubscribe(eventType string, ch chan Event) {
	ee.mu.Lock()
	defer ee.mu.Unlock()

	subscribers := ee.subscribers[eventType]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			ee.subscribers[eventType] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
} 
