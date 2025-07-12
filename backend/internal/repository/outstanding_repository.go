package repository

import (
	"database/sql"
	"kriuk/internal/models"
)

// OutstandingRepo interface for outstanding deposit operations.
type OutstandingRepo interface {
	GetOutstanding() ([]*models.OutstandingDeposit, error)
	GetOutstandingByEmployee(employeeID string) ([]*models.OutstandingDeposit, error)
	AddOutstanding(outstanding *models.OutstandingDeposit) error
	UpdateOutstanding(outstanding *models.OutstandingDeposit) error
}

// OutstandingRepoImp struct for OutstandingRepo implementation.
type OutstandingRepoImp struct {
	Db *sql.DB
}

// NewOutstandingRepoImp function for creating a new instance of OutstandingRepoImp.
func NewOutstandingRepoImp(db *sql.DB) OutstandingRepo {
	return &OutstandingRepoImp{
		Db: db,
	}
}

// AddOutstanding function for adding a new outstanding deposit.
func (s *OutstandingRepoImp) AddOutstanding(outstanding *models.OutstandingDeposit) error {
	query := `INSERT INTO outstanding_deposits (id, employeeid, transactionid, amount, ispaid, note, date, createon, updateon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.Db.Exec(query, outstanding.ID, outstanding.EmployeeID, outstanding.TransactionID, outstanding.Amount, outstanding.IsPaid, outstanding.Note, outstanding.Date, outstanding.CreateOn, outstanding.UpdateOn)
	return err
}

// UpdateOutstanding function for updating an existing outstanding deposit.
func (s *OutstandingRepoImp) UpdateOutstanding(outstanding *models.OutstandingDeposit) error {
	query := `UPDATE outstanding_deposits SET employeeid=$1, transactionid=$2, amount=$3, ispaid=$4, note=$5, updateon=$6 WHERE id=$7`

	_, err := s.Db.Exec(query, outstanding.EmployeeID, outstanding.TransactionID, outstanding.Amount, outstanding.IsPaid, outstanding.Note, outstanding.UpdateOn, outstanding.ID)
	return err
}

// GetOutstanding function for retrieving all outstanding deposits.
func (s *OutstandingRepoImp) GetOutstanding() ([]*models.OutstandingDeposit, error) {
	query := `SELECT id, employeeid, transactionid, amount, ispaid, note, date, createon, updateon FROM outstanding_deposits`
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outstanding []*models.OutstandingDeposit
	for rows.Next() {
		dep := &models.OutstandingDeposit{}
		err := rows.Scan(&dep.ID, &dep.EmployeeID, &dep.TransactionID, &dep.Amount, &dep.IsPaid, &dep.Note, &dep.Date, &dep.CreateOn, &dep.UpdateOn)
		if err != nil {
			return nil, err
		}
		outstanding = append(outstanding, dep)
	}

	return outstanding, nil
}

// GetOutstandingByEmployee function for retrieving outstanding deposits by employee ID.
func (s *OutstandingRepoImp) GetOutstandingByEmployee(employeeID string) ([]*models.OutstandingDeposit, error) {
	query := `SELECT id, employeeid, transactionid, amount, ispaid, note, date, createon, updateon FROM outstanding_deposits WHERE employeeid = $1`
	rows, err := s.Db.Query(query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var outstanding []*models.OutstandingDeposit
	for rows.Next() {
		dep := &models.OutstandingDeposit{}
		err := rows.Scan(&dep.ID, &dep.EmployeeID, &dep.TransactionID, &dep.Amount, &dep.IsPaid, &dep.Note, &dep.Date, &dep.CreateOn, &dep.UpdateOn)
		if err != nil {
			return nil, err
		}
		outstanding = append(outstanding, dep)
	}

	return outstanding, nil
}
