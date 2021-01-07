package lib

type MovieCategory struct {
	Id   int
	Name string
}

type Movie struct {
	Id int
	Name string
	CategoryId int
}

type User struct {
	Id            int
	Name          string
	WatchedCategories map[int]int
}

const (
	MoviesSeedFile = "seeds/movies.txt"
	MovieCategoriesSeedFile = "seeds/categories.txt"
	UsersNumber           = 50000
	MaxPredictCategories  = 5
	MaxNumberOfWatchedCategories = 5
	NearestNeighborsBatch  = 1000
)
