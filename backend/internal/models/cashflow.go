package models

import "time"

type CashFlowType string

const (
	CashIn  CashFlowType = "in"
	CashOut CashFlowType = "out"
)

type CashFlow struct {
	ID       string       `json:"id"`
	Type     CashFlowType `json:"type"`
	Amount   float64      `json:"amount"`
	Note     string       `json:"note"`
	Date     time.Time    `json:"date"`
	CreateOn time.Time    `json:"createon"`
	UpdateOn time.Time    `json:"updateon"`
}

func (c *CashFlow) TableName() string {
	return "cashflow"
}
