# Movie recomendations
Sample KNN algorithm implementation in Go.
The goal of this script is to recommend categories with movies user likes the most.

User has `WatchedMovies` counter of watched movies from specific categories ({category_id: watched_times}).
- user has many watched categories
- category has many movies
- movie has one category

## How to run
`go run service.go`

## How to run tests
```
cd tests
go test
```
