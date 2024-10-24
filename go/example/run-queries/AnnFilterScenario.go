package run_queries

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type AnnFilterScenario struct {
	Database          string
	Collection        string
	Embedding         []float64
	VectorSearchStage bson.D
	ProjectStage      bson.D
	Expected          []ProjectedMovieResultWithFilter
	Testing           *testing.T
}
