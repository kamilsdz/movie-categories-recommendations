# Movie recomendations
Sample KNN algorithm implementation in Go.
The goal of this script is to recommend categories with movies user likes the most.

User has `WatchedCategories` counter of watched movies from specific categories (`{category_id: watched_x_times}`).
- user has many watched categories
- movie has one category

The assumption is that if a user watches the most videos from certain categories, he will likely be interested in videos from similar categories that are watched by the people who watch them.
Simply put - if user X watches only ACTION movies only and user B watches ACTION and DRAMA movies - the algorithm will recommend ACTION and DRAMA categories for person A. It's how kNN algotithm works.

## Performance
Not too bad. When there is 50k users and 20 categories, it takes ~20 ms to calculate a recommendations on my MBP.

## How to run
`go run main.go`

## How to run tests
`go test -v ./tests`
