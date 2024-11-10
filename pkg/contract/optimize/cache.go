package optimize

import (
	"sync"
	"time"

	"github.com/wanglilind/qqq/pkg/contract"
)

// åçº¦ç¼å­ç®¡çå?
type ContractCache struct {
	mu            sync.RWMutex
	codeCache     map[string][]byte
	stateCache    map[string]map[string]interface{}
	resultCache   map[string]CachedResult
	maxSize       int
	expiration    time.Duration
}

type CachedResult struct {
	Result     interface{}
	Timestamp  time.Time
	AccessCount int
}

func NewContractCache(maxSize int, expiration time.Duration) *ContractCache {
	return &ContractCache{
		codeCache:   make(map[string][]byte),
		stateCache:  make(map[string]map[string]interface{}),
		resultCache: make(map[string]CachedResult),
		maxSize:     maxSize,
		expiration:  expiration,
	}
}

// ç¼å­åçº¦ä»£ç 
func (cc *ContractCache) CacheCode(address string, code []byte) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	if len(cc.codeCache) >= cc.maxSize {
		cc.evictOldest()
	}

	cc.codeCache[address] = code
}

// ç¼å­åçº¦ç¶æ?
func (cc *ContractCache) CacheState(address string, state map[string]interface{}) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	cc.stateCache[address] = state
}

// ç¼å­æ§è¡ç»æ
func (cc *ContractCache) CacheResult(key string, result interface{}) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	if len(cc.resultCache) >= cc.maxSize {
		cc.evictLeastUsed()
	}

	cc.resultCache[key] = CachedResult{
		Result:     result,
		Timestamp:  time.Now(),
		AccessCount: 0,
	}
}

// è·åç¼å­ç»æ
func (cc *ContractCache) GetResult(key string) (interface{}, bool) {
	cc.mu.RLock()
	defer cc.mu.RUnlock()

	if cached, exists := cc.resultCache[key]; exists {
		if time.Since(cached.Timestamp) > cc.expiration {
			delete(cc.resultCache, key)
			return nil, false
		}
		
		cached.AccessCount++
		cc.resultCache[key] = cached
		return cached.Result, true
	}
	return nil, false
}

// æ¸çè¿æç¼å­
func (cc *ContractCache) CleanupExpired() {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	now := time.Now()
	for key, cached := range cc.resultCache {
		if now.Sub(cached.Timestamp) > cc.expiration {
			delete(cc.resultCache, key)
		}
	}
}

// é©±éææ§çç¼å­
func (cc *ContractCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, cached := range cc.resultCache {
		if oldestKey == "" || cached.Timestamp.Before(oldestTime) {
			oldestKey = key
			oldestTime = cached.Timestamp
		}
	}

	if oldestKey != "" {
		delete(cc.resultCache, oldestKey)
	}
}

// é©±éæå°ä½¿ç¨çç¼å­
func (cc *ContractCache) evictLeastUsed() {
	var leastUsedKey string
	var leastCount int

	for key, cached := range cc.resultCache {
		if leastUsedKey == "" || cached.AccessCount < leastCount {
			leastUsedKey = key
			leastCount = cached.AccessCount
		}
	}

	if leastUsedKey != "" {
		delete(cc.resultCache, leastUsedKey)
	}
} 
