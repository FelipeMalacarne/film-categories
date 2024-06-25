package queries

import (
	cd "github.com/felipemalacarne/back-prod-sup/internal/category/domain"
	"github.com/felipemalacarne/back-prod-sup/internal/film/application/responses"
	fd "github.com/felipemalacarne/back-prod-sup/internal/film/domain"
)

type ListFilmsQuery struct{}

type ListFilmsHandler struct {
	filmRepository     fd.FilmRepository
	categoryRepository cd.CategoryRepository
}

func NewListFilmsHandler(filmRepository fd.FilmRepository, categoryRepository cd.CategoryRepository) *ListFilmsHandler {
	return &ListFilmsHandler{filmRepository, categoryRepository}
}

func (h ListFilmsHandler) Handle(query ListFilmsQuery) ([]responses.FilmResponse, error) {
	films, err := h.filmRepository.FindAll()
	if err != nil {
		return []responses.FilmResponse{}, err
	}

	var filmResponses []responses.FilmResponse
	for _, film := range films {
		if film.CategoryID == nil {
			filmResponses = append(filmResponses, *responses.NewFilmResponse(&film, nil))
			continue
		}

		category, err := h.categoryRepository.FindByID(*film.CategoryID)
		if err != nil {
			return []responses.FilmResponse{}, err
		}

		filmResponses = append(filmResponses, *responses.NewFilmResponse(&film, &category))
	}

	return filmResponses, nil
}
