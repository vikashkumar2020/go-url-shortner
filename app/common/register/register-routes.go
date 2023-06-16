package register

import (
	"net/http"
	"github.com/gin-gonic/gin"
	config "github.com/vikashkumar2020/go-url-shortner/config"
)

// health ckeck api
func healhCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health check",
	})
}

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func Routes(router *gin.Engine, serverConfig *config.ServerConfig) {
	router.Use(CORSMiddleware())
	webV1RouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1)
	RegiterWebRoutes(webV1RouterGroup)
	router.GET("/health", healhCheck)

}

// rest api routes 
func RegiterWebRoutes(router *gin.RouterGroup) {
	
}