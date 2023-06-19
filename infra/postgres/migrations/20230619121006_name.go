package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
	model "github.com/vikashkumar2020/go-url-shortner/app/models"
	pgdatabase "github.com/vikashkumar2020/go-url-shortner/infra/postgres/database"
	config "github.com/vikashkumar2020/go-url-shortner/config"
	"gorm.io/gorm"
)

var db *gorm.DB
func init() {
		config.LoadEnv()
		config := config.NewDBConfig()
		database := pgdatabase.Database{}
		database.NewDBConnection(config)
		db = database.DB
	goose.AddMigration(upName, downName)
}

func upName(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return db.Migrator().CreateTable(&model.URL{})
	// return nil
}

func downName(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return db.Migrator().CreateTable(&model.URL{})
}
