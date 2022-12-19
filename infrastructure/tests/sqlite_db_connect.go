package tests

import (
	"io/ioutil"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteDBConnect struct {
	dataDir      string
	dbConnection *gorm.DB
}

func NewSqliteDBConnect() (*SqliteDBConnect, error) {
	var tempDir = ""

	// Temp file setup
	tempDir, err := ioutil.TempDir("", "tests-")
	if err != nil {
		return nil, err
	}

	// Database setup
	tempDir = path.Join(tempDir, "test.sqlite3")
	dialector := sqlite.Open(tempDir)

	dbConn, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqliteDB, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	sqliteDB.SetMaxOpenConns(1)
	dbConn.Debug()

	return &SqliteDBConnect{dataDir: tempDir, dbConnection: dbConn}, nil
}

func (conn *SqliteDBConnect) Connection() *gorm.DB {
	return conn.dbConnection
}

func (conn *SqliteDBConnect) CleanUp() error {
	var err error

	if len(conn.dataDir) > 0 {
		err = os.RemoveAll(conn.dataDir)
	}

	return err
}
