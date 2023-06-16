package pgdatabase

import (
	config "github.com/vikashkumar2020/go-url-shortner/config"
	db "github.com/vikashkumar2020/go-url-shortner/infra/postgres"
	"github.com/vikashkumar2020/go-url-shortner/utils"

	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

// Database type structure
type Database struct {
	DB *gorm.DB
}

var databaseInstance *Database

// GetDB Using this function to get a connection, you can create your connection pool here.
func (database *Database) NewDBConnection(config *config.DBConfig) {
	postgresConnection := postgres.Open(db.GetOrmConfig(config))
	// change this to postgres connection
	db, err := gorm.Open(postgresConnection, &gorm.Config{})
	if err != nil {
		utils.LogFatal(err)
	}
	database.DB = db
}

// Singleton fun	ction
func GetDBInstance() *Database {
	if databaseInstance == nil {
		databaseInstance = &Database{}
	}
	return databaseInstance
}

// GetDB Using this function to get a connection, you can create your connection pool here.

func (database *Database) GetDB() *gorm.DB {
	return database.DB
}



