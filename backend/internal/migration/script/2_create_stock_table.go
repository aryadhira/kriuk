package script

import (
	"database/sql"
)

func CreateStockTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS stock(
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		qty INTEGER DEFAULT 0,
		unit TEXT,
		price NUMERIC,
		createon TIMESTAMP,
		updateon TIMESTAMP

	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
