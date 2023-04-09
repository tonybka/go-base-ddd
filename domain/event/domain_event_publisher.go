package event

import (
	"gorm.io/gorm"
)

// =----------------------
// DomainEventPublisher
// =----------------------
type DomainEventPublisher struct {
	eventHandlers map[string][]IDomainEvenHandler
}

// RegisterSubscriber registers new handlers of given event
func (publisher *DomainEventPublisher) RegisterSubscriber(event IBaseDomainEvent, newHandlers ...IDomainEvenHandler) error {
	eventName := event.Name()

	for _, handler := range newHandlers {
		currentHandlers := publisher.eventHandlers[eventName]
		handlers := append(currentHandlers, handler)
		publisher.eventHandlers[eventName] = handlers
	}

	return nil
}

// Publish notifies all registered subscribers about the given events
func (publisher *DomainEventPublisher) Publish(tx *gorm.DB, events ...IBaseDomainEvent) error {
	for _, event := range events {
		eventName := event.Name()

		handlers := publisher.eventHandlers[eventName]
		if handlers == nil {
			return nil
		}

		for _, handler := range handlers {
			err := handler.Notify(event)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// =-------------------------------
// DomainEventPublisher singleton
// =-------------------------------
var singletonEventPublisher *DomainEventPublisher

func GetDomainEventPublisher() *DomainEventPublisher {
	return singletonEventPublisher
}

// InitDomainEventPublisher initialize Domain Event Publisher
func InitDomainEventPublisher() *DomainEventPublisher {
	singletonEventPublisher = &DomainEventPublisher{
		eventHandlers: map[string][]IDomainEvenHandler{},
	}

	return singletonEventPublisher
}
