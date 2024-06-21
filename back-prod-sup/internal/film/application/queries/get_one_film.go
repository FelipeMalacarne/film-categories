package queries

import (
	"github.com/felipemalacarne/back-prod-sup/internal/film/domain"
	"github.com/google/uuid"
)

type GetOneFilmQuery struct {
	ID uuid.UUID
}

type GetOneFilmHandler struct {
	repository domain.FilmRepository
}

func NewGetOneFilmHandler(repository domain.FilmRepository) *GetOneFilmHandler {
	return &GetOneFilmHandler{repository}
}

func (h *GetOneFilmHandler) Handle(query GetOneFilmQuery) (domain.Film, error) {
	return h.repository.FindByID(query.ID)
}
