package account

import (
	"github.com/jackc/pgx/v5"
	"github.com/tonybka/go-base-ddd/domain/entity"
)

const DBTblNameAccounts = "accounts"

type Account struct {
	entity.BaseEntity
	AccountName string
}

func (s *Account) ScanRow(row pgx.Row) error {
	return row.Scan(
		&s.ID,
		&s.AccountName,
		&s.CreatedAt,
	)
}
