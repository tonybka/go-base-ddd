package entity

import "github.com/tonybka/go-base-ddd/domain/event"

type BaseAggregateRoot struct {
	BaseEntity

	domainEvents []event.IBaseDomainEvent
}

func NewBaseAggregateRoot() (BaseAggregateRoot, error) {
	base, err := NewBaseEntity()
	if err != nil {
		return BaseAggregateRoot{}, nil
	}

	events := make([]event.IBaseDomainEvent, 0)

	return BaseAggregateRoot{
		BaseEntity:   base,
		domainEvents: events,
	}, nil
}

func (aggregate *BaseAggregateRoot) AddEvent(event event.IBaseDomainEvent) {
	aggregate.domainEvents = append(aggregate.domainEvents, event)
}

func (aggregate *BaseAggregateRoot) PendingEvents() []event.IBaseDomainEvent {
	return aggregate.domainEvents
}
