package handler

import (
	"net/http"

	"github.com/Chandra5468/movie-recommendation-system/internal/service"
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
