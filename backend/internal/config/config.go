package config

import (
	"os"
)

type Config struct {
	Port          string
	DatabaseURL   string
	RedisURL      string
	JWTSecret     string
	KYCLabsAPIKey string
	BlockchainRPC string
	PrivateKey    string
	Environment   string
}

func Load() (*Config, error) {
	return &Config{
		Port:          getEnv("PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://user:password@localhost/notabot?sslmode=disable"),
		RedisURL:      getEnv("REDIS_URL", "redis://localhost:6379"),
		JWTSecret:     getEnv("JWT_SECRET", "change-me-in-production"),
		KYCLabsAPIKey: getEnv("KYC_LABS_API_KEY", ""),
		BlockchainRPC: getEnv("BLOCKCHAIN_RPC", ""),
		PrivateKey:    getEnv("PRIVATE_KEY", ""),
		Environment:   getEnv("ENVIRONMENT", "development"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
