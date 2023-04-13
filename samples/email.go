package account

import (
	"github.com/tonybka/go-base-ddd/domain/entity"
)

const TableNameEmails = "emails"

type SampleEmail struct {
	entity.BaseEntity
	Mail string
}
