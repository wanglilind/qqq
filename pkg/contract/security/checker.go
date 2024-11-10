package security

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	
	"github.com/wanglilind/qqq/pkg/contract"
)

type SecurityChecker struct {
	knownVulnerabilities map[string]string
	maxCodeSize         int
	maxStackDepth       int
	bannedOpcodes      map[byte]bool
}

func NewSecurityChecker() *SecurityChecker {
	return &SecurityChecker{
		knownVulnerabilities: make(map[string]string),
		maxCodeSize:         1024 * 1024, // 1MB
		maxStackDepth:       1000,
		bannedOpcodes:      map[byte]bool{
			0xF4: true, // SELFDESTRUCT
			0xFF: true, // DELEGATECALL
		},
	}
}

// æ£æ¥åçº¦ä»£ç å®å¨æ?
func (sc *SecurityChecker) CheckContract(code []byte) ([]SecurityIssue, error) {
	var issues []SecurityIssue

	// æ£æ¥ä»£ç å¤§å°?
	if len(code) > sc.maxCodeSize {
		issues = append(issues, SecurityIssue{
			Type:        "CODE_SIZE",
			Severity:    "HIGH",
			Description: "Contract code size exceeds maximum limit",
		})
	}

	// æ£æ¥å·²ç¥æ¼æ´?
	codeHash := sha256.Sum256(code)
	if vulnerability, exists := sc.knownVulnerabilities[hex.EncodeToString(codeHash[:])]; exists {
		issues = append(issues, SecurityIssue{
			Type:        "KNOWN_VULNERABILITY",
			Severity:    "CRITICAL",
			Description: vulnerability,
		})
	}

	// æ£æ¥å±é©æä½ç 
	issues = append(issues, sc.checkOpcodes(code)...)

	// æ£æ¥éå¥æ¼æ´?
	issues = append(issues, sc.checkReentrancy(code)...)

	// æ£æ¥æ´æ°æº¢å?
	issues = append(issues, sc.checkIntegerOverflow(code)...)

	return issues, nil
}

// æ£æ¥æä½ç å®å¨æ?
func (sc *SecurityChecker) checkOpcodes(code []byte) []SecurityIssue {
	var issues []SecurityIssue
	
	for i := 0; i < len(code); i++ {
		if sc.bannedOpcodes[code[i]] {
			issues = append(issues, SecurityIssue{
				Type:        "BANNED_OPCODE",
				Severity:    "HIGH",
				Description: "Contract contains banned opcode",
				Location:    i,
			})
		}
	}
	
	return issues
}

// æ£æ¥éå¥æ¼æ´?
func (sc *SecurityChecker) checkReentrancy(code []byte) []SecurityIssue {
	var issues []SecurityIssue
	
	// æ£æ¥ç¶æåæ´å¨å¤é¨è°ç¨ä¹å
	if bytes.Contains(code, []byte{0xf1, 0x15}) { // CALL followed by SSTORE
		issues = append(issues, SecurityIssue{
			Type:        "REENTRANCY",
			Severity:    "HIGH",
			Description: "Potential reentrancy vulnerability detected",
		})
	}
	
	return issues
}

// æ£æ¥æ´æ°æº¢å?
func (sc *SecurityChecker) checkIntegerOverflow(code []byte) []SecurityIssue {
	var issues []SecurityIssue
	
	// æ£æ¥æ¯å¦ææªæ£æ¥çç®æ¯æä½
	if bytes.Contains(code, []byte{0x01}) && !bytes.Contains(code, []byte{0x10}) { // ADD without overflow check
		issues = append(issues, SecurityIssue{
			Type:        "INTEGER_OVERFLOW",
			Severity:    "MEDIUM",
			Description: "Potential integer overflow detected",
		})
	}
	
	return issues
}

type SecurityIssue struct {
	Type        string
	Severity    string
	Description string
	Location    int
} 
