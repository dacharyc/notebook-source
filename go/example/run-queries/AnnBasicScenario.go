package run_queries

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type AnnBasicScenario struct {
	Database          string
	Collection        string
	Embedding         []float64
	VectorSearchStage bson.D
	ProjectStage      bson.D
	Expected          []ProjectedMovieResult
	Testing           *testing.T
}
