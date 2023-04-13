package account

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tonybka/go-base-ddd/persistence"
)

const (
	sqlCreateEmail     = `INSERT INTO emails (id, mail) VALUES ($1, $2) RETURNING emails.id;`
	sqlSelectAllEmails = `SELECT id, mail created_at FROM emails;`
)

type EmailRepository struct {
	pgDBConn *pgxpool.Pool
}

func NewEmailRepository(pg *pgxpool.Pool) *EmailRepository {
	return &EmailRepository{pgDBConn: pg}
}

// Create creates new email
func (repo *EmailRepository) Create(email SampleEmail) (*persistence.ResultRowId, error) {
	mailId := &persistence.ResultRowId{}

	err := repo.pgDBConn.QueryRow(context.Background(), sqlCreateEmail, email.ID, email.Mail).Scan(&(*mailId).Id)
	if err != nil {
		return nil, err
	}

	return mailId, nil
}
