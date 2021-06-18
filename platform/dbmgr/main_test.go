package dbmgr

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbSource = "./auth_test.db"
)

func TestMain(m *testing.M) {

	os.Remove(dbSource)

	// instance := GetInstance()

	os.Exit(m.Run())

}
