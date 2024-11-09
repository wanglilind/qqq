package vm

import (
	"fmt"
	"sync"

	"github.com/wanglilind/qqq/pkg/state"
)

type VirtualMachine struct {
	mu           sync.RWMutex
	instructions map[string]Instruction
	memory       []byte
	stack        []interface{}
	pc          uint64
}

type ExecutionEnvironment struct {
	Contract *Contract
	Caller   string
	Value    uint64
	State    *state.StateManager
}

type Instruction interface {
	Execute(vm *VirtualMachine, env *ExecutionEnvironment) error
}

func NewVirtualMachine() *VirtualMachine {
	vm := &VirtualMachine{
		instructions: make(map[string]Instruction),
		memory:      make([]byte, 65536), // 64KB åå­
		stack:       make([]interface{}, 0),
		pc:         0,
	}

	// æ³¨ååºæ¬æä»¤é?
	vm.registerInstructions()

	return vm
}

func (vm *VirtualMachine) ValidateCode(code []byte) error {
	// éªè¯ä»£ç æ ¼å¼åå®å¨æ?
	return nil
}

func (vm *VirtualMachine) Execute(env *ExecutionEnvironment, method string, args []interface{}) (interface{}, error) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	// éç½®VMç¶æ?
	vm.reset()

	// å è½½åçº¦ä»£ç 
	if err := vm.loadCode(env.Contract.Code); err != nil {
		return nil, err
	}

	// è®¾ç½®åæ°
	if err := vm.setArgs(args); err != nil {
		return nil, err
	}

	// æ§è¡ä»£ç 
	for vm.pc < uint64(len(env.Contract.Code)) {
		opcode := env.Contract.Code[vm.pc]
		instruction, exists := vm.instructions[string(opcode)]
		if !exists {
			return nil, fmt.Errorf("invalid opcode: %d", opcode)
		}

		if err := instruction.Execute(vm, env); err != nil {
			return nil, err
		}

		vm.pc++
	}

	// è·åæ§è¡ç»æ
	if len(vm.stack) == 0 {
		return nil, nil
	}
	return vm.stack[len(vm.stack)-1], nil
}

func (vm *VirtualMachine) reset() {
	vm.pc = 0
	vm.stack = vm.stack[:0]
	for i := range vm.memory {
		vm.memory[i] = 0
	}
}

func (vm *VirtualMachine) registerInstructions() {
	// æ³¨ååºæ¬æä»¤
	vm.instructions["PUSH"] = &PushInstruction{}
	vm.instructions["POP"] = &PopInstruction{}
	vm.instructions["ADD"] = &AddInstruction{}
	vm.instructions["SUB"] = &SubInstruction{}
	vm.instructions["MUL"] = &MulInstruction{}
	vm.instructions["DIV"] = &DivInstruction{}
	vm.instructions["STORE"] = &StoreInstruction{}
	vm.instructions["LOAD"] = &LoadInstruction{}
	vm.instructions["CALL"] = &CallInstruction{}
	vm.instructions["RETURN"] = &ReturnInstruction{}
}

// åºæ¬æä»¤å®ç°
type PushInstruction struct{}
func (i *PushInstruction) Execute(vm *VirtualMachine, env *ExecutionEnvironment) error {
	vm.pc++
	value := vm.memory[vm.pc]
	vm.stack = append(vm.stack, value)
	return nil
}

type PopInstruction struct{}
func (i *PopInstruction) Execute(vm *VirtualMachine, env *ExecutionEnvironment) error {
	if len(vm.stack) == 0 {
		return fmt.Errorf("stack underflow")
	}
	vm.stack = vm.stack[:len(vm.stack)-1]
	return nil
}

// å¶ä»æä»¤å®ç°... 
