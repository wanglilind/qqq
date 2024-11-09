package dao

import (
	"context"
	"time"

	"github.com/wanglilind/qqq/pkg/database"
)

// DAOæ¥å£å®ä¹
type DAO interface {
	IdentityDAO
	AccountDAO
	TransactionDAO
	SystemDAO
}

// èº«ä»½ç¸å³DAOæ¥å£
type IdentityDAO interface {
	CreateIdentity(ctx context.Context, identity *Identity) error
	GetIdentity(ctx context.Context, id string) (*Identity, error)
	UpdateIdentity(ctx context.Context, identity *Identity) error
	VerifyIdentity(ctx context.Context, id string, biometricData []byte) (bool, error)
	ListIdentities(ctx context.Context, filter *IdentityFilter) ([]*Identity, error)
}

// è´¦æ·ç¸å³DAOæ¥å£
type AccountDAO interface {
	CreateAccount(ctx context.Context, account *Account) error
	GetAccount(ctx context.Context, id string) (*Account, error)
	UpdateBalance(ctx context.Context, id string, amount float64) error
	GetAccountHistory(ctx context.Context, id string, startTime, endTime time.Time) ([]*AccountActivity, error)
	ListAccounts(ctx context.Context, filter *AccountFilter) ([]*Account, error)
}

// äº¤æç¸å³DAOæ¥å£
type TransactionDAO interface {
	CreateTransaction(ctx context.Context, tx *Transaction) error
	GetTransaction(ctx context.Context, id string) (*Transaction, error)
	UpdateTransactionStatus(ctx context.Context, id string, status string) error
	ListTransactions(ctx context.Context, filter *TransactionFilter) ([]*Transaction, error)
	GetTransactionStats(ctx context.Context, accountId string) (*TransactionStatistics, error)
}

// ç³»ç»åæ°ç¸å³DAOæ¥å£
type SystemDAO interface {
	GetParameter(ctx context.Context, key string) (string, error)
	SetParameter(ctx context.Context, key string, value string) error
	ListParameters(ctx context.Context) (map[string]string, error)
}

// DAOå®ç°
type daoImpl struct {
	store *database.StoreManager
}

func NewDAO(store *database.StoreManager) DAO {
	return &daoImpl{
		store: store,
	}
}

// èº«ä»½DAOå®ç°
func (d *daoImpl) CreateIdentity(ctx context.Context, identity *Identity) error {
	return d.store.ExecSQL(ctx, func(tx *database.Tx) error {
		query := `
			INSERT INTO identities (
				id, user_id, biometric_hash, national_id, country_code, 
				birth_date, status, verification_count, metadata
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`
		_, err := tx.Exec(ctx, query,
			identity.ID, identity.UserID, identity.BiometricHash,
			identity.NationalID, identity.CountryCode, identity.BirthDate,
			identity.Status, identity.VerificationCount, identity.Metadata,
		)
		return err
	})
}

// è´¦æ·DAOå®ç°
func (d *daoImpl) UpdateBalance(ctx context.Context, id string, amount float64) error {
	return d.store.ExecSQL(ctx, func(tx *database.Tx) error {
		// å¼å§äºå?
		query := `
			UPDATE currency_accounts 
			SET current_balance = current_balance + $1,
				last_decay_date = CURRENT_TIMESTAMP
			WHERE id = $2
		`
		result, err := tx.Exec(ctx, query, amount, id)
		if err != nil {
			return err
		}

		// è®°å½è´¦æ·æ´»å¨
		activityQuery := `
			INSERT INTO account_activities (
				id, account_id, activity_type, amount, 
				balance_before, balance_after, timestamp
			) VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP)
		`
		_, err = tx.Exec(ctx, activityQuery,
			NewUUID(), id, "BALANCE_UPDATE", amount,
			0, 0, // è¿ééè¦è·åå®éçä½é¢åå
		)
		return err
	})
}

// äº¤æDAOå®ç°
func (d *daoImpl) CreateTransaction(ctx context.Context, tx *Transaction) error {
	return d.store.ExecSQL(ctx, func(dbTx *database.Tx) error {
		query := `
			INSERT INTO transactions (
				id, sender_id, recipient_id, amount, 
				transaction_type, status, signature
			) VALUES ($1, $2, $3, $4, $5, $6, $7)
		`
		_, err := dbTx.Exec(ctx, query,
			tx.ID, tx.SenderID, tx.RecipientID,
			tx.Amount, tx.Type, tx.Status, tx.Signature,
		)
		if err != nil {
			return err
		}

		// æ´æ°ç»è®¡æ°æ®
		statsQuery := `
			INSERT INTO transaction_statistics (
				id, account_id, period_start, period_end,
				total_transactions, total_amount
			) VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (account_id, period_start) DO UPDATE
			SET total_transactions = transaction_statistics.total_transactions + 1,
				total_amount = transaction_statistics.total_amount + EXCLUDED.total_amount
		`
		_, err = dbTx.Exec(ctx, statsQuery,
			NewUUID(), tx.SenderID,
			time.Now().Truncate(24*time.Hour),
			time.Now().Truncate(24*time.Hour).Add(24*time.Hour),
			1, tx.Amount,
		)
		return err
	})
}

// ç³»ç»åæ°DAOå®ç°
func (d *daoImpl) GetParameter(ctx context.Context, key string) (string, error) {
	var value string
	err := d.store.ExecSQL(ctx, func(tx *database.Tx) error {
		return tx.QueryRow(ctx,
			"SELECT parameter_value FROM system_parameters WHERE parameter_key = $1",
			key,
		).Scan(&value)
	})
	return value, err
} 
