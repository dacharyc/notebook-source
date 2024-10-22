package tests

import (
	"test-poc/example/manage-indexes"
	"test-poc/example/run-queries"
	"testing"
)

func TestAnnQueryBasic(t *testing.T) {
	// Test creating the index and performing a query that relies on the index
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
