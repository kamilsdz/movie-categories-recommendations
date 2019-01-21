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
	maxPredictCategories  = 10
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
	predictedPreferredCategories := PredictPreferredCategories(needyUser, nearestNeighbors, movieCategories)
	predictedPreferredCategoriesPrint(predictedPreferredCategories)

	benchmarkStop := time.Since(benchmarkStart)
	fmt.Printf("Benchmark time: %s", benchmarkStop)
}

func PredictPreferredCategories(needyUser User, nearestNeighbors map[float64]User, movieCategories []MovieCategory) []MovieCategory {
	predictedCategoriesRating := make(map[int]int)
	var predictedCategories []MovieCategory
	for _, v := range nearestNeighbors {
		userWatchedMovies := v.watchedMovies
		for k, v := range userWatchedMovies {
			predictedCategoriesRating[k] += v
		}
	}
	predictedCategoriesRating = filteredVal(predictedCategoriesRating)

	for k, _ := range predictedCategoriesRating {
		predictedCategories = append(predictedCategories, movieCategories[k-1])
	}
	return predictedCategories
}

func predictedPreferredCategoriesPrint(movieCategories []MovieCategory) {
	for _, category := range movieCategories {
		fmt.Printf("Predicted category: %v\n", category.name)
	}
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
		movieCategory := MovieCategory{i, fmt.Sprintf("Movie category %d", i)}
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

func filteredVal(data map[int]int) map[int]int {
	filteredVal := make(map[int]int)
	for k, v := range data {
		lowestKey, lowestValue := lowestMapValue(filteredVal)
		if len(filteredVal) == maxPredictCategories && lowestValue < v {
			delete(filteredVal, lowestKey)
			filteredVal[k] = v
		} else if len(filteredVal) < maxPredictCategories {
			filteredVal[k] = v
		}
	}
	return filteredVal
}

func lowestMapValue(data map[int]int) (int, int) {
	minVal := int(^uint(0) >> 1)
	var key int
	var value int

	for k, v := range data {
		if v < minVal {
			minVal = v
			key = k
			value = v
		}
	}
	return key, value
}
