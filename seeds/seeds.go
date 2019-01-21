package seeds

import (
	"../lib"
	"fmt"
	"math/rand"
	"time"
)

func AddMovieCategories(movieCategories *[]lib.MovieCategory) {
	for i := 1; i <= lib.MovieCategoriesNumber; i++ {
		movieCategory := lib.MovieCategory{i, fmt.Sprintf("Movie category %d", i)}
		*movieCategories = append(*movieCategories, movieCategory)
	}
}

func AddUsers(users *[]lib.User, movieCategories *[]lib.MovieCategory) {
	for i := 1; i <= lib.UsersNumber; i++ {
		watchedMoviesMap := generateWatchedMoviesForUser()
		user := lib.User{i, fmt.Sprintf("User %d", i), watchedMoviesMap}
		*users = append(*users, user)
	}
}

func generateWatchedMoviesForUser() map[int]int {
	watchedMoviesMap := make(map[int]int)
	for i := 1; i <= lib.MovieCategoriesNumber; i++ {
		maximumNumberOfVideosWatched := 20
		watchedMoviesMap[i] = randomInt(0, maximumNumberOfVideosWatched)
	}
	return watchedMoviesMap
}
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
