package app

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port          string
	DatabaseDSN   string
	JWTSecret     string
	AdminUsername string
	AdminPassword string
	UploadDir     string
	PublicDir     string
	PublicBaseURL string
	MaxUploadSize int64
	TokenTTL      time.Duration
}

func LoadConfig() Config {
	return Config{
		Port:          getEnv("PORT", "8080"),
		DatabaseDSN:   getEnv("DATABASE_DSN", "byml:byml_password@tcp(127.0.0.1:3306)/byml?charset=utf8mb4&parseTime=True&loc=Local"),
		JWTSecret:     getEnv("JWT_SECRET", "change-me-in-production"),
		AdminUsername: getEnv("ADMIN_USERNAME", "admin"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123456"),
		UploadDir:     getEnv("UPLOAD_DIR", "/data/uploads"),
		PublicDir:     getEnv("PUBLIC_DIR", "/app/public"),
		PublicBaseURL: getEnv("PUBLIC_BASE_URL", ""),
		MaxUploadSize: getEnvInt64("MAX_UPLOAD_SIZE", 100<<20),
		TokenTTL:      time.Duration(getEnvInt64("TOKEN_TTL_HOURS", 72)) * time.Hour,
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid integer env %s=%q", key, value))
	}
	return parsed
}
