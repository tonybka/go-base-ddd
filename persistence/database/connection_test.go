package database

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     "localhost",
		Port:     "54320",
		Name:     "postgres",
		UserName: "postgres",
		Password: "postgres",
	}
}

// Make sure you have run the postgresQL database from docker compose
func TestConnectDB(t *testing.T) {
	dbConfig := testConfig()
	assert.NotNil(t, dbConfig)

	conn, err := NewDBConnection(context.Background(), dbConfig)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
}

func TestBuildConnectionDSN(t *testing.T) {
	dbConfig := testConfig()
	dsn := dbConfig.ToConnectionURL()

	assert.Greater(t, len(dsn), 0)
	assert.True(t, strings.HasPrefix(dsn, "postgres://"))
}
