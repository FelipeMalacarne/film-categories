package persistence

import (
	"time"

	"github.com/felipemalacarne/back-film-categories/internal/film/domain"
	"github.com/felipemalacarne/back-film-categories/utils"
	"github.com/google/uuid"
)

type dynamoFilm struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CategoryID  *string `json:"category_id,omitempty"`
	Author      string  `json:"author"`
	Duration    uint16  `json:"duration"`
}

func toDynamoFilm(film *domain.Film) *dynamoFilm {
	var categoryID *string
	if film.CategoryID != nil {
		idStr := film.CategoryID.String()
		categoryID = &idStr
	}

	return &dynamoFilm{
		ID:          film.ID.String(),
		Name:        film.Name,
		Description: film.Description,
		ReleaseDate: film.ReleaseDate.Format(time.RFC3339),
		CreatedAt:   film.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   film.UpdatedAt.Format(time.RFC3339),
		Duration:    film.Duration,
		Author:      film.Author,
		CategoryID:  categoryID,
	}
}

func toFilm(df *dynamoFilm) *domain.Film {
	var categoryID *uuid.UUID
    if df.CategoryID != nil {
        id, _ := uuid.Parse(*df.CategoryID)
        categoryID = &id
    }

	return &domain.Film{
		ID:          uuid.MustParse(df.ID),
		Name:        df.Name,
		Description: df.Description,
		ReleaseDate: utils.ParseTime(df.ReleaseDate),
		CreatedAt:   utils.ParseTime(df.CreatedAt),
		UpdatedAt:   utils.ParseTime(df.UpdatedAt),
		Duration:    df.Duration,
		CategoryID:  categoryID,
		Author:      df.Author,
	}
}
