package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetDistinctGenreMovies(column string) ([]string, error) {
	var query string

	// Use an if-else block to build the correct query
	if column == "director" {
		query = fmt.Sprintf("SELECT DISTINCT %s FROM movies", column)
	} else {
		// This handles arrays like 'genre' and 'actors'
		query = fmt.Sprintf("SELECT DISTINCT UNNEST(%s) FROM movies", column)
	}

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, err
		}
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

// func (r *Repository) SaveMovie(movie Movie) error {
// 	_, err := r.db.Exec("INSERT INTO movies (title, director, release_year) VALUES ($1, $2, $3)", movie.Title, movie.Director, movie.ReleaseYear)
// 	return err
// }
