package database

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	err := Database(":memory:")
	if err != nil {
		t.Errorf("Database connection failed: %v", err)
	}
	if DB == nil {
		t.Fatalf("Database connection not established: %v", DB)
	} else {
		DB.Close()
	}
}