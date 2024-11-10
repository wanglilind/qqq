package stdlib

import (
	"time"
	"math/big"
	
	"github.com/wanglilind/qqq/pkg/contract"
)

// æ ååºåçº¦æ¥å?
type StandardContract interface {
	Initialize() error
	GetBalance() (*big.Int, error)
	Transfer(to string, amount *big.Int) error
	GetOwner() string
	IsExpired() bool
}

// åºç¡åçº¦å®ç°
type BaseContract struct {
	contract.Contract
	decayRate     *big.Int
	creationTime  time.Time
	maxLifespan   time.Duration
}

func NewBaseContract() *BaseContract {
	return &BaseContract{
		decayRate:    big.NewInt(120), // 120å¹´è¡°å?
		maxLifespan:  120 * 365 * 24 * time.Hour,
		creationTime: time.Now(),
	}
}

// è®¡ç®å½åä½é¢ï¼èèè¡°åï¼?
func (bc *BaseContract) GetBalance() (*big.Int, error) {
	currentBalance := new(big.Int).Set(bc.State["balance"].(*big.Int))
	elapsedTime := time.Since(bc.creationTime)
	
	// è®¡ç®è¡°å
	if elapsedTime > 0 {
		decayFactor := new(big.Int).Mul(
			big.NewInt(int64(elapsedTime.Hours())/24/365),
			bc.decayRate,
		)
		currentBalance.Sub(currentBalance, decayFactor)
		if currentBalance.Sign() < 0 {
			currentBalance = big.NewInt(0)
		}
	}
	
	return currentBalance, nil
}

// è½¬è´¦å®ç°
func (bc *BaseContract) Transfer(to string, amount *big.Int) error {
	balance, err := bc.GetBalance()
	if err != nil {
		return err
	}
	
	if balance.Cmp(amount) < 0 {
		return ErrInsufficientBalance
	}
	
	// æ§è¡è½¬è´¦
	bc.State["balance"].(*big.Int).Sub(bc.State["balance"].(*big.Int), amount)
	return nil
}

// æ£æ¥åçº¦æ¯å¦è¿æ?
func (bc *BaseContract) IsExpired() bool {
	return time.Since(bc.creationTime) > bc.maxLifespan
} 
