package main

import (
	"log"

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

	server.POST("/shortener", func(ctx *gin.Context) {
		handler.AddUrlHandler(ctx, db)
	})
	server.GET("/urls", func(ctx *gin.Context) {
		handler.GetAllUrlHandler(ctx, db)
	})
	server.GET("/:name", func(ctx *gin.Context) {
		handler.Redirect(ctx, db)
	})
	server.Run(":8080")
}
