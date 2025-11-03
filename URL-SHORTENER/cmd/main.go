package main

import (
	"log"
	"time"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yafiakmal/golang-mini-project/url-shortener/config"
	"github.com/yafiakmal/golang-mini-project/url-shortener/handler"
)

func main() {
	godotenv.Load()

	// initialize postgres
	db, err := config.NewGormConnection(config.GetDBConfig())
	if err != nil {
		log.Panic(err)
	}

	server := gin.Default()

 server.Use(cors.New(cors.Config{
    AllowOrigins:  []string{"*"}, // atau ganti dengan domain frontend kamu
    AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
    ExposeHeaders: []string{"Content-Length"},
    MaxAge:        12 * time.Hour,
  }))

	server.POST("/shortener", func(ctx *gin.Context) {
		handler.AddUrlHandler(ctx, db)
	})
	server.GET("/urls", func(ctx *gin.Context) {
		handler.GetAllUrlHandler(ctx, db)
	})
	server.GET("/:name", func(ctx *gin.Context) {
		handler.Redirect(ctx, db)
	})
	server.Run(":80")
}
