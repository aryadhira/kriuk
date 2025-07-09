package main

import (
	"kriuk/internal/database"
	"kriuk/internal/migration"
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
}
