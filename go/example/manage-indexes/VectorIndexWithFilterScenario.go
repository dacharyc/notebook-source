package manage_indexes

import "testing"

type VectorIndexWithFilterScenario struct {
	Name          string
	Path          string
	NumDimensions int
	Similarity    string
	Filters       []filterField
	Testing       *testing.T
	Expectation   IndexExpectation
}

type filterField struct {
	Type string `bson:"type"`
	Path string `bson:"path"`
}
