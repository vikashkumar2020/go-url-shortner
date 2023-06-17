package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upName, downName)
}

func upUpdateTable(tx *sql.Tx) error {
	_, err := tx.Exec("ALTER SEQUENCE urls_id_seq RESTART WITH 5000")
	if err != nil {
		return fmt.Errorf("failed to restart sequence: %w", err)
	}
	return nil
}

func downTableUpdate(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
