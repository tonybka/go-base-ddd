package event

var EventSource *EventStorage

func init() {
	if EventSource == nil {
		EventSource = &EventStorage{
			pendingEvents: make(map[string][]IBaseDomainEvent),
		}
	}
}

type EventStorage struct {
	pendingEvents map[string][]IBaseDomainEvent
}

func (storage *EventStorage) AddEvent(dataModel string, domainEvent IBaseDomainEvent) {
	if storage.pendingEvents[dataModel] == nil {
		storage.pendingEvents[dataModel] = make([]IBaseDomainEvent, 0)
	}

	storage.pendingEvents[dataModel] = append(storage.pendingEvents[dataModel], domainEvent)
}

func (storage *EventStorage) GetPendingEvents(dataModel string) []IBaseDomainEvent {
	return storage.pendingEvents[dataModel]
}
