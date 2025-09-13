package types

import "github.com/google/uuid"

type Movie struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Genre    []string  `json:"genre"`
	Actors   []string  `json:"actors"`
	Budget   int       `json:"budget"`
	Director string    `json:"director"`
}
