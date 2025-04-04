package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var testDB *sql.DB

const (
	dbUser = "user"
	dbPassword = "pass"
	dbPort = "3306"
	dbHost = "localhost"
	dbName = "sampledb"
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func connectDB() error {
	var err error

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	testDB, err = sql.Open("mysql", dbConn)

	if err != nil {
		return err
	}
	return nil
}

// テストの前処理
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	return nil
}

// 前テスト共通の後処理を書く
func teardown() {
	testDB.Close()
}