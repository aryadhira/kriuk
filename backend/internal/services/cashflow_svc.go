package services

import (
	"encoding/json"
	"io"
	"kriuk/internal/models"
	"kriuk/internal/repository"
	"kriuk/utils"
	"net/http"
	"time"
)

type CashflowSvc struct {
	db repository.CashFlowRepo
}

func NewCashflowSvc(db repository.CashFlowRepo) *CashflowSvc {
	return &CashflowSvc{
		db: db,
	}
}

func (s *CashflowSvc) GetCashflow(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filterDate := params.Get("date")

	datas := []*models.CashFlow{}
	if filterDate != "" {
		dt, err := time.Parse("20060102", filterDate)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		}

		cashflows, err := s.db.GetCashFlowByDate(dt)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		}

		datas = append(datas, cashflows)
	} else {
		cashflows, err := s.db.GetCashFlow()
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		}

		datas = append(datas, cashflows...)
	}

	utils.WriteJSON(w, http.StatusOK, "", datas)
}

func (s *CashflowSvc) UpdateCashflow(w http.ResponseWriter, r *http.Request) {
	isUpdate := false

	if r.Method == http.MethodPut {
		isUpdate = true
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
	}

	defer r.Body.Close()

	var payload map[string]interface{}

	err = json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	cashflow := new(models.CashFlow)
	cashflow.ID = utils.InterfaceToString(payload["id"])
	dt, err := time.Parse("20060102", utils.InterfaceToString("date"))
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
	}
	cashflow.Date = dt
	cashType := utils.InterfaceToString(payload["type"])
	cashflow.Type = models.CashFlowType(cashType)
	cashflow.Amount = utils.InterfaceToFloat(payload["amount"])
	cashflow.Note = utils.InterfaceToString(payload["note"])

	if isUpdate {
		err = s.db.UpdateCashFlow(cashflow)
	} else {
		err = s.db.AddCashFlow(cashflow)
	}

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusOK, "Cashflow Successfully Saved", nil)
}

func (s *CashflowSvc) DeleteCashFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}

	params := r.URL.Query()

	id := params.Get("id")

	if id == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "Id cannot empty", nil)
	}

	err := s.db.DeleteCashFlow(id)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusOK, "Cashflow Delete Successfully", nil)
}
