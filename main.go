package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// r.GET("/:shortURL", redirect)
	// r.POST("/create", create)
	r.Run()
}