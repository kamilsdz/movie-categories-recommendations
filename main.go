package main

import (
	"./knn"
	"./lib"
	"./printer"
	"./seeds"
	"./user"
	"fmt"
	"time"
)

func main() {
	var movieCategories []lib.MovieCategory
	var users []lib.User

	addSeeds(&movieCategories, &users)

	currentUser := user.Build(movieCategories, &users)
	users = append(users, currentUser)

	benchmarkStart := time.Now()

	nearestNeighborsModel := knn.NearestObjects(currentUser, users, movieCategories)
	predictedPreferredCategories := knn.PredictPreferredCategories(currentUser, nearestNeighborsModel, movieCategories)
	mostSimilarUser := knn.FindMostSimilarUser(currentUser, nearestNeighborsModel)

	benchmarkStop := time.Since(benchmarkStart)

	printer.SimilarUsers(nearestNeighborsModel)
	printer.MostSimilarUser(mostSimilarUser)
	printer.PredictedPreferredCategories(predictedPreferredCategories)
	printer.BenchmarkResult(benchmarkStop)
}

func addSeeds(movieCategories *[]lib.MovieCategory, users *[]lib.User) {
	fmt.Println("Seeding.. [creating sample categories and users in memory]")
	seeds.AddMovieCategories(movieCategories)
	seeds.AddUsers(users, movieCategories)
	fmt.Println("Done!")
}
