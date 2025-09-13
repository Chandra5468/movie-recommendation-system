package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Chandra5468/movie-recommendation-system/internal/service"
	"github.com/Chandra5468/movie-recommendation-system/types"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	// movies, err := h.service.FetchMovies()
	// if err != nil {
	//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//     return
	// }
	// c.JSON(http.StatusOK, movies)
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// hand.ServeHTTP(w, r)
	// })
	wg := &sync.WaitGroup{}
	data := make(chan []string, 3)
	list := [3]string{"genre", "actors", "director"}
	for i := range list {
		wg.Add(1)
		go h.service.FetchMovies(list[i], wg, data)
	}
	wg.Wait()
	close(data)
	// result := make(map[string][]string)
	allResults := [][]string{}
	for res := range data { // here reading data from channel
		allResults = append(allResults, res)
	}
	// Combining them into a single slice.
	var MasterList []string
	for _, lst := range allResults {
		MasterList = append(MasterList, lst...)
	}

	movie := &types.Movie{
		Name:     "Inception",
		Genre:    []string{"Action", "Sci-Fi"},
		Actors:   []string{"Leonardo DiCaprio", "Joseph Gordon-Levitt"},
		Budget:   160000000,
		Director: "Christopher Nolan",
	}
	res := createFeatureVector(movie, MasterList)
	// create a new table for vector stores. Refer it to movies
	// Do dot product with other movies and recommend top N movies
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// func (h *Handler) CreateMovie(c *gin.Context) {
// 	var movie service.Movie
// 	if err := c.ShouldBindJSON(&movie); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := h.service.AddMovie(&movie); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, movie)
// }
