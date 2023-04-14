package entity

import "github.com/tonybka/go-base-ddd/domain/event"

type BaseAggregateRoot struct {
	BaseEntity
}

func NewBaseAggregateRoot(rootEntityId uint) BaseAggregateRoot {
	base := NewBaseEntity(rootEntityId)

	return BaseAggregateRoot{
		BaseEntity: base,
	}
}

func (base *BaseAggregateRoot) AddEvent(tableName string, domainEvent event.IBaseDomainEvent) {
	event.EventSource.AddEvent(tableName, domainEvent)
}
