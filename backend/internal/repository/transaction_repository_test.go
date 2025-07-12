package repository

import (
	"database/sql"
	"fmt"
	"kriuk/internal/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransactionRepoImp(t *testing.T) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5432", "kriuk_db")
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	defer db.Close()

	transactionRepo := NewTransactionRepoImp(db)

	// Test AddTransaction
	newTransaction := &models.Transaction{
		ID:           uuid.NewString(),
		EmployeeName: "Test Employee",
		Deposit:      100.0,
		Date:         time.Now(),
		CreateOn:     time.Now(),
		UpdateOn:     time.Now(),
		ProductItems: []models.ProductItem{
			{ProductName: "Product 1", Total: 50.0},
			{ProductName: "Product 2", Total: 30.0},
		},
	}
	err = transactionRepo.AddTransaction(newTransaction)
	assert.NoError(t, err)

	// Test GetTransactions
	transactions, err := transactionRepo.GetTransactions()
	require.NoError(t, err)
	assert.NotEmpty(t, transactions)

	// Test GetTransactionByDate
	date := time.Now()
	transactionByDate, err := transactionRepo.GetTransactionByDate(date)
	require.NoError(t, err)
	assert.NotNil(t, transactionByDate)

	// Test UpdateTransaction
	newTransaction.Deposit = 200.0
	err = transactionRepo.UpdateTransaction(newTransaction)
	assert.NoError(t, err)

	// Test DeleteTransaction
	err = transactionRepo.DeleteTransaction(newTransaction.ID)
	assert.NoError(t, err)
}
