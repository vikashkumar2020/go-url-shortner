package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	utils "github.com/vikashkumar2020/go-url-shortner/utils"
)

func create(c *gin.Context) {
	var url URL
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
	var id int
	err := db.QueryRow("INSERT INTO urls(long_url) VALUES($1) RETURNING id", url.LongURL).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Convert the ID to a short URL
	url.ShortURL = utils.Base62Encode(id)

	// Update the record with the short URL
	_, err = db.Exec("UPDATE urls SET short_url = $1 WHERE id = $2", url.ShortURL, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, url)
}

// Check if the URL is valid
func isValidURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}