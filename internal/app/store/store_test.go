package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost username=nikolasmelui dbname=go_rest_api password=password sslmode=disable"
	}

	os.Exit(m.Run())
}
