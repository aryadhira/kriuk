package services

import (
	"encoding/json"
	"io"
	"kriuk/internal/models"
	"kriuk/internal/repository"
	"kriuk/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TransactionSvc struct {
	db  repository.TransactionRepo
	emp repository.EmployeeRepo
	stk repository.StockRepo
	out repository.OutstandingRepo
}

func NewTransactionSvc(
	db repository.TransactionRepo,
	emp repository.EmployeeRepo,
	stk repository.StockRepo,
	out repository.OutstandingRepo,
) *TransactionSvc {
	return &TransactionSvc{
		db:  db,
		emp: emp,
		stk: stk,
		out: out,
	}
}

func (t *TransactionSvc) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
	}

	defer r.Body.Close()

	var transaction *models.Transaction
	err = json.Unmarshal(bodyBytes, &transaction)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
	}

	now := time.Now()
	trxID := uuid.NewString()

	stocks, err := t.stk.GetStocks()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}

	stockMap := make(map[string]*models.Stocks)
	for _, each := range stocks {
		stockMap[each.Name] = each
	}

	transaction.ID = trxID
	transaction.CreateOn = now
	transaction.UpdateOn = now

	// update stock
	totalamount := 0.0
	for _, items := range transaction.ProductItems {
		product := stockMap[items.ProductName]
		product.Quantity = product.Quantity - items.Total
		product.UpdateOn = now
		err = t.stk.UpdateStock(product)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "Stock update failed", nil)
		}
		totalamount += float64(items.Total) * product.Price
	}

	// Create outstanding deposit
	outstanding := new(models.OutstandingDeposit)
	outstanding.ID = uuid.NewString()
	outstanding.TransactionID = trxID
	outstanding.Amount = totalamount
	outstanding.IsPaid = false
	outstanding.Note = ""
	outstanding.Date = now
	outstanding.CreateOn = now
	outstanding.UpdateOn = now

	err = t.out.AddOutstanding(outstanding)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Outstanding save failed", nil)
	}

	// save transaction
	err = t.db.AddTransaction(transaction)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "Error save transaction", nil)
	}

}

func (s *TransactionSvc) GetTransactions(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	var results []*models.Transaction

	if date != "" {
		dateVal, err := time.Parse("20060102", date)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		}
		trx, err := s.db.GetTransactionByDate(dateVal)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		results = append(results, trx)
	} else {
		transactions, err := s.db.GetTransactions()
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		results = transactions
	}

	utils.WriteJSON(w, http.StatusOK, "", results)
}
