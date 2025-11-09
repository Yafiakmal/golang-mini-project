package main

import (
	"log"

	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/database"

	"github.com/yafiakmal/golang-mini-project/url-shortener/config"
)

func main() {
	config.EnvLoad()

	_, err := database.Connect(config.GetDBConfig())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// if len(os.Args) < 2 {
	// 	log.Fatalf("missing command (up, down, status, etc.)")
	// }
	// log.Println(os.Args)

	// dir := "./migrations"

	// // Jalankan perintah goose (up, down, status, dll)
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("failed to get sql.DB: %v", err)
	// }
	// if err := goose.RunContext(context.Background(), os.Args[1], sqlDB, dir); err != nil {
	// 	log.Fatalf("migration failed: %v", err)
	// }
}
