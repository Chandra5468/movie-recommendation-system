package service

import (
	"fmt"
	"sync"

	"github.com/Chandra5468/movie-recommendation-system/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FetchMovies(category string, wg *sync.WaitGroup, data chan<- []string) {
	defer wg.Done()
	res, err := s.repo.GetDistinctGenreMovies(category)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", category, err)
		return
	}
	data <- res
}

// func (s *Service) AddMovie(movie Movie) error {
// 	return s.repo.SaveMovie(movie)
// }

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}
