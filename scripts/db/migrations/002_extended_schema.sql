-- 用户身份表扩展
ALTER TABLE identities 
ADD COLUMN last_verification_time TIMESTAMP WITH TIME ZONE,
ADD COLUMN verification_method VARCHAR(50),
ADD COLUMN risk_score DECIMAL(5,2),
ADD COLUMN metadata JSONB;

-- 生物特征历史表
CREATE TABLE biometric_history (
    id UUID PRIMARY KEY,
    identity_id UUID REFERENCES identities(id),
    biometric_type VARCHAR(50) NOT NULL,
    biometric_hash VARCHAR(512) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    verified_at TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50),
    confidence_score DECIMAL(5,2)
);

-- 账户活动表
CREATE TABLE account_activities (
    id UUID PRIMARY KEY,
    account_id UUID REFERENCES currency_accounts(id),
    activity_type VARCHAR(50) NOT NULL,
    amount DECIMAL(20,8),
    balance_before DECIMAL(20,8),
    balance_after DECIMAL(20,8),
    decay_rate DECIMAL(10,8),
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    metadata JSONB
);

-- 交易统计表
CREATE TABLE transaction_statistics (
    id UUID PRIMARY KEY,
    account_id UUID REFERENCES currency_accounts(id),
    period_start TIMESTAMP WITH TIME ZONE,
    period_end TIMESTAMP WITH TIME ZONE,
    total_transactions INTEGER,
    total_amount DECIMAL(20,8),
    average_amount DECIMAL(20,8),
    max_amount DECIMAL(20,8),
    transaction_frequency DECIMAL(10,2),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 系统参数表
CREATE TABLE system_parameters (
    id UUID PRIMARY KEY,
    parameter_key VARCHAR(100) UNIQUE NOT NULL,
    parameter_value TEXT NOT NULL,
    description TEXT,
    last_modified TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_by VARCHAR(255),
    version INTEGER DEFAULT 1,
    status VARCHAR(50) DEFAULT 'ACTIVE'
);

-- 索引
CREATE INDEX idx_biometric_history_identity ON biometric_history(identity_id);
CREATE INDEX idx_account_activities_account ON account_activities(account_id);
CREATE INDEX idx_account_activities_type ON account_activities(activity_type);
CREATE INDEX idx_transaction_stats_account ON transaction_statistics(account_id);
CREATE INDEX idx_transaction_stats_period ON transaction_statistics(period_start, period_end); 