package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/models"
	"gorm.io/gorm"
)

func DeleteUrl(c *gin.Context, db *gorm.DB) {
	url := c.Param("name")
	urls := &models.Url{}
	res := db.WithContext(context.Background()).Where("short_url = ?", url).Delete(urls)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusConflict, gin.H{"Error": "Entity not found"})
	}
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusConflict, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}
