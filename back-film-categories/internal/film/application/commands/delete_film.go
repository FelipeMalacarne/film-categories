package commands

import (
	"github.com/felipemalacarne/back-film-categories/internal/film/domain"
	"github.com/google/uuid"
)

type DeleteFilmCommand struct {
	ID uuid.UUID `json:"id"`
}

type DeleteFilmHandler struct {
	repository domain.FilmRepository
}

func NewDeleteFilmHandler(repository domain.FilmRepository) *DeleteFilmHandler {
	return &DeleteFilmHandler{repository}
}

func (h DeleteFilmHandler) Handle(command DeleteFilmCommand) error {
    return h.repository.Delete(command.ID)
}
