package run_queries

type ProjectedMovieResult struct {
	Title string  `bson:"title"`
	Plot  string  `bson:"plot"`
	Score float64 `bson:"score"`
}
