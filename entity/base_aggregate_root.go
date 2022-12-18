package entity

import "github.com/tonybka/go-base-ddd/event"

type BaseAggregateRoot struct {
	BaseEntity
}

func (aggregate *BaseAggregateRoot) AddEvent(event event.IBaseDomainEvent) {

}
