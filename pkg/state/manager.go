package state

import (
	"context"
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/blockchain"
	"github.com/wanglilind/qqq/pkg/database"
)

type StateManager struct {
	db          *database.PostgresDB
	cache       map[string]interface{}
	mu          sync.RWMutex
	subscribers map[string][]chan StateUpdate
	ctx         context.Context
	cancel      context.CancelFunc
}

type StateUpdate struct {
	Type      string
	Key       string
	Value     interface{}
	BlockHash string
	Timestamp time.Time
}

func NewStateManager(db *database.PostgresDB) *StateManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &StateManager{
		db:          db,
		cache:       make(map[string]interface{}),
		subscribers: make(map[string][]chan StateUpdate),
		ctx:         ctx,
		cancel:      cancel,
	}
}

func (sm *StateManager) ApplyBlock(block *blockchain.Block) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// å¼å§æ°æ®åºäºå¡
	tx, err := sm.db.BeginTx(sm.ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// åºç¨æ¯ä¸ªäº¤æ
	for _, txn := range block.Body.Transactions {
		if err := sm.applyTransaction(tx, txn); err != nil {
			return err
		}
	}

	// æäº¤äºå¡
	if err := tx.Commit(); err != nil {
		return err
	}

	// æ´æ°ç¼å­
	sm.updateCache(block)

	// éç¥è®¢éè?
	sm.notifySubscribers(block)

	return nil
}

func (sm *StateManager) GetState(key string) (interface{}, error) {
	sm.mu.RLock()
	if value, exists := sm.cache[key]; exists {
		sm.mu.RUnlock()
		return value, nil
	}
	sm.mu.RUnlock()

	// ä»æ°æ®åºå è½½
	value, err := sm.loadFromDB(key)
	if err != nil {
		return nil, err
	}

	sm.mu.Lock()
	sm.cache[key] = value
	sm.mu.Unlock()

	return value, nil
}

func (sm *StateManager) Subscribe(stateType string) chan StateUpdate {
	ch := make(chan StateUpdate, 100)
	
	sm.mu.Lock()
	sm.subscribers[stateType] = append(sm.subscribers[stateType], ch)
	sm.mu.Unlock()

	return ch
}

func (sm *StateManager) applyTransaction(tx *database.Tx, transaction blockchain.Transaction) error {
	// å®ç°å·ä½çäº¤æåºç¨é»è¾
	return nil
}

func (sm *StateManager) updateCache(block *blockchain.Block) {
	// æ´æ°åå­ç¼å­
}

func (sm *StateManager) notifySubscribers(block *blockchain.Block) {
	update := StateUpdate{
		BlockHash: block.Hash,
		Timestamp: block.Header.Timestamp,
	}

	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, subs := range sm.subscribers {
		for _, ch := range subs {
			select {
			case ch <- update:
			default:
				// å¦æchannelå·²æ»¡ï¼è·³è¿?
			}
		}
	}
}

func (sm *StateManager) loadFromDB(key string) (interface{}, error) {
	// ä»æ°æ®åºå è½½ç¶æ?
	return nil, nil
}

func (sm *StateManager) Close() {
	sm.cancel()
	// æ¸çèµæº
} 
