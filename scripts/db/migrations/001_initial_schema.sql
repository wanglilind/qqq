-- 升级脚本
-- +migrate Up
-- 用户身份表
CREATE TABLE identities (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) UNIQUE NOT NULL,
    biometric_hash VARCHAR(512) NOT NULL,
    national_id VARCHAR(255) UNIQUE,
    country_code VARCHAR(3) NOT NULL,
    birth_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'ACTIVE',
    verification_count INT DEFAULT 0
);

-- 货币账户表
CREATE TABLE currency_accounts (
    id UUID PRIMARY KEY,
    identity_id UUID REFERENCES identities(id),
    initial_balance DECIMAL(20,8) NOT NULL,
    current_balance DECIMAL(20,8) NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_decay_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) DEFAULT 'ACTIVE',
    CONSTRAINT positive_balance CHECK (current_balance >= 0)
);

-- 交易历史表
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    sender_id UUID REFERENCES currency_accounts(id),
    recipient_id UUID REFERENCES currency_accounts(id),
    amount DECIMAL(20,8) NOT NULL,
    transaction_type VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    confirmed_at TIMESTAMP WITH TIME ZONE,
    block_height BIGINT,
    signature VARCHAR(512) NOT NULL
);

-- 创建索引
CREATE INDEX idx_identities_user_id ON identities(user_id);
CREATE INDEX idx_identities_national_id ON identities(national_id);
CREATE INDEX idx_transactions_sender ON transactions(sender_id);
CREATE INDEX idx_transactions_recipient ON transactions(recipient_id);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);

-- 回滚脚本
-- +migrate Down
DROP TABLE transactions;
DROP TABLE currency_accounts;
DROP TABLE identities; 