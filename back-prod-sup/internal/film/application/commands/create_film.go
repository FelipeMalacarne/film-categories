package commands

import (
	"time"

	"github.com/felipemalacarne/back-prod-sup/internal/film/domain"
)

type CreateFilmCommand struct {
	ReleaseDate time.Time `json:"release_date"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Duration    uint16    `json:"duration"`
}

type CreateFilmHandler struct {
	repository domain.FilmRepository
}

func NewCreateFilmHandler(repository domain.FilmRepository) *CreateFilmHandler {
    return &CreateFilmHandler{repository}
}

func (h CreateFilmHandler) Handle(command CreateFilmCommand) (domain.Film, error) {
    film := domain.NewFilm(command.Name, command.Description, command.Duration, command.ReleaseDate, command.Author)
    err := film.Validate()
    if err != nil {
        return domain.Film{}, err
    }
    return h.repository.Create(film)
}
