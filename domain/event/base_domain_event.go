package event

import "time"

type IBaseDomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}
