package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/internal/models"
	"gorm.io/gorm"
)

func DeleteUrl(c *gin.Context, db *gorm.DB) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}
	urls := &models.Url{}
	res := db.WithContext(context.Background()).Where("name = ?", name).Delete(urls)
	if res.RowsAffected == 0 {
		log.Println("DATABASE ERROR: ", "Entity not found")
		c.JSON(http.StatusConflict, gin.H{"error": "Entity not found"})
	}
	if res.Error != nil {
		log.Println("DATABASE ERROR: ", res.Error)
		c.JSON(http.StatusConflict, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}
