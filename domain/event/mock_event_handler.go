package event

type MockEventHandler struct {
	Notified bool
}

func (handler *MockEventHandler) Notify(event IBaseDomainEvent) error {
	handler.Notified = true
	return nil
}
