package config

import (
	"os"
	"time"
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

// NewConnection creates a new database connection
// func NewSqlConnection(config Config) (*sql.DB, error) {
// 	dsn := fmt.Sprintf(
// 		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
// 		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode, config.TimeZone,
// 	)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open database: %w", err)
// 	}

// 	// Test the connection
// 	if err := db.Ping(); err != nil {
// 		return nil, fmt.Errorf("failed to ping database: %w", err)
// 	}

// 	log.Println("Successfully connected to database")
// 	return db, nil
// }
