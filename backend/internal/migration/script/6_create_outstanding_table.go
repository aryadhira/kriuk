package script

import "database/sql"

func CreateOutstandingTable(db *sql.DB) error {
	query := `CREATE TABLE outstanding_deposits (
		id TEXT PRIMARY KEY,
		employeeid TEXT NOT NULL,
		transactionid TEXT, -- optional: nullable if payment not tied to a transaction
		amount NUMERIC(12, 2) NOT NULL, -- positive = debt, negative = payment
		note TEXT,
		ispaid BOOLEAN DEFAULT FALSE,
		date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
		createon TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
		updateon TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),

		-- Foreign key references
		CONSTRAINT fk_employee FOREIGN KEY (employeeid) REFERENCES employee(id),
		CONSTRAINT fk_transaction FOREIGN KEY (transactionid) REFERENCES transaction(id)
	);

	`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil

}
