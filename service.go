package main

import (
	"./knn"
	"./lib"
	"./seeds"
	"fmt"
	"time"
	"strconv"
	"log"
)

func main() {
	var movieCategories []lib.MovieCategory
	var users []lib.User

	addSeeds(&movieCategories, &users)
	currentUser := getCurrentUser(movieCategories, &users)

	benchmarkStart := time.Now()

	nearestNeighbors := knn.OptimizedFindNearestNeighborsForUser(currentUser, users, movieCategories)
	knn.SimilarUsersPrint(nearestNeighbors)

	predictedPreferredCategories := knn.PredictPreferredCategories(currentUser, nearestNeighbors, movieCategories)
	knn.PredictedPreferredCategoriesPrint(predictedPreferredCategories)

	mostSimilarUser := knn.FindMostSimilarUser(currentUser, nearestNeighbors)
	fmt.Printf("Most similar user: %s\n", mostSimilarUser.Name)
	benchmarkStop := time.Since(benchmarkStart)
	fmt.Printf("Benchmark time: %s\n", benchmarkStop)
}


func getCurrentUser(movieCategories []lib.MovieCategory, users *[]lib.User) lib.User {
	lastUser := (*users)[len(*users)-1]

	watchedCategoriesMap := getFavouritesCategories()

	currentUser := lib.User{lastUser.Id+1, "Current User", watchedCategoriesMap}
	*users = append(*users, currentUser)
	return currentUser
}

func getFavouritesCategories() map[int]int {
  categories := make(map[int]int)
  movies := seeds.Movies()
  var selectedMovies []lib.Movie

  printSampleMovies(movies)

  fmt.Println("\nSelect 5 movies you like the most.")

	for i := 0; i < 5; i++ {
	  fmt.Print("Enter movie's ID: ")
	      var input string
    fmt.Scanln(&input)
	  selectedMovie := findMovieWithId(input, movies)
	  selectedMovies = append(selectedMovies, selectedMovie)
	}


	for _, movie := range selectedMovies {
		categories[movie.CategoryId] = categories[movie.CategoryId] + 1
	}


  return categories
}

func findMovieWithId(id string, movies []lib.Movie) lib.Movie {
	var selectedMovie lib.Movie
	convertedId, err := strconv.Atoi(id)

    if err != nil {
        log.Fatal(err)
    }

for _, movie := range movies {
    if movie.Id == convertedId {
        selectedMovie = movie
    }
}

if selectedMovie.Id == 0 {
	log.Fatal("Incorrect ID!")
}

	fmt.Println("Selected:", selectedMovie.Name)
	return selectedMovie
}

func printSampleMovies(movies []lib.Movie) {
	fmt.Print("\n\nHere is a list of sample movies:\n")

	for _, movie := range movies {
		fmt.Printf("ID: %v, title: %s\n", movie.Id, movie.Name)
	}
}

func addSeeds(movieCategories *[]lib.MovieCategory, users *[]lib.User) {
	fmt.Println("Seeding.. [creating sample categories and users in memory]")
	seeds.AddMovieCategories(movieCategories)
	seeds.AddUsers(users, movieCategories)
	fmt.Println("Done!\n")
}
