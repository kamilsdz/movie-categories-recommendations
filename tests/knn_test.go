package knn

import (
	"../knn"
	"../lib"
	"testing"
)

func buildNearestNeighbors() (lib.User, []lib.User, []lib.MovieCategory, map[float64]lib.User) {
	movieCategories := []lib.MovieCategory{
		lib.MovieCategory{1, "Horror"},
		lib.MovieCategory{2, "Drama"},
		lib.MovieCategory{3, "Action"},
		lib.MovieCategory{4, "Spy movie"},
		lib.MovieCategory{5, "Comedy"},
		lib.MovieCategory{6, "Thriller"},
		lib.MovieCategory{7, "Science Fiction"},
		lib.MovieCategory{8, "Fantasy"},
		lib.MovieCategory{9, "Reality Show"},
		lib.MovieCategory{10, "Documentary"},
	}
	currentUserWatchedMovies := map[int]int{1: 30, 2: 10, 3: 0, 4: 0}
	mostSimilarWatchedMovies := map[int]int{1: 40, 2: 10, 3: 10, 4: 0, 10: 10}
	lessSimilarWatchedMovies := map[int]int{1: 20, 2: 0, 3: 20, 4: 10}
	nonSimilarWatchedMovies := map[int]int{1: 0, 2: 0, 3: 40, 4: 0}
	users := []lib.User{
		lib.User{1, "John Wick", mostSimilarWatchedMovies},
		lib.User{2, "Lara Croft", lessSimilarWatchedMovies},
		lib.User{3, "Indiana Jones", nonSimilarWatchedMovies},
		lib.User{4, "Jan Nowak", currentUserWatchedMovies},
		lib.User{5, "Freddie Mercury", nonSimilarWatchedMovies},
	}
	currentUser := users[3]

	nearestNeighbors := knn.OptimizedFindNearestNeighborsForUser(currentUser, users, movieCategories)
	return currentUser, users, movieCategories, nearestNeighbors
}

func TestNeighborsQuantity(t *testing.T) {
	_, _, _, nearestNeighbors := buildNearestNeighbors()
	if len(nearestNeighbors) != 1 {
		t.Errorf("Incorrect nearest neighbors quantity, got: %d, want: %d.", len(nearestNeighbors), 2)
	}
}

func TestNearestNeighbors(t *testing.T) {
	_, _, _, nearestNeighbors := buildNearestNeighbors()
	for _, v := range nearestNeighbors {
		isCorrect := (v.Name == "Lara Croft" || v.Name == "John Wick")
		if !isCorrect {
			t.Errorf("Incorrect nearest neighbors, got: %v, want: Lara Croft & John Wick.", nearestNeighbors)
		}
	}
}
func TestPredictPreferredCategories(t *testing.T) {
	currentUser, _, movieCategories, nearestNeighbors := buildNearestNeighbors()
	predictedPreferredCategories := knn.PredictPreferredCategories(currentUser, nearestNeighbors, movieCategories)
	expectedPreferredCategories := []lib.MovieCategory{movieCategories[0], movieCategories[1], movieCategories[2], movieCategories[3], movieCategories[9]}
	if !predictedCategoriesAreTheSame(predictedPreferredCategories, expectedPreferredCategories) {
		t.Errorf("Incorrect predicted preferred category, got: %v, want: %v.", predictedPreferredCategories, expectedPreferredCategories)
	}
}

func TestFindMostSimilarUser(t *testing.T) {
	currentUser, users, _, nearestNeighbors := buildNearestNeighbors()
	mostSimilarUser := knn.FindMostSimilarUser(currentUser, nearestNeighbors)
	if mostSimilarUser.Id != users[0].Id {
		t.Errorf("Incorrect most similar user, got: %v, want: %v.", mostSimilarUser, users[0])
	}
}

func predictedCategoriesAreTheSame(predictedPreferredCategories []lib.MovieCategory, expectedPreferredCategories []lib.MovieCategory) bool {
	if len(predictedPreferredCategories) != len(expectedPreferredCategories) {
		return false
	}
	for _, predictedCategory := range predictedPreferredCategories {
		var findStatus bool
		findStatus = false
		for _, expectedCategory := range expectedPreferredCategories {
			if predictedCategory.Id == expectedCategory.Id {
				findStatus = true
			}
		}
		if findStatus == false {
			return false
		}
	}

	return true
}
