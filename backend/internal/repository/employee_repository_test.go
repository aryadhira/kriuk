package repository

import (
	"database/sql"
	"fmt"
	"kriuk/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func TestEmployeeRepoImp(t *testing.T) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "password", "localhost", "5432", "kriuk_db")
	db, err := sql.Open("postgres", connString)
	require.NoError(t, err)
	defer db.Close()

	employeeRepo := NewEmployeeRepoImp(db)

	// Test AddEmployee
	newEmployee := &models.Employee{
		ID:   "1",
		Name: "John Doe",
	}
	err = employeeRepo.AddEmployee(newEmployee)
	assert.NoError(t, err)

	// Test GetEmployees
	employees, err := employeeRepo.GetEmployees()
	require.NoError(t, err)
	assert.NotEmpty(t, employees)
	assert.Contains(t, employees, newEmployee)

	// Test GetEmployeeByName
	employee, err := employeeRepo.GetEmployeeByName("John Doe")
	require.NoError(t, err)
	assert.NotNil(t, employee)
	assert.Equal(t, newEmployee.ID, employee.ID)
	assert.Equal(t, newEmployee.Name, employee.Name)

}
