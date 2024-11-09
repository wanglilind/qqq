package upgrade

import (
	"sync"

	"github.com/wanglilind/qqq/pkg/contract"
	"github.com/wanglilind/qqq/pkg/state"
)

// ä»£çåçº¦ï¼ç¨äºå®ç°åçº¦åçº?
type ContractProxy struct {
	mu              sync.RWMutex
	implementation  string        // å®ç°åçº¦å°å
	admin           string        // ç®¡çåå°å
	state          *state.StateManager
	upgradeHistory []UpgradeRecord
}

type UpgradeRecord struct {
	OldImplementation string
	NewImplementation string
	Timestamp        int64
	Reason           string
}

func NewContractProxy(admin string, implementation string, state *state.StateManager) *ContractProxy {
	return &ContractProxy{
		admin:          admin,
		implementation: implementation,
		state:         state,
		upgradeHistory: make([]UpgradeRecord, 0),
	}
}

// åçº§åçº¦å®ç°
func (cp *ContractProxy) Upgrade(newImplementation string, reason string) error {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	// è®°å½åçº§åå²
	record := UpgradeRecord{
		OldImplementation: cp.implementation,
		NewImplementation: newImplementation,
		Timestamp:        getCurrentTimestamp(),
		Reason:          reason,
	}
	cp.upgradeHistory = append(cp.upgradeHistory, record)

	// æ´æ°å®ç°å°å
	cp.implementation = newImplementation
	return nil
}

// å§æè°ç¨
func (cp *ContractProxy) DelegateCall(method string, args []interface{}) (interface{}, error) {
	cp.mu.RLock()
	implementation := cp.implementation
	cp.mu.RUnlock()

	// è·åå®ç°åçº¦
	contract, err := cp.state.GetContract(implementation)
	if err != nil {
		return nil, err
	}

	// æ§è¡è°ç¨
	return contract.Execute(method, args)
}

// è·ååçº§åå²
func (cp *ContractProxy) GetUpgradeHistory() []UpgradeRecord {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	
	history := make([]UpgradeRecord, len(cp.upgradeHistory))
	copy(history, cp.upgradeHistory)
	return history
} 
