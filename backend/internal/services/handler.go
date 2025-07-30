package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Kriuk struct {
	stock       *StockSvc
	employee    *EmployeeSvc
	cashflow    *CashflowSvc
	transaction *TransactionSvc
}

func NewKriuk(stock *StockSvc, employee *EmployeeSvc, cashflow *CashflowSvc, transaction *TransactionSvc) *Kriuk {
	return &Kriuk{
		stock:       stock,
		employee:    employee,
		cashflow:    cashflow,
		transaction: transaction,
	}
}

func (h *Kriuk) registerHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/stocks", h.stock.GetStock).Methods(http.MethodGet)
	router.HandleFunc("/stocks", h.stock.UpdateStock).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/employees", h.employee.AddEmployee).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/employees", h.employee.GetEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employeesbyname", h.employee.GetEmployeeByName).Methods(http.MethodGet)
	router.HandleFunc("/cashflows", h.cashflow.GetCashflow).Methods(http.MethodGet)
	router.HandleFunc("/cashflows", h.cashflow.UpdateCashflow).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/cashflows", h.cashflow.DeleteCashFlow).Methods(http.MethodDelete)
	router.HandleFunc("/transactions", h.transaction.GetTransactions).Methods(http.MethodGet)
	router.HandleFunc("/transactions", h.transaction.CreateTransaction).Methods(http.MethodPost, http.MethodPut)

	return router
}

func (h *Kriuk) Start() error {
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")
	listenAddr := fmt.Sprintf("%s:%s",apiHost,apiPort)

	router := h.registerHandler()

	server := new(http.Server)
	server.Handler = router
	server.Addr = listenAddr

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, _ := route.GetMethods()
		fmt.Printf("Path: %s, Methods: %v\n", pathTemplate, methods)
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Kriuk services running on: ", listenAddr)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}