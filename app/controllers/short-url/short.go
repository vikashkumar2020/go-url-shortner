package controllers

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	utils "github.com/vikashkumar2020/go-url-shortner/utils"
)

var cache = &sync.Map{}

func redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	id := utils.Base62Decode(shortURL)
	// Try to get the long URL from the cache
	if longURL, ok := cache.Load(shortURL); ok {
		c.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}
	// If it's not in the cache, get it from the database
	var url URL
	err := db.Get(&url, "SELECT long_url FROM urls WHERE id=$1", id)
	if err != nil {
		if err == sql.ErrNoRows {
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