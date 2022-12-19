package account

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	customgorm "github.com/tonybka/go-base-ddd/infrastructure/custom_gorm"
	"github.com/tonybka/go-base-ddd/infrastructure/persistence"
	"github.com/tonybka/go-base-ddd/infrastructure/tests"
	"gorm.io/gorm"
)

type AccountRepositoryTestSuite struct {
	suite.Suite

	sqliteConnect *tests.SqliteDBConnect
	dbConn        *gorm.DB

	accountRepo *AccountRepository
}

func (ts *AccountRepositoryTestSuite) SetupSuite() {
	sqliteConn, err := tests.NewSqliteDBConnect()
	require.NoError(ts.T(), err)

	ts.sqliteConnect = sqliteConn
	ts.dbConn = sqliteConn.Connection()

	ts.dbConn.AutoMigrate(&AccountModel{})

	ts.accountRepo = NewAccountRepository(ts.dbConn)
}

func (ts *AccountRepositoryTestSuite) TestCreateAccount() {
	entityId := customgorm.CustomTypeUUIDv1FromString(uuid.New().String())

	account := AccountModel{
		BaseModel: persistence.BaseModel{
			ID: entityId,
		},
		AccountName: "abc",
	}

	err := ts.accountRepo.Create(account)
	ts.NoError(err)

	all, err := ts.accountRepo.GetAll()
	ts.NoError(err)
	ts.Greater(len(all), 0)

	queriedAccount, err := ts.accountRepo.FindById(uuid.UUID(account.ID))
	ts.NoError(err)
	ts.Equal(account.AccountName, queriedAccount.AccountName)
	ts.Equal(account.ID, queriedAccount.ID)
}

func (ts *AccountRepositoryTestSuite) TearDownSuite() {
	err := ts.sqliteConnect.CleanUp()
	ts.NoError(err)
}

func TestSuiteRunnerAccountRepository(t *testing.T) {
	ts := new(AccountRepositoryTestSuite)
	suite.Run(t, ts)
}
