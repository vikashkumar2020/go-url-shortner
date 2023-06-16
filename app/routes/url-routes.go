package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/go-url-shortner/app/controllers/create-url"
	"github.com/vikashkumar2020/go-url-shortner/app/controllers/short-url"
)

func RegisterShortUrlRoutes(router *gin.RouterGroup) {

	router.GET("/:shortURL",short.Redirect )
    router.POST("/create", create.Create)

}