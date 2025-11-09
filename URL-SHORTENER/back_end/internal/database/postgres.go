package database

import (
	"fmt"

	"github.com/yafiakmal/golang-mini-project/url-shortener/config"
	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbConfig *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password,
		dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // ✅ ini cara yang benar
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}
	/*
		defer sqlDB.Close()
	*/

	poolConfig := config.GetPoolConfig()
	sqlDB.SetMaxIdleConns(poolConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(poolConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(poolConfig.ConnMaxLifetime)

	if config.APP_ENV == "development" {
		if err := models.AutoMigrate(db); err != nil {
			return nil, fmt.Errorf("failed to migrate: %w", err)
		}
	}

	return db, nil
}
