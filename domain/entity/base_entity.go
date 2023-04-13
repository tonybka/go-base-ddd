package entity

import (
	"time"

	"github.com/tonybka/go-base-ddd/domain/event"
)

type BaseEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntity(id uint) BaseEntity {
	return BaseEntity{ID: id, CreatedAt: time.Now()}
}

func (base *BaseEntity) AddEvent(tableName string, domainEvent event.IBaseDomainEvent) {
	event.EventSource.AddEvent(tableName, domainEvent)
}
