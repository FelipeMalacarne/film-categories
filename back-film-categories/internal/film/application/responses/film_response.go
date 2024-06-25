package responses

import (
	"time"

	cd "github.com/felipemalacarne/back-film-categories/internal/category/domain"
	fd "github.com/felipemalacarne/back-film-categories/internal/film/domain"
	"github.com/google/uuid"
)

type FilmResponse struct {
	ReleaseDate time.Time    `json:"release_date"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Category    *cd.Category `json:"category,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Author      string       `json:"author"`
	Duration    uint16       `json:"duration"`
	ID          uuid.UUID    `json:"id"`
}

func NewFilmResponse(film *fd.Film, category *cd.Category) *FilmResponse {
	return &FilmResponse{
		ID:          film.ID,
		Name:        film.Name,
		Description: film.Description,
		ReleaseDate: film.ReleaseDate,
		CreatedAt:   film.CreatedAt,
		UpdatedAt:   film.UpdatedAt,
        Category:    category,
		Author:   film.Author,
		Duration: film.Duration,
	}
}
