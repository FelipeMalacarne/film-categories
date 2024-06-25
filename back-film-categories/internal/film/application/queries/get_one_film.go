package queries

import (
	cd "github.com/felipemalacarne/back-film-categories/internal/category/domain"
	"github.com/felipemalacarne/back-film-categories/internal/film/application/responses"
	fd "github.com/felipemalacarne/back-film-categories/internal/film/domain"
	"github.com/google/uuid"
)

type GetOneFilmQuery struct {
	ID uuid.UUID
}

type GetOneFilmHandler struct {
	filmRepository     fd.FilmRepository
	categoryRepository cd.CategoryRepository
}

func NewGetOneFilmHandler(filmRepository fd.FilmRepository, categoryRepository cd.CategoryRepository) *GetOneFilmHandler {
	return &GetOneFilmHandler{filmRepository, categoryRepository}
}

func (h *GetOneFilmHandler) Handle(query GetOneFilmQuery) (responses.FilmResponse, error) {
	film, err := h.filmRepository.FindByID(query.ID)
	if err != nil {
		return responses.FilmResponse{}, err
	}

    if film.CategoryID == nil {
        return *responses.NewFilmResponse(&film, nil), nil
    }

	category, err := h.categoryRepository.FindByID(*film.CategoryID)
    if err != nil {
        return responses.FilmResponse{}, err
    }

	return *responses.NewFilmResponse(&film, &category), nil
}
