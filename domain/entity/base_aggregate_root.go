package entity

import "github.com/tonybka/go-base-ddd/domain/event"

type BaseAggregateRoot struct {
	BaseEntity
}

func (aggregate *BaseAggregateRoot) AddEvent(event event.IBaseDomainEvent) {

}
