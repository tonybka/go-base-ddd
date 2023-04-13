package database

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// attemptToConnect retries to establish connection
func attemptToConnect(fn func() error, attemtps int, delay time.Duration) (err error) {
	if attemtps == 0 {
		attemtps = 5
	}

	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--
			continue
		}
		return nil
	}
	return
}

// NewDBConnection creates new connection to database
func NewDBConnection(ctx context.Context, cfg *DatabaseConfig) (*pgxpool.Pool, error) {
	const timeDelta = 2 * time.Second
	var dbpool *pgxpool.Pool
	var err error

	var once sync.Once
	once.Do(func() {
		err = attemptToConnect(
			func() error {
				ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()

				dbpool, err = pgxpool.New(ctx, cfg.ToConnectionURL())
				if err != nil {
					return err
				}

				return nil
			},
			cfg.ConnectAttempts,
			5*time.Second,
		)

		if err != nil {
			log.Fatal("could not connect to database after retries")
		}
	})

	return dbpool, nil
}
