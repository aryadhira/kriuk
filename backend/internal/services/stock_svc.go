package services

import (
	"encoding/json"
	"io"
	"kriuk/internal/models"
	"kriuk/internal/repository"
	"kriuk/utils"
	"net/http"
)

type StockSvc struct {
	db repository.StockRepo
}

func NewStockSvc(db repository.StockRepo) *StockSvc {
	return &StockSvc{
		db: db,
	}
}

func (s *StockSvc) GetStock(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filterName := params.Get("name")

	datas := []*models.Stocks{}
	if filterName != "" {
		stocks, err := s.db.GetStocksByName(filterName)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		}

		datas = append(datas, stocks)
	} else {
		stocks, err := s.db.GetStocks()
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		}

		datas = append(datas, stocks...)
	}

	utils.WriteJSON(w, http.StatusOK, "", datas)
}

func (s *StockSvc) UpdateStock(w http.ResponseWriter, r *http.Request) {
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

	stocks := new(models.Stocks)
	stocks.ID = utils.InterfaceToString(payload["id"])
	stocks.Name = utils.InterfaceToString(payload["name"])
	stocks.Unit = utils.InterfaceToString(payload["unit"])
	stocks.Quantity = utils.InterfaceToInt(payload["qty"])
	stocks.Price = utils.InterfaceToFloat(payload["qty"])

	if isUpdate {
		err = s.db.UpdateStock(stocks)
	} else {
		err = s.db.AddStock(stocks)
	}

	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteJSON(w, http.StatusOK, "Stock Successfully Saved", nil)
}
