package service

import (
	"github.com/Chandra5468/movie-recommendation-system/internal/repository"
	"github.com/Chandra5468/movie-recommendation-system/types"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FetchMovies() ([]types.Movie, error) {
	return s.repo.GetDistinctGenreMovies()
}

// func (s *Service) AddMovie(movie Movie) error {
// 	return s.repo.SaveMovie(movie)
// }

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
