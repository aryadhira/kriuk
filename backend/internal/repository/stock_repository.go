package repository

import (
	"database/sql"
	"fmt"
	"kriuk/internal/models"
	"time"
)

type StockRepo interface {
	AddStock(stock *models.Stocks) error
	GetStocks() ([]*models.Stocks, error)
	GetStocksByName(name string) (*models.Stocks, error)
	UpdateStock(stock *models.Stocks) error
}

type StockRepoImp struct {
	Db *sql.DB
}

func (s *StockRepoImp) NewStockRepoImp(db *sql.DB) StockRepo {
	return &StockRepoImp{
		Db: db,
	}
}

func (s *StockRepoImp) AddStock(stock *models.Stocks) error {
	query := `INSERT INTO stocks (id, name, quantity, unit, price, createon, updateon)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	now := time.Now()
	stock.CreateOn = now
	stock.UpdateOn = now

	_, err := s.Db.Exec(query,
		stock.ID,
		stock.Name,
		stock.Quantity,
		stock.Unit,
		stock.Price,
		stock.CreateOn,
		stock.UpdateOn,
	)

	return err
}
func (s *StockRepoImp) GetStocks() ([]*models.Stocks, error) {
	query := `SELECT id, name, quantity, unit, price, createon, updateon FROM stocks`
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []*models.Stocks
	for rows.Next() {
		stock := &models.Stocks{}
		err := rows.Scan(
			&stock.ID,
			&stock.Name,
			&stock.Quantity,
			&stock.Unit,
			&stock.Price,
			&stock.CreateOn,
			&stock.UpdateOn,
		)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}
func (s *StockRepoImp) GetStocksByName(name string) (*models.Stocks, error) {
	query := `SELECT id, name, quantity, unit, price, createon, updateon FROM stocks WHERE name = $1 LIMIT 1`
	row := s.Db.QueryRow(query, name)

	stock := &models.Stocks{}
	err := row.Scan(
		&stock.ID,
		&stock.Name,
		&stock.Quantity,
		&stock.Unit,
		&stock.Price,
		&stock.CreateOn,
		&stock.UpdateOn,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return stock, nil
}
func (s *StockRepoImp) UpdateStock(stock *models.Stocks) error {
	query := `UPDATE stocks SET name = $1, quantity = $2, unit = $3, price = $4, updateon = $5 WHERE id = $6`

	stock.UpdateOn = time.Now()

	res, err := s.Db.Exec(query,
		stock.Name,
		stock.Quantity,
		stock.Unit,
		stock.Price,
		stock.UpdateOn,
		stock.ID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no stock found with ID: %s", stock.ID)
	}

	return nil
}
