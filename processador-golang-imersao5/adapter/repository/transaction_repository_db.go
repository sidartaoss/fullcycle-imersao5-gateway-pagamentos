package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	*sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{
		DB: db,
	}
}

func (t *TransactionRepositoryDb) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err := t.DB.Prepare(`insert into transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		values(?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
