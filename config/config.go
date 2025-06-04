package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSchema   string
	JWTSecret  string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	return Config{
		Port:       getEnv("PORT", "8082"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("PG_USER", "postgres"),
		DBPassword: getEnv("PG_PASSWORD", ""),
		DBName:     getEnv("PG_DB", "price_monitor"),
		DBSchema:   getEnv("PG_SCHEMA", "public"),
		JWTSecret:  getEnv("JWT_SECRET", "supersecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
