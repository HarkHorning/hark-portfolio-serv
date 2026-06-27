package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Environment string
	Server      ServerConfig
	Database    DatabaseConfig
	Admin       AdminConfig
	Storage     StorageConfig
}

type AdminConfig struct {
	Username      string
	Password      string
	SessionSecret string
}

type StorageConfig struct {
	LocalStoragePath string
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	SeedData        bool // <--- ADD THIS LINE
}

func Load() Config {
	return Config{
		Environment: getEnv("ENVIRONMENT", "production"),
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "mysql"),
			Port:            getEnvInt("DB_PORT", 3306),
			User:            getEnv("DB_USER", "root"),
			Password:        getEnv("DB_PASSWORD", "your_secure_password"),
			Database:        getEnv("DB_NAME", "portfolio"),
			MaxOpenConns:    25,
			MaxIdleConns:    5,
			ConnMaxLifetime: 5 * time.Minute,
			SeedData:        getEnvBool("DB_SEED_DATA", false), // <--- ADD THIS LINE
		},
		Admin: AdminConfig{
			Username:      getEnv("ADMIN_USERNAME", "admin"),
			Password:      getEnv("ADMIN_PASSWORD", "changeme"),
			SessionSecret: getEnv("ADMIN_SESSION_SECRET", "very-secret-string"),
		},
		Storage: StorageConfig{
			LocalStoragePath: getEnv("STORAGE_PATH", "/app/data/images"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

// ADD THIS HELPER FUNCTION
func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	return fallback
}
