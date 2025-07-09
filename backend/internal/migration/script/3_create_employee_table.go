package script

import "database/sql"

func CreateEmployeeTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS employee (
		id TEXT PRIMARY KEY,
  		name TEXT NOT NULL

	)`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
