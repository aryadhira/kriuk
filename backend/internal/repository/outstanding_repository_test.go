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

func TestOutstandingRepoImp(t *testing.T) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5432", "kriuk_db")
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	defer db.Close()

	outstandingRepo := NewOutstandingRepoImp(db)

	// Test AddOutstanding
	newOutstanding := &models.OutstandingDeposit{
		ID:            "1",
		EmployeeID:    "1",
		TransactionID: "",
		Amount:        100.0,
		IsPaid:        false,
		Note:          "Test outstanding deposit",
		Date:          time.Now(),
		CreateOn:      time.Now(),
		UpdateOn:      time.Now(),
	}
	err = outstandingRepo.AddOutstanding(newOutstanding)
	assert.NoError(t, err)

	// Test GetOutstanding
	outstandings, err := outstandingRepo.GetOutstanding()
	require.NoError(t, err)
	assert.NotEmpty(t, outstandings)

	// Test GetOutstandingByEmployee
	employeeID := newOutstanding.EmployeeID
	outstandings, err = outstandingRepo.GetOutstandingByEmployee(employeeID)
	require.NoError(t, err)
	assert.NotEmpty(t, outstandings)

	// Test UpdateOutstanding
	newOutstanding.Amount = 200.0
	err = outstandingRepo.UpdateOutstanding(newOutstanding)
	assert.NoError(t, err)

}
