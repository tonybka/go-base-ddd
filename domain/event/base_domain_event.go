package event

import "time"

type IBaseDomainEvent interface {
	Name() string
	ToJson() (string, error)
	ID() string
	OccurredAt() time.Time
}
