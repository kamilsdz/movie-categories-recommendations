package seeds

import (
	"../lib"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var cachedCategories = readCategories()

func Movies() []lib.Movie {
	var movies []lib.Movie
	file, err := os.Open(lib.MoviesSeedFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawMovieData := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(rawMovieData[0])
		title := rawMovieData[1]
		categoryId, _ := strconv.Atoi(rawMovieData[2])
		movie := lib.Movie{id, title, categoryId}

		movies = append(movies, movie)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return movies
}

func AddMovieCategories(movieCategories *[]lib.MovieCategory) {
	for _, movieCategory := range cachedCategories {
		*movieCategories = append(*movieCategories, movieCategory)
	}
}

func AddUsers(users *[]lib.User, movieCategories *[]lib.MovieCategory) {
	for i := 1; i <= lib.UsersNumber; i++ {
		watchedMoviesMap := generateWatchedCategoriesForUser()
		user := lib.User{i, fmt.Sprintf("User %d", i), watchedMoviesMap}
		*users = append(*users, user)
	}
}

func generateWatchedCategoriesForUser() map[int]int {
	watchedCategoriesMap := make(map[int]int)
	for i := 1; i <= categoriesSize(); i++ {
		watchedCategoriesMap[i] = randomInt(0, lib.MaxNumberOfWatchedCategories)
	}
	return watchedCategoriesMap
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func readCategories() []lib.MovieCategory {
	var categories []lib.MovieCategory
	file, err := os.Open(lib.MovieCategoriesSeedFile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawCategoryData := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(rawCategoryData[0])
		name := rawCategoryData[1]
		category := lib.MovieCategory{id, name}

		categories = append(categories, category)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return categories
}

func categoriesSize() int {
	return len(cachedCategories)
}
