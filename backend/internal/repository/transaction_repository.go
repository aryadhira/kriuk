package repository

import (
	"database/sql"
	"kriuk/internal/models"
	"time"
)

type TransactionRepo interface {
	AddTransaction(transaction *models.Transaction) error
	GetTransactions() ([]*models.Transaction, error)
	GetTransactionByDate(date time.Time) (*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id string) error
}

type TransactionRepoImp struct {
	Db *sql.DB
}

func NewTransactionRepoImp(db *sql.DB) TransactionRepo {
	return &TransactionRepoImp{
		Db: db,
	}
}

func (t *TransactionRepoImp) AddTransaction(transaction *models.Transaction) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	now := time.Now()
	transaction.CreateOn = now
	transaction.UpdateOn = now

	// Insert main transaction
	query := `INSERT INTO transaction (id, employeename, deposit, date, createon, updateon)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.Exec(query, transaction.ID, transaction.EmployeeName, transaction.Deposit, transaction.Date, transaction.CreateOn, transaction.UpdateOn)
	if err != nil {
		return err
	}

	// Insert each product item
	itemQuery := `INSERT INTO transaction_items (transaction_id, productname, total)
	              VALUES ($1, $2, $3)`
	for _, item := range transaction.ProductItems {
		_, err := tx.Exec(itemQuery, transaction.ID, item.ProductName, item.Total)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (t *TransactionRepoImp) GetTransactions() ([]*models.Transaction, error) {
	query := `SELECT id, employeename, deposit, date, createon, updateon FROM transaction`
	rows, err := t.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*models.Transaction
	for rows.Next() {
		trx := &models.Transaction{}
		err := rows.Scan(&trx.ID, &trx.EmployeeName, &trx.Deposit, &trx.Date, &trx.CreateOn, &trx.UpdateOn)
		if err != nil {
			return nil, err
		}

		items, err := t.getTransactionItems(trx.ID)
		if err != nil {
			return nil, err
		}
		trx.ProductItems = items
		results = append(results, trx)
	}

	return results, nil
}

func (t *TransactionRepoImp) GetTransactionByDate(date time.Time) (*models.Transaction, error) {
	query := `SELECT id, employeename, deposit, date, createon, updateon
	          FROM transaction WHERE DATE(date) = DATE($1) LIMIT 1`

	trx := &models.Transaction{}
	err := t.Db.QueryRow(query, date).Scan(&trx.ID, &trx.EmployeeName, &trx.Deposit, &trx.Date, &trx.CreateOn, &trx.UpdateOn)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	trx.ProductItems, err = t.getTransactionItems(trx.ID)
	if err != nil {
		return nil, err
	}

	return trx, nil
}

func (t *TransactionRepoImp) UpdateTransaction(transaction *models.Transaction) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	transaction.UpdateOn = time.Now()

	query := `UPDATE transaction
	          SET employeename = $1, deposit = $2, date = $3, updateon = $4
	          WHERE id = $5`

	_, err = tx.Exec(query, transaction.EmployeeName, transaction.Deposit, transaction.Date, transaction.UpdateOn, transaction.ID)
	if err != nil {
		return err
	}

	// Delete old product items
	_, err = tx.Exec(`DELETE FROM transaction_items WHERE transaction_id = $1`, transaction.ID)
	if err != nil {
		return err
	}

	// Re-insert updated product items
	itemQuery := `INSERT INTO transaction_items (transaction_id, productname, total)
	              VALUES ($1, $2, $3)`
	for _, item := range transaction.ProductItems {
		_, err := tx.Exec(itemQuery, transaction.ID, item.ProductName, item.Total)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (t *TransactionRepoImp) DeleteTransaction(id string) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM transaction_items WHERE transaction_id = $1`, id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM transaction WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t *TransactionRepoImp) getTransactionItems(transactionID string) ([]models.ProductItem, error) {
	query := `SELECT productname, total FROM transaction_items WHERE transaction_id = $1`
	rows, err := t.Db.Query(query, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.ProductItem
	for rows.Next() {
		var item models.ProductItem
		err := rows.Scan(&item.ProductName, &item.Total)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
