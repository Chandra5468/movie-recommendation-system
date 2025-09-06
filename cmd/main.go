package main

import (
	"log"
	"net/http"

	"github.com/Chandra5468/movie-recommendation-system/internal/handler"
	"github.com/Chandra5468/movie-recommendation-system/internal/repository"
	"github.com/Chandra5468/movie-recommendation-system/internal/service"

	"github.com/Chandra5468/movie-recommendation-system/pkg/db"
)

func main() {
	// Initialize database connection
	db, err := db.Connect("postgres://postgres:user@localhost/cfq?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Initialize repository, service, and handler
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Set up HTTP routes
	http.HandleFunc("/movies", h.GetMovies)
	// http.HandleFunc("/movies/create", h.CreateMovie)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
