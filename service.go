package main

import (
	"./knn"
	"./lib"
	"./seeds"
	"fmt"
	"time"
)

func main() {
	var users []lib.User
	var movieCategories []lib.MovieCategory
	seeds.AddMovieCategories(&movieCategories)
	seeds.AddUsers(&users, &movieCategories)
	needyUser := users[5]

	benchmarkStart := time.Now()

	nearestNeighbors := knn.FindNearestNeighborsForUser(needyUser, users, movieCategories)
	knn.SimilarUsersPrint(nearestNeighbors)
	predictedPreferredCategories := knn.PredictPreferredCategories(needyUser, nearestNeighbors, movieCategories)
	knn.PredictedPreferredCategoriesPrint(predictedPreferredCategories)

	benchmarkStop := time.Since(benchmarkStart)
	fmt.Printf("Benchmark time: %s\n", benchmarkStop)
}
