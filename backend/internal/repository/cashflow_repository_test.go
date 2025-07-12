package repository

import (
	"database/sql"
	"fmt"
	"kriuk/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func TestCashFlowRepoImp(t *testing.T) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5432", "kriuk_db")
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	defer db.Close()

	cashFlowRepo := NewCashFlowRepoImp(db)

	// Test AddCashFlow
	newCashFlow := &models.CashFlow{
		ID:       "1",
		Type:     "in",
		Amount:   100.0,
		Note:     "Test cash flow",
		Date:     time.Now(),
		CreateOn: time.Now(),
		UpdateOn: time.Now(),
	}
	err = cashFlowRepo.AddCashFlow(newCashFlow)
	assert.NoError(t, err)

	// Test GetCashFlow
	cashFlows, err := cashFlowRepo.GetCashFlow()
	require.NoError(t, err)
	assert.NotEmpty(t, cashFlows)

	// Test GetCashFlowByDate
	date := newCashFlow.Date
	cashFlow, err := cashFlowRepo.GetCashFlowByDate(date)
	require.NoError(t, err)
	assert.NotNil(t, cashFlow)

	// Test UpdateCashFlow
	newCashFlow.Amount = 200.0
	err = cashFlowRepo.UpdateCashFlow(newCashFlow)
	assert.NoError(t, err)

	// Test DeleteCashFlow
	err = cashFlowRepo.DeleteCashFlow(newCashFlow.ID)
	assert.NoError(t, err)
}
