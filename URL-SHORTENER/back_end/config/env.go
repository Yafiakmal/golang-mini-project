package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_ENV     = os.Getenv("APP_ENV")
	DB_HOST     = ""
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_PORT     = ""
	DB_SSLMODE  = ""
	DB_TIMEZONE = ""
)

func EnvLoad() {
	if APP_ENV == "" {
		APP_ENV = "development"
	}

	filename := fmt.Sprintf(".env.%s", APP_ENV)

	err := godotenv.Load(filename)
	if err != nil {
		msg := filename + ` NOT FOUND`
		panic(msg)
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SSLMODE = os.Getenv("DB_SSLMODE")
	DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	if DB_HOST == "" || DB_USER == "" || DB_PASSWORD == "" || DB_NAME == "" || DB_PORT == "" || DB_SSLMODE == "" || DB_TIMEZONE == "" {
		panic("Database environment variables are not set properly")
	}

}
