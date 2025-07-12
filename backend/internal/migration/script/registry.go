package script

import (
	"database/sql"
)

type MigrationScript struct {
	Version int
	Migrate func(db *sql.DB) error
}

var Migrations = []MigrationScript{
	{Version: 2, Migrate: CreateStockTable},
	{Version: 3, Migrate: CreateEmployeeTable},
	{Version: 4, Migrate: CreateCashFlowTable},
	{Version: 5, Migrate: CreateTransactionTable},
	{Version: 6, Migrate: CreateOutstandingTable},
}
