package db

import (
	"fmt"
	"log"
	"price-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s search_path=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSchema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get raw DB: %v", err)
	}

	// Проверим подключение
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database unreachable: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return db
}
