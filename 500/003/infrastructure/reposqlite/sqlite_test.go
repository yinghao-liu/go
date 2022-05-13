package reposqlite_test

import (
	"ddd/infrastructure/reposqlite"
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	reposqlite.DatabaseInit()
	reposqlite.DatabaseHandle()
}
