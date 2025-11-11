package config

import "github.com/yafiakmal/golang-mini-project/blog/setup"

func GetPostgresURL() string {
	// "postgres://user:password@localhost:5432/mydb?sslmode=disable"
	dsn := "postgres://" +
		setup.DB_USER + ":" +
		setup.DB_PASSWORD + "@" +
		setup.DB_HOST + ":" +
		setup.DB_PORT + "/" +
		setup.DB_DBNAME +
		"?sslmode=" + setup.DB_SSLMODE +
		"&TimeZone=" + setup.DB_TIMEZONE
	return dsn
}
