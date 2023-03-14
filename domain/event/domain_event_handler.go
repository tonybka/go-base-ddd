package event

type IDomainEvenHandler interface {
	Notify(event IBaseDomainEvent) error
}
