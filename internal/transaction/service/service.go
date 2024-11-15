package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/wanglilind/qqq/pkg/config"
	"github.com/wanglilind/qqq/pkg/database"
	pb "github.com/wanglilind/qqq/api/proto"
)

const (
	ErrInvalidRequest = "无效的请求"
	ErrInsufficientBalance = "余额不足"
	ErrInvalidSignature = "无效的签名"
)

type TransactionError struct {
	Code    string
	Message string
}

func (e *TransactionError) Error() string {
	return e.Message
}

type TransactionService struct {
	config *config.Config
	db     *database.PostgresDB
	mu     sync.RWMutex
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(cfg *config.Config) *TransactionService {
	// 初始化数据库连接
	db, err := database.NewPostgresDB(&database.PostgresConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.DBName,
	})
	if err != nil {
		panic(err)
	}

	return &TransactionService{
		config: cfg,
		db:     db,
	}
}

// CreateTransaction 创建新交易
func (s *TransactionService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 验证交易请求
	if err := s.validateTransactionRequest(req); err != nil {
		return nil, err
	}

	// 检查余额
	if err := s.checkBalance(ctx, req.SenderId, req.Amount); err != nil {
		return nil, err
	}

	// 创建交易记录
	tx := &pb.Transaction{
		TransactionId: generateTransactionID(),
		SenderId:     req.SenderId,
		RecipientId:  req.RecipientId,
		Amount:       req.Amount,
		Type:        req.Type,
		Timestamp:    time.Now().Unix(),
		Status:      "PENDING",
	}

	// 执行交易
	if err := s.executeTransaction(ctx, tx); err != nil {
		return nil, err
	}

	return &pb.CreateTransactionResponse{
		TransactionId: tx.TransactionId,
		Status:       "SUCCESS",
		Timestamp:    tx.Timestamp,
		Message:      "Transaction created successfully",
	}, nil
}

// GetTransactionHistory 获取交易历史
func (s *TransactionService) GetTransactionHistory(ctx context.Context, req *pb.GetTransactionHistoryRequest) (*pb.GetTransactionHistoryResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 查询交易历史
	transactions, err := s.queryTransactionHistory(ctx, req)
	if err != nil {
		return nil, err
	}

	// 计算总数
	totalCount := len(transactions)
	hasMore := false
	if totalCount > req.PageSize {
		hasMore = true
		transactions = transactions[:req.PageSize]
	}

	return &pb.GetTransactionHistoryResponse{
		Transactions: transactions,
		TotalCount:  int32(totalCount),
		HasMore:     hasMore,
	}, nil
}

// ValidateTransaction 验证交易
func (s *TransactionService) ValidateTransaction(ctx context.Context, req *pb.ValidateTransactionRequest) (*pb.ValidateTransactionResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 验证交易签名
	if err := s.verifyTransactionSignature(req.Transaction); err != nil {
		return &pb.ValidateTransactionResponse{
			Valid:   false,
			Message: err.Error(),
		}, nil
	}

	// 验证交易状态
	if err := s.verifyTransactionStatus(ctx, req.Transaction); err != nil {
		return &pb.ValidateTransactionResponse{
			Valid:   false,
			Message: err.Error(),
		}, nil
	}

	return &pb.ValidateTransactionResponse{
		Valid:   true,
		Message: "Transaction is valid",
	}, nil
}

// GetBalance 获取账户余额
func (s *TransactionService) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 查询账户余额
	balance, updateTime, err := s.queryBalance(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetBalanceResponse{
		Balance:        balance,
		UpdateTime:     updateTime,
		CurrencyStatus: "ACTIVE",
	}, nil
}

// 内部辅助方法
func (s *TransactionService) validateTransactionRequest(req *pb.CreateTransactionRequest) error {
	if req.SenderId == "" || req.RecipientId == "" {
		return &TransactionError{
			Code:    ErrInvalidRequest,
			Message: "发送者和接收者ID不能为空",
		}
	}
	if req.Amount == 0 {
		return &TransactionError{
			Code:    ErrInvalidRequest,
			Message: "交易金额必须大于0",
		}
	}
	return nil
}

func (s *TransactionService) checkBalance(ctx context.Context, userID string, amount uint64) error {
	// 实现余额检查逻辑
	return nil
}

func (s *TransactionService) executeTransaction(ctx context.Context, tx *pb.Transaction) error {
	// 实现交易执行逻辑
	return nil
}

func (s *TransactionService) queryTransactionHistory(ctx context.Context, req *pb.GetTransactionHistoryRequest) ([]*pb.Transaction, error) {
	// 实现交易历史查询逻辑
	return nil, nil
}

func (s *TransactionService) verifyTransactionSignature(tx *pb.Transaction) error {
	// 实现交易签名验证逻辑
	return nil
}

func (s *TransactionService) verifyTransactionStatus(ctx context.Context, tx *pb.Transaction) error {
	// 实现交易状态验证逻辑
	return nil
}

func (s *TransactionService) queryBalance(ctx context.Context, userID string) (uint64, int64, error) {
	// 实现余额查询逻辑
	return 0, 0, nil
}

func generateTransactionID() string {
	return fmt.Sprintf("TX_%d_%s", time.Now().UnixNano(), uuid.New().String()[:8])
}

func (s *TransactionService) initDatabase() error {
	// 检查必要的数据库表是否存在，如果不存在则创建
	queries := []string{
		`CREATE TABLE IF NOT EXISTS transactions (
			id TEXT PRIMARY KEY,
			sender_id TEXT NOT NULL,
			recipient_id TEXT NOT NULL,
			amount BIGINT NOT NULL,
			type TEXT NOT NULL,
			status TEXT NOT NULL,
			timestamp BIGINT NOT NULL,
			signature BYTEA
		)`,
		`CREATE TABLE IF NOT EXISTS balances (
			user_id TEXT PRIMARY KEY,
			amount BIGINT NOT NULL,
			update_time BIGINT NOT NULL
		)`,
	}
	
	for _, query := range queries {
		if _, err := s.db.Exec(query); err != nil {
			return fmt.Errorf("初始化数据库失败: %v", err)
		}
	}
	return nil
}
