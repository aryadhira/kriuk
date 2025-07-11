package services

import (
	"encoding/json"
	"io"
	"kriuk/internal/models"
	"kriuk/internal/repository"
	"kriuk/utils"
	"net/http"
)

type EmployeeSvc struct {
	db repository.EmployeeRepo
}

func NewEmployeeSvc(db repository.EmployeeRepo) *EmployeeSvc {
	return &EmployeeSvc{
		db: db,
	}
}

func (s *EmployeeSvc) AddEmployee(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
	}

	defer r.Body.Close()

	var employee models.Employee
	err = json.Unmarshal(bodyBytes, &employee)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	err = s.db.AddEmployee(&employee)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusCreated, "Employee Successfully Added", nil)
}

func (s *EmployeeSvc) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := s.db.GetEmployees()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusOK, "", employees)
}

func (s *EmployeeSvc) GetEmployeeByName(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")

	employee, err := s.db.GetEmployeeByName(name)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	if employee == nil {
		utils.WriteJSON(w, http.StatusNotFound, "Employee not found", nil)
	} else {
		utils.WriteJSON(w, http.StatusOK, "", employee)
	}
}
