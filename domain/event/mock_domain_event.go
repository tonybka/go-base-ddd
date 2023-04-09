package event

import (
	"encoding/json"
	"time"
)

// MockDomainEventStruct is the simulation of Domain Event structure in domain layer that would
// be referred in persistence layer for publishing
type MockDomainEventStruct struct {
	EventName string
	EventID   string
}

func (event *MockDomainEventStruct) Name() string {
	return event.EventName
}

func (event *MockDomainEventStruct) ToJson() (string, error) {
	jsonString, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func (event *MockDomainEventStruct) ID() string {
	return event.EventID
}

func (event *MockDomainEventStruct) OccurredAt() time.Time {
	return time.Now()
}
