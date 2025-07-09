package models

import "time"

type Stocks struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"qty"`
	Unit     string    `json:"unit"`
	Price    float64   `json:"price"`
	CreateOn time.Time `json:"createon"`
	UpdateOn time.Time `json:"updateon"`
}

func (s *Stocks) TableName() string {
	return "stocks"
}
