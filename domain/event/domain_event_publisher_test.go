package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRegisterEventHandler ensure handler registration works
func TestRegisterEventHandler(t *testing.T) {
	eventPublisher := InitDomainEventPublisher()
	assert.NotNil(t, eventPublisher)

	handler := &MockEventHandler{}
	event := &MockDomainEventStruct{}

	eventPublisher.RegisterSubscriber(event, handler)
}

// TestPublishDomainEvent ensure the handler was notified once event triggered
func TestPublishDomainEvent(t *testing.T) {
	eventPublisher := InitDomainEventPublisher()
	assert.NotNil(t, eventPublisher)

	handler := &MockEventHandler{}
	event := &MockDomainEventStruct{}

	eventPublisher.RegisterSubscriber(event, handler)

	// Before notification
	assert.False(t, handler.Notified)

	eventPublisher.Publish(nil, event)

	// After notification
	assert.True(t, handler.Notified)
}
