package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/models"
	"gorm.io/gorm"
)

type UrlInput struct {
	URL  string `json:"url" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func AddUrlHandler(c *gin.Context, db *gorm.DB) {
	var input UrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("BINDING ERROR: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// add to database
	url := &models.Url{
		Url:    input.URL,
		Name:   input.Name,
		UserID: 1,
	}
	res := db.WithContext(context.Background()).Create(url)
	if res.Error != nil {
		log.Println("DATABASE ERROR: ", res.Error)
		c.JSON(http.StatusConflict, gin.H{"error": "failed to create short url, name might be already used"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully", "data": url})
}
