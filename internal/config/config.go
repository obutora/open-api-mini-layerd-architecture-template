package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config はアプリケーション設定を保持します
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
	Log      LogConfig
}

// ServerConfig はサーバー関連の設定を保持します
type ServerConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

// DatabaseConfig はデータベース関連の設定を保持します
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// AuthConfig は認証関連の設定を保持します
type AuthConfig struct {
	JWTSecret     string
	TokenDuration int // 分単位
}

// LogConfig はログ関連の設定を保持します
type LogConfig struct {
	Level string
	File  string
}

// DSN はデータベース接続文字列を返します
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

// Load は環境変数から設定を読み込みます
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvAsInt("SERVER_READ_TIMEOUT", 10),
			WriteTimeout: getEnvAsInt("SERVER_WRITE_TIMEOUT", 10),
		},
		Database: DatabaseConfig{
			Driver:   getEnv("DB_DRIVER", "mysql"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 3306),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "myapp"),
		},
		Auth: AuthConfig{
			JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
			TokenDuration: getEnvAsInt("TOKEN_DURATION", 60),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "info"),
			File:  getEnv("LOG_FILE", ""),
		},
	}
}

// getEnv は環境変数を取得し、設定されていない場合はデフォルト値を返します
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt は環境変数を整数として取得し、設定されていない場合はデフォルト値を返します
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
