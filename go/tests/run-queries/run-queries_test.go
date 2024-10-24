package tests

import (
	"os"
	"test-poc/example/manage-indexes"
	"test-poc/example/run-queries"
	"testing"
)

func TestAnnQueryBasic(t *testing.T) {
	// Test creating the index and performing a query that relies on the index
	if os.Getenv("ENV_CI") != "" {
		t.Skip("Skipping this test in CI because it doesn't work on local Atlas")
	}
	manage_indexes.ExampleCreateIndexBasic(t)
	run_queries.ExampleAnnBasicQuery(t)

	// Drop the index to clear state for future tests
	manage_indexes.ExampleDropIndex()
}

func TestAnnQueryWithFilter(t *testing.T) {
	// Test creating the index and performing a query that relies on the index
	manage_indexes.ExampleCreateIndexFilter(t)
	run_queries.ExampleAnnFilterQuery(t)

	// Drop the index to clear state for future tests
	manage_indexes.ExampleDropIndex()
}
