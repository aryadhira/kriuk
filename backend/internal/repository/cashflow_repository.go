package repository

import (
	"database/sql"
	"fmt"
	"kriuk/internal/models"
	"time"
)

type CashFlowRepo interface {
	AddCashFlow(cashflow *models.CashFlow) error
	GetCashFlow() ([]*models.CashFlow, error)
	GetCashFlowByDate(date time.Time) (*models.CashFlow, error)
	UpdateCashFlow(cash *models.CashFlow) error
	DeleteCashFlow(cashid string) error
}

type CashFlowRepoImp struct {
	Db *sql.DB
}

func (s *CashFlowRepoImp) NewCashFlowRepoImp(db *sql.DB) CashFlowRepo {
	return &CashFlowRepoImp{
		Db: db,
	}
}

func (c *CashFlowRepoImp) AddCashFlow(cashflow *models.CashFlow) error {
	query := `INSERT INTO cashflow (id, type, amount, note, date, createon, updateon)
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	now := time.Now()
	cashflow.CreateOn = now
	cashflow.UpdateOn = now

	_, err := c.Db.Exec(query,
		cashflow.ID,
		cashflow.Type,
		cashflow.Amount,
		cashflow.Note,
		cashflow.Date,
		cashflow.CreateOn,
		cashflow.UpdateOn,
	)

	return err
}

func (c *CashFlowRepoImp) GetCashFlow() ([]*models.CashFlow, error) {
	query := `SELECT id, type, amount, note, date, createon, updateon FROM cashflow`
	rows, err := c.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.CashFlow
	for rows.Next() {
		c := &models.CashFlow{}
		err := rows.Scan(
			&c.ID,
			&c.Type,
			&c.Amount,
			&c.Note,
			&c.Date,
			&c.CreateOn,
			&c.UpdateOn,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, c)
	}

	return results, nil
}

func (c *CashFlowRepoImp) GetCashFlowByDate(date time.Time) (*models.CashFlow, error) {
	query := `SELECT id, type, amount, note, date, createon, updateon 
	          FROM cashflow 
	          WHERE DATE(date) = DATE($1) 
	          LIMIT 1`

	row := c.Db.QueryRow(query, date)

	cash := &models.CashFlow{}
	err := row.Scan(
		&cash.ID,
		&cash.Type,
		&cash.Amount,
		&cash.Note,
		&cash.Date,
		&cash.CreateOn,
		&cash.UpdateOn,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return cash, nil
}

func (c *CashFlowRepoImp) UpdateCashFlow(cash *models.CashFlow) error {
	query := `UPDATE cashflow 
	          SET type = $1, amount = $2, note = $3, date = $4, updateon = $5 
	          WHERE id = $6`

	cash.UpdateOn = time.Now()

	res, err := c.Db.Exec(query,
		cash.Type,
		cash.Amount,
		cash.Note,
		cash.Date,
		cash.UpdateOn,
		cash.ID,
	)

	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no cashflow found with ID: %s", cash.ID)
	}

	return nil
}

func (c *CashFlowRepoImp) DeleteCashFlow(cashid string) error {
	query := `DELETE FROM cashflow WHERE id = $1`
	res, err := c.Db.Exec(query, cashid)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("no cashflow found with ID: %s", cashid)
	}

	return nil
}
