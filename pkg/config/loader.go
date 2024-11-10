package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
		Port    int    `yaml:"port"`
	} `yaml:"service"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`

	Security struct {
		JWTSecret           string        `yaml:"jwt_secret"`
		TokenExpiration     time.Duration `yaml:"token_expiration"`
		MaxLoginAttempts    int           `yaml:"max_login_attempts"`
		PasswordMinLength   int           `yaml:"password_min_length"`
		BiometricThreshold float64       `yaml:"biometric_threshold"`
	} `yaml:"security"`

	Monitoring struct {
		MetricsInterval      time.Duration `yaml:"metrics_interval"`
		AlertRetentionPeriod time.Duration `yaml:"alert_retention_period"`
		LogLevel            string        `yaml:"log_level"`
	} `yaml:"monitoring"`

	Consensus struct {
		Algorithm           string        `yaml:"algorithm"`
		BlockTime          time.Duration `yaml:"block_time"`
		MinValidators      int           `yaml:"min_validators"`
		MaxValidators      int           `yaml:"max_validators"`
	} `yaml:"consensus"`

	Network struct {
		P2PPort       int      `yaml:"p2p_port"`
		BootstrapNodes []string `yaml:"bootstrap_nodes"`
		MaxPeers      int      `yaml:"max_peers"`
	} `yaml:"network"`
}

func Load(serviceName string) (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	configPath := filepath.Join("configs", env, fmt.Sprintf("%s.yaml", serviceName))
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	// éªè¯éç½®
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	// åºç¨ç¯å¢åéè¦ç
	applyEnvironmentOverrides(config)

	return config, nil
}

func validateConfig(config *Config) error {
	if config.Service.Name == "" {
		return fmt.Errorf("service name is required")
	}

	if config.Service.Port <= 0 {
		return fmt.Errorf("invalid service port")
	}

	// æ·»å æ´å¤éªè¯...

	return nil
}

func applyEnvironmentOverrides(config *Config) {
	// ä»ç¯å¢åéè¦çéç½?
	if port := os.Getenv("SERVICE_PORT"); port != "" {
		// è½¬æ¢å¹¶è®¾ç½®ç«¯å?
	}

	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.Database.Host = dbHost
	}

	// æ·»å æ´å¤ç¯å¢åéè¦ç...
} 
