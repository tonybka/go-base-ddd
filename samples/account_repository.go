package account

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tonybka/go-base-ddd/persistence"
)

const (
	sqlCreateAccount     = `INSERT INTO accounts (id, account_name) VALUES ($1, $2) RETURNING accounts.id;`
	sqlQueryAccountById  = `SELECT id, account_name, created_at FROM accounts WHERE id = $1;`
	sqlSelectAllAccounts = `SELECT id, account_name, created_at FROM accounts;`
)

type AccountRepository struct {
	pgDBConn *pgxpool.Pool
}

func NewAccountRepository(pg *pgxpool.Pool) *AccountRepository {
	return &AccountRepository{pgDBConn: pg}
}

// Create creates new account
func (repo *AccountRepository) Create(account Account) (*persistence.ResultRowId, error) {
	accountId := &persistence.ResultRowId{}

	err := repo.pgDBConn.QueryRow(context.Background(), sqlCreateAccount, account.ID, account.AccountName).Scan(&(*accountId).Id)
	if err != nil {
		return nil, err
	}

	return accountId, nil
}

// FindById query account by it's identity
func (repo *AccountRepository) FindById(id uint) (*Account, error) {
	account := &Account{}

	row := repo.pgDBConn.QueryRow(context.Background(), sqlQueryAccountById, id)
	err := account.ScanRow(row)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetAll returns all accounts in the table
func (repo *AccountRepository) GetAll() ([]*Account, error) {
	var accounts []*Account

	rows, err := repo.pgDBConn.Query(context.Background(), sqlSelectAllAccounts)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		account := &Account{}
		err = account.ScanRow(rows)
		accounts = append(accounts, account)
	}

	return accounts, nil
}
