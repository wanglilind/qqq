package validator

import (
	"context"
	"sync"

	"github.com/wanglilind/qqq/pkg/contract"
	"github.com/wanglilind/qqq/pkg/contract/security"
)

// åçº¦éªè¯å?
type ContractValidator struct {
	securityChecker *security.SecurityChecker
	rules          map[string]ValidationRule
	results        map[string]ValidationResult
	mu             sync.RWMutex
}

type ValidationRule struct {
	Name        string
	Description string
	Severity    string
	Validator   func([]byte) error
}

type ValidationResult struct {
	Valid       bool
	Errors      []ValidationError
	Warnings    []ValidationWarning
	Suggestions []string
}

type ValidationError struct {
	Rule        string
	Message     string
	Location    string
	Severity    string
}

type ValidationWarning struct {
	Rule        string
	Message     string
	Location    string
}

func NewContractValidator() *ContractValidator {
	validator := &ContractValidator{
		securityChecker: security.NewSecurityChecker(),
		rules:          make(map[string]ValidationRule),
		results:        make(map[string]ValidationResult),
	}

	// æ³¨åé»è®¤éªè¯è§å
	validator.registerDefaultRules()
	return validator
}

// éªè¯åçº¦
func (cv *ContractValidator) ValidateContract(ctx context.Context, code []byte) (*ValidationResult, error) {
	cv.mu.Lock()
	defer cv.mu.Unlock()

	result := &ValidationResult{
		Valid:       true,
		Errors:      make([]ValidationError, 0),
		Warnings:    make([]ValidationWarning, 0),
		Suggestions: make([]string, 0),
	}

	// è¿è¡ææéªè¯è§å?
	for _, rule := range cv.rules {
		if err := rule.Validator(code); err != nil {
			result.Valid = false
			result.Errors = append(result.Errors, ValidationError{
				Rule:     rule.Name,
				Message: err.Error(),
				Severity: rule.Severity,
			})
		}
	}

	// è¿è¡å®å¨æ£æ?
	securityIssues, err := cv.securityChecker.CheckContract(code)
	if err != nil {
		return nil, err
	}

	// å¤çå®å¨é®é¢
	for _, issue := range securityIssues {
		if issue.Severity == "HIGH" || issue.Severity == "CRITICAL" {
			result.Valid = false
			result.Errors = append(result.Errors, ValidationError{
				Rule:     "SecurityCheck",
				Message: issue.Description,
				Severity: issue.Severity,
			})
		} else {
			result.Warnings = append(result.Warnings, ValidationWarning{
				Rule:     "SecurityCheck",
				Message: issue.Description,
			})
		}
	}

	// çæä¼åå»ºè®®
	result.Suggestions = cv.generateSuggestions(code)

	// å­å¨ç»æ
	cv.results[string(code)] = *result

	return result, nil
}

// æ³¨åéªè¯è§å
func (cv *ContractValidator) RegisterRule(rule ValidationRule) {
	cv.mu.Lock()
	defer cv.mu.Unlock()
	cv.rules[rule.Name] = rule
}

// æ³¨åé»è®¤è§å
func (cv *ContractValidator) registerDefaultRules() {
	// ä»£ç å¤§å°éå¶
	cv.RegisterRule(ValidationRule{
		Name:        "CodeSize",
		Description: "Check contract code size",
		Severity:    "HIGH",
		Validator:   cv.validateCodeSize,
	})

	// å¤æåº¦æ£æ?
	cv.RegisterRule(ValidationRule{
		Name:        "Complexity",
		Description: "Check contract complexity",
		Severity:    "MEDIUM",
		Validator:   cv.validateComplexity,
	})

	// èµæºä½¿ç¨æ£æ?
	cv.RegisterRule(ValidationRule{
		Name:        "ResourceUsage",
		Description: "Check resource usage",
		Severity:    "HIGH",
		Validator:   cv.validateResourceUsage,
	})
}

// éªè¯ä»£ç å¤§å°
func (cv *ContractValidator) validateCodeSize(code []byte) error {
	// å®ç°ä»£ç å¤§å°éªè¯é»è¾
	return nil
}

// éªè¯å¤æåº?
func (cv *ContractValidator) validateComplexity(code []byte) error {
	// å®ç°å¤æåº¦éªè¯é»è¾
	return nil
}

// éªè¯èµæºä½¿ç¨
func (cv *ContractValidator) validateResourceUsage(code []byte) error {
	// å®ç°èµæºä½¿ç¨éªè¯é»è¾
	return nil
}

// çæä¼åå»ºè®®
func (cv *ContractValidator) generateSuggestions(code []byte) []string {
	// å®ç°ä¼åå»ºè®®çæé»è¾
	return nil
} 
