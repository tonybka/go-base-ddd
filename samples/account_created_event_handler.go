package account

import (
	"fmt"
	"time"

	"github.com/tonybka/go-base-ddd/domain/event"
)

// AccountCreatedEventHandler triggered once new account created
type AccountCreatedEventHandler struct {
	accountRepo *AccountRepository
	isNotified  bool
	isCompleted bool
}

func NewAccountCreatedEventHandler(accountRepo *AccountRepository) *AccountCreatedEventHandler {
	return &AccountCreatedEventHandler{accountRepo: accountRepo}
}

func (handler *AccountCreatedEventHandler) Notify(event event.IBaseDomainEvent) error {
	fmt.Println("AccountCreatedEventHandler.Notify: get notified")
	handler.isNotified = true
	accountCreatedEvent := event.(*AccountCreatedEvent)
	time.Sleep(1 * time.Second)

	account, err := handler.accountRepo.FindById(accountCreatedEvent.AggregateID())
	if err != nil {
		fmt.Println("Cound not query account")
		return err
	}

	fmt.Printf("Account ID: %d\n", account.ID)
	fmt.Printf("Account Name: %s\n", account.AccountName)

	fmt.Println("AccountCreatedEventHandler.Notify: able to complete")
	handler.isCompleted = true
	return nil
}
