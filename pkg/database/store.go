package database

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/go-redis/redis/v8"
)

// StoreManager manages different types of storage
type StoreManager struct {
	sqlDB    *sql.DB          // Database connection
	levelDB  *leveldb.DB      // LevelDB instance
	redisDB  *redis.Client    // Redis client
	mu       sync.RWMutex
}

// StoreConfig contains configuration options
type StoreConfig struct {
	DatabaseURL string
	LevelDBPath string
	RedisURL    string
}

func NewStoreManager(config *StoreConfig) (*StoreManager, error) {
	sm := &StoreManager{}
	
	// Initialize database connection
	sqlDB, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	sm.sqlDB = sqlDB

	// Initialize LevelDB
	levelDB, err := leveldb.OpenFile(config.LevelDBPath, nil)
	if err != nil {
		return nil, err
	}
	sm.levelDB = levelDB

	// Initialize Redis connection
	redisDB := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})
	sm.redisDB = redisDB

	return sm, nil
}

// ExecSQL executes SQL operations in a transaction
func (sm *StoreManager) ExecSQL(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := sm.sqlDB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit()
}

// PutBlock stores blockchain data
func (sm *StoreManager) PutBlock(key []byte, value []byte) error {
	return sm.levelDB.Put(key, value, nil)
}

func (sm *StoreManager) GetBlock(key []byte) ([]byte, error) {
	return sm.levelDB.Get(key, nil)
}

// SetCache handles hot data caching
func (sm *StoreManager) SetCache(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return sm.redisDB.Set(ctx, key, value, ttl).Err()
}

func (sm *StoreManager) GetCache(ctx context.Context, key string) (string, error) {
	return sm.redisDB.Get(ctx, key).Result()
}

// Close closes all connections
func (sm *StoreManager) Close() error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if err := sm.sqlDB.Close(); err != nil {
		return err
	}
	if err := sm.levelDB.Close(); err != nil {
		return err
	}
	if err := sm.redisDB.Close(); err != nil {
		return err
	}
	return nil
}
