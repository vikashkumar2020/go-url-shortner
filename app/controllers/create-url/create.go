package create

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	model "github.com/vikashkumar2020/go-url-shortner/app/models"
	pgdatabase "github.com/vikashkumar2020/go-url-shortner/infra/postgres/database"
	utils "github.com/vikashkumar2020/go-url-shortner/utils"
)


func Create(c *gin.Context) {
	var url model.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Validate the URL
	if !isValidURL(url.LongURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Insert the long URL into the database and get the generated ID
	db := pgdatabase.GetDBInstance().GetDB()
	if err := db.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Convert the ID to a short URL
	url.ShortURL = fmt.Sprintf("http://%s:%s/api/v1/%s",os.Getenv("DB_HOST"),os.Getenv("APP_PORT"),utils.Base62Encode(int(url.ID)))

	// Update the record with the short URL
	if err := db.Save(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusCreated, url)
}

// Check if the URL is valid
func isValidURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}
