package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yafiakmal/golang-mini-project/url-shortener/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
	TimeZone string
}

func GetDBConfig() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}
}

type PoolConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func GetPoolConfig() PoolConfig {
	return PoolConfig{
		MaxIdleConns:    5,
		MaxOpenConns:    20,
		ConnMaxLifetime: 30 * time.Minute,
	}
}

// postgresql://postgres:[YOUR_PASSWORD]@db.rtocusuptglvetwispqn.supabase.co:5432/postgres
func NewGormConnection(dbConfig *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password,
		dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	poolConfig := GetPoolConfig()
	sqlDB.SetMaxIdleConns(poolConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(poolConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(poolConfig.ConnMaxLifetime)

	if err := models.AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}

// NewConnection creates a new database connection
func NewSqlConnection(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode, config.TimeZone,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database")
	return db, nil
}
