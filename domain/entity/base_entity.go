package entity

import (
	"time"
)

type BaseEntity struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntity(id uint) BaseEntity {
	return BaseEntity{ID: id, CreatedAt: time.Now()}
}
