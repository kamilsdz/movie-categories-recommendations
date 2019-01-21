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
	UsersNumber           = 10000
	MaxPredictCategories  = 10
)
