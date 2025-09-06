package repository

import (
	"database/sql"

	"github.com/Chandra5468/movie-recommendation-system/types"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetDistinctGenreMovies() ([]types.Movie, error) {
	rows, err := r.db.Query("SELECT id, title, director, release_year FROM movies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []types.Movie
	for rows.Next() {
		var movie types.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.ReleaseYear); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

// func (r *Repository) SaveMovie(movie Movie) error {
// 	_, err := r.db.Exec("INSERT INTO movies (title, director, release_year) VALUES ($1, $2, $3)", movie.Title, movie.Director, movie.ReleaseYear)
// 	return err
// }
