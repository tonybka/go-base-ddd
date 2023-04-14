package account

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tonybka/go-base-ddd/domain/entity"
	"github.com/tonybka/go-base-ddd/domain/event"
	"github.com/tonybka/go-base-ddd/persistence/database"
	"github.com/tonybka/go-base-ddd/persistence/trigger"
	"github.com/tonybka/go-base-ddd/tests"
)

const createAccountTableSmt = `
CREATE TABLE accounts
(
		id	serial UNIQUE,
		account_name	varchar(80) UNIQUE NOT NULL,
		created_at timestamp	NOT NULL DEFAULT now(),
		PRIMARY KEY (id)
);
`

const createEmailTableSmt = `
CREATE TABLE emails
(
		id	serial UNIQUE,
		mail	varchar(80) UNIQUE NOT NULL,
		created_at timestamp	NOT NULL DEFAULT now(),
		PRIMARY KEY (id)
);
`

type AccountRepositoryTestSuite struct {
	suite.Suite
	accountRepo *AccountRepository
	emailRepo   *EmailRepository
}

func (ts *AccountRepositoryTestSuite) SetupSuite() {

	// Init global domain publisher
	event.InitDomainEventPublisher()
	publisher := event.GetDomainEventPublisher()
	require.NotNil(ts.T(), publisher)

	dbClient, err := database.NewDBConnection(
		context.Background(),
		&database.DatabaseConfig{
			Host:     "localhost",
			Port:     "54320",
			Name:     "postgres",
			UserName: "postgres",
			Password: "postgres",
		},
	)
	require.NoError(ts.T(), err)
	require.NotNil(ts.T(), dbClient)

	ts.accountRepo = NewAccountRepository(dbClient)
	ts.emailRepo = NewEmailRepository(dbClient)

	_, err = dbClient.Query(context.Background(), createAccountTableSmt)
	require.NoError(ts.T(), err)
	_, err = dbClient.Query(context.Background(), createEmailTableSmt)
	require.NoError(ts.T(), err)

	err = trigger.RegisterModelHooks(dbClient, []string{DBTblNameAccounts, TableNameEmails})
	require.NoError(ts.T(), err)

	// Register handlers of domain event
	accountCreatedSubscribers := []event.IDomainEvenHandler{NewAccountCreatedEventHandler(ts.accountRepo)}
	publisher.RegisterSubscriber(&AccountCreatedEvent{}, accountCreatedSubscribers...)

	// Reset random seed to make sure the generated value is unique
	rand.Seed(time.Now().UnixNano())
}

func (ts *AccountRepositoryTestSuite) TestCreateAccount() {
	randId := rand.Intn(99999)

	account := Account{
		BaseAggregateRoot: entity.NewBaseAggregateRoot(uint(randId)),
		AccountName:       tests.RandomString(),
	}

	result, err := ts.accountRepo.Create(account)
	ts.NoError(err)
	ts.NotNil(result)

	all, err := ts.accountRepo.GetAll()
	ts.NoError(err)
	ts.Greater(len(all), 0)

	queriedAccount, err := ts.accountRepo.FindById(account.ID)
	ts.NoError(err)
	ts.Equal(account.AccountName, queriedAccount.AccountName)
	ts.Equal(account.ID, queriedAccount.ID)
}

func (ts *AccountRepositoryTestSuite) TestCreateEmail() {
	randId := rand.Intn(99999)

	email := SampleEmail{
		BaseEntity: entity.NewBaseEntity(uint(randId)),
		Mail:       tests.RandomString(),
	}

	result, err := ts.emailRepo.Create(email)
	ts.NoError(err)
	ts.NotNil(result)
}

func (ts *AccountRepositoryTestSuite) TestAccountWithEvent() {
	randId := rand.Intn(99999)

	account := Account{
		BaseAggregateRoot: entity.NewBaseAggregateRoot(uint(randId)),
		AccountName:       tests.RandomString(),
	}

	account.AddEvent(DBTblNameAccounts, NewAccountCreatedEvent(uint(randId), nil))

	_, err := ts.accountRepo.Create(account)
	ts.NoError(err)
}

func (ts *AccountRepositoryTestSuite) TearDownSuite() {
}

func TestSuiteRunnerAccountRepository(t *testing.T) {
	ts := new(AccountRepositoryTestSuite)
	suite.Run(t, ts)
}
