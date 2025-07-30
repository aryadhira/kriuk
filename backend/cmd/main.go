package main

import (
	"kriuk/internal/database"
	"kriuk/internal/migration"
	"kriuk/internal/repository"
	"kriuk/internal/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	// initiate DB
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// start migration
	dbMigration := migration.NewDBMigration(db)
	err = dbMigration.StartMigration()
	if err != nil {
		log.Fatal(err)
	}

	// starting service
	stockRepo := repository.NewStockRepoImp(db)
	stock := services.NewStockSvc(stockRepo)

	employeeRepo := repository.NewEmployeeRepoImp(db)
	employee := services.NewEmployeeSvc(employeeRepo)

	cashflowRepo := repository.NewCashFlowRepoImp(db)
	cashflow := services.NewCashflowSvc(cashflowRepo)

	transactionRepo := repository.NewTransactionRepoImp(db)
	outstandingRepo := repository.NewOutstandingRepoImp(db)
	transaction := services.NewTransactionSvc(transactionRepo,employeeRepo,stockRepo,outstandingRepo)

	kriuk := services.NewKriuk(stock, employee, cashflow, transaction)

	err = kriuk.Start()
	if err != nil {
		log.Fatal("error starting kriuk service: ",err)
	}
}
