package short

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	model "github.com/vikashkumar2020/go-url-shortner/app/models"
	pgdatabase "github.com/vikashkumar2020/go-url-shortner/infra/postgres/database"
	utils "github.com/vikashkumar2020/go-url-shortner/utils"
	"gorm.io/gorm"
)

var cache = &sync.Map{}

func Redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	id := utils.Base62Decode(shortURL)

	// Try to get the long URL from the cache
	if longURL, ok := cache.Load(shortURL); ok {
		c.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}

	// If it's not in the cache, get it from the database
	var url model.URL
	db := pgdatabase.GetDBInstance().GetDB()
	if err := db.First(&url, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Store the long URL in the cache
	cache.Store(shortURL, url.LongURL)

	c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
