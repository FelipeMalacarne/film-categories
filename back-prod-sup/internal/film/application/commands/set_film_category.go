package commands

import (
	"errors"

	cd "github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	fd "github.com/felipemalacarne/back-prod-sup/internal/film/domain"
	"github.com/google/uuid"
)

type SetFilmCategoryCommand struct {
	FilmID     uuid.UUID
	CategoryID uuid.UUID
}

type SetFilmCategoryHandler struct {
	filmRepository     fd.FilmRepository
	categoryRepository cd.CategoryRepository
}

func NewSetFilmCategoryHandler(filmRepository fd.FilmRepository, categoryRepository cd.CategoryRepository) *SetFilmCategoryHandler {
	return &SetFilmCategoryHandler{filmRepository, categoryRepository}
}

func (h SetFilmCategoryHandler) Handle(command SetFilmCategoryCommand) (fd.Film, error) {
	catetory, err := h.categoryRepository.FindByID(command.CategoryID)
	if err != nil {
		return fd.Film{}, err
	}

	if catetory.ID == uuid.Nil {
		return fd.Film{}, errors.New("category not found")
	}

	film, err := h.filmRepository.FindByID(command.FilmID)
	if err != nil {
		return fd.Film{}, err
	}

	film.CategoryID = &catetory.ID

	return h.filmRepository.Update(&film)
}
