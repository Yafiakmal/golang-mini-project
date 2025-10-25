package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/models"
	"gorm.io/gorm"
)

type UrlInput struct {
	URL      string `json:"url" binding:"required"`
	ShortUrl string `json:"short_url" binding:"required"`
}

func AddUrlHandler(c *gin.Context, db *gorm.DB) {
	var input UrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// add to database
	url := &models.Url{
		Url:      input.URL,
		ShortUrl: input.ShortUrl,
		UserID:   "user_1",
	}
	res := db.WithContext(context.Background()).Create(url)
	if res.Error != nil {
		c.JSON(http.StatusConflict, nil)
	}
	c.JSON(http.StatusOK, url)
}
