package account

import (
	"github.com/tonybka/go-base-ddd/domain/event"
)

type AccountCreatedEvent struct {
	*event.BaseDomainEvent
}

func NewAccountCreatedEvent(aggregateID uint, data interface{}) *AccountCreatedEvent {
	base := event.NewBaseDomainEvent("accounts", aggregateID, "event.account.created", data)
	return &AccountCreatedEvent{
		BaseDomainEvent: base,
	}
}

func (event *AccountCreatedEvent) Name() string {
	return "event.account.created"
}
