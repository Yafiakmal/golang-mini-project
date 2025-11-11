package setup

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_HOST     string = ""
	DB_PORT     string = ""
	DB_USER     string = ""
	DB_PASSWORD string = ""
	DB_DBNAME   string = ""
	DB_SSLMODE  string = ""
	DB_TIMEZONE string = ""
)

func LoadEnv() {
	// load environment variables from .env file or system
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	err := godotenv.Load(".env." + env)
	if err != nil {

		log.Fatalf("Error loading .env.%s file", env)
	}
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DBNAME = os.Getenv("DB_DBNAME")
	DB_SSLMODE = os.Getenv("DB_SSLMODE")
	DB_TIMEZONE = os.Getenv("DB_TIMEZONE")
	if DB_HOST == "" || DB_PORT == "" || DB_USER == "" || DB_PASSWORD == "" || DB_DBNAME == "" || DB_SSLMODE == "" || DB_TIMEZONE == "" {
		log.Fatal("One or more required environment variables are missing")
	}

}
