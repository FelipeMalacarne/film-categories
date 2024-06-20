package commands

import (
	"time"

	"github.com/felipemalacarne/back-prod-sup/internal/film/domain"
	"github.com/google/uuid"
)

type UpdateFilmCommand struct {
	ReleaseDate *time.Time `json:"release_date"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Duration    *uint16    `json:"duration"`
	ID          uuid.UUID  `json:"id"`
}

type UpdateFilmHandler struct {
	repository domain.FilmRepository
}

func NewUpdateFilmHandler(repository domain.FilmRepository) *UpdateFilmHandler {
	return &UpdateFilmHandler{repository}
}

func (h UpdateFilmHandler) Handle(command UpdateFilmCommand) (domain.Film, error) {
	film, err := h.repository.FindByID(command.ID)
	if err != nil {
		return domain.Film{}, err
	}

	film.Update(command.Name, command.Description, command.Duration, command.ReleaseDate)

	err = film.Validate()
	if err != nil {
		return domain.Film{}, err
	}
	return h.repository.Update(&film)
}
