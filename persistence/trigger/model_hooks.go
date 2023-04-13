package trigger

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tonybka/go-base-ddd/domain/event"
)

// modelHooks get called whenever there is new create-or-update event triggered
func modelHooks(pid uint, table string, rowID uint, action string) error {
	pendingEvents := event.EventSource.GetPendingEvents(table)

	publisher := event.GetDomainEventPublisher()
	err := publisher.Publish(pendingEvents...)
	if err != nil {
		return err
	}

	return nil
}

// RegisterModelHooks register listener of triggers
func RegisterModelHooks(dbPool *pgxpool.Pool, tables []string) error {

	// Create trigger function
	_, err := dbPool.Query(context.Background(), SQLCreateInsertOrUpdateTrigger)
	if err != nil {
		return err
	}

	// Create triggers for database tables
	for _, table := range tables {
		sql := fmt.Sprintf(sqlCreateInsertUpdateTrigger, table, table)
		_, err := dbPool.Query(context.Background(), sql)

		if err != nil {
			return err
		}
	}

	// Called at last, in goroutine
	go listenToTriggers(dbPool, modelHooks)

	return nil
}
