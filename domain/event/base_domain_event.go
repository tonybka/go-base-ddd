package event

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type IBaseDomainEvent interface {
	ID() uuid.UUID
	AggregateID() uint
	Aggregate() string
	Name() string
	ToJson() (string, error)
	OccurredAt() time.Time
	Data() interface{}
}

type domainEventProps struct {
	ID          uuid.UUID
	AggregateID uint
	Aggregate   string
	Name        string
	OccuredAt   time.Time
	Data        interface{}
}

type BaseDomainEvent struct {
	props domainEventProps
}

func NewBaseDomainEvent(aggregate string, aggregateID uint, name string, data interface{}) *BaseDomainEvent {
	return &BaseDomainEvent{
		props: domainEventProps{
			ID:          uuid.New(),
			Aggregate:   aggregate,
			AggregateID: aggregateID,
			Name:        name,
			OccuredAt:   time.Now(),
			Data:        data,
		},
	}
}

func (event *BaseDomainEvent) ID() uuid.UUID {
	return event.props.ID
}

func (event *BaseDomainEvent) AggregateID() uint {
	return event.props.AggregateID
}

func (event *BaseDomainEvent) Aggregate() string {
	return event.props.Aggregate
}

func (event *BaseDomainEvent) ToJson() (string, error) {
	b, err := json.Marshal(event.props)
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (event *BaseDomainEvent) OccurredAt() time.Time {
	return event.props.OccuredAt
}

func (event *BaseDomainEvent) Data() interface{} {
	return event.props.Data
}
