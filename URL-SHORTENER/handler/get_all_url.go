package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/models"
	"gorm.io/gorm"
)

func GetAllUrlHandler(c *gin.Context, db *gorm.DB) {
	// add to database
	url := &[]models.Url{}
	res := db.WithContext(context.Background()).Find(url)
	if res.Error != nil {
		c.JSON(http.StatusConflict, nil)
	}
	c.JSON(http.StatusOK, url)
}
