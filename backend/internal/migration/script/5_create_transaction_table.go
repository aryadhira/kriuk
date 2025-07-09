package script

import "database/sql"

func CreateTransactionTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS transaction (
		id TEXT PRIMARY KEY,
		employeename TEXT NOT NULL,
		deposit INTEGER NOT NULL,
		date TIMESTAMP NOT NULL,
		createon TIMESTAMP NOT NULL,
		updateon TIMESTAMP NOT NULL
	);
	CREATE TABLE transaction_items (
		transaction_id TEXT REFERENCES transaction(id) ON DELETE CASCADE,
		productname TEXT NOT NULL,
		total INTEGER NOT NULL
	)
	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
