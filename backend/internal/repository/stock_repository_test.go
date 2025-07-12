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

func TestStockRepoImp(t *testing.T) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5432", "kriuk_db")
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	defer db.Close()

	stockRepo := NewStockRepoImp(db)

	// Test AddStock
	newStock := &models.Stocks{
		ID:       uuid.NewString(),
		Name:     "Test Stock",
		Quantity: 10,
		Unit:     "piece",
		Price:    5.0,
		CreateOn: time.Now(),
		UpdateOn: time.Now(),
	}
	err = stockRepo.AddStock(newStock)
	assert.NoError(t, err)

	// Test GetStocks
	stocks, err := stockRepo.GetStocks()
	require.NoError(t, err)
	assert.NotEmpty(t, stocks)

	// Test GetStocksByName
	stockByName, err := stockRepo.GetStocksByName("Test Stock")
	require.NoError(t, err)
	assert.NotNil(t, stockByName)

	// Test UpdateStock
	newStock.Quantity = 20
	err = stockRepo.UpdateStock(stockByName)
	assert.NoError(t, err)
}
