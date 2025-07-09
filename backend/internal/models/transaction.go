package models

import "time"

type ProductItem struct {
	ProductName string `json:"productname"`
	Total       int    `json:"total"`
}

type Transaction struct {
	ID           string        `json:"id"`
	EmployeeName string        `json:"employeename"`
	Deposit      int           `json:"deposit"`
	ProductItems []ProductItem `json:"productitems"`
	Date         time.Time     `json:"date"`
	CreateOn     time.Time     `json:"createon"`
	UpdateOn     time.Time     `json:"updateon"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
