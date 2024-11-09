package contract

import (
	"context"
	"sync"

	"github.com/wanglilind/qqq/pkg/state"
	"github.com/wanglilind/qqq/pkg/vm"
)

type ContractEngine struct {
	vm          *vm.VirtualMachine
	state       *state.StateManager
	contracts   map[string]*Contract
	mu          sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
}

type Contract struct {
	Address     string
	Code        []byte
	State       map[string]interface{}
	Owner       string
	CreateTime  int64
	UpdateTime  int64
}

type ContractCall struct {
	ContractAddress string
	Method         string
	Args           []interface{}
	Caller         string
	Value          uint64
}

func NewContractEngine(stateManager *state.StateManager) *ContractEngine {
	ctx, cancel := context.WithCancel(context.Background())
	return &ContractEngine{
		vm:        vm.NewVirtualMachine(),
		state:     stateManager,
		contracts: make(map[string]*Contract),
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (ce *ContractEngine) DeployContract(code []byte, owner string) (string, error) {
	ce.mu.Lock()
	defer ce.mu.Unlock()

	// éªè¯åçº¦ä»£ç 
	if err := ce.vm.ValidateCode(code); err != nil {
		return "", err
	}

	// çæåçº¦å°å
	address := generateContractAddress(code, owner)

	// åå»ºåçº¦å®ä¾
	contract := &Contract{
		Address:    address,
		Code:       code,
		State:      make(map[string]interface{}),
		Owner:      owner,
		CreateTime: getCurrentTimestamp(),
		UpdateTime: getCurrentTimestamp(),
	}

	// åå§ååçº¦ç¶æ?
	if err := ce.vm.InitializeContract(contract); err != nil {
		return "", err
	}

	// å­å¨åçº¦
	ce.contracts[address] = contract

	return address, nil
}

func (ce *ContractEngine) ExecuteContract(call ContractCall) (interface{}, error) {
	ce.mu.RLock()
	contract, exists := ce.contracts[call.ContractAddress]
	ce.mu.RUnlock()

	if !exists {
		return nil, ErrContractNotFound
	}

	// åå»ºæ§è¡ç¯å¢
	env := &vm.ExecutionEnvironment{
		Contract: contract,
		Caller:   call.Caller,
		Value:    call.Value,
		State:    ce.state,
	}

	// æ§è¡åçº¦
	result, err := ce.vm.Execute(env, call.Method, call.Args)
	if err != nil {
		return nil, err
	}

	// æ´æ°åçº¦ç¶æ?
	ce.mu.Lock()
	contract.UpdateTime = getCurrentTimestamp()
	ce.mu.Unlock()

	return result, nil
}

func (ce *ContractEngine) GetContractState(address string) (map[string]interface{}, error) {
	ce.mu.RLock()
	defer ce.mu.RUnlock()

	contract, exists := ce.contracts[address]
	if !exists {
		return nil, ErrContractNotFound
	}

	return contract.State, nil
} 
