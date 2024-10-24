package manage_indexes

import "testing"

type VectorIndexScenario struct {
	Name   string
	Fields []struct {
		Type          string `bson:"type"`
		Path          string `bson:"path"`
		NumDimensions int    `bson:"numDimensions"`
		Similarity    string `bson:"similarity"`
	}
	Testing *testing.T
}
