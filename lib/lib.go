package lib

type MovieCategory struct {
	Id   int
	Name string
}

type User struct {
	Id            int
	Name          string
	WatchedMovies map[int]int
}

const (
	MovieCategoriesNumber = 30
	UsersNumber           = 50000
	MaxPredictCategories  = 10
	NearestNeighborsPart  = 1000
)
