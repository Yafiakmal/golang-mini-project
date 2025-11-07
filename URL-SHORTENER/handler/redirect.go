package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yafiakmal/golang-mini-project/url-shortener/models"
	"gorm.io/gorm"
)

func Redirect(c *gin.Context, db *gorm.DB) {
	// get path param
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	// search is param exist
	urls := &models.Url{}
	res := db.WithContext(context.Background()).Where("short_url = ?", name).First(urls)
	if res.Error != nil {
		log.Println(res.Error)
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	// redirect user to the link
	c.Redirect(http.StatusMovedPermanently, urls.Url)
}
