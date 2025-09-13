package handler

import (
	"github.com/Chandra5468/movie-recommendation-system/types"
)

func createFeatureVector(movie *types.Movie, masterList []string) []int {
	featureVector := make([]int, len(masterList))

	genresMap := convertToMap(movie.Genre)
	actorsMap := convertToMap(movie.Actors)

	for index, feature := range masterList {
		if _, ok := genresMap[feature]; ok {
			featureVector[index] = 1
			continue
		}
		if _, ok := actorsMap[feature]; ok {
			featureVector[index] = 1
			continue
		}
		if feature == movie.Director {
			featureVector[index] = 1
			continue
		}
	}
	return featureVector
}

func convertToMap(items []string) map[string]bool {
	set := make(map[string]bool, len(items))
	for _, item := range items {
		set[item] = true
	}
	return set
}
