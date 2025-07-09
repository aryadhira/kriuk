package script

import "database/sql"

func CreateCashFlowTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS cashflow (
		id TEXT PRIMARY KEY,
		type TEXT NOT NULL CHECK (type IN ('in', 'out')),
		amount NUMERIC NOT NULL,
		note TEXT,
		date TIMESTAMP NOT NULL,
		createon TIMESTAMP NOT NULL,
		updateon TIMESTAMP NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
