package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type ErrorCode int

const (
	// 系统级错误码
	ErrSystem ErrorCode = iota + 1000
	ErrDatabase
	ErrCache
	ErrNetwork
	
	// 业务级错误码
	ErrInvalidInput ErrorCode = iota + 2000
	ErrUnauthorized
	ErrForbidden
	ErrNotFound
	ErrDuplicate
	
	// 自定义错误码
	ErrIdentityVerification ErrorCode = iota + 3000
	ErrTransactionFailed
	ErrConsensusTimeout
)

type Error struct {
	Code    ErrorCode
	Message string
	Cause   error
	Stack   string
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Stack:   getStackTrace(),
	}
}

func Wrap(err error, code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   err,
		Stack:   getStackTrace(),
	}
}

func getStackTrace() string {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	
	var builder strings.Builder
	for {
		frame, more := frames.Next()
		if !strings.Contains(frame.File, "runtime/") {
			fmt.Fprintf(&builder, "%s:%d\n", frame.File, frame.Line)
		}
		if !more {
			break
		}
	}
	return builder.String()
} 
