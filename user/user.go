package user

import (
	"../lib"
	"../printer"
	"../seeds"
	"fmt"
	"log"
	"strconv"
)

func Build(movieCategories []lib.MovieCategory, users *[]lib.User) lib.User {
	lastUser := (*users)[len(*users)-1]

	watchedCategoriesMap := getFavouritesCategories()
	currentUser := lib.User{lastUser.Id + 1, "Current User", watchedCategoriesMap}

	return currentUser
}

func getFavouritesCategories() map[int]int {
	var selectedMovies []lib.Movie

	categories := make(map[int]int)
	movies := seeds.Movies()

	printer.SampleMovies(movies)
	askForFavouriteMovies(movies, &selectedMovies)
	fillUserWatchedCategories(&categories, selectedMovies)

	return categories
}

func askForFavouriteMovies(movies []lib.Movie, selectedMovies *[]lib.Movie) {
	fmt.Println("\nSelect 5 movies you like the most.")

	for i := 0; i < 5; i++ {
		fmt.Print("Enter movie's ID: ")
		var input string
		fmt.Scanln(&input)
		selectedMovie := findMovieWithId(input, movies)
		*selectedMovies = append(*selectedMovies, selectedMovie)
	}
}

func fillUserWatchedCategories(categories *map[int]int, selectedMovies []lib.Movie) {

	for _, movie := range selectedMovies {
		(*categories)[movie.CategoryId] = (*categories)[movie.CategoryId] + 1
	}
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
