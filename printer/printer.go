package printer

import (
	"../lib"
	"fmt"
	"time"
)

func SimilarUsers(nearestNeighbors map[float64]lib.User) {
	emptyLine()
	for k, v := range nearestNeighbors {
		fmt.Printf("Similar user to current user: Name: %s, Path: %f\n", v.Name, k)
	}
}

func PredictedPreferredCategories(movieCategories []lib.MovieCategory) {
	emptyLine()
	for _, category := range movieCategories {
		fmt.Printf("Predicted recommended category for the currentUser: %v\n", category.Name)
	}
}

func MostSimilarUser(mostSimilarUser lib.User) {
	emptyLine()
	fmt.Printf("Most similar user: %s\n", mostSimilarUser.Name)
}


func SampleMovies(movies []lib.Movie) {
	emptyLine()
	fmt.Print("Here is a list of sample movies:\n")

	for _, movie := range movies {
		fmt.Printf("ID: %v, title: %s\n", movie.Id, movie.Name)
	}
}

func BenchmarkResult(benchmarkStop time.Duration) {
	emptyLine()
	fmt.Printf("Benchmark time: %s\n", benchmarkStop)
}

func emptyLine() {
	fmt.Print("\n")
}
