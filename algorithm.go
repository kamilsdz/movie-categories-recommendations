package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type MovieCategory struct {
	id   int
	name string
}

type User struct {
	id            int
	name          string
	watchedMovies map[int]int
}

const (
	movieCategoriesNumber = 30
	usersNumber           = 10000
)

func main() {
	var users []User
	var movieCategories []MovieCategory
	addMovieCategories(&movieCategories)
	addUsers(&users, &movieCategories)
	needyUser := users[5]

	benchmarkStart := time.Now()

	nearestNeighbors := FindNearestNeighborsForUser(needyUser.id, users, movieCategories)
	similarUsersPrint(nearestNeighbors)

	benchmarkStop := time.Since(benchmarkStart)
	fmt.Printf("Benchmark time: %s", benchmarkStop)
}

func similarUsersPrint(neighbors map[float64]User) {
	for i, v := range neighbors {
		fmt.Printf("Rating: %f - Username: %s\n", i, v.name)
	}
}

func FindNearestNeighborsForUser(needy_id int, users []User, movieCategories []MovieCategory) map[float64]User {
	neighborsMaxQuantity := int(math.Sqrt(usersNumber))
	index := needy_id - 2
	needyUser := users[index]
	users = append(users[:index], users[index+1:]...)
	similarObjects := make(map[float64]User)

	for i := 0; i < usersNumber-1; i++ {
		pathLength := countPathLength(needyUser, users[i])
		if len(similarObjects) <= neighborsMaxQuantity || float64(highestKeyValue(similarObjects)) > pathLength {
			if len(similarObjects) == neighborsMaxQuantity {
				delete(similarObjects, highestKeyValue(similarObjects))
			}
			similarObjects[pathLength] = users[i]
		}
	}
	return similarObjects
}

func highestKeyValue(similarObjects map[float64]User) float64 {
	value := 0.0
	for k, _ := range similarObjects {
		if k > value {
			value = k
		}
	}
	return value
}

func countPathLength(needyUser User, comparedUser User) float64 {
	var dimensionResult float64
	needyWatchedMovies := needyUser.watchedMovies
	comparedWatchedMovies := comparedUser.watchedMovies
	for k, v := range comparedWatchedMovies {
		dimensionResult += math.Pow(float64(v-needyWatchedMovies[k]), 2)
	}
	return math.Sqrt(float64(dimensionResult))

}

func addMovieCategories(movieCategories *[]MovieCategory) {
	for i := 1; i <= movieCategoriesNumber; i++ {
		movieCategory := MovieCategory{i, fmt.Sprintf("Movie category %s", i)}
		*movieCategories = append(*movieCategories, movieCategory)
	}
}

func addUsers(users *[]User, movieCategories *[]MovieCategory) {
	for i := 1; i <= usersNumber; i++ {
		watchedMoviesMap := generateWatchedMoviesForUser()
		user := User{i, fmt.Sprintf("User %d", i), watchedMoviesMap}
		*users = append(*users, user)
	}
}

func generateWatchedMoviesForUser() map[int]int {
	watchedMoviesMap := make(map[int]int)
	for i := 1; i <= movieCategoriesNumber; i++ {
		maximumNumberOfVideosWatched := 20
		watchedMoviesMap[i] = randomInt(0, maximumNumberOfVideosWatched)
	}
	return watchedMoviesMap
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
