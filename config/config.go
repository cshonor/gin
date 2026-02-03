package config

import (
	"os"
)

// Config 应用配置
type Config struct {
	Port         string
	DatabaseURL  string // SQLite: "gin.db" 或完整路径, MySQL: "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local", PostgreSQL: "host=localhost user=postgres password=postgres dbname=gin port=5432 sslmode=disable"
	JWTSecret    string
	Environment  string // development 或 production
}

// Load 加载配置
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

