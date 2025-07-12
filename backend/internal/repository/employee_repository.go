package repository

import (
	"database/sql"
	"kriuk/internal/models"
)

type EmployeeRepo interface {
	AddEmployee(employee *models.Employee) error
	GetEmployees() ([]*models.Employee, error)
	GetEmployeeByName(name string) (*models.Employee, error)
}

type EmployeeRepoImp struct {
	Db *sql.DB
}

func NewEmployeeRepoImp(db *sql.DB) EmployeeRepo {
	return &EmployeeRepoImp{
		Db: db,
	}
}

func (e *EmployeeRepoImp) AddEmployee(employee *models.Employee) error {
	query := `INSERT INTO employee (id, name) VALUES ($1, $2)`

	_, err := e.Db.Exec(query, employee.ID, employee.Name)
	return err
}

func (e *EmployeeRepoImp) GetEmployees() ([]*models.Employee, error) {
	query := `SELECT id, name FROM employee`
	rows, err := e.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		emp := &models.Employee{}
		err := rows.Scan(&emp.ID, &emp.Name)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func (e *EmployeeRepoImp) GetEmployeeByName(name string) (*models.Employee, error) {
	query := `SELECT id, name FROM employee WHERE name = $1 LIMIT 1`
	row := e.Db.QueryRow(query, name)

	emp := &models.Employee{}
	err := row.Scan(&emp.ID, &emp.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return emp, nil
}
