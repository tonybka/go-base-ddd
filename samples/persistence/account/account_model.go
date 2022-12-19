package account

import (
	"github.com/tonybka/go-base-ddd/infrastructure/persistence"
)

type AccountModel struct {
	persistence.BaseModel

	AccountName string `gorm:"column:account_name;unique"`
}
