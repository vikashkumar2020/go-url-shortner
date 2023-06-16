package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/go-url-shortner/config"
	"github.com/vikashkumar2020/go-url-shortner/utils"
	pgdatabase "github.com/vikashkumar2020/go-url-shortner/infra/postgres/database"
	register "github.com/vikashkumar2020/go-url-shortner/app/common/register"
)


func main() {

	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	utils.LogInfo("server config loaded")

	// Initialize the database
	dbConfig := config.NewDBConfig()
	utils.LogInfo("db config loaded")

	// initialize database
	database := pgdatabase.GetDBInstance();
	database.NewDBConnection(dbConfig);
	utils.LogInfo("db connection established")

	router := gin.Default()
	register.Routes(router, serverConfig)
	
	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("server started")
}