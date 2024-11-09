package interop

import (
	"context"
	"encoding/json"

	"github.com/wanglilind/qqq/pkg/contract"
	"github.com/wanglilind/qqq/pkg/state"
)

// åçº¦äºæä½æ¥å?
type ContractInterface struct {
	address     string
	abi         []ABIMethod
	state       *state.StateManager
	permissions map[string]bool
}

type ABIMethod struct {
	Name       string
	Inputs     []ABIParameter
	Outputs    []ABIParameter
	Constant   bool
	Payable    bool
	Visibility string
}

type ABIParameter struct {
	Name    string
	Type    string
	Indexed bool
}

func NewContractInterface(address string, abiJSON []byte) (*ContractInterface, error) {
	var abi []ABIMethod
	if err := json.Unmarshal(abiJSON, &abi); err != nil {
		return nil, err
	}

	return &ContractInterface{
		address:     address,
		abi:         abi,
		permissions: make(map[string]bool),
	}, nil
}

// è°ç¨å¶ä»åçº¦
func (ci *ContractInterface) Call(ctx context.Context, method string, args ...interface{}) (interface{}, error) {
	// æ£æ¥æé?
	if !ci.hasPermission(method) {
		return nil, ErrPermissionDenied
	}

	// éªè¯æ¹æ³ååæ?
	if err := ci.validateMethod(method, args); err != nil {
		return nil, err
	}

	// åå»ºè°ç¨
	call := contract.ContractCall{
		ContractAddress: ci.address,
		Method:         method,
		Args:          args,
	}

	// æ§è¡è°ç¨
	return ci.executeCall(ctx, call)
}

// ææè®¿é®
func (ci *ContractInterface) GrantAccess(method string) {
	ci.permissions[method] = true
}

// æ¤éè®¿é®
func (ci *ContractInterface) RevokeAccess(method string) {
	delete(ci.permissions, method)
}

// éªè¯æ¹æ³
func (ci *ContractInterface) validateMethod(method string, args []interface{}) error {
	for _, m := range ci.abi {
		if m.Name == method {
			if len(m.Inputs) != len(args) {
				return ErrInvalidArguments
			}
			// éªè¯åæ°ç±»å
			return nil
		}
	}
	return ErrMethodNotFound
}

// æ£æ¥æé?
func (ci *ContractInterface) hasPermission(method string) bool {
	return ci.permissions[method]
}

// æ§è¡è°ç¨
func (ci *ContractInterface) executeCall(ctx context.Context, call contract.ContractCall) (interface{}, error) {
	// å®ç°è·¨åçº¦è°ç¨é»è¾
	return nil, nil
} 
