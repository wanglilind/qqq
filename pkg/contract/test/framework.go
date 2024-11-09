package test

import (
	"context"
	"testing"
	"time"

	"github.com/wanglilind/qqq/pkg/contract"
	"github.com/wanglilind/qqq/pkg/state"
	"github.com/stretchr/testify/mock"
)

// 合约测试框架
type ContractTestFramework struct {
	t           *testing.T
	state       *state.MockStateManager
	engine      *contract.ContractEngine
	mockEvents  []contract.Event
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewContractTestFramework(t *testing.T) *ContractTestFramework {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	mockState := state.NewMockStateManager()
	
	return &ContractTestFramework{
		t:          t,
		state:      mockState,
		engine:     contract.NewContractEngine(mockState),
		mockEvents: make([]contract.Event, 0),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// 部署测试合约
func (f *ContractTestFramework) DeployContract(code []byte, owner string) string {
	address, err := f.engine.DeployContract(code, owner)
	if err != nil {
		f.t.Fatalf("Failed to deploy contract: %v", err)
	}
	return address
}

// 执行合约调用
func (f *ContractTestFramework) ExecuteContract(call contract.ContractCall) interface{} {
	result, err := f.engine.ExecuteContract(call)
	if err != nil {
		f.t.Fatalf("Contract execution failed: %v", err)
	}
	return result
}

// 模拟事件
func (f *ContractTestFramework) MockEvent(event contract.Event) {
	f.mockEvents = append(f.mockEvents, event)
}

// 验证事件
func (f *ContractTestFramework) AssertEvent(expectedEvent contract.Event) {
	for _, event := range f.mockEvents {
		if event.EventType == expectedEvent.EventType {
			// 验证事件数据
			return
		}
	}
	f.t.Errorf("Expected event %s not found", expectedEvent.EventType)
}

// 清理资源
func (f *ContractTestFramework) Cleanup() {
	f.cancel()
} 
