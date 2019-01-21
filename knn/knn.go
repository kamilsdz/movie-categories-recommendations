package knn

import (
	"../lib"
	"fmt"
	"math"
)

func PredictPreferredCategories(needyUser lib.User, nearestNeighbors map[float64]lib.User, movieCategories []lib.MovieCategory) []lib.MovieCategory {
	predictedCategoriesRating := make(map[int]int)
	var predictedCategories []lib.MovieCategory
	for _, v := range nearestNeighbors {
		userWatchedMovies := v.WatchedMovies
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

func FindMostSimilarUser(user lib.User, nearestNeighbors map[float64]lib.User) lib.User {
	lowestValue := math.MaxFloat64
	var mostSimilarUser lib.User
	for k, v := range nearestNeighbors {
		if k < lowestValue {
			lowestValue = k
			mostSimilarUser = v
		}
	}
	return mostSimilarUser
}

func PredictedPreferredCategoriesPrint(movieCategories []lib.MovieCategory) {
	for _, category := range movieCategories {
		fmt.Printf("Predicted recommended category for the needyUser: %v\n", category.Name)
	}
}

func SimilarUsersPrint(neighbors map[float64]lib.User) {
	for k, v := range neighbors {
		fmt.Printf("Similar user to needyUser: Name: %s, Rating: %f\n", v.Name, k)
	}
}

func FindNearestNeighborsForUser(needyUser lib.User, users []lib.User, movieCategories []lib.MovieCategory) map[float64]lib.User {
	neighborsMaxQuantity := int(math.Sqrt(lib.UsersNumber))
	index := indexOfUser(needyUser, users)
	users = append(users[:index], users[index+1:]...)
	similarObjects := make(map[float64]lib.User)

	for i := 0; i < lib.UsersNumber-1; i++ {
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

func countPathLength(needyUser lib.User, comparedUser lib.User) float64 {
	var dimensionResult float64
	needyWatchedMovies := needyUser.WatchedMovies
	comparedWatchedMovies := comparedUser.WatchedMovies
	for k, v := range comparedWatchedMovies {
		dimensionResult += math.Pow(float64(v-needyWatchedMovies[k]), 2)
	}
	return math.Sqrt(float64(dimensionResult))

}

func filteredVal(data map[int]int) map[int]int {
	filteredVal := make(map[int]int)
	for k, v := range data {
		lowestKey, lowestValue := lowestMapValue(filteredVal)
		if len(filteredVal) == lib.MaxPredictCategories && lowestValue < v {
			delete(filteredVal, lowestKey)
			filteredVal[k] = v
		} else if len(filteredVal) < lib.MaxPredictCategories {
			filteredVal[k] = v
		}
	}
	return filteredVal
}

func highestKeyValue(similarObjects map[float64]lib.User) float64 {
	value := 0.0
	for k, _ := range similarObjects {
		if k > value {
			value = k
		}
	}
	return value
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

func indexOfUser(user lib.User, collection []lib.User) int {
	var mid int
	var low int
	high := len(collection)
	for low <= high {
		mid := (low + high) / 2
		guess := collection[mid]
		if user.Id == guess.Id {
			return mid
		} else if guess.Id > user.Id {
			high = mid - 1
		} else if guess.Id < user.Id {
			low = mid + 1
		} else {
			return -1
		}
	}
	return mid
}
