package models

import "time"

type OutstandingDeposit struct {
	ID            string    `json:"id"`
	EmployeeID    string    `json:"employeeid"`
	TransactionID string    `json:"transactionid,omitempty"` // optional: link to a transaction, if applicable
	Amount        float64   `json:"amount"`                  // positive for debt, negative for repayment
	IsPaid        bool      `json:"ispaid"`                  // true if paid (optional depending on model)
	Note          string    `json:"note"`
	Date          time.Time `json:"date"`
	CreateOn      time.Time `json:"createon"`
	UpdateOn      time.Time `json:"updateon"`
}
